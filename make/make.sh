#!/bin/bash

#!/usr/bin/env bash
set -e

# @NOTE Do not do any logic or functionality in this file
#    as it may in some circumstances be sources in an
#    escalated permission environment


if [ -z "$GOPATH" ]; then
    echo "WARNING: No GOPATH exists in your environment.  Certain components such as TESTs may produce weird errors"
fi

# We should be determining these automatically somehow?
export GOOS="${GOOS:-linux}" # Perhaps you would prefer "osx" ?
export GOARCH="${GOARCH:-amd64}"
export GOVERSION="latest"

export COACH_PKG_ROOT='github.com/CoachApplication/app-cli/'
export COACH_BUILD_PATH="./bin"
[ -z "${COACH_INSTALL_PATH}" ] && export COACH_INSTALL_PATH="${GOPATH}/bin"
export COACH_COMMANDS_PATH="cmd" # each folder in here should be built into a binary

COACH_ROOT="./"
COACH_DEPS="
vendor/github.com/CoachApplication/api
vendor/github.com/CoachApplication/base
vendor/github.com/CoachApplication/utils
vendor/github.com/CoachApplication/config
vendor/github.com/CoachApplication/project
vendor/github.com/CoachApplication/handler-dockercli
vendor/github.com/CoachApplication/handler-local
"
COACH_GO_TARGETS="$COACH_ROOT $COACH_DEPS"

# Build a bundle
bundle() {
 local bundle="$1"; shift
 echo "---> Make-bundle: $(basename "$bundle") (in $DEST)"
 source "make/$bundle" "$@"
}

if [ $# -gt 0 ]; then
 bundles=($@)
 for bundle in ${bundles[@]}; do
     export DEST=.
     ABS_DEST="$(cd "$DEST" && pwd -P)"
     bundle "$bundle"
     echo
 done
fi