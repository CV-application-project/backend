create table if not exists timekeeping_history
(
    `id`         bigint auto_increment primary key,
    `user_id`    bigint   not null,
    `day`        smallint not null,
    `month`      smallint not null,
    `year`       smallint not null,
    `is_active`  boolean           default false,
    `data`       text collate utf8mb4_general_ci,
    `created_at` datetime not null default now(),
    `updated_at` datetime not null default now() on update now(),

    unique key (user_id, `day`, `month`, `year`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;