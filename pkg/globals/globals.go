// Copyright 2015 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package globals

import (
	"log"

	"github.com/vmwaresamples/go-reminders/pkg/stats"
)

// Global application context variables.
var (
	ListenPort  int
	Port        int
	Host        string
	Admin       string
	Passwd      string
	DBName      string
	DBType      string
	ContentRoot string
	APIAddress  string
	CfgType     string
	CfgSrc      string
	Insecure    bool
	Stats       stats.Stats = stats.New()
)

func init() {
	log.Printf("Initialized globals package.")
}
