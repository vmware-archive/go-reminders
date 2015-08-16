// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package template

import (
	"encoding/json"
	"log"
)

// Print out JSON errors on the log.
func jsonError(err error, jsonData string) {
	if err != nil {
		log.Printf("%T\n%s\n%#v\n", err, err, err)
		switch v := err.(type) {
		case *json.SyntaxError:
			log.Println(string(jsonData[v.Offset-40 : v.Offset]))
		}
	}
}

// Print out generic errors on the log.
func perror(err error) {
	if err != nil {
		panic(err)
	}
}
