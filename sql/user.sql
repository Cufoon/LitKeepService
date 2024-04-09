create or replace table litkeep.user
(
    id         int auto_increment
        primary key,
    id_user    char(8)                              not null comment '8位随机字符串',
    passwd     varchar(20)                          null,
    email      varchar(200)                         not null,
    nick_name  varchar(255)                         null,
    icon       longblob                             null,
    status     char                                 null comment '用户状态，包括正常和禁止登录，默认NULL或者''A''为状态正常Active，''B''为禁止登录Forbidden',
    created_at datetime default current_timestamp() not null,
    updated_at datetime                             null on update current_timestamp(),
    deleted_at datetime                             null
);

INSERT INTO litkeep.user (id, id_user, passwd, email, nick_name, icon, status, created_at, updated_at, deleted_at) VALUES (1, 'LitAdmin', null, 'admin', null, null, 'B', '2024-01-01 00:00:00', null, null);
INSERT INTO litkeep.user (id, id_user, passwd, email, nick_name, icon, status, created_at, updated_at, deleted_at) VALUES (2, 'TEST0001', null, 'test', 'Test Account', null, null, '2024-01-01 00:00:00', null, null);
