/*
 Navicat Premium Data Transfer

 Source Server         : root
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : finders_test

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 14/08/2020 22:19:45
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activities
-- ----------------------------
DROP TABLE IF EXISTS `activities`;
CREATE TABLE `activities` (
                              `activity_id` varchar(50) NOT NULL AUTO_INCREMENT COMMENT '帖子ID',
                              `activity_status` int DEFAULT NULL COMMENT '帖子类型',
                              `activity_info` text COMMENT '帖子内容',
                              `collect_num` int NOT NULL COMMENT '收藏次数',
                              `comment_num` int NOT NULL COMMENT '评论次数',
                              `read_num` int NOT NULL COMMENT '阅读次数',
                              `media_ids` varchar(5000) DEFAULT NULL COMMENT '帖子媒体ID',
                              `user_id` varchar(50) DEFAULT NULL COMMENT '发表用户ID',
                              `activity_title` varchar(50) DEFAULT NULL COMMENT '帖子标题',
                              `community_id` int NOT NULL COMMENT '所属社区ID',
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                              PRIMARY KEY (`activity_id`),
                              KEY `fk_activity_user` (`user_id`),
                              KEY `fk_activity_community` (`community_id`),
                              CONSTRAINT `fk_activity_community` FOREIGN KEY (`community_id`) REFERENCES `communities` (`community_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                              CONSTRAINT `fk_activity_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for activity_likes
-- ----------------------------
DROP TABLE IF EXISTS like_maps;
CREATE TABLE `activity_likes` (
                                  `activity_id` varchar(50) NOT NULL COMMENT '帖子ID',
                                  `user_id` varchar(50) NOT NULL COMMENT '用户ID',
                                  `id` int NOT NULL AUTO_INCREMENT,
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '点赞时间',
                                  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                                  PRIMARY KEY (`id`),
                                  KEY `fk_likes_activity` (`activity_id`),
                                  KEY `fk_likes_user` (`user_id`),
                                  CONSTRAINT `fk_likes_activity` FOREIGN KEY (`activity_id`) REFERENCES `activities` (`activity_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                                  CONSTRAINT `fk_likes_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for admins
-- ----------------------------
DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins` (
                          `admin_id` varchar(50) NOT NULL,
                          `admin_name` varchar(50) DEFAULT NULL,
                          `admin_password` varchar(30) DEFAULT NULL,
                          `admin_phone` varchar(30) DEFAULT NULL,
                          `permission` int DEFAULT NULL,
                          `created_at` datetime DEFAULT NULL,
                          `updated_at` datetime DEFAULT NULL,
                          `deleted_at` datetime DEFAULT NULL,
                          PRIMARY KEY (`admin_id`),
                          UNIQUE KEY `unique_admin_phone` (`admin_phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for collections
-- ----------------------------
DROP TABLE IF EXISTS `collections`;
CREATE TABLE `collections` (
                               `collection_id` int NOT NULL AUTO_INCREMENT COMMENT '收藏ID',
                               `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '收藏时间',
                               `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                               `collection_status` varchar(10) NOT NULL COMMENT '收藏状态',
                               `collection_type` int NOT NULL COMMENT '收藏类型',
                               `user_id` varchar(50) NOT NULL,
                               `link` varchar(100) NOT NULL COMMENT '收藏链接',
                               PRIMARY KEY (`collection_id`),
                               KEY `fk_collection_user` (`user_id`),
                               CONSTRAINT `fk_collection_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
                            `comment_id` int NOT NULL AUTO_INCREMENT COMMENT '评论ID',
                            `item_id` varchar(50) NOT NULL COMMENT '评论的帖子ID',
                            `item_type` int NOT NULL,
                            `content` text COMMENT '评论内容',
                            `from_uid` varchar(50) NOT NULL COMMENT '评论用户ID',
                            `to_uid` varchar(50) DEFAULT NULL,
                            `status` int DEFAULT NULL COMMENT '评论状态',
                            `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                            PRIMARY KEY (`comment_id`),
                            KEY `fk_comment_user` (`from_uid`),
                            KEY `fk_comment_activity` (`item_id`),
                            CONSTRAINT `fk_comment_user` FOREIGN KEY (`from_uid`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for communities
-- ----------------------------
DROP TABLE IF EXISTS `communities`;
CREATE TABLE `communities` (
                               `community_id` int NOT NULL AUTO_INCREMENT COMMENT '社区ID',
                               `community_creator` varchar(50) NOT NULL COMMENT '社区创建者（圈主）',
                               `community_name` varchar(100) NOT NULL COMMENT '社区名称',
                               `community_description` text COMMENT '社区简介',
                               `community_status` int NOT NULL COMMENT '社区状态',
                               `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                               `background` varchar(200) DEFAULT NULL COMMENT '背景',
                               `community_avatar` varchar(200) DEFAULT NULL COMMENT '圈子头像',
                               PRIMARY KEY (`community_id`),
                               KEY `fk_com_user` (`community_creator`),
                               CONSTRAINT `fk_com_user` FOREIGN KEY (`community_creator`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for community_managers
-- ----------------------------
DROP TABLE IF EXISTS `community_managers`;
CREATE TABLE `community_managers` (
                                      `community_id` int NOT NULL COMMENT '社区ID',
                                      `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
                                      `manager_id` varchar(50) DEFAULT NULL COMMENT '管理员ID',
                                      `permission` int DEFAULT NULL COMMENT '圈子管理员权限',
                                      `status` int NOT NULL COMMENT '管理员状态',
                                      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                      `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                      `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                                      PRIMARY KEY (`id`),
                                      KEY `fk_manager_community` (`community_id`),
                                      KEY `fk_manager_user` (`manager_id`),
                                      CONSTRAINT `fk_manager_community` FOREIGN KEY (`community_id`) REFERENCES `communities` (`community_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                                      CONSTRAINT `fk_manager_user` FOREIGN KEY (`manager_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for community_users
-- ----------------------------
DROP TABLE IF EXISTS `community_users`;
CREATE TABLE `community_users` (
                                   `id` int NOT NULL AUTO_INCREMENT,
                                   `user_id` varchar(50) DEFAULT NULL COMMENT '社区用户ID',
                                   `community_id` int DEFAULT NULL COMMENT '社区ID',
                                   `status` int DEFAULT NULL COMMENT '用户当前在圈子的状态',
                                   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                   `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                                   PRIMARY KEY (`id`),
                                   KEY `fk_community_user_user` (`user_id`),
                                   KEY `fk_community_user_community` (`community_id`),
                                   CONSTRAINT `fk_community_user_community` FOREIGN KEY (`community_id`) REFERENCES `communities` (`community_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                                   CONSTRAINT `fk_community_user_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for logins
-- ----------------------------
DROP TABLE IF EXISTS `logins`;
CREATE TABLE `logins` (
                          `login_id` int NOT NULL AUTO_INCREMENT COMMENT '登陆ID',
                          `user_id` varchar(50) NOT NULL COMMENT '用户ID',
                          `access_token` varchar(250) NOT NULL COMMENT 'api token',
                          `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '登陆时间',
                          `expired_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '过期时间',
                          PRIMARY KEY (`login_id`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for medias
-- ----------------------------
DROP TABLE IF EXISTS `medias`;
CREATE TABLE `medias` (
                          `media_id` varchar(50) NOT NULL,
                          `media_url` varchar(200) DEFAULT NULL,
                          `media_type` int DEFAULT NULL,
                          `user_id` varchar(50) DEFAULT NULL,
                                   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                   `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                          PRIMARY KEY (`media_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- ----------------------------
-- Table structure for relations
-- ----------------------------
DROP TABLE IF EXISTS `relations`;
CREATE TABLE `relations` (
                             `relation_id` int NOT NULL AUTO_INCREMENT COMMENT '用户关系ID',
                             `relation_type` int NOT NULL COMMENT '关系类型',
                             `relation_group` varchar(20) NOT NULL COMMENT '关系组名',
                             `from_uid` varchar(50) NOT NULL COMMENT '用户ID',
                             `to_uid` varchar(50) NOT NULL COMMENT '被关注用户ID',
                             `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '关系建立时间',
                             `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '关系更新时间',
                             `deleted_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
                             PRIMARY KEY (`relation_id`),
                             UNIQUE KEY `unique_user_user` (`from_uid`,`to_uid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for tag_map
-- ----------------------------
DROP TABLE IF EXISTS `tag_map`;
CREATE TABLE `tag_map` (
                           `item_id` varchar(100) NOT NULL,
                           `item_type` int DEFAULT NULL,
                           `tag_id` int NOT NULL,
                           `created_at` datetime DEFAULT NULL,
                           `updated_at` datetime DEFAULT NULL,
                           `deleted_at` datetime DEFAULT NULL,
                           PRIMARY KEY (`item_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
                        `tag_id` int NOT NULL AUTO_INCREMENT COMMENT '标签的id',
                        `tag_name` varchar(50) DEFAULT NULL,
                        `tag_type` int DEFAULT NULL,
                        `created_at` datetime DEFAULT NULL,
                        `updated_at` datetime DEFAULT NULL,
                        `deleted_at` datetime DEFAULT NULL,
                        PRIMARY KEY (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user_infos
-- ----------------------------
DROP TABLE IF EXISTS `user_infos`;
CREATE TABLE `user_infos` (
                              `user_id` varchar(50) NOT NULL,
                              `truename` varchar(40) DEFAULT NULL,
                              `address` varchar(200) DEFAULT NULL,
                              `sex` varchar(4) DEFAULT NULL,
                              `sexual` varchar(8) DEFAULT NULL,
                              `feeling` varchar(20) DEFAULT NULL,
                              `birthday` varchar(20) DEFAULT NULL,
                              `introduction` varchar(400) DEFAULT NULL,
                              `signature` varchar(400) DEFAULT NULL,
                              `blood_type` varchar(8) DEFAULT NULL,
                              `eamil` varchar(60) DEFAULT NULL,
                              `qq` varchar(30) DEFAULT NULL,
                              `wechat` varchar(30) DEFAULT NULL,
                              `profession` varchar(60) DEFAULT NULL,
                              `school` varchar(30) DEFAULT NULL,
                              `constellation` varchar(40) DEFAULT NULL,
                              `created_at` datetime DEFAULT NULL,
                              `updated_at` datetime DEFAULT NULL,
                              `credit` int DEFAULT NULL,
                              `user_tag` text,
                              `deleted_at` datetime DEFAULT NULL,
                              `age` int DEFAULT NULL,
                              PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `user_id` varchar(50) NOT NULL,
                         `phone` varchar(30) DEFAULT NULL,
                         `password` varchar(100) DEFAULT NULL,
                         `nickname` varchar(30) DEFAULT NULL,
                         `created_at` datetime DEFAULT NULL,
                         `status` int DEFAULT NULL,
                         `deleted_at` datetime DEFAULT NULL,
                         `avatar` varchar(100) DEFAULT NULL,
                         `username` varchar(50) DEFAULT NULL,
                         `updated_at` datetime DEFAULT NULL,
                         PRIMARY KEY (`user_id`),
                         UNIQUE KEY `unique_phone` (`phone`),
                         UNIQUE KEY `unique_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for activities
-- ----------------------------
DROP TABLE IF EXISTS `moments`;
CREATE TABLE `moments` (
                              `moment_id` varchar(50) NOT NULL  COMMENT '动态ID',
                              `moment_status` int DEFAULT NULL COMMENT '动态类型',
                              `moment_info` text COMMENT '动态内容',
                              `read_num` int NOT NULL COMMENT '阅读次数',
                              `media_ids` varchar(5000) DEFAULT NULL COMMENT '动态的媒体IDs',
                              `user_id` varchar(50) DEFAULT NULL COMMENT '发表用户ID',
                              `location` varchar(500) DEFAULT NULL COMMENT '发表位置',
                              `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
                              PRIMARY KEY (`moment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for message_box
-- ----------------------------
DROP TABLE IF EXISTS `question_box`;
CREATE TABLE `question_box` (
                             `question_box_id` int NOT NULL AUTO_INCREMENT COMMENT '提问箱id',
                             `user_id` varchar(50) NOT NULL COMMENT '发出问题的用户id',
                             `question_box_status` int NOT NULL COMMENT '提问箱状态，封禁或正常',
                             `question_box_info` text NOT NULL COMMENT '提问箱内容',
                             `use_num` int DEFAULT  0 COMMENT '使用次数',
                             `reply_num` int DEFAULT 0 COMMENT '回复次数',
                             `like_num` int DEFAULT 0 COMMENT '喜欢的次数',
                             `tag_names` varchar(5000) NOT NULL COMMENT '关联内容的标签，使用;分割',
                             `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '关系建立时间',
                             `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '关系更新时间',
                             `deleted_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
                             PRIMARY KEY (`question_box_id`)

) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;



