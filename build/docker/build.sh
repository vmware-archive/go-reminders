#!/bin/bash
#
# Copyright 2015-2019 VMware, Inc. All Rights Reserved.
# Author: Tom Hite (thite@vmware.com)
#
# SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
#

if [ -z "${CONTAINER}" ]; then
    echo 'Environment variables CONTAINER is not set.'
    echo 'It should be set to the container name to build and push.'
    exit 1
fi

# Grab the latest build output
cp -a ../../cmd/go-reminders/go-reminders .

# Copy the relevant skeleton code
mkdir -p html/skeleton
cp -a ../../web/static/Skeleton/css html/skeleton/
cp -a ../../web/static/Skeleton/images html/skeleton/

# Copy the template files
cp -a ../../web/templates/tmpl html/
cp -a ../../web/templates/stats html/

# Build and push the container
docker build --rm -t ${CONTAINER} .
docker push ${CONTAINER}
