package client

import (
	"github.com/spf13/cobra"
)

// function to create the cli handler
type CLIHandlerFn func() *cobra.Command

// The combined type for a proposal handler for both cli and rest
type ExecutionHandle struct {
	CLIHandler CLIHandlerFn
}

// NewExecutionHandler creates a new ExecutionHandle object
func NewExecutionHandler(cliHandler CLIHandlerFn) ExecutionHandle {
	return ExecutionHandle{
		CLIHandler: cliHandler,
	}
}
