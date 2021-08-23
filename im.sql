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

 Date: 23/08/2021 17:13:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for im_friends
-- ----------------------------
DROP TABLE IF EXISTS `im_friends`;
CREATE TABLE `im_friends` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `m_id` int(11) DEFAULT NULL,
  `f_id` int(11) DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL COMMENT '0 未通过 1 已添加 2 已拒绝',
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of im_friends
-- ----------------------------
BEGIN;
COMMIT;

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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=157 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


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
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of im_groups
-- ----------------------------
BEGIN;
INSERT INTO `im_groups` VALUES (44, 30, 'Go学习交流群', '2021-08-23 13:59:05', '暂无', 'https://api.pltrue.top/400x400.png');
INSERT INTO `im_groups` VALUES (45, 30, 'vue社区', '2021-08-23 14:27:09', '暂无', 'https://api.pltrue.top/400x400.png');
COMMIT;

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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'admin', '', NULL, '$2a$10$ntZsWa6w97AdpGwlwJK.z.rA/b9inQfkOCwnQBFCJN2kt2DE7HRD6', NULL, '2021-07-01 14:54:46', NULL, 'https://cdn.learnku.com//uploads/communities/Y7fElYYwCFjTTXCdwPNW.png!/both/44x44', '', 0, NULL, 1, 1, NULL);
INSERT INTO `users` VALUES (2, 'hhhhh', NULL, NULL, '$2a$10$ntZsWa6w97AdpGwlwJK.z.rA/b9inQfkOCwnQBFCJN2kt2DE7HRD6', NULL, NULL, NULL, 'https://cdn.learnku.com//uploads/communities/sNljssWWQoW6J88O9G37.png!/both/44x44', NULL, 0, NULL, NULL, 1, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
