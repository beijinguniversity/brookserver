# brookserver
brookserver

基于Beego。开发的Brook流控的后端（服务端）


## 服务端搭建教程

0.把myBrookServer.tar.gz上传并解压到服务器上

1.设置服务器的时区

```linux
		不会请Google
```

2.安装iptables
```linux
		yum install iptables
		不会请Google
```

3.安装lsof
```linux
		yum install lsof
		不会请Google
```

4.修改conf/app.conf

4.1.修改mysql配置

4.2.修改lp_brook_server_id(后端添加节点后的id)

5.你需要把项目跑起来！

```linux
		 nohup ./myBrookServer &
```

### 说明
以上步骤无特殊说明，需必做，否则跑不起来

后端端口范围 1024-60000 因此这个项目的用户量最大为 60000-1024

服务器上的其他应用程序也`不要`占用此范围的端口

目前还没有自动删除长期未使用的用户，所以你只需要知道就好

### 站在巨人的肩上

BeeGo: https://github.com/astaxie/beego

Brook: https://github.com/txthinking/brook

Jquery

Bootstrap

Mysql

Redis


