package show

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"

	cmdutil "github.com/i2tsuki/mkr-graph/cmd/cmdutil"
	cmderror "github.com/i2tsuki/mkr-graph/cmd/error"

	asciigraph "github.com/i2tsuki/asciigraph"
	cobra "github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

// RunOptions is subcommand options struct
type RunOptions struct {
	Identifier string
	Height     int
	Point      int
	Offset     int

	Attribute string
}

// NewRunOptions generate options
func NewRunOptions() *RunOptions {
	return &RunOptions{}
}

func addRunFlags(cmd *cobra.Command, opt *RunOptions) {
	cmd.Flags().StringVar(&opt.Identifier, "host", "", "Show host metrics graph")
	cmd.Flags().StringVar(&opt.Identifier, "service", "", "Show service metrics graph")
	cmd.Flags().IntVar(&opt.Height, "height", 12, "Graph height (Default: 12)")
	cmd.Flags().IntVar(&opt.Point, "point", 60, "Number of time point before now (Default: 60)")
	cmd.Flags().IntVar(&opt.Offset, "offset", 3, "X-Axis offset of graph (Default: 3)")

	cmd.Flags().SetNormalizeFunc(
		func(f *flag.FlagSet, name string) flag.NormalizedName {
			return flag.NormalizedName(name)
		},
	)
}

// NewCmd construct show subcommand
func NewCmd(f cmdutil.Factory) *cobra.Command {
	o := NewRunOptions()

	cmd := &cobra.Command{
		Use: `show <--host [HOST ID]> <--service [SERIVCE]>`,
		DisableFlagsInUseLine: true,
		Short:   "Show graph by host, service-role, service",
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
	if o.Point > 300 {
		err := errors.New("point is too much")
		err = cmderror.NewError(f, err)
		return err
	}

	if o.Height > 72 {
		err := errors.New("height is so large")
		err = cmderror.NewError(f, err)
		return err
	}

	o.Attribute = cmdutil.LastMatchArgOptionName(o.Identifier)

	return nil
}

func (o *RunOptions) run(f cmdutil.Factory, cmd *cobra.Command, args []string) error {
	var names []string
	if len(args) > 0 {
		names = args
	} else {
		// Read metric name from os.Stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			names = append(names, scanner.Text())
		}
	}

	now := time.Now().Unix()

	to := now
	from := now - int64(o.Point*60)

	switch o.Attribute {

	case "host":
		for _, name := range names {
			var series []float64

			metrics, err := f.Client.FetchHostMetricValues(o.Identifier, name, from, to)
			if err != nil {
				err := cmderror.NewError(f, err)
				return err
			}
			if len(metrics) == 0 {
				err := errors.New("metrics is empty")
				err = cmderror.NewError(f, err)
				return err
			}

			for _, metric := range metrics {
				switch v := metric.Value.(type) {
				case float32:
					series = append(series, (float64(v)))
				case float64:
					series = append(series, (float64(v)))
				default:
					err := errors.New("metric interface is invalid")
					err = cmderror.NewError(f, err)
					return err
				}
			}

			graph := asciigraph.Plot(
				series,
				asciigraph.Caption(name),
				asciigraph.Height(o.Height),
				asciigraph.Offset(o.Offset),
				asciigraph.Upper(1.0),
				asciigraph.Lower(0.0),
			)
			fmt.Printf("%s\n", graph)
		}

	case "service":
		for _, name := range names {
			var series []float64

			metrics, err := f.Client.FetchServiceMetricValues(o.Identifier, name, from, to)
			if err != nil {
				err := cmderror.NewError(f, err)
				return err
			}
			if len(metrics) == 0 {
				err := errors.New("metrics is empty")
				err = cmderror.NewError(f, err)
				return err
			}

			for _, metric := range metrics {
				switch v := metric.Value.(type) {
				case float32:
					series = append(series, (float64(v)))
				case float64:
					series = append(series, (float64(v)))
				default:
					err := errors.New("metric interface is invalid")
					err = cmderror.NewError(f, err)
					return err
				}
			}

			graph := asciigraph.Plot(
				series,
				asciigraph.Caption(name),
				asciigraph.Height(o.Height),
				asciigraph.Offset(o.Offset),
				asciigraph.Upper(1.0),
				asciigraph.Lower(0.0),
			)
			fmt.Printf("%s\n", graph)
		}
	}

	return nil
}
