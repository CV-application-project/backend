/* name: GetUserInfoByUsernameOrEmail :one */
select *
from user_info
where employee_id = ?
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
insert into user_info (user_id, employee_id, role, email, token, expired_at, department)
VALUES (?, ?, ?, ?, ?, ?, ?);

/* name: UpdateUserCard :execresult */
update user_info
set front_card = ?,
    back_card  = ?
where user_id = ?;