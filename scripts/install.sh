#!/bin/sh

VERSION=6.18.0

cat scripts/pkgs.txt | sed -e 's/\/.*$//' | xargs -I {} yarn add {}@$VERSION
