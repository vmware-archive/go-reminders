// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package vro

import (
	"log"
)

type BasicAuth struct {
	user     string `json: "user"`
	passwd   string `json: "password`
	insecure bool   `json: "insecure"`
}

type vROCreds struct {
	host   string `json: "Host"`
	port   string `json: "port"`
	admin  string `json: "Admin"`
	passwd string `json: "Passwd"`
}

func init() {
	log.Println("Initialized package vro.")
}

func New(user string, passwd string, insecure bool) BasicAuth {
	return BasicAuth{
		user:     user,
		passwd:   passwd,
		insecure: insecure,
	}
}
