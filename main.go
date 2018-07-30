package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"

	app "github.com/i2tsuki/mkr-graph/app"
	config "github.com/i2tsuki/mkr-graph/config"

	cmdutil "github.com/i2tsuki/mkr-graph/cmd/cmdutil"
	describe "github.com/i2tsuki/mkr-graph/cmd/describe"
	show "github.com/i2tsuki/mkr-graph/cmd/show"

	// ToDo: submit pull request to mackerelio/mackerel-client-go
	mkr "github.com/i2tsuki/mackerel-client-go"
	cobra "github.com/spf13/cobra"
)

func main() {
	app := app.NewApp()

	apiKey, err := config.LoadApikey()
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf(app.Msg, app.Name, app.Version))
		fmt.Println(app.Msg, err)
		os.Exit(1)
	}

	client := mkr.NewClient(apiKey)

	f := cmdutil.NewFactory(app, client)

	cmd := &cobra.Command{
		Use:   app.Name,
		Short: app.Short,
		Long:  app.Long,
	}
	cmd.AddCommand(describe.NewCmd(f))
	cmd.AddCommand(show.NewCmd(f))

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
