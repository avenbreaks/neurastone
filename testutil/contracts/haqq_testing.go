package contracts

import (
	_ "embed"
	"encoding/json"

	evmtypes "github.com/avenbreaks/neurastone/x/evm/types"
)

var (
	//go:embed compiled_contracts/neura_testing.json
	neuraTestingJSON []byte // nolint: golint

	// neuraTestingContract is the compiled dummy contract
	neuraTestingContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(neuraTestingJSON, &neuraTestingContract)
	if err != nil {
		panic(err)
	}

	if len(neuraTestingContract.Bin) == 0 {
		panic("load contract failed")
	}
}
