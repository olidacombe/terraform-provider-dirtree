package dirtree

import (
	"fmt"
	"os"
	"testing"

	"path/filepath"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func createTestDirStructure(dir string) {
	os.MkdirAll(filepath.Join(dir, "a", "b"), os.ModePerm)
	os.Mkdir(filepath.Join(dir, "a", "c"), os.ModePerm)
	os.Create(filepath.Join(dir, "1"))
	os.Create(filepath.Join(dir, "a", "2"))
	os.Create(filepath.Join(dir, "a", "3"))
	os.Create(filepath.Join(dir, "a", "b", "4"))
}

func TestAcctDirtree_basic(t *testing.T) {
	test_dir := t.TempDir()
	createTestDirStructure(test_dir)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"dirtree": func() (tfprotov6.ProviderServer, error) {
				return tfsdk.NewProtocol6Server(New()), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "dirtree_files" "test" {
						root = "%s"
					}
				`, test_dir),
				Check: resource.TestCheckResourceAttr("data.dirtree_files.test", "tree", "{\"1\":null,\"a\":{\"2\":null,\"3\":null,\"b\":{\"4\":null},\"c\":{}}}"),
			},
		},
	})
}
