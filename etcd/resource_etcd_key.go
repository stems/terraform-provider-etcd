package etcd

import (
	"context"
	"fmt"

	"github.com/coreos/etcd/client"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKey() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},

		Create: resourceKeySet,
		Read:   resourceKeyRead,
		Update: resourceKeySet,
		Delete: resourceKeyDelete,
	}
}

func resourceKeySet(d *schema.ResourceData, meta interface{}) error {
	c := meta.(client.Client)
        kapi := client.NewKeysAPI(c)

	key := d.Get("key").(string)
	value := d.Get("value").(string)

	_, err := kapi.Set(context.Background(), key, value, nil)
	if err != nil {
		return fmt.Errorf("could not set key %s: %s", key, err)
	}

        // fetch the value back - this normalizes the key with leading slash
	response, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		return fmt.Errorf("could not read key %s: %s", key, err)
	}

        if err := d.Set("key", response.Node.Key); err != nil {
                return err
        }
        if err := d.Set("value", response.Node.Value); err != nil {
                return err
        }

	d.SetId(response.Node.Key)
	if err := d.Set("key", response.Node.Key); err != nil {
		return err
	}
	if err := d.Set("value", response.Node.Value); err != nil {
		return err
	}
	return nil
}

func resourceKeyRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(client.Client)
        kapi := client.NewKeysAPI(c)

	key := string(d.Id())
	response, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
                d.SetId("")
                return nil
		//return fmt.Errorf("could not read key %s: %s", key, err)
	}

        if err := d.Set("key", response.Node.Key); err != nil {
                return err
        }
        if err := d.Set("value", response.Node.Value); err != nil {
                return err
        }
        return nil
}

func resourceKeyDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(client.Client)
        kapi := client.NewKeysAPI(c)

	_, err := kapi.Delete(context.Background(), d.Id(), nil)
	if err != nil {
		return fmt.Errorf("could not delete key %s: %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}
