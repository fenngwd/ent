// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package group

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldExpire holds the string denoting the expire field in the database.
	FieldExpire = "expire"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldMaxUsers holds the string denoting the max_users field in the database.
	FieldMaxUsers = "max_users"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// EdgeBlocked holds the string denoting the blocked edge name in mutations.
	EdgeBlocked = "blocked"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeInfo holds the string denoting the info edge name in mutations.
	EdgeInfo = "info"

	// Table holds the table name of the group in the database.
	Table = "groups"
	// FilesTable is the table the holds the files relation/edge.
	FilesTable = "files"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "files"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "group_files"
	// BlockedTable is the table the holds the blocked relation/edge.
	BlockedTable = "users"
	// BlockedInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	BlockedInverseTable = "users"
	// BlockedColumn is the table column denoting the blocked relation/edge.
	BlockedColumn = "group_blocked"
	// UsersTable is the table the holds the users relation/edge. The primary key declared below.
	UsersTable = "user_groups"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// InfoTable is the table the holds the info relation/edge.
	InfoTable = "groups"
	// InfoInverseTable is the table name for the GroupInfo entity.
	// It exists in this package in order to avoid circular dependency with the "groupinfo" package.
	InfoInverseTable = "group_infos"
	// InfoColumn is the table column denoting the info relation/edge.
	InfoColumn = "group_info"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldActive,
	FieldExpire,
	FieldType,
	FieldMaxUsers,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Group type.
var ForeignKeys = []string{
	"group_info",
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "group_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultActive holds the default value on creation for the active field.
	DefaultActive bool
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultMaxUsers holds the default value on creation for the max_users field.
	DefaultMaxUsers int
	// MaxUsersValidator is a validator for the "max_users" field. It is called by the builders before save.
	MaxUsersValidator func(int) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// comment from another template.
