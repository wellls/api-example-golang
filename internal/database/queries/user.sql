-- name: GetUserByID :one
select * from users u where u.id = $1;