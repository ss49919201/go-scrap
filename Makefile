# 既に環境変数として定義されている場合は、そちらを優先したいので`?=`を使う
GENERATE_PATH ?= ./generate

PHONY: gen
gen:
	go generate $(GENERATE_PATH)
