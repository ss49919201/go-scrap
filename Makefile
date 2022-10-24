# コマンドと同名のファイルがあった時に困らないように`.PHONY`を使ってコマンド名をダミーターゲットとして定義しておく
# 変数の値が追いづらくなるので再帰変数展開`=`や上乗せ`+=`は使わない
# 引数->変数->環境変数の優先順序であることに注意

# `?=`は変数が未定義の時だけ値を代入する
# つまり未定義時は`:=`と同等である

# 既に環境変数として定義されている場合は、そちらを優先したいので`?=`を使う
# e.g
# `GENERATE_PATH := ./generate` だと `GENERATE_PATH=hoge make gen` -> `go generate ./generate`
# `GENERATE_PATH ?= ./generate` だと `GENERATE_PATH=hoge make gen` -> `go generate hoge` になる

# ref: https://nekonenene.hatenablog.com/entry/make-define-variables

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

PHONY: test
test:
	docker-compose -p test -f docker-compose.test.yml run test sh -c " \
	CGO_ENABLED=0 go test -cover -coverprofile=./coverage.out ./... -count=1  && \
	go tool cover -func=coverage.out | grep "total:" \
	"

# TODO: 動かす
# ERRO Running error: context loading failed: failed to load packages: timed out to load packages: context deadline exceeded 
# ERRO Timeout exceeded: try increasing it by passing --timeout option
# PHONY: lint
# lint:
# 	docker compose -p golang-lint -f docker-compose.lint.yml run lint golangci-lint run
