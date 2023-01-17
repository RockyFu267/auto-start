# auto-start
## 监控端口，不通就执行对应的启动脚本

## 用法
### 修改配置
+ interval 配置间隔时间
+ service.Port   要监听的服务端口
+ service.Startpath    对应端口的进程服务启动脚本

### 赋权
+ chmod +x auto-start

### 启动
+ nohup ./auto-start &

### 日志
+ 观察nohup.out日志文件不生成新的信息即可
+ 如果服务起不来，可以看该日志来排查，大概率为启动脚本内容有误
