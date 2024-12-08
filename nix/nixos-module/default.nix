{ ... }:
{
  imports = [
    ./neurad-service.nix
    ./grafana-agent.nix
    ./delete-old-backups.nix
  ];
}
