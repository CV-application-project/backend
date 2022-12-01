create table if not exists user_info (
    `id` bigint auto_increment primary key,
    `name` nvarchar(100) not null,
    `extra_data` text COLLATE utf8mb4_general_ci,
    `created_at` datetime not null default now(),
    `updated_at` datetime not null default now() on update now()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;