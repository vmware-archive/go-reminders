// Copyright 2015 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
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
