#!/bin/bash

cur_dir=$(pwd)
pb_dir=$(pwd)/pkg/pb
mkdir -p ${pb_dir}

cd api
proto_dir=$(pwd)
file_names=$(ls $proto_dir)

for file_name in $file_names
do
    file_path=$proto_dir/$file_name
    if [ "${file_path##*.}x" == "proto"x ]; then
        echo ${cur_dir}
        ${cur_dir}/tools/protoc/bin/protoc --proto_path="" -I . --go_out=${pb_dir} --go-grpc_out=${pb_dir} $file_name
        
        code=$(echo $?)
        if [ $code = 0 ]; then
            echo -e "编译文件: $file_path => [\033[31m成功\033[0m] "
        else
            echo -e "[\033[34mFatal\033[0m]: 编译文件: [$file_path] => [\033[34m失败\033[0m] "
            echo -e "\t <<<<<<<<<<<< 编译过程意外退出，已终止  <<<<<<<<<<<<"
            exit
        fi
    fi 
done
