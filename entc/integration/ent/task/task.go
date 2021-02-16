// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package task

import (
	"github.com/storskegg/ent/entc/integration/ent/schema"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"

	// Table holds the table name of the task in the database.
	Table = "tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldPriority,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultPriority holds the default value on creation for the "priority" field.
	DefaultPriority schema.Priority
	// PriorityValidator is a validator for the "priority" field. It is called by the builders before save.
	PriorityValidator func(int) error
)

// comment from another template.
