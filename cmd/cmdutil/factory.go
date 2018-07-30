package cmdutil

import (
	mkr "github.com/i2tsuki/mackerel-client-go"
	app "github.com/i2tsuki/mkr-graph/app"
)

// Factory provides abstractions that allow the mkr-graph command to be extended across multiple types of resources and different API sets.
type Factory struct {
	App app.App
	// Client gives you back an internal, generated client
	Client *mkr.Client
}

func NewFactory(app app.App, client *mkr.Client) Factory {
	f := Factory{
		App:    app,
		Client: client,
	}
	return f

}
