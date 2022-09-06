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

# `-`はエラーが発生しても続行する
# FIXME: ターゲットも引数にしたい
PHONY: test
test:
	-CGO_ENABLED=0 go test -cover -coverprofile=./coverage.out ./... -count=1
	-go tool cover -func=coverage.out | grep "total:"
	rm coverage.out
