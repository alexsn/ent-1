// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/item"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// ItemQuery is the builder for querying Item entities.
type ItemQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Item
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Where adds a new predicate for the builder.
func (iq *ItemQuery) Where(ps ...predicate.Item) *ItemQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *ItemQuery) Limit(limit int) *ItemQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *ItemQuery) Offset(offset int) *ItemQuery {
	iq.offset = &offset
	return iq
}

// Order adds an order step to the query.
func (iq *ItemQuery) Order(o ...Order) *ItemQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// First returns the first Item entity in the query. Returns *ErrNotFound when no item was found.
func (iq *ItemQuery) First(ctx context.Context) (*Item, error) {
	is, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(is) == 0 {
		return nil, &ErrNotFound{item.Label}
	}
	return is[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *ItemQuery) FirstX(ctx context.Context) *Item {
	i, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return i
}

// FirstID returns the first Item id in the query. Returns *ErrNotFound when no id was found.
func (iq *ItemQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{item.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (iq *ItemQuery) FirstXID(ctx context.Context) string {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Item entity in the query, returns an error if not exactly one entity was returned.
func (iq *ItemQuery) Only(ctx context.Context) (*Item, error) {
	is, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(is) {
	case 1:
		return is[0], nil
	case 0:
		return nil, &ErrNotFound{item.Label}
	default:
		return nil, &ErrNotSingular{item.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *ItemQuery) OnlyX(ctx context.Context) *Item {
	i, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return i
}

// OnlyID returns the only Item id in the query, returns an error if not exactly one id was returned.
func (iq *ItemQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{item.Label}
	default:
		err = &ErrNotSingular{item.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (iq *ItemQuery) OnlyXID(ctx context.Context) string {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Items.
func (iq *ItemQuery) All(ctx context.Context) ([]*Item, error) {
	switch iq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return iq.sqlAll(ctx)
	case dialect.Gremlin:
		return iq.gremlinAll(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// AllX is like All, but panics if an error occurs.
func (iq *ItemQuery) AllX(ctx context.Context) []*Item {
	is, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return is
}

// IDs executes the query and returns a list of Item ids.
func (iq *ItemQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := iq.Select(item.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *ItemQuery) IDsX(ctx context.Context) []string {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *ItemQuery) Count(ctx context.Context) (int, error) {
	switch iq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return iq.sqlCount(ctx)
	case dialect.Gremlin:
		return iq.gremlinCount(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// CountX is like Count, but panics if an error occurs.
func (iq *ItemQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *ItemQuery) Exist(ctx context.Context) (bool, error) {
	switch iq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return iq.sqlExist(ctx)
	case dialect.Gremlin:
		return iq.gremlinExist(ctx)
	default:
		return false, errors.New("ent: unsupported dialect")
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *ItemQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *ItemQuery) Clone() *ItemQuery {
	return &ItemQuery{
		config:     iq.config,
		limit:      iq.limit,
		offset:     iq.offset,
		order:      append([]Order{}, iq.order...),
		unique:     append([]string{}, iq.unique...),
		predicates: append([]predicate.Item{}, iq.predicates...),
		// clone intermediate queries.
		sql:     iq.sql.Clone(),
		gremlin: iq.gremlin.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (iq *ItemQuery) GroupBy(field string, fields ...string) *ItemGroupBy {
	group := &ItemGroupBy{config: iq.config}
	group.fields = append([]string{field}, fields...)
	switch iq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		group.sql = iq.sqlQuery()
	case dialect.Gremlin:
		group.gremlin = iq.gremlinQuery()
	}
	return group
}

// Select one or more fields from the given query.
func (iq *ItemQuery) Select(field string, fields ...string) *ItemSelect {
	selector := &ItemSelect{config: iq.config}
	selector.fields = append([]string{field}, fields...)
	switch iq.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		selector.sql = iq.sqlQuery()
	case dialect.Gremlin:
		selector.gremlin = iq.gremlinQuery()
	}
	return selector
}

func (iq *ItemQuery) sqlAll(ctx context.Context) ([]*Item, error) {
	rows := &sql.Rows{}
	selector := iq.sqlQuery()
	if unique := iq.unique; len(unique) == 0 {
		selector.Distinct()
	}
	query, args := selector.Query()
	if err := iq.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var is Items
	if err := is.FromRows(rows); err != nil {
		return nil, err
	}
	is.config(iq.config)
	return is, nil
}

func (iq *ItemQuery) sqlCount(ctx context.Context) (int, error) {
	rows := &sql.Rows{}
	selector := iq.sqlQuery()
	unique := []string{item.FieldID}
	if len(iq.unique) > 0 {
		unique = iq.unique
	}
	selector.Count(sql.Distinct(selector.Columns(unique...)...))
	query, args := selector.Query()
	if err := iq.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, errors.New("ent: no rows found")
	}
	var n int
	if err := rows.Scan(&n); err != nil {
		return 0, fmt.Errorf("ent: failed reading count: %v", err)
	}
	return n, nil
}

func (iq *ItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (iq *ItemQuery) sqlQuery() *sql.Selector {
	t1 := sql.Table(item.Table)
	selector := sql.Select(t1.Columns(item.Columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(item.Columns...)...)
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

func (iq *ItemQuery) gremlinAll(ctx context.Context) ([]*Item, error) {
	res := &gremlin.Response{}
	query, bindings := iq.gremlinQuery().ValueMap(true).Query()
	if err := iq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var is Items
	if err := is.FromResponse(res); err != nil {
		return nil, err
	}
	is.config(iq.config)
	return is, nil
}

func (iq *ItemQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := iq.gremlinQuery().Count().Query()
	if err := iq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (iq *ItemQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := iq.gremlinQuery().HasNext().Query()
	if err := iq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (iq *ItemQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(item.Label)
	if iq.gremlin != nil {
		v = iq.gremlin.Clone()
	}
	for _, p := range iq.predicates {
		p(v)
	}
	if len(iq.order) > 0 {
		v.Order()
		for _, p := range iq.order {
			p(v)
		}
	}
	switch limit, offset := iq.limit, iq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := iq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// ItemGroupBy is the builder for group-by Item entities.
type ItemGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *ItemGroupBy) Aggregate(fns ...Aggregate) *ItemGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scan the result into the given value.
func (igb *ItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	switch igb.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return igb.sqlScan(ctx, v)
	case dialect.Gremlin:
		return igb.gremlinScan(ctx, v)
	default:
		return errors.New("igb: unsupported dialect")
	}
}

// ScanX is like Scan, but panics if an error occurs.
func (igb *ItemGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := igb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (igb *ItemGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: ItemGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (igb *ItemGroupBy) StringsX(ctx context.Context) []string {
	v, err := igb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (igb *ItemGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: ItemGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (igb *ItemGroupBy) IntsX(ctx context.Context) []int {
	v, err := igb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (igb *ItemGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: ItemGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (igb *ItemGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := igb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (igb *ItemGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: ItemGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (igb *ItemGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := igb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (igb *ItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := igb.sqlQuery().Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *ItemGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql
	columns := make([]string, 0, len(igb.fields)+len(igb.fns))
	columns = append(columns, igb.fields...)
	for _, fn := range igb.fns {
		columns = append(columns, fn.SQL(selector))
	}
	return selector.Select(columns...).GroupBy(igb.fields...)
}

func (igb *ItemGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := igb.gremlinQuery().Query()
	if err := igb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(igb.fields)+len(igb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (igb *ItemGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range igb.fns {
		name, tr := fn.Gremlin("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range igb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return igb.gremlin.Group().
		By(__.Values(igb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}

// ItemSelect is the builder for select fields of Item entities.
type ItemSelect struct {
	config
	fields []string
	// intermediate queries.
	sql     *sql.Selector
	gremlin *dsl.Traversal
}

// Scan applies the selector query and scan the result into the given value.
func (is *ItemSelect) Scan(ctx context.Context, v interface{}) error {
	switch is.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return is.sqlScan(ctx, v)
	case dialect.Gremlin:
		return is.gremlinScan(ctx, v)
	default:
		return errors.New("ItemSelect: unsupported dialect")
	}
}

// ScanX is like Scan, but panics if an error occurs.
func (is *ItemSelect) ScanX(ctx context.Context, v interface{}) {
	if err := is.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (is *ItemSelect) Strings(ctx context.Context) ([]string, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: ItemSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (is *ItemSelect) StringsX(ctx context.Context) []string {
	v, err := is.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (is *ItemSelect) Ints(ctx context.Context) ([]int, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: ItemSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (is *ItemSelect) IntsX(ctx context.Context) []int {
	v, err := is.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (is *ItemSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: ItemSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (is *ItemSelect) Float64sX(ctx context.Context) []float64 {
	v, err := is.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (is *ItemSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: ItemSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (is *ItemSelect) BoolsX(ctx context.Context) []bool {
	v, err := is.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (is *ItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := is.sqlQuery().Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (is *ItemSelect) sqlQuery() sql.Querier {
	view := "item_view"
	return sql.Select(is.fields...).From(is.sql.As(view))
}

func (is *ItemSelect) gremlinScan(ctx context.Context, v interface{}) error {
	var (
		traversal *dsl.Traversal
		res       = &gremlin.Response{}
	)
	if len(is.fields) == 1 {
		if is.fields[0] != item.FieldID {
			traversal = is.gremlin.Values(is.fields...)
		} else {
			traversal = is.gremlin.ID()
		}
	} else {
		fields := make([]interface{}, len(is.fields))
		for i, f := range is.fields {
			fields[i] = f
		}
		traversal = is.gremlin.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := is.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(is.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
