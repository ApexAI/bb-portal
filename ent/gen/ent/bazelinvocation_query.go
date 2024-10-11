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
	"github.com/buildbarn/bb-portal/ent/gen/ent/bazelinvocationproblem"
	"github.com/buildbarn/bb-portal/ent/gen/ent/build"
	"github.com/buildbarn/bb-portal/ent/gen/ent/eventfile"
	"github.com/buildbarn/bb-portal/ent/gen/ent/metrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetpair"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testcollection"
)

// BazelInvocationQuery is the builder for querying BazelInvocation entities.
type BazelInvocationQuery struct {
	config
	ctx                     *QueryContext
	order                   []bazelinvocation.OrderOption
	inters                  []Interceptor
	predicates              []predicate.BazelInvocation
	withEventFile           *EventFileQuery
	withBuild               *BuildQuery
	withProblems            *BazelInvocationProblemQuery
	withMetrics             *MetricsQuery
	withTestCollection      *TestCollectionQuery
	withTargets             *TargetPairQuery
	withFKs                 bool
	modifiers               []func(*sql.Selector)
	loadTotal               []func(context.Context, []*BazelInvocation) error
	withNamedProblems       map[string]*BazelInvocationProblemQuery
	withNamedTestCollection map[string]*TestCollectionQuery
	withNamedTargets        map[string]*TargetPairQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BazelInvocationQuery builder.
func (biq *BazelInvocationQuery) Where(ps ...predicate.BazelInvocation) *BazelInvocationQuery {
	biq.predicates = append(biq.predicates, ps...)
	return biq
}

// Limit the number of records to be returned by this query.
func (biq *BazelInvocationQuery) Limit(limit int) *BazelInvocationQuery {
	biq.ctx.Limit = &limit
	return biq
}

// Offset to start from.
func (biq *BazelInvocationQuery) Offset(offset int) *BazelInvocationQuery {
	biq.ctx.Offset = &offset
	return biq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (biq *BazelInvocationQuery) Unique(unique bool) *BazelInvocationQuery {
	biq.ctx.Unique = &unique
	return biq
}

// Order specifies how the records should be ordered.
func (biq *BazelInvocationQuery) Order(o ...bazelinvocation.OrderOption) *BazelInvocationQuery {
	biq.order = append(biq.order, o...)
	return biq
}

// QueryEventFile chains the current query on the "event_file" edge.
func (biq *BazelInvocationQuery) QueryEventFile() *EventFileQuery {
	query := (&EventFileClient{config: biq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := biq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := biq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bazelinvocation.Table, bazelinvocation.FieldID, selector),
			sqlgraph.To(eventfile.Table, eventfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, bazelinvocation.EventFileTable, bazelinvocation.EventFileColumn),
		)
		fromU = sqlgraph.SetNeighbors(biq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBuild chains the current query on the "build" edge.
func (biq *BazelInvocationQuery) QueryBuild() *BuildQuery {
	query := (&BuildClient{config: biq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := biq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := biq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bazelinvocation.Table, bazelinvocation.FieldID, selector),
			sqlgraph.To(build.Table, build.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bazelinvocation.BuildTable, bazelinvocation.BuildColumn),
		)
		fromU = sqlgraph.SetNeighbors(biq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProblems chains the current query on the "problems" edge.
func (biq *BazelInvocationQuery) QueryProblems() *BazelInvocationProblemQuery {
	query := (&BazelInvocationProblemClient{config: biq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := biq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := biq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bazelinvocation.Table, bazelinvocation.FieldID, selector),
			sqlgraph.To(bazelinvocationproblem.Table, bazelinvocationproblem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bazelinvocation.ProblemsTable, bazelinvocation.ProblemsColumn),
		)
		fromU = sqlgraph.SetNeighbors(biq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMetrics chains the current query on the "metrics" edge.
func (biq *BazelInvocationQuery) QueryMetrics() *MetricsQuery {
	query := (&MetricsClient{config: biq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := biq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := biq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bazelinvocation.Table, bazelinvocation.FieldID, selector),
			sqlgraph.To(metrics.Table, metrics.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, bazelinvocation.MetricsTable, bazelinvocation.MetricsColumn),
		)
		fromU = sqlgraph.SetNeighbors(biq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTestCollection chains the current query on the "test_collection" edge.
func (biq *BazelInvocationQuery) QueryTestCollection() *TestCollectionQuery {
	query := (&TestCollectionClient{config: biq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := biq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := biq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bazelinvocation.Table, bazelinvocation.FieldID, selector),
			sqlgraph.To(testcollection.Table, testcollection.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bazelinvocation.TestCollectionTable, bazelinvocation.TestCollectionColumn),
		)
		fromU = sqlgraph.SetNeighbors(biq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTargets chains the current query on the "targets" edge.
func (biq *BazelInvocationQuery) QueryTargets() *TargetPairQuery {
	query := (&TargetPairClient{config: biq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := biq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := biq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bazelinvocation.Table, bazelinvocation.FieldID, selector),
			sqlgraph.To(targetpair.Table, targetpair.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bazelinvocation.TargetsTable, bazelinvocation.TargetsColumn),
		)
		fromU = sqlgraph.SetNeighbors(biq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BazelInvocation entity from the query.
// Returns a *NotFoundError when no BazelInvocation was found.
func (biq *BazelInvocationQuery) First(ctx context.Context) (*BazelInvocation, error) {
	nodes, err := biq.Limit(1).All(setContextOp(ctx, biq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bazelinvocation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (biq *BazelInvocationQuery) FirstX(ctx context.Context) *BazelInvocation {
	node, err := biq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BazelInvocation ID from the query.
// Returns a *NotFoundError when no BazelInvocation ID was found.
func (biq *BazelInvocationQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = biq.Limit(1).IDs(setContextOp(ctx, biq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bazelinvocation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (biq *BazelInvocationQuery) FirstIDX(ctx context.Context) int {
	id, err := biq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BazelInvocation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BazelInvocation entity is found.
// Returns a *NotFoundError when no BazelInvocation entities are found.
func (biq *BazelInvocationQuery) Only(ctx context.Context) (*BazelInvocation, error) {
	nodes, err := biq.Limit(2).All(setContextOp(ctx, biq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bazelinvocation.Label}
	default:
		return nil, &NotSingularError{bazelinvocation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (biq *BazelInvocationQuery) OnlyX(ctx context.Context) *BazelInvocation {
	node, err := biq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BazelInvocation ID in the query.
// Returns a *NotSingularError when more than one BazelInvocation ID is found.
// Returns a *NotFoundError when no entities are found.
func (biq *BazelInvocationQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = biq.Limit(2).IDs(setContextOp(ctx, biq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bazelinvocation.Label}
	default:
		err = &NotSingularError{bazelinvocation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (biq *BazelInvocationQuery) OnlyIDX(ctx context.Context) int {
	id, err := biq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BazelInvocations.
func (biq *BazelInvocationQuery) All(ctx context.Context) ([]*BazelInvocation, error) {
	ctx = setContextOp(ctx, biq.ctx, "All")
	if err := biq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BazelInvocation, *BazelInvocationQuery]()
	return withInterceptors[[]*BazelInvocation](ctx, biq, qr, biq.inters)
}

// AllX is like All, but panics if an error occurs.
func (biq *BazelInvocationQuery) AllX(ctx context.Context) []*BazelInvocation {
	nodes, err := biq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BazelInvocation IDs.
func (biq *BazelInvocationQuery) IDs(ctx context.Context) (ids []int, err error) {
	if biq.ctx.Unique == nil && biq.path != nil {
		biq.Unique(true)
	}
	ctx = setContextOp(ctx, biq.ctx, "IDs")
	if err = biq.Select(bazelinvocation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (biq *BazelInvocationQuery) IDsX(ctx context.Context) []int {
	ids, err := biq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (biq *BazelInvocationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, biq.ctx, "Count")
	if err := biq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, biq, querierCount[*BazelInvocationQuery](), biq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (biq *BazelInvocationQuery) CountX(ctx context.Context) int {
	count, err := biq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (biq *BazelInvocationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, biq.ctx, "Exist")
	switch _, err := biq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (biq *BazelInvocationQuery) ExistX(ctx context.Context) bool {
	exist, err := biq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BazelInvocationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (biq *BazelInvocationQuery) Clone() *BazelInvocationQuery {
	if biq == nil {
		return nil
	}
	return &BazelInvocationQuery{
		config:             biq.config,
		ctx:                biq.ctx.Clone(),
		order:              append([]bazelinvocation.OrderOption{}, biq.order...),
		inters:             append([]Interceptor{}, biq.inters...),
		predicates:         append([]predicate.BazelInvocation{}, biq.predicates...),
		withEventFile:      biq.withEventFile.Clone(),
		withBuild:          biq.withBuild.Clone(),
		withProblems:       biq.withProblems.Clone(),
		withMetrics:        biq.withMetrics.Clone(),
		withTestCollection: biq.withTestCollection.Clone(),
		withTargets:        biq.withTargets.Clone(),
		// clone intermediate query.
		sql:  biq.sql.Clone(),
		path: biq.path,
	}
}

// WithEventFile tells the query-builder to eager-load the nodes that are connected to
// the "event_file" edge. The optional arguments are used to configure the query builder of the edge.
func (biq *BazelInvocationQuery) WithEventFile(opts ...func(*EventFileQuery)) *BazelInvocationQuery {
	query := (&EventFileClient{config: biq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	biq.withEventFile = query
	return biq
}

// WithBuild tells the query-builder to eager-load the nodes that are connected to
// the "build" edge. The optional arguments are used to configure the query builder of the edge.
func (biq *BazelInvocationQuery) WithBuild(opts ...func(*BuildQuery)) *BazelInvocationQuery {
	query := (&BuildClient{config: biq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	biq.withBuild = query
	return biq
}

// WithProblems tells the query-builder to eager-load the nodes that are connected to
// the "problems" edge. The optional arguments are used to configure the query builder of the edge.
func (biq *BazelInvocationQuery) WithProblems(opts ...func(*BazelInvocationProblemQuery)) *BazelInvocationQuery {
	query := (&BazelInvocationProblemClient{config: biq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	biq.withProblems = query
	return biq
}

// WithMetrics tells the query-builder to eager-load the nodes that are connected to
// the "metrics" edge. The optional arguments are used to configure the query builder of the edge.
func (biq *BazelInvocationQuery) WithMetrics(opts ...func(*MetricsQuery)) *BazelInvocationQuery {
	query := (&MetricsClient{config: biq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	biq.withMetrics = query
	return biq
}

// WithTestCollection tells the query-builder to eager-load the nodes that are connected to
// the "test_collection" edge. The optional arguments are used to configure the query builder of the edge.
func (biq *BazelInvocationQuery) WithTestCollection(opts ...func(*TestCollectionQuery)) *BazelInvocationQuery {
	query := (&TestCollectionClient{config: biq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	biq.withTestCollection = query
	return biq
}

// WithTargets tells the query-builder to eager-load the nodes that are connected to
// the "targets" edge. The optional arguments are used to configure the query builder of the edge.
func (biq *BazelInvocationQuery) WithTargets(opts ...func(*TargetPairQuery)) *BazelInvocationQuery {
	query := (&TargetPairClient{config: biq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	biq.withTargets = query
	return biq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		InvocationID uuid.UUID `json:"invocation_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BazelInvocation.Query().
//		GroupBy(bazelinvocation.FieldInvocationID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (biq *BazelInvocationQuery) GroupBy(field string, fields ...string) *BazelInvocationGroupBy {
	biq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BazelInvocationGroupBy{build: biq}
	grbuild.flds = &biq.ctx.Fields
	grbuild.label = bazelinvocation.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		InvocationID uuid.UUID `json:"invocation_id,omitempty"`
//	}
//
//	client.BazelInvocation.Query().
//		Select(bazelinvocation.FieldInvocationID).
//		Scan(ctx, &v)
func (biq *BazelInvocationQuery) Select(fields ...string) *BazelInvocationSelect {
	biq.ctx.Fields = append(biq.ctx.Fields, fields...)
	sbuild := &BazelInvocationSelect{BazelInvocationQuery: biq}
	sbuild.label = bazelinvocation.Label
	sbuild.flds, sbuild.scan = &biq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BazelInvocationSelect configured with the given aggregations.
func (biq *BazelInvocationQuery) Aggregate(fns ...AggregateFunc) *BazelInvocationSelect {
	return biq.Select().Aggregate(fns...)
}

func (biq *BazelInvocationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range biq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, biq); err != nil {
				return err
			}
		}
	}
	for _, f := range biq.ctx.Fields {
		if !bazelinvocation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if biq.path != nil {
		prev, err := biq.path(ctx)
		if err != nil {
			return err
		}
		biq.sql = prev
	}
	return nil
}

func (biq *BazelInvocationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BazelInvocation, error) {
	var (
		nodes       = []*BazelInvocation{}
		withFKs     = biq.withFKs
		_spec       = biq.querySpec()
		loadedTypes = [6]bool{
			biq.withEventFile != nil,
			biq.withBuild != nil,
			biq.withProblems != nil,
			biq.withMetrics != nil,
			biq.withTestCollection != nil,
			biq.withTargets != nil,
		}
	)
	if biq.withEventFile != nil || biq.withBuild != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, bazelinvocation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BazelInvocation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BazelInvocation{config: biq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(biq.modifiers) > 0 {
		_spec.Modifiers = biq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, biq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := biq.withEventFile; query != nil {
		if err := biq.loadEventFile(ctx, query, nodes, nil,
			func(n *BazelInvocation, e *EventFile) { n.Edges.EventFile = e }); err != nil {
			return nil, err
		}
	}
	if query := biq.withBuild; query != nil {
		if err := biq.loadBuild(ctx, query, nodes, nil,
			func(n *BazelInvocation, e *Build) { n.Edges.Build = e }); err != nil {
			return nil, err
		}
	}
	if query := biq.withProblems; query != nil {
		if err := biq.loadProblems(ctx, query, nodes,
			func(n *BazelInvocation) { n.Edges.Problems = []*BazelInvocationProblem{} },
			func(n *BazelInvocation, e *BazelInvocationProblem) { n.Edges.Problems = append(n.Edges.Problems, e) }); err != nil {
			return nil, err
		}
	}
	if query := biq.withMetrics; query != nil {
		if err := biq.loadMetrics(ctx, query, nodes, nil,
			func(n *BazelInvocation, e *Metrics) { n.Edges.Metrics = e }); err != nil {
			return nil, err
		}
	}
	if query := biq.withTestCollection; query != nil {
		if err := biq.loadTestCollection(ctx, query, nodes,
			func(n *BazelInvocation) { n.Edges.TestCollection = []*TestCollection{} },
			func(n *BazelInvocation, e *TestCollection) {
				n.Edges.TestCollection = append(n.Edges.TestCollection, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := biq.withTargets; query != nil {
		if err := biq.loadTargets(ctx, query, nodes,
			func(n *BazelInvocation) { n.Edges.Targets = []*TargetPair{} },
			func(n *BazelInvocation, e *TargetPair) { n.Edges.Targets = append(n.Edges.Targets, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range biq.withNamedProblems {
		if err := biq.loadProblems(ctx, query, nodes,
			func(n *BazelInvocation) { n.appendNamedProblems(name) },
			func(n *BazelInvocation, e *BazelInvocationProblem) { n.appendNamedProblems(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range biq.withNamedTestCollection {
		if err := biq.loadTestCollection(ctx, query, nodes,
			func(n *BazelInvocation) { n.appendNamedTestCollection(name) },
			func(n *BazelInvocation, e *TestCollection) { n.appendNamedTestCollection(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range biq.withNamedTargets {
		if err := biq.loadTargets(ctx, query, nodes,
			func(n *BazelInvocation) { n.appendNamedTargets(name) },
			func(n *BazelInvocation, e *TargetPair) { n.appendNamedTargets(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range biq.loadTotal {
		if err := biq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (biq *BazelInvocationQuery) loadEventFile(ctx context.Context, query *EventFileQuery, nodes []*BazelInvocation, init func(*BazelInvocation), assign func(*BazelInvocation, *EventFile)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*BazelInvocation)
	for i := range nodes {
		if nodes[i].event_file_bazel_invocation == nil {
			continue
		}
		fk := *nodes[i].event_file_bazel_invocation
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(eventfile.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "event_file_bazel_invocation" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (biq *BazelInvocationQuery) loadBuild(ctx context.Context, query *BuildQuery, nodes []*BazelInvocation, init func(*BazelInvocation), assign func(*BazelInvocation, *Build)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*BazelInvocation)
	for i := range nodes {
		if nodes[i].build_invocations == nil {
			continue
		}
		fk := *nodes[i].build_invocations
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(build.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "build_invocations" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (biq *BazelInvocationQuery) loadProblems(ctx context.Context, query *BazelInvocationProblemQuery, nodes []*BazelInvocation, init func(*BazelInvocation), assign func(*BazelInvocation, *BazelInvocationProblem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*BazelInvocation)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.BazelInvocationProblem(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(bazelinvocation.ProblemsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.bazel_invocation_problems
		if fk == nil {
			return fmt.Errorf(`foreign-key "bazel_invocation_problems" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "bazel_invocation_problems" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (biq *BazelInvocationQuery) loadMetrics(ctx context.Context, query *MetricsQuery, nodes []*BazelInvocation, init func(*BazelInvocation), assign func(*BazelInvocation, *Metrics)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*BazelInvocation)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.Metrics(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(bazelinvocation.MetricsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.bazel_invocation_metrics
		if fk == nil {
			return fmt.Errorf(`foreign-key "bazel_invocation_metrics" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "bazel_invocation_metrics" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (biq *BazelInvocationQuery) loadTestCollection(ctx context.Context, query *TestCollectionQuery, nodes []*BazelInvocation, init func(*BazelInvocation), assign func(*BazelInvocation, *TestCollection)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*BazelInvocation)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TestCollection(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(bazelinvocation.TestCollectionColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.bazel_invocation_test_collection
		if fk == nil {
			return fmt.Errorf(`foreign-key "bazel_invocation_test_collection" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "bazel_invocation_test_collection" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (biq *BazelInvocationQuery) loadTargets(ctx context.Context, query *TargetPairQuery, nodes []*BazelInvocation, init func(*BazelInvocation), assign func(*BazelInvocation, *TargetPair)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*BazelInvocation)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.TargetPair(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(bazelinvocation.TargetsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.bazel_invocation_targets
		if fk == nil {
			return fmt.Errorf(`foreign-key "bazel_invocation_targets" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "bazel_invocation_targets" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (biq *BazelInvocationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := biq.querySpec()
	if len(biq.modifiers) > 0 {
		_spec.Modifiers = biq.modifiers
	}
	_spec.Node.Columns = biq.ctx.Fields
	if len(biq.ctx.Fields) > 0 {
		_spec.Unique = biq.ctx.Unique != nil && *biq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, biq.driver, _spec)
}

func (biq *BazelInvocationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bazelinvocation.Table, bazelinvocation.Columns, sqlgraph.NewFieldSpec(bazelinvocation.FieldID, field.TypeInt))
	_spec.From = biq.sql
	if unique := biq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if biq.path != nil {
		_spec.Unique = true
	}
	if fields := biq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bazelinvocation.FieldID)
		for i := range fields {
			if fields[i] != bazelinvocation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := biq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := biq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := biq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := biq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (biq *BazelInvocationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(biq.driver.Dialect())
	t1 := builder.Table(bazelinvocation.Table)
	columns := biq.ctx.Fields
	if len(columns) == 0 {
		columns = bazelinvocation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if biq.sql != nil {
		selector = biq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if biq.ctx.Unique != nil && *biq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range biq.predicates {
		p(selector)
	}
	for _, p := range biq.order {
		p(selector)
	}
	if offset := biq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := biq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedProblems tells the query-builder to eager-load the nodes that are connected to the "problems"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (biq *BazelInvocationQuery) WithNamedProblems(name string, opts ...func(*BazelInvocationProblemQuery)) *BazelInvocationQuery {
	query := (&BazelInvocationProblemClient{config: biq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if biq.withNamedProblems == nil {
		biq.withNamedProblems = make(map[string]*BazelInvocationProblemQuery)
	}
	biq.withNamedProblems[name] = query
	return biq
}

// WithNamedTestCollection tells the query-builder to eager-load the nodes that are connected to the "test_collection"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (biq *BazelInvocationQuery) WithNamedTestCollection(name string, opts ...func(*TestCollectionQuery)) *BazelInvocationQuery {
	query := (&TestCollectionClient{config: biq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if biq.withNamedTestCollection == nil {
		biq.withNamedTestCollection = make(map[string]*TestCollectionQuery)
	}
	biq.withNamedTestCollection[name] = query
	return biq
}

// WithNamedTargets tells the query-builder to eager-load the nodes that are connected to the "targets"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (biq *BazelInvocationQuery) WithNamedTargets(name string, opts ...func(*TargetPairQuery)) *BazelInvocationQuery {
	query := (&TargetPairClient{config: biq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if biq.withNamedTargets == nil {
		biq.withNamedTargets = make(map[string]*TargetPairQuery)
	}
	biq.withNamedTargets[name] = query
	return biq
}

// BazelInvocationGroupBy is the group-by builder for BazelInvocation entities.
type BazelInvocationGroupBy struct {
	selector
	build *BazelInvocationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bigb *BazelInvocationGroupBy) Aggregate(fns ...AggregateFunc) *BazelInvocationGroupBy {
	bigb.fns = append(bigb.fns, fns...)
	return bigb
}

// Scan applies the selector query and scans the result into the given value.
func (bigb *BazelInvocationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bigb.build.ctx, "GroupBy")
	if err := bigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BazelInvocationQuery, *BazelInvocationGroupBy](ctx, bigb.build, bigb, bigb.build.inters, v)
}

func (bigb *BazelInvocationGroupBy) sqlScan(ctx context.Context, root *BazelInvocationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bigb.fns))
	for _, fn := range bigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bigb.flds)+len(bigb.fns))
		for _, f := range *bigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BazelInvocationSelect is the builder for selecting fields of BazelInvocation entities.
type BazelInvocationSelect struct {
	*BazelInvocationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bis *BazelInvocationSelect) Aggregate(fns ...AggregateFunc) *BazelInvocationSelect {
	bis.fns = append(bis.fns, fns...)
	return bis
}

// Scan applies the selector query and scans the result into the given value.
func (bis *BazelInvocationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bis.ctx, "Select")
	if err := bis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BazelInvocationQuery, *BazelInvocationSelect](ctx, bis.BazelInvocationQuery, bis, bis.inters, v)
}

func (bis *BazelInvocationSelect) sqlScan(ctx context.Context, root *BazelInvocationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bis.fns))
	for _, fn := range bis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
