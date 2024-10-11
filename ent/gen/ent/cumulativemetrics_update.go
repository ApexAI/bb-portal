// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/cumulativemetrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/metrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
)

// CumulativeMetricsUpdate is the builder for updating CumulativeMetrics entities.
type CumulativeMetricsUpdate struct {
	config
	hooks    []Hook
	mutation *CumulativeMetricsMutation
}

// Where appends a list predicates to the CumulativeMetricsUpdate builder.
func (cmu *CumulativeMetricsUpdate) Where(ps ...predicate.CumulativeMetrics) *CumulativeMetricsUpdate {
	cmu.mutation.Where(ps...)
	return cmu
}

// SetNumAnalyses sets the "num_analyses" field.
func (cmu *CumulativeMetricsUpdate) SetNumAnalyses(i int32) *CumulativeMetricsUpdate {
	cmu.mutation.ResetNumAnalyses()
	cmu.mutation.SetNumAnalyses(i)
	return cmu
}

// SetNillableNumAnalyses sets the "num_analyses" field if the given value is not nil.
func (cmu *CumulativeMetricsUpdate) SetNillableNumAnalyses(i *int32) *CumulativeMetricsUpdate {
	if i != nil {
		cmu.SetNumAnalyses(*i)
	}
	return cmu
}

// AddNumAnalyses adds i to the "num_analyses" field.
func (cmu *CumulativeMetricsUpdate) AddNumAnalyses(i int32) *CumulativeMetricsUpdate {
	cmu.mutation.AddNumAnalyses(i)
	return cmu
}

// ClearNumAnalyses clears the value of the "num_analyses" field.
func (cmu *CumulativeMetricsUpdate) ClearNumAnalyses() *CumulativeMetricsUpdate {
	cmu.mutation.ClearNumAnalyses()
	return cmu
}

// SetNumBuilds sets the "num_builds" field.
func (cmu *CumulativeMetricsUpdate) SetNumBuilds(i int32) *CumulativeMetricsUpdate {
	cmu.mutation.ResetNumBuilds()
	cmu.mutation.SetNumBuilds(i)
	return cmu
}

// SetNillableNumBuilds sets the "num_builds" field if the given value is not nil.
func (cmu *CumulativeMetricsUpdate) SetNillableNumBuilds(i *int32) *CumulativeMetricsUpdate {
	if i != nil {
		cmu.SetNumBuilds(*i)
	}
	return cmu
}

// AddNumBuilds adds i to the "num_builds" field.
func (cmu *CumulativeMetricsUpdate) AddNumBuilds(i int32) *CumulativeMetricsUpdate {
	cmu.mutation.AddNumBuilds(i)
	return cmu
}

// ClearNumBuilds clears the value of the "num_builds" field.
func (cmu *CumulativeMetricsUpdate) ClearNumBuilds() *CumulativeMetricsUpdate {
	cmu.mutation.ClearNumBuilds()
	return cmu
}

// SetMetricsID sets the "metrics" edge to the Metrics entity by ID.
func (cmu *CumulativeMetricsUpdate) SetMetricsID(id int) *CumulativeMetricsUpdate {
	cmu.mutation.SetMetricsID(id)
	return cmu
}

// SetNillableMetricsID sets the "metrics" edge to the Metrics entity by ID if the given value is not nil.
func (cmu *CumulativeMetricsUpdate) SetNillableMetricsID(id *int) *CumulativeMetricsUpdate {
	if id != nil {
		cmu = cmu.SetMetricsID(*id)
	}
	return cmu
}

// SetMetrics sets the "metrics" edge to the Metrics entity.
func (cmu *CumulativeMetricsUpdate) SetMetrics(m *Metrics) *CumulativeMetricsUpdate {
	return cmu.SetMetricsID(m.ID)
}

// Mutation returns the CumulativeMetricsMutation object of the builder.
func (cmu *CumulativeMetricsUpdate) Mutation() *CumulativeMetricsMutation {
	return cmu.mutation
}

// ClearMetrics clears the "metrics" edge to the Metrics entity.
func (cmu *CumulativeMetricsUpdate) ClearMetrics() *CumulativeMetricsUpdate {
	cmu.mutation.ClearMetrics()
	return cmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cmu *CumulativeMetricsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cmu.sqlSave, cmu.mutation, cmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cmu *CumulativeMetricsUpdate) SaveX(ctx context.Context) int {
	affected, err := cmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cmu *CumulativeMetricsUpdate) Exec(ctx context.Context) error {
	_, err := cmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmu *CumulativeMetricsUpdate) ExecX(ctx context.Context) {
	if err := cmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cmu *CumulativeMetricsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(cumulativemetrics.Table, cumulativemetrics.Columns, sqlgraph.NewFieldSpec(cumulativemetrics.FieldID, field.TypeInt))
	if ps := cmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmu.mutation.NumAnalyses(); ok {
		_spec.SetField(cumulativemetrics.FieldNumAnalyses, field.TypeInt32, value)
	}
	if value, ok := cmu.mutation.AddedNumAnalyses(); ok {
		_spec.AddField(cumulativemetrics.FieldNumAnalyses, field.TypeInt32, value)
	}
	if cmu.mutation.NumAnalysesCleared() {
		_spec.ClearField(cumulativemetrics.FieldNumAnalyses, field.TypeInt32)
	}
	if value, ok := cmu.mutation.NumBuilds(); ok {
		_spec.SetField(cumulativemetrics.FieldNumBuilds, field.TypeInt32, value)
	}
	if value, ok := cmu.mutation.AddedNumBuilds(); ok {
		_spec.AddField(cumulativemetrics.FieldNumBuilds, field.TypeInt32, value)
	}
	if cmu.mutation.NumBuildsCleared() {
		_spec.ClearField(cumulativemetrics.FieldNumBuilds, field.TypeInt32)
	}
	if cmu.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   cumulativemetrics.MetricsTable,
			Columns: []string{cumulativemetrics.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metrics.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmu.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   cumulativemetrics.MetricsTable,
			Columns: []string{cumulativemetrics.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metrics.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cumulativemetrics.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cmu.mutation.done = true
	return n, nil
}

// CumulativeMetricsUpdateOne is the builder for updating a single CumulativeMetrics entity.
type CumulativeMetricsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CumulativeMetricsMutation
}

// SetNumAnalyses sets the "num_analyses" field.
func (cmuo *CumulativeMetricsUpdateOne) SetNumAnalyses(i int32) *CumulativeMetricsUpdateOne {
	cmuo.mutation.ResetNumAnalyses()
	cmuo.mutation.SetNumAnalyses(i)
	return cmuo
}

// SetNillableNumAnalyses sets the "num_analyses" field if the given value is not nil.
func (cmuo *CumulativeMetricsUpdateOne) SetNillableNumAnalyses(i *int32) *CumulativeMetricsUpdateOne {
	if i != nil {
		cmuo.SetNumAnalyses(*i)
	}
	return cmuo
}

// AddNumAnalyses adds i to the "num_analyses" field.
func (cmuo *CumulativeMetricsUpdateOne) AddNumAnalyses(i int32) *CumulativeMetricsUpdateOne {
	cmuo.mutation.AddNumAnalyses(i)
	return cmuo
}

// ClearNumAnalyses clears the value of the "num_analyses" field.
func (cmuo *CumulativeMetricsUpdateOne) ClearNumAnalyses() *CumulativeMetricsUpdateOne {
	cmuo.mutation.ClearNumAnalyses()
	return cmuo
}

// SetNumBuilds sets the "num_builds" field.
func (cmuo *CumulativeMetricsUpdateOne) SetNumBuilds(i int32) *CumulativeMetricsUpdateOne {
	cmuo.mutation.ResetNumBuilds()
	cmuo.mutation.SetNumBuilds(i)
	return cmuo
}

// SetNillableNumBuilds sets the "num_builds" field if the given value is not nil.
func (cmuo *CumulativeMetricsUpdateOne) SetNillableNumBuilds(i *int32) *CumulativeMetricsUpdateOne {
	if i != nil {
		cmuo.SetNumBuilds(*i)
	}
	return cmuo
}

// AddNumBuilds adds i to the "num_builds" field.
func (cmuo *CumulativeMetricsUpdateOne) AddNumBuilds(i int32) *CumulativeMetricsUpdateOne {
	cmuo.mutation.AddNumBuilds(i)
	return cmuo
}

// ClearNumBuilds clears the value of the "num_builds" field.
func (cmuo *CumulativeMetricsUpdateOne) ClearNumBuilds() *CumulativeMetricsUpdateOne {
	cmuo.mutation.ClearNumBuilds()
	return cmuo
}

// SetMetricsID sets the "metrics" edge to the Metrics entity by ID.
func (cmuo *CumulativeMetricsUpdateOne) SetMetricsID(id int) *CumulativeMetricsUpdateOne {
	cmuo.mutation.SetMetricsID(id)
	return cmuo
}

// SetNillableMetricsID sets the "metrics" edge to the Metrics entity by ID if the given value is not nil.
func (cmuo *CumulativeMetricsUpdateOne) SetNillableMetricsID(id *int) *CumulativeMetricsUpdateOne {
	if id != nil {
		cmuo = cmuo.SetMetricsID(*id)
	}
	return cmuo
}

// SetMetrics sets the "metrics" edge to the Metrics entity.
func (cmuo *CumulativeMetricsUpdateOne) SetMetrics(m *Metrics) *CumulativeMetricsUpdateOne {
	return cmuo.SetMetricsID(m.ID)
}

// Mutation returns the CumulativeMetricsMutation object of the builder.
func (cmuo *CumulativeMetricsUpdateOne) Mutation() *CumulativeMetricsMutation {
	return cmuo.mutation
}

// ClearMetrics clears the "metrics" edge to the Metrics entity.
func (cmuo *CumulativeMetricsUpdateOne) ClearMetrics() *CumulativeMetricsUpdateOne {
	cmuo.mutation.ClearMetrics()
	return cmuo
}

// Where appends a list predicates to the CumulativeMetricsUpdate builder.
func (cmuo *CumulativeMetricsUpdateOne) Where(ps ...predicate.CumulativeMetrics) *CumulativeMetricsUpdateOne {
	cmuo.mutation.Where(ps...)
	return cmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cmuo *CumulativeMetricsUpdateOne) Select(field string, fields ...string) *CumulativeMetricsUpdateOne {
	cmuo.fields = append([]string{field}, fields...)
	return cmuo
}

// Save executes the query and returns the updated CumulativeMetrics entity.
func (cmuo *CumulativeMetricsUpdateOne) Save(ctx context.Context) (*CumulativeMetrics, error) {
	return withHooks(ctx, cmuo.sqlSave, cmuo.mutation, cmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cmuo *CumulativeMetricsUpdateOne) SaveX(ctx context.Context) *CumulativeMetrics {
	node, err := cmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cmuo *CumulativeMetricsUpdateOne) Exec(ctx context.Context) error {
	_, err := cmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmuo *CumulativeMetricsUpdateOne) ExecX(ctx context.Context) {
	if err := cmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cmuo *CumulativeMetricsUpdateOne) sqlSave(ctx context.Context) (_node *CumulativeMetrics, err error) {
	_spec := sqlgraph.NewUpdateSpec(cumulativemetrics.Table, cumulativemetrics.Columns, sqlgraph.NewFieldSpec(cumulativemetrics.FieldID, field.TypeInt))
	id, ok := cmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CumulativeMetrics.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cumulativemetrics.FieldID)
		for _, f := range fields {
			if !cumulativemetrics.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cumulativemetrics.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmuo.mutation.NumAnalyses(); ok {
		_spec.SetField(cumulativemetrics.FieldNumAnalyses, field.TypeInt32, value)
	}
	if value, ok := cmuo.mutation.AddedNumAnalyses(); ok {
		_spec.AddField(cumulativemetrics.FieldNumAnalyses, field.TypeInt32, value)
	}
	if cmuo.mutation.NumAnalysesCleared() {
		_spec.ClearField(cumulativemetrics.FieldNumAnalyses, field.TypeInt32)
	}
	if value, ok := cmuo.mutation.NumBuilds(); ok {
		_spec.SetField(cumulativemetrics.FieldNumBuilds, field.TypeInt32, value)
	}
	if value, ok := cmuo.mutation.AddedNumBuilds(); ok {
		_spec.AddField(cumulativemetrics.FieldNumBuilds, field.TypeInt32, value)
	}
	if cmuo.mutation.NumBuildsCleared() {
		_spec.ClearField(cumulativemetrics.FieldNumBuilds, field.TypeInt32)
	}
	if cmuo.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   cumulativemetrics.MetricsTable,
			Columns: []string{cumulativemetrics.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metrics.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmuo.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   cumulativemetrics.MetricsTable,
			Columns: []string{cumulativemetrics.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metrics.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CumulativeMetrics{config: cmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cumulativemetrics.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cmuo.mutation.done = true
	return _node, nil
}
