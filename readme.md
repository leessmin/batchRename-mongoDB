# 介绍

这是一个用于批量修改mongoDB中文档字段的工具

# 使用
在exe可执行程序同级目录下创建一个config.conf的配置文件格式如仓库中的config.conf

config.conf采用 key = value 的形式进行配置

格式如下
| key      | valve                               | 描述               |
| -------- | ----------------------------------- | ------------------ |
| MongoUrl | 例如：mongodb://192.168.1.104:27017 | mongoDB的地址      |
| DBName   | 例如：test                          | 数据库名字         |
| CollName | 例如：[article, blog_info]          | 需要扫描的集合名字 |
| OldVal   | 例如：https://192.168.1.104         | 被替换的值         |
| NewVal   | 例如：http://1.1.1.1                | 替换的值           |

# 注
这是一个用go编写的程序，这也是Leessmin的第一个用Go语言编写的程序哦。可能写得不是很好，还请理解，有bug向我提交Issues哦，我会尽快解决哒，谢谢大家的使用！！！

# License

[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2022-present, Leessmin(李思敏)
