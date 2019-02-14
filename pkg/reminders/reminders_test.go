// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package reminders

import (
	"strconv"
	"testing"
)

const (
	testMsg = "reminder"
)

var g_Reminders Reminders

func testMemDBSave(t *testing.T) {
	c := DBCreds{}
	c.Init("", 0, "", "", "mem", "", "", "")

	// Get a new test environment storage interface
	g_Reminders, err := newTestReminders(c, true)
	if err != nil {
		t.Error("Failed to create new Reminders, cannot continue.")
		t.Fail()
		return
	}

	// add a bunch of reminders
	for i := 1; i <= 6; i++ {
		g, err := newGuid()
		if err != nil {
			t.Errorf("Failed to create new guid for reminder because %v\n", err)
			t.Fail()
			break
		}
		r := Reminder{
			Message: testMsg + strconv.Itoa(i),
			Guid:    g.String(),
		}
		if err := g_Reminders.s.Save(r); err != nil {
			t.Errorf("Failed to add reminder %s because %v\n", r.Message, err)
			t.Fail()
			break
		}
	}
}

func TestMemDB(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping: short mode ignores tests.")
		return
	}

	testMemDBSave(t)
}
