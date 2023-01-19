CREATE TABLE `payment_info` (
                                `id` int(11) NOT NULL AUTO_INCREMENT,
                                `order_num` char(22) NOT NULL DEFAULT '' COMMENT '订单号',
                                `prices` decimal(2,0) NOT NULL DEFAULT '0' COMMENT '支付金额',
                                `order_user_id` int(11) NOT NULL DEFAULT '0' COMMENT '订单所属用户ID',
                                `payment_user_id` int(11) NOT NULL DEFAULT '0' COMMENT '支付用户ID(正常情况与order_user_id相同，当他人帮忙支付时不同)',
                                `payment_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '支付类型 1：微信 2：支付宝 3：银联 4：visa信用卡 5：账号余额',
                                `payment_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '支付ID',
                                `payment_code` varchar(255) NOT NULL DEFAULT '' COMMENT '支付串',
                                `payment_state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '支付状态 1:支付不成功 2：支付中 3：支付成功 4：余额不足',
                                `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                PRIMARY KEY (`id`),
                                KEY `user` (`order_user_id`) USING BTREE,
                                KEY `payment` (`payment_id`) USING BTREE,
                                KEY `order` (`order_num`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='支付记录表';