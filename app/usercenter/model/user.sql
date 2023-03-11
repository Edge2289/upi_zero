CREATE TABLE `roles` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `name` varchar(255) NOT NULL,
                         `guard_name` varchar(255) NOT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `roles_name_guard_name_unique` (`name`,`guard_name`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8

CREATE TABLE `roles_resources` (
                                   `roles_resources_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
                                   `roles_id` bigint(20) unsigned NOT NULL COMMENT '功能资源id',
                                   `function_resources_id` int(11) unsigned NOT NULL COMMENT '功能资源id',
                                   `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '生成时间',
                                   `updated_at` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
                                   PRIMARY KEY (`roles_resources_id`),
                                   KEY `roles_id` (`roles_id`),
                                   KEY `function_resources_id` (`function_resources_id`)
) ENGINE=InnoDB AUTO_INCREMENT=680 DEFAULT CHARSET=utf8 COMMENT='角色-资源表'

CREATE TABLE `oa_users` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `staff_id` int(11) NOT NULL COMMENT 'oa员工id',
                            `project` varchar(64) NOT NULL COMMENT '所属项目',
                            `account` varchar(64) NOT NULL COMMENT '登录账号',
                            `staff_name` varchar(64) NOT NULL,
                            `entity_id` int(11) NOT NULL COMMENT '单位id',
                            `entity_name` varchar(64) NOT NULL COMMENT '单位名称',
                            `pre_entity_id` int(11) NOT NULL COMMENT '上级单位id',
                            `dw_id` int(11) NOT NULL COMMENT '单位下小组ID',
                            `dw_name` varchar(64) NOT NULL COMMENT '单位下小组名称',
                            `dw_type` varchar(64) NOT NULL COMMENT '单位类型',
                            `dw_level` varchar(64) NOT NULL COMMENT '单位级别',
                            `dw_type_id` int(11) NOT NULL COMMENT '单位类型ID',
                            `staff_type_id` int(11) NOT NULL COMMENT '员工类型ID',
                            `staff_type` varchar(64) NOT NULL COMMENT '员工类型',
                            `staff_level` varchar(64) NOT NULL COMMENT '员工级别',
                            `dq_id` int(11) NOT NULL COMMENT '大区ID',
                            `dq_name` varchar(64) NOT NULL COMMENT '大区名称',
                            `pq_id` int(11) NOT NULL COMMENT '片区id',
                            `pq_name` varchar(64) NOT NULL COMMENT '片区名称',
                            `qq_code` varchar(64) NOT NULL COMMENT 'qq号',
                            `avatar_img` varchar(190) NOT NULL DEFAULT '',
                            `identity_card` varchar(64) NOT NULL,
                            `created` int(11) NOT NULL,
                            `user_id` int(11) NOT NULL COMMENT '用户ID,关联user表id',
                            `created_at` timestamp NULL DEFAULT NULL,
                            `updated_at` timestamp NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `staff_id` (`staff_id`),
                            KEY `account` (`account`),
                            KEY `entity_id` (`entity_id`),
                            KEY `pre_entity_id` (`pre_entity_id`),
                            KEY `dw_id` (`dw_id`),
                            KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 COMMENT='OA用户表'

CREATE TABLE `model_has_roles` (
                                   `role_id` bigint(20) unsigned NOT NULL,
                                   `model_type` varchar(255) NOT NULL,
                                   `model_id` bigint(20) unsigned NOT NULL,
                                   PRIMARY KEY (`role_id`,`model_id`,`model_type`),
                                   KEY `model_has_roles_model_id_model_type_index` (`model_id`,`model_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

CREATE TABLE `users_operation_log` (
                                       `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '操作记录id',
                                       `users_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户表id',
                                       `operation_time` int(11) NOT NULL DEFAULT '0' COMMENT '操作时间',
                                       `operator` varchar(20) NOT NULL DEFAULT '' COMMENT '操作人',
                                       `operator_id` int(11) NOT NULL DEFAULT '0' COMMENT '操作人id',
                                       `oper_type` varchar(10) NOT NULL DEFAULT '' COMMENT '操作类型',
                                       `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注信息',
                                       `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '生成时间',
                                       PRIMARY KEY (`id`),
                                       KEY `users_id` (`users_id`)
) ENGINE=InnoDB AUTO_INCREMENT=656 DEFAULT CHARSET=utf8 COMMENT='用户操作记录表'

CREATE TABLE `users` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `name` varchar(255) NOT NULL,
                         `email` varchar(255) NOT NULL,
                         `email_verified_at` timestamp NULL DEFAULT NULL,
                         `password` varchar(255) NOT NULL,
                         `remember_token` varchar(100) DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8

CREATE TABLE `user_unfinished` (
                                   `id` bigint(20) NOT NULL AUTO_INCREMENT,
                                   `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
                                   `content` varchar(255) NOT NULL DEFAULT '' COMMENT '通知的内容',
                                   `is_know` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否知道 1：知道；0：不知道；',
                                   `created_at` int(14) NOT NULL DEFAULT '0' COMMENT '创建时间',
                                   `updated_at` int(14) NOT NULL DEFAULT '0' COMMENT '更新时间',
                                   PRIMARY KEY (`id`),
                                   KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3749 DEFAULT CHARSET=utf8 COMMENT='待办事项通知表'

