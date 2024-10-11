// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/actioncachestatistics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/actionsummary"
	"github.com/buildbarn/bb-portal/ent/gen/ent/missdetail"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
)

// ActionCacheStatisticsUpdate is the builder for updating ActionCacheStatistics entities.
type ActionCacheStatisticsUpdate struct {
	config
	hooks    []Hook
	mutation *ActionCacheStatisticsMutation
}

// Where appends a list predicates to the ActionCacheStatisticsUpdate builder.
func (acsu *ActionCacheStatisticsUpdate) Where(ps ...predicate.ActionCacheStatistics) *ActionCacheStatisticsUpdate {
	acsu.mutation.Where(ps...)
	return acsu
}

// SetSizeInBytes sets the "size_in_bytes" field.
func (acsu *ActionCacheStatisticsUpdate) SetSizeInBytes(u uint64) *ActionCacheStatisticsUpdate {
	acsu.mutation.ResetSizeInBytes()
	acsu.mutation.SetSizeInBytes(u)
	return acsu
}

// SetNillableSizeInBytes sets the "size_in_bytes" field if the given value is not nil.
func (acsu *ActionCacheStatisticsUpdate) SetNillableSizeInBytes(u *uint64) *ActionCacheStatisticsUpdate {
	if u != nil {
		acsu.SetSizeInBytes(*u)
	}
	return acsu
}

// AddSizeInBytes adds u to the "size_in_bytes" field.
func (acsu *ActionCacheStatisticsUpdate) AddSizeInBytes(u int64) *ActionCacheStatisticsUpdate {
	acsu.mutation.AddSizeInBytes(u)
	return acsu
}

// ClearSizeInBytes clears the value of the "size_in_bytes" field.
func (acsu *ActionCacheStatisticsUpdate) ClearSizeInBytes() *ActionCacheStatisticsUpdate {
	acsu.mutation.ClearSizeInBytes()
	return acsu
}

// SetSaveTimeInMs sets the "save_time_in_ms" field.
func (acsu *ActionCacheStatisticsUpdate) SetSaveTimeInMs(u uint64) *ActionCacheStatisticsUpdate {
	acsu.mutation.ResetSaveTimeInMs()
	acsu.mutation.SetSaveTimeInMs(u)
	return acsu
}

// SetNillableSaveTimeInMs sets the "save_time_in_ms" field if the given value is not nil.
func (acsu *ActionCacheStatisticsUpdate) SetNillableSaveTimeInMs(u *uint64) *ActionCacheStatisticsUpdate {
	if u != nil {
		acsu.SetSaveTimeInMs(*u)
	}
	return acsu
}

// AddSaveTimeInMs adds u to the "save_time_in_ms" field.
func (acsu *ActionCacheStatisticsUpdate) AddSaveTimeInMs(u int64) *ActionCacheStatisticsUpdate {
	acsu.mutation.AddSaveTimeInMs(u)
	return acsu
}

// ClearSaveTimeInMs clears the value of the "save_time_in_ms" field.
func (acsu *ActionCacheStatisticsUpdate) ClearSaveTimeInMs() *ActionCacheStatisticsUpdate {
	acsu.mutation.ClearSaveTimeInMs()
	return acsu
}

// SetLoadTimeInMs sets the "load_time_in_ms" field.
func (acsu *ActionCacheStatisticsUpdate) SetLoadTimeInMs(i int64) *ActionCacheStatisticsUpdate {
	acsu.mutation.ResetLoadTimeInMs()
	acsu.mutation.SetLoadTimeInMs(i)
	return acsu
}

// SetNillableLoadTimeInMs sets the "load_time_in_ms" field if the given value is not nil.
func (acsu *ActionCacheStatisticsUpdate) SetNillableLoadTimeInMs(i *int64) *ActionCacheStatisticsUpdate {
	if i != nil {
		acsu.SetLoadTimeInMs(*i)
	}
	return acsu
}

// AddLoadTimeInMs adds i to the "load_time_in_ms" field.
func (acsu *ActionCacheStatisticsUpdate) AddLoadTimeInMs(i int64) *ActionCacheStatisticsUpdate {
	acsu.mutation.AddLoadTimeInMs(i)
	return acsu
}

// ClearLoadTimeInMs clears the value of the "load_time_in_ms" field.
func (acsu *ActionCacheStatisticsUpdate) ClearLoadTimeInMs() *ActionCacheStatisticsUpdate {
	acsu.mutation.ClearLoadTimeInMs()
	return acsu
}

// SetHits sets the "hits" field.
func (acsu *ActionCacheStatisticsUpdate) SetHits(i int32) *ActionCacheStatisticsUpdate {
	acsu.mutation.ResetHits()
	acsu.mutation.SetHits(i)
	return acsu
}

// SetNillableHits sets the "hits" field if the given value is not nil.
func (acsu *ActionCacheStatisticsUpdate) SetNillableHits(i *int32) *ActionCacheStatisticsUpdate {
	if i != nil {
		acsu.SetHits(*i)
	}
	return acsu
}

// AddHits adds i to the "hits" field.
func (acsu *ActionCacheStatisticsUpdate) AddHits(i int32) *ActionCacheStatisticsUpdate {
	acsu.mutation.AddHits(i)
	return acsu
}

// ClearHits clears the value of the "hits" field.
func (acsu *ActionCacheStatisticsUpdate) ClearHits() *ActionCacheStatisticsUpdate {
	acsu.mutation.ClearHits()
	return acsu
}

// SetMisses sets the "misses" field.
func (acsu *ActionCacheStatisticsUpdate) SetMisses(i int32) *ActionCacheStatisticsUpdate {
	acsu.mutation.ResetMisses()
	acsu.mutation.SetMisses(i)
	return acsu
}

// SetNillableMisses sets the "misses" field if the given value is not nil.
func (acsu *ActionCacheStatisticsUpdate) SetNillableMisses(i *int32) *ActionCacheStatisticsUpdate {
	if i != nil {
		acsu.SetMisses(*i)
	}
	return acsu
}

// AddMisses adds i to the "misses" field.
func (acsu *ActionCacheStatisticsUpdate) AddMisses(i int32) *ActionCacheStatisticsUpdate {
	acsu.mutation.AddMisses(i)
	return acsu
}

// ClearMisses clears the value of the "misses" field.
func (acsu *ActionCacheStatisticsUpdate) ClearMisses() *ActionCacheStatisticsUpdate {
	acsu.mutation.ClearMisses()
	return acsu
}

// SetActionSummaryID sets the "action_summary" edge to the ActionSummary entity by ID.
func (acsu *ActionCacheStatisticsUpdate) SetActionSummaryID(id int) *ActionCacheStatisticsUpdate {
	acsu.mutation.SetActionSummaryID(id)
	return acsu
}

// SetNillableActionSummaryID sets the "action_summary" edge to the ActionSummary entity by ID if the given value is not nil.
func (acsu *ActionCacheStatisticsUpdate) SetNillableActionSummaryID(id *int) *ActionCacheStatisticsUpdate {
	if id != nil {
		acsu = acsu.SetActionSummaryID(*id)
	}
	return acsu
}

// SetActionSummary sets the "action_summary" edge to the ActionSummary entity.
func (acsu *ActionCacheStatisticsUpdate) SetActionSummary(a *ActionSummary) *ActionCacheStatisticsUpdate {
	return acsu.SetActionSummaryID(a.ID)
}

// AddMissDetailIDs adds the "miss_details" edge to the MissDetail entity by IDs.
func (acsu *ActionCacheStatisticsUpdate) AddMissDetailIDs(ids ...int) *ActionCacheStatisticsUpdate {
	acsu.mutation.AddMissDetailIDs(ids...)
	return acsu
}

// AddMissDetails adds the "miss_details" edges to the MissDetail entity.
func (acsu *ActionCacheStatisticsUpdate) AddMissDetails(m ...*MissDetail) *ActionCacheStatisticsUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return acsu.AddMissDetailIDs(ids...)
}

// Mutation returns the ActionCacheStatisticsMutation object of the builder.
func (acsu *ActionCacheStatisticsUpdate) Mutation() *ActionCacheStatisticsMutation {
	return acsu.mutation
}

// ClearActionSummary clears the "action_summary" edge to the ActionSummary entity.
func (acsu *ActionCacheStatisticsUpdate) ClearActionSummary() *ActionCacheStatisticsUpdate {
	acsu.mutation.ClearActionSummary()
	return acsu
}

// ClearMissDetails clears all "miss_details" edges to the MissDetail entity.
func (acsu *ActionCacheStatisticsUpdate) ClearMissDetails() *ActionCacheStatisticsUpdate {
	acsu.mutation.ClearMissDetails()
	return acsu
}

// RemoveMissDetailIDs removes the "miss_details" edge to MissDetail entities by IDs.
func (acsu *ActionCacheStatisticsUpdate) RemoveMissDetailIDs(ids ...int) *ActionCacheStatisticsUpdate {
	acsu.mutation.RemoveMissDetailIDs(ids...)
	return acsu
}

// RemoveMissDetails removes "miss_details" edges to MissDetail entities.
func (acsu *ActionCacheStatisticsUpdate) RemoveMissDetails(m ...*MissDetail) *ActionCacheStatisticsUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return acsu.RemoveMissDetailIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acsu *ActionCacheStatisticsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, acsu.sqlSave, acsu.mutation, acsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acsu *ActionCacheStatisticsUpdate) SaveX(ctx context.Context) int {
	affected, err := acsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acsu *ActionCacheStatisticsUpdate) Exec(ctx context.Context) error {
	_, err := acsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acsu *ActionCacheStatisticsUpdate) ExecX(ctx context.Context) {
	if err := acsu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (acsu *ActionCacheStatisticsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(actioncachestatistics.Table, actioncachestatistics.Columns, sqlgraph.NewFieldSpec(actioncachestatistics.FieldID, field.TypeInt))
	if ps := acsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acsu.mutation.SizeInBytes(); ok {
		_spec.SetField(actioncachestatistics.FieldSizeInBytes, field.TypeUint64, value)
	}
	if value, ok := acsu.mutation.AddedSizeInBytes(); ok {
		_spec.AddField(actioncachestatistics.FieldSizeInBytes, field.TypeUint64, value)
	}
	if acsu.mutation.SizeInBytesCleared() {
		_spec.ClearField(actioncachestatistics.FieldSizeInBytes, field.TypeUint64)
	}
	if value, ok := acsu.mutation.SaveTimeInMs(); ok {
		_spec.SetField(actioncachestatistics.FieldSaveTimeInMs, field.TypeUint64, value)
	}
	if value, ok := acsu.mutation.AddedSaveTimeInMs(); ok {
		_spec.AddField(actioncachestatistics.FieldSaveTimeInMs, field.TypeUint64, value)
	}
	if acsu.mutation.SaveTimeInMsCleared() {
		_spec.ClearField(actioncachestatistics.FieldSaveTimeInMs, field.TypeUint64)
	}
	if value, ok := acsu.mutation.LoadTimeInMs(); ok {
		_spec.SetField(actioncachestatistics.FieldLoadTimeInMs, field.TypeInt64, value)
	}
	if value, ok := acsu.mutation.AddedLoadTimeInMs(); ok {
		_spec.AddField(actioncachestatistics.FieldLoadTimeInMs, field.TypeInt64, value)
	}
	if acsu.mutation.LoadTimeInMsCleared() {
		_spec.ClearField(actioncachestatistics.FieldLoadTimeInMs, field.TypeInt64)
	}
	if value, ok := acsu.mutation.Hits(); ok {
		_spec.SetField(actioncachestatistics.FieldHits, field.TypeInt32, value)
	}
	if value, ok := acsu.mutation.AddedHits(); ok {
		_spec.AddField(actioncachestatistics.FieldHits, field.TypeInt32, value)
	}
	if acsu.mutation.HitsCleared() {
		_spec.ClearField(actioncachestatistics.FieldHits, field.TypeInt32)
	}
	if value, ok := acsu.mutation.Misses(); ok {
		_spec.SetField(actioncachestatistics.FieldMisses, field.TypeInt32, value)
	}
	if value, ok := acsu.mutation.AddedMisses(); ok {
		_spec.AddField(actioncachestatistics.FieldMisses, field.TypeInt32, value)
	}
	if acsu.mutation.MissesCleared() {
		_spec.ClearField(actioncachestatistics.FieldMisses, field.TypeInt32)
	}
	if acsu.mutation.ActionSummaryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   actioncachestatistics.ActionSummaryTable,
			Columns: []string{actioncachestatistics.ActionSummaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(actionsummary.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acsu.mutation.ActionSummaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   actioncachestatistics.ActionSummaryTable,
			Columns: []string{actioncachestatistics.ActionSummaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(actionsummary.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if acsu.mutation.MissDetailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   actioncachestatistics.MissDetailsTable,
			Columns: []string{actioncachestatistics.MissDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missdetail.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acsu.mutation.RemovedMissDetailsIDs(); len(nodes) > 0 && !acsu.mutation.MissDetailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   actioncachestatistics.MissDetailsTable,
			Columns: []string{actioncachestatistics.MissDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acsu.mutation.MissDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   actioncachestatistics.MissDetailsTable,
			Columns: []string{actioncachestatistics.MissDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{actioncachestatistics.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	acsu.mutation.done = true
	return n, nil
}

// ActionCacheStatisticsUpdateOne is the builder for updating a single ActionCacheStatistics entity.
type ActionCacheStatisticsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActionCacheStatisticsMutation
}

// SetSizeInBytes sets the "size_in_bytes" field.
func (acsuo *ActionCacheStatisticsUpdateOne) SetSizeInBytes(u uint64) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ResetSizeInBytes()
	acsuo.mutation.SetSizeInBytes(u)
	return acsuo
}

// SetNillableSizeInBytes sets the "size_in_bytes" field if the given value is not nil.
func (acsuo *ActionCacheStatisticsUpdateOne) SetNillableSizeInBytes(u *uint64) *ActionCacheStatisticsUpdateOne {
	if u != nil {
		acsuo.SetSizeInBytes(*u)
	}
	return acsuo
}

// AddSizeInBytes adds u to the "size_in_bytes" field.
func (acsuo *ActionCacheStatisticsUpdateOne) AddSizeInBytes(u int64) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.AddSizeInBytes(u)
	return acsuo
}

// ClearSizeInBytes clears the value of the "size_in_bytes" field.
func (acsuo *ActionCacheStatisticsUpdateOne) ClearSizeInBytes() *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ClearSizeInBytes()
	return acsuo
}

// SetSaveTimeInMs sets the "save_time_in_ms" field.
func (acsuo *ActionCacheStatisticsUpdateOne) SetSaveTimeInMs(u uint64) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ResetSaveTimeInMs()
	acsuo.mutation.SetSaveTimeInMs(u)
	return acsuo
}

// SetNillableSaveTimeInMs sets the "save_time_in_ms" field if the given value is not nil.
func (acsuo *ActionCacheStatisticsUpdateOne) SetNillableSaveTimeInMs(u *uint64) *ActionCacheStatisticsUpdateOne {
	if u != nil {
		acsuo.SetSaveTimeInMs(*u)
	}
	return acsuo
}

// AddSaveTimeInMs adds u to the "save_time_in_ms" field.
func (acsuo *ActionCacheStatisticsUpdateOne) AddSaveTimeInMs(u int64) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.AddSaveTimeInMs(u)
	return acsuo
}

// ClearSaveTimeInMs clears the value of the "save_time_in_ms" field.
func (acsuo *ActionCacheStatisticsUpdateOne) ClearSaveTimeInMs() *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ClearSaveTimeInMs()
	return acsuo
}

// SetLoadTimeInMs sets the "load_time_in_ms" field.
func (acsuo *ActionCacheStatisticsUpdateOne) SetLoadTimeInMs(i int64) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ResetLoadTimeInMs()
	acsuo.mutation.SetLoadTimeInMs(i)
	return acsuo
}

// SetNillableLoadTimeInMs sets the "load_time_in_ms" field if the given value is not nil.
func (acsuo *ActionCacheStatisticsUpdateOne) SetNillableLoadTimeInMs(i *int64) *ActionCacheStatisticsUpdateOne {
	if i != nil {
		acsuo.SetLoadTimeInMs(*i)
	}
	return acsuo
}

// AddLoadTimeInMs adds i to the "load_time_in_ms" field.
func (acsuo *ActionCacheStatisticsUpdateOne) AddLoadTimeInMs(i int64) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.AddLoadTimeInMs(i)
	return acsuo
}

// ClearLoadTimeInMs clears the value of the "load_time_in_ms" field.
func (acsuo *ActionCacheStatisticsUpdateOne) ClearLoadTimeInMs() *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ClearLoadTimeInMs()
	return acsuo
}

// SetHits sets the "hits" field.
func (acsuo *ActionCacheStatisticsUpdateOne) SetHits(i int32) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ResetHits()
	acsuo.mutation.SetHits(i)
	return acsuo
}

// SetNillableHits sets the "hits" field if the given value is not nil.
func (acsuo *ActionCacheStatisticsUpdateOne) SetNillableHits(i *int32) *ActionCacheStatisticsUpdateOne {
	if i != nil {
		acsuo.SetHits(*i)
	}
	return acsuo
}

// AddHits adds i to the "hits" field.
func (acsuo *ActionCacheStatisticsUpdateOne) AddHits(i int32) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.AddHits(i)
	return acsuo
}

// ClearHits clears the value of the "hits" field.
func (acsuo *ActionCacheStatisticsUpdateOne) ClearHits() *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ClearHits()
	return acsuo
}

// SetMisses sets the "misses" field.
func (acsuo *ActionCacheStatisticsUpdateOne) SetMisses(i int32) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ResetMisses()
	acsuo.mutation.SetMisses(i)
	return acsuo
}

// SetNillableMisses sets the "misses" field if the given value is not nil.
func (acsuo *ActionCacheStatisticsUpdateOne) SetNillableMisses(i *int32) *ActionCacheStatisticsUpdateOne {
	if i != nil {
		acsuo.SetMisses(*i)
	}
	return acsuo
}

// AddMisses adds i to the "misses" field.
func (acsuo *ActionCacheStatisticsUpdateOne) AddMisses(i int32) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.AddMisses(i)
	return acsuo
}

// ClearMisses clears the value of the "misses" field.
func (acsuo *ActionCacheStatisticsUpdateOne) ClearMisses() *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ClearMisses()
	return acsuo
}

// SetActionSummaryID sets the "action_summary" edge to the ActionSummary entity by ID.
func (acsuo *ActionCacheStatisticsUpdateOne) SetActionSummaryID(id int) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.SetActionSummaryID(id)
	return acsuo
}

// SetNillableActionSummaryID sets the "action_summary" edge to the ActionSummary entity by ID if the given value is not nil.
func (acsuo *ActionCacheStatisticsUpdateOne) SetNillableActionSummaryID(id *int) *ActionCacheStatisticsUpdateOne {
	if id != nil {
		acsuo = acsuo.SetActionSummaryID(*id)
	}
	return acsuo
}

// SetActionSummary sets the "action_summary" edge to the ActionSummary entity.
func (acsuo *ActionCacheStatisticsUpdateOne) SetActionSummary(a *ActionSummary) *ActionCacheStatisticsUpdateOne {
	return acsuo.SetActionSummaryID(a.ID)
}

// AddMissDetailIDs adds the "miss_details" edge to the MissDetail entity by IDs.
func (acsuo *ActionCacheStatisticsUpdateOne) AddMissDetailIDs(ids ...int) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.AddMissDetailIDs(ids...)
	return acsuo
}

// AddMissDetails adds the "miss_details" edges to the MissDetail entity.
func (acsuo *ActionCacheStatisticsUpdateOne) AddMissDetails(m ...*MissDetail) *ActionCacheStatisticsUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return acsuo.AddMissDetailIDs(ids...)
}

// Mutation returns the ActionCacheStatisticsMutation object of the builder.
func (acsuo *ActionCacheStatisticsUpdateOne) Mutation() *ActionCacheStatisticsMutation {
	return acsuo.mutation
}

// ClearActionSummary clears the "action_summary" edge to the ActionSummary entity.
func (acsuo *ActionCacheStatisticsUpdateOne) ClearActionSummary() *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ClearActionSummary()
	return acsuo
}

// ClearMissDetails clears all "miss_details" edges to the MissDetail entity.
func (acsuo *ActionCacheStatisticsUpdateOne) ClearMissDetails() *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.ClearMissDetails()
	return acsuo
}

// RemoveMissDetailIDs removes the "miss_details" edge to MissDetail entities by IDs.
func (acsuo *ActionCacheStatisticsUpdateOne) RemoveMissDetailIDs(ids ...int) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.RemoveMissDetailIDs(ids...)
	return acsuo
}

// RemoveMissDetails removes "miss_details" edges to MissDetail entities.
func (acsuo *ActionCacheStatisticsUpdateOne) RemoveMissDetails(m ...*MissDetail) *ActionCacheStatisticsUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return acsuo.RemoveMissDetailIDs(ids...)
}

// Where appends a list predicates to the ActionCacheStatisticsUpdate builder.
func (acsuo *ActionCacheStatisticsUpdateOne) Where(ps ...predicate.ActionCacheStatistics) *ActionCacheStatisticsUpdateOne {
	acsuo.mutation.Where(ps...)
	return acsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acsuo *ActionCacheStatisticsUpdateOne) Select(field string, fields ...string) *ActionCacheStatisticsUpdateOne {
	acsuo.fields = append([]string{field}, fields...)
	return acsuo
}

// Save executes the query and returns the updated ActionCacheStatistics entity.
func (acsuo *ActionCacheStatisticsUpdateOne) Save(ctx context.Context) (*ActionCacheStatistics, error) {
	return withHooks(ctx, acsuo.sqlSave, acsuo.mutation, acsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acsuo *ActionCacheStatisticsUpdateOne) SaveX(ctx context.Context) *ActionCacheStatistics {
	node, err := acsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acsuo *ActionCacheStatisticsUpdateOne) Exec(ctx context.Context) error {
	_, err := acsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acsuo *ActionCacheStatisticsUpdateOne) ExecX(ctx context.Context) {
	if err := acsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (acsuo *ActionCacheStatisticsUpdateOne) sqlSave(ctx context.Context) (_node *ActionCacheStatistics, err error) {
	_spec := sqlgraph.NewUpdateSpec(actioncachestatistics.Table, actioncachestatistics.Columns, sqlgraph.NewFieldSpec(actioncachestatistics.FieldID, field.TypeInt))
	id, ok := acsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ActionCacheStatistics.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, actioncachestatistics.FieldID)
		for _, f := range fields {
			if !actioncachestatistics.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != actioncachestatistics.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acsuo.mutation.SizeInBytes(); ok {
		_spec.SetField(actioncachestatistics.FieldSizeInBytes, field.TypeUint64, value)
	}
	if value, ok := acsuo.mutation.AddedSizeInBytes(); ok {
		_spec.AddField(actioncachestatistics.FieldSizeInBytes, field.TypeUint64, value)
	}
	if acsuo.mutation.SizeInBytesCleared() {
		_spec.ClearField(actioncachestatistics.FieldSizeInBytes, field.TypeUint64)
	}
	if value, ok := acsuo.mutation.SaveTimeInMs(); ok {
		_spec.SetField(actioncachestatistics.FieldSaveTimeInMs, field.TypeUint64, value)
	}
	if value, ok := acsuo.mutation.AddedSaveTimeInMs(); ok {
		_spec.AddField(actioncachestatistics.FieldSaveTimeInMs, field.TypeUint64, value)
	}
	if acsuo.mutation.SaveTimeInMsCleared() {
		_spec.ClearField(actioncachestatistics.FieldSaveTimeInMs, field.TypeUint64)
	}
	if value, ok := acsuo.mutation.LoadTimeInMs(); ok {
		_spec.SetField(actioncachestatistics.FieldLoadTimeInMs, field.TypeInt64, value)
	}
	if value, ok := acsuo.mutation.AddedLoadTimeInMs(); ok {
		_spec.AddField(actioncachestatistics.FieldLoadTimeInMs, field.TypeInt64, value)
	}
	if acsuo.mutation.LoadTimeInMsCleared() {
		_spec.ClearField(actioncachestatistics.FieldLoadTimeInMs, field.TypeInt64)
	}
	if value, ok := acsuo.mutation.Hits(); ok {
		_spec.SetField(actioncachestatistics.FieldHits, field.TypeInt32, value)
	}
	if value, ok := acsuo.mutation.AddedHits(); ok {
		_spec.AddField(actioncachestatistics.FieldHits, field.TypeInt32, value)
	}
	if acsuo.mutation.HitsCleared() {
		_spec.ClearField(actioncachestatistics.FieldHits, field.TypeInt32)
	}
	if value, ok := acsuo.mutation.Misses(); ok {
		_spec.SetField(actioncachestatistics.FieldMisses, field.TypeInt32, value)
	}
	if value, ok := acsuo.mutation.AddedMisses(); ok {
		_spec.AddField(actioncachestatistics.FieldMisses, field.TypeInt32, value)
	}
	if acsuo.mutation.MissesCleared() {
		_spec.ClearField(actioncachestatistics.FieldMisses, field.TypeInt32)
	}
	if acsuo.mutation.ActionSummaryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   actioncachestatistics.ActionSummaryTable,
			Columns: []string{actioncachestatistics.ActionSummaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(actionsummary.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acsuo.mutation.ActionSummaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   actioncachestatistics.ActionSummaryTable,
			Columns: []string{actioncachestatistics.ActionSummaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(actionsummary.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if acsuo.mutation.MissDetailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   actioncachestatistics.MissDetailsTable,
			Columns: []string{actioncachestatistics.MissDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missdetail.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acsuo.mutation.RemovedMissDetailsIDs(); len(nodes) > 0 && !acsuo.mutation.MissDetailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   actioncachestatistics.MissDetailsTable,
			Columns: []string{actioncachestatistics.MissDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := acsuo.mutation.MissDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   actioncachestatistics.MissDetailsTable,
			Columns: []string{actioncachestatistics.MissDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(missdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ActionCacheStatistics{config: acsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{actioncachestatistics.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	acsuo.mutation.done = true
	return _node, nil
}
