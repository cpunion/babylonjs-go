#!/bin/sh

# read version from package.json
export BABYLONJS_VERSION=$(node -p "require('./package.json').version")
