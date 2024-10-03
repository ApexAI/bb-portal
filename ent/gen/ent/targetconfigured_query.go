// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetconfigured"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetpair"
)

// TargetConfiguredQuery is the builder for querying TargetConfigured entities.
type TargetConfiguredQuery struct {
	config
	ctx                 *QueryContext
	order               []targetconfigured.OrderOption
	inters              []Interceptor
	predicates          []predicate.TargetConfigured
	withTargetPair      *TargetPairQuery
	modifiers           []func(*sql.Selector)
	loadTotal           []func(context.Context, []*TargetConfigured) error
	withNamedTargetPair map[string]*TargetPairQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TargetConfiguredQuery builder.
func (tcq *TargetConfiguredQuery) Where(ps ...predicate.TargetConfigured) *TargetConfiguredQuery {
	tcq.predicates = append(tcq.predicates, ps...)
	return tcq
}

// Limit the number of records to be returned by this query.
func (tcq *TargetConfiguredQuery) Limit(limit int) *TargetConfiguredQuery {
	tcq.ctx.Limit = &limit
	return tcq
}

// Offset to start from.
func (tcq *TargetConfiguredQuery) Offset(offset int) *TargetConfiguredQuery {
	tcq.ctx.Offset = &offset
	return tcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tcq *TargetConfiguredQuery) Unique(unique bool) *TargetConfiguredQuery {
	tcq.ctx.Unique = &unique
	return tcq
}

// Order specifies how the records should be ordered.
func (tcq *TargetConfiguredQuery) Order(o ...targetconfigured.OrderOption) *TargetConfiguredQuery {
	tcq.order = append(tcq.order, o...)
	return tcq
}

// QueryTargetPair chains the current query on the "target_pair" edge.
func (tcq *TargetConfiguredQuery) QueryTargetPair() *TargetPairQuery {
	query := (&TargetPairClient{config: tcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(targetconfigured.Table, targetconfigured.FieldID, selector),
			sqlgraph.To(targetpair.Table, targetpair.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, targetconfigured.TargetPairTable, targetconfigured.TargetPairColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TargetConfigured entity from the query.
// Returns a *NotFoundError when no TargetConfigured was found.
func (tcq *TargetConfiguredQuery) First(ctx context.Context) (*TargetConfigured, error) {
	nodes, err := tcq.Limit(1).All(setContextOp(ctx, tcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{targetconfigured.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tcq *TargetConfiguredQuery) FirstX(ctx context.Context) *TargetConfigured {
	node, err := tcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TargetConfigured ID from the query.
// Returns a *NotFoundError when no TargetConfigured ID was found.
func (tcq *TargetConfiguredQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(1).IDs(setContextOp(ctx, tcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{targetconfigured.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tcq *TargetConfiguredQuery) FirstIDX(ctx context.Context) int {
	id, err := tcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TargetConfigured entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TargetConfigured entity is found.
// Returns a *NotFoundError when no TargetConfigured entities are found.
func (tcq *TargetConfiguredQuery) Only(ctx context.Context) (*TargetConfigured, error) {
	nodes, err := tcq.Limit(2).All(setContextOp(ctx, tcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{targetconfigured.Label}
	default:
		return nil, &NotSingularError{targetconfigured.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tcq *TargetConfiguredQuery) OnlyX(ctx context.Context) *TargetConfigured {
	node, err := tcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TargetConfigured ID in the query.
// Returns a *NotSingularError when more than one TargetConfigured ID is found.
// Returns a *NotFoundError when no entities are found.
func (tcq *TargetConfiguredQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(2).IDs(setContextOp(ctx, tcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{targetconfigured.Label}
	default:
		err = &NotSingularError{targetconfigured.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tcq *TargetConfiguredQuery) OnlyIDX(ctx context.Context) int {
	id, err := tcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TargetConfigureds.
func (tcq *TargetConfiguredQuery) All(ctx context.Context) ([]*TargetConfigured, error) {
	ctx = setContextOp(ctx, tcq.ctx, "All")
	if err := tcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TargetConfigured, *TargetConfiguredQuery]()
	return withInterceptors[[]*TargetConfigured](ctx, tcq, qr, tcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tcq *TargetConfiguredQuery) AllX(ctx context.Context) []*TargetConfigured {
	nodes, err := tcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TargetConfigured IDs.
func (tcq *TargetConfiguredQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tcq.ctx.Unique == nil && tcq.path != nil {
		tcq.Unique(true)
	}
	ctx = setContextOp(ctx, tcq.ctx, "IDs")
	if err = tcq.Select(targetconfigured.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tcq *TargetConfiguredQuery) IDsX(ctx context.Context) []int {
	ids, err := tcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tcq *TargetConfiguredQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tcq.ctx, "Count")
	if err := tcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tcq, querierCount[*TargetConfiguredQuery](), tcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tcq *TargetConfiguredQuery) CountX(ctx context.Context) int {
	count, err := tcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tcq *TargetConfiguredQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tcq.ctx, "Exist")
	switch _, err := tcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tcq *TargetConfiguredQuery) ExistX(ctx context.Context) bool {
	exist, err := tcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TargetConfiguredQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tcq *TargetConfiguredQuery) Clone() *TargetConfiguredQuery {
	if tcq == nil {
		return nil
	}
	return &TargetConfiguredQuery{
		config:         tcq.config,
		ctx:            tcq.ctx.Clone(),
		order:          append([]targetconfigured.OrderOption{}, tcq.order...),
		inters:         append([]Interceptor{}, tcq.inters...),
		predicates:     append([]predicate.TargetConfigured{}, tcq.predicates...),
		withTargetPair: tcq.withTargetPair.Clone(),
		// clone intermediate query.
		sql:  tcq.sql.Clone(),
		path: tcq.path,
	}
}

// WithTargetPair tells the query-builder to eager-load the nodes that are connected to
// the "target_pair" edge. The optional arguments are used to configure the query builder of the edge.
func (tcq *TargetConfiguredQuery) WithTargetPair(opts ...func(*TargetPairQuery)) *TargetConfiguredQuery {
	query := (&TargetPairClient{config: tcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tcq.withTargetPair = query
	return tcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Tag []string `json:"tag,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TargetConfigured.Query().
//		GroupBy(targetconfigured.FieldTag).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tcq *TargetConfiguredQuery) GroupBy(field string, fields ...string) *TargetConfiguredGroupBy {
	tcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TargetConfiguredGroupBy{build: tcq}
	grbuild.flds = &tcq.ctx.Fields
	grbuild.label = targetconfigured.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Tag []string `json:"tag,omitempty"`
//	}
//
//	client.TargetConfigured.Query().
//		Select(targetconfigured.FieldTag).
//		Scan(ctx, &v)
func (tcq *TargetConfiguredQuery) Select(fields ...string) *TargetConfiguredSelect {
	tcq.ctx.Fields = append(tcq.ctx.Fields, fields...)
	sbuild := &TargetConfiguredSelect{TargetConfiguredQuery: tcq}
	sbuild.label = targetconfigured.Label
	sbuild.flds, sbuild.scan = &tcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TargetConfiguredSelect configured with the given aggregations.
func (tcq *TargetConfiguredQuery) Aggregate(fns ...AggregateFunc) *TargetConfiguredSelect {
	return tcq.Select().Aggregate(fns...)
}

func (tcq *TargetConfiguredQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tcq); err != nil {
				return err
			}
		}
	}
	for _, f := range tcq.ctx.Fields {
		if !targetconfigured.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tcq.path != nil {
		prev, err := tcq.path(ctx)
		if err != nil {
			return err
		}
		tcq.sql = prev
	}
	return nil
}

func (tcq *TargetConfiguredQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TargetConfigured, error) {
	var (
		nodes       = []*TargetConfigured{}
		_spec       = tcq.querySpec()
		loadedTypes = [1]bool{
			tcq.withTargetPair != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TargetConfigured).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TargetConfigured{config: tcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tcq.modifiers) > 0 {
		_spec.Modifiers = tcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tcq.withTargetPair; query != nil {
		if err := tcq.loadTargetPair(ctx, query, nodes,
			func(n *TargetConfigured) { n.Edges.TargetPair = []*TargetPair{} },
			func(n *TargetConfigured, e *TargetPair) { n.Edges.TargetPair = append(n.Edges.TargetPair, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tcq.withNamedTargetPair {
		if err := tcq.loadTargetPair(ctx, query, nodes,
			func(n *TargetConfigured) { n.appendNamedTargetPair(name) },
			func(n *TargetConfigured, e *TargetPair) { n.appendNamedTargetPair(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range tcq.loadTotal {
		if err := tcq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tcq *TargetConfiguredQuery) loadTargetPair(ctx context.Context, query *TargetPairQuery, nodes []*TargetConfigured, init func(*TargetConfigured), assign func(*TargetConfigured, *TargetPair)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*TargetConfigured)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TargetPair(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(targetconfigured.TargetPairColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.target_pair_configuration
		if fk == nil {
			return fmt.Errorf(`foreign-key "target_pair_configuration" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "target_pair_configuration" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tcq *TargetConfiguredQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tcq.querySpec()
	if len(tcq.modifiers) > 0 {
		_spec.Modifiers = tcq.modifiers
	}
	_spec.Node.Columns = tcq.ctx.Fields
	if len(tcq.ctx.Fields) > 0 {
		_spec.Unique = tcq.ctx.Unique != nil && *tcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tcq.driver, _spec)
}

func (tcq *TargetConfiguredQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(targetconfigured.Table, targetconfigured.Columns, sqlgraph.NewFieldSpec(targetconfigured.FieldID, field.TypeInt))
	_spec.From = tcq.sql
	if unique := tcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tcq.path != nil {
		_spec.Unique = true
	}
	if fields := tcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, targetconfigured.FieldID)
		for i := range fields {
			if fields[i] != targetconfigured.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tcq *TargetConfiguredQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tcq.driver.Dialect())
	t1 := builder.Table(targetconfigured.Table)
	columns := tcq.ctx.Fields
	if len(columns) == 0 {
		columns = targetconfigured.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tcq.sql != nil {
		selector = tcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tcq.ctx.Unique != nil && *tcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tcq.predicates {
		p(selector)
	}
	for _, p := range tcq.order {
		p(selector)
	}
	if offset := tcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedTargetPair tells the query-builder to eager-load the nodes that are connected to the "target_pair"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tcq *TargetConfiguredQuery) WithNamedTargetPair(name string, opts ...func(*TargetPairQuery)) *TargetConfiguredQuery {
	query := (&TargetPairClient{config: tcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if tcq.withNamedTargetPair == nil {
		tcq.withNamedTargetPair = make(map[string]*TargetPairQuery)
	}
	tcq.withNamedTargetPair[name] = query
	return tcq
}

// TargetConfiguredGroupBy is the group-by builder for TargetConfigured entities.
type TargetConfiguredGroupBy struct {
	selector
	build *TargetConfiguredQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tcgb *TargetConfiguredGroupBy) Aggregate(fns ...AggregateFunc) *TargetConfiguredGroupBy {
	tcgb.fns = append(tcgb.fns, fns...)
	return tcgb
}

// Scan applies the selector query and scans the result into the given value.
func (tcgb *TargetConfiguredGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcgb.build.ctx, "GroupBy")
	if err := tcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TargetConfiguredQuery, *TargetConfiguredGroupBy](ctx, tcgb.build, tcgb, tcgb.build.inters, v)
}

func (tcgb *TargetConfiguredGroupBy) sqlScan(ctx context.Context, root *TargetConfiguredQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tcgb.fns))
	for _, fn := range tcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tcgb.flds)+len(tcgb.fns))
		for _, f := range *tcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TargetConfiguredSelect is the builder for selecting fields of TargetConfigured entities.
type TargetConfiguredSelect struct {
	*TargetConfiguredQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tcs *TargetConfiguredSelect) Aggregate(fns ...AggregateFunc) *TargetConfiguredSelect {
	tcs.fns = append(tcs.fns, fns...)
	return tcs
}

// Scan applies the selector query and scans the result into the given value.
func (tcs *TargetConfiguredSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcs.ctx, "Select")
	if err := tcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TargetConfiguredQuery, *TargetConfiguredSelect](ctx, tcs.TargetConfiguredQuery, tcs, tcs.inters, v)
}

func (tcs *TargetConfiguredSelect) sqlScan(ctx context.Context, root *TargetConfiguredQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tcs.fns))
	for _, fn := range tcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}