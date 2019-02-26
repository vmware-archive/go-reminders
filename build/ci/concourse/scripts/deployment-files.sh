#!/bin/bash
# go-reminders deployment-files.sh
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

set -e -x

echo "------Environment Variables------"
set

# Save current directory
TOP="$(pwd)"

ret=0
if [ -z "${container}" ]; then
    echo "ERROR: container not supplied. Aborting!"
    ret=1
fi
if [ $ret -ne 0 ]; then
    exit $ret
fi

tag="$(cat version/version)"

# Make the output area if it does not exist
mkdir -p ${TOP}/kubernetes

# create the kubernetes deployment and service manifests
repl="${container//\//\\/}:${tag}"
echo "Replacement regexp provided: ${repl} from ${container}"
sed -e "s/{{docker-registry-repo}}/${repl}/g" git-reminders-repo/deployments/kubernetes/deployment.yml >kubernetes/deployment.yml
cp git-reminders-repo/deployments/kubernetes/service.yml kubernetes/service.yml

# create the helm chart copy
cp -a git-reminders-repo/deployments/helm kubernetes/

# Check whats here
echo "List out the output directory:"
ls -laRt kubernetes
echo ""
