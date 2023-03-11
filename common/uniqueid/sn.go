package uniqueid

import (
	"fmt"
	"time"
	"upi/common/tool"
)

// SnPrefix 单号前缀
type SnPrefix string

const (
	SN_RETAIL_ORDER SnPrefix = "RO" // 零售订单前缀
)

// GenSn 生成单号
func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("200601021"), tool.Krand(7, tool.KC_RAND_KIND_NUM))
}

// GenUserIdSn 生成带用户ID的单号
func GenUserIdSn(snPrefix SnPrefix, userId int64) string {
	return fmt.Sprintf("%s%s%d%s", snPrefix, time.Now().Format("200601021"), userId, tool.Krand(6, tool.KC_RAND_KIND_NUM))
}