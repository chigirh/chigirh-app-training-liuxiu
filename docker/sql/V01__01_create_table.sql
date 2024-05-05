CREATE TABLE IF NOT EXISTS `admin_users`(
    `user_id` VARCHAR(64) comment 'user id',
    `password` VARCHAR(64) NOT NULL comment 'password',
    PRIMARY KEY(`user_id`)
);

CREATE TABLE IF NOT EXISTS `users`(
    `user_id` VARCHAR(64) comment 'user id',
    `session_key` VARCHAR(64) comment 'session key',
    `theme_id` VARCHAR(64) comment 'theme id',
    PRIMARY KEY(`user_id`),
    UNIQUE (`session_key`)
);

CREATE TABLE IF NOT EXISTS `themes`(
    `theme_id` VARCHAR(64) comment 'theme id',
    `theme` VARCHAR(64) comment 'theme',
    `description` TEXT comment 'themes description',
    PRIMARY KEY(`theme_id`)
);

CREATE TABLE IF NOT EXISTS `chapters`(
    `chapter_id` VARCHAR(64) NOT NULL comment 'chapter id',
    `main_code` TEXT NOT NULL comment 'Main class and running code',
    `example_code` TEXT NOT NULL comment 'example code',
    `expected` TEXT NOT NULL comment 'expected',
    `best_practice_code` TEXT comment 'answer best practice',
    `level` INTEGER NOT NULL comment 'level',
    `exercise` TEXT comment 'exercise is markdown text',
    PRIMARY KEY(`chapter_id`)
);

CREATE TABLE IF NOT EXISTS `themes_chapter_intersections`(
    `theme_id` VARCHAR(64) comment 'theme id',
    `chapter_id` VARCHAR(64) comment 'chapter id',
    `order` INTEGER comment 'chapter order',
    PRIMARY KEY(`theme_id`, `chapter_id`)
);

CREATE TABLE IF NOT EXISTS `archivements`(
    `archivement_id` VARCHAR(64) NOT NULL comment 'archivements id',
    `user_id` VARCHAR(64) NOT NULL comment 'user id',
    `chapter_id` VARCHAR(64) NOT NULL comment 'chapter id',
    `status` CHAR(1) NOT NULL comment '1:pending,2:completed',
    PRIMARY KEY(`user_id`, `chapter_id`),
    UNIQUE(`archivement_id`)
);

CREATE TABLE IF NOT EXISTS `revisions`(
    `archivement_id` VARCHAR(64) NOT NULL comment 'archivements id',
    `version` INTEGER NOT NULL comment 'version',
    `status` CHAR(1) NOT NULL comment '1:pending,2:completed',
    `code` TEXT NOT NULL comment 'code',
    `comment` TEXT NOT NULL comment 'comment',
    `result` TEXT NOT NULL comment 'result',
    `is_compile_error` BOOLEAN NOT NULL comment 'result',
    PRIMARY KEY(`archivement_id`, `version`)
);