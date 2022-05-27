# 介绍
利用outline的API，执行导出数据的操作
# 使用
```
docker run --rm -e URL=<https://outline.xxx.com> -e TOKEN=<Your TOKEN> -v /etc/timezone:/etc/timezone:ro -v /etc/localtime:/etc/localtime:ro -v /path:/backup 601096721/outline-export

```

触发之后outline默认把备份发送到你的邮箱（前提是你设置了发件邮箱）,然后根据时间作为文件名保存在`/backup`下面

github仓库：[https://github.com/Sagit-chu/outline-export](https://github.com/Sagit-chu/outline-export)
