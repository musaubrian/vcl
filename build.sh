#!/bin/sh

rm -r ./bin
go build  -o bin/raygo .
if [ $# -eq 0 ]; then
    echo "
YOU NEED TO PARSE THE PATH TO THE FILE

./build.sh ~/path/to/music/file
    "
else
    v=$(echo "$1" | tr '_' '#' | tr ' ' '_')
    echo $v
    ./bin/raygo -f $v
fi


