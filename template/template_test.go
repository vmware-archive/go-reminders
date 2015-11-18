// Copyright (c) 2015 VMware
// Author: Tom Hite (thite@vmware.com)
//
// License: MIT (see https://github.com/tdhite/go-reminders/LICENSE).
//
package template

import (
	"testing"
)

func TestGenerateAPIUrl(t *testing.T) {
	tmpl := New("/", "www.website.com/")

	expected := "http://www.website.com/testing"
	actual := tmpl.generateAPIUrl("testing")

	if expected != actual {
		t.Errorf("Test failed!, expected: '%s', got: '%s'", expected, actual)
	}
}
