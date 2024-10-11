// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/garbagemetrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/memorymetrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/metrics"
)

// MemoryMetricsCreate is the builder for creating a MemoryMetrics entity.
type MemoryMetricsCreate struct {
	config
	mutation *MemoryMetricsMutation
	hooks    []Hook
}

// SetPeakPostGcHeapSize sets the "peak_post_gc_heap_size" field.
func (mmc *MemoryMetricsCreate) SetPeakPostGcHeapSize(i int64) *MemoryMetricsCreate {
	mmc.mutation.SetPeakPostGcHeapSize(i)
	return mmc
}

// SetNillablePeakPostGcHeapSize sets the "peak_post_gc_heap_size" field if the given value is not nil.
func (mmc *MemoryMetricsCreate) SetNillablePeakPostGcHeapSize(i *int64) *MemoryMetricsCreate {
	if i != nil {
		mmc.SetPeakPostGcHeapSize(*i)
	}
	return mmc
}

// SetUsedHeapSizePostBuild sets the "used_heap_size_post_build" field.
func (mmc *MemoryMetricsCreate) SetUsedHeapSizePostBuild(i int64) *MemoryMetricsCreate {
	mmc.mutation.SetUsedHeapSizePostBuild(i)
	return mmc
}

// SetNillableUsedHeapSizePostBuild sets the "used_heap_size_post_build" field if the given value is not nil.
func (mmc *MemoryMetricsCreate) SetNillableUsedHeapSizePostBuild(i *int64) *MemoryMetricsCreate {
	if i != nil {
		mmc.SetUsedHeapSizePostBuild(*i)
	}
	return mmc
}

// SetPeakPostGcTenuredSpaceHeapSize sets the "peak_post_gc_tenured_space_heap_size" field.
func (mmc *MemoryMetricsCreate) SetPeakPostGcTenuredSpaceHeapSize(i int64) *MemoryMetricsCreate {
	mmc.mutation.SetPeakPostGcTenuredSpaceHeapSize(i)
	return mmc
}

// SetNillablePeakPostGcTenuredSpaceHeapSize sets the "peak_post_gc_tenured_space_heap_size" field if the given value is not nil.
func (mmc *MemoryMetricsCreate) SetNillablePeakPostGcTenuredSpaceHeapSize(i *int64) *MemoryMetricsCreate {
	if i != nil {
		mmc.SetPeakPostGcTenuredSpaceHeapSize(*i)
	}
	return mmc
}

// SetMetricsID sets the "metrics" edge to the Metrics entity by ID.
func (mmc *MemoryMetricsCreate) SetMetricsID(id int) *MemoryMetricsCreate {
	mmc.mutation.SetMetricsID(id)
	return mmc
}

// SetNillableMetricsID sets the "metrics" edge to the Metrics entity by ID if the given value is not nil.
func (mmc *MemoryMetricsCreate) SetNillableMetricsID(id *int) *MemoryMetricsCreate {
	if id != nil {
		mmc = mmc.SetMetricsID(*id)
	}
	return mmc
}

// SetMetrics sets the "metrics" edge to the Metrics entity.
func (mmc *MemoryMetricsCreate) SetMetrics(m *Metrics) *MemoryMetricsCreate {
	return mmc.SetMetricsID(m.ID)
}

// AddGarbageMetricIDs adds the "garbage_metrics" edge to the GarbageMetrics entity by IDs.
func (mmc *MemoryMetricsCreate) AddGarbageMetricIDs(ids ...int) *MemoryMetricsCreate {
	mmc.mutation.AddGarbageMetricIDs(ids...)
	return mmc
}

// AddGarbageMetrics adds the "garbage_metrics" edges to the GarbageMetrics entity.
func (mmc *MemoryMetricsCreate) AddGarbageMetrics(g ...*GarbageMetrics) *MemoryMetricsCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return mmc.AddGarbageMetricIDs(ids...)
}

// Mutation returns the MemoryMetricsMutation object of the builder.
func (mmc *MemoryMetricsCreate) Mutation() *MemoryMetricsMutation {
	return mmc.mutation
}

// Save creates the MemoryMetrics in the database.
func (mmc *MemoryMetricsCreate) Save(ctx context.Context) (*MemoryMetrics, error) {
	return withHooks(ctx, mmc.sqlSave, mmc.mutation, mmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mmc *MemoryMetricsCreate) SaveX(ctx context.Context) *MemoryMetrics {
	v, err := mmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mmc *MemoryMetricsCreate) Exec(ctx context.Context) error {
	_, err := mmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mmc *MemoryMetricsCreate) ExecX(ctx context.Context) {
	if err := mmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mmc *MemoryMetricsCreate) check() error {
	return nil
}

func (mmc *MemoryMetricsCreate) sqlSave(ctx context.Context) (*MemoryMetrics, error) {
	if err := mmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mmc.mutation.id = &_node.ID
	mmc.mutation.done = true
	return _node, nil
}

func (mmc *MemoryMetricsCreate) createSpec() (*MemoryMetrics, *sqlgraph.CreateSpec) {
	var (
		_node = &MemoryMetrics{config: mmc.config}
		_spec = sqlgraph.NewCreateSpec(memorymetrics.Table, sqlgraph.NewFieldSpec(memorymetrics.FieldID, field.TypeInt))
	)
	if value, ok := mmc.mutation.PeakPostGcHeapSize(); ok {
		_spec.SetField(memorymetrics.FieldPeakPostGcHeapSize, field.TypeInt64, value)
		_node.PeakPostGcHeapSize = value
	}
	if value, ok := mmc.mutation.UsedHeapSizePostBuild(); ok {
		_spec.SetField(memorymetrics.FieldUsedHeapSizePostBuild, field.TypeInt64, value)
		_node.UsedHeapSizePostBuild = value
	}
	if value, ok := mmc.mutation.PeakPostGcTenuredSpaceHeapSize(); ok {
		_spec.SetField(memorymetrics.FieldPeakPostGcTenuredSpaceHeapSize, field.TypeInt64, value)
		_node.PeakPostGcTenuredSpaceHeapSize = value
	}
	if nodes := mmc.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   memorymetrics.MetricsTable,
			Columns: []string{memorymetrics.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metrics.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.metrics_memory_metrics = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mmc.mutation.GarbageMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   memorymetrics.GarbageMetricsTable,
			Columns: []string{memorymetrics.GarbageMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(garbagemetrics.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemoryMetricsCreateBulk is the builder for creating many MemoryMetrics entities in bulk.
type MemoryMetricsCreateBulk struct {
	config
	err      error
	builders []*MemoryMetricsCreate
}

// Save creates the MemoryMetrics entities in the database.
func (mmcb *MemoryMetricsCreateBulk) Save(ctx context.Context) ([]*MemoryMetrics, error) {
	if mmcb.err != nil {
		return nil, mmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mmcb.builders))
	nodes := make([]*MemoryMetrics, len(mmcb.builders))
	mutators := make([]Mutator, len(mmcb.builders))
	for i := range mmcb.builders {
		func(i int, root context.Context) {
			builder := mmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemoryMetricsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mmcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mmcb *MemoryMetricsCreateBulk) SaveX(ctx context.Context) []*MemoryMetrics {
	v, err := mmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mmcb *MemoryMetricsCreateBulk) Exec(ctx context.Context) error {
	_, err := mmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mmcb *MemoryMetricsCreateBulk) ExecX(ctx context.Context) {
	if err := mmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
