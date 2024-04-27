CREATE TABLE IF NOT EXISTS `users`(
    `user_id` VARCHAR(64) comment 'user id',
    PRIMARY KEY(`user_id`)
);

CREATE TABLE IF NOT EXISTS `admin_users`(
    `user_id` VARCHAR(64) comment 'user id',
    `password` VARCHAR(64) NOT NULL comment 'password',
    PRIMARY KEY(`user_id`)
);

CREATE TABLE IF NOT EXISTS `chapters`(
    `chapter_id` VARCHAR(64) NOT NULL comment 'chapter id',
    `main_execute_code` TEXT NOT NULL comment 'Main class and running code',
    `init_code` TEXT NOT NULL comment 'Init code',
    `expected` TEXT NOT NULL comment 'expected',
    `answer_code` TEXT comment 'answer best practice',
    `level` INTEGER NOT NULL comment 'answer best practice',
    PRIMARY KEY(`chapter_id`)
);

CREATE TABLE IF NOT EXISTS `themes`(
    `theme_id` VARCHAR(64) comment 'theme id',
    `theme` VARCHAR(64) comment 'theme',
    `description` TEXT comment 'themes description',
    PRIMARY KEY(`theme_id`)
);

CREATE TABLE IF NOT EXISTS `themes_chapter_intersections`(
    `theme_id` VARCHAR(64) comment 'theme id',
    `chapter_id` VARCHAR(64) comment 'chapter id',
    `order` INTEGER comment 'chapter order',
    PRIMARY KEY(`theme_id`, `chapter_id`)
);