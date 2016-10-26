# VMovie
22v.net 的网站源代码，使用beego 1.6.1开发，golang 1.6编译 

最新版已经使用golang 1.7 / beego 1.7编译

###环境
mysql  
golang 1.6/golang 1.7  
nginx  
supervisord

###数据库
`create database vmovie`

###数据库连接配置
`vi conf/app.conf`

### 安装及运行
`go get github.com/zituocn/22v`  
`更改22v目录为VMovie`	
`go build` 或 `bee run`	
`./vmovie`  
会自动建表

###后台登录
`http://ip:port/vvadmin/`  
`app.conf中可配置后台目录地址`
请手动向user_info表中写入管理员登录帐号/密码(大写32位md5)的记录

###最近更新 
2016.10.26
>增加向baidu站长平台push网址的功能，使用beego的task功能，默认为10分钟一次，频率可以app.conf中配置。task的功能详情请查看 https://beego.me/docs/module/toolbox.md#task

2016.9.8 
>一、增加后台图片上传水印功能，水印图片指定为/static/img/mark.png  
>二、修改部分样式和模板文件  
>三、网站备案成功，放回国内。

###网站地址
https://22v.net

###手机版
https://m.22v.net