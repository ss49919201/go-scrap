#/bin/bash

NOW=`date "+%Y%m%d"`
DIR=${1:-$NOW}
QESTION_TYPES=(a b c d)

# 1回目のイテレートで親ディレクトリもまとめて作る
for Q in ${QESTION_TYPES[@]}; do
    TARGET=${DIR}/${Q}
    mkdir -p ${TARGET}
    cp template/main.go ${TARGET}/main.go
done
