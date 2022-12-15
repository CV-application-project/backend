/* name: GetCICByUserId :one */
select *
from citizen_identify_card
where `user_id` = ?;

/* name: CreateCICByUserId :execresult */
insert into citizen_identify_card (user_id, data)
VALUES (?, ?);

/* name: UpdateCICByUserId :execresult */
update citizen_identify_card
set data = ?
where user_id = ?;

/* name: GetCICByUserIdForUpdate :one */
select *
from citizen_identify_card
where user_id = ? for
update;