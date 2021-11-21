/*
 Navicat Premium Data Transfer

 Source Server         : 阿里云
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : 139.224.11.85:3306
 Source Schema         : nuxt-blog

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 23/01/2021 14:33:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article_tags
-- ----------------------------
DROP TABLE IF EXISTS `article_tags`;
CREATE TABLE `article_tags`  (
  `article_id` int(11) NOT NULL,
  `tag_id` int(11) NOT NULL,
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0)
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章标签表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`  (
  `article_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `article_title` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标题',
  `article_author` varchar(25) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'bigfool' COMMENT '作者',
  `article_read` int(11) NOT NULL DEFAULT 0 COMMENT '阅读量',
  `article_status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态 0-草稿 1-发布',
  `article_desc` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '描述',
  `article_content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`article_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 63 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags`  (
  `tag_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标签名称',
  `tag_icon` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标签图标',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`tag_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '标签表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for vistors
-- ----------------------------
DROP TABLE IF EXISTS `vistors`;
CREATE TABLE `vistors`  (
  `vistor_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `vistor_ip` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'IP地址',
  `user_agent` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `vistor_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '路径',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '创建时间',
  PRIMARY KEY (`vistor_id`) USING BTREE,
  INDEX `idx_vistor_create_time`(`create_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8670 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '访客表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
