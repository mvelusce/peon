#!/usr/bin/env bash 

SCOPE="$1"

# Get last commit and parse text
if [[ -z "$SCOPE" ]]; then
    SCOPE=$(git log -1 | egrep -ohi '(MAJOR|MINOR|PATCH):' | head -1 | tr '[:upper:]' '[:lower:]')
fi

if [[ -z "$SCOPE" ]]; then
  SCOPE="auto"
fi

echo "Using scope $SCOPE"

# Get the next version, without tagging
echo "Getting next version"
nextversion="$(source semtag final -fos $SCOPE)"
echo "Publishing with version: $nextversion"

# Update the tag with the new version
source semtag final -f -v $nextversion

export PROG_VERSION=$nextversion

# Build
sh build.sh

# Create release notes
git log -1 | tail -n +5 > release-notes.md
