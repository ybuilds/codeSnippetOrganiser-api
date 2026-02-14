package models

import (
	"database/sql"

	"ybuilds.in/codesnippet-api/database"
)

var db *sql.DB = database.DB
