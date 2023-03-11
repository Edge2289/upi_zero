CREATE TABLE `split_bills_summary`
(
    `sbs_id`            int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `the_month`         varchar(10)      NOT NULL DEFAULT '' COMMENT '月份',
    `buyer_sub_acct_no` bigint(20)       NOT NULL DEFAULT '0' COMMENT '购方公司',
    `saler_sub_acct_no` bigint(20)       NOT NULL DEFAULT '0' COMMENT '销方公司',
    `the_type`          tinyint(1)       NOT NULL DEFAULT '1' COMMENT '类型{1:推广费,2:配送服务费,3:平台费,4:技术服务费,5:软件服务费,6:服务费}',
    `settle_in`         decimal(10, 2)   NOT NULL DEFAULT '0.00' COMMENT '月初结入',
    `into_amount`       decimal(10, 2)   NOT NULL DEFAULT '0.00' COMMENT '本月转入开票',
    `surplus_amount`    decimal(10, 2)   NOT NULL DEFAULT '0.00' COMMENT '剩余可开金额(含税)',
    `offline_amount`    decimal(10, 2)   NOT NULL DEFAULT '0.00' COMMENT '月末结出',
    `the_change_amount` decimal(10, 2)   NOT NULL DEFAULT '0.00' COMMENT '本月发生',
    `invoice_type`      tinyint(1)       NOT NULL DEFAULT '1' COMMENT '默认开票种类{1:普通电子发票,2:专用发票}',
    `serial_number`     int(4)           NOT NULL DEFAULT '1' COMMENT '流水号',
    `created_at`        int(11)          NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at`        int(11)          NOT NULL DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`sbs_id`),
    KEY `the_month` (`the_month`),
    KEY `buyer_sub_acct_no` (`buyer_sub_acct_no`),
    KEY `saler_sub_acct_no` (`saler_sub_acct_no`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 58
  DEFAULT CHARSET = utf8 COMMENT ='分账账务累计';

CREATE TABLE `split_bills_summary_apply`
(
    `sbsa_id`           int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `sbs_id`            int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分账账务累计id',
    `the_kp_month`      varchar(10)      NOT NULL DEFAULT '' COMMENT '开票日期',
    `buyer_sub_acct_no` bigint(20)       NOT NULL DEFAULT '0' COMMENT '购方公司',
    `saler_sub_acct_no` bigint(20)       NOT NULL DEFAULT '0' COMMENT '销方公司',
    `invoice_type`      tinyint(1)       NOT NULL DEFAULT '1' COMMENT '开票种类{1:蓝字发票,2:红字发票}',
    `invoice_code`      varchar(20)      NOT NULL DEFAULT '' COMMENT '发票代码',
    `invoice_no`        varchar(20)      NOT NULL DEFAULT '' COMMENT '发票号码',
    `tax_amount`        decimal(10, 2)   NOT NULL DEFAULT '0.00' COMMENT '含税金额',
    `created_at`        int(11)          NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_at`        int(11)          NOT NULL DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`sbsa_id`),
    KEY `sbs_id` (`sbs_id`),
    KEY `the_kp_month` (`the_kp_month`),
    KEY `buyer_sub_acct_no` (`buyer_sub_acct_no`),
    KEY `saler_sub_acct_no` (`saler_sub_acct_no`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='分账账务线下开票';

CREATE TABLE `split_bills_summary_log`
(
    `sbsl_id`     int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `sbs_id`      int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分账账务累计id',
    `operator`    varchar(10)      NOT NULL DEFAULT '' COMMENT '操作人',
    `the_content` varchar(30)      NOT NULL DEFAULT '' COMMENT '操作内容',
    `created_at`  int(11)          NOT NULL DEFAULT '0' COMMENT '创建时间',
    PRIMARY KEY (`sbsl_id`),
    KEY `sbs_id` (`sbs_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 101
  DEFAULT CHARSET = utf8 COMMENT ='分账账务累计操作记录表';

