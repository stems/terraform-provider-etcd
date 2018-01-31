package main

import (
	"context"
	"fmt"

	"github.com/coreos/etcd/clientv3"

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
	c := meta.(clientv3.Client)
	kapi := clientv3.NewKV(&c)

	key := d.Get("key").(string)
	value := d.Get("value").(string)

	_, err := kapi.Put(context.Background(), key, value)
	if err != nil {
		return fmt.Errorf("could not set key %s: %s", key, err)
	}

	d.SetId(key)
	if err := d.Set("key", key); err != nil {
		return err
	}
	if err := d.Set("value", value); err != nil {
		return err
	}
	return nil
}

func resourceKeyRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(clientv3.Client)
	kapi := clientv3.NewKV(&c)

	key := string(d.Id())
	response, err := kapi.Get(context.Background(), key)
	if err != nil {
		return fmt.Errorf("could not read key %s: %s", key, err)
	}

	if response.Count <= 0 {
		return fmt.Errorf("response was empty for key: %s", key)
	}
	for _, ev := range response.Kvs {
		newKey := string(ev.Key)
		newValue := string(ev.Value)
		if newKey == key {
			if err := d.Set("key", newKey); err != nil {
				return err
			}
			if err := d.Set("value", newValue); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("No value was found for key: %s", key)
}

func resourceKeyDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(clientv3.Client)
	kapi := clientv3.NewKV(&c)

	_, err := kapi.Delete(context.Background(), d.Id())
	if err != nil {
		return fmt.Errorf("could not delete key %s: %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}
