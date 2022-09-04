#/bin/bash

# NOW=`date "+%Y%m%d"`
# DIR=${1:-$NOW}
FILES=`ls pattern`

# 1回目のイテレートで親ディレクトリもまとめて作る
for FILE in ${FILES[@]}; do
    # TARGET=${DIR}/${Q}
    # mkdir -p ${TARGET}
    # cp template/main.go ${TARGET}/main.go
    BASENAME=${FILE%.*}
    mkdir examples/${BASENAME}
    cp pattern/${FILE} examples/${BASENAME}/main.go
    rm pattern/${FILE}
done
