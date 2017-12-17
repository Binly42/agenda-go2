[![Build Status](https://travis-ci.org/Binly42/agenda-go2.svg?branch=master)](https://travis-ci.org/Binly42/agenda-go2)
[![codecov](https://codecov.io/gh/LIANGTJ/agenda-go2/branch/master/graph/badge.svg)](https://codecov.io/gh/LIANGTJ/agenda-go2)

## 简介

 agenda-go2 是在 [agenda-go](https://github.com/Binly42/agenda-go) 的基础上继续开发作业的, 对应 pml老师 的 [这篇博客](http://blog.csdn.net/pmlpml/article/details/78727210) 及 `ex-service-agenda.html`(*服务程序开发实战 - Agenda*) 。


## 获取 代码

```shell
go get -u github.com/Binly42/agenda-go2/...
```

## 获取 docker 镜像

```shell
sudo docker pull binly/agenda-go2
```


## *service* 用法

```shell
Usage:
  -p string
        The PORT to be listened by agenda. (default "8080")
```

## *cli* 用法

```shell
Usage:
  agenda [flags]
  agenda [command]

Available Commands and local flags:
  createM    -s startTime //create Meeting 
  			-e endTime 
  			-t title 
  			-p participator 
  			      
            
 
  help        Help about any command
  
  login      -u username  //login 
  			-p password
  
  logout     			 //logout
  
  register   -u username  //register for further use
  			-p password
  			[-e] email
  			[-t] phone
  			
  query      -u           //search users or meetings
  			-m			// note that -u & -m can't appear at the same time
  			-s startTime
  			-e endTime

Root Flags:
  -a, --author string         Author name for copyright attribution (default "YOUR NAME")
      --config string         config file (default is $HOME/.cobra.yaml) (default "./.cobra.yaml")
  -h, --help                  help for agenda
  -l, --license licensetext   Name of license for the project (can provide licensetext in config)
  -b, --projectbase string    base project directory eg. github.com/spf13/
      --viper                 Use Viper for configuration (default true)

Use "agenda [command] --help" for more information about a command.
```


## service 实现原理

 大致上, 与 [agenda-go](https://github.com/Binly42/agenda-go) 的区别主要有:

> * 借助 **gorm** 与 splite3 数据库 进行交互, 具体实现包装在 [*model* 模块](https://github.com/Binly42/agenda-go2/tree/master/service/vendor/model), 由某个 *service* 模块 中提供的 各项 service 调用, 这里还没分得那么细 ... 所以 *service* 也直接放在 [*agenda* 模块](https://github.com/Binly42/agenda-go2/tree/master/service/vendor/agenda) 中了 ... ;

> * 本来打算沿用 *User as Actor* 形式的, 这里还没有作适配, 所以也基本没有 log ... ;

> * server 方面, 则直接在 对请求作点处理 之后 调用 各项 service ;



## 样例

先把镜像跑起来:
```shell
sudo docker run -it -p 8080:8080 binly/agenda-go2
```
再在容器内启动服务器:
```shell
service -p 8080
```

register 

```shell
mock测试：
PS E:\GoWorkSpace\src\github.com\LIANGTJ\agenda-go2\cli> .\main.exe register -u root -p 123
register called
register successfully

服务端：
【200 response】
liangtj@ubuntu:~/Desktop/GoWorkSpace/src/github.com/LIANGTJ/agenda-go2/cli$ ./main register -u ltj  -p 123 -e ltj@163.com -t 12345
[info]2017/12/16 07:30:27 root.go:67: Can't read config: open ./.cobra.yaml: no such file or directory
register called
[Register] Response:  
register successfully ltj


【非200 response】:

liangtj@ubuntu:~/Desktop/GoWorkSpace/src/github.com/LIANGTJ/agenda-go2/cli$ ./main register
[info]2017/12/16 07:29:30 root.go:67: Can't read config: open ./.cobra.yaml: no such file or directory
register called
Error[registerd]： user regiestered invalid

liangtj@ubuntu:~/Desktop/GoWorkSpace/src/github.com/LIANGTJ/agenda-go2/cli$ ./main register -u root -p 123 -e ltj@163.com -t 12345
[info]2017/12/16 07:30:16 root.go:67: Can't read config: open ./.cobra.yaml: no such file or directory
register called
Error[registerd]： the user has been existed


```

login

```shell
mock测试：
PS E:\GoWorkSpace\src\github.com\LIANGTJ\agenda-go2\cli> .\main.exe login -u root -p 123
login called by root
login with info password: 123
Login Sucessfully root

服务端测试：
liangtj@ubuntu:~/Desktop/GoWorkSpace/src/github.com/LIANGTJ/agenda-go2/cli$ ./main login  -u ltj  -p 123
[info]2017/12/16 07:35:32 root.go:67: Can't read config: open ./.cobra.yaml: no such file or directory
login called by ltj
login with info password: 123
[Login] Response:  
Login Sucessfully ltj
```

因为没多大意义，所以以下都不再展示mock测试

query user

```shell
liangtj@ubuntu:~/Desktop/GoWorkSpace/src/github.com/LIANGTJ/agenda-go2/cli$ ./main query -u
[info]2017/12/16 07:36:05 root.go:67: Can't read config: open ./.cobra.yaml: no such file or directory
query called
[QueryAccountAll] Response:  
+--------+-------+-------+
|  NAME  | EMAIL | PHONE |
+--------+-------+-------+
| root   |       |       |
| matrix |       |       |
| ltj    |       | 12345 |
+--------+-------+-------+

```

create meeting

```shell
liangtj@ubuntu:~/Desktop/GoWorkSpace/src/github.com/LIANGTJ/agenda-go2/cli$ ./main createM -s "2011-01-01 10:00:34" -e "2011-01-02 08:00:34" -t MatrixShareMeeting -p lrd
create Meeting called
start: 2011-01-01 10:00:34 +0000 UTC end: 2011-01-02 08:00:34 +0000 UTC
sucessfully create meeting
```

 logout

```shell
liangtj@ubuntu:~/Desktop/GoWorkSpace/src/github.com/LIANGTJ/agenda-go2/cli$ ./main logout
[info]2017/12/16 07:38:39 root.go:67: Can't read config: open ./.cobra.yaml: no such file or directory
logout called
logout sucessfully
```








