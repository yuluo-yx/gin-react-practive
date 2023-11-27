/*
 Navicat Premium Data Transfer

 Source Server         : local-mysql
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : dimeng

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 26/11/2023 21:20:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `create_time` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_message_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (4, '2023-11-07 08:26:02.151', '2023-11-07 08:26:02.151', '2023-11-07 08:29:47.744', '这是一个留言信息！-牧生', 2, NULL);
INSERT INTO `message` VALUES (5, '2023-11-07 08:56:43.791', '2023-11-07 08:56:43.791', NULL, '这是一个留言信息！-5', 4, NULL);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password_digest` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_name`(`user_name`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (2, '2023-11-25 22:47:50.000', '2023-11-25 22:47:53.000', NULL, '牧生', '$2a$12$rUw8P6UiJLY5YaYErVIY6eMDgyl2K917gNwrSOSS4BNeifSXvA4U2');
INSERT INTO `user` VALUES (3, '2023-11-07 07:11:14.111', '2023-11-07 07:11:14.111', NULL, 'yuluo', '$2a$12$8jT.s3wofig51cV72eSh0.SqyBINpGuMchQ8eTBfd4O/dhSC8GKqO');
INSERT INTO `user` VALUES (4, '2023-11-07 07:12:08.310', '2023-11-07 07:12:08.310', NULL, 'test', '$2a$12$rUw8P6UiJLY5YaYErVIY6eMDgyl2K917gNwrSOSS4BNeifSXvA4U2');
INSERT INTO `user` VALUES (6, '2023-11-26 01:32:10.481', '2023-11-26 01:32:10.481', NULL, 'test-1', '$2a$12$Khpp/rkXnmQ4AtCeQ/VPO.SQM6Uzn5qHqZYZ.tBdtg4tY2tRf7YAe');
INSERT INTO `user` VALUES (8, '2023-11-26 01:36:16.315', '2023-11-26 01:36:16.315', NULL, 'yuluo-1', '$2a$12$mhUPjtxbbMh7T8Rn1d3oVOqsUvpp3fEddYbJKhdP.eqQVintkfJte');

SET FOREIGN_KEY_CHECKS = 1;
