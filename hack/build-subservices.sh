#!/bin/sh

[ -z $GOPATH ] && GOPATH=$(go env GOPATH)

[ -z $GOPATH/src/github.com/gomeet-examples/svc-echo ] \
	&& echo "$GOPATH/src/github.com/gomeet-examples/svc-echo not found" \
	&& exit 1

[ -z $GOPATH/src/github.com/gomeet-examples/svc-profile ] \
	&& echo "$GOPATH/src/github.com/gomeet-examples/svc-profile not found" \
	&& exit 1

cd $GOPATH/src/github.com/gomeet-examples/svc-echo \
echo "Build $GOPATH/src/github.com/gomeet-examples/svc-echo"
make

cd $GOPATH/src/github.com/gomeet-examples/svc-profile \
echo "Build $GOPATH/src/github.com/gomeet-examples/svc-profile"
make
