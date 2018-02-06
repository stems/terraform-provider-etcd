package etcd

import (
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"      
  "fmt"
//  "github.com/hashicorp/hil/ast"
  "github.com/hashicorp/terraform/helper/schema"
)

func dataSourceJsonKVs() *schema.Resource {
  return &schema.Resource{
    Read: dataSourceContentsRead,

    Schema: map[string]*schema.Schema{
      "json": &schema.Schema{
        Type: schema.TypeString,
        Required: true,
        Description: "json object whose properties will be merged with output",
      },
      "interned": &schema.Schema{
        Type: schema.TypeMap,
        Default: nil,
        Computed: true,
        Description: "json object properties",
      },
    },
  }
}

func dataSourceContentsRead(d *schema.ResourceData, meta interface{}) error {
  json_data := d.Get("json").(string)
  var dat map[string]interface{}

  if err := json.Unmarshal([]byte(json_data), &dat); err != nil {
    return fmt.Errorf("failed to parse incoming json")
  }

//  outputs := make(map[string]ast.Variable)
  outputs := make(map[string]interface{})
  sha := sha256.New()
  for k, v := range dat {
    s, ok := v.(string)
    if !ok {
      return fmt.Errorf("unexpected type for key %q: %T", k, v)
    }
    outputs[k] = s
//    ast.Variable{
//      Value: s,
//      Type: ast.TypeString,
//    }
    sha.Write([]byte(k))
    sha.Write([]byte(s))
  }

  d.Set("interned", outputs)
  d.SetId(hex.EncodeToString(sha.Sum(nil)))
  return nil
}

