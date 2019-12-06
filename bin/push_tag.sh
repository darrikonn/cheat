#!/bin/bash

set -euo pipefail

cheat_version=$(go run cheat --version)
version=v${cheat_version#"cheat version "}
git tag -a $version -m "Cheat version $version"
git push origin $version
