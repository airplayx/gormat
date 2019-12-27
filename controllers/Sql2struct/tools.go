package Sql2struct

import (
	"bytes"
	"encoding/json"
	"log"
)

func InStringSlice(f string, a []string) bool {
	for _, s := range a {
		if f == s {
			return true
		}
	}
	return false
}

func JSONMethod(content interface{}) map[string]string {
	var name map[string]string
	if marshalContent, err := json.Marshal(content); err != nil {
		log.Println(err.Error())
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber()
		if err := d.Decode(&name); err != nil {
			log.Println(err.Error())
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}
