# nameservice
debug调试方案
1. 编译
进入到项目下/cmd/nsd/文件夹下
go build -gcflags "all=-N -l"

2. 开启远程监听工具功能
2.1 安装dlv

启动节点并监听
* linux

dlv --listen=:2345 --headless=true --api-version=2 exec ./nsd start

* windows
dlv --listen=:2345 --headless=true --api-version=2 exec ./nsd.exe start

3.goland配置go remote