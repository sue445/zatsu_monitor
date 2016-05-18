#!/bin/bash -e

# build and compress file

readonly DIST_DIR="dist"
readonly BIN_NAME="zatsu_monitor"
# TODO: get from version file
readonly VERSION="0.1.0"

function build(){
    local goos=$1
    local goarch=$2
    local zip_name="${BIN_NAME}_${VERSION}_${goos}_${goarch}.zip"

    GOOS=$goos GOARCH=$goarch go build -o "${DIST_DIR}/${BIN_NAME}" *.go

    pushd $DIST_DIR > /dev/null
    rm -f $zip_name
    zip -m -q $zip_name $BIN_NAME
    popd > /dev/null

    echo "Write: ${DIST_DIR}/${zip_name}"
}

build "darwin" "amd64"
build "linux" "amd64"
build "linux" "arm"
