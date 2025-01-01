#!/bin/bash

gitLog=$(git log -1 --pretty=format:"%H")
gitLog=9a4bcb7f2eaf9c6b820f42b6b8758d986b38fd1f
version=$1
nowDate=$(date +"%F")

make VERSION=${version} BuildAt=${nowDate} GITCOMMIT=${gitLog}

