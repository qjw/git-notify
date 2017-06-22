#!/usr/bin/env bash

error()
{
    echo "$@"
    return 1
}

[ -n "${GOPATH}" -a "${GOPATH}" != "" ] || error "GOPATH not exist" || exit 1

export PATH=$PATH:$GOPATH/bin

ensure_bin(){
    which "${1}" 2>/dev/null
    if [ $? -ne 0 ];then
        echo "install "${1}" tool"
        go get -u "${2}" || return 1
    else
        echo "${1} tool installed yet"
    fi
    return 0
}

ensure_bin govendor "github.com/kardianos/govendor" || exit 1
ensure_bin go-bindata "github.com/jteeuwen/go-bindata/..." || exit 1

if ! [ -f vendor/vendor.json ];then
    echo "init vendor repertory"
    govendor init
else
    echo "vendor init yet"
fi

govendor sync