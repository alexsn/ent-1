// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/ent/user"

	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "number", Type: field.TypeString},
		{Name: "owner_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "cards_users_card",
				Columns: []*schema.Column{CardsColumns[4]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "unique_int", Type: field.TypeInt, Unique: true},
		{Name: "unique_float", Type: field.TypeFloat64, Unique: true},
		{Name: "nillable_int", Type: field.TypeInt, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:        "comments",
		Columns:     CommentsColumns,
		PrimaryKey:  []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FieldTypesColumns holds the columns for the "field_types" table.
	FieldTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "int", Type: field.TypeInt},
		{Name: "int8", Type: field.TypeInt8},
		{Name: "int16", Type: field.TypeInt16},
		{Name: "int32", Type: field.TypeInt32},
		{Name: "int64", Type: field.TypeInt64},
		{Name: "optional_int", Type: field.TypeInt, Nullable: true},
		{Name: "optional_int8", Type: field.TypeInt8, Nullable: true},
		{Name: "optional_int16", Type: field.TypeInt16, Nullable: true},
		{Name: "optional_int32", Type: field.TypeInt32, Nullable: true},
		{Name: "optional_int64", Type: field.TypeInt64, Nullable: true},
		{Name: "nillable_int", Type: field.TypeInt, Nullable: true},
		{Name: "nillable_int8", Type: field.TypeInt8, Nullable: true},
		{Name: "nillable_int16", Type: field.TypeInt16, Nullable: true},
		{Name: "nillable_int32", Type: field.TypeInt32, Nullable: true},
		{Name: "nillable_int64", Type: field.TypeInt64, Nullable: true},
		{Name: "validate_optional_int32", Type: field.TypeInt32, Nullable: true},
		{Name: "state", Type: field.TypeEnum, Nullable: true, Enums: []string{"on", "off"}},
	}
	// FieldTypesTable holds the schema information for the "field_types" table.
	FieldTypesTable = &schema.Table{
		Name:        "field_types",
		Columns:     FieldTypesColumns,
		PrimaryKey:  []*schema.Column{FieldTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "size", Type: field.TypeInt, Default: file.DefaultSize},
		{Name: "name", Type: field.TypeString},
		{Name: "user", Type: field.TypeString, Nullable: true},
		{Name: "group", Type: field.TypeString, Nullable: true},
		{Name: "type_id", Type: field.TypeInt, Nullable: true},
		{Name: "group_file_id", Type: field.TypeInt, Nullable: true},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "files_file_types_files",
				Columns: []*schema.Column{FilesColumns[5]},

				RefColumns: []*schema.Column{FileTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "files_groups_files",
				Columns: []*schema.Column{FilesColumns[6]},

				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "files_users_files",
				Columns: []*schema.Column{FilesColumns[7]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "name_size",
				Unique:  false,
				Columns: []*schema.Column{FilesColumns[2], FilesColumns[1]},
			},
			{
				Name:    "name_user",
				Unique:  true,
				Columns: []*schema.Column{FilesColumns[2], FilesColumns[3]},
			},
			{
				Name:    "name_owner_id_type_id",
				Unique:  true,
				Columns: []*schema.Column{FilesColumns[2], FilesColumns[7], FilesColumns[5]},
			},
		},
	}
	// FileTypesColumns holds the columns for the "file_types" table.
	FileTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// FileTypesTable holds the schema information for the "file_types" table.
	FileTypesTable = &schema.Table{
		Name:        "file_types",
		Columns:     FileTypesColumns,
		PrimaryKey:  []*schema.Column{FileTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "active", Type: field.TypeBool, Default: group.DefaultActive},
		{Name: "expire", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString, Nullable: true},
		{Name: "max_users", Type: field.TypeInt, Nullable: true, Default: group.DefaultMaxUsers},
		{Name: "name", Type: field.TypeString},
		{Name: "info_id", Type: field.TypeInt, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "groups_group_infos_info",
				Columns: []*schema.Column{GroupsColumns[6]},

				RefColumns: []*schema.Column{GroupInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupInfosColumns holds the columns for the "group_infos" table.
	GroupInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "desc", Type: field.TypeString},
		{Name: "max_users", Type: field.TypeInt, Default: groupinfo.DefaultMaxUsers},
	}
	// GroupInfosTable holds the schema information for the "group_infos" table.
	GroupInfosTable = &schema.Table{
		Name:        "group_infos",
		Columns:     GroupInfosColumns,
		PrimaryKey:  []*schema.Column{GroupInfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:        "items",
		Columns:     ItemsColumns,
		PrimaryKey:  []*schema.Column{ItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// NodesColumns holds the columns for the "nodes" table.
	NodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeInt, Nullable: true},
		{Name: "prev_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// NodesTable holds the schema information for the "nodes" table.
	NodesTable = &schema.Table{
		Name:       "nodes",
		Columns:    NodesColumns,
		PrimaryKey: []*schema.Column{NodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "nodes_nodes_next",
				Columns: []*schema.Column{NodesColumns[2]},

				RefColumns: []*schema.Column{NodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true},
		{Name: "team_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "pets_users_pets",
				Columns: []*schema.Column{PetsColumns[2]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "pets_users_team",
				Columns: []*schema.Column{PetsColumns[3]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString},
		{Name: "last", Type: field.TypeString, Default: user.DefaultLast},
		{Name: "nickname", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "phone", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "group_blocked_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_spouse_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_groups_blocked",
				Columns: []*schema.Column{UsersColumns[6]},

				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_users_spouse",
				Columns: []*schema.Column{UsersColumns[7]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "users_users_parent",
				Columns: []*schema.Column{UsersColumns[8]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0], UserGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_groups_user_id",
				Columns: []*schema.Column{UserGroupsColumns[0]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "user_groups_group_id",
				Columns: []*schema.Column{UserGroupsColumns[1]},

				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserFriendsColumns holds the columns for the "user_friends" table.
	UserFriendsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// UserFriendsTable holds the schema information for the "user_friends" table.
	UserFriendsTable = &schema.Table{
		Name:       "user_friends",
		Columns:    UserFriendsColumns,
		PrimaryKey: []*schema.Column{UserFriendsColumns[0], UserFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_friends_user_id",
				Columns: []*schema.Column{UserFriendsColumns[0]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "user_friends_friend_id",
				Columns: []*schema.Column{UserFriendsColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserFollowingColumns holds the columns for the "user_following" table.
	UserFollowingColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "follower_id", Type: field.TypeInt},
	}
	// UserFollowingTable holds the schema information for the "user_following" table.
	UserFollowingTable = &schema.Table{
		Name:       "user_following",
		Columns:    UserFollowingColumns,
		PrimaryKey: []*schema.Column{UserFollowingColumns[0], UserFollowingColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_following_user_id",
				Columns: []*schema.Column{UserFollowingColumns[0]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "user_following_follower_id",
				Columns: []*schema.Column{UserFollowingColumns[1]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CardsTable,
		CommentsTable,
		FieldTypesTable,
		FilesTable,
		FileTypesTable,
		GroupsTable,
		GroupInfosTable,
		ItemsTable,
		NodesTable,
		PetsTable,
		UsersTable,
		UserGroupsTable,
		UserFriendsTable,
		UserFollowingTable,
	}
)

func init() {
	CardsTable.ForeignKeys[0].RefTable = UsersTable
	FilesTable.ForeignKeys[0].RefTable = FileTypesTable
	FilesTable.ForeignKeys[1].RefTable = GroupsTable
	FilesTable.ForeignKeys[2].RefTable = UsersTable
	GroupsTable.ForeignKeys[0].RefTable = GroupInfosTable
	NodesTable.ForeignKeys[0].RefTable = NodesTable
	PetsTable.ForeignKeys[0].RefTable = UsersTable
	PetsTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = GroupsTable
	UsersTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[2].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[1].RefTable = GroupsTable
	UserFriendsTable.ForeignKeys[0].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[1].RefTable = UsersTable
	UserFollowingTable.ForeignKeys[0].RefTable = UsersTable
	UserFollowingTable.ForeignKeys[1].RefTable = UsersTable
}
