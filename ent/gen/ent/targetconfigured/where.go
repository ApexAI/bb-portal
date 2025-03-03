// Code generated by ent, DO NOT EDIT.

package targetconfigured

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldLTE(FieldID, id))
}

// TargetKind applies equality check predicate on the "target_kind" field. It's identical to TargetKindEQ.
func TargetKind(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldEQ(FieldTargetKind, v))
}

// StartTimeInMs applies equality check predicate on the "start_time_in_ms" field. It's identical to StartTimeInMsEQ.
func StartTimeInMs(v int64) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldEQ(FieldStartTimeInMs, v))
}

// TagIsNil applies the IsNil predicate on the "tag" field.
func TagIsNil() predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldIsNull(FieldTag))
}

// TagNotNil applies the NotNil predicate on the "tag" field.
func TagNotNil() predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNotNull(FieldTag))
}

// TargetKindEQ applies the EQ predicate on the "target_kind" field.
func TargetKindEQ(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldEQ(FieldTargetKind, v))
}

// TargetKindNEQ applies the NEQ predicate on the "target_kind" field.
func TargetKindNEQ(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNEQ(FieldTargetKind, v))
}

// TargetKindIn applies the In predicate on the "target_kind" field.
func TargetKindIn(vs ...string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldIn(FieldTargetKind, vs...))
}

// TargetKindNotIn applies the NotIn predicate on the "target_kind" field.
func TargetKindNotIn(vs ...string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNotIn(FieldTargetKind, vs...))
}

// TargetKindGT applies the GT predicate on the "target_kind" field.
func TargetKindGT(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldGT(FieldTargetKind, v))
}

// TargetKindGTE applies the GTE predicate on the "target_kind" field.
func TargetKindGTE(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldGTE(FieldTargetKind, v))
}

// TargetKindLT applies the LT predicate on the "target_kind" field.
func TargetKindLT(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldLT(FieldTargetKind, v))
}

// TargetKindLTE applies the LTE predicate on the "target_kind" field.
func TargetKindLTE(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldLTE(FieldTargetKind, v))
}

// TargetKindContains applies the Contains predicate on the "target_kind" field.
func TargetKindContains(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldContains(FieldTargetKind, v))
}

// TargetKindHasPrefix applies the HasPrefix predicate on the "target_kind" field.
func TargetKindHasPrefix(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldHasPrefix(FieldTargetKind, v))
}

// TargetKindHasSuffix applies the HasSuffix predicate on the "target_kind" field.
func TargetKindHasSuffix(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldHasSuffix(FieldTargetKind, v))
}

// TargetKindIsNil applies the IsNil predicate on the "target_kind" field.
func TargetKindIsNil() predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldIsNull(FieldTargetKind))
}

// TargetKindNotNil applies the NotNil predicate on the "target_kind" field.
func TargetKindNotNil() predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNotNull(FieldTargetKind))
}

// TargetKindEqualFold applies the EqualFold predicate on the "target_kind" field.
func TargetKindEqualFold(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldEqualFold(FieldTargetKind, v))
}

// TargetKindContainsFold applies the ContainsFold predicate on the "target_kind" field.
func TargetKindContainsFold(v string) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldContainsFold(FieldTargetKind, v))
}

// StartTimeInMsEQ applies the EQ predicate on the "start_time_in_ms" field.
func StartTimeInMsEQ(v int64) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldEQ(FieldStartTimeInMs, v))
}

// StartTimeInMsNEQ applies the NEQ predicate on the "start_time_in_ms" field.
func StartTimeInMsNEQ(v int64) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNEQ(FieldStartTimeInMs, v))
}

// StartTimeInMsIn applies the In predicate on the "start_time_in_ms" field.
func StartTimeInMsIn(vs ...int64) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldIn(FieldStartTimeInMs, vs...))
}

// StartTimeInMsNotIn applies the NotIn predicate on the "start_time_in_ms" field.
func StartTimeInMsNotIn(vs ...int64) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNotIn(FieldStartTimeInMs, vs...))
}

// StartTimeInMsGT applies the GT predicate on the "start_time_in_ms" field.
func StartTimeInMsGT(v int64) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldGT(FieldStartTimeInMs, v))
}

// StartTimeInMsGTE applies the GTE predicate on the "start_time_in_ms" field.
func StartTimeInMsGTE(v int64) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldGTE(FieldStartTimeInMs, v))
}

// StartTimeInMsLT applies the LT predicate on the "start_time_in_ms" field.
func StartTimeInMsLT(v int64) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldLT(FieldStartTimeInMs, v))
}

// StartTimeInMsLTE applies the LTE predicate on the "start_time_in_ms" field.
func StartTimeInMsLTE(v int64) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldLTE(FieldStartTimeInMs, v))
}

// StartTimeInMsIsNil applies the IsNil predicate on the "start_time_in_ms" field.
func StartTimeInMsIsNil() predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldIsNull(FieldStartTimeInMs))
}

// StartTimeInMsNotNil applies the NotNil predicate on the "start_time_in_ms" field.
func StartTimeInMsNotNil() predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNotNull(FieldStartTimeInMs))
}

// TestSizeEQ applies the EQ predicate on the "test_size" field.
func TestSizeEQ(v TestSize) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldEQ(FieldTestSize, v))
}

// TestSizeNEQ applies the NEQ predicate on the "test_size" field.
func TestSizeNEQ(v TestSize) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNEQ(FieldTestSize, v))
}

// TestSizeIn applies the In predicate on the "test_size" field.
func TestSizeIn(vs ...TestSize) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldIn(FieldTestSize, vs...))
}

// TestSizeNotIn applies the NotIn predicate on the "test_size" field.
func TestSizeNotIn(vs ...TestSize) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNotIn(FieldTestSize, vs...))
}

// TestSizeIsNil applies the IsNil predicate on the "test_size" field.
func TestSizeIsNil() predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldIsNull(FieldTestSize))
}

// TestSizeNotNil applies the NotNil predicate on the "test_size" field.
func TestSizeNotNil() predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.FieldNotNull(FieldTestSize))
}

// HasTargetPair applies the HasEdge predicate on the "target_pair" edge.
func HasTargetPair() predicate.TargetConfigured {
	return predicate.TargetConfigured(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TargetPairTable, TargetPairColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTargetPairWith applies the HasEdge predicate on the "target_pair" edge with a given conditions (other predicates).
func HasTargetPairWith(preds ...predicate.TargetPair) predicate.TargetConfigured {
	return predicate.TargetConfigured(func(s *sql.Selector) {
		step := newTargetPairStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TargetConfigured) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TargetConfigured) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TargetConfigured) predicate.TargetConfigured {
	return predicate.TargetConfigured(sql.NotPredicates(p))
}
