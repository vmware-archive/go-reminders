// Copyright 2015 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package stats

import (
	"encoding/json"
	"log"
	"sync"
)

// store hits per URL
type Stats struct {
	hits map[string]int `json:"hits"`
	lock sync.RWMutex
}

func init() {
	log.Println("Initialized stats package.")
}

func New() Stats {
	return Stats{
		hits: make(map[string]int),
		lock: sync.RWMutex{},
	}
}

// Convert a JSON string to Go struct and return.
func HitsFromJson(jsonData []byte) (map[string]int, error) {
	var hits map[string]int
	err := json.Unmarshal([]byte(jsonData), &hits)
	if err != nil {
		log.Printf("%T\n%s\n%#v\n", err, err, err)
		switch v := err.(type) {
		case *json.SyntaxError:
			log.Println(string(jsonData[v.Offset-40 : v.Offset]))
		}
	}

	return hits, err
}
