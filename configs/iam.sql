
CREATE DATABASE  IF NOT EXISTS `l-iam` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `l-iam`;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(128) DEFAULT '' NOT NULL COMMENT '昵称',
  `password` varchar(255) DEFAULT '' NOT NULL COMMENT '密码',
  `email` varchar(256) DEFAULT '' NOT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT '' NOT NULL COMMENT '手机号',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
  # 为username添加唯一约束会带来很多麻烦，故取消
  # UNIQUE KEY `idx_nickname` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `policy`;
CREATE TABLE `policy` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '策略ID',
  `name` varchar(128) DEFAULT '' NOT NULL COMMENT '名称',
  `created_at` datetime(3) NOT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL COMMENT '更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `policy` longtext COLLATE utf8mb4_unicode_ci COMMENT '昵称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci