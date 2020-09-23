#!/usr/bin/env bash 

SCOPE="$1"

if [ -z "$SCOPE" ]; then
    echo "Scope is empty. Trying to get from commit message..."
    SCOPE=$(git log -1 | egrep -ohi '(MAJOR|MINOR|PATCH):' | head -1 | tr '[:upper:]' '[:lower:]')
fi

if [ -z "$SCOPE" ]; then
  echo "Scope is still empty. Setting it to patch..."
  SCOPE="patch"
fi

echo "Using scope $SCOPE"

last_version=$(git describe --match "v[0-9]*" --tags | egrep -o '[0-9]+\.[0-9]+\.[0-9]+')
echo "Last version: $last_version"

echo "Getting next version, without tagging"
next_version="$(./tools/shell-semver/increment_version.sh -p $last_version)"
if [ "$SCOPE" = "major" ]; then
    next_version=-Prelease.releaseVersion=$(./tools/shell-semver/increment_version.sh -M $last_version)
fi
if [ "$SCOPE" = "minor" ]; then
    next_version=-Prelease.releaseVersion=$(./tools/shell-semver/increment_version.sh -m $last_version)
fi

echo "Publishing with version: $next_version"

echo "Creating new tag"
git tag "v$next_version"
echo "Pushing tag to origin"
git push origin --tags

export PROG_VERSION=$next_version

echo "Building"
sh build.sh

echo "Create release notes"
git log -1 | tail -n +5 > release-notes.md
