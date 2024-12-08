set -euxo pipefail
echo "Bootstrapping"

neuraD_DIR="$HOME"/.neurad
BIN_DIR="$neuraD_DIR"/cosmovisor/genesis/bin

if [ -f "$neuraD_DIR"/.bootstrapped ]; then
    echo "neurad already bootstrapped"
    exit 0
fi
echo "Bootstrapping ~/.neurad"
# cp -r ${neuradBinary}/share/neurad/init "$neuraD_DIR"
# chmod -R 0770 "$neuraD_DIR"
id

mkdir -p "$BIN_DIR"

tmpDir=$(mktemp -d -p /tmp)
cd "$tmpDir"

neurad_path=$(realpath "$(which neurad)")
echo "$neurad_path"

cp "$neurad_path" "$BIN_DIR"
chmod -R 0771 "$BIN_DIR"/neurad

export PATH="$BIN_DIR":$PATH
neurad config chain-id neura_11235-1
neurad init "neura-node" --chain-id neura_11235-1

curl -L https://raw.githubusercontent.com/neura-network/mainnet/master/genesis.json -o "$neuraD_DIR"/config/genesis.json
# curl -L https://raw.githubusercontent.com/neura-network/mainnet/master/addrbook.json -o "$neuraD_DIR"/config/addrbook.json

touch "$neuraD_DIR"/.bootstrapped
