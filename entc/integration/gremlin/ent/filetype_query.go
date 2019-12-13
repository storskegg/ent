// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"math"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
)

// FileTypeQuery is the builder for querying FileType entities.
type FileTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.FileType
	// intermediate query.
	gremlin *dsl.Traversal
}

// Where adds a new predicate for the builder.
func (ftq *FileTypeQuery) Where(ps ...predicate.FileType) *FileTypeQuery {
	ftq.predicates = append(ftq.predicates, ps...)
	return ftq
}

// Limit adds a limit step to the query.
func (ftq *FileTypeQuery) Limit(limit int) *FileTypeQuery {
	ftq.limit = &limit
	return ftq
}

// Offset adds an offset step to the query.
func (ftq *FileTypeQuery) Offset(offset int) *FileTypeQuery {
	ftq.offset = &offset
	return ftq
}

// Order adds an order step to the query.
func (ftq *FileTypeQuery) Order(o ...Order) *FileTypeQuery {
	ftq.order = append(ftq.order, o...)
	return ftq
}

// QueryFiles chains the current query on the files edge.
func (ftq *FileTypeQuery) QueryFiles() *FileQuery {
	query := &FileQuery{config: ftq.config}
	gremlin := ftq.gremlinQuery()
	query.gremlin = gremlin.OutE(filetype.FilesLabel).InV()
	return query
}

// First returns the first FileType entity in the query. Returns *ErrNotFound when no filetype was found.
func (ftq *FileTypeQuery) First(ctx context.Context) (*FileType, error) {
	fts, err := ftq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(fts) == 0 {
		return nil, &ErrNotFound{filetype.Label}
	}
	return fts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftq *FileTypeQuery) FirstX(ctx context.Context) *FileType {
	ft, err := ftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ft
}

// FirstID returns the first FileType id in the query. Returns *ErrNotFound when no id was found.
func (ftq *FileTypeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ftq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{filetype.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (ftq *FileTypeQuery) FirstXID(ctx context.Context) string {
	id, err := ftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only FileType entity in the query, returns an error if not exactly one entity was returned.
func (ftq *FileTypeQuery) Only(ctx context.Context) (*FileType, error) {
	fts, err := ftq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(fts) {
	case 1:
		return fts[0], nil
	case 0:
		return nil, &ErrNotFound{filetype.Label}
	default:
		return nil, &ErrNotSingular{filetype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftq *FileTypeQuery) OnlyX(ctx context.Context) *FileType {
	ft, err := ftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ft
}

// OnlyID returns the only FileType id in the query, returns an error if not exactly one id was returned.
func (ftq *FileTypeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ftq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{filetype.Label}
	default:
		err = &ErrNotSingular{filetype.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (ftq *FileTypeQuery) OnlyXID(ctx context.Context) string {
	id, err := ftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FileTypes.
func (ftq *FileTypeQuery) All(ctx context.Context) ([]*FileType, error) {
	return ftq.gremlinAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ftq *FileTypeQuery) AllX(ctx context.Context) []*FileType {
	fts, err := ftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return fts
}

// IDs executes the query and returns a list of FileType ids.
func (ftq *FileTypeQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := ftq.Select(filetype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ftq *FileTypeQuery) IDsX(ctx context.Context) []string {
	ids, err := ftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftq *FileTypeQuery) Count(ctx context.Context) (int, error) {
	return ftq.gremlinCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ftq *FileTypeQuery) CountX(ctx context.Context) int {
	count, err := ftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftq *FileTypeQuery) Exist(ctx context.Context) (bool, error) {
	return ftq.gremlinExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ftq *FileTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ftq *FileTypeQuery) Clone() *FileTypeQuery {
	return &FileTypeQuery{
		config:     ftq.config,
		limit:      ftq.limit,
		offset:     ftq.offset,
		order:      append([]Order{}, ftq.order...),
		unique:     append([]string{}, ftq.unique...),
		predicates: append([]predicate.FileType{}, ftq.predicates...),
		// clone intermediate query.
		gremlin: ftq.gremlin.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FileType.Query().
//		GroupBy(filetype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ftq *FileTypeQuery) GroupBy(field string, fields ...string) *FileTypeGroupBy {
	group := &FileTypeGroupBy{config: ftq.config}
	group.fields = append([]string{field}, fields...)
	group.gremlin = ftq.gremlinQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.FileType.Query().
//		Select(filetype.FieldName).
//		Scan(ctx, &v)
//
func (ftq *FileTypeQuery) Select(field string, fields ...string) *FileTypeSelect {
	selector := &FileTypeSelect{config: ftq.config}
	selector.fields = append([]string{field}, fields...)
	selector.gremlin = ftq.gremlinQuery()
	return selector
}

func (ftq *FileTypeQuery) gremlinAll(ctx context.Context) ([]*FileType, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery().ValueMap(true).Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var fts FileTypes
	if err := fts.FromResponse(res); err != nil {
		return nil, err
	}
	fts.config(ftq.config)
	return fts, nil
}

func (ftq *FileTypeQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery().Count().Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (ftq *FileTypeQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery().HasNext().Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (ftq *FileTypeQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(filetype.Label)
	if ftq.gremlin != nil {
		v = ftq.gremlin.Clone()
	}
	for _, p := range ftq.predicates {
		p(v)
	}
	if len(ftq.order) > 0 {
		v.Order()
		for _, p := range ftq.order {
			p(v)
		}
	}
	switch limit, offset := ftq.limit, ftq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := ftq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// FileTypeGroupBy is the builder for group-by FileType entities.
type FileTypeGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	gremlin *dsl.Traversal
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftgb *FileTypeGroupBy) Aggregate(fns ...Aggregate) *FileTypeGroupBy {
	ftgb.fns = append(ftgb.fns, fns...)
	return ftgb
}

// Scan applies the group-by query and scan the result into the given value.
func (ftgb *FileTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	return ftgb.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ftgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FileTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FileTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := ftgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FileTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FileTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := ftgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FileTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FileTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ftgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FileTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FileTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ftgb *FileTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ftgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftgb *FileTypeGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := ftgb.gremlinQuery().Query()
	if err := ftgb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ftgb.fields)+len(ftgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (ftgb *FileTypeGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range ftgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range ftgb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return ftgb.gremlin.Group().
		By(__.Values(ftgb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}

// FileTypeSelect is the builder for select fields of FileType entities.
type FileTypeSelect struct {
	config
	fields []string
	// intermediate queries.
	gremlin *dsl.Traversal
}

// Scan applies the selector query and scan the result into the given value.
func (fts *FileTypeSelect) Scan(ctx context.Context, v interface{}) error {
	return fts.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fts *FileTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FileTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fts *FileTypeSelect) StringsX(ctx context.Context) []string {
	v, err := fts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FileTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fts *FileTypeSelect) IntsX(ctx context.Context) []int {
	v, err := fts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FileTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fts *FileTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (fts *FileTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FileTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fts *FileTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := fts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fts *FileTypeSelect) gremlinScan(ctx context.Context, v interface{}) error {
	var (
		traversal *dsl.Traversal
		res       = &gremlin.Response{}
	)
	if len(fts.fields) == 1 {
		if fts.fields[0] != filetype.FieldID {
			traversal = fts.gremlin.Values(fts.fields...)
		} else {
			traversal = fts.gremlin.ID()
		}
	} else {
		fields := make([]interface{}, len(fts.fields))
		for i, f := range fts.fields {
			fields[i] = f
		}
		traversal = fts.gremlin.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := fts.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(fts.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
