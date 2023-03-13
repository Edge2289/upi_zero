package model


type SpaAccountsInfo struct {
	Id uint64 `json:"id,omitempty"`
	SubAcctNo int64 `json:"sub_acct_no,omitempty"` // 见证子账户的账号
	TranNetMemberCode string `json:"tran_net_member_code,omitempty"` // 交易网会员代码
	MemberName string `json:"member_name,omitempty"` // 会员名称
	MemberGlobalType string `json:"member_global_type,omitempty"` // 会员证件类型 1:身份证 3:港澳台居民通行证（即回乡证） 4:中国护照 5:台湾居民来往大陆通行证（即台胞证） 19:外国护照 52:组织机构代码证 68:营业执照 73:统一社会信用代码
	MemberGlobalId string `json:"member_global_id,omitempty"` // 会员证件号码
	UserNickname string `json:"user_nickname,omitempty"` // 用户昵称
	Mobile string `json:"mobile,omitempty"` // 手机号码
	MemberProperty string `json:"member_property,omitempty"` // 会员属性 SH-商户子账户（默认）
	ReservedMsg string `json:"reserved_msg,omitempty"` // 保留域
	Status string `json:"status,omitempty"` // 状态：1:开立成功
	PassTime int64 `json:"pass_time,omitempty"` // 账户开立时间
	Type string `json:"type,omitempty"` // 账户类型 1:经销商 2:供应商 3:经销商和供应商
	EiconBankBranchId string `json:"eicon_bank_branch_id,omitempty"` // 会员绑定账户的开户行的超级网银行号
	MemberAcctNo string `json:"member_acct_no,omitempty"` // 会员绑定账户的账号 提现的银行卡
	CnapsBranchId string `json:"cnaps_branch_id,omitempty"` // 会员绑定账户的开户行的联行号
	AcctOpenBranchName string `json:"acct_open_branch_name,omitempty"` // 会员绑定账户的开户行名称
	BankType string `json:"bank_type,omitempty"` // 会员绑定账户的本他行类型 1：本行 2：他行
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type SpaAccountBindingsInfo struct {
	Id uint64 `json:"id,omitempty"`
	UserId string `json:"user_id,omitempty"` // users表id
	SubAcctNo string `json:"sub_acct_no,omitempty"` // spa_accounts表sub_acct_no
	OperatorId string `json:"operator_id,omitempty"` // 操作人ID 0：系统操作
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func (s *SpaAccountBindingsInfo) TableName() string {
	return "spa_account_bindings"
}

type SpaAccountOaEntitiesInfo struct {
	Id uint64 `json:"id,omitempty"`
	SubAcctNo string `json:"sub_acct_no,omitempty"` // 见证子账户的账号
	EntityId string `json:"entity_id,omitempty"` // 单位ID
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type SpaAccountHistoriesInfo struct {
	Id uint64 `json:"id,omitempty"`
	SubAcctNo string `json:"sub_acct_no,omitempty"` // 见证子账户的账号
	TranNetMemberCode string `json:"tran_net_member_code,omitempty"` // 交易网会员代码
	MemberName string `json:"member_name,omitempty"` // 会员名称
	MemberGlobalType string `json:"member_global_type,omitempty"` // 会员证件类型 1:身份证 3:港澳台居民通行证（即回乡证） 4:中国护照 5:台湾居民来往大陆通行证（即台胞证） 19:外国护照 52:组织机构代码证 68:营业执照 73:统一社会信用代码
	MemberGlobalId string `json:"member_global_id,omitempty"` // 会员证件号码
	UserNickname string `json:"user_nickname,omitempty"` // 用户昵称
	Mobile string `json:"mobile,omitempty"` // 手机号码
	MemberProperty string `json:"member_property,omitempty"` // 会员属性 SH-商户子账户（默认）
	ReservedMsg string `json:"reserved_msg,omitempty"` // 保留域
	Status string `json:"status,omitempty"` // 状态：1:开立成功
	PassTime string `json:"pass_time,omitempty"` // 账户开立时间
	Type string `json:"type,omitempty"` // 账户类型 1:经销商 2:供应商 3:经销商和供应商
	EiconBankBranchId string `json:"eicon_bank_branch_id,omitempty"` // 会员绑定账户的开户行的超级网银行号
	MemberAcctNo string `json:"member_acct_no,omitempty"` // 会员绑定账户的账号 提现的银行卡
	CnapsBranchId string `json:"cnaps_branch_id,omitempty"` // 会员绑定账户的开户行的联行号
	AcctOpenBranchName string `json:"acct_open_branch_name,omitempty"` // 会员绑定账户的开户行名称
	BankType string `json:"bank_type,omitempty"` // 会员绑定账户的本他行类型 1：本行 2：他行
	DeleteTime string `json:"delete_time,omitempty"` // 删除时间
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type SpaCustAccountsInfo struct {
	Id uint64 `json:"id,omitempty"`
	SubAcctNo string `json:"sub_acct_no,omitempty"` // 见证功能子账号
	MemberName string `json:"member_name,omitempty"` // 见证功能子账名
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type SpaSupAccountBillsInfo struct {
	Id uint64 `json:"id,omitempty"`
	TranTime string `json:"tran_time,omitempty"` // 交易日期 格式：YmdHis
	FundSummaryAcctNo string `json:"fund_summary_acct_no,omitempty"` // 汇总账号
	Flag string `json:"flag,omitempty"` // 借贷标志 D：借；C：贷
	TranAmt string `json:"tran_amt,omitempty"` // 交易金额
	BookBal string `json:"book_bal,omitempty"` // 账面余额
	MemberAcctNo string `json:"member_acct_no,omitempty"` // 对方账号
	MemberAcctName string `json:"member_acct_name,omitempty"` // 对方户名
	FrontSeqNo string `json:"front_seq_no,omitempty"`
	OrderNo string `json:"order_no,omitempty"` // 订单号
	CnsmrSeqNo string `json:"cnsmr_seq_no,omitempty"` // 交易网流水号
	Describe string `json:"describe,omitempty"` // 摘要描述
	Remark string `json:"remark,omitempty"` // 备注
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}