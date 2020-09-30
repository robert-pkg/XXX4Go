
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE IF NOT EXISTS `sys_user` (
  `sys_user_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键，自增长',
  `mobile` varchar(40) COLLATE utf8mb4_bin NOT NULL COMMENT '手机号码',
  `name` varchar(40) COLLATE utf8mb4_bin NOT NULL COMMENT '名字',
  `is_enabled` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否有效，默认 1启用，0禁用',
  `is_deleted` tinyint(4) DEFAULT '1' COMMENT '是否删除，1未删除，0已删除',
  `created_tm` datetime(6) DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updated_tm` datetime(6) DEFAULT '0000-00-00 00:00:00.000000' ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '修改时间',
  PRIMARY KEY (`sys_user_id`),
  UNIQUE KEY `uinx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='系统用户';

DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE IF NOT EXISTS `sys_user_role` (
  `sys_user_role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键，自增长',
  `sys_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '系统用户ID',
  `sys_role_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '系统角色ID',
  `is_deleted` tinyint(4) DEFAULT '1' COMMENT '是否删除，1未删除，0已删除',
  `created_tm` datetime(6) DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updated_tm` datetime(6) DEFAULT '0000-00-00 00:00:00.000000' ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '修改时间',
  PRIMARY KEY (`sys_user_role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COMMENT='系统用户角色表';

DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE IF NOT EXISTS `sys_role` (
  `sys_role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键，自增长',
  `name` varchar(40) COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `description` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '描述',
  `scope` tinyint(4) DEFAULT '1' COMMENT '权限范围 1全部 2自定义',
  `is_deleted` tinyint(4) DEFAULT '1' COMMENT '是否删除，1未删除，0已删除',
  `created_tm` datetime(6) DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updated_tm` datetime(6) DEFAULT '0000-00-00 00:00:00.000000' ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '修改时间',
  PRIMARY KEY (`sys_role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='系统角色表';

DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE IF NOT EXISTS `sys_role_menu` (
  `sys_role_menu_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键，自增长',
  `sys_role_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '系统角色ID',
  `sys_menu_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '系统菜单ID',
  `is_deleted` tinyint(4) DEFAULT '1' COMMENT '是否删除，1未删除，0已删除',
  `created_tm` datetime(6) DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updated_tm` datetime(6) DEFAULT '0000-00-00 00:00:00.000000' ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '修改时间',
  PRIMARY KEY (`sys_role_menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COMMENT='系统角色菜单关联表';

DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE IF NOT EXISTS `sys_menu` (
  `sys_menu_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键，自增长',
  `parent_menu_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '上级菜单ID',
  `type` int(11) NOT NULL DEFAULT '1' COMMENT '类型 1菜单 2按钮',
  `name` varchar(40) COLLATE utf8mb4_bin NOT NULL COMMENT '名称',
  `icon_url` varchar(250) COLLATE utf8mb4_bin DEFAULT '' COMMENT 'type为1时有效, 菜单图标url',
  `path` varchar(250) COLLATE utf8mb4_bin DEFAULT '' COMMENT 'type为1时有效, 菜单路径',
  `button_uid` varchar(50) COLLATE utf8mb4_bin DEFAULT '' COMMENT 'type为2时有效, 按钮唯一标识ID',
  `sort_idx` int(11) DEFAULT '0' COMMENT '同级顺序',
  `remark` varchar(250) COLLATE utf8mb4_bin DEFAULT '' COMMENT '备注',
  `is_deleted` tinyint(4) DEFAULT '1' COMMENT '是否删除，1未删除，0已删除',
  `created_by` bigint(20) DEFAULT '0' COMMENT '创建人',
  `created_tm` datetime(6) DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updated_by` bigint(20) DEFAULT '0' COMMENT '修改人',
  `updated_tm` datetime(6) DEFAULT '0000-00-00 00:00:00.000000' ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '修改时间',
  PRIMARY KEY (`sys_menu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='系统角色菜单关联表';

insert into sys_user (sys_user_id,mobile,name) values (1,'18000000000', '超级管理员');
insert into sys_role (sys_role_id,name,description,scope) values (1, '超级管理员', '-', 1);
insert into sys_menu (sys_menu_id,type,name,sort_idx) values (1, 1, 'xxx运营平台', 1);
