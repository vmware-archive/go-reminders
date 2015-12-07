#!/bin/bash
# Copyright 2015 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

# Grab the latest build output
cp -a ../cmd/go-reminders/go-reminders .

# Copy the relevant skeleton code
mkdir -p html/skeleton
cp -a ../web/static/Skeleton/css html/skeleton/
cp -a ../web/static/Skeleton/images html/skeleton/

# Copy the template files
cp -a ../web/templates/tmpl html/
cp -a ../web/templates/stats html/

# Run the thing. Check it at http://localhost:8080
./go-reminders -dbtype=mem
