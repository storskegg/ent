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
	"github.com/storskegg/ent/entc/integration/privacy/ent/predicate"
	"github.com/storskegg/ent/entc/integration/privacy/ent/task"
	"github.com/storskegg/ent/schema/field"
)

// TaskDelete is the builder for deleting a Task entity.
type TaskDelete struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where adds a new predicate to the TaskDelete builder.
func (td *TaskDelete) Where(ps ...predicate.Task) *TaskDelete {
	td.mutation.predicates = append(td.mutation.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TaskDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(td.hooks) == 0 {
		affected, err = td.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			td.mutation = mutation
			affected, err = td.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(td.hooks) - 1; i >= 0; i-- {
			mut = td.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, td.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TaskDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TaskDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: task.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, td.driver, _spec)
}

// TaskDeleteOne is the builder for deleting a single Task entity.
type TaskDeleteOne struct {
	td *TaskDelete
}

// Exec executes the deletion query.
func (tdo *TaskDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{task.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TaskDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}
