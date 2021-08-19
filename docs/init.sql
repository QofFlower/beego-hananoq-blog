drop table blog_article;
CREATE TABLE blog_article
(
    id                 int NOT NULL,
    article_head_pic   varchar(255) DEFAULT NULL,
    article_name       varchar(255) DEFAULT NULL,
    article_tag        varchar(255) DEFAULT NULL,
    article_remark     varchar(255) DEFAULT NULL,
    article_read_count int NULL     DEFAULT 0,
    article_state      int NULL     DEFAULT 0,
    manager_id         int NULL     DEFAULT NULL,
    manager_name       varchar(50),
    article_content    text,
    create_time        varchar(255),
    article_type       int NULL     DEFAULT NULL,
    article_star_num   int NULL     DEFAULT NULL,
    article_con_num    int NULL     DEFAULT NULL,
    enclosure          varchar(255)
);
-- create index article_id_index on blog_article(id);
-- create index article_name_index on blog_article(article_name);
CREATE SEQUENCE article_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

alter table "hananoq_blog".hananoq_blog."blog_article"
    alter column id set default nextval('article_id_seq');

------
CREATE TABLE blog_category
(
    id   int          NOT NULL,
    name varchar(255) NULL DEFAULT NULL
);
-- create index category_id_index on blog_category(id)
CREATE SEQUENCE category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

alter table "hananoq_blog".hananoq_blog."blog_category"
    alter column id set default nextval('category_id_seq');


------
CREATE TABLE blog_comment
(
    id            int          NOT NULL,
    article_id    int          NULL DEFAULT NULL,
    content       varchar(255) NULL DEFAULT NULL,
    create_time   varchar(255) NULL DEFAULT NULL,
    by_manager_id int          NULL DEFAULT NULL,
    pid           int          NULL DEFAULT NULL,
    nickname      varchar(255) NULL DEFAULT NULL
);
-- create index comment_id_index on blog_comment(id);
CREATE SEQUENCE comment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

alter table "hananoq_blog".hananoq_blog."blog_comment"
    alter column id set default nextval('comment_id_seq');


------
drop table blog_user;
CREATE TABLE blog_user
(
    id           int          NOT NULL,
    name         varchar(255) NULL DEFAULT NULL,
    username     varchar(255) NULL DEFAULT NULL,
    agi_password varchar(255) NULL DEFAULT NULL,
    password     varchar(255) NULL DEFAULT NULL,
    head_pic     varchar(255) NULL DEFAULT NULL,
    create_time  timestamptz,
    update_time  timestamptz,
    status       int2         NULL DEFAULT 1,
    type         varchar(255) NULL DEFAULT NULL
);
CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

alter table "hananoq_blog".hananoq_blog."blog_user"
    alter column id set default nextval('user_id_seq');

INSERT INTO "hananoq_blog".hananoq_blog."blog_user"("username", "password", "status", "create_time", "update_time") VALUES ('c5898c3cc9fc32e7eef24686204d10fb', 'c4d038b4bed09fdb1471ef51ec3a32cd', 1, '2021-08-13 02:03:50.751426+00', '2021-08-13 02:03:50.751426+00');