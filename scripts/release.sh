#!/bin/sh

source scripts/version.sh

# Check version exists on git tags
if git rev-parse $BABYLONJS_VERSION >/dev/null 2>&1; then
  echo "Version $BABYLONJS_VERSION already exists"
  exit 0
fi

# Add files to git
TAG=v$BABYLONJS_VERSION
git add -A
git commit -m "Release $TAG"
git tag $TAG
git push origin main
git push origin $TAG
