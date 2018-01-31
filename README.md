# terraform-provider-etcd

A plugin for Terraform enabling it to manipulate Etcd keys.


## Installation

  1. 
  2. `go get github.com/stems/terraform-provider-etcd`
  3. `cd $GOPATH/src/github.com/stems/terraform-provider-etcd`
  4. `go install` - which will fail with errors, but not before doing some things we need it to do
  5. You should now have both the `coreos/etcd` and `hashicorp/terraform` codebases in your $GOPATH
  6. Switch to the `v0.10.8` tag of `hashicorp/terraform`
  7. Delete vendor/google.golang.org/grpc from both of those repos - they conflict with each other.
  8. `go get google.golang.org/grpc`
  9. Now you can finally build both terraform and the terraform-provider-etcd plugin.
  10. In github.com/hashicorp/terraform, run `make dev` to build a binary for local development. `make bin` builds for about 7 different platforms and will take much longer.
  11. In github.com/stems/terraform-provider-etcd, execute `go install` which should now complete without errors.  
  12. I think both binaries end up in $GOPATH/bin.  But terraform wherever your old terraform is - probably /usr/local/bin and put the plugin in $HOME/.terraform.d/plugins/ (which you will probably have to create).  You should now be able to use terraform with templates that include the etcd provider.

## API stability

Both Terraform and Etcd are 0.x projects. Expect incompatible changes.

  [1]: https://terraform.io/
