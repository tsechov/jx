package cmd

import (
	"github.com/spf13/cobra"
)

// StepEnvOptions contains the command line flags
type StepEnvOptions struct {
	StepOptions
}

// NewCmdStepEnv Steps a command object for the "step" command
func NewCmdStepEnv(commonOpts *CommonOptions) *cobra.Command {
	options := &StepEnvOptions{
		StepOptions: StepOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:   "env",
		Short: "env [command]",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			CheckErr(err)
		},
	}
	cmd.AddCommand(NewCmdStepEnvApply(commonOpts))
	return cmd
}

// Run implements this command
func (o *StepEnvOptions) Run() error {
	return o.Cmd.Help()
}
