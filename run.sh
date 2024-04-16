#!/bin/bash

app_path="/home/workspace/myBlog/myBlogWeb/server"

# 获取pid，用于停止程序
pid=$(pgrep myBlog)

#if [ $pid -eq 0 ]; then
#    # 程序在运行
#    kill $pid
#
#fi

while [ $pid > /dev/null ]
do
    echo "myBlog pid is :$pid, it will be kill and restart soon"
    kill $pid
    pid=$(pgrep myBlog)
done

cd "$app_path"; rm -rf myBlog; go build -o myBlog;  ./myBlog &

if [ -f "$app_path/myBlog" ]; then
    echo "myBlog is running!"
else
    echo "myBlog start failed!"
fi
