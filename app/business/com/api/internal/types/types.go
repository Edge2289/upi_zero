// Code generated by goctl. DO NOT EDIT.
package types

type SplitBillsSummaryListReq struct {
	Page              string `form:"page"`
	Per_page          string `form:"perPage"`
	buyerCompanyName string `form:"buyer_company_name"`
	TheMonthStart     string `form:"the_month_start"`
	TheMonthEnd       string `form:"the_month_end"`
}

type SplitBillsSummaryListResp struct {
	List []SplitBillsSummaryListmap `json:"list"`
}

type SplitBillsSummaryListmap struct {
	SbsID            int    `json:"sbs_id"`
	TheMonth         string `json:"the_month"`
	BuyerSubAcctNo   int64  `json:"buyer_sub_acct_no"`
	SalerSubAcctNo   int64  `json:"saler_sub_acct_no"`
	TheType          string `json:"the_type"`
	SettleIn         string `json:"settle_in"`
	IntoAmount       string `json:"into_amount"`
	SurplusAmount    string `json:"surplus_amount"`
	OfflineAmount    string `json:"offline_amount"`
	TheChangeAmount  int    `json:"the_change_amount"`
	InvoiceType      string `json:"invoice_type"`
	SerialNumber     int    `json:"serial_number"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	BuyerCompanyName string `json:"buyer_company_name"`
	SalerCompanyName string `json:"saler_company_name"`
}

type NullResp struct {
}
