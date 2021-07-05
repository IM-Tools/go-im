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
BEGIN;
INSERT INTO `im_messages` VALUES (1, '测试', '2021-07-02 18:24:30', 7, 29, 'channel_7_29');
INSERT INTO `im_messages` VALUES (2, '你好啊', '2021-07-02 18:24:39', 7, 29, 'channel_7_29');
INSERT INTO `im_messages` VALUES (3, '你哈', '2021-07-02 18:24:46', 29, 7, 'channel_29_7');
INSERT INTO `im_messages` VALUES (4, ' 😊', '2021-07-02 18:24:52', 29, 7, 'channel_29_7');
INSERT INTO `im_messages` VALUES (5, '你还', '2021-07-02 18:25:04', 7, 29, 'channel_7_29');
INSERT INTO `im_messages` VALUES (6, ' 🤔', '2021-07-02 18:25:24', 7, 29, 'channel_7_29');
INSERT INTO `im_messages` VALUES (7, '测试', '2021-07-02 18:44:22', 29, 7, 'channel_29_7');
INSERT INTO `im_messages` VALUES (8, ' 🥰', '2021-07-02 18:44:33', 29, 7, 'channel_29_7');
INSERT INTO `im_messages` VALUES (9, '3333', '2021-07-04 21:45:40', 31, 30, 'channel_31_30');
INSERT INTO `im_messages` VALUES (10, '3333', '2021-07-04 21:45:54', 30, 31, 'channel_30_31');
INSERT INTO `im_messages` VALUES (11, ' 😂', '2021-07-04 21:46:01', 30, 31, 'channel_30_31');
INSERT INTO `im_messages` VALUES (12, '你好喔', '2021-07-04 21:46:06', 30, 31, 'channel_30_31');
INSERT INTO `im_messages` VALUES (13, '为啥看不到你的声音', '2021-07-04 21:46:13', 30, 31, 'channel_30_31');
INSERT INTO `im_messages` VALUES (14, '222', '2021-07-04 21:55:38', 30, 31, 'channel_30_31');
INSERT INTO `im_messages` VALUES (15, '222', '2021-07-04 22:10:11', 31, 30, 'channel_31_30');
INSERT INTO `im_messages` VALUES (16, '222', '2021-07-04 22:10:11', 31, 30, 'channel_31_30');
INSERT INTO `im_messages` VALUES (17, ' 😀', '2021-07-04 22:10:49', 30, 31, 'channel_30_31');
INSERT INTO `im_messages` VALUES (18, 'gg', '2021-07-04 22:10:56', 31, 30, 'channel_31_30');
INSERT INTO `im_messages` VALUES (19, 'gg', '2021-07-04 22:10:56', 31, 30, 'channel_31_30');
INSERT INTO `im_messages` VALUES (20, '调整', '2021-07-05 10:58:06', 30, 31, 'channel_30_31');
INSERT INTO `im_messages` VALUES (21, '呵呵', '2021-07-05 10:58:12', 30, 31, 'channel_30_31');
INSERT INTO `im_messages` VALUES (22, '测试', '2021-07-05 10:58:20', 30, 31, 'channel_30_31');
COMMIT;

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
INSERT INTO `users` VALUES (7, 'Hi叫我李荣浩', '', NULL, '$2a$10$yBlCzxlErLHcvMlVKH3TtOr9MKhJoAAIJjLdcWMo5nWGzWeXb13r2', NULL, '2021-06-24 10:41:11', NULL, 'https://tvax2.sinaimg.cn/crop.0.0.1002.1002.180/006pP2Laly8gqcj17wce9j30ru0ru0up.jpg?KID=imgbed,tva&Expires=1624513271&ssig=lmrNTiDUPy', '5878370732', 0, NULL, 1, 0);
INSERT INTO `users` VALUES (19, '五西路', '', NULL, '$2y$10$BwsPNYO.rPGqmPNxBAK/QevuAx3Izqz5jZTA5EYNmO/VL.Fmjp53.', NULL, '2021-06-09 10:50:56', '2021-06-09 10:50:56', 'https://tva2.sinaimg.cn/crop.0.0.180.180.180/6e02f304jw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1623217855&ssig=%2F%2BcSVB5SU3', '1845687044', 0, NULL, NULL, 0);
INSERT INTO `users` VALUES (20, '不要告诉我什么不行', '', NULL, '$2y$10$9qo/YoUxfJjU/TiBLshU5.IhfG.UQzc.uxTGoSqzF1GQfUnC6UD7u', NULL, '2021-06-09 11:14:52', '2021-06-09 11:14:52', 'https://tvax1.sinaimg.cn/crop.837.0.1067.1067.180/5322c24fly8fix6ugirgmj21hd0u0abi.jpg?KID=imgbed,tva&Expires=1623219291&ssig=%2BAENnk4B76', '1394786895', 0, NULL, NULL, 0);
INSERT INTO `users` VALUES (21, '用户6970206330', '', NULL, '$2y$10$OxR0Pyqy7uKNFOiRKfSbM.tYOxOCdKlsK/nFiY4vjj9dq8wa0hqTi', NULL, '2021-06-11 11:44:31', '2021-06-11 11:44:31', 'https://tvax3.sinaimg.cn/crop.0.0.132.132.180/007BIh4Kly4fzfjulsz5ij303o03odfn.jpg?KID=imgbed,tva&Expires=1623393871&ssig=xv%2Bw3WXv%2Fn', '6970206330', 0, NULL, NULL, 0);
INSERT INTO `users` VALUES (22, '解心释神莫然无魂', '', NULL, '$2y$10$BbccK69TJEwVv0CwqYmnTe73ZKwcAtnIgveDYXqPtov4eKCW/NbTa', NULL, '2021-06-15 00:40:07', '2021-06-15 00:40:07', 'https://tvax2.sinaimg.cn/crop.0.0.639.639.180/74889c63ly8fdpsdi7jfaj20hs0hrq3j.jpg?KID=imgbed,tva&Expires=1623699607&ssig=hZ0BuKjyH8', '1955109987', 0, NULL, NULL, 0);
INSERT INTO `users` VALUES (23, '用户6134663166', '', NULL, '$2y$10$v4R7D7UW8kUGUFqATOLl3.zcyps5puY7g8i/HLExdk/jtVv34VT1e', NULL, '2021-06-15 18:29:33', '2021-06-15 18:29:33', 'https://tvax1.sinaimg.cn/default/images/default_avatar_male_180.gif?KID=imgbed,tva&Expires=1623763773&ssig=OQNmLhNUTF', '6134663166', 0, NULL, NULL, 0);
INSERT INTO `users` VALUES (24, '嘘里嘘气的', '', NULL, '$2y$10$2MQm9IiUY2zMofTt3B.H4OY3nBnpikmdThytS9gZ2kl9gjz0o4Zn.', NULL, '2021-06-16 00:54:32', '2021-06-16 00:54:32', 'https://tvax4.sinaimg.cn/crop.0.0.690.690.180/0071NFG6ly8fmmdnzfqtzj30j60j6wkz.jpg?KID=imgbed,tva&Expires=1623786871&ssig=7%2BFnImcbhd', '6439544446', 0, NULL, NULL, 0);
INSERT INTO `users` VALUES (25, 'ZL_at_DGUT', '', NULL, '$2y$10$DtuXe8RTM4W/J9OnliNXeu8lOtC5Xc64BmC/hX9w79TjtqLwdej7y', NULL, '2021-06-16 10:46:39', '2021-06-16 10:46:39', 'https://tva1.sinaimg.cn/crop.0.0.180.180.180/6c02df7djw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1623822398&ssig=ucvfP4ZNir', '1812127613', 0, NULL, NULL, 0);
INSERT INTO `users` VALUES (26, 'LBJ-敏', '', NULL, '$2y$10$YXaheHjgAhfk3wna94emmOFbD5SUPyJlNQ1cV3MtOLKrzEbWvtTWu', NULL, '2021-06-16 14:07:23', '2021-06-16 14:07:23', 'https://tva1.sinaimg.cn/crop.0.0.180.180.180/be8f8d5bjw1e8qgp5bmzyj2050050aa8.jpg?KID=imgbed,tva&Expires=1623834442&ssig=KpxfJVxd8p', '3197078875', 0, NULL, NULL, 0);
INSERT INTO `users` VALUES (27, '恐惧丶恶魔', '', NULL, '$2y$10$EDKVNqr0.R91.IZudtce.O1ISBj2BSam.HVsubWcNXiMVK/Ly.k.W', NULL, '2021-06-16 20:41:25', '2021-06-16 20:41:25', 'https://tva1.sinaimg.cn/crop.1.0.198.198.180/006bDctWjw1evdlkvho0rj305k05kgln.jpg?KID=imgbed,tva&Expires=1623858085&ssig=fy0YElw978', '5668679464', 0, NULL, NULL, 0);
INSERT INTO `users` VALUES (28, 'myronChina', '', NULL, '$2a$10$Ee.HQ9xk3vK8ProJD4YEzu..4BV3btJSARyfdVF9892vNO4FIr1t6', NULL, '2021-06-28 10:04:03', NULL, 'https://tvax2.sinaimg.cn/crop.0.0.1080.1080.180/c1364b26ly8fy8zoijjy0j20u00u040c.jpg?KID=imgbed,tva&Expires=1624856643&ssig=YVN7NYj6PQ', '3241560870', 0, NULL, 1, 0);
INSERT INTO `users` VALUES (29, '用户6200924559', '', NULL, '$2a$10$JbrbIfWE7u9BCVjWeVXz1es7Au6Ljlc/3qNARedRs9X2ysT9oR69O', NULL, '2021-06-29 20:26:36', NULL, 'https://tvax1.sinaimg.cn/default/images/default_avatar_male_180.gif?KID=imgbed,tva&Expires=1624980396&ssig=0Vi0pF4H4D', '6200924559', 0, NULL, 1, 0);
INSERT INTO `users` VALUES (30, 'admin', '', NULL, '$2a$10$ntZsWa6w97AdpGwlwJK.z.rA/b9inQfkOCwnQBFCJN2kt2DE7HRD6', NULL, '2021-07-01 14:54:46', NULL, 'https://tvax2.sinaimg.cn/crop.0.0.1002.1002.180/006pP2Laly8gqcj17wce9j30ru0ru0up.jpg?KID=imgbed,tva&Expires=1624513271&ssig=lmrNTiDUPy', '', 0, NULL, 1, 0);
INSERT INTO `users` VALUES (31, 'hhhhh', NULL, NULL, '$2a$10$ntZsWa6w97AdpGwlwJK.z.rA/b9inQfkOCwnQBFCJN2kt2DE7HRD6', NULL, NULL, NULL, 'https://tvax1.sinaimg.cn/default/images/default_avatar_male_180.gif?KID=imgbed,tva&Expires=1624980396&ssig=0Vi0pF4H4D', NULL, 0, NULL, NULL, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
