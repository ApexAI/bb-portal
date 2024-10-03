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
	"github.com/buildbarn/bb-portal/ent/gen/ent/metrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/networkmetrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/systemnetworkstats"
)

// NetworkMetricsQuery is the builder for querying NetworkMetrics entities.
type NetworkMetricsQuery struct {
	config
	ctx                         *QueryContext
	order                       []networkmetrics.OrderOption
	inters                      []Interceptor
	predicates                  []predicate.NetworkMetrics
	withMetrics                 *MetricsQuery
	withSystemNetworkStats      *SystemNetworkStatsQuery
	modifiers                   []func(*sql.Selector)
	loadTotal                   []func(context.Context, []*NetworkMetrics) error
	withNamedMetrics            map[string]*MetricsQuery
	withNamedSystemNetworkStats map[string]*SystemNetworkStatsQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NetworkMetricsQuery builder.
func (nmq *NetworkMetricsQuery) Where(ps ...predicate.NetworkMetrics) *NetworkMetricsQuery {
	nmq.predicates = append(nmq.predicates, ps...)
	return nmq
}

// Limit the number of records to be returned by this query.
func (nmq *NetworkMetricsQuery) Limit(limit int) *NetworkMetricsQuery {
	nmq.ctx.Limit = &limit
	return nmq
}

// Offset to start from.
func (nmq *NetworkMetricsQuery) Offset(offset int) *NetworkMetricsQuery {
	nmq.ctx.Offset = &offset
	return nmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nmq *NetworkMetricsQuery) Unique(unique bool) *NetworkMetricsQuery {
	nmq.ctx.Unique = &unique
	return nmq
}

// Order specifies how the records should be ordered.
func (nmq *NetworkMetricsQuery) Order(o ...networkmetrics.OrderOption) *NetworkMetricsQuery {
	nmq.order = append(nmq.order, o...)
	return nmq
}

// QueryMetrics chains the current query on the "metrics" edge.
func (nmq *NetworkMetricsQuery) QueryMetrics() *MetricsQuery {
	query := (&MetricsClient{config: nmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(networkmetrics.Table, networkmetrics.FieldID, selector),
			sqlgraph.To(metrics.Table, metrics.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, networkmetrics.MetricsTable, networkmetrics.MetricsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(nmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySystemNetworkStats chains the current query on the "system_network_stats" edge.
func (nmq *NetworkMetricsQuery) QuerySystemNetworkStats() *SystemNetworkStatsQuery {
	query := (&SystemNetworkStatsClient{config: nmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(networkmetrics.Table, networkmetrics.FieldID, selector),
			sqlgraph.To(systemnetworkstats.Table, systemnetworkstats.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, networkmetrics.SystemNetworkStatsTable, networkmetrics.SystemNetworkStatsColumn),
		)
		fromU = sqlgraph.SetNeighbors(nmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NetworkMetrics entity from the query.
// Returns a *NotFoundError when no NetworkMetrics was found.
func (nmq *NetworkMetricsQuery) First(ctx context.Context) (*NetworkMetrics, error) {
	nodes, err := nmq.Limit(1).All(setContextOp(ctx, nmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{networkmetrics.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nmq *NetworkMetricsQuery) FirstX(ctx context.Context) *NetworkMetrics {
	node, err := nmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NetworkMetrics ID from the query.
// Returns a *NotFoundError when no NetworkMetrics ID was found.
func (nmq *NetworkMetricsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nmq.Limit(1).IDs(setContextOp(ctx, nmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{networkmetrics.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nmq *NetworkMetricsQuery) FirstIDX(ctx context.Context) int {
	id, err := nmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NetworkMetrics entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NetworkMetrics entity is found.
// Returns a *NotFoundError when no NetworkMetrics entities are found.
func (nmq *NetworkMetricsQuery) Only(ctx context.Context) (*NetworkMetrics, error) {
	nodes, err := nmq.Limit(2).All(setContextOp(ctx, nmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{networkmetrics.Label}
	default:
		return nil, &NotSingularError{networkmetrics.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nmq *NetworkMetricsQuery) OnlyX(ctx context.Context) *NetworkMetrics {
	node, err := nmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NetworkMetrics ID in the query.
// Returns a *NotSingularError when more than one NetworkMetrics ID is found.
// Returns a *NotFoundError when no entities are found.
func (nmq *NetworkMetricsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nmq.Limit(2).IDs(setContextOp(ctx, nmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{networkmetrics.Label}
	default:
		err = &NotSingularError{networkmetrics.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nmq *NetworkMetricsQuery) OnlyIDX(ctx context.Context) int {
	id, err := nmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NetworkMetricsSlice.
func (nmq *NetworkMetricsQuery) All(ctx context.Context) ([]*NetworkMetrics, error) {
	ctx = setContextOp(ctx, nmq.ctx, "All")
	if err := nmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NetworkMetrics, *NetworkMetricsQuery]()
	return withInterceptors[[]*NetworkMetrics](ctx, nmq, qr, nmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nmq *NetworkMetricsQuery) AllX(ctx context.Context) []*NetworkMetrics {
	nodes, err := nmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NetworkMetrics IDs.
func (nmq *NetworkMetricsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if nmq.ctx.Unique == nil && nmq.path != nil {
		nmq.Unique(true)
	}
	ctx = setContextOp(ctx, nmq.ctx, "IDs")
	if err = nmq.Select(networkmetrics.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nmq *NetworkMetricsQuery) IDsX(ctx context.Context) []int {
	ids, err := nmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nmq *NetworkMetricsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nmq.ctx, "Count")
	if err := nmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nmq, querierCount[*NetworkMetricsQuery](), nmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nmq *NetworkMetricsQuery) CountX(ctx context.Context) int {
	count, err := nmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nmq *NetworkMetricsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nmq.ctx, "Exist")
	switch _, err := nmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nmq *NetworkMetricsQuery) ExistX(ctx context.Context) bool {
	exist, err := nmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NetworkMetricsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nmq *NetworkMetricsQuery) Clone() *NetworkMetricsQuery {
	if nmq == nil {
		return nil
	}
	return &NetworkMetricsQuery{
		config:                 nmq.config,
		ctx:                    nmq.ctx.Clone(),
		order:                  append([]networkmetrics.OrderOption{}, nmq.order...),
		inters:                 append([]Interceptor{}, nmq.inters...),
		predicates:             append([]predicate.NetworkMetrics{}, nmq.predicates...),
		withMetrics:            nmq.withMetrics.Clone(),
		withSystemNetworkStats: nmq.withSystemNetworkStats.Clone(),
		// clone intermediate query.
		sql:  nmq.sql.Clone(),
		path: nmq.path,
	}
}

// WithMetrics tells the query-builder to eager-load the nodes that are connected to
// the "metrics" edge. The optional arguments are used to configure the query builder of the edge.
func (nmq *NetworkMetricsQuery) WithMetrics(opts ...func(*MetricsQuery)) *NetworkMetricsQuery {
	query := (&MetricsClient{config: nmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nmq.withMetrics = query
	return nmq
}

// WithSystemNetworkStats tells the query-builder to eager-load the nodes that are connected to
// the "system_network_stats" edge. The optional arguments are used to configure the query builder of the edge.
func (nmq *NetworkMetricsQuery) WithSystemNetworkStats(opts ...func(*SystemNetworkStatsQuery)) *NetworkMetricsQuery {
	query := (&SystemNetworkStatsClient{config: nmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nmq.withSystemNetworkStats = query
	return nmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (nmq *NetworkMetricsQuery) GroupBy(field string, fields ...string) *NetworkMetricsGroupBy {
	nmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NetworkMetricsGroupBy{build: nmq}
	grbuild.flds = &nmq.ctx.Fields
	grbuild.label = networkmetrics.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (nmq *NetworkMetricsQuery) Select(fields ...string) *NetworkMetricsSelect {
	nmq.ctx.Fields = append(nmq.ctx.Fields, fields...)
	sbuild := &NetworkMetricsSelect{NetworkMetricsQuery: nmq}
	sbuild.label = networkmetrics.Label
	sbuild.flds, sbuild.scan = &nmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NetworkMetricsSelect configured with the given aggregations.
func (nmq *NetworkMetricsQuery) Aggregate(fns ...AggregateFunc) *NetworkMetricsSelect {
	return nmq.Select().Aggregate(fns...)
}

func (nmq *NetworkMetricsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nmq); err != nil {
				return err
			}
		}
	}
	for _, f := range nmq.ctx.Fields {
		if !networkmetrics.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nmq.path != nil {
		prev, err := nmq.path(ctx)
		if err != nil {
			return err
		}
		nmq.sql = prev
	}
	return nil
}

func (nmq *NetworkMetricsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NetworkMetrics, error) {
	var (
		nodes       = []*NetworkMetrics{}
		_spec       = nmq.querySpec()
		loadedTypes = [2]bool{
			nmq.withMetrics != nil,
			nmq.withSystemNetworkStats != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NetworkMetrics).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NetworkMetrics{config: nmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(nmq.modifiers) > 0 {
		_spec.Modifiers = nmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nmq.withMetrics; query != nil {
		if err := nmq.loadMetrics(ctx, query, nodes,
			func(n *NetworkMetrics) { n.Edges.Metrics = []*Metrics{} },
			func(n *NetworkMetrics, e *Metrics) { n.Edges.Metrics = append(n.Edges.Metrics, e) }); err != nil {
			return nil, err
		}
	}
	if query := nmq.withSystemNetworkStats; query != nil {
		if err := nmq.loadSystemNetworkStats(ctx, query, nodes,
			func(n *NetworkMetrics) { n.Edges.SystemNetworkStats = []*SystemNetworkStats{} },
			func(n *NetworkMetrics, e *SystemNetworkStats) {
				n.Edges.SystemNetworkStats = append(n.Edges.SystemNetworkStats, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range nmq.withNamedMetrics {
		if err := nmq.loadMetrics(ctx, query, nodes,
			func(n *NetworkMetrics) { n.appendNamedMetrics(name) },
			func(n *NetworkMetrics, e *Metrics) { n.appendNamedMetrics(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range nmq.withNamedSystemNetworkStats {
		if err := nmq.loadSystemNetworkStats(ctx, query, nodes,
			func(n *NetworkMetrics) { n.appendNamedSystemNetworkStats(name) },
			func(n *NetworkMetrics, e *SystemNetworkStats) { n.appendNamedSystemNetworkStats(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range nmq.loadTotal {
		if err := nmq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nmq *NetworkMetricsQuery) loadMetrics(ctx context.Context, query *MetricsQuery, nodes []*NetworkMetrics, init func(*NetworkMetrics), assign func(*NetworkMetrics, *Metrics)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*NetworkMetrics)
	nids := make(map[int]map[*NetworkMetrics]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(networkmetrics.MetricsTable)
		s.Join(joinT).On(s.C(metrics.FieldID), joinT.C(networkmetrics.MetricsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(networkmetrics.MetricsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(networkmetrics.MetricsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*NetworkMetrics]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Metrics](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "metrics" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (nmq *NetworkMetricsQuery) loadSystemNetworkStats(ctx context.Context, query *SystemNetworkStatsQuery, nodes []*NetworkMetrics, init func(*NetworkMetrics), assign func(*NetworkMetrics, *SystemNetworkStats)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*NetworkMetrics)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.SystemNetworkStats(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(networkmetrics.SystemNetworkStatsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.network_metrics_system_network_stats
		if fk == nil {
			return fmt.Errorf(`foreign-key "network_metrics_system_network_stats" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "network_metrics_system_network_stats" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (nmq *NetworkMetricsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nmq.querySpec()
	if len(nmq.modifiers) > 0 {
		_spec.Modifiers = nmq.modifiers
	}
	_spec.Node.Columns = nmq.ctx.Fields
	if len(nmq.ctx.Fields) > 0 {
		_spec.Unique = nmq.ctx.Unique != nil && *nmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nmq.driver, _spec)
}

func (nmq *NetworkMetricsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(networkmetrics.Table, networkmetrics.Columns, sqlgraph.NewFieldSpec(networkmetrics.FieldID, field.TypeInt))
	_spec.From = nmq.sql
	if unique := nmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nmq.path != nil {
		_spec.Unique = true
	}
	if fields := nmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, networkmetrics.FieldID)
		for i := range fields {
			if fields[i] != networkmetrics.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nmq *NetworkMetricsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nmq.driver.Dialect())
	t1 := builder.Table(networkmetrics.Table)
	columns := nmq.ctx.Fields
	if len(columns) == 0 {
		columns = networkmetrics.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nmq.sql != nil {
		selector = nmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nmq.ctx.Unique != nil && *nmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nmq.predicates {
		p(selector)
	}
	for _, p := range nmq.order {
		p(selector)
	}
	if offset := nmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedMetrics tells the query-builder to eager-load the nodes that are connected to the "metrics"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (nmq *NetworkMetricsQuery) WithNamedMetrics(name string, opts ...func(*MetricsQuery)) *NetworkMetricsQuery {
	query := (&MetricsClient{config: nmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if nmq.withNamedMetrics == nil {
		nmq.withNamedMetrics = make(map[string]*MetricsQuery)
	}
	nmq.withNamedMetrics[name] = query
	return nmq
}

// WithNamedSystemNetworkStats tells the query-builder to eager-load the nodes that are connected to the "system_network_stats"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (nmq *NetworkMetricsQuery) WithNamedSystemNetworkStats(name string, opts ...func(*SystemNetworkStatsQuery)) *NetworkMetricsQuery {
	query := (&SystemNetworkStatsClient{config: nmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if nmq.withNamedSystemNetworkStats == nil {
		nmq.withNamedSystemNetworkStats = make(map[string]*SystemNetworkStatsQuery)
	}
	nmq.withNamedSystemNetworkStats[name] = query
	return nmq
}

// NetworkMetricsGroupBy is the group-by builder for NetworkMetrics entities.
type NetworkMetricsGroupBy struct {
	selector
	build *NetworkMetricsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nmgb *NetworkMetricsGroupBy) Aggregate(fns ...AggregateFunc) *NetworkMetricsGroupBy {
	nmgb.fns = append(nmgb.fns, fns...)
	return nmgb
}

// Scan applies the selector query and scans the result into the given value.
func (nmgb *NetworkMetricsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nmgb.build.ctx, "GroupBy")
	if err := nmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NetworkMetricsQuery, *NetworkMetricsGroupBy](ctx, nmgb.build, nmgb, nmgb.build.inters, v)
}

func (nmgb *NetworkMetricsGroupBy) sqlScan(ctx context.Context, root *NetworkMetricsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(nmgb.fns))
	for _, fn := range nmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*nmgb.flds)+len(nmgb.fns))
		for _, f := range *nmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*nmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NetworkMetricsSelect is the builder for selecting fields of NetworkMetrics entities.
type NetworkMetricsSelect struct {
	*NetworkMetricsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nms *NetworkMetricsSelect) Aggregate(fns ...AggregateFunc) *NetworkMetricsSelect {
	nms.fns = append(nms.fns, fns...)
	return nms
}

// Scan applies the selector query and scans the result into the given value.
func (nms *NetworkMetricsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nms.ctx, "Select")
	if err := nms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NetworkMetricsQuery, *NetworkMetricsSelect](ctx, nms.NetworkMetricsQuery, nms, nms.inters, v)
}

func (nms *NetworkMetricsSelect) sqlScan(ctx context.Context, root *NetworkMetricsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nms.fns))
	for _, fn := range nms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}