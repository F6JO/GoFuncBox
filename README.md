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

* PrintCmdOutput(comm string, proc_line func(str string)) string 传入一行命令并执行，传入处理函数处理每一行的输出，返回值为全部输出
* exec_comm(comm string) 传入一句命令，解析后返回*exec.Cmd