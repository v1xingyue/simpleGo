#!/bin/sh
rm DownloadTaskRunner
rm /data0/shareGo/bin/DownloadTaskRunner
go build
go install
ps aux | grep DownloadTaskRunner | grep -v "grep" | awk '{print $2}' | xargs kill -9 
/data0/shareGo/bin/DownloadTaskRunner >> /dev/null 2>&1 &
