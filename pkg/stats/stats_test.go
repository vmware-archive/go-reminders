// Copyright 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
//
package stats

import (
	"encoding/json"
	"testing"
)

const url = "localhost:8080/api/stats"

func testAddHit(t *testing.T) {
	s := New()

	s.AddHit(url)

	if s.hits[url] != 1 {
		t.Fail()
	}
}

func testJson(t *testing.T) {
	hits := make(map[string]int)

	hits["/api/reminders"] = 1
	hits["/stats/hits"] = 1
	jsonData, err := json.Marshal(hits)
	if err != nil {
		t.Errorf("Failed to marshal JSON data: %v\n", err)
		t.Fail()
	}

	if _, err := HitsFromJson(jsonData); err != nil {
		t.Errorf("Error: %v\n", err)
		t.Fail()
	}
}

func TestStats(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping: short mode ignores tests.")
		return
	}

	testAddHit(t)
	testJson(t)
}
