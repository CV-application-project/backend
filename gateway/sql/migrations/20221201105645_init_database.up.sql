create table if not exists user_info
(
    `id`          bigint auto_increment primary key,
    `user_id`     bigint        not null,
    `employee_id` nvarchar(255) not null,
    `role`        nvarchar(20)           default 'STAFF',
    `email`       nvarchar(255) not null,
    `token`       text          not null collate utf8mb4_general_ci,
    `created_at`  datetime      not null default now(),
    `updated_at`  datetime      not null default now() on update now(),
    `expired_at`  datetime      not null
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

alter table user_info
    add unique key (employee_id);
alter table user_info
    add unique key (user_id);
alter table user_info
    add unique key (email);