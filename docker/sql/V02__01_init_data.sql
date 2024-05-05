INSERT INTO
    admin_users(`user_id`, `password`)
VALUES
    ('admin', 'admin');

INSERT INTO
    themes(`theme_id`, `theme`, `description`)
VALUES
    ('THEME-01', 'kensyu', "kensyu dayo");

INSERT INTO
    users(`user_id`, `session_key`, `theme_id`)
VALUES
    ('user', 'xxxxxxxxxx-01', "THEME-01");

INSERT INTO
    themes_chapter_intersections(`theme_id`, `chapter_id`, `order`)
VALUES
    ('THEME-01', 'CHAPTER-AA-01', 1),
    ('THEME-01', 'CHAPTER-AA-02', 2),
    ('THEME-01', 'CHAPTER-AA-03', 3),
    ('THEME-01', 'CHAPTER-AA-04', 4),
    ('THEME-01', 'CHAPTER-AA-05', 5);