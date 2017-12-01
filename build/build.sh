#!/bin/bash -e
cd $(cd `dirname "$0"`; cd ..; pwd)

VERSION=$1
if [ -z $VERSION ]; then
	echo "Usage: $0 VERSION"
	exit 1
fi

TARGETS="darwin_amd64 linux_amd64"

for target in $TARGETS; do
	t=(${target//_/ })

	export GOOS=${t[0]}
	export GOARCH=${t[1]}
	export NAME=zanroo-inventory

	if [ $GOOS == "windows" ]; then
		NAME=${NAME}.exe
	fi

	go build -ldflags "-X main.buildVersion=$VERSION" -o pkg/${NAME}

	pushd pkg
	zip ${NAME}_${VERSION}_${GOOS}_${GOARCH}.zip ${NAME}
	rm ${NAME}
	popd
done
