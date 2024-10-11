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
	"github.com/buildbarn/bb-portal/ent/gen/ent/bazelinvocation"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testcollection"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testresultbes"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testsummary"
)

// TestCollectionQuery is the builder for querying TestCollection entities.
type TestCollectionQuery struct {
	config
	ctx                  *QueryContext
	order                []testcollection.OrderOption
	inters               []Interceptor
	predicates           []predicate.TestCollection
	withBazelInvocation  *BazelInvocationQuery
	withTestSummary      *TestSummaryQuery
	withTestResults      *TestResultBESQuery
	withFKs              bool
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*TestCollection) error
	withNamedTestResults map[string]*TestResultBESQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestCollectionQuery builder.
func (tcq *TestCollectionQuery) Where(ps ...predicate.TestCollection) *TestCollectionQuery {
	tcq.predicates = append(tcq.predicates, ps...)
	return tcq
}

// Limit the number of records to be returned by this query.
func (tcq *TestCollectionQuery) Limit(limit int) *TestCollectionQuery {
	tcq.ctx.Limit = &limit
	return tcq
}

// Offset to start from.
func (tcq *TestCollectionQuery) Offset(offset int) *TestCollectionQuery {
	tcq.ctx.Offset = &offset
	return tcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tcq *TestCollectionQuery) Unique(unique bool) *TestCollectionQuery {
	tcq.ctx.Unique = &unique
	return tcq
}

// Order specifies how the records should be ordered.
func (tcq *TestCollectionQuery) Order(o ...testcollection.OrderOption) *TestCollectionQuery {
	tcq.order = append(tcq.order, o...)
	return tcq
}

// QueryBazelInvocation chains the current query on the "bazel_invocation" edge.
func (tcq *TestCollectionQuery) QueryBazelInvocation() *BazelInvocationQuery {
	query := (&BazelInvocationClient{config: tcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testcollection.Table, testcollection.FieldID, selector),
			sqlgraph.To(bazelinvocation.Table, bazelinvocation.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, testcollection.BazelInvocationTable, testcollection.BazelInvocationColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTestSummary chains the current query on the "test_summary" edge.
func (tcq *TestCollectionQuery) QueryTestSummary() *TestSummaryQuery {
	query := (&TestSummaryClient{config: tcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testcollection.Table, testcollection.FieldID, selector),
			sqlgraph.To(testsummary.Table, testsummary.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, testcollection.TestSummaryTable, testcollection.TestSummaryColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTestResults chains the current query on the "test_results" edge.
func (tcq *TestCollectionQuery) QueryTestResults() *TestResultBESQuery {
	query := (&TestResultBESClient{config: tcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testcollection.Table, testcollection.FieldID, selector),
			sqlgraph.To(testresultbes.Table, testresultbes.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, testcollection.TestResultsTable, testcollection.TestResultsColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TestCollection entity from the query.
// Returns a *NotFoundError when no TestCollection was found.
func (tcq *TestCollectionQuery) First(ctx context.Context) (*TestCollection, error) {
	nodes, err := tcq.Limit(1).All(setContextOp(ctx, tcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testcollection.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tcq *TestCollectionQuery) FirstX(ctx context.Context) *TestCollection {
	node, err := tcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TestCollection ID from the query.
// Returns a *NotFoundError when no TestCollection ID was found.
func (tcq *TestCollectionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(1).IDs(setContextOp(ctx, tcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testcollection.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tcq *TestCollectionQuery) FirstIDX(ctx context.Context) int {
	id, err := tcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TestCollection entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TestCollection entity is found.
// Returns a *NotFoundError when no TestCollection entities are found.
func (tcq *TestCollectionQuery) Only(ctx context.Context) (*TestCollection, error) {
	nodes, err := tcq.Limit(2).All(setContextOp(ctx, tcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testcollection.Label}
	default:
		return nil, &NotSingularError{testcollection.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tcq *TestCollectionQuery) OnlyX(ctx context.Context) *TestCollection {
	node, err := tcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TestCollection ID in the query.
// Returns a *NotSingularError when more than one TestCollection ID is found.
// Returns a *NotFoundError when no entities are found.
func (tcq *TestCollectionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(2).IDs(setContextOp(ctx, tcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testcollection.Label}
	default:
		err = &NotSingularError{testcollection.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tcq *TestCollectionQuery) OnlyIDX(ctx context.Context) int {
	id, err := tcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TestCollections.
func (tcq *TestCollectionQuery) All(ctx context.Context) ([]*TestCollection, error) {
	ctx = setContextOp(ctx, tcq.ctx, "All")
	if err := tcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TestCollection, *TestCollectionQuery]()
	return withInterceptors[[]*TestCollection](ctx, tcq, qr, tcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tcq *TestCollectionQuery) AllX(ctx context.Context) []*TestCollection {
	nodes, err := tcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TestCollection IDs.
func (tcq *TestCollectionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tcq.ctx.Unique == nil && tcq.path != nil {
		tcq.Unique(true)
	}
	ctx = setContextOp(ctx, tcq.ctx, "IDs")
	if err = tcq.Select(testcollection.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tcq *TestCollectionQuery) IDsX(ctx context.Context) []int {
	ids, err := tcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tcq *TestCollectionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tcq.ctx, "Count")
	if err := tcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tcq, querierCount[*TestCollectionQuery](), tcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tcq *TestCollectionQuery) CountX(ctx context.Context) int {
	count, err := tcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tcq *TestCollectionQuery) Exist(ctx context.Context) (bool, error) {
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
func (tcq *TestCollectionQuery) ExistX(ctx context.Context) bool {
	exist, err := tcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestCollectionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tcq *TestCollectionQuery) Clone() *TestCollectionQuery {
	if tcq == nil {
		return nil
	}
	return &TestCollectionQuery{
		config:              tcq.config,
		ctx:                 tcq.ctx.Clone(),
		order:               append([]testcollection.OrderOption{}, tcq.order...),
		inters:              append([]Interceptor{}, tcq.inters...),
		predicates:          append([]predicate.TestCollection{}, tcq.predicates...),
		withBazelInvocation: tcq.withBazelInvocation.Clone(),
		withTestSummary:     tcq.withTestSummary.Clone(),
		withTestResults:     tcq.withTestResults.Clone(),
		// clone intermediate query.
		sql:  tcq.sql.Clone(),
		path: tcq.path,
	}
}

// WithBazelInvocation tells the query-builder to eager-load the nodes that are connected to
// the "bazel_invocation" edge. The optional arguments are used to configure the query builder of the edge.
func (tcq *TestCollectionQuery) WithBazelInvocation(opts ...func(*BazelInvocationQuery)) *TestCollectionQuery {
	query := (&BazelInvocationClient{config: tcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tcq.withBazelInvocation = query
	return tcq
}

// WithTestSummary tells the query-builder to eager-load the nodes that are connected to
// the "test_summary" edge. The optional arguments are used to configure the query builder of the edge.
func (tcq *TestCollectionQuery) WithTestSummary(opts ...func(*TestSummaryQuery)) *TestCollectionQuery {
	query := (&TestSummaryClient{config: tcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tcq.withTestSummary = query
	return tcq
}

// WithTestResults tells the query-builder to eager-load the nodes that are connected to
// the "test_results" edge. The optional arguments are used to configure the query builder of the edge.
func (tcq *TestCollectionQuery) WithTestResults(opts ...func(*TestResultBESQuery)) *TestCollectionQuery {
	query := (&TestResultBESClient{config: tcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tcq.withTestResults = query
	return tcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Label string `json:"label,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TestCollection.Query().
//		GroupBy(testcollection.FieldLabel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tcq *TestCollectionQuery) GroupBy(field string, fields ...string) *TestCollectionGroupBy {
	tcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TestCollectionGroupBy{build: tcq}
	grbuild.flds = &tcq.ctx.Fields
	grbuild.label = testcollection.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Label string `json:"label,omitempty"`
//	}
//
//	client.TestCollection.Query().
//		Select(testcollection.FieldLabel).
//		Scan(ctx, &v)
func (tcq *TestCollectionQuery) Select(fields ...string) *TestCollectionSelect {
	tcq.ctx.Fields = append(tcq.ctx.Fields, fields...)
	sbuild := &TestCollectionSelect{TestCollectionQuery: tcq}
	sbuild.label = testcollection.Label
	sbuild.flds, sbuild.scan = &tcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TestCollectionSelect configured with the given aggregations.
func (tcq *TestCollectionQuery) Aggregate(fns ...AggregateFunc) *TestCollectionSelect {
	return tcq.Select().Aggregate(fns...)
}

func (tcq *TestCollectionQuery) prepareQuery(ctx context.Context) error {
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
		if !testcollection.ValidColumn(f) {
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

func (tcq *TestCollectionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TestCollection, error) {
	var (
		nodes       = []*TestCollection{}
		withFKs     = tcq.withFKs
		_spec       = tcq.querySpec()
		loadedTypes = [3]bool{
			tcq.withBazelInvocation != nil,
			tcq.withTestSummary != nil,
			tcq.withTestResults != nil,
		}
	)
	if tcq.withBazelInvocation != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, testcollection.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TestCollection).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TestCollection{config: tcq.config}
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
	if query := tcq.withBazelInvocation; query != nil {
		if err := tcq.loadBazelInvocation(ctx, query, nodes, nil,
			func(n *TestCollection, e *BazelInvocation) { n.Edges.BazelInvocation = e }); err != nil {
			return nil, err
		}
	}
	if query := tcq.withTestSummary; query != nil {
		if err := tcq.loadTestSummary(ctx, query, nodes, nil,
			func(n *TestCollection, e *TestSummary) { n.Edges.TestSummary = e }); err != nil {
			return nil, err
		}
	}
	if query := tcq.withTestResults; query != nil {
		if err := tcq.loadTestResults(ctx, query, nodes,
			func(n *TestCollection) { n.Edges.TestResults = []*TestResultBES{} },
			func(n *TestCollection, e *TestResultBES) { n.Edges.TestResults = append(n.Edges.TestResults, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range tcq.withNamedTestResults {
		if err := tcq.loadTestResults(ctx, query, nodes,
			func(n *TestCollection) { n.appendNamedTestResults(name) },
			func(n *TestCollection, e *TestResultBES) { n.appendNamedTestResults(name, e) }); err != nil {
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

func (tcq *TestCollectionQuery) loadBazelInvocation(ctx context.Context, query *BazelInvocationQuery, nodes []*TestCollection, init func(*TestCollection), assign func(*TestCollection, *BazelInvocation)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TestCollection)
	for i := range nodes {
		if nodes[i].bazel_invocation_test_collection == nil {
			continue
		}
		fk := *nodes[i].bazel_invocation_test_collection
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(bazelinvocation.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "bazel_invocation_test_collection" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tcq *TestCollectionQuery) loadTestSummary(ctx context.Context, query *TestSummaryQuery, nodes []*TestCollection, init func(*TestCollection), assign func(*TestCollection, *TestSummary)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*TestCollection)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.TestSummary(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(testcollection.TestSummaryColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.test_collection_test_summary
		if fk == nil {
			return fmt.Errorf(`foreign-key "test_collection_test_summary" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "test_collection_test_summary" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (tcq *TestCollectionQuery) loadTestResults(ctx context.Context, query *TestResultBESQuery, nodes []*TestCollection, init func(*TestCollection), assign func(*TestCollection, *TestResultBES)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*TestCollection)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TestResultBES(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(testcollection.TestResultsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.test_collection_test_results
		if fk == nil {
			return fmt.Errorf(`foreign-key "test_collection_test_results" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "test_collection_test_results" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tcq *TestCollectionQuery) sqlCount(ctx context.Context) (int, error) {
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

func (tcq *TestCollectionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(testcollection.Table, testcollection.Columns, sqlgraph.NewFieldSpec(testcollection.FieldID, field.TypeInt))
	_spec.From = tcq.sql
	if unique := tcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tcq.path != nil {
		_spec.Unique = true
	}
	if fields := tcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testcollection.FieldID)
		for i := range fields {
			if fields[i] != testcollection.FieldID {
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

func (tcq *TestCollectionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tcq.driver.Dialect())
	t1 := builder.Table(testcollection.Table)
	columns := tcq.ctx.Fields
	if len(columns) == 0 {
		columns = testcollection.Columns
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

// WithNamedTestResults tells the query-builder to eager-load the nodes that are connected to the "test_results"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (tcq *TestCollectionQuery) WithNamedTestResults(name string, opts ...func(*TestResultBESQuery)) *TestCollectionQuery {
	query := (&TestResultBESClient{config: tcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if tcq.withNamedTestResults == nil {
		tcq.withNamedTestResults = make(map[string]*TestResultBESQuery)
	}
	tcq.withNamedTestResults[name] = query
	return tcq
}

// TestCollectionGroupBy is the group-by builder for TestCollection entities.
type TestCollectionGroupBy struct {
	selector
	build *TestCollectionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tcgb *TestCollectionGroupBy) Aggregate(fns ...AggregateFunc) *TestCollectionGroupBy {
	tcgb.fns = append(tcgb.fns, fns...)
	return tcgb
}

// Scan applies the selector query and scans the result into the given value.
func (tcgb *TestCollectionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcgb.build.ctx, "GroupBy")
	if err := tcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestCollectionQuery, *TestCollectionGroupBy](ctx, tcgb.build, tcgb, tcgb.build.inters, v)
}

func (tcgb *TestCollectionGroupBy) sqlScan(ctx context.Context, root *TestCollectionQuery, v any) error {
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

// TestCollectionSelect is the builder for selecting fields of TestCollection entities.
type TestCollectionSelect struct {
	*TestCollectionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tcs *TestCollectionSelect) Aggregate(fns ...AggregateFunc) *TestCollectionSelect {
	tcs.fns = append(tcs.fns, fns...)
	return tcs
}

// Scan applies the selector query and scans the result into the given value.
func (tcs *TestCollectionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcs.ctx, "Select")
	if err := tcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestCollectionQuery, *TestCollectionSelect](ctx, tcs.TestCollectionQuery, tcs, tcs.inters, v)
}

func (tcs *TestCollectionSelect) sqlScan(ctx context.Context, root *TestCollectionQuery, v any) error {
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
