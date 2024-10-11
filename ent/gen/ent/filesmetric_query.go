// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/artifactmetrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/filesmetric"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
)

// FilesMetricQuery is the builder for querying FilesMetric entities.
type FilesMetricQuery struct {
	config
	ctx                 *QueryContext
	order               []filesmetric.OrderOption
	inters              []Interceptor
	predicates          []predicate.FilesMetric
	withArtifactMetrics *ArtifactMetricsQuery
	withFKs             bool
	modifiers           []func(*sql.Selector)
	loadTotal           []func(context.Context, []*FilesMetric) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FilesMetricQuery builder.
func (fmq *FilesMetricQuery) Where(ps ...predicate.FilesMetric) *FilesMetricQuery {
	fmq.predicates = append(fmq.predicates, ps...)
	return fmq
}

// Limit the number of records to be returned by this query.
func (fmq *FilesMetricQuery) Limit(limit int) *FilesMetricQuery {
	fmq.ctx.Limit = &limit
	return fmq
}

// Offset to start from.
func (fmq *FilesMetricQuery) Offset(offset int) *FilesMetricQuery {
	fmq.ctx.Offset = &offset
	return fmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fmq *FilesMetricQuery) Unique(unique bool) *FilesMetricQuery {
	fmq.ctx.Unique = &unique
	return fmq
}

// Order specifies how the records should be ordered.
func (fmq *FilesMetricQuery) Order(o ...filesmetric.OrderOption) *FilesMetricQuery {
	fmq.order = append(fmq.order, o...)
	return fmq
}

// QueryArtifactMetrics chains the current query on the "artifact_metrics" edge.
func (fmq *FilesMetricQuery) QueryArtifactMetrics() *ArtifactMetricsQuery {
	query := (&ArtifactMetricsClient{config: fmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(filesmetric.Table, filesmetric.FieldID, selector),
			sqlgraph.To(artifactmetrics.Table, artifactmetrics.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, filesmetric.ArtifactMetricsTable, filesmetric.ArtifactMetricsColumn),
		)
		fromU = sqlgraph.SetNeighbors(fmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FilesMetric entity from the query.
// Returns a *NotFoundError when no FilesMetric was found.
func (fmq *FilesMetricQuery) First(ctx context.Context) (*FilesMetric, error) {
	nodes, err := fmq.Limit(1).All(setContextOp(ctx, fmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{filesmetric.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fmq *FilesMetricQuery) FirstX(ctx context.Context) *FilesMetric {
	node, err := fmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FilesMetric ID from the query.
// Returns a *NotFoundError when no FilesMetric ID was found.
func (fmq *FilesMetricQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fmq.Limit(1).IDs(setContextOp(ctx, fmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{filesmetric.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fmq *FilesMetricQuery) FirstIDX(ctx context.Context) int {
	id, err := fmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FilesMetric entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FilesMetric entity is found.
// Returns a *NotFoundError when no FilesMetric entities are found.
func (fmq *FilesMetricQuery) Only(ctx context.Context) (*FilesMetric, error) {
	nodes, err := fmq.Limit(2).All(setContextOp(ctx, fmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{filesmetric.Label}
	default:
		return nil, &NotSingularError{filesmetric.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fmq *FilesMetricQuery) OnlyX(ctx context.Context) *FilesMetric {
	node, err := fmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FilesMetric ID in the query.
// Returns a *NotSingularError when more than one FilesMetric ID is found.
// Returns a *NotFoundError when no entities are found.
func (fmq *FilesMetricQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fmq.Limit(2).IDs(setContextOp(ctx, fmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{filesmetric.Label}
	default:
		err = &NotSingularError{filesmetric.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fmq *FilesMetricQuery) OnlyIDX(ctx context.Context) int {
	id, err := fmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FilesMetrics.
func (fmq *FilesMetricQuery) All(ctx context.Context) ([]*FilesMetric, error) {
	ctx = setContextOp(ctx, fmq.ctx, "All")
	if err := fmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FilesMetric, *FilesMetricQuery]()
	return withInterceptors[[]*FilesMetric](ctx, fmq, qr, fmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fmq *FilesMetricQuery) AllX(ctx context.Context) []*FilesMetric {
	nodes, err := fmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FilesMetric IDs.
func (fmq *FilesMetricQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fmq.ctx.Unique == nil && fmq.path != nil {
		fmq.Unique(true)
	}
	ctx = setContextOp(ctx, fmq.ctx, "IDs")
	if err = fmq.Select(filesmetric.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fmq *FilesMetricQuery) IDsX(ctx context.Context) []int {
	ids, err := fmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fmq *FilesMetricQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fmq.ctx, "Count")
	if err := fmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fmq, querierCount[*FilesMetricQuery](), fmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fmq *FilesMetricQuery) CountX(ctx context.Context) int {
	count, err := fmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fmq *FilesMetricQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fmq.ctx, "Exist")
	switch _, err := fmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fmq *FilesMetricQuery) ExistX(ctx context.Context) bool {
	exist, err := fmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FilesMetricQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fmq *FilesMetricQuery) Clone() *FilesMetricQuery {
	if fmq == nil {
		return nil
	}
	return &FilesMetricQuery{
		config:              fmq.config,
		ctx:                 fmq.ctx.Clone(),
		order:               append([]filesmetric.OrderOption{}, fmq.order...),
		inters:              append([]Interceptor{}, fmq.inters...),
		predicates:          append([]predicate.FilesMetric{}, fmq.predicates...),
		withArtifactMetrics: fmq.withArtifactMetrics.Clone(),
		// clone intermediate query.
		sql:  fmq.sql.Clone(),
		path: fmq.path,
	}
}

// WithArtifactMetrics tells the query-builder to eager-load the nodes that are connected to
// the "artifact_metrics" edge. The optional arguments are used to configure the query builder of the edge.
func (fmq *FilesMetricQuery) WithArtifactMetrics(opts ...func(*ArtifactMetricsQuery)) *FilesMetricQuery {
	query := (&ArtifactMetricsClient{config: fmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fmq.withArtifactMetrics = query
	return fmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SizeInBytes int64 `json:"size_in_bytes,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FilesMetric.Query().
//		GroupBy(filesmetric.FieldSizeInBytes).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fmq *FilesMetricQuery) GroupBy(field string, fields ...string) *FilesMetricGroupBy {
	fmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FilesMetricGroupBy{build: fmq}
	grbuild.flds = &fmq.ctx.Fields
	grbuild.label = filesmetric.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SizeInBytes int64 `json:"size_in_bytes,omitempty"`
//	}
//
//	client.FilesMetric.Query().
//		Select(filesmetric.FieldSizeInBytes).
//		Scan(ctx, &v)
func (fmq *FilesMetricQuery) Select(fields ...string) *FilesMetricSelect {
	fmq.ctx.Fields = append(fmq.ctx.Fields, fields...)
	sbuild := &FilesMetricSelect{FilesMetricQuery: fmq}
	sbuild.label = filesmetric.Label
	sbuild.flds, sbuild.scan = &fmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FilesMetricSelect configured with the given aggregations.
func (fmq *FilesMetricQuery) Aggregate(fns ...AggregateFunc) *FilesMetricSelect {
	return fmq.Select().Aggregate(fns...)
}

func (fmq *FilesMetricQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fmq); err != nil {
				return err
			}
		}
	}
	for _, f := range fmq.ctx.Fields {
		if !filesmetric.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fmq.path != nil {
		prev, err := fmq.path(ctx)
		if err != nil {
			return err
		}
		fmq.sql = prev
	}
	return nil
}

func (fmq *FilesMetricQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FilesMetric, error) {
	var (
		nodes       = []*FilesMetric{}
		withFKs     = fmq.withFKs
		_spec       = fmq.querySpec()
		loadedTypes = [1]bool{
			fmq.withArtifactMetrics != nil,
		}
	)
	if fmq.withArtifactMetrics != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, filesmetric.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FilesMetric).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FilesMetric{config: fmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(fmq.modifiers) > 0 {
		_spec.Modifiers = fmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fmq.withArtifactMetrics; query != nil {
		if err := fmq.loadArtifactMetrics(ctx, query, nodes, nil,
			func(n *FilesMetric, e *ArtifactMetrics) { n.Edges.ArtifactMetrics = e }); err != nil {
			return nil, err
		}
	}
	for i := range fmq.loadTotal {
		if err := fmq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fmq *FilesMetricQuery) loadArtifactMetrics(ctx context.Context, query *ArtifactMetricsQuery, nodes []*FilesMetric, init func(*FilesMetric), assign func(*FilesMetric, *ArtifactMetrics)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*FilesMetric)
	for i := range nodes {
		if nodes[i].artifact_metrics_top_level_artifacts == nil {
			continue
		}
		fk := *nodes[i].artifact_metrics_top_level_artifacts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(artifactmetrics.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "artifact_metrics_top_level_artifacts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fmq *FilesMetricQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fmq.querySpec()
	if len(fmq.modifiers) > 0 {
		_spec.Modifiers = fmq.modifiers
	}
	_spec.Node.Columns = fmq.ctx.Fields
	if len(fmq.ctx.Fields) > 0 {
		_spec.Unique = fmq.ctx.Unique != nil && *fmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fmq.driver, _spec)
}

func (fmq *FilesMetricQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(filesmetric.Table, filesmetric.Columns, sqlgraph.NewFieldSpec(filesmetric.FieldID, field.TypeInt))
	_spec.From = fmq.sql
	if unique := fmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fmq.path != nil {
		_spec.Unique = true
	}
	if fields := fmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filesmetric.FieldID)
		for i := range fields {
			if fields[i] != filesmetric.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fmq *FilesMetricQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fmq.driver.Dialect())
	t1 := builder.Table(filesmetric.Table)
	columns := fmq.ctx.Fields
	if len(columns) == 0 {
		columns = filesmetric.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fmq.sql != nil {
		selector = fmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fmq.ctx.Unique != nil && *fmq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fmq.predicates {
		p(selector)
	}
	for _, p := range fmq.order {
		p(selector)
	}
	if offset := fmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FilesMetricGroupBy is the group-by builder for FilesMetric entities.
type FilesMetricGroupBy struct {
	selector
	build *FilesMetricQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fmgb *FilesMetricGroupBy) Aggregate(fns ...AggregateFunc) *FilesMetricGroupBy {
	fmgb.fns = append(fmgb.fns, fns...)
	return fmgb
}

// Scan applies the selector query and scans the result into the given value.
func (fmgb *FilesMetricGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fmgb.build.ctx, "GroupBy")
	if err := fmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FilesMetricQuery, *FilesMetricGroupBy](ctx, fmgb.build, fmgb, fmgb.build.inters, v)
}

func (fmgb *FilesMetricGroupBy) sqlScan(ctx context.Context, root *FilesMetricQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fmgb.fns))
	for _, fn := range fmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fmgb.flds)+len(fmgb.fns))
		for _, f := range *fmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FilesMetricSelect is the builder for selecting fields of FilesMetric entities.
type FilesMetricSelect struct {
	*FilesMetricQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fms *FilesMetricSelect) Aggregate(fns ...AggregateFunc) *FilesMetricSelect {
	fms.fns = append(fms.fns, fns...)
	return fms
}

// Scan applies the selector query and scans the result into the given value.
func (fms *FilesMetricSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fms.ctx, "Select")
	if err := fms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FilesMetricQuery, *FilesMetricSelect](ctx, fms.FilesMetricQuery, fms, fms.inters, v)
}

func (fms *FilesMetricSelect) sqlScan(ctx context.Context, root *FilesMetricQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fms.fns))
	for _, fn := range fms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
