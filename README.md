`【CODE IS AS ELEGANT AS POETRY】代码就像诗歌一样优雅`

BelongTo:Shinnosuke

### 项目简
upiGo 版本

#### 常用的命令
```go
rpc命令
cd pd
goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../

api命令
cd desc
goctl api go -api account.api -dir ../

model命令
goctl model mysql ddl -src accounts.sql -dir .  -cache=true --style=goZero
```
### 各个服务使用的端口汇总
```go
rpc使用etcd管理rpc，所以不需要理会rpc的端口，只要保证同一容器下各个rpc端口不冲突就可以了

unionpay
    api: 8832
    rpc: 3001

account
    api: 8822
    rpc: 2001

splitBillsSummary
    api: 8842
    rpc: 4001


```
