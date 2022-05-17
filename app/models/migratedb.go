package models

/*
  @Author : Mustang Kong
*/

import (
	"golang-base-flamego/app/models/auth"
	"golang-base-flamego/app/models/email"
	"golang-base-flamego/pkg/connection"
)

func AutoMigrateTable() {
	connection.DB.Self.AutoMigrate(
		// auth
		&auth.User{},

		&email.EmailTextContent{},
	)
}
