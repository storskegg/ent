// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/storskegg/ent/dialect/sql/sqlgraph"
	"github.com/storskegg/ent/examples/edgeindex/ent/city"
	"github.com/storskegg/ent/examples/edgeindex/ent/street"
	"github.com/storskegg/ent/schema/field"
)

// StreetCreate is the builder for creating a Street entity.
type StreetCreate struct {
	config
	mutation *StreetMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *StreetCreate) SetName(s string) *StreetCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetCityID sets the "city" edge to the City entity by ID.
func (sc *StreetCreate) SetCityID(id int) *StreetCreate {
	sc.mutation.SetCityID(id)
	return sc
}

// SetNillableCityID sets the "city" edge to the City entity by ID if the given value is not nil.
func (sc *StreetCreate) SetNillableCityID(id *int) *StreetCreate {
	if id != nil {
		sc = sc.SetCityID(*id)
	}
	return sc
}

// SetCity sets the "city" edge to the City entity.
func (sc *StreetCreate) SetCity(c *City) *StreetCreate {
	return sc.SetCityID(c.ID)
}

// Mutation returns the StreetMutation object of the builder.
func (sc *StreetCreate) Mutation() *StreetMutation {
	return sc.mutation
}

// Save creates the Street in the database.
func (sc *StreetCreate) Save(ctx context.Context) (*Street, error) {
	var (
		err  error
		node *Street
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StreetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StreetCreate) SaveX(ctx context.Context) *Street {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (sc *StreetCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	return nil
}

func (sc *StreetCreate) sqlSave(ctx context.Context) (*Street, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *StreetCreate) createSpec() (*Street, *sqlgraph.CreateSpec) {
	var (
		_node = &Street{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: street.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: street.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: street.FieldName,
		})
		_node.Name = value
	}
	if nodes := sc.mutation.CityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   street.CityTable,
			Columns: []string{street.CityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: city.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StreetCreateBulk is the builder for creating many Street entities in bulk.
type StreetCreateBulk struct {
	config
	builders []*StreetCreate
}

// Save creates the Street entities in the database.
func (scb *StreetCreateBulk) Save(ctx context.Context) ([]*Street, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Street, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StreetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StreetCreateBulk) SaveX(ctx context.Context) []*Street {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
