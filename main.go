package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"terraform-multipass-provider/multipass"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	startPlugin(debugMode)
}

func startPlugin(debugMode bool) {
	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return multipass.New()
		},
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "roberts126/tf-provider/multipass", opts)
		if err != nil {
			log.Fatal(err.Error())
		}

		return
	}

	plugin.Serve(opts)
}
