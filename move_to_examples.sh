#/bin/bash

DIR=${1}
FILES=`ls ${DIR}`

# 1回目のイテレートで親ディレクトリもまとめて作る
for FILE in ${FILES[@]}; do
    BASENAME=${FILE%.*}
    mkdir examples/${BASENAME}
    cp ${DIR}/${FILE} examples/${BASENAME}/main.go
    rm ${DIR}/${FILE}
done