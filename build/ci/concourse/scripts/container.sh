#!/bin/bash
# go-reminders container.sh
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

# Copy buikld results to the output
echo "Copy artifacts to output"
cp -a ${TOP}/build/go-reminders ${TOP}/container/
cp -a ${TOP}/git-reminders-repo/build/docker/Dockerfile ${TOP}/container/
mkdir -p ${TOP}/container/html/skeleton
cp -a ${TOP}/git-reminders-repo/web/static/Skeleton/css ${TOP}/container/html/skeleton/
cp -a ${TOP}/git-reminders-repo/web/static/Skeleton/images ${TOP}/container/html/skeleton/

# Check what got laid down
echo "List out the container directory"
ls -laRt ${TOP}/container
echo ""
