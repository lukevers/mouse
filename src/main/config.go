package main

import (
	"github.com/spf13/viper"
	"logger/stderr"
	"reflect"
)

type Config struct {
	Servers map[string]Server
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/mouse/")
	viper.AddConfigPath("$HOME/.mouse")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		stderr.Fatalf("Could not read config file:", err)
	}

	// Temporary(?) fix for HCL problems.
	//
	// As of me writing this, the problem is that currently the HCL parser adds
	// extra arrays for some weird reason.
	//
	// This:
	//
	// servers "fc00" {
	//   ...
	// }
	//
	// Is converted to this [[map:[]]] when it should be converted to [map:[]]
	// on every map.
	//
	// @see https://github.com/hashicorp/hcl/pull/24#issuecomment-69821965
	servers := viper.Get("servers")
	if reflect.ValueOf(servers).Kind() == reflect.Slice {
		servers = servers.([]map[string]interface{})[0]
		for key, s := range servers.(map[string]interface{}) {
			if reflect.ValueOf(s).Kind() == reflect.Slice {
				servers.(map[string]interface{})[key] = s.([]map[string]interface{})[0]
				for k, v := range servers.(map[string]interface{})[key].(map[string]interface{}) {
					if reflect.ValueOf(v).Kind() == reflect.Slice {
						switch v.(type) {
						case []interface{}:
							continue
						case []map[string]interface{}:
							for p, b := range v.([]map[string]interface{})[0] {
								v.([]map[string]interface{})[0][p] = b.([]map[string]interface{})[0]
							}

							servers.(map[string]interface{})[key].(map[string]interface{})[k] = v.([]map[string]interface{})[0]
						}
					}
				}
			}
		}

		viper.Set("servers", servers)
	}

	if err := viper.Unmarshal(&config); err != nil {
		stderr.Fatalf("Could not unmarshal config:", err)
	}
}
