/* name: GetTimekeepingHistoryInYearByUserId :many */
select *
from timekeeping_history
where `user_id` = ?
  and `year` = ?;

/* name: GetTimekeepingHistoryAtMonthByUserId :many */
select *
from timekeeping_history
where `user_id` = ?
  and `month` = ?
  and `year` = ?
  and is_active = true;

/* name: GetTimekeepingHistoryInDayByUserId :one */
select *
from timekeeping_history
where `user_id` = ?
  and `day` = ?
  and `month` = ?
  and `year` = ?
  and is_active = true;

/* name: CreateTimekeepingHistory :execresult */
insert into timekeeping_history (user_id, `day`, `month`, `year`, `data`, `is_active`)
VALUES (?, ?, ?, ?, ?, true);

/* name: UpdateTimekeepingHistoryInDay :execresult */
update timekeeping_history
set `data` = ?
where `user_id` = ?
  and `day` = ?
  and `month` = ?
  and `year` = ?
  and is_active = true;