TARGET=gocrawler
PKG=$(TARGET)
TAG=latest
IMAGE_PREFIX?=liuqi04
IMAGE_PREFIX_PRD=liuqi04
TARGET_IMAGE_DEV=$(IMAGE_PREFIX)/$(TARGET):$(TAG)
TARGET_IMAGE_PRD=$(IMAGE_PREFIX_PRD)/$(TARGET):$(TAG)
all: image

$(TARGET):
	CGO_ENABLED=0 go build -o dist/$(TARGET) $(PKG)

git log:

target:
	mkdir -p dist
	git log | head -n 1
