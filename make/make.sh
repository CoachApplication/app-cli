#!/bin/bash

#!/usr/bin/env bash
set -e

# @NOTE Do not do any logic or functionality in this file
#    as it may in some circumstances be sources in an
#    escalated permission environment

# We should be determining these automatically somehow?
export GOOS="${GOOS:-linux}" # Perhaps you would prefer "osx" ?
export GOARCH="${GOARCH:-amd64}"
export GOVERSION="latest"

export COACH_PKG='github.com/CoachApplication/coach-cli'
export COACH_BUILD_PATH="./bin"
export COACH_BINARY_NAME="coach"

export COACH_BUILD_BINARY_PATH="${COACH_BUILD_PATH}/${COACH_BINARY_NAME}"

[ -z "${COACH_INSTALL_PATH}" ] && export COACH_INSTALL_PATH="${GOPATH}/bin"
export COACH_INSTALL_BINARY="${COACH_INSTALL_PATH}/${COACH_BINARY_NAME}"

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