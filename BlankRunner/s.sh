#!/bin/sh

ps aux | grep DownloadTaskRunner | grep -v "grep" | awk '{print $2}' | xargs kill -9 
/data0/shareGo/bin/DownloadTaskRunner > /data0/shareGo/logs/DownloadTask/run.log 2>&1 &
