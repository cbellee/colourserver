VERSION := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date -R)
NAME := $(shell basename `git rev-parse --show-toplevel`)
VENDOR := $(shell whoami)

print:
	@echo VERSION=${VERSION}
	@echo BUILD_DATE=${BUILD_DATE}
	@echo NAME=${NAME}
	@echo VENDOR=${VENDOR}

build:
	docker build -t colourserver:${VERSION} \
	--build-arg VERSION="${VERSION}" \
	--build-arg BUILD_DATE="${BUILD_DATE}" \
	--build-arg NAME="${NAME}" \
	--build-arg VENDOR="${VENDOR}" .

tag:
	docker tag colourserver:${VERSION} belstarr/colourserver:latest

push:
	docker push belstarr/colourserver:latest

run:
	docker run -d -p 8080:8080 \
	colourserver:${SEMVER}
