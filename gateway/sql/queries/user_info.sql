/* name: GetUserInfoById :one */
select * from user_info where id = ? limit 1;

/* name: CreateNewUserInfo :execresult */
insert into user_info (name, username, password, extra_data, email) VALUES (?, ?, ?, ?, ?);

/* name: UpdateUserInfoByUsername :execresult */
update user_info set extra_data = ? where username = ?;