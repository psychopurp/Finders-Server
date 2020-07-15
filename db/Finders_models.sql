CREATE TABLE `users` (
`user_id` varchar(30) NOT NULL COMMENT '用户ID',
`phone` varchar(30) NOT NULL COMMENT '手机号',
`password` varchar(30) NOT NULL COMMENT '密码',
`nickname` varchar(30) NOT NULL COMMENT '昵称',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '注册时间',
`status` int NOT NULL COMMENT '用户状态',
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
`updated_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`avatar` varchar(100) NOT NULL COMMENT '用户头像',
PRIMARY KEY (`user_id`) ,
UNIQUE INDEX `unique_phone` (`phone` ASC)
);
CREATE TABLE `user_infos` (
`user_id` varchar(30) NOT NULL COMMENT '用户ID',
`truename` varchar(40) NULL COMMENT '真实姓名',
`address` varchar(200) NULL COMMENT '所在地',
`sex` varchar(4) NULL COMMENT '性别',
`sexual` varchar(8) NULL COMMENT '性取向',
`feeling` varchar(20) NULL COMMENT '感情状况',
`birthday` varchar(20) NULL COMMENT '生日',
`introduction` varchar(400) NULL COMMENT '简介',
`blood_type` varchar(8) NULL COMMENT '血型',
`eamil` varchar(60) NULL COMMENT '邮箱',
`qq` varchar(30) NULL COMMENT 'QQ',
`wechat` varchar(30) NULL COMMENT '微信',
`profession` varchar(60) NULL COMMENT '职业信息',
`school` varchar(30) NULL COMMENT '学校',
`constellation` varchar(40) NULL COMMENT '星座',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`credit` int NOT NULL COMMENT '用户信誉积分',
`user_tag` text CHARACTER SET 用户标签 NULL,
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
`age` int NULL COMMENT '年龄',
PRIMARY KEY (`user_id`) 
);
CREATE TABLE `admins` (
`admin_id` varchar(30) NOT NULL COMMENT '管理员ID',
`admin_name` varchar(30) NOT NULL COMMENT '管理员名称',
`admin_password` varchar(30) NOT NULL COMMENT '管理员密码',
`permission` int NOT NULL COMMENT '管理员权限',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '注册时间',
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '注销管理员的时间',
PRIMARY KEY (`admin_id`) 
);
CREATE TABLE `relations` (
`relation_id` int NOT NULL AUTO_INCREMENT COMMENT '用户关系ID',
`relation_type` int NOT NULL COMMENT '关系类型',
`relation_group` varchar(20) NOT NULL COMMENT '关系组名',
`from_uid` varchar(50) NOT NULL COMMENT '用户ID',
`to_uid` varchar(50) NOT NULL COMMENT '被关注用户ID',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '关系建立时间',
`updated_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '关系更新时间',
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
PRIMARY KEY (`relation_id`) ,
UNIQUE INDEX `unique_user_user` (`from_uid` ASC, `to_uid` ASC) USING BTREE
);
CREATE TABLE `messages` (
`message_id` int NOT NULL COMMENT '消息ID',
`message_info` text NULL COMMENT '消息内容',
`message_status` int NOT NULL COMMENT '消息状态状态',
`send_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '发送时间',
`from_uid` varchar(30) NOT NULL COMMENT '发送用户ID',
`to_uid` varchar(30) NOT NULL COMMENT '接受用户ID',
`message_type` int NULL COMMENT '消息类型',
`link` varchar(100) NULL COMMENT '消息链接',
PRIMARY KEY (`message_id`) 
);
CREATE TABLE `logins` (
`login_id` int  NOT NULL COMMENT '登陆ID',
`user_id` varchar(50) NOT NULL COMMENT '用户ID',
`access_token` varchar(250) NOT NULL COMMENT 'api token',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '登陆时间',
`expired_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '过期时间',
PRIMARY KEY (`login_id`) 
);
CREATE TABLE `collections` (
`collection_id` int NOT NULL COMMENT '收藏ID',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '收藏时间',
`collection_status` varchar(10) NOT NULL COMMENT '收藏状态',
`collection_type` int NOT NULL COMMENT '收藏类型',
`user_id` varchar(30) NOT NULL,
`link` varchar(100) NOT NULL COMMENT '收藏链接',
PRIMARY KEY (`collection_id`) 
);
CREATE TABLE `communities` (
`community_id` int NOT NULL AUTO_INCREMEN COMMENT '社区ID',
`community_creator` varchar(30) NOT NULL COMMENT '社区创建者（圈主）',
`community_name` varchar(100) NOT NULL COMMENT '社区名称',
`community_description` text COMMENT '社区简介' NULL,
`community_status` int NOT NULL COMMENT '社区状态',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
`background` varchar(200) NULL COMMENT '背景',
PRIMARY KEY (`community_id`)
);
CREATE TABLE `pictures` (
`picture_id` varchar(30) NOT NULL COMMENT '图片ID',
`picture_url` varchar(200) NULL COMMENT '图片地址',
`picture_type` int NULL COMMENT '图片类型',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`user_id` varchar(30) NOT NULL COMMENT '用户ID',
PRIMARY KEY (`picture_id`)
);
CREATE TABLE `tags` (
`tag_id` int NOT NULL COMMENT '标签ID',
`tag_name` varchar(50) NOT NULL COMMENT '标签名',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`tag_type` int  NULL COMMENT '标签类型',
PRIMARY KEY (`tag_id`)
);
CREATE TABLE `activities` (
`activity_id` varchar(50) NOT NULL COMMENT '帖子ID',
`activity_status` int NULL COMMENT '帖子类型',
`activity_info` text COMMENT '帖子内容' NULL,
`collect_num` int NOT NULL COMMENT '收藏次数',
`comment_num` int NOT NULL COMMENT '评论次数',
`read_num` int NOT NULL COMMENT '阅读次数',
`media_id` varchar(50) NULL COMMENT '帖子媒体ID',
`media_type` int NULL COMMENT '帖子媒体类型',
`user_id` varchar(50) NULL COMMENT '发表用户ID',
`community_id` int NOT NULL COMMENT '所属社区ID',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
PRIMARY KEY (`activity_id`) 
);
CREATE TABLE `community_managers` (
`community_id` int NOT NULL COMMENT '社区ID',
`id` int AUTO_INCREMENT NOT NULL COMMENT 'ID',
`manager_id` varchar(30) NULL COMMENT '管理员ID',
`permission` int NULL COMMENT '圈子管理员权限',
`status` int NOT NULL COMMENT '管理员状态',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
PRIMARY KEY (`id`) 
);
CREATE TABLE `community_users` (
`id` int AUTO_INCREMENT NOT NULL,
`user_id` varchar(50) NULL COMMENT '社区用户ID',
`community_id` int NULL COMMENT '社区ID',
`status` int NULL COMMENT '用户当前在圈子的状态',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
PRIMARY KEY (`id`) 
);
CREATE TABLE `activity_likes` (
`activity_id` varchar(50) NOT NULL COMMENT '帖子ID',
`user_id` varchar(50) NOT NULL COMMENT '用户ID',
`id` int AUTO_INCREMENT NOT NULL,
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '点赞时间',
`updated_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
PRIMARY KEY (`id`) 
);
CREATE TABLE `comments` (
`comment_id` int NOT NULL COMMENT '评论ID',
`activity_id` varchar(30) NULL COMMENT '评论的帖子ID',
`activity_type` varchar(100) NULL COMMENT '帖子类型',
`content` text CHARACTER SET 评论内容 NULL,
`from_uid` varchar(30) NULL COMMENT '评论用户ID',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`status` int NULL COMMENT '评论状态',
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
PRIMARY KEY (`comment_id`) 
);
CREATE TABLE `replies` (
`reply_id` int NOT NULL COMMENT '回复ID',
`comment_id` int NULL COMMENT '评论ID',
`reply_type` int NULL COMMENT '回复类型',
`to_reply_id` int NULL COMMENT '回复目标ID',
`content` text CHARACTER SET 回复内容 NULL,
`from_uid` varchar(30) NULL COMMENT '回复用户ID',
`to_uid` varchar(30) NULL COMMENT '目标用户ID',
`created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '删除时间',
`status` int NULL COMMENT '回复状态',
PRIMARY KEY (`reply_id`) 
);
CREATE TABLE `at_users` (
`atusers_id` int NOT NULL COMMENT '艾特ID',
`acitvity_id` varchar(30) NULL COMMENT '帖子ID',
`from_uid` varchar(30) NULL COMMENT '用户ID',
`to_uid` varchar(30) NULL COMMENT '被艾特用户ID',
`created_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '艾特时间',
PRIMARY KEY (`atusers_id`) 
);

ALTER TABLE `user_infos` ADD CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `relations` ADD CONSTRAINT `fk_user_id1` FOREIGN KEY (`from_uid`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `relations` ADD CONSTRAINT `fk_user_id2` FOREIGN KEY (`to_uid`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `messages` ADD CONSTRAINT `fk_letter_user_id1` FOREIGN KEY (`from_uid`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `messages` ADD CONSTRAINT `fk_letter_user_id2` FOREIGN KEY (`to_uid`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `logins` ADD CONSTRAINT `fk_login_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `collections` ADD CONSTRAINT `fk_collection_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `pictures` ADD CONSTRAINT `fk_picture_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE `communities` ADD CONSTRAINT `fk_com_user` FOREIGN KEY (`community_creator`) REFERENCES `users` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE `activities` ADD CONSTRAINT `fk_activity_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `activities` ADD CONSTRAINT `fk_activity_picture` FOREIGN KEY (`picture_id`) REFERENCES `pictures` (`picture_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE `activities` ADD CONSTRAINT `fk_activity_community` FOREIGN KEY (`community_id`) REFERENCES `communities` (`community_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `community_managers` ADD CONSTRAINT `fk_manager_community` FOREIGN KEY (`community_id`) REFERENCES `communities` (`community_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `community_managers` ADD CONSTRAINT `fk_manager_user` FOREIGN KEY (`manager_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `community_users` ADD CONSTRAINT `fk_community_user_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `community_users` ADD CONSTRAINT `fk_community_user_community` FOREIGN KEY (`community_id`) REFERENCES `communities` (`community_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `activity_likes` ADD CONSTRAINT `fk_likes_activity` FOREIGN KEY (`activity_id`) REFERENCES `activities` (`activity_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `activity_likes` ADD CONSTRAINT `fk_likes_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE `comments` ADD CONSTRAINT `fk_comment_user` FOREIGN KEY (`from_uid`) REFERENCES `users` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE `comments` ADD CONSTRAINT `fk_comment_activity` FOREIGN KEY (`activity_id`) REFERENCES `activities` (`activity_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `replies` ADD CONSTRAINT `fk_reply_comment` FOREIGN KEY (`comment_id`) REFERENCES `comments` (`comment_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `replies` ADD CONSTRAINT `fk_reply_reply` FOREIGN KEY (`to_reply_id`) REFERENCES `replies` (`reply_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE `replies` ADD CONSTRAINT `fk_reply_from_user` FOREIGN KEY (`from_uid`) REFERENCES `users` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE `replies` ADD CONSTRAINT `fk_reply_to_user` FOREIGN KEY (`to_uid`) REFERENCES `users` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE `at_users` ADD CONSTRAINT `fk_at_activity` FOREIGN KEY (`acitvity_id`) REFERENCES `activities` (`activity_id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE `at_users` ADD CONSTRAINT `fk_at_user1` FOREIGN KEY (`from_uid`) REFERENCES `users` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE `at_users` ADD CONSTRAINT `fk_at_user2` FOREIGN KEY (`to_uid`) REFERENCES `users` (`user_id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

