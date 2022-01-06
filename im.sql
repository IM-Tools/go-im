/*
 Navicat Premium Data Transfer

 Source Server         : laradock_mysql
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : localhost:3306
 Source Schema         : im

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 04/01/2022 14:26:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for im_friend_records
-- ----------------------------
DROP TABLE IF EXISTS `im_friend_records`;
CREATE TABLE `im_friend_records` (
                                     `id` int(11) NOT NULL AUTO_INCREMENT,
                                     `user_id` int(11) NOT NULL,
                                     `f_id` int(11) NOT NULL,
                                     `status` tinyint(1) DEFAULT NULL COMMENT '0 等待通过 1 已通过 2 已拒绝',
                                     `created_at` timestamp NULL DEFAULT NULL,
                                     `information` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求信息',
                                     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for im_friends
-- ----------------------------
DROP TABLE IF EXISTS `im_friends`;
CREATE TABLE `im_friends` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `m_id` int(11) DEFAULT NULL,
                              `f_id` int(11) DEFAULT NULL,
                              `created_at` timestamp NULL DEFAULT NULL,
                              `note` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
                              `top_time` timestamp NULL DEFAULT NULL,
                              `status` tinyint(1) DEFAULT '0' COMMENT '0.未置顶 1.已置顶',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for im_group_messages
-- ----------------------------
DROP TABLE IF EXISTS `im_group_messages`;
CREATE TABLE `im_group_messages` (
                                     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                                     `msg_type` tinyint(1) NOT NULL COMMENT '消息类型 1.文本消息 2.图文消息 3.语音消息',
                                     `msg` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '消息内容',
                                     `group_id` int(11) NOT NULL COMMENT '群聊id',
                                     `from_id` int(11) NOT NULL COMMENT '消息发送人',
                                     `created_at` timestamp NULL DEFAULT NULL,
                                     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for im_group_users
-- ----------------------------
DROP TABLE IF EXISTS `im_group_users`;
CREATE TABLE `im_group_users` (
                                  `id` int(11) NOT NULL AUTO_INCREMENT,
                                  `user_id` int(11) NOT NULL,
                                  `created_at` timestamp NULL DEFAULT NULL,
                                  `group_id` int(11) DEFAULT NULL,
                                  `remark` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
                                  `avatar` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
                                  `name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
                                  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=178 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for im_groups
-- ----------------------------
DROP TABLE IF EXISTS `im_groups`;
CREATE TABLE `im_groups` (
                             `id` int(11) NOT NULL AUTO_INCREMENT,
                             `user_id` int(11) DEFAULT NULL,
                             `group_name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
                             `created_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                             `info` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
                             `group_avatar` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
                             PRIMARY KEY (`id`),
                             KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for im_messages
-- ----------------------------
DROP TABLE IF EXISTS `im_messages`;
CREATE TABLE `im_messages` (
                               `id` int(11) NOT NULL AUTO_INCREMENT,
                               `msg` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL,
                               `created_at` timestamp NULL DEFAULT NULL,
                               `from_id` int(11) DEFAULT NULL,
                               `to_id` int(11) DEFAULT NULL,
                               `channel` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
                               `is_read` tinyint(1) DEFAULT NULL COMMENT '0 未读 1已读',
                               `msg_type` tinyint(1) DEFAULT '1',
                               `channel_type` tinyint(1) DEFAULT '1' COMMENT '1.好友 2.群聊',
                               `status` tinyint(1) DEFAULT NULL,
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=158 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for im_sessions
-- ----------------------------
DROP TABLE IF EXISTS `im_sessions`;
CREATE TABLE `im_sessions` (
                               `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '会话表',
                               `m_id` int(11) NOT NULL,
                               `f_id` int(11) NOT NULL,
                               `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               `top_status` tinyint(1) DEFAULT '0' COMMENT '0.否 1.是',
                               `top_time` timestamp NULL DEFAULT NULL,
                               `note` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '备注',
                               `channel_type` tinyint(1) DEFAULT '0' COMMENT '0.单聊 1.群聊',
                               `name` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '会话名称',
                               `avatar` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL COMMENT '会话头像',
                               `status` tinyint(1) DEFAULT '0' COMMENT '会话状态 0.正常 1.禁用',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for im_users
-- ----------------------------
DROP TABLE IF EXISTS `im_users`;
CREATE TABLE `im_users` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                            `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                            `email_verified_at` timestamp NULL DEFAULT NULL,
                            `password` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                            `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                            `created_at` timestamp NULL DEFAULT NULL,
                            `updated_at` timestamp NULL DEFAULT NULL,
                            `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '头像',
                            `oauth_id` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方id',
                            `bound_oauth` tinyint(1) DEFAULT '0' COMMENT '1\\github 2\\gitee',
                            `deleted_at` timestamp NULL DEFAULT NULL,
                            `oauth_type` tinyint(1) DEFAULT NULL COMMENT '1.微博 2.github',
                            `status` tinyint(1) DEFAULT '0' COMMENT '0 离线 1 在线',
                            `bio` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '用户简介',
                            `sex` tinyint(1) DEFAULT '0' COMMENT '0 未知 1.男 2.女',
                            `client_type` tinyint(1) DEFAULT NULL COMMENT '1.web 2.pc 3.app',
                            `age` int(3) DEFAULT NULL,
                            `last_login_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

SET FOREIGN_KEY_CHECKS = 1;
