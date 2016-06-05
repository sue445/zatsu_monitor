#!/bin/bash -e

# build and compress file
#
# If you use Mac, install GNU sed
#   $ brew install gnu-sed

readonly DIST_DIR="dist"
readonly BIN_NAME="zatsu_monitor"
readonly MODE=$1

# get version from version.go
function get_version(){
    if [  `which gsed | wc -l` = "1" ]; then
        # Use GNU sed instead of BSD sed (for Mac)
        local _sed="gsed"
    else
        local _sed="sed"
    fi

    grep VERSION version.go | $_sed -r 's/.+?"(.+)".+?/\1/g'
}

function build(){
    local goos=$1
    local goarch=$2
    local zip_name="${BIN_NAME}_${version}_${goos}_${goarch}.zip"

    GOOS=$goos GOARCH=$goarch go build -o "${DIST_DIR}/${BIN_NAME}" *.go

    pushd $DIST_DIR > /dev/null
    rm -f $zip_name
    zip -m -q $zip_name $BIN_NAME
    popd > /dev/null

    echo "Write: ${DIST_DIR}/${zip_name}"
}

version=`get_version`

build "darwin" "amd64"
build "linux" "amd64"
build "linux" "arm"

if [  "${MODE}" = "release" ]; then
    git tag -a ${version} -m "Release v${version}"
    echo "Add tag: ${version}"

    git push origin master
    git push --tags
fi
