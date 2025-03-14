// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/bazelinvocation"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetcomplete"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetconfigured"
	"github.com/buildbarn/bb-portal/ent/gen/ent/targetpair"
)

// TargetPairUpdate is the builder for updating TargetPair entities.
type TargetPairUpdate struct {
	config
	hooks    []Hook
	mutation *TargetPairMutation
}

// Where appends a list predicates to the TargetPairUpdate builder.
func (tpu *TargetPairUpdate) Where(ps ...predicate.TargetPair) *TargetPairUpdate {
	tpu.mutation.Where(ps...)
	return tpu
}

// SetLabel sets the "label" field.
func (tpu *TargetPairUpdate) SetLabel(s string) *TargetPairUpdate {
	tpu.mutation.SetLabel(s)
	return tpu
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (tpu *TargetPairUpdate) SetNillableLabel(s *string) *TargetPairUpdate {
	if s != nil {
		tpu.SetLabel(*s)
	}
	return tpu
}

// ClearLabel clears the value of the "label" field.
func (tpu *TargetPairUpdate) ClearLabel() *TargetPairUpdate {
	tpu.mutation.ClearLabel()
	return tpu
}

// SetDurationInMs sets the "duration_in_ms" field.
func (tpu *TargetPairUpdate) SetDurationInMs(i int64) *TargetPairUpdate {
	tpu.mutation.ResetDurationInMs()
	tpu.mutation.SetDurationInMs(i)
	return tpu
}

// SetNillableDurationInMs sets the "duration_in_ms" field if the given value is not nil.
func (tpu *TargetPairUpdate) SetNillableDurationInMs(i *int64) *TargetPairUpdate {
	if i != nil {
		tpu.SetDurationInMs(*i)
	}
	return tpu
}

// AddDurationInMs adds i to the "duration_in_ms" field.
func (tpu *TargetPairUpdate) AddDurationInMs(i int64) *TargetPairUpdate {
	tpu.mutation.AddDurationInMs(i)
	return tpu
}

// ClearDurationInMs clears the value of the "duration_in_ms" field.
func (tpu *TargetPairUpdate) ClearDurationInMs() *TargetPairUpdate {
	tpu.mutation.ClearDurationInMs()
	return tpu
}

// SetSuccess sets the "success" field.
func (tpu *TargetPairUpdate) SetSuccess(b bool) *TargetPairUpdate {
	tpu.mutation.SetSuccess(b)
	return tpu
}

// SetNillableSuccess sets the "success" field if the given value is not nil.
func (tpu *TargetPairUpdate) SetNillableSuccess(b *bool) *TargetPairUpdate {
	if b != nil {
		tpu.SetSuccess(*b)
	}
	return tpu
}

// ClearSuccess clears the value of the "success" field.
func (tpu *TargetPairUpdate) ClearSuccess() *TargetPairUpdate {
	tpu.mutation.ClearSuccess()
	return tpu
}

// SetTargetKind sets the "target_kind" field.
func (tpu *TargetPairUpdate) SetTargetKind(s string) *TargetPairUpdate {
	tpu.mutation.SetTargetKind(s)
	return tpu
}

// SetNillableTargetKind sets the "target_kind" field if the given value is not nil.
func (tpu *TargetPairUpdate) SetNillableTargetKind(s *string) *TargetPairUpdate {
	if s != nil {
		tpu.SetTargetKind(*s)
	}
	return tpu
}

// ClearTargetKind clears the value of the "target_kind" field.
func (tpu *TargetPairUpdate) ClearTargetKind() *TargetPairUpdate {
	tpu.mutation.ClearTargetKind()
	return tpu
}

// SetTestSize sets the "test_size" field.
func (tpu *TargetPairUpdate) SetTestSize(ts targetpair.TestSize) *TargetPairUpdate {
	tpu.mutation.SetTestSize(ts)
	return tpu
}

// SetNillableTestSize sets the "test_size" field if the given value is not nil.
func (tpu *TargetPairUpdate) SetNillableTestSize(ts *targetpair.TestSize) *TargetPairUpdate {
	if ts != nil {
		tpu.SetTestSize(*ts)
	}
	return tpu
}

// ClearTestSize clears the value of the "test_size" field.
func (tpu *TargetPairUpdate) ClearTestSize() *TargetPairUpdate {
	tpu.mutation.ClearTestSize()
	return tpu
}

// SetAbortReason sets the "abort_reason" field.
func (tpu *TargetPairUpdate) SetAbortReason(tr targetpair.AbortReason) *TargetPairUpdate {
	tpu.mutation.SetAbortReason(tr)
	return tpu
}

// SetNillableAbortReason sets the "abort_reason" field if the given value is not nil.
func (tpu *TargetPairUpdate) SetNillableAbortReason(tr *targetpair.AbortReason) *TargetPairUpdate {
	if tr != nil {
		tpu.SetAbortReason(*tr)
	}
	return tpu
}

// ClearAbortReason clears the value of the "abort_reason" field.
func (tpu *TargetPairUpdate) ClearAbortReason() *TargetPairUpdate {
	tpu.mutation.ClearAbortReason()
	return tpu
}

// SetBazelInvocationID sets the "bazel_invocation" edge to the BazelInvocation entity by ID.
func (tpu *TargetPairUpdate) SetBazelInvocationID(id int) *TargetPairUpdate {
	tpu.mutation.SetBazelInvocationID(id)
	return tpu
}

// SetNillableBazelInvocationID sets the "bazel_invocation" edge to the BazelInvocation entity by ID if the given value is not nil.
func (tpu *TargetPairUpdate) SetNillableBazelInvocationID(id *int) *TargetPairUpdate {
	if id != nil {
		tpu = tpu.SetBazelInvocationID(*id)
	}
	return tpu
}

// SetBazelInvocation sets the "bazel_invocation" edge to the BazelInvocation entity.
func (tpu *TargetPairUpdate) SetBazelInvocation(b *BazelInvocation) *TargetPairUpdate {
	return tpu.SetBazelInvocationID(b.ID)
}

// SetConfigurationID sets the "configuration" edge to the TargetConfigured entity by ID.
func (tpu *TargetPairUpdate) SetConfigurationID(id int) *TargetPairUpdate {
	tpu.mutation.SetConfigurationID(id)
	return tpu
}

// SetNillableConfigurationID sets the "configuration" edge to the TargetConfigured entity by ID if the given value is not nil.
func (tpu *TargetPairUpdate) SetNillableConfigurationID(id *int) *TargetPairUpdate {
	if id != nil {
		tpu = tpu.SetConfigurationID(*id)
	}
	return tpu
}

// SetConfiguration sets the "configuration" edge to the TargetConfigured entity.
func (tpu *TargetPairUpdate) SetConfiguration(t *TargetConfigured) *TargetPairUpdate {
	return tpu.SetConfigurationID(t.ID)
}

// SetCompletionID sets the "completion" edge to the TargetComplete entity by ID.
func (tpu *TargetPairUpdate) SetCompletionID(id int) *TargetPairUpdate {
	tpu.mutation.SetCompletionID(id)
	return tpu
}

// SetNillableCompletionID sets the "completion" edge to the TargetComplete entity by ID if the given value is not nil.
func (tpu *TargetPairUpdate) SetNillableCompletionID(id *int) *TargetPairUpdate {
	if id != nil {
		tpu = tpu.SetCompletionID(*id)
	}
	return tpu
}

// SetCompletion sets the "completion" edge to the TargetComplete entity.
func (tpu *TargetPairUpdate) SetCompletion(t *TargetComplete) *TargetPairUpdate {
	return tpu.SetCompletionID(t.ID)
}

// Mutation returns the TargetPairMutation object of the builder.
func (tpu *TargetPairUpdate) Mutation() *TargetPairMutation {
	return tpu.mutation
}

// ClearBazelInvocation clears the "bazel_invocation" edge to the BazelInvocation entity.
func (tpu *TargetPairUpdate) ClearBazelInvocation() *TargetPairUpdate {
	tpu.mutation.ClearBazelInvocation()
	return tpu
}

// ClearConfiguration clears the "configuration" edge to the TargetConfigured entity.
func (tpu *TargetPairUpdate) ClearConfiguration() *TargetPairUpdate {
	tpu.mutation.ClearConfiguration()
	return tpu
}

// ClearCompletion clears the "completion" edge to the TargetComplete entity.
func (tpu *TargetPairUpdate) ClearCompletion() *TargetPairUpdate {
	tpu.mutation.ClearCompletion()
	return tpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tpu *TargetPairUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tpu.sqlSave, tpu.mutation, tpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tpu *TargetPairUpdate) SaveX(ctx context.Context) int {
	affected, err := tpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tpu *TargetPairUpdate) Exec(ctx context.Context) error {
	_, err := tpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpu *TargetPairUpdate) ExecX(ctx context.Context) {
	if err := tpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpu *TargetPairUpdate) check() error {
	if v, ok := tpu.mutation.TestSize(); ok {
		if err := targetpair.TestSizeValidator(v); err != nil {
			return &ValidationError{Name: "test_size", err: fmt.Errorf(`ent: validator failed for field "TargetPair.test_size": %w`, err)}
		}
	}
	if v, ok := tpu.mutation.AbortReason(); ok {
		if err := targetpair.AbortReasonValidator(v); err != nil {
			return &ValidationError{Name: "abort_reason", err: fmt.Errorf(`ent: validator failed for field "TargetPair.abort_reason": %w`, err)}
		}
	}
	return nil
}

func (tpu *TargetPairUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(targetpair.Table, targetpair.Columns, sqlgraph.NewFieldSpec(targetpair.FieldID, field.TypeInt))
	if ps := tpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpu.mutation.Label(); ok {
		_spec.SetField(targetpair.FieldLabel, field.TypeString, value)
	}
	if tpu.mutation.LabelCleared() {
		_spec.ClearField(targetpair.FieldLabel, field.TypeString)
	}
	if value, ok := tpu.mutation.DurationInMs(); ok {
		_spec.SetField(targetpair.FieldDurationInMs, field.TypeInt64, value)
	}
	if value, ok := tpu.mutation.AddedDurationInMs(); ok {
		_spec.AddField(targetpair.FieldDurationInMs, field.TypeInt64, value)
	}
	if tpu.mutation.DurationInMsCleared() {
		_spec.ClearField(targetpair.FieldDurationInMs, field.TypeInt64)
	}
	if value, ok := tpu.mutation.Success(); ok {
		_spec.SetField(targetpair.FieldSuccess, field.TypeBool, value)
	}
	if tpu.mutation.SuccessCleared() {
		_spec.ClearField(targetpair.FieldSuccess, field.TypeBool)
	}
	if value, ok := tpu.mutation.TargetKind(); ok {
		_spec.SetField(targetpair.FieldTargetKind, field.TypeString, value)
	}
	if tpu.mutation.TargetKindCleared() {
		_spec.ClearField(targetpair.FieldTargetKind, field.TypeString)
	}
	if value, ok := tpu.mutation.TestSize(); ok {
		_spec.SetField(targetpair.FieldTestSize, field.TypeEnum, value)
	}
	if tpu.mutation.TestSizeCleared() {
		_spec.ClearField(targetpair.FieldTestSize, field.TypeEnum)
	}
	if value, ok := tpu.mutation.AbortReason(); ok {
		_spec.SetField(targetpair.FieldAbortReason, field.TypeEnum, value)
	}
	if tpu.mutation.AbortReasonCleared() {
		_spec.ClearField(targetpair.FieldAbortReason, field.TypeEnum)
	}
	if tpu.mutation.BazelInvocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   targetpair.BazelInvocationTable,
			Columns: []string{targetpair.BazelInvocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bazelinvocation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpu.mutation.BazelInvocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   targetpair.BazelInvocationTable,
			Columns: []string{targetpair.BazelInvocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bazelinvocation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tpu.mutation.ConfigurationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   targetpair.ConfigurationTable,
			Columns: []string{targetpair.ConfigurationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(targetconfigured.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpu.mutation.ConfigurationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   targetpair.ConfigurationTable,
			Columns: []string{targetpair.ConfigurationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(targetconfigured.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tpu.mutation.CompletionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   targetpair.CompletionTable,
			Columns: []string{targetpair.CompletionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(targetcomplete.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpu.mutation.CompletionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   targetpair.CompletionTable,
			Columns: []string{targetpair.CompletionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(targetcomplete.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{targetpair.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tpu.mutation.done = true
	return n, nil
}

// TargetPairUpdateOne is the builder for updating a single TargetPair entity.
type TargetPairUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TargetPairMutation
}

// SetLabel sets the "label" field.
func (tpuo *TargetPairUpdateOne) SetLabel(s string) *TargetPairUpdateOne {
	tpuo.mutation.SetLabel(s)
	return tpuo
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (tpuo *TargetPairUpdateOne) SetNillableLabel(s *string) *TargetPairUpdateOne {
	if s != nil {
		tpuo.SetLabel(*s)
	}
	return tpuo
}

// ClearLabel clears the value of the "label" field.
func (tpuo *TargetPairUpdateOne) ClearLabel() *TargetPairUpdateOne {
	tpuo.mutation.ClearLabel()
	return tpuo
}

// SetDurationInMs sets the "duration_in_ms" field.
func (tpuo *TargetPairUpdateOne) SetDurationInMs(i int64) *TargetPairUpdateOne {
	tpuo.mutation.ResetDurationInMs()
	tpuo.mutation.SetDurationInMs(i)
	return tpuo
}

// SetNillableDurationInMs sets the "duration_in_ms" field if the given value is not nil.
func (tpuo *TargetPairUpdateOne) SetNillableDurationInMs(i *int64) *TargetPairUpdateOne {
	if i != nil {
		tpuo.SetDurationInMs(*i)
	}
	return tpuo
}

// AddDurationInMs adds i to the "duration_in_ms" field.
func (tpuo *TargetPairUpdateOne) AddDurationInMs(i int64) *TargetPairUpdateOne {
	tpuo.mutation.AddDurationInMs(i)
	return tpuo
}

// ClearDurationInMs clears the value of the "duration_in_ms" field.
func (tpuo *TargetPairUpdateOne) ClearDurationInMs() *TargetPairUpdateOne {
	tpuo.mutation.ClearDurationInMs()
	return tpuo
}

// SetSuccess sets the "success" field.
func (tpuo *TargetPairUpdateOne) SetSuccess(b bool) *TargetPairUpdateOne {
	tpuo.mutation.SetSuccess(b)
	return tpuo
}

// SetNillableSuccess sets the "success" field if the given value is not nil.
func (tpuo *TargetPairUpdateOne) SetNillableSuccess(b *bool) *TargetPairUpdateOne {
	if b != nil {
		tpuo.SetSuccess(*b)
	}
	return tpuo
}

// ClearSuccess clears the value of the "success" field.
func (tpuo *TargetPairUpdateOne) ClearSuccess() *TargetPairUpdateOne {
	tpuo.mutation.ClearSuccess()
	return tpuo
}

// SetTargetKind sets the "target_kind" field.
func (tpuo *TargetPairUpdateOne) SetTargetKind(s string) *TargetPairUpdateOne {
	tpuo.mutation.SetTargetKind(s)
	return tpuo
}

// SetNillableTargetKind sets the "target_kind" field if the given value is not nil.
func (tpuo *TargetPairUpdateOne) SetNillableTargetKind(s *string) *TargetPairUpdateOne {
	if s != nil {
		tpuo.SetTargetKind(*s)
	}
	return tpuo
}

// ClearTargetKind clears the value of the "target_kind" field.
func (tpuo *TargetPairUpdateOne) ClearTargetKind() *TargetPairUpdateOne {
	tpuo.mutation.ClearTargetKind()
	return tpuo
}

// SetTestSize sets the "test_size" field.
func (tpuo *TargetPairUpdateOne) SetTestSize(ts targetpair.TestSize) *TargetPairUpdateOne {
	tpuo.mutation.SetTestSize(ts)
	return tpuo
}

// SetNillableTestSize sets the "test_size" field if the given value is not nil.
func (tpuo *TargetPairUpdateOne) SetNillableTestSize(ts *targetpair.TestSize) *TargetPairUpdateOne {
	if ts != nil {
		tpuo.SetTestSize(*ts)
	}
	return tpuo
}

// ClearTestSize clears the value of the "test_size" field.
func (tpuo *TargetPairUpdateOne) ClearTestSize() *TargetPairUpdateOne {
	tpuo.mutation.ClearTestSize()
	return tpuo
}

// SetAbortReason sets the "abort_reason" field.
func (tpuo *TargetPairUpdateOne) SetAbortReason(tr targetpair.AbortReason) *TargetPairUpdateOne {
	tpuo.mutation.SetAbortReason(tr)
	return tpuo
}

// SetNillableAbortReason sets the "abort_reason" field if the given value is not nil.
func (tpuo *TargetPairUpdateOne) SetNillableAbortReason(tr *targetpair.AbortReason) *TargetPairUpdateOne {
	if tr != nil {
		tpuo.SetAbortReason(*tr)
	}
	return tpuo
}

// ClearAbortReason clears the value of the "abort_reason" field.
func (tpuo *TargetPairUpdateOne) ClearAbortReason() *TargetPairUpdateOne {
	tpuo.mutation.ClearAbortReason()
	return tpuo
}

// SetBazelInvocationID sets the "bazel_invocation" edge to the BazelInvocation entity by ID.
func (tpuo *TargetPairUpdateOne) SetBazelInvocationID(id int) *TargetPairUpdateOne {
	tpuo.mutation.SetBazelInvocationID(id)
	return tpuo
}

// SetNillableBazelInvocationID sets the "bazel_invocation" edge to the BazelInvocation entity by ID if the given value is not nil.
func (tpuo *TargetPairUpdateOne) SetNillableBazelInvocationID(id *int) *TargetPairUpdateOne {
	if id != nil {
		tpuo = tpuo.SetBazelInvocationID(*id)
	}
	return tpuo
}

// SetBazelInvocation sets the "bazel_invocation" edge to the BazelInvocation entity.
func (tpuo *TargetPairUpdateOne) SetBazelInvocation(b *BazelInvocation) *TargetPairUpdateOne {
	return tpuo.SetBazelInvocationID(b.ID)
}

// SetConfigurationID sets the "configuration" edge to the TargetConfigured entity by ID.
func (tpuo *TargetPairUpdateOne) SetConfigurationID(id int) *TargetPairUpdateOne {
	tpuo.mutation.SetConfigurationID(id)
	return tpuo
}

// SetNillableConfigurationID sets the "configuration" edge to the TargetConfigured entity by ID if the given value is not nil.
func (tpuo *TargetPairUpdateOne) SetNillableConfigurationID(id *int) *TargetPairUpdateOne {
	if id != nil {
		tpuo = tpuo.SetConfigurationID(*id)
	}
	return tpuo
}

// SetConfiguration sets the "configuration" edge to the TargetConfigured entity.
func (tpuo *TargetPairUpdateOne) SetConfiguration(t *TargetConfigured) *TargetPairUpdateOne {
	return tpuo.SetConfigurationID(t.ID)
}

// SetCompletionID sets the "completion" edge to the TargetComplete entity by ID.
func (tpuo *TargetPairUpdateOne) SetCompletionID(id int) *TargetPairUpdateOne {
	tpuo.mutation.SetCompletionID(id)
	return tpuo
}

// SetNillableCompletionID sets the "completion" edge to the TargetComplete entity by ID if the given value is not nil.
func (tpuo *TargetPairUpdateOne) SetNillableCompletionID(id *int) *TargetPairUpdateOne {
	if id != nil {
		tpuo = tpuo.SetCompletionID(*id)
	}
	return tpuo
}

// SetCompletion sets the "completion" edge to the TargetComplete entity.
func (tpuo *TargetPairUpdateOne) SetCompletion(t *TargetComplete) *TargetPairUpdateOne {
	return tpuo.SetCompletionID(t.ID)
}

// Mutation returns the TargetPairMutation object of the builder.
func (tpuo *TargetPairUpdateOne) Mutation() *TargetPairMutation {
	return tpuo.mutation
}

// ClearBazelInvocation clears the "bazel_invocation" edge to the BazelInvocation entity.
func (tpuo *TargetPairUpdateOne) ClearBazelInvocation() *TargetPairUpdateOne {
	tpuo.mutation.ClearBazelInvocation()
	return tpuo
}

// ClearConfiguration clears the "configuration" edge to the TargetConfigured entity.
func (tpuo *TargetPairUpdateOne) ClearConfiguration() *TargetPairUpdateOne {
	tpuo.mutation.ClearConfiguration()
	return tpuo
}

// ClearCompletion clears the "completion" edge to the TargetComplete entity.
func (tpuo *TargetPairUpdateOne) ClearCompletion() *TargetPairUpdateOne {
	tpuo.mutation.ClearCompletion()
	return tpuo
}

// Where appends a list predicates to the TargetPairUpdate builder.
func (tpuo *TargetPairUpdateOne) Where(ps ...predicate.TargetPair) *TargetPairUpdateOne {
	tpuo.mutation.Where(ps...)
	return tpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tpuo *TargetPairUpdateOne) Select(field string, fields ...string) *TargetPairUpdateOne {
	tpuo.fields = append([]string{field}, fields...)
	return tpuo
}

// Save executes the query and returns the updated TargetPair entity.
func (tpuo *TargetPairUpdateOne) Save(ctx context.Context) (*TargetPair, error) {
	return withHooks(ctx, tpuo.sqlSave, tpuo.mutation, tpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tpuo *TargetPairUpdateOne) SaveX(ctx context.Context) *TargetPair {
	node, err := tpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tpuo *TargetPairUpdateOne) Exec(ctx context.Context) error {
	_, err := tpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tpuo *TargetPairUpdateOne) ExecX(ctx context.Context) {
	if err := tpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpuo *TargetPairUpdateOne) check() error {
	if v, ok := tpuo.mutation.TestSize(); ok {
		if err := targetpair.TestSizeValidator(v); err != nil {
			return &ValidationError{Name: "test_size", err: fmt.Errorf(`ent: validator failed for field "TargetPair.test_size": %w`, err)}
		}
	}
	if v, ok := tpuo.mutation.AbortReason(); ok {
		if err := targetpair.AbortReasonValidator(v); err != nil {
			return &ValidationError{Name: "abort_reason", err: fmt.Errorf(`ent: validator failed for field "TargetPair.abort_reason": %w`, err)}
		}
	}
	return nil
}

func (tpuo *TargetPairUpdateOne) sqlSave(ctx context.Context) (_node *TargetPair, err error) {
	if err := tpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(targetpair.Table, targetpair.Columns, sqlgraph.NewFieldSpec(targetpair.FieldID, field.TypeInt))
	id, ok := tpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TargetPair.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, targetpair.FieldID)
		for _, f := range fields {
			if !targetpair.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != targetpair.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tpuo.mutation.Label(); ok {
		_spec.SetField(targetpair.FieldLabel, field.TypeString, value)
	}
	if tpuo.mutation.LabelCleared() {
		_spec.ClearField(targetpair.FieldLabel, field.TypeString)
	}
	if value, ok := tpuo.mutation.DurationInMs(); ok {
		_spec.SetField(targetpair.FieldDurationInMs, field.TypeInt64, value)
	}
	if value, ok := tpuo.mutation.AddedDurationInMs(); ok {
		_spec.AddField(targetpair.FieldDurationInMs, field.TypeInt64, value)
	}
	if tpuo.mutation.DurationInMsCleared() {
		_spec.ClearField(targetpair.FieldDurationInMs, field.TypeInt64)
	}
	if value, ok := tpuo.mutation.Success(); ok {
		_spec.SetField(targetpair.FieldSuccess, field.TypeBool, value)
	}
	if tpuo.mutation.SuccessCleared() {
		_spec.ClearField(targetpair.FieldSuccess, field.TypeBool)
	}
	if value, ok := tpuo.mutation.TargetKind(); ok {
		_spec.SetField(targetpair.FieldTargetKind, field.TypeString, value)
	}
	if tpuo.mutation.TargetKindCleared() {
		_spec.ClearField(targetpair.FieldTargetKind, field.TypeString)
	}
	if value, ok := tpuo.mutation.TestSize(); ok {
		_spec.SetField(targetpair.FieldTestSize, field.TypeEnum, value)
	}
	if tpuo.mutation.TestSizeCleared() {
		_spec.ClearField(targetpair.FieldTestSize, field.TypeEnum)
	}
	if value, ok := tpuo.mutation.AbortReason(); ok {
		_spec.SetField(targetpair.FieldAbortReason, field.TypeEnum, value)
	}
	if tpuo.mutation.AbortReasonCleared() {
		_spec.ClearField(targetpair.FieldAbortReason, field.TypeEnum)
	}
	if tpuo.mutation.BazelInvocationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   targetpair.BazelInvocationTable,
			Columns: []string{targetpair.BazelInvocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bazelinvocation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpuo.mutation.BazelInvocationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   targetpair.BazelInvocationTable,
			Columns: []string{targetpair.BazelInvocationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bazelinvocation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tpuo.mutation.ConfigurationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   targetpair.ConfigurationTable,
			Columns: []string{targetpair.ConfigurationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(targetconfigured.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpuo.mutation.ConfigurationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   targetpair.ConfigurationTable,
			Columns: []string{targetpair.ConfigurationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(targetconfigured.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tpuo.mutation.CompletionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   targetpair.CompletionTable,
			Columns: []string{targetpair.CompletionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(targetcomplete.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tpuo.mutation.CompletionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   targetpair.CompletionTable,
			Columns: []string{targetpair.CompletionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(targetcomplete.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TargetPair{config: tpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{targetpair.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tpuo.mutation.done = true
	return _node, nil
}
