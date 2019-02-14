// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package main

import (
	"testing"
)

func TestGoReminders(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping: short mode ignores tests.")
		return
	}

	t.Log("This is the app -- testing occurs externally for now.")
}
