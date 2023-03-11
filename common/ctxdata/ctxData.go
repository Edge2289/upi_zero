package ctxdata

import (
	"context"
	"encoding/json"
)

// CtxUidKey jwt中的key值
var CtxUidKey = "jwtUserId"

// CtxSubAcctNoKey jwt中的key值
var CtxSubAcctNoKey = "jwtSubAcctNo"

// GetUidFromCtx 获取jwt头的uid参数
func GetUidFromCtx(ctx context.Context) int64 {
	return 1
	//uid, _ := ctx.Value(CtxUidKey).(json.Number).Int64()
	//return uid
}

// GetSubAcctNoFromCtx 获取jwt头的 sub_acct_no 参数
func GetSubAcctNoFromCtx(ctx context.Context) string {
	s := ctx.Value(CtxSubAcctNoKey).(json.Number).String()
	return s
}