//go:build ignore

package main_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/stretchr/testify/require"

	"github.com/avenbreaks/neurastone/app"
	neurd "github.com/avenbreaks/neurastone/cmd/neurd"
	"github.com/avenbreaks/neurastone/utils"
)

func TestInitCmd(t *testing.T) {
	rootCmd, _ := neurd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"init",      // Test the init cmd
		"haqq-test", // Moniker
		fmt.Sprintf("--%s=%s", cli.FlagOverwrite, "true"), // Overwrite genesis.json, in case it already exists
		fmt.Sprintf("--%s=%s", flags.FlagChainID, utils.TestEdge2ChainID+"-3"),
	})

	err := svrcmd.Execute(rootCmd, "neurd", app.DefaultNodeHome)
	require.NoError(t, err)
}

func TestAddKeyLedgerCmd(t *testing.T) {
	rootCmd, _ := neurd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"keys",
		"add",
		"dev0",
		fmt.Sprintf("--%s", flags.FlagUseLedger),
	})

	err := svrcmd.Execute(rootCmd, "neurd", app.DefaultNodeHome)
	require.Error(t, err)
}
