#!/usr/bin/env bash 

SCOPE="$1"

# Get last commit and parse text
if [ -z "$SCOPE" ]; then
    echo "Scope is empty. Trying to get from commit message..."
    SCOPE=$(git log -1 | egrep -ohi '(MAJOR|MINOR|PATCH):' | head -1 | tr '[:upper:]' '[:lower:]')
fi

if [ -z "$SCOPE" ]; then
  echo "Scope is still empty. Setting it to auto..."
  SCOPE="auto"
fi

echo "Using scope $SCOPE"

echo "Getting next version, without tagging"
nextversion="$(sh semtag final -fos $SCOPE)"
echo "Publishing with version: $nextversion"

echo "Update the tag with the new version"
sh semtag final -f -v $nextversion

export PROG_VERSION=$nextversion

echo "Building"
sh build.sh

echo "Create release notes"
git log -1 | tail -n +5 > release-notes.md
