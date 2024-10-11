// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/exectioninfo"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testcollection"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testfile"
	"github.com/buildbarn/bb-portal/ent/gen/ent/testresultbes"
)

// TestResultBESUpdate is the builder for updating TestResultBES entities.
type TestResultBESUpdate struct {
	config
	hooks    []Hook
	mutation *TestResultBESMutation
}

// Where appends a list predicates to the TestResultBESUpdate builder.
func (trbu *TestResultBESUpdate) Where(ps ...predicate.TestResultBES) *TestResultBESUpdate {
	trbu.mutation.Where(ps...)
	return trbu
}

// SetTestStatus sets the "test_status" field.
func (trbu *TestResultBESUpdate) SetTestStatus(ts testresultbes.TestStatus) *TestResultBESUpdate {
	trbu.mutation.SetTestStatus(ts)
	return trbu
}

// SetNillableTestStatus sets the "test_status" field if the given value is not nil.
func (trbu *TestResultBESUpdate) SetNillableTestStatus(ts *testresultbes.TestStatus) *TestResultBESUpdate {
	if ts != nil {
		trbu.SetTestStatus(*ts)
	}
	return trbu
}

// ClearTestStatus clears the value of the "test_status" field.
func (trbu *TestResultBESUpdate) ClearTestStatus() *TestResultBESUpdate {
	trbu.mutation.ClearTestStatus()
	return trbu
}

// SetStatusDetails sets the "status_details" field.
func (trbu *TestResultBESUpdate) SetStatusDetails(s string) *TestResultBESUpdate {
	trbu.mutation.SetStatusDetails(s)
	return trbu
}

// SetNillableStatusDetails sets the "status_details" field if the given value is not nil.
func (trbu *TestResultBESUpdate) SetNillableStatusDetails(s *string) *TestResultBESUpdate {
	if s != nil {
		trbu.SetStatusDetails(*s)
	}
	return trbu
}

// ClearStatusDetails clears the value of the "status_details" field.
func (trbu *TestResultBESUpdate) ClearStatusDetails() *TestResultBESUpdate {
	trbu.mutation.ClearStatusDetails()
	return trbu
}

// SetLabel sets the "label" field.
func (trbu *TestResultBESUpdate) SetLabel(s string) *TestResultBESUpdate {
	trbu.mutation.SetLabel(s)
	return trbu
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (trbu *TestResultBESUpdate) SetNillableLabel(s *string) *TestResultBESUpdate {
	if s != nil {
		trbu.SetLabel(*s)
	}
	return trbu
}

// ClearLabel clears the value of the "label" field.
func (trbu *TestResultBESUpdate) ClearLabel() *TestResultBESUpdate {
	trbu.mutation.ClearLabel()
	return trbu
}

// SetWarning sets the "warning" field.
func (trbu *TestResultBESUpdate) SetWarning(s []string) *TestResultBESUpdate {
	trbu.mutation.SetWarning(s)
	return trbu
}

// AppendWarning appends s to the "warning" field.
func (trbu *TestResultBESUpdate) AppendWarning(s []string) *TestResultBESUpdate {
	trbu.mutation.AppendWarning(s)
	return trbu
}

// ClearWarning clears the value of the "warning" field.
func (trbu *TestResultBESUpdate) ClearWarning() *TestResultBESUpdate {
	trbu.mutation.ClearWarning()
	return trbu
}

// SetCachedLocally sets the "cached_locally" field.
func (trbu *TestResultBESUpdate) SetCachedLocally(b bool) *TestResultBESUpdate {
	trbu.mutation.SetCachedLocally(b)
	return trbu
}

// SetNillableCachedLocally sets the "cached_locally" field if the given value is not nil.
func (trbu *TestResultBESUpdate) SetNillableCachedLocally(b *bool) *TestResultBESUpdate {
	if b != nil {
		trbu.SetCachedLocally(*b)
	}
	return trbu
}

// ClearCachedLocally clears the value of the "cached_locally" field.
func (trbu *TestResultBESUpdate) ClearCachedLocally() *TestResultBESUpdate {
	trbu.mutation.ClearCachedLocally()
	return trbu
}

// SetTestAttemptStartMillisEpoch sets the "test_attempt_start_millis_epoch" field.
func (trbu *TestResultBESUpdate) SetTestAttemptStartMillisEpoch(i int64) *TestResultBESUpdate {
	trbu.mutation.ResetTestAttemptStartMillisEpoch()
	trbu.mutation.SetTestAttemptStartMillisEpoch(i)
	return trbu
}

// SetNillableTestAttemptStartMillisEpoch sets the "test_attempt_start_millis_epoch" field if the given value is not nil.
func (trbu *TestResultBESUpdate) SetNillableTestAttemptStartMillisEpoch(i *int64) *TestResultBESUpdate {
	if i != nil {
		trbu.SetTestAttemptStartMillisEpoch(*i)
	}
	return trbu
}

// AddTestAttemptStartMillisEpoch adds i to the "test_attempt_start_millis_epoch" field.
func (trbu *TestResultBESUpdate) AddTestAttemptStartMillisEpoch(i int64) *TestResultBESUpdate {
	trbu.mutation.AddTestAttemptStartMillisEpoch(i)
	return trbu
}

// ClearTestAttemptStartMillisEpoch clears the value of the "test_attempt_start_millis_epoch" field.
func (trbu *TestResultBESUpdate) ClearTestAttemptStartMillisEpoch() *TestResultBESUpdate {
	trbu.mutation.ClearTestAttemptStartMillisEpoch()
	return trbu
}

// SetTestAttemptStart sets the "test_attempt_start" field.
func (trbu *TestResultBESUpdate) SetTestAttemptStart(s string) *TestResultBESUpdate {
	trbu.mutation.SetTestAttemptStart(s)
	return trbu
}

// SetNillableTestAttemptStart sets the "test_attempt_start" field if the given value is not nil.
func (trbu *TestResultBESUpdate) SetNillableTestAttemptStart(s *string) *TestResultBESUpdate {
	if s != nil {
		trbu.SetTestAttemptStart(*s)
	}
	return trbu
}

// ClearTestAttemptStart clears the value of the "test_attempt_start" field.
func (trbu *TestResultBESUpdate) ClearTestAttemptStart() *TestResultBESUpdate {
	trbu.mutation.ClearTestAttemptStart()
	return trbu
}

// SetTestAttemptDurationMillis sets the "test_attempt_duration_millis" field.
func (trbu *TestResultBESUpdate) SetTestAttemptDurationMillis(i int64) *TestResultBESUpdate {
	trbu.mutation.ResetTestAttemptDurationMillis()
	trbu.mutation.SetTestAttemptDurationMillis(i)
	return trbu
}

// SetNillableTestAttemptDurationMillis sets the "test_attempt_duration_millis" field if the given value is not nil.
func (trbu *TestResultBESUpdate) SetNillableTestAttemptDurationMillis(i *int64) *TestResultBESUpdate {
	if i != nil {
		trbu.SetTestAttemptDurationMillis(*i)
	}
	return trbu
}

// AddTestAttemptDurationMillis adds i to the "test_attempt_duration_millis" field.
func (trbu *TestResultBESUpdate) AddTestAttemptDurationMillis(i int64) *TestResultBESUpdate {
	trbu.mutation.AddTestAttemptDurationMillis(i)
	return trbu
}

// ClearTestAttemptDurationMillis clears the value of the "test_attempt_duration_millis" field.
func (trbu *TestResultBESUpdate) ClearTestAttemptDurationMillis() *TestResultBESUpdate {
	trbu.mutation.ClearTestAttemptDurationMillis()
	return trbu
}

// SetTestAttemptDuration sets the "test_attempt_duration" field.
func (trbu *TestResultBESUpdate) SetTestAttemptDuration(i int64) *TestResultBESUpdate {
	trbu.mutation.ResetTestAttemptDuration()
	trbu.mutation.SetTestAttemptDuration(i)
	return trbu
}

// SetNillableTestAttemptDuration sets the "test_attempt_duration" field if the given value is not nil.
func (trbu *TestResultBESUpdate) SetNillableTestAttemptDuration(i *int64) *TestResultBESUpdate {
	if i != nil {
		trbu.SetTestAttemptDuration(*i)
	}
	return trbu
}

// AddTestAttemptDuration adds i to the "test_attempt_duration" field.
func (trbu *TestResultBESUpdate) AddTestAttemptDuration(i int64) *TestResultBESUpdate {
	trbu.mutation.AddTestAttemptDuration(i)
	return trbu
}

// ClearTestAttemptDuration clears the value of the "test_attempt_duration" field.
func (trbu *TestResultBESUpdate) ClearTestAttemptDuration() *TestResultBESUpdate {
	trbu.mutation.ClearTestAttemptDuration()
	return trbu
}

// SetTestCollectionID sets the "test_collection" edge to the TestCollection entity by ID.
func (trbu *TestResultBESUpdate) SetTestCollectionID(id int) *TestResultBESUpdate {
	trbu.mutation.SetTestCollectionID(id)
	return trbu
}

// SetNillableTestCollectionID sets the "test_collection" edge to the TestCollection entity by ID if the given value is not nil.
func (trbu *TestResultBESUpdate) SetNillableTestCollectionID(id *int) *TestResultBESUpdate {
	if id != nil {
		trbu = trbu.SetTestCollectionID(*id)
	}
	return trbu
}

// SetTestCollection sets the "test_collection" edge to the TestCollection entity.
func (trbu *TestResultBESUpdate) SetTestCollection(t *TestCollection) *TestResultBESUpdate {
	return trbu.SetTestCollectionID(t.ID)
}

// AddTestActionOutputIDs adds the "test_action_output" edge to the TestFile entity by IDs.
func (trbu *TestResultBESUpdate) AddTestActionOutputIDs(ids ...int) *TestResultBESUpdate {
	trbu.mutation.AddTestActionOutputIDs(ids...)
	return trbu
}

// AddTestActionOutput adds the "test_action_output" edges to the TestFile entity.
func (trbu *TestResultBESUpdate) AddTestActionOutput(t ...*TestFile) *TestResultBESUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return trbu.AddTestActionOutputIDs(ids...)
}

// SetExecutionInfoID sets the "execution_info" edge to the ExectionInfo entity by ID.
func (trbu *TestResultBESUpdate) SetExecutionInfoID(id int) *TestResultBESUpdate {
	trbu.mutation.SetExecutionInfoID(id)
	return trbu
}

// SetNillableExecutionInfoID sets the "execution_info" edge to the ExectionInfo entity by ID if the given value is not nil.
func (trbu *TestResultBESUpdate) SetNillableExecutionInfoID(id *int) *TestResultBESUpdate {
	if id != nil {
		trbu = trbu.SetExecutionInfoID(*id)
	}
	return trbu
}

// SetExecutionInfo sets the "execution_info" edge to the ExectionInfo entity.
func (trbu *TestResultBESUpdate) SetExecutionInfo(e *ExectionInfo) *TestResultBESUpdate {
	return trbu.SetExecutionInfoID(e.ID)
}

// Mutation returns the TestResultBESMutation object of the builder.
func (trbu *TestResultBESUpdate) Mutation() *TestResultBESMutation {
	return trbu.mutation
}

// ClearTestCollection clears the "test_collection" edge to the TestCollection entity.
func (trbu *TestResultBESUpdate) ClearTestCollection() *TestResultBESUpdate {
	trbu.mutation.ClearTestCollection()
	return trbu
}

// ClearTestActionOutput clears all "test_action_output" edges to the TestFile entity.
func (trbu *TestResultBESUpdate) ClearTestActionOutput() *TestResultBESUpdate {
	trbu.mutation.ClearTestActionOutput()
	return trbu
}

// RemoveTestActionOutputIDs removes the "test_action_output" edge to TestFile entities by IDs.
func (trbu *TestResultBESUpdate) RemoveTestActionOutputIDs(ids ...int) *TestResultBESUpdate {
	trbu.mutation.RemoveTestActionOutputIDs(ids...)
	return trbu
}

// RemoveTestActionOutput removes "test_action_output" edges to TestFile entities.
func (trbu *TestResultBESUpdate) RemoveTestActionOutput(t ...*TestFile) *TestResultBESUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return trbu.RemoveTestActionOutputIDs(ids...)
}

// ClearExecutionInfo clears the "execution_info" edge to the ExectionInfo entity.
func (trbu *TestResultBESUpdate) ClearExecutionInfo() *TestResultBESUpdate {
	trbu.mutation.ClearExecutionInfo()
	return trbu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (trbu *TestResultBESUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, trbu.sqlSave, trbu.mutation, trbu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (trbu *TestResultBESUpdate) SaveX(ctx context.Context) int {
	affected, err := trbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (trbu *TestResultBESUpdate) Exec(ctx context.Context) error {
	_, err := trbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trbu *TestResultBESUpdate) ExecX(ctx context.Context) {
	if err := trbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (trbu *TestResultBESUpdate) check() error {
	if v, ok := trbu.mutation.TestStatus(); ok {
		if err := testresultbes.TestStatusValidator(v); err != nil {
			return &ValidationError{Name: "test_status", err: fmt.Errorf(`ent: validator failed for field "TestResultBES.test_status": %w`, err)}
		}
	}
	return nil
}

func (trbu *TestResultBESUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := trbu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(testresultbes.Table, testresultbes.Columns, sqlgraph.NewFieldSpec(testresultbes.FieldID, field.TypeInt))
	if ps := trbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := trbu.mutation.TestStatus(); ok {
		_spec.SetField(testresultbes.FieldTestStatus, field.TypeEnum, value)
	}
	if trbu.mutation.TestStatusCleared() {
		_spec.ClearField(testresultbes.FieldTestStatus, field.TypeEnum)
	}
	if value, ok := trbu.mutation.StatusDetails(); ok {
		_spec.SetField(testresultbes.FieldStatusDetails, field.TypeString, value)
	}
	if trbu.mutation.StatusDetailsCleared() {
		_spec.ClearField(testresultbes.FieldStatusDetails, field.TypeString)
	}
	if value, ok := trbu.mutation.Label(); ok {
		_spec.SetField(testresultbes.FieldLabel, field.TypeString, value)
	}
	if trbu.mutation.LabelCleared() {
		_spec.ClearField(testresultbes.FieldLabel, field.TypeString)
	}
	if value, ok := trbu.mutation.Warning(); ok {
		_spec.SetField(testresultbes.FieldWarning, field.TypeJSON, value)
	}
	if value, ok := trbu.mutation.AppendedWarning(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, testresultbes.FieldWarning, value)
		})
	}
	if trbu.mutation.WarningCleared() {
		_spec.ClearField(testresultbes.FieldWarning, field.TypeJSON)
	}
	if value, ok := trbu.mutation.CachedLocally(); ok {
		_spec.SetField(testresultbes.FieldCachedLocally, field.TypeBool, value)
	}
	if trbu.mutation.CachedLocallyCleared() {
		_spec.ClearField(testresultbes.FieldCachedLocally, field.TypeBool)
	}
	if value, ok := trbu.mutation.TestAttemptStartMillisEpoch(); ok {
		_spec.SetField(testresultbes.FieldTestAttemptStartMillisEpoch, field.TypeInt64, value)
	}
	if value, ok := trbu.mutation.AddedTestAttemptStartMillisEpoch(); ok {
		_spec.AddField(testresultbes.FieldTestAttemptStartMillisEpoch, field.TypeInt64, value)
	}
	if trbu.mutation.TestAttemptStartMillisEpochCleared() {
		_spec.ClearField(testresultbes.FieldTestAttemptStartMillisEpoch, field.TypeInt64)
	}
	if value, ok := trbu.mutation.TestAttemptStart(); ok {
		_spec.SetField(testresultbes.FieldTestAttemptStart, field.TypeString, value)
	}
	if trbu.mutation.TestAttemptStartCleared() {
		_spec.ClearField(testresultbes.FieldTestAttemptStart, field.TypeString)
	}
	if value, ok := trbu.mutation.TestAttemptDurationMillis(); ok {
		_spec.SetField(testresultbes.FieldTestAttemptDurationMillis, field.TypeInt64, value)
	}
	if value, ok := trbu.mutation.AddedTestAttemptDurationMillis(); ok {
		_spec.AddField(testresultbes.FieldTestAttemptDurationMillis, field.TypeInt64, value)
	}
	if trbu.mutation.TestAttemptDurationMillisCleared() {
		_spec.ClearField(testresultbes.FieldTestAttemptDurationMillis, field.TypeInt64)
	}
	if value, ok := trbu.mutation.TestAttemptDuration(); ok {
		_spec.SetField(testresultbes.FieldTestAttemptDuration, field.TypeInt64, value)
	}
	if value, ok := trbu.mutation.AddedTestAttemptDuration(); ok {
		_spec.AddField(testresultbes.FieldTestAttemptDuration, field.TypeInt64, value)
	}
	if trbu.mutation.TestAttemptDurationCleared() {
		_spec.ClearField(testresultbes.FieldTestAttemptDuration, field.TypeInt64)
	}
	if trbu.mutation.TestCollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testresultbes.TestCollectionTable,
			Columns: []string{testresultbes.TestCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcollection.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := trbu.mutation.TestCollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testresultbes.TestCollectionTable,
			Columns: []string{testresultbes.TestCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcollection.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if trbu.mutation.TestActionOutputCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testresultbes.TestActionOutputTable,
			Columns: []string{testresultbes.TestActionOutputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := trbu.mutation.RemovedTestActionOutputIDs(); len(nodes) > 0 && !trbu.mutation.TestActionOutputCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testresultbes.TestActionOutputTable,
			Columns: []string{testresultbes.TestActionOutputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := trbu.mutation.TestActionOutputIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testresultbes.TestActionOutputTable,
			Columns: []string{testresultbes.TestActionOutputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if trbu.mutation.ExecutionInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   testresultbes.ExecutionInfoTable,
			Columns: []string{testresultbes.ExecutionInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exectioninfo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := trbu.mutation.ExecutionInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   testresultbes.ExecutionInfoTable,
			Columns: []string{testresultbes.ExecutionInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exectioninfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, trbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testresultbes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	trbu.mutation.done = true
	return n, nil
}

// TestResultBESUpdateOne is the builder for updating a single TestResultBES entity.
type TestResultBESUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TestResultBESMutation
}

// SetTestStatus sets the "test_status" field.
func (trbuo *TestResultBESUpdateOne) SetTestStatus(ts testresultbes.TestStatus) *TestResultBESUpdateOne {
	trbuo.mutation.SetTestStatus(ts)
	return trbuo
}

// SetNillableTestStatus sets the "test_status" field if the given value is not nil.
func (trbuo *TestResultBESUpdateOne) SetNillableTestStatus(ts *testresultbes.TestStatus) *TestResultBESUpdateOne {
	if ts != nil {
		trbuo.SetTestStatus(*ts)
	}
	return trbuo
}

// ClearTestStatus clears the value of the "test_status" field.
func (trbuo *TestResultBESUpdateOne) ClearTestStatus() *TestResultBESUpdateOne {
	trbuo.mutation.ClearTestStatus()
	return trbuo
}

// SetStatusDetails sets the "status_details" field.
func (trbuo *TestResultBESUpdateOne) SetStatusDetails(s string) *TestResultBESUpdateOne {
	trbuo.mutation.SetStatusDetails(s)
	return trbuo
}

// SetNillableStatusDetails sets the "status_details" field if the given value is not nil.
func (trbuo *TestResultBESUpdateOne) SetNillableStatusDetails(s *string) *TestResultBESUpdateOne {
	if s != nil {
		trbuo.SetStatusDetails(*s)
	}
	return trbuo
}

// ClearStatusDetails clears the value of the "status_details" field.
func (trbuo *TestResultBESUpdateOne) ClearStatusDetails() *TestResultBESUpdateOne {
	trbuo.mutation.ClearStatusDetails()
	return trbuo
}

// SetLabel sets the "label" field.
func (trbuo *TestResultBESUpdateOne) SetLabel(s string) *TestResultBESUpdateOne {
	trbuo.mutation.SetLabel(s)
	return trbuo
}

// SetNillableLabel sets the "label" field if the given value is not nil.
func (trbuo *TestResultBESUpdateOne) SetNillableLabel(s *string) *TestResultBESUpdateOne {
	if s != nil {
		trbuo.SetLabel(*s)
	}
	return trbuo
}

// ClearLabel clears the value of the "label" field.
func (trbuo *TestResultBESUpdateOne) ClearLabel() *TestResultBESUpdateOne {
	trbuo.mutation.ClearLabel()
	return trbuo
}

// SetWarning sets the "warning" field.
func (trbuo *TestResultBESUpdateOne) SetWarning(s []string) *TestResultBESUpdateOne {
	trbuo.mutation.SetWarning(s)
	return trbuo
}

// AppendWarning appends s to the "warning" field.
func (trbuo *TestResultBESUpdateOne) AppendWarning(s []string) *TestResultBESUpdateOne {
	trbuo.mutation.AppendWarning(s)
	return trbuo
}

// ClearWarning clears the value of the "warning" field.
func (trbuo *TestResultBESUpdateOne) ClearWarning() *TestResultBESUpdateOne {
	trbuo.mutation.ClearWarning()
	return trbuo
}

// SetCachedLocally sets the "cached_locally" field.
func (trbuo *TestResultBESUpdateOne) SetCachedLocally(b bool) *TestResultBESUpdateOne {
	trbuo.mutation.SetCachedLocally(b)
	return trbuo
}

// SetNillableCachedLocally sets the "cached_locally" field if the given value is not nil.
func (trbuo *TestResultBESUpdateOne) SetNillableCachedLocally(b *bool) *TestResultBESUpdateOne {
	if b != nil {
		trbuo.SetCachedLocally(*b)
	}
	return trbuo
}

// ClearCachedLocally clears the value of the "cached_locally" field.
func (trbuo *TestResultBESUpdateOne) ClearCachedLocally() *TestResultBESUpdateOne {
	trbuo.mutation.ClearCachedLocally()
	return trbuo
}

// SetTestAttemptStartMillisEpoch sets the "test_attempt_start_millis_epoch" field.
func (trbuo *TestResultBESUpdateOne) SetTestAttemptStartMillisEpoch(i int64) *TestResultBESUpdateOne {
	trbuo.mutation.ResetTestAttemptStartMillisEpoch()
	trbuo.mutation.SetTestAttemptStartMillisEpoch(i)
	return trbuo
}

// SetNillableTestAttemptStartMillisEpoch sets the "test_attempt_start_millis_epoch" field if the given value is not nil.
func (trbuo *TestResultBESUpdateOne) SetNillableTestAttemptStartMillisEpoch(i *int64) *TestResultBESUpdateOne {
	if i != nil {
		trbuo.SetTestAttemptStartMillisEpoch(*i)
	}
	return trbuo
}

// AddTestAttemptStartMillisEpoch adds i to the "test_attempt_start_millis_epoch" field.
func (trbuo *TestResultBESUpdateOne) AddTestAttemptStartMillisEpoch(i int64) *TestResultBESUpdateOne {
	trbuo.mutation.AddTestAttemptStartMillisEpoch(i)
	return trbuo
}

// ClearTestAttemptStartMillisEpoch clears the value of the "test_attempt_start_millis_epoch" field.
func (trbuo *TestResultBESUpdateOne) ClearTestAttemptStartMillisEpoch() *TestResultBESUpdateOne {
	trbuo.mutation.ClearTestAttemptStartMillisEpoch()
	return trbuo
}

// SetTestAttemptStart sets the "test_attempt_start" field.
func (trbuo *TestResultBESUpdateOne) SetTestAttemptStart(s string) *TestResultBESUpdateOne {
	trbuo.mutation.SetTestAttemptStart(s)
	return trbuo
}

// SetNillableTestAttemptStart sets the "test_attempt_start" field if the given value is not nil.
func (trbuo *TestResultBESUpdateOne) SetNillableTestAttemptStart(s *string) *TestResultBESUpdateOne {
	if s != nil {
		trbuo.SetTestAttemptStart(*s)
	}
	return trbuo
}

// ClearTestAttemptStart clears the value of the "test_attempt_start" field.
func (trbuo *TestResultBESUpdateOne) ClearTestAttemptStart() *TestResultBESUpdateOne {
	trbuo.mutation.ClearTestAttemptStart()
	return trbuo
}

// SetTestAttemptDurationMillis sets the "test_attempt_duration_millis" field.
func (trbuo *TestResultBESUpdateOne) SetTestAttemptDurationMillis(i int64) *TestResultBESUpdateOne {
	trbuo.mutation.ResetTestAttemptDurationMillis()
	trbuo.mutation.SetTestAttemptDurationMillis(i)
	return trbuo
}

// SetNillableTestAttemptDurationMillis sets the "test_attempt_duration_millis" field if the given value is not nil.
func (trbuo *TestResultBESUpdateOne) SetNillableTestAttemptDurationMillis(i *int64) *TestResultBESUpdateOne {
	if i != nil {
		trbuo.SetTestAttemptDurationMillis(*i)
	}
	return trbuo
}

// AddTestAttemptDurationMillis adds i to the "test_attempt_duration_millis" field.
func (trbuo *TestResultBESUpdateOne) AddTestAttemptDurationMillis(i int64) *TestResultBESUpdateOne {
	trbuo.mutation.AddTestAttemptDurationMillis(i)
	return trbuo
}

// ClearTestAttemptDurationMillis clears the value of the "test_attempt_duration_millis" field.
func (trbuo *TestResultBESUpdateOne) ClearTestAttemptDurationMillis() *TestResultBESUpdateOne {
	trbuo.mutation.ClearTestAttemptDurationMillis()
	return trbuo
}

// SetTestAttemptDuration sets the "test_attempt_duration" field.
func (trbuo *TestResultBESUpdateOne) SetTestAttemptDuration(i int64) *TestResultBESUpdateOne {
	trbuo.mutation.ResetTestAttemptDuration()
	trbuo.mutation.SetTestAttemptDuration(i)
	return trbuo
}

// SetNillableTestAttemptDuration sets the "test_attempt_duration" field if the given value is not nil.
func (trbuo *TestResultBESUpdateOne) SetNillableTestAttemptDuration(i *int64) *TestResultBESUpdateOne {
	if i != nil {
		trbuo.SetTestAttemptDuration(*i)
	}
	return trbuo
}

// AddTestAttemptDuration adds i to the "test_attempt_duration" field.
func (trbuo *TestResultBESUpdateOne) AddTestAttemptDuration(i int64) *TestResultBESUpdateOne {
	trbuo.mutation.AddTestAttemptDuration(i)
	return trbuo
}

// ClearTestAttemptDuration clears the value of the "test_attempt_duration" field.
func (trbuo *TestResultBESUpdateOne) ClearTestAttemptDuration() *TestResultBESUpdateOne {
	trbuo.mutation.ClearTestAttemptDuration()
	return trbuo
}

// SetTestCollectionID sets the "test_collection" edge to the TestCollection entity by ID.
func (trbuo *TestResultBESUpdateOne) SetTestCollectionID(id int) *TestResultBESUpdateOne {
	trbuo.mutation.SetTestCollectionID(id)
	return trbuo
}

// SetNillableTestCollectionID sets the "test_collection" edge to the TestCollection entity by ID if the given value is not nil.
func (trbuo *TestResultBESUpdateOne) SetNillableTestCollectionID(id *int) *TestResultBESUpdateOne {
	if id != nil {
		trbuo = trbuo.SetTestCollectionID(*id)
	}
	return trbuo
}

// SetTestCollection sets the "test_collection" edge to the TestCollection entity.
func (trbuo *TestResultBESUpdateOne) SetTestCollection(t *TestCollection) *TestResultBESUpdateOne {
	return trbuo.SetTestCollectionID(t.ID)
}

// AddTestActionOutputIDs adds the "test_action_output" edge to the TestFile entity by IDs.
func (trbuo *TestResultBESUpdateOne) AddTestActionOutputIDs(ids ...int) *TestResultBESUpdateOne {
	trbuo.mutation.AddTestActionOutputIDs(ids...)
	return trbuo
}

// AddTestActionOutput adds the "test_action_output" edges to the TestFile entity.
func (trbuo *TestResultBESUpdateOne) AddTestActionOutput(t ...*TestFile) *TestResultBESUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return trbuo.AddTestActionOutputIDs(ids...)
}

// SetExecutionInfoID sets the "execution_info" edge to the ExectionInfo entity by ID.
func (trbuo *TestResultBESUpdateOne) SetExecutionInfoID(id int) *TestResultBESUpdateOne {
	trbuo.mutation.SetExecutionInfoID(id)
	return trbuo
}

// SetNillableExecutionInfoID sets the "execution_info" edge to the ExectionInfo entity by ID if the given value is not nil.
func (trbuo *TestResultBESUpdateOne) SetNillableExecutionInfoID(id *int) *TestResultBESUpdateOne {
	if id != nil {
		trbuo = trbuo.SetExecutionInfoID(*id)
	}
	return trbuo
}

// SetExecutionInfo sets the "execution_info" edge to the ExectionInfo entity.
func (trbuo *TestResultBESUpdateOne) SetExecutionInfo(e *ExectionInfo) *TestResultBESUpdateOne {
	return trbuo.SetExecutionInfoID(e.ID)
}

// Mutation returns the TestResultBESMutation object of the builder.
func (trbuo *TestResultBESUpdateOne) Mutation() *TestResultBESMutation {
	return trbuo.mutation
}

// ClearTestCollection clears the "test_collection" edge to the TestCollection entity.
func (trbuo *TestResultBESUpdateOne) ClearTestCollection() *TestResultBESUpdateOne {
	trbuo.mutation.ClearTestCollection()
	return trbuo
}

// ClearTestActionOutput clears all "test_action_output" edges to the TestFile entity.
func (trbuo *TestResultBESUpdateOne) ClearTestActionOutput() *TestResultBESUpdateOne {
	trbuo.mutation.ClearTestActionOutput()
	return trbuo
}

// RemoveTestActionOutputIDs removes the "test_action_output" edge to TestFile entities by IDs.
func (trbuo *TestResultBESUpdateOne) RemoveTestActionOutputIDs(ids ...int) *TestResultBESUpdateOne {
	trbuo.mutation.RemoveTestActionOutputIDs(ids...)
	return trbuo
}

// RemoveTestActionOutput removes "test_action_output" edges to TestFile entities.
func (trbuo *TestResultBESUpdateOne) RemoveTestActionOutput(t ...*TestFile) *TestResultBESUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return trbuo.RemoveTestActionOutputIDs(ids...)
}

// ClearExecutionInfo clears the "execution_info" edge to the ExectionInfo entity.
func (trbuo *TestResultBESUpdateOne) ClearExecutionInfo() *TestResultBESUpdateOne {
	trbuo.mutation.ClearExecutionInfo()
	return trbuo
}

// Where appends a list predicates to the TestResultBESUpdate builder.
func (trbuo *TestResultBESUpdateOne) Where(ps ...predicate.TestResultBES) *TestResultBESUpdateOne {
	trbuo.mutation.Where(ps...)
	return trbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (trbuo *TestResultBESUpdateOne) Select(field string, fields ...string) *TestResultBESUpdateOne {
	trbuo.fields = append([]string{field}, fields...)
	return trbuo
}

// Save executes the query and returns the updated TestResultBES entity.
func (trbuo *TestResultBESUpdateOne) Save(ctx context.Context) (*TestResultBES, error) {
	return withHooks(ctx, trbuo.sqlSave, trbuo.mutation, trbuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (trbuo *TestResultBESUpdateOne) SaveX(ctx context.Context) *TestResultBES {
	node, err := trbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (trbuo *TestResultBESUpdateOne) Exec(ctx context.Context) error {
	_, err := trbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trbuo *TestResultBESUpdateOne) ExecX(ctx context.Context) {
	if err := trbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (trbuo *TestResultBESUpdateOne) check() error {
	if v, ok := trbuo.mutation.TestStatus(); ok {
		if err := testresultbes.TestStatusValidator(v); err != nil {
			return &ValidationError{Name: "test_status", err: fmt.Errorf(`ent: validator failed for field "TestResultBES.test_status": %w`, err)}
		}
	}
	return nil
}

func (trbuo *TestResultBESUpdateOne) sqlSave(ctx context.Context) (_node *TestResultBES, err error) {
	if err := trbuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(testresultbes.Table, testresultbes.Columns, sqlgraph.NewFieldSpec(testresultbes.FieldID, field.TypeInt))
	id, ok := trbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TestResultBES.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := trbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testresultbes.FieldID)
		for _, f := range fields {
			if !testresultbes.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != testresultbes.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := trbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := trbuo.mutation.TestStatus(); ok {
		_spec.SetField(testresultbes.FieldTestStatus, field.TypeEnum, value)
	}
	if trbuo.mutation.TestStatusCleared() {
		_spec.ClearField(testresultbes.FieldTestStatus, field.TypeEnum)
	}
	if value, ok := trbuo.mutation.StatusDetails(); ok {
		_spec.SetField(testresultbes.FieldStatusDetails, field.TypeString, value)
	}
	if trbuo.mutation.StatusDetailsCleared() {
		_spec.ClearField(testresultbes.FieldStatusDetails, field.TypeString)
	}
	if value, ok := trbuo.mutation.Label(); ok {
		_spec.SetField(testresultbes.FieldLabel, field.TypeString, value)
	}
	if trbuo.mutation.LabelCleared() {
		_spec.ClearField(testresultbes.FieldLabel, field.TypeString)
	}
	if value, ok := trbuo.mutation.Warning(); ok {
		_spec.SetField(testresultbes.FieldWarning, field.TypeJSON, value)
	}
	if value, ok := trbuo.mutation.AppendedWarning(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, testresultbes.FieldWarning, value)
		})
	}
	if trbuo.mutation.WarningCleared() {
		_spec.ClearField(testresultbes.FieldWarning, field.TypeJSON)
	}
	if value, ok := trbuo.mutation.CachedLocally(); ok {
		_spec.SetField(testresultbes.FieldCachedLocally, field.TypeBool, value)
	}
	if trbuo.mutation.CachedLocallyCleared() {
		_spec.ClearField(testresultbes.FieldCachedLocally, field.TypeBool)
	}
	if value, ok := trbuo.mutation.TestAttemptStartMillisEpoch(); ok {
		_spec.SetField(testresultbes.FieldTestAttemptStartMillisEpoch, field.TypeInt64, value)
	}
	if value, ok := trbuo.mutation.AddedTestAttemptStartMillisEpoch(); ok {
		_spec.AddField(testresultbes.FieldTestAttemptStartMillisEpoch, field.TypeInt64, value)
	}
	if trbuo.mutation.TestAttemptStartMillisEpochCleared() {
		_spec.ClearField(testresultbes.FieldTestAttemptStartMillisEpoch, field.TypeInt64)
	}
	if value, ok := trbuo.mutation.TestAttemptStart(); ok {
		_spec.SetField(testresultbes.FieldTestAttemptStart, field.TypeString, value)
	}
	if trbuo.mutation.TestAttemptStartCleared() {
		_spec.ClearField(testresultbes.FieldTestAttemptStart, field.TypeString)
	}
	if value, ok := trbuo.mutation.TestAttemptDurationMillis(); ok {
		_spec.SetField(testresultbes.FieldTestAttemptDurationMillis, field.TypeInt64, value)
	}
	if value, ok := trbuo.mutation.AddedTestAttemptDurationMillis(); ok {
		_spec.AddField(testresultbes.FieldTestAttemptDurationMillis, field.TypeInt64, value)
	}
	if trbuo.mutation.TestAttemptDurationMillisCleared() {
		_spec.ClearField(testresultbes.FieldTestAttemptDurationMillis, field.TypeInt64)
	}
	if value, ok := trbuo.mutation.TestAttemptDuration(); ok {
		_spec.SetField(testresultbes.FieldTestAttemptDuration, field.TypeInt64, value)
	}
	if value, ok := trbuo.mutation.AddedTestAttemptDuration(); ok {
		_spec.AddField(testresultbes.FieldTestAttemptDuration, field.TypeInt64, value)
	}
	if trbuo.mutation.TestAttemptDurationCleared() {
		_spec.ClearField(testresultbes.FieldTestAttemptDuration, field.TypeInt64)
	}
	if trbuo.mutation.TestCollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testresultbes.TestCollectionTable,
			Columns: []string{testresultbes.TestCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcollection.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := trbuo.mutation.TestCollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   testresultbes.TestCollectionTable,
			Columns: []string{testresultbes.TestCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testcollection.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if trbuo.mutation.TestActionOutputCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testresultbes.TestActionOutputTable,
			Columns: []string{testresultbes.TestActionOutputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := trbuo.mutation.RemovedTestActionOutputIDs(); len(nodes) > 0 && !trbuo.mutation.TestActionOutputCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testresultbes.TestActionOutputTable,
			Columns: []string{testresultbes.TestActionOutputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := trbuo.mutation.TestActionOutputIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   testresultbes.TestActionOutputTable,
			Columns: []string{testresultbes.TestActionOutputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(testfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if trbuo.mutation.ExecutionInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   testresultbes.ExecutionInfoTable,
			Columns: []string{testresultbes.ExecutionInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exectioninfo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := trbuo.mutation.ExecutionInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   testresultbes.ExecutionInfoTable,
			Columns: []string{testresultbes.ExecutionInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exectioninfo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TestResultBES{config: trbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, trbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testresultbes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	trbuo.mutation.done = true
	return _node, nil
}
