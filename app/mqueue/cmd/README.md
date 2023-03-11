scheduler 设置定时任务
```md
    设置定时任务，塞进 task 里面，由 job 进行消费
```

job 作为消费定时任务

jobType 任务名称的定义

md 异步任务，允许在业务中添加的
```md
    定义：在`md - internal - listen` 下的 `Mqs` 函数中添加对应需要监听的 `Topic`
            `customer` 接收的是一个 `string` 化参数
    发布：在api-rpc中 servicecontext 中添加 *kq.Pusher 
                然后使用 l.svcCtx.XXX.Push(string)
    例子：unionPay 中的 Rpc部分
```