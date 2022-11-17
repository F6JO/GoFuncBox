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

* PrintCmdOutput(comm string, proc_line func(str string)) 
  * 传入一行命令并执行，传入处理函数处理每一行的输出，返回值为全部输出
  * 参数1：要执行的命令
  * 参数2：对每一行处理的函数
* exec_comm(comm string) 
  * 传入一句命令，解析后返回*exec.Cmd

### 文件操作相关

* PathExists(path string) 
  * 判断文件或目录是否存在，返回布尔值和err
* Py_Writefile(fileName string, nr string, fangshi string) 
  * python的文件写入方式
  * 参数1：文件路径 
  * 参数2：写入内容 
  * 参数3：方式，w覆盖，a追加