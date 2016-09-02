# VMovie
22v.net 的网站源代码，使用beego 1.6.1开发，golang 1.6编译 
最新版已经使用golang 1.7/beego 1.7编译

###环境
mysql  
golang 1.6+ /golang 1.7

###数据库
`create database vmovie`

###数据库连接配置
`vi conf/app.conf`

### 安装及运行
`go get github.com/zituocn/22v`	
`更改22v目录为VMovie`	
`go build` 或 `bee run`	
`./22v`  
会自动建表

###后台登录
`http://ip:port/vvadmin/`  
`app.conf中可配置后台目录地址`
请手动向user_info表中写入管理员登录帐号/密码的记录



###网站地址
https://22v.net

###手机版
https://m.22v.net