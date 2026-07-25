-- name: CreateStudent :execresult
INSERT INTO students(
    name,
    age,
    course,
    marks
)
VALUES(
    ?,
    ?,
    ?,
    ?
);


-- name: GetStudents :many
SELECT *
FROM students;


-- name: GetStudentByID :one
SELECT *
FROM students
WHERE id=?;


-- name: UpdateStudent :exec
UPDATE students
SET
    name=?,
    age=?,
    course=?,
    marks=?
WHERE id=?;


-- name: DeleteStudent :exec
DELETE FROM students
WHERE id=?;