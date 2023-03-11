CREATE TABLE `spa_accounts`
(
    `id`                    bigint(20) unsigned                     NOT NULL AUTO_INCREMENT,
    `sub_acct_no`           bigint(20)                              NOT NULL DEFAULT '0' COMMENT '见证子账户的账号',
    `tran_net_member_code`  varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '交易网会员代码',
    `member_name`           varchar(30) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '会员名称',
    `member_global_type`    varchar(4) COLLATE utf8mb4_unicode_ci   NOT NULL COMMENT '会员证件类型 1:身份证 3:港澳台居民通行证（即回乡证） 4:中国护照 5:台湾居民来往大陆通行证（即台胞证） 19:外国护照 52:组织机构代码证 68:营业执照 73:统一社会信用代码',
    `member_global_id`      varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '会员证件号码',
    `user_nickname`         varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户昵称',
    `mobile`                varchar(12) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '手机号码',
    `member_property`       varchar(2) COLLATE utf8mb4_unicode_ci   NOT NULL DEFAULT 'SH' COMMENT '会员属性 SH-商户子账户（默认）',
    `reserved_msg`          varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '保留域',
    `status`                varchar(2) COLLATE utf8mb4_unicode_ci   NOT NULL COMMENT '状态：1:开立成功',
    `pass_time`             bigint(20)                              NOT NULL DEFAULT '0' COMMENT '账户开立时间',
    `type`                  varchar(2) COLLATE utf8mb4_unicode_ci   NOT NULL COMMENT '账户类型 1:经销商 2:供应商 3:经销商和供应商',
    `eicon_bank_branch_id`  varchar(14) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '会员绑定账户的开户行的超级网银行号',
    `member_acct_no`        varchar(32) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '会员绑定账户的账号 提现的银行卡',
    `cnaps_branch_id`       varchar(14) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '会员绑定账户的开户行的联行号',
    `acct_open_branch_name` varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '会员绑定账户的开户行名称',
    `bank_type`             varchar(1) COLLATE utf8mb4_unicode_ci   NOT NULL COMMENT '会员绑定账户的本他行类型 1：本行 2：他行',
    `created_at`            timestamp                               NULL     DEFAULT NULL,
    `updated_at`            timestamp                               NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `sub_acct_no` (`sub_acct_no`),
    KEY `member_name` (`member_name`),
    KEY `member_global_id` (`member_global_id`),
    KEY `pass_time` (`pass_time`),
    KEY `idx_tran_net_member_code` (`tran_net_member_code`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='见证子账表';


CREATE TABLE `spa_account_bindings`
(
    `id`          bigint(20) unsigned                    NOT NULL AUTO_INCREMENT,
    `user_id`     varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'users表id',
    `sub_acct_no` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'spa_accounts表sub_acct_no',
    `operator_id` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '操作人ID 0：系统操作',
    `created_at`  timestamp                              NULL DEFAULT NULL,
    `updated_at`  timestamp                              NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`),
    KEY `sub_acct_no` (`sub_acct_no`),
    KEY `operator_id` (`operator_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 103
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户关联见证子账表';

CREATE TABLE `spa_account_oa_entities`
(
    `id`          bigint(20) unsigned                    NOT NULL AUTO_INCREMENT,
    `sub_acct_no` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '见证子账户的账号',
    `entity_id`   varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '单位ID',
    `created_at`  timestamp                              NULL DEFAULT NULL,
    `updated_at`  timestamp                              NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `sub_acct_no` (`sub_acct_no`),
    KEY `entity_id` (`entity_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 329
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE TABLE `spa_account_histories`
(
    `id`                    bigint(20) unsigned                     NOT NULL AUTO_INCREMENT,
    `sub_acct_no`           varchar(32) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '见证子账户的账号',
    `tran_net_member_code`  varchar(32) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '交易网会员代码',
    `member_name`           varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '会员名称',
    `member_global_type`    varchar(4) COLLATE utf8mb4_unicode_ci   NOT NULL COMMENT '会员证件类型 1:身份证 3:港澳台居民通行证（即回乡证） 4:中国护照 5:台湾居民来往大陆通行证（即台胞证） 19:外国护照 52:组织机构代码证 68:营业执照 73:统一社会信用代码',
    `member_global_id`      varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '会员证件号码',
    `user_nickname`         varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户昵称',
    `mobile`                varchar(12) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '手机号码',
    `member_property`       varchar(2) COLLATE utf8mb4_unicode_ci   NOT NULL DEFAULT 'SH' COMMENT '会员属性 SH-商户子账户（默认）',
    `reserved_msg`          varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '保留域',
    `status`                varchar(2) COLLATE utf8mb4_unicode_ci   NOT NULL COMMENT '状态：1:开立成功',
    `pass_time`             varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '账户开立时间',
    `type`                  varchar(2) COLLATE utf8mb4_unicode_ci   NOT NULL COMMENT '账户类型 1:经销商 2:供应商 3:经销商和供应商',
    `eicon_bank_branch_id`  varchar(14) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '会员绑定账户的开户行的超级网银行号',
    `member_acct_no`        varchar(32) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '会员绑定账户的账号 提现的银行卡',
    `cnaps_branch_id`       varchar(14) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '会员绑定账户的开户行的联行号',
    `acct_open_branch_name` varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '会员绑定账户的开户行名称',
    `bank_type`             varchar(1) COLLATE utf8mb4_unicode_ci   NOT NULL COMMENT '会员绑定账户的本他行类型 1：本行 2：他行',
    `delete_time`           timestamp                               NULL     DEFAULT NULL COMMENT '删除时间',
    `created_at`            timestamp                               NULL     DEFAULT NULL,
    `updated_at`            timestamp                               NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `sub_acct_no` (`sub_acct_no`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2008
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='见证子账历史表';

CREATE TABLE `spa_cust_accounts`
(
    `id`          bigint(20) unsigned                     NOT NULL AUTO_INCREMENT,
    `sub_acct_no` varchar(32) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '见证功能子账号',
    `member_name` varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '见证功能子账名',
    `created_at`  timestamp                               NULL DEFAULT NULL,
    `updated_at`  timestamp                               NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `sub_acct_no` (`sub_acct_no`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 10
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='见证功能子账户表';

CREATE TABLE `spa_sup_account_bills`
(
    `id`                   bigint(20) unsigned                     NOT NULL AUTO_INCREMENT,
    `tran_time`            varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '交易日期 格式：YmdHis',
    `fund_summary_acct_no` varchar(32) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '汇总账号',
    `flag`                 varchar(2) COLLATE utf8mb4_unicode_ci   NOT NULL COMMENT '借贷标志 D：借；C：贷',
    `tran_amt`             varchar(32) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '交易金额',
    `book_bal`             varchar(32) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '账面余额',
    `member_acct_no`       varchar(32) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '对方账号',
    `member_acct_name`     varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '对方户名',
    `front_seq_no`         varchar(22) COLLATE utf8mb4_unicode_ci       DEFAULT NULL,
    `order_no`             varchar(64) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '订单号',
    `cnsmr_seq_no`         varchar(22) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '交易网流水号',
    `describe`             varchar(32) COLLATE utf8mb4_unicode_ci  NOT NULL COMMENT '摘要描述',
    `remark`               varchar(120) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '备注',
    `created_at`           timestamp                               NULL DEFAULT NULL,
    `updated_at`           timestamp                               NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `tran_time` (`tran_time`),
    KEY `fund_summary_acct_no` (`fund_summary_acct_no`),
    KEY `flag` (`flag`),
    KEY `front_seq_no` (`front_seq_no`),
    KEY `order_no` (`order_no`),
    KEY `cnsmr_seq_no` (`cnsmr_seq_no`),
    KEY `member_acct_no` (`member_acct_no`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='见证汇总账户明细文件表';