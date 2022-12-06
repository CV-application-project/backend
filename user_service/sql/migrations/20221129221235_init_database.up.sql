create table if not exists user
(
    `id`         bigint auto_increment primary key,
    `name`       nvarchar(255) not null,
    `username`   nvarchar(255) not null,
    `password`   blob          not null COLLATE utf8mb4_general_ci,
    `data`       text COLLATE utf8mb4_general_ci,
    `created_at` timestamp     not null default now(),
    `updated_at` timestamp     not null default now() on update now()
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

create table if not exists user_token
(
    `id`         bigint auto_increment primary key,
    `user_id`    bigint    not null,
    `token`      text      not null collate utf8mb4_general_ci,
    `created_at` timestamp not null default now(),
    `expired_at` timestamp not null
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;