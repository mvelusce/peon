#!/usr/bin/env bash

TOKEN="$GITHUB_PAT"
REPO="skyveluscekm/peon"
FILE="peon.zip"
VERSION="latest"
GITHUB="https://api.github.com"

alias errcho='>&2 echo'

current_peon_version=""
if [ -f .peon-version ]; then
    current_peon_version=$(cat .peon-version)
fi
latest_peon_version=$(curl -s -H "Authorization: token $TOKEN" \
    https://api.github.com/repos/skyveluscekm/peon/releases/latest |\
    grep tag_name | sed -E 's/.*(v[0-9]+\.[0-9]+\.[0-9]+).*/\1/g')

if [ "$current_peon_version" = "$latest_peon_version" ]; then
    echo "Peon is up-to-date"
    exit 0
fi

echo "Peon is not up-to-date. Downloading latest version..."

function gh_curl() {
  curl -H "Authorization: token $TOKEN" \
       -H "Accept: application/vnd.github.v3.raw" \
       $@
}

if [ "$VERSION" = "latest" ]; then
  # Github should return the latest release first.
  parser=".[0].assets | map(select(.name == \"$FILE\"))[0].id"
else
  parser=". | map(select(.tag_name == \"$VERSION\"))[0].assets | map(select(.name == \"$FILE\"))[0].id"
fi;

asset_id=`gh_curl -s $GITHUB/repos/$REPO/releases | jq "$parser"`
if [ "$asset_id" = "null" ]; then
  errcho "ERROR: version not found $VERSION"
  exit 1
fi;

wget -q --auth-no-challenge --header='Accept:application/octet-stream' \
  https://$TOKEN:@api.github.com/repos/$REPO/releases/assets/$asset_id \
  -O peon-latest.zip

unzip -o peon-latest.zip

rm peon-latest.zip

echo "$latest_peon_version" > .peon-version
