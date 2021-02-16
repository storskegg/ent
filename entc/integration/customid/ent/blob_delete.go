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
	"github.com/storskegg/ent/entc/integration/customid/ent/blob"
	"github.com/storskegg/ent/entc/integration/customid/ent/predicate"
	"github.com/storskegg/ent/schema/field"
)

// BlobDelete is the builder for deleting a Blob entity.
type BlobDelete struct {
	config
	hooks    []Hook
	mutation *BlobMutation
}

// Where adds a new predicate to the BlobDelete builder.
func (bd *BlobDelete) Where(ps ...predicate.Blob) *BlobDelete {
	bd.mutation.predicates = append(bd.mutation.predicates, ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BlobDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bd.hooks) == 0 {
		affected, err = bd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bd.mutation = mutation
			affected, err = bd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bd.hooks) - 1; i >= 0; i-- {
			mut = bd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BlobDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BlobDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: blob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		},
	}
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
}

// BlobDeleteOne is the builder for deleting a single Blob entity.
type BlobDeleteOne struct {
	bd *BlobDelete
}

// Exec executes the deletion query.
func (bdo *BlobDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{blob.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BlobDeleteOne) ExecX(ctx context.Context) {
	bdo.bd.ExecX(ctx)
}
