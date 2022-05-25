# 介绍
利用outline的API，执行导出数据的操作
# 使用
```
docker build -t out .
docker run -e URL=<https://outline.xxx.com> -e TOKEN=<Your TOKEN> out

```

触发之后outline默认把备份发送到你的邮箱（前提是你设置了发件邮箱）
