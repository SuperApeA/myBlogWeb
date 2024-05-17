/*
 Navicat Premium Data Transfer

 Source Server Type    : MySQL
 Source Server Version : 80036
 Source Schema         : myblog

 Target Server Type    : MySQL
 Target Server Version : 80036
 File Encoding         : 65001

 Date: 17/05/2024 18:21:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_category
-- ----------------------------
DROP TABLE IF EXISTS `blog_category`;
CREATE TABLE `blog_category`  (
  `cid` int NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`cid`) USING BTREE,
  UNIQUE INDEX `idx_cid`(`cid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_category
-- ----------------------------
INSERT INTO `blog_category` VALUES (1, 'Go', '2022-02-14 23:55:28', '2022-02-14 23:55:28');
INSERT INTO `blog_category` VALUES (2, 'Java', '2022-02-17 17:16:07', '2022-02-17 17:16:07');

-- ----------------------------
-- Table structure for blog_post
-- ----------------------------
DROP TABLE IF EXISTS `blog_post`;
CREATE TABLE `blog_post`  (
  `pid` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `markdown` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `category_id` int NOT NULL,
  `user_id` bigint NOT NULL,
  `view_count` int NOT NULL DEFAULT 0,
  `type` int NOT NULL DEFAULT 0,
  `slug` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_at` datetime NOT NULL,
  `update_at` datetime NOT NULL,
  PRIMARY KEY (`pid`) USING BTREE,
  UNIQUE INDEX `idx_pid`(`pid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_post
-- ----------------------------
INSERT INTO `blog_post` VALUES (1, 'test', '<p>123<br>asdfasdf</p>\n', '123\nasdfasdf\n\n', 1, 1, 9, 0, '', '2024-05-16 16:22:57', '2024-05-17 18:01:01');
INSERT INTO `blog_post` VALUES (2, 'test', '<p><img src=\"/home/workspace/myBlog/myBlogWeb/viewsrc/markdown/image/ironman.png\" alt=\"\">测试测试</p>\n', '![](/home/workspace/myBlog/myBlogWeb/viewsrc/markdown/image/ironman.png)测试测试', 1, 1, 16, 0, '', '2024-04-20 02:16:04', '2024-04-29 22:31:07');
INSERT INTO `blog_post` VALUES (3, 'test1', '<p>123<br>asdfasdf</p>\n', '123\nasdfasdf\n\n', 2, 1, 3, 0, '', '2024-05-17 16:17:19', '2024-05-17 18:04:59');
INSERT INTO `blog_post` VALUES (4, 'test2', '<p>231231231</p>\n', '231231231\n\n', 2, 1, 0, 0, '', '2024-05-17 16:23:25', '2024-05-17 16:27:40');
INSERT INTO `blog_post` VALUES (5, 'test2', '<p>231231231</p>\n', '231231231\n\n', 1, 1, 0, 0, '', '2024-05-17 16:34:12', '2024-05-17 16:34:12');
INSERT INTO `blog_post` VALUES (6, 'test2asfasdf', '<p>231231231</p>\n', '231231231\n\n', 1, 1, 4, 2, '', '2024-05-17 16:34:47', '2024-05-17 18:02:28');
INSERT INTO `blog_post` VALUES (7, 'test2asfasdf', '<p>231231231asf<br>asfasf<br>asfasdfasdfasd<br>asfasf<br>asfasdf</p>\n', '231231231asf\nasfasf\nasfasdfasdfasd\nasfasf\nasfasdf\n\n', 1, 1, 1, 0, '', '2024-05-17 17:08:59', '2024-05-17 17:09:18');
INSERT INTO `blog_post` VALUES (8, 'test2asfasdf', '<p>231231231asf<br>asfasf<br>asfasdfasdfasd<br>asfasf<br>asfasdf<br>asfasdf</p>\n', '231231231asf\nasfasf\nasfasdfasdfasd\nasfasf\nasfasdf\nasfasdf\n\n', 1, 1, 1, 0, '', '2024-05-17 17:09:26', '2024-05-17 17:09:41');
INSERT INTO `blog_post` VALUES (9, 'test2asfasdf', '<p>231231231asf<br>asfasf<br>asfasdfasdfasd<br>asfasf<br>asfasdf<br>asfasdf</p>\n', '231231231asf\nasfasf\nasfasdfasdfasd\nasfasf\nasfasdf\nasfasdf\n\n', 1, 1, 2, 0, '', '2024-05-17 17:10:09', '2024-05-17 17:10:25');

-- ----------------------------
-- Table structure for blog_user
-- ----------------------------
DROP TABLE IF EXISTS `blog_user`;
CREATE TABLE `blog_user`  (
  `uid` bigint NOT NULL,
  `user_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `passwd` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `avatar` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_at` datetime NOT NULL,
  `update_at` datetime NOT NULL,
  PRIMARY KEY (`uid`) USING BTREE,
  UNIQUE INDEX `idx_user_name`(`user_name` ASC) USING BTREE,
  UNIQUE INDEX `idx_uid`(`uid` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_user
-- ----------------------------
INSERT INTO `blog_user` VALUES (1, 'admin', '6d902979ff0894173e31bd879d58648d', '/resource/images/avatar.jpeg', '2024-04-17 23:29:33', '2024-04-17 23:29:38');

SET FOREIGN_KEY_CHECKS = 1;
