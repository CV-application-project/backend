/* name: GetFaceByUserId :one */
select *
from user_face
where `user_id` = ?;

/* name: CreateFaceByUserId :execresult */
insert into user_face (user_id, data)
VALUES (?, ?);

/* name: UpdateFaceByUserId :execresult */
update user_face
set data = ?
where user_id = ?;

/* name: GetFaceByUserIdForUpdate :one */
select *
from user_face
where user_id = ? for
update;