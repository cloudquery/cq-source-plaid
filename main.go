package main

import (
	"github.com/cloudquery/cq-source-plaid/plugin"
	"github.com/cloudquery/plugin-sdk/v3/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
