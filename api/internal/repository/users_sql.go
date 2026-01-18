package repository

import _ "embed"

//go:embed sql/users_create.sql
var queryCreateUser string

//go:embed sql/users_get_by_id.sql
var queryGetUserByID string

//go:embed sql/users_get_all.sql
var queryGetAllUsers string
