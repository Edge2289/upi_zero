package capitalManager

type SplitBillsSummary struct {
	SbsId uint32 `json:"sbs_id,omitempty"` // 自增id
	TheMonth string `json:"the_month,omitempty"` // 月份
	BuyerSubAcctNo int64 `json:"buyer_sub_acct_no,omitempty"` // 购方公司
	SalerSubAcctNo int64 `json:"saler_sub_acct_no,omitempty"` // 销方公司
	TheType int8 `json:"the_type,omitempty"` // 类型{1:推广费,2:配送服务费,3:平台费,4:技术服务费,5:软件服务费,6:服务费}
	SettleIn string `json:"settle_in,omitempty"` // 月初结入
	IntoAmount string `json:"into_amount,omitempty"` // 本月转入开票
	SurplusAmount string `json:"surplus_amount,omitempty"` // 剩余可开金额(含税)
	OfflineAmount string `json:"offline_amount,omitempty"` // 月末结出
	TheChangeAmount string `json:"the_change_amount,omitempty"` // 本月发生
	InvoiceType int8 `json:"invoice_type,omitempty"` // 默认开票种类{1:普通电子发票,2:专用发票}
	SerialNumber int32 `json:"serial_number,omitempty"` // 流水号
	CreatedAt int32 `json:"created_at,omitempty"` // 创建时间
	UpdatedAt int32 `json:"updated_at,omitempty"` // 更新时间
}

func (s SplitBillsSummary) TableName() string {
	return "split_bills_summary"
}

type SplitBillsSummaryApply struct {
	SbsaId uint32 `json:"sbsa_id,omitempty"` // 自增id
	SbsId uint32 `json:"sbs_id,omitempty"` // 分账账务累计id
	TheKpMonth string `json:"the_kp_month,omitempty"` // 开票日期
	BuyerSubAcctNo int64 `json:"buyer_sub_acct_no,omitempty"` // 购方公司
	SalerSubAcctNo int64 `json:"saler_sub_acct_no,omitempty"` // 销方公司
	InvoiceType int8 `json:"invoice_type,omitempty"` // 开票种类{1:蓝字发票,2:红字发票}
	InvoiceCode string `json:"invoice_code,omitempty"` // 发票代码
	InvoiceNo string `json:"invoice_no,omitempty"` // 发票号码
	TaxAmount string `json:"tax_amount,omitempty"` // 含税金额
	CreatedAt int32 `json:"created_at,omitempty"` // 创建时间
	UpdatedAt int32 `json:"updated_at,omitempty"` // 更新时间
}

type SplitBillsSummaryLog struct {
	SbslId uint32 `json:"sbsl_id,omitempty"` // 自增id
	SbsId uint32 `json:"sbs_id,omitempty"` // 分账账务累计id
	Operator string `json:"operator,omitempty"` // 操作人
	TheContent string `json:"the_content,omitempty"` // 操作内容
	CreatedAt int32 `json:"created_at,omitempty"` // 创建时间
}
