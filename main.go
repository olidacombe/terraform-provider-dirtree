package main

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

import (
	"context"
	"github.com/olidacombe/terraform-provider-dirtree/dirtree"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func main() {
	tfsdk.Serve(context.Background(), dirtree.New, tfsdk.ServeOpts{
		Name: "dirtree",
	})
}
