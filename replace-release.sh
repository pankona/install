#!/bin/bash -eux

git tag --delete $1 || true
git tag $1
git push origin $1 -f
gh release delete $1 --repo pankona/install || true
make release
