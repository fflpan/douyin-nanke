/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : golearn

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 07/06/2022 01:16:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `cid` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL DEFAULT '0',
  `video_id` bigint NOT NULL DEFAULT '0',
  `content` varchar(10) NOT NULL DEFAULT '',
  `create_date` datetime NOT NULL,
  PRIMARY KEY (`cid`),
  KEY `videoIndex` (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of comment
-- ----------------------------
BEGIN;
INSERT INTO `comment` VALUES (1, 1, 5, 'hello', '2022-05-24 17:18:07');
INSERT INTO `comment` VALUES (2, 1, 5, 'hello', '2022-05-24 17:18:27');
INSERT INTO `comment` VALUES (3, 1, 5, 'hello', '2022-05-24 17:37:13');
INSERT INTO `comment` VALUES (4, 1, 1, 'text', '2022-05-24 17:39:19');
INSERT INTO `comment` VALUES (5, 1, 1, 'text', '2022-05-24 17:40:01');
INSERT INTO `comment` VALUES (6, 1, 5, 'nih', '2022-05-24 19:14:55');
INSERT INTO `comment` VALUES (7, 1, 6, 'pretty', '2022-05-30 11:21:11');
COMMIT;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
  `favorite_id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL DEFAULT '0',
  `video_id` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`favorite_id`),
  KEY `index1` (`user_id`,`video_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of favorite
-- ----------------------------
BEGIN;
INSERT INTO `favorite` VALUES (20, 1, 5);
INSERT INTO `favorite` VALUES (21, 1, 6);
COMMIT;

-- ----------------------------
-- Table structure for follow
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow` (
  `follow_id` bigint NOT NULL AUTO_INCREMENT,
  `follower_id` bigint NOT NULL DEFAULT '0',
  `be_follow_id` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`follow_id`),
  KEY `followIndex` (`follower_id`),
  KEY `beFollowIndex` (`be_follow_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of follow
-- ----------------------------
BEGIN;
INSERT INTO `follow` VALUES (9, 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `user_id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `follow_count` bigint NOT NULL DEFAULT '1',
  `follower_count` bigint NOT NULL DEFAULT '1',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'scj', 'fflpan', 1, 1);
COMMIT;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `video_id` bigint NOT NULL AUTO_INCREMENT,
  `author_id` bigint NOT NULL DEFAULT '0',
  `play_url` varchar(255) NOT NULL DEFAULT '',
  `cover_url` varchar(255) NOT NULL DEFAULT '',
  `favorite_count` bigint NOT NULL DEFAULT '0',
  `comment_count` bigint NOT NULL DEFAULT '0',
  `pub_time` datetime NOT NULL,
  `title` varchar(300) NOT NULL DEFAULT '',
  PRIMARY KEY (`video_id`),
  KEY `authorIndex` (`author_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of video
-- ----------------------------
BEGIN;
INSERT INTO `video` VALUES (2, 1, 'http://10.180.139.161:8080/video/scj5577006791947779410av1.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 0, 0, '2022-05-13 09:52:54', '士兵突击');
INSERT INTO `video` VALUES (5, 1, 'http://10.180.139.161:8080/video/8674665223082153551av2.mp4', 'https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg', 1, 4, '2022-05-13 11:48:00', '西藏自驾');
INSERT INTO `video` VALUES (6, 1, 'http://10.180.139.161:8080/video/1653876672687033000pretty.mp4', 'http://10.180.139.161:8080/img/1653876672687033000.jpg', 1, 1, '2022-05-30 10:11:13', 'HANGUO');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
