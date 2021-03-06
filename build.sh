#!/bin/bash

#
# Build wundertools in a container
#
# @NOTE to specify a different os/arch:
#    - GOOS : linux darwin windows
#    - GOARCH : amd64 arm arm64
#
# @NOTE !does not install it yet
#  (installs it, but inside the container)
#

QUIET='no'
INSTALL='ask'

# @TODO we can make this more advanced.
if [ $1 == '--automated' ]; then
  QUIET="yes"
  INSTALL="no"
  shift
fi

source make/.os-detect
source make/make.sh

EXEC_PATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
INTERNAL_LIBRARY_PATH="github.com/CoachApplication/app-cli"

echo "***** Building COACH cli client.

This will build the coach-cli as a 'coach' binary for '$GOOS-$GOARCH'.

(Override this by setting \$GOOS and \$GOARCH environment variables)

 **** Building in containerized golang environment
 "

# some sanity stuff, to prevent docker related permissions issues
mkdir -p "${COACH_BUILD_PATH}"
mkdir -p ".git/modules/vendor"
chmod u+x Makefile

# Run the build inside a container
#
#  - volumify the submodule changes
#  - build in a valid gopath to get active vendor dependencies
#  - pass in env variables for environment control
docker run --rm -ti \
	-v "${EXEC_PATH}:/go/src/${INTERNAL_LIBRARY_PATH}" \
	-v "/go/src/${INTERNAL_LIBRARY_PATH}/.git/modules/vendor" \
	-v "/go/src/${INTERNAL_LIBRARY_PATH}/vendor" \
	-e "GOOS=${GOOS}" \
	-w "/go/src/${INTERNAL_LIBRARY_PATH}" \
	golang:${GOVERSION} \
	make build

echo " 

Finished building the application inside the container.  If an error occured
during the golang compile, then you would have seen it reported above.

"

echo " **** Containerized build complete 

an executable binary has (hopefully) now been built 
in ${COACH_BUILD_BINARY_PATH}

"

# @TODO implement some improved logic for determining
#    Install path, and sudo

export COACH_INSTALL_PATH="/usr/local/bin"

if [ "$INSTALL" == 'ask' ]; then

    echo " **** Installation

    This installer can now install the built binary for you,
    if you don't want to do it manually.

    The planned installation path is : ${COACH_INSTALL_PATH}

    Would you like to me install a binary to that location? (y/n)
    "
    read  yninstall
    case "$yninstall" in
        [Yy]* )

            INSTALL="yes"

            ;;
        *)
            echo " "
            echo "skipped installation"
            ;;
    esac

fi

if [ "$INSTALL" == "yes" ]; then
    if [ -w "COACH_INSTALL_PATH" ] ; then
        export COACH_INSTALL_SUDO=""
    else
        export COACH_INSTALL_SUDO="`which sudo`  -E"
        echo "--> detected that sudo will be required, as you don't have write privilege to the target path"
    fi

    ${COACH_INSTALL_SUDO} make install
fi
