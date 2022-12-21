/* name: GetUserInfoByUsernameOrEmail :one */
select *
from user_info
where username = ?
   or email = ?
limit 1;

/* name: GetUserInfoByToken :one */
select *
from user_info
where token = ?
  and expired_at > now();

/* name: UpdateUserInfoTokenByUserId :execresult */
update user_info
set token      = ?,
    expired_at = ?
where user_id = ?;

/* name: CreateUserInfo :execresult */
insert into user_info (user_id, username, email, token, expired_at)
VALUES (?, ?, ?, ?, ?);