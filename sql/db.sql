
create database xxx;

DROP TABLE IF EXISTS `sms_login_vcode`;
CREATE TABLE IF NOT EXISTS `sms_login_vcode`(
  `pid` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号码',
  `vcode` varchar(20) NOT NULL DEFAULT '' COMMENT '短信验证码',
  `effect_ts` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '开始生效时间戳',
  `expire_ts` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '过期时间戳',
  `is_deleted` tinyint(4) DEFAULT '0' COMMENT '记录逻辑状态，是否被删除，1未删除， 2已删除',
  `created_tm` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  PRIMARY KEY (`pid`),
  KEY `inx_m_v` (`mobile`,`vcode`)
) ENGINE=InnoDB AUTO_INCREMENT=1041 DEFAULT CHARSET=utf8mb4 COMMENT='登录短信验证码'

DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键，自增长',
  `mobile` varchar(40) COLLATE utf8mb4_bin NOT NULL COMMENT '手机号码',
  `is_deleted` tinyint(4) DEFAULT '1' COMMENT '是否删除，1 未删除，2 已删除',
  `created_tm` datetime(6) DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updated_tm` datetime(6) DEFAULT '0000-00-00 00:00:00.000000' ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '修改时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `uinx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COMMENT='会员信息表';

DROP TABLE IF EXISTS `user_token`;
CREATE TABLE IF NOT EXISTS `user_token`(
  `pid` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `device_type` bigint(20) NOT NULL DEFAULT '1' COMMENT '设备类型. 1:web, 2:android, 3:ios',
  `token` varchar(64) NOT NULL DEFAULT '' COMMENT '令牌',
  `expire_ts` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '过期时间戳',
  `is_deleted` tinyint(4) DEFAULT '1' COMMENT '是否被删除，1未删除， 2已删除',
  `created_tm` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  PRIMARY KEY (`pid`),
   KEY `inx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COMMENT='用户令牌'

select * from user_token;




