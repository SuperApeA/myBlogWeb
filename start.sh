#!/bin/bash

app_path="/home/workspace/myBlog/myBlogWeb/server"

./myBlog &

if [ -f "$app_path/myBlog" ]; then
    echo "myBlog is running!"
else
    echo "myBlog start failed!"
fi