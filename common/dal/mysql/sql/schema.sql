CREATE TABLE `test_user` (
    `id` bigint unsigned NOT NULL COMMENT 'id',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT= 'user';

CREATE TABLE `users` (
  `user_id` bigint unsigned NOT NULL,
  `username` longtext,
  `password` longtext,
  `nick_name` longtext,
  `follower_cnt` bigint DEFAULT NULL,
  `follow_cnt` bigint DEFAULT NULL,
  `register_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT= 'users';

CREATE TABLE `likes` (
  `key_id` bigint unsigned NOT NULL,
  `owner_id` bigint unsigned DEFAULT NULL,
  `video_id` bigint unsigned DEFAULT NULL,
  `like_time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`key_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT= 'likes';

CREATE TABLE `videos` (
  `video_id` bigint unsigned NOT NULL,
  `auther_id` bigint unsigned DEFAULT NULL,
  `play_url` longtext,
  `cover_url` longtext,
  `like_count` bigint DEFAULT NULL,
  `comment_count` bigint DEFAULT NULL,
  `title` longtext,
  `abstract` longtext,
  PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='videos';