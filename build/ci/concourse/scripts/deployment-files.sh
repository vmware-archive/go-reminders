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
cp -a git-reminders-repo/deployments/kubernetes/base kubernetes/
cp -a git-reminders-repo/deployments/kubernetes/overlays kubernetes/
cat >kubernetes/base/kustomization.yaml <<EOF
  resources:
  - service.yaml
  - deployment.yaml
  images:
  - name: docker-registry-repo
    newTag: ${tag}
    newName: ${container}
EOF

# create the helm chart copy and fixup a values file
cp -a git-reminders-repo/deployments/helm kubernetes/
cd kubernetes/helm

# fixup the values file
cp values-minikube.yml values.yml
repl="${container//\//\\/}"
sed -i -e "s/repository: tdhite\/go-reminders/repository: ${repl}/g" values.yml
sed -i -e "s/tag: latest/tag: ${tag}/g" values.yml

# Check whats here
cd "${TOP}"
echo "List out the output directory:"
ls -laRt kubernetes
echo ""
