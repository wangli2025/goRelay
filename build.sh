#!/bin/bash

gitLog=$(git log -1 --pretty=format:"%H")
version=$1

make VERSION=${version} GITCOMMIT=${gitLog}

