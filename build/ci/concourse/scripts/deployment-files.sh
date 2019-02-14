#!/bin/bash
# go-winery deployment-files.sh
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

# Make the output area if it does not exist
mkdir -p ${TOP}/docker

# List out params
echo "version: ${version}"

# create the docker image and files
echo "${docker_private_repo}/${docker_container}" >docker/repository
cat ${version} >docker/tag

# Check whats here
echo "List out the output directory:"
ls -lat docker
echo ""

# List the data
echo "Values of files:"
echo "repository: $(cat docker/repository)"
echo "tag: $(cat docker/tag)"
echo ""
