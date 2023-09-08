#!/bin/sh

source scripts/version.sh

cat scripts/pkgs.txt | sed -e 's/\/.*$//' | xargs -I {} yarn add {}@$BABYLONJS_VERSION
