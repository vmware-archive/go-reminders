#!/bin/bash
# go-reminders submodules.sh
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

set -e -x

# Save current directory
TOP="$(pwd)"

# check for git
which git
if [ $? -ne 0 ]; then
    "Git is not avaiabale in the image."
    exit 1
fi

# initialize submodules
cd git-reminders-repo
git submodule init
if [ $? -ne 0 ]; then
    "Git submodule init failed."
    exit 2
fi

git submodule update --recursive
if [ $? -ne 0 ]; then
    "Git submodule init failed."
    exit 3
fi
