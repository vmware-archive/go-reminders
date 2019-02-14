// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package main

import (
	"os"

	"github.com/vmware/go-reminders/internal/app"
)

func main() {
	// Delegate to realMain so defered operations can happen (os.Exit exits
	// the program without servicing defer statements)
	os.Exit(app.RealMain())
}
