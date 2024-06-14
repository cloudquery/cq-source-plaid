package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Version = "development"
	Kind    = "source"
	Name    = "example"
	Team    = "example-team"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"plaid",
		Version,
		Configure,
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
		plugin.WithConnectionTester(TestConnection),
	)
}
