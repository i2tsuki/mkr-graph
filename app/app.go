package app

import (
	"fmt"
)

// Name is mkr-graph application name
const Name = "mkr-graph"

// Name is mkr-graph version
const Version = "0.1.0"

// App include application name and version
type App struct {
	Name    string
	Version string
	Short   string
	Long    string
	Msg     string
}

// NewApp create mkr-graph app struct
func NewApp() App {
	app := App{
		Name:    Name,
		Version: Version,
		Short:   "mkr-graph is show metric graph in terminal by mackerel.io",
		Long: `mkr-graph:

Describe metric name and show graph by host, role, service hosting by mackerel.io.
This code is hosted by https://github.com/i2tsuki/mkr-graph.`,
		Msg: fmt.Sprintf("%v-%v failed: ", Name, Version),
	}

	return app
}
