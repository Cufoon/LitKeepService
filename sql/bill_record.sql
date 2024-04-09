create or replace table litkeep.bill_record
(
    id         int auto_increment comment '主键'
        primary key,
    id_user    char(8)                              not null,
    id_kind    char(16)                             null comment '表征具体消费类型，如食品支出、衣物支出',
    type       tinyint                              not null comment '表征记录的类型：
  0 收入 Income
  1 支出 Outcome
  2 账户互转 Own Account transfer',
    value      double                               not null comment '具体的钱数',
    mark       varchar(200)                         null comment '此条交易的备注信息 comment',
    time       datetime                             not null comment '此条交易发生的时间',
    created_at datetime default current_timestamp() not null,
    updated_at datetime                             null on update current_timestamp(),
    deleted_at datetime                             null
);
