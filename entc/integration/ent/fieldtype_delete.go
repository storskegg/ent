// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/storskegg/ent/dialect/sql"
	"github.com/storskegg/ent/dialect/sql/sqlgraph"
	"github.com/storskegg/ent/entc/integration/ent/fieldtype"
	"github.com/storskegg/ent/entc/integration/ent/predicate"
	"github.com/storskegg/ent/schema/field"
)

// FieldTypeDelete is the builder for deleting a FieldType entity.
type FieldTypeDelete struct {
	config
	hooks    []Hook
	mutation *FieldTypeMutation
}

// Where adds a new predicate to the FieldTypeDelete builder.
func (ftd *FieldTypeDelete) Where(ps ...predicate.FieldType) *FieldTypeDelete {
	ftd.mutation.predicates = append(ftd.mutation.predicates, ps...)
	return ftd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ftd *FieldTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ftd.hooks) == 0 {
		affected, err = ftd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FieldTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftd.mutation = mutation
			affected, err = ftd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ftd.hooks) - 1; i >= 0; i-- {
			mut = ftd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FieldTypeDelete) ExecX(ctx context.Context) int {
	n, err := ftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftd *FieldTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: fieldtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fieldtype.FieldID,
			},
		},
	}
	if ps := ftd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ftd.driver, _spec)
}

// FieldTypeDeleteOne is the builder for deleting a single FieldType entity.
type FieldTypeDeleteOne struct {
	ftd *FieldTypeDelete
}

// Exec executes the deletion query.
func (ftdo *FieldTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ftdo.ftd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fieldtype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FieldTypeDeleteOne) ExecX(ctx context.Context) {
	ftdo.ftd.ExecX(ctx)
}
