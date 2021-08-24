# Examples

Here's a simple example demoing how a directory tree translates to a map in terraform.

Given an example root:

```
root
├── 1
├── a
│   ├── 2
│   ├── 3
│   └── b
│       ├── 4
│       ├── 5
│       ├── 6
│       └── c
└── apps
    ├── app1
    │   ├── deeper
    │   │   └── data
    │   ├── pipeline1
    │   ├── pipeline2
    │   └── pipeline3
    ├── app2
    │   ├── deeper
    │   │   └── data
    │   ├── pipeline1
    │   ├── pipeline2
    │   └── pipeline3
    ├── app3
    │   ├── deeper
    │   │   └── data
    │   ├── pipeline1
    │   ├── pipeline2
    │   └── pipeline3
    └── config
```

The data provider generates this kind of map for use in terraform:

```
  + root = {
      + 1    = null
      + a    = {
          + 2 = null
          + 3 = null
          + b = {
              + 4 = null
              + 5 = null
              + 6 = null
              + c = {
                  + .gitkeep = null
                }
            }
        }
      + apps = {
          + app1   = {
              + deeper    = {
                  + data = null
                }
              + pipeline1 = null
              + pipeline2 = null
              + pipeline3 = null
            }
          + app2   = {
              + deeper    = {
                  + data = null
                }
              + pipeline1 = null
              + pipeline2 = null
              + pipeline3 = null
            }
          + app3   = {
              + deeper    = {
                  + data = null
                }
              + pipeline1 = null
              + pipeline2 = null
              + pipeline3 = null
            }
          + config = null
        }
    }
```

There is also an output demonstrating how you might pick just the `app*` directories and normal `pipeline*` files into a simple data structure.

## Run

```
terraform init
terraform plan
terraform apply
```
