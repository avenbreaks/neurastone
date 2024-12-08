{ pkgs, lib, config, ... }:
let
  cfg = config.services.neurad-supervised;

  defaultCfg = pkgs.callPackage ./default-config.nix {
    neuradPackage = cfg.initialPackage;
  };

  neuradUserName = "neurad";
  neuradGroupName = neuradUserName;

  defaultCfgConfigToml = lib.recursiveUpdate (lib.importTOML "${defaultCfg}/config.toml") {
    instrumentation.prometheus = true;
    chain-id = "neura_11235-1";
    p2p.seeds = "c45991e0098b9cacb8603caf4e1cdb7e6e5f87c0@eu.seed.neura.network:26656,e37cb47590ba46b503269ef255873e9698244d8b@us.seed.neura.network:26656,c593e93e1fb8be8b48d4e7bab514a227aa620bf8@as.seed.neura.network:26656";
  };
  defaultCfgAppToml = lib.recursiveUpdate (lib.importTOML "${defaultCfg}/app.toml")
    {
      telemetry = {
        enabled = true;
      };
    };
in
{
  imports = [
    ./grafana-agent.nix
    # ./nginx.nix
  ];

  options.services.neurad-supervised =
    {
      enable = lib.mkOption {
        type = lib.types.bool;
        default = false;
      };

      deleteOldBackups = lib.mkOption {
        type = lib.types.int;
        default = 7;
      };

      initialPackage = lib.mkOption {
        type = lib.types.package;
        default = pkgs.neura;
      };

      config = lib.mkOption {
        type = lib.types.attrs;
        default = { };
      };

      userHome = lib.mkOption {
        type = lib.types.str;
        default = "/var/lib/neurad";
      };

      app = lib.mkOption
        {
          type = lib.types.attrs;
          default = { };
        };

      grafana = {
        enable = lib.mkOption {
          type = lib.types.bool;
          default = true;
        };

        package = lib.mkOption {
          type = lib.types.package;
          default = pkgs.grafana-agent-unstable;
        };

        instance = lib.mkOption {
          type = lib.types.str;
          default = "neurad";
        };

        metricsUrl = lib.mkOption {
          type = lib.types.str;
        };

        logsUrl = lib.mkOption {
          type = lib.types.str;
        };

        secretKeyPath = lib.mkOption {
          type = lib.types.path;
        };
      };
    };

  config = lib.mkIf cfg.enable {
    # to support launching of binaries downloaded by cosmovisor from github releases
    programs.nix-ld.enable = true;

    users.users.${neuradUserName} =
      {
        isSystemUser = true;
        home = cfg.userHome;
        createHome = true;
        group = neuradGroupName;
      };

    users.groups.${neuradGroupName} =
      { };


    systemd.services =
      {
        neurad-bootstrap = {
          serviceConfig =
            let
              neurad-init = pkgs.writeShellApplication {
                name = "neurad-bootstrap";
                runtimeInputs = with pkgs; [
                  cfg.initialPackage
                  curl
                  gnused
                  coreutils
                  which

                  dig.dnsutils
                ];

                text = builtins.readFile ./neurad-bootstrap.sh;
              };
            in
            {
              User = neuradUserName;
              Type = "oneshot";
              ExecStart = ''
                ${neurad-init}/bin/neurad-bootstrap
              '';
              LimitNOFILE = "infinity";
            };

          environment = {
            NIX_LD = config.environment.variables.NIX_LD;
          };

          before = [ "neurad.service" ];
          after = [
            "network-online.target"
            "nss-lookup.target"
          ];
        };

        neurad =
          {
            serviceConfig =
              let
                format = pkgs.formats.toml { };
                tomlConfig = format.generate "config.toml" (lib.attrsets.recursiveUpdate defaultCfgConfigToml cfg.config);
                appConfig = format.generate "app.toml" (lib.attrsets.recursiveUpdate defaultCfgAppToml cfg.app);
                start = pkgs.writeShellApplication {
                  name = "neurad-start";
                  text = ''
                    ln -snf ${tomlConfig} .neurad/config/config.toml
                    ln -snf ${appConfig} .neurad/config/app.toml
                    ${pkgs.cosmovisor}/bin/cosmovisor run start
                  '';
                };
              in
              {
                User = neuradUserName;
                ExecStart = ''${start}/bin/neurad-start'';
                WorkingDirectory = cfg.userHome;
                Restart = "on-failure";
                RestartSec = 30;
                LimitNOFILE = "infinity";
              };

            environment = {
              DAEMON_HOME = "${cfg.userHome}/.neurad";
              DAEMON_NAME = "neurad";
              DAEMON_ALLOW_DOWNLOAD_BINARIES = "true";
              UNSAFE_SKIP_BACKUP = "false";

              NIX_LD = config.environment.variables.NIX_LD;
            };

            wantedBy = [ "multi-user.target" ];
            requires = [ "neurad-bootstrap.service" ];
          };
      };
  };
}
