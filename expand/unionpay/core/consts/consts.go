package consts

import "time"

// HTTP 请求报文 Header 相关常量
const (
	Authorization = "Authorization"  // Header 中的 Authorization 字段
	Accept        = "Accept"         // Header 中的 Accept 字段
	ContentType   = "Content-Type"   // Header 中的 ContentType 字段
	ContentLength = "Content-Length" // Header 中的 ContentLength 字段
	UserAgent     = "User-Agent"     // Header 中的 UserAgent 字段
)


// SDK 相关信息
const (
	Version         = "0.1.0"                      // SDK 版本
	UserAgentFormat = "UnionPay-Go/%s (%s) GO/%s" // UserAgent中的信息
)

// 常用 ContentType
const (
	ApplicationJSON = "application/json"
	ImageJPG        = "image/jpg"
	ImagePNG        = "image/png"
	VideoMP4        = "video/mp4"
)

// 时间相关常量
const (
	FiveMinute     = 5 * 60           // 回包校验最长时间（秒）
	DefaultTimeout = 30 * time.Second // HTTP 请求默认超时时间
)