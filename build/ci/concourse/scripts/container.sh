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

# Copy build results and static content to the output
echo "Copy build artifacts to output"
cp -a ${TOP}/build/go-reminders ${TOP}/container/
# Copy the docker build file
cp -a ${TOP}/git-reminders-repo/build/docker/Dockerfile ${TOP}/container/
mkdir -p ${TOP}/container/html/skeleton
# Copy the Skeleton files
cp -a ${TOP}/git-reminders-repo/web/static/Skeleton/css ${TOP}/container/html/skeleton/
cp -a ${TOP}/git-reminders-repo/web/static/Skeleton/images ${TOP}/container/html/skeleton/
# Copy the html template files
cp -a ${TOP}/git-reminders-repo/web/templates/tmpl ${TOP}/container/html/
cp -a ${TOP}/git-reminders-repo/web/templates/stats ${TOP}/container/html/

# Check what got laid down
echo "List out the container directory"
ls -laRt ${TOP}/container
echo ""
