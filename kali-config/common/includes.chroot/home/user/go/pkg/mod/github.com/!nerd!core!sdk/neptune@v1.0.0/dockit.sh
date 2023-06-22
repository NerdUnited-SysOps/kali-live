#!/usr/bin/env sh
#

IMAGE_NAME="bridge/1.1"
IMAGE_VERSION="latest"

docker run \
	-it \
	--rm \
	--user $(id -u):$(id -g) \
	--volume "$(pwd)":/code \
	"$IMAGE_NAME:$IMAGE_VERSION" "$@" \
	
