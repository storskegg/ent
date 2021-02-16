// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/storskegg/ent/dialect/gremlin"
	"github.com/storskegg/ent/dialect/gremlin/graph/dsl"
	"github.com/storskegg/ent/dialect/gremlin/graph/dsl/__"
	"github.com/storskegg/ent/dialect/gremlin/graph/dsl/g"
	"github.com/storskegg/ent/entc/integration/gremlin/ent/predicate"
	"github.com/storskegg/ent/entc/integration/gremlin/ent/spec"
)

// SpecUpdate is the builder for updating Spec entities.
type SpecUpdate struct {
	config
	hooks    []Hook
	mutation *SpecMutation
}

// Where adds a new predicate for the SpecUpdate builder.
func (su *SpecUpdate) Where(ps ...predicate.Spec) *SpecUpdate {
	su.mutation.predicates = append(su.mutation.predicates, ps...)
	return su
}

// AddCardIDs adds the "card" edge to the Card entity by IDs.
func (su *SpecUpdate) AddCardIDs(ids ...string) *SpecUpdate {
	su.mutation.AddCardIDs(ids...)
	return su
}

// AddCard adds the "card" edges to the Card entity.
func (su *SpecUpdate) AddCard(c ...*Card) *SpecUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddCardIDs(ids...)
}

// Mutation returns the SpecMutation object of the builder.
func (su *SpecUpdate) Mutation() *SpecMutation {
	return su.mutation
}

// ClearCard clears all "card" edges to the Card entity.
func (su *SpecUpdate) ClearCard() *SpecUpdate {
	su.mutation.ClearCard()
	return su
}

// RemoveCardIDs removes the "card" edge to Card entities by IDs.
func (su *SpecUpdate) RemoveCardIDs(ids ...string) *SpecUpdate {
	su.mutation.RemoveCardIDs(ids...)
	return su
}

// RemoveCard removes "card" edges to Card entities.
func (su *SpecUpdate) RemoveCard(c ...*Card) *SpecUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveCardIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SpecUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.gremlinSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SpecUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SpecUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SpecUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SpecUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := su.gremlin().Query()
	if err := su.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (su *SpecUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(spec.Label)
	for _, p := range su.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	for _, id := range su.mutation.RemovedCardIDs() {
		tr := rv.Clone().OutE(spec.CardLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range su.mutation.CardIDs() {
		v.AddE(spec.CardLabel).To(g.V(id)).OutV()
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// SpecUpdateOne is the builder for updating a single Spec entity.
type SpecUpdateOne struct {
	config
	hooks    []Hook
	mutation *SpecMutation
}

// AddCardIDs adds the "card" edge to the Card entity by IDs.
func (suo *SpecUpdateOne) AddCardIDs(ids ...string) *SpecUpdateOne {
	suo.mutation.AddCardIDs(ids...)
	return suo
}

// AddCard adds the "card" edges to the Card entity.
func (suo *SpecUpdateOne) AddCard(c ...*Card) *SpecUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddCardIDs(ids...)
}

// Mutation returns the SpecMutation object of the builder.
func (suo *SpecUpdateOne) Mutation() *SpecMutation {
	return suo.mutation
}

// ClearCard clears all "card" edges to the Card entity.
func (suo *SpecUpdateOne) ClearCard() *SpecUpdateOne {
	suo.mutation.ClearCard()
	return suo
}

// RemoveCardIDs removes the "card" edge to Card entities by IDs.
func (suo *SpecUpdateOne) RemoveCardIDs(ids ...string) *SpecUpdateOne {
	suo.mutation.RemoveCardIDs(ids...)
	return suo
}

// RemoveCard removes "card" edges to Card entities.
func (suo *SpecUpdateOne) RemoveCard(c ...*Card) *SpecUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveCardIDs(ids...)
}

// Save executes the query and returns the updated Spec entity.
func (suo *SpecUpdateOne) Save(ctx context.Context) (*Spec, error) {
	var (
		err  error
		node *Spec
	)
	if len(suo.hooks) == 0 {
		node, err = suo.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpecMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SpecUpdateOne) SaveX(ctx context.Context) *Spec {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SpecUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SpecUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SpecUpdateOne) gremlinSave(ctx context.Context) (*Spec, error) {
	res := &gremlin.Response{}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Spec.ID for update")}
	}
	query, bindings := suo.gremlin(id).Query()
	if err := suo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	s := &Spec{config: suo.config}
	if err := s.FromResponse(res); err != nil {
		return nil, err
	}
	return s, nil
}

func (suo *SpecUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	for _, id := range suo.mutation.RemovedCardIDs() {
		tr := rv.Clone().OutE(spec.CardLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range suo.mutation.CardIDs() {
		v.AddE(spec.CardLabel).To(g.V(id)).OutV()
	}
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
