{ pkgs, overlay, ... }:
pkgs.nixosTest rec {
  name = "service-test";
  enableOCR = true;
  globalTimeout = 60 * 60 * 6; #  hours (statesync is very slow until we migrate to iavlv1)

  nodes.machine = { config, ... }: {
    virtualisation = {
      memorySize = 16 * 1024;
      diskSize = 50 * 1024;
      cores = 8;

      # https://wiki.qemu.org/Documentation/9psetup#Performance_Considerations
      # from console output:
      # 9pnet: Limiting 'msize' to 512000 as this is the maximum supported by transport virtio
      # so setting it to maximum allowed
      msize = 512000;

      graphics = false;

      forwardPorts = [
        # expose tendermint p2p on host to make it discoverable
        {
          from = "host";
          proto = "tcp";

          guest.port = 26656;
          host.port = 26656;
        }

        # ssh into the test machine to debug it
        {
          from = "host";
          proto = "tcp";

          guest.port = 22;
          host.port = 2222;
          host.address = "127.0.0.1";
        }
      ];
    };

    imports = [
      ../nixos-module
      ./configuration.nix
    ];

    nixpkgs.overlays = [ overlay ];

    services.openssh.enable = true;
    users.users.neuraer = {
      isNormalUser = true;
      extraGroups = [ "wheel" ];

      # generated by mkpasswd
      # password is the same as user name: neuraer
      initialHashedPassword = "$y$j9T$AuOCK4gHHdQJu25LsNr2T1$nGOcshPcyIw49crYzfkt647LPS43ebinUP3WLSnzku/";
    };
    users.groups.neuraer = { };

    environment.systemPackages = with pkgs;
      [
        curl
        dig.host
        dig.dnsutils
        inetutils
        htop
      ];
  };

  # https://nix.dev/tutorials/nixos/integration-testing-using-virtual-machines.html
  testScript = ''
    machine.start()

    machine.wait_for_file("/var/lib/neurad/.neurad/.bootstrapped")
    
    machine.wait_for_open_port(26656)

    timeout = 60 * 5 # 5 minutes
    text = "Applied snapshot chunk"

    machine.wait_until_succeeds(f"journalctl -u neurad.service --since -1m --grep='{text}'", timeout= timeout)
    
    timeout = ${toString globalTimeout}
    text = "commit synced commit"

    machine.wait_until_succeeds(f"journalctl -u neurad.service --since -1m --grep='{text}'", timeout = timeout)
  '';
}
