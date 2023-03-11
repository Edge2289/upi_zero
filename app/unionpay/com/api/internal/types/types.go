// Code generated by goctl. DO NOT EDIT.
package types

type PayReq struct {
	TradeNo        string `form:"tradeNo" validate:"required,max=28"`
	Amount         string `form:"amount" validate:"required,max=123"`
	PayMode        string `form:"payMode" validate:"required,max=30"`
	CallbackUrl    string `form:"callbackUrl,optional" validate:"max=128"`
	AuthAppId      string `form:"authAppId,optional" validate:"max=32"`
	OpenId         string `form:"openId,optional" validate:"max=32"`
	Memo           string `form:"memo,optional" validate:"max=30"`
	TerminalDevice string `form:"terminalDevice,optional" validate:"max=2"`
	IsRaw          string `form:"isRaw,optional" validate:"max=1"`
	LimitCreditPay string `form:"limitCreditPay,optional" validate:"max=10"`
	HbFqNum        string `form:"hbFqNum,optional" validate:"max=10"`
	BankId         string `form:"bankId,optional" validate:"max=12"`
	ExpireDate     string `form:"expireDate,optional" validate:"max=14"`
	TimeExpire     string `form:"timeExpire,optional" validate:"max=10"`
	TerminalIp     string `form:"terminalIp,optional" validate:"max=32"`
	AuthCode       string `form:"authCode,optional" validate:"max=32"`
}

type PayResp struct {
	Data interface{} `json:"data"`
}

type QueryReq struct {
	TradeNo string `form:"tradeNo"`
}

type QueryResp struct {
	Amount     string `json:"amount"`
	App_id     string `json:"app_id"`
	Result     string `json:"result"`
	Trade_date string `json:"trade_date"`
	Trans_id   string `json:"trans_id"`
}

type ModesResp struct {
	Modes interface{} `json:"modes"`
}

type CancelReq struct {
	TradeNo  string `form:"tradeNo"`
	CancelNo string `form:"cancelNo"`
	Remark   string `form:"remark"`
}

type CancelResp struct {
	App_id   string `json:"appId"`
	Result   string `json:"result"`
	Trans_id string `json:"trans_id"`
}

type NotifyTradeReq struct {
	Account      string `json:"account"`
	Amount       string `json:"amount"`
	AppID        string `json:"app_id"`
	Fee          string `json:"fee"`
	PayMode      string `json:"pay_mode"`
	Rate         string `json:"rate"`
	Result       string `json:"result"`
	SettleAmount string `json:"settle_amount"`
	Sign         string `json:"sign"`
	SuccTime     string `json:"succ_time"`
	TransID      string `json:"trans_id"`
}
