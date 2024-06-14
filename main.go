package main

import (
	"context"
	"log"

	"github.com/cloudquery/cq-source-plaid/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	p := serve.Plugin(plugin.Plugin())

	if err := p.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
