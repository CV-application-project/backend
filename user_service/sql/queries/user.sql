/* name: GetUserInfoById :one */
select *
from user
where id = ?
limit 1;

/* name: CreateNewUserInfo :execresult */
insert into user (name, employee_id, password, phone, address, gender, department, position, role, data, email,
                  front_card, back_card) value (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

/* name: UpdateUserInfoById :execresult */
update user
set phone      = ?,
    address    = ?,
    department = ?,
    position   = ?,
    role       = ?,
    front_card = ?,
    back_card = ?
where id = ?;

/* name: GetUserByUsernameOrEmail :one */
select *
from user
where employee_id = ?
   or email = ?
limit 1;

/* name: GetUsersByDepartment :many */
select *
from user
where department = ?
  and role = 'STAFF';

/* name: GetAllUsers :many */
select *
from user
where role != 'HR';