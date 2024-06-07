#!/bin/sh -eu
#
# Copyright 2024 Intel Corporation.
#
# SPDX-License-Identifier: Apache-2.0
#
# Invoke this script with a version as parameter
# and it will update all hard-coded devel versions
# to the tag versions in the source code.
#
# Adapted from https://github.com/intel/intel-device-plugins-for-kubernetes/

if [ $# != 1 ] || [ "$1" = "?" ] || [ "$1" = "--help" ]; then
    echo "Please provide TAG version as an argument. Usage: $0 <tag_version>" >&2
    exit 1
fi

devel_link="intel/intel-technology-enabling-for-openshift/main/"
tag_link="intel/intel-technology-enabling-for-openshift/$1/"

files=$(git grep -lF $devel_link -- '*.md')

for file in $files; do
    sed -i -e "s|$devel_link|$tag_link|g" "$file";
done