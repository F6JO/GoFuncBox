# 协程池

***

* github.com/F6JO/gopool	导入库

* Pool := gopool.NewSimplePoll(数量)    设置协程池数量

* Pool.Add(函数(参数))    给协程池添加任务，可以用循环批量添加

* Pool.Run()    协程池启动