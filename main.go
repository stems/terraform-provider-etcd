package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/stems/terraform-provider-etcd/etcd"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: etcd.Provider,
	})
}
