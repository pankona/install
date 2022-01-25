#!/bin/bash -e

git tag --delete v0.0.17
git tag v0.0.17
git push origin v0.0.17 -f
make release
