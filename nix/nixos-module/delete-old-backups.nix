{ pkgs, lib, config, ... }:
let cfg = config.services.neurad-supervised; in
{
  systemd.timers = lib.mkIf (cfg.enable && cfg.deleteOldBackups > 0)
    {
      neurad-delete-old-backups = {
        timerConfig = {
          OnCalendar = "hourly";
          # OnCalendar = "*:0/1";
          Persistent = true;
          Unit = "neurad-delete-old-backups.service";
        };

        wantedBy = [ "timers.target" ];
      };
    };

  systemd.services.neurad-delete-old-backups = lib.mkIf (cfg.enable && cfg.deleteOldBackups > 0)
    {
      serviceConfig =
        let
          deleteOldBackups = pkgs.writeShellApplication {
            name = "neurad-delete-old-backups";
            # buildInputs = [ pkgs.coreutils ];
            text = ''
              set -x
              find ${config.users.users.neurad.home}/.neurad -type d -name "data-backup-*" -mtime +${builtins.toString cfg.deleteOldBackups} -exec rm -rf {} +
            '';
          };
        in
        {
          User = config.users.users.neurad.name;
          Type = "oneshot";
          ExecStart = ''
            ${deleteOldBackups}/bin/neurad-delete-old-backups
          '';
        };
    };
}
