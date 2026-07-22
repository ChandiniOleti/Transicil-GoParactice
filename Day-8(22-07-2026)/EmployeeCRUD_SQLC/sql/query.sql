-- name: CreateEmployee :exec

INSERT INTO employeesqlc(

name,
age,
department,
salary,
city

)

VALUES(?,?,?,?,?);


-- name: ListEmployees :many

SELECT *

FROM employeesqlc;

-- name: GetEmployee :one

SELECT *

FROM employeesqlc

WHERE id=?;


-- name: UpdateEmployee :exec

UPDATE employeesqlc

SET

name=?,
age=?,
department=?,
salary=?,
city=?

WHERE id=?;



-- name: DeleteEmployee :exec

DELETE

FROM employeesqlc

WHERE id=?;