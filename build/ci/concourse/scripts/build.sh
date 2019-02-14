#!/bin/bash
# go-reminders build.sh
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

set -e -x

# Save current directory
TOP="$(pwd)"

# Show current setup
echo "GOPATH is: " $GOPATH
echo "TOP is: " $TOP
echo ""

# Build the beast
cd git-reminders-repo
CONTAINER=nomatter make cmd/go-reminders/go-reminders

# Check static linked binary
echo "Check static link status:"
if ldd cmd/go-reminders/go-reminders; then
    echo "The go-reminders binary is dynamically linked, cannot use it."
    exit 1
fi

# Copy build artifacts to the output directory
cp -a cmd/go-reminders/go-reminders ${TOP}/build/
