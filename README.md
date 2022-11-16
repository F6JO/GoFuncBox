# 自用方法

github.com/F6JO/GoFuncBox

***

### 协程池

* Pool := GoFuncBox.NewSimplePoll(数量)	设置协程池数量
* Pool.Add(函数(参数))    给协程池添加任务

* Pool.Run()    协程池启动

### 请求头字符串转map

* GetHeaders(path)	返回map与err

### 命令行相关方法

* PrintCmdOutput(cmd *exec.Cmd) 实时打印命令输出