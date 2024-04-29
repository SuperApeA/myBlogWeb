#!/bin/bash

app_path="/home/workspace/myBlog/myBlogWeb/server"

# 获取pid，用于停止程序
pid=$(pgrep myBlog)

if [ -n "$pid" ]; then
  echo "myBlog is running with pid: $pid"
  while kill -0 $pid 2>/dev/null; do
    echo "Killing myBlog with pid: $pid"
    kill $pid
    sleep 1  # 添加延时，防止立即重试导致的问题
    pid=$(pgrep myBlog)
  done
  echo "myBlog has been killed."
else
  echo "myBlog is not running"
fi