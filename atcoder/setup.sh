#/bin/bash

function create_dir() {
    NOW=$1
    DIR=$2
    QESTION_TYPES=$3

    # 1回目のイテレートで親ディレクトリもまとめて作る
    for Q in ${QESTION_TYPES[@]}; do
        TARGET=${DIR}/${Q}
        mkdir -p ${TARGET}
        cp template/main.go ${TARGET}/main.go
    done
}

NOW=`date "+%Y%m%d"`
DIR=${1:-$NOW}
QESTION_TYPES=(a b c d)

create_dir $NOW $DIR $QESTION_TYPES
