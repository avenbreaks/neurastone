{ pkgs, neuradPackage }:
let
  inherit (pkgs) stdenv callPackage;
in
stdenv.mkDerivation {
  name = "neurad-default-config";
  version = neuradPackage.version;

  buildInputs = [ neuradPackage ];

  dontUnpack = true;

  CHAIN_ID = "neurad_11235-1";

  buildPhase = ''
    neurad init "neurad-node" --chain-id $CHAIN_ID --home .

    mkdir $out
    mv config/app.toml $out
    mv config/config.toml $out
  '';
}
