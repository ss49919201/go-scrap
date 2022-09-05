# 既に環境変数として定義されている場合は、そちらを優先したいので`?=`を使う
GENERATE_PATH ?= ./generate
GO_SRC_PATH := ./

PHONY: gen
gen:
	go generate $(GENERATE_PATH)

PHONY: build
build:
	go build -o ./bin/ -tags netgo -installsuffix netgo -ldflags '-s -w' $(GO_DUBUG_FLAGS) $(GO_SRC_PATH)

PHONY: build-debug
build-debug:
	$(MAKE) build GO_DUBUG_FLAGS="-gcflags \"all=-N -l\""
