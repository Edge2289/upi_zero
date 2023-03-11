package trade

import (
	"context"
	"upi/common/tool"
)

// BeforeUnionPay 支付前置操作
func (w *WebPay) BeforeUnionPay(ctx context.Context, req *UnifiedPayRequest) error {

	req.Syscode = w.Config.SysCode
	req.Account = w.Config.Account

	signMap := tool.Ksort(req.ToMap())
	var dataStr string
	for v, k := range signMap {
		// 为空的不参与签名
		if v == "" {
			continue
		}
		dataStr = dataStr + k + "=" + v + "&"
	}
	f := dataStr[0:len(dataStr) - 1]
	// 剔除空值作为传参
	sign, err := tool.RSASign([]byte(f), w.Config.PrivateKey)
	if err != nil {
		return err
	}
	req.Sign = sign

	return nil
}

// ToMap 结构体转map
func (req UnifiedPayRequest) ToMap() map[string]string {
	var signMap = make(map[string]string)
	signMap["syscode"] = req.Syscode
	signMap["account"] = req.Account
	signMap["trade_no"] = req.TradeNo
	signMap["amount"] = req.Amount
	signMap["pay_mode"] = req.PayMode
	signMap["callback_url"] = req.CallbackUrl
	signMap["auth_app_id"] = req.AuthAppId
	signMap["open_id"] = req.OpenId
	signMap["memo"] = req.Memo
	signMap["terminal_device"] = req.TerminalDevice
	signMap["is_raw"] = req.IsRaw
	signMap["limit_credit_pay"] = req.LimitCreditPay
	signMap["bank_id"] = req.BankId
	signMap["hb_fq_num"] = req.HbFqNum
	signMap["expire_date"] = req.ExpireDate
	signMap["time_expire"] = req.TimeExpire
	signMap["terminal_ip"] = req.TerminalIp
	signMap["auth_code"] = req.AuthCode

	return signMap
}