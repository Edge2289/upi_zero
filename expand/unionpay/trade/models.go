package trade

// UnifiedPayRequest 请求参数
type UnifiedPayRequest struct {
	Syscode        string `json:"syscode"`
	Account        string `json:"account"`
	Sign           string `json:"sign"`
	TradeNo        string `json:"trade_no"`
	Amount         string `json:"amount"`
	PayMode        string `json:"pay_mode"`
	CallbackUrl    string `json:"callback_url"`
	AuthAppId      string `json:"auth_app_id"`
	OpenId         string `json:"open_id"`
	Memo           string `json:"memo"`
	TerminalDevice string `json:"terminal_device"`
	IsRaw          string `json:"is_raw"`
	LimitCreditPay string `json:"limit_credit_pay"`
	BankId         string `json:"bank_id"`
	HbFqNum        string `json:"hb_fq_num"`
	ExpireDate     string `json:"expire_date"`
	TimeExpire     string `json:"time_expire"`
	TerminalIp     string `json:"terminal_ip"`
	AuthCode       string `json:"auth_code"`
}

// UnifiedPayResponse 响应结果
type UnifiedPayResponse struct {
	Code     int64  `json:"code,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Response struct {
		Response struct {
			Amount  string `json:"amount,omitempty"`
			AppId   string `json:"app_id,omitempty"`
			PayUrl  string `json:"pay_url,omitempty"`
			TransId string `json:"trans_id,omitempty"`
			UrlType string `json:"url_type,omitempty"`
		} `json:"response"`
		Sign string `json:"sign,omitempty"`
	} `json:"response"`
}