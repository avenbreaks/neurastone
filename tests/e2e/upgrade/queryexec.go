package upgrade

import (
	"fmt"
)

// CreateModuleQueryExec creates a neura module query
func (m *Manager) CreateModuleQueryExec(moduleName, subCommand, chainID string) (string, error) {
	cmd := []string{
		"neurd",
		"q",
		moduleName,
		subCommand,
		fmt.Sprintf("--chain-id=%s", chainID),
		"--keyring-backend=test",
		"--log_format=json",
	}
	return m.CreateExec(cmd, m.ContainerID())
}
