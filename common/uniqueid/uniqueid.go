package uniqueid

import (
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logx"
)

var flake *sonyflake.Sonyflake

func init() {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

// GenId 获取唯一ID
func GenId() int64 {

	id, err := flake.NextID()
	if err != nil {
		logx.Severef("flake NextID failed with %s \n", err)
		panic(err)
	}

	return int64(id)
}