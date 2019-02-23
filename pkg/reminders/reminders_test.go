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

var gReminders Reminders

func testMemDBSave(t *testing.T) {
	c := DBCreds{}
	c.Init("", 0, "", "", "mem", "", "", "")

	// Get a new test environment storage interface
	gReminders, err := newTestReminders(c, true)
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
			GUID:    g.String(),
		}
		if err := gReminders.s.Save(r); err != nil {
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

func TestMemDBDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping: short mode ignores tests.")
		return
	}

	c := DBCreds{}
	c.Init("", 0, "", "", "mem", "", "", "")

	// Get a new test environment storage interface
	gReminders, err := newTestReminders(c, true)
	if err != nil {
		t.Fatal("Failed to create new Reminders, cannot continue.")
	}

	// add a reminder of reminders
	g, err := newGuid()
	if err != nil {
		t.Fatalf("Failed to create new guid for reminder because %v\n", err)
	}
	r := Reminder{
		Message: testMsg + strconv.Itoa(1),
		GUID:    g.String(),
	}
	if err := gReminders.s.Save(r); err != nil {
		t.Fatalf("Failed to add reminder %s because %v\n", r.Message, err)
	}

	reminderList, _ := gReminders.s.GetAll()
	countBeforeDelete := len(*reminderList)
	gReminders.s.DeleteGUID(g.String())
	reminderList, _ = gReminders.s.GetAll()
	countAfterDelete := len(*reminderList)
	if (countBeforeDelete - 1) != countAfterDelete {
		t.Logf("The count before was %d\n", countBeforeDelete)
		t.Logf("The count after was %d\n", countAfterDelete)
		t.Fatalf("Delete of the reminder failed")
	}

}
