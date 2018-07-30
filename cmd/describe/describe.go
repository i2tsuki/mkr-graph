package describe

import (
	"fmt"

	cmdutil "github.com/i2tsuki/mkr-graph/cmd/cmdutil"
	cmderror "github.com/i2tsuki/mkr-graph/cmd/error"

	cobra "github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

// RunOptions is subcommand options struct
type RunOptions struct {
	Identifier string

	Attribute string
}

// NewRunOptions generate options
func NewRunOptions() *RunOptions {
	return &RunOptions{}
}

func addRunFlags(cmd *cobra.Command, opt *RunOptions) {
	cmd.Flags().StringVar(&opt.Identifier, "host", "", "Describe host metrics name")
	cmd.Flags().StringVar(&opt.Identifier, "service", "", "Describe service metrics name")

	cmd.Flags().SetNormalizeFunc(
		func(f *flag.FlagSet, name string) flag.NormalizedName {
			return flag.NormalizedName(name)
		},
	)
}

// NewCmd construct describe subcommand
func NewCmd(f cmdutil.Factory) *cobra.Command {
	o := NewRunOptions()

	cmd := &cobra.Command{
		Use: `describe <--host [HOST ID]> <--service [SERIVCE]>`,
		DisableFlagsInUseLine: true,
		Short:   "Describe metrics by host, service",
		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			o.complete(f, cmd)
			o.run(f, cmd, args)
		},
	}
	addRunFlags(cmd, o)

	return cmd
}

func (o *RunOptions) complete(f cmdutil.Factory, cmd *cobra.Command) error {
	o.Attribute = cmdutil.LastMatchArgOptionName(o.Identifier)

	return nil
}

func (o *RunOptions) run(f cmdutil.Factory, cmd *cobra.Command, args []string) error {
	var names []string
	var err error

	switch o.Attribute {
	case "host":
		names, err = f.Client.FetchHostMetricNames(o.Identifier)
		if err != nil {
			err := cmderror.NewError(f, err)
			fmt.Println(f.App.Msg, err)
		}
	case "service":
		names, err = f.Client.FetchServiceMetricNames(o.Identifier)
		if err != nil {
			err := cmderror.NewError(f, err)
			fmt.Println(f.App.Msg, err)
		}
	}

	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}
