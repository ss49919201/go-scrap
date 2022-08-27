#/bin/bash

NOW=`date "+%Y%m%d"`
DIRS=(a b c d)

# 1回目のイテレートで日付ディレクトリもまとめて作る
for DIR in ${DIRS[@]}; do
    TERGET=${NOW}/${DIR}
    mkdir -p ${TERGET}
    cp template/main.go ${TERGET}/main.go
done
