/* name: GetUserTokenByUserId :one */
select * from user_token where user_id = ?;

/* name: CreateNewUserTokenByUserId :execresult */
insert into user_token (user_id, token, expired_at) values (?, ?, ?);