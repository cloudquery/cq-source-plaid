package main

import (
	"github.com/cloudquery/cq-source-plaid/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

func main() {
	serve.Source(plugin.Plugin())
}
