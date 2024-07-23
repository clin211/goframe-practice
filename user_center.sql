/*
 Navicat Premium Data Transfer

 Source Server         : sh
 Source Server Type    : MySQL
 Source Server Version : 80200
 Source Host           : 127.0.0.1:3306
 Source Schema         : user_center

 Target Server Type    : MySQL
 Target Server Version : 80200
 File Encoding         : 65001

 Date: 23/07/2024 11:21:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned zerofill NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
