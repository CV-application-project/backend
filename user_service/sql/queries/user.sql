/* name: GetUserInfoById :one */
select * from user where id = ? limit 1;

/* name: CreateNewUserInfo :execresult */
insert into user (name, username, password, data, email) values (?, ?, ?, ?, ?);

/* name: UpdateUserInfoById :execresult */
update user set data = ? where id = ?;

/* name: GetUserByUsernameOrEmail :one */
select * from user where username = ? or email = ? limit 1;