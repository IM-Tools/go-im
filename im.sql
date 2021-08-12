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

 Date: 05/07/2021 11:01:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of im_messages
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
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
  `status` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (30, 'admin', '', NULL, '$2a$10$ntZsWa6w97AdpGwlwJK.z.rA/b9inQfkOCwnQBFCJN2kt2DE7HRD6', NULL, '2021-07-01 14:54:46', NULL, 'https://tvax2.sinaimg.cn/crop.0.0.1002.1002.180/006pP2Laly8gqcj17wce9j30ru0ru0up.jpg?KID=imgbed,tva&Expires=1624513271&ssig=lmrNTiDUPy', '', 0, NULL, 1, 0);
INSERT INTO `users` VALUES (31, 'hhhhh', NULL, NULL, '$2a$10$ntZsWa6w97AdpGwlwJK.z.rA/b9inQfkOCwnQBFCJN2kt2DE7HRD6', NULL, NULL, NULL, 'https://tvax1.sinaimg.cn/default/images/default_avatar_male_180.gif?KID=imgbed,tva&Expires=1624980396&ssig=0Vi0pF4H4D', NULL, 0, NULL, NULL, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
