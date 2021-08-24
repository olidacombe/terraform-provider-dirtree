terraform {
  required_providers {
    dirtree = {
      source = "github.com/olidacombe/dirtree"
    }
  }
}

provider "dirtree" {}

data "dirtree_files" "example" {
  root = "${path.module}/root"
}

output "full_tree" {
  value = jsondecode(data.dirtree_files.example.tree)
}

data "dirtree_files" "apps" {
  root = "${path.module}/root/apps"
}

# Demo how we can filter for dirs and files, and disregard bits of the tree
# (e.g. the deeper/data files under each app)
# The tree consists of maps where keys are names, and values are:
#   + null for regular files
#   + maps for directories (wherein the definition repeats)
output "apps" {
  # apps are defined by all top-level directories in some root "apps"
  value = { for app, pipelines in jsondecode(data.dirtree_files.apps.tree) : app => [
    # anything in an app directory that's not itself a directory is a pipeline definition
    # in this imaginary codebase
    for pipeline, file in pipelines : pipeline if file == null
  ] if pipelines != null }
}
