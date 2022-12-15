create table if not exists citizen_identify_card
(
    `user_id`    bigint   not null primary key,
    `data`       text collate utf8mb4_general_ci,
    `created_at` datetime not null default now(),
    `updated_at` datetime not null default now() on update now()
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

create table if not exists user_face
(
    `user_id`    bigint   not null primary key,
    `data`       text collate utf8mb4_general_ci,
    `created_at` datetime not null default now(),
    `updated_at` datetime not null default now() on update now()
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;