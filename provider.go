package main

import (
	"github.com/coreos/etcd/clientv3"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns the ResourceProvider implemented by this package. Serve
// this with the Terraform plugin helper and you are golden.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoints": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
                        "username": &schema.Schema{
                                Type: schema.TypeString,
                                Required: true,
                        },
                        "password": &schema.Schema{
                                Type: schema.TypeString,
                                Required: true,
                        },
		},
		ResourcesMap: map[string]*schema.Resource{
			"etcd_key": resourceKey(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	ifEndpoints := d.Get("endpoints").([]interface{})
	strEndpoints := make([]string, len(ifEndpoints))
	for i, v := range ifEndpoints {
		strEndpoints[i] = v.(string)
	}

	username := d.Get("username").(string)
	password := d.Get("password").(string)
        cfg := clientv3.Config{
                Endpoints: strEndpoints,
                Username: username,
                Password: password,
        }
	c, err := clientv3.New(cfg)
	return *c, err
}
