package dirtree

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"crypto/sha256"

	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

type dataSourceDirType struct{}

func (r dataSourceDirType) GetSchema(_ context.Context) (schema.Schema, []*tfprotov6.Diagnostic) {
	return schema.Schema{
		Description: "Local directory tree data source",
		Attributes: map[string]schema.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"tree": {
				Description: "Tree of tile names, represented as single json string",
				Type:        types.StringType,
				Computed:    true,
			},
			"root": {
				Description: "Path to the root directory from which you want to retrieve the tree",
				Type:        types.StringType,
				Required:    true,
			},
		},
	}, nil
}

func (r dataSourceDirType) NewDataSource(ctx context.Context, p tfsdk.Provider) (tfsdk.DataSource, []*tfprotov6.Diagnostic) {
	return dataSourceDir{
		p: *(p.(*provider)),
	}, nil
}

type dataSourceDir struct {
	p provider
}

type dataSourceDirConfig struct {
	Id   types.String `tfsdk:"id"`
	Root string       `tfsdk:"root"`
	Tree types.String `tfsdk:"tree"`
}

type fsDir map[string]*fsDir

func newFsDirFromPath(path string) *fsDir {
	r := fsDir{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}

	for _, file := range files {
		name := file.Name()
		r[name] = newFsDirFromPath(filepath.Join(path, name))
	}

	return &r
}

func typesStringFromString(s string) types.String {
	return types.String{
		Unknown: false,
		Null:    false,
		Value:   s,
	}
}

func (r dataSourceDir) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	// Declare struct that this function will set to this data source's state
	var resourceState struct {
		Id   types.String `tfsdk:"id"`
		Tree types.String `tfsdk:"tree"`
		Root types.String `tfsdk:"root"`
	}

	var config dataSourceDirConfig
	err := req.Config.Get(ctx, &config)
	if err != nil {
		// TODO
		panic(err)
	}

	resourceState.Root = typesStringFromString(config.Root)

	walked := newFsDirFromPath(config.Root)
	tree, err := json.Marshal(walked)
	if err != nil {
		tree = []byte("{}")
	}

	resourceState.Tree = typesStringFromString(string(tree))

	id := sha256.Sum256(tree)
	resourceState.Id = typesStringFromString(string(id[:]))

	// Set state
	err = resp.State.Set(ctx, &resourceState)
	if err != nil {
		resp.Diagnostics = append(resp.Diagnostics, &tfprotov6.Diagnostic{
			Severity: tfprotov6.DiagnosticSeverityError,
			Summary:  "Error traversing root",
			Detail:   fmt.Sprintf("An unexpected error was encountered while setting the tree state blob %+v", err.Error()),
		})
		return
	}
}
