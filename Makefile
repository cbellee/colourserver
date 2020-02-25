VERSION := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date -R)
NAME := $(shell basename `git rev-parse --show-toplevel`)
VENDOR := $(shell whoami)
SEMVER := 0.1.3 #$(shell cat ./app/semver)
ENV := dev
COLOUR := blue

print:
	@echo VERSION=${VERSION}
	@echo SEMVER=${SEMVER}
	@echo BUILD_DATE=${BUILD_DATE}
	@echo NAME=${NAME}
	@echo VENDOR=${VENDOR}
	@echo COLOUR=${COLOUR}

build:
	docker build \
	-t belstarr/colourserver:${ENV}-${SEMVER} \
	--build-arg VERSION="${VERSION}" \
	--build-arg SEMVER="${SEMVER}" \
	--build-arg COLOUR="${COLOUR}" \
	--build-arg BUILD_DATE="${BUILD_DATE}" \
	--build-arg NAME="${NAME}" \
	--build-arg VENDOR="${VENDOR}" .

push:
	docker push belstarr/colourserver:${ENV}-${SEMVER}

run:
	docker run \
	-d \
	-p 8080:80 \
	-e VERSION=${SEMVER} \
	belstarr/colourserver:${ENV}-${SEMVER}
