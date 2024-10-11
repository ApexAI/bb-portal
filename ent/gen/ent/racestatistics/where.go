// Code generated by ent, DO NOT EDIT.

package racestatistics

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLTE(FieldID, id))
}

// Mnemonic applies equality check predicate on the "mnemonic" field. It's identical to MnemonicEQ.
func Mnemonic(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldMnemonic, v))
}

// LocalRunner applies equality check predicate on the "local_runner" field. It's identical to LocalRunnerEQ.
func LocalRunner(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldLocalRunner, v))
}

// RemoteRunner applies equality check predicate on the "remote_runner" field. It's identical to RemoteRunnerEQ.
func RemoteRunner(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldRemoteRunner, v))
}

// LocalWins applies equality check predicate on the "local_wins" field. It's identical to LocalWinsEQ.
func LocalWins(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldLocalWins, v))
}

// RenoteWins applies equality check predicate on the "renote_wins" field. It's identical to RenoteWinsEQ.
func RenoteWins(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldRenoteWins, v))
}

// MnemonicEQ applies the EQ predicate on the "mnemonic" field.
func MnemonicEQ(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldMnemonic, v))
}

// MnemonicNEQ applies the NEQ predicate on the "mnemonic" field.
func MnemonicNEQ(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNEQ(FieldMnemonic, v))
}

// MnemonicIn applies the In predicate on the "mnemonic" field.
func MnemonicIn(vs ...string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldIn(FieldMnemonic, vs...))
}

// MnemonicNotIn applies the NotIn predicate on the "mnemonic" field.
func MnemonicNotIn(vs ...string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNotIn(FieldMnemonic, vs...))
}

// MnemonicGT applies the GT predicate on the "mnemonic" field.
func MnemonicGT(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGT(FieldMnemonic, v))
}

// MnemonicGTE applies the GTE predicate on the "mnemonic" field.
func MnemonicGTE(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGTE(FieldMnemonic, v))
}

// MnemonicLT applies the LT predicate on the "mnemonic" field.
func MnemonicLT(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLT(FieldMnemonic, v))
}

// MnemonicLTE applies the LTE predicate on the "mnemonic" field.
func MnemonicLTE(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLTE(FieldMnemonic, v))
}

// MnemonicContains applies the Contains predicate on the "mnemonic" field.
func MnemonicContains(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldContains(FieldMnemonic, v))
}

// MnemonicHasPrefix applies the HasPrefix predicate on the "mnemonic" field.
func MnemonicHasPrefix(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldHasPrefix(FieldMnemonic, v))
}

// MnemonicHasSuffix applies the HasSuffix predicate on the "mnemonic" field.
func MnemonicHasSuffix(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldHasSuffix(FieldMnemonic, v))
}

// MnemonicIsNil applies the IsNil predicate on the "mnemonic" field.
func MnemonicIsNil() predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldIsNull(FieldMnemonic))
}

// MnemonicNotNil applies the NotNil predicate on the "mnemonic" field.
func MnemonicNotNil() predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNotNull(FieldMnemonic))
}

// MnemonicEqualFold applies the EqualFold predicate on the "mnemonic" field.
func MnemonicEqualFold(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEqualFold(FieldMnemonic, v))
}

// MnemonicContainsFold applies the ContainsFold predicate on the "mnemonic" field.
func MnemonicContainsFold(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldContainsFold(FieldMnemonic, v))
}

// LocalRunnerEQ applies the EQ predicate on the "local_runner" field.
func LocalRunnerEQ(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldLocalRunner, v))
}

// LocalRunnerNEQ applies the NEQ predicate on the "local_runner" field.
func LocalRunnerNEQ(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNEQ(FieldLocalRunner, v))
}

// LocalRunnerIn applies the In predicate on the "local_runner" field.
func LocalRunnerIn(vs ...string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldIn(FieldLocalRunner, vs...))
}

// LocalRunnerNotIn applies the NotIn predicate on the "local_runner" field.
func LocalRunnerNotIn(vs ...string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNotIn(FieldLocalRunner, vs...))
}

// LocalRunnerGT applies the GT predicate on the "local_runner" field.
func LocalRunnerGT(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGT(FieldLocalRunner, v))
}

// LocalRunnerGTE applies the GTE predicate on the "local_runner" field.
func LocalRunnerGTE(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGTE(FieldLocalRunner, v))
}

// LocalRunnerLT applies the LT predicate on the "local_runner" field.
func LocalRunnerLT(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLT(FieldLocalRunner, v))
}

// LocalRunnerLTE applies the LTE predicate on the "local_runner" field.
func LocalRunnerLTE(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLTE(FieldLocalRunner, v))
}

// LocalRunnerContains applies the Contains predicate on the "local_runner" field.
func LocalRunnerContains(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldContains(FieldLocalRunner, v))
}

// LocalRunnerHasPrefix applies the HasPrefix predicate on the "local_runner" field.
func LocalRunnerHasPrefix(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldHasPrefix(FieldLocalRunner, v))
}

// LocalRunnerHasSuffix applies the HasSuffix predicate on the "local_runner" field.
func LocalRunnerHasSuffix(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldHasSuffix(FieldLocalRunner, v))
}

// LocalRunnerIsNil applies the IsNil predicate on the "local_runner" field.
func LocalRunnerIsNil() predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldIsNull(FieldLocalRunner))
}

// LocalRunnerNotNil applies the NotNil predicate on the "local_runner" field.
func LocalRunnerNotNil() predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNotNull(FieldLocalRunner))
}

// LocalRunnerEqualFold applies the EqualFold predicate on the "local_runner" field.
func LocalRunnerEqualFold(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEqualFold(FieldLocalRunner, v))
}

// LocalRunnerContainsFold applies the ContainsFold predicate on the "local_runner" field.
func LocalRunnerContainsFold(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldContainsFold(FieldLocalRunner, v))
}

// RemoteRunnerEQ applies the EQ predicate on the "remote_runner" field.
func RemoteRunnerEQ(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldRemoteRunner, v))
}

// RemoteRunnerNEQ applies the NEQ predicate on the "remote_runner" field.
func RemoteRunnerNEQ(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNEQ(FieldRemoteRunner, v))
}

// RemoteRunnerIn applies the In predicate on the "remote_runner" field.
func RemoteRunnerIn(vs ...string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldIn(FieldRemoteRunner, vs...))
}

// RemoteRunnerNotIn applies the NotIn predicate on the "remote_runner" field.
func RemoteRunnerNotIn(vs ...string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNotIn(FieldRemoteRunner, vs...))
}

// RemoteRunnerGT applies the GT predicate on the "remote_runner" field.
func RemoteRunnerGT(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGT(FieldRemoteRunner, v))
}

// RemoteRunnerGTE applies the GTE predicate on the "remote_runner" field.
func RemoteRunnerGTE(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGTE(FieldRemoteRunner, v))
}

// RemoteRunnerLT applies the LT predicate on the "remote_runner" field.
func RemoteRunnerLT(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLT(FieldRemoteRunner, v))
}

// RemoteRunnerLTE applies the LTE predicate on the "remote_runner" field.
func RemoteRunnerLTE(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLTE(FieldRemoteRunner, v))
}

// RemoteRunnerContains applies the Contains predicate on the "remote_runner" field.
func RemoteRunnerContains(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldContains(FieldRemoteRunner, v))
}

// RemoteRunnerHasPrefix applies the HasPrefix predicate on the "remote_runner" field.
func RemoteRunnerHasPrefix(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldHasPrefix(FieldRemoteRunner, v))
}

// RemoteRunnerHasSuffix applies the HasSuffix predicate on the "remote_runner" field.
func RemoteRunnerHasSuffix(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldHasSuffix(FieldRemoteRunner, v))
}

// RemoteRunnerIsNil applies the IsNil predicate on the "remote_runner" field.
func RemoteRunnerIsNil() predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldIsNull(FieldRemoteRunner))
}

// RemoteRunnerNotNil applies the NotNil predicate on the "remote_runner" field.
func RemoteRunnerNotNil() predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNotNull(FieldRemoteRunner))
}

// RemoteRunnerEqualFold applies the EqualFold predicate on the "remote_runner" field.
func RemoteRunnerEqualFold(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEqualFold(FieldRemoteRunner, v))
}

// RemoteRunnerContainsFold applies the ContainsFold predicate on the "remote_runner" field.
func RemoteRunnerContainsFold(v string) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldContainsFold(FieldRemoteRunner, v))
}

// LocalWinsEQ applies the EQ predicate on the "local_wins" field.
func LocalWinsEQ(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldLocalWins, v))
}

// LocalWinsNEQ applies the NEQ predicate on the "local_wins" field.
func LocalWinsNEQ(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNEQ(FieldLocalWins, v))
}

// LocalWinsIn applies the In predicate on the "local_wins" field.
func LocalWinsIn(vs ...int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldIn(FieldLocalWins, vs...))
}

// LocalWinsNotIn applies the NotIn predicate on the "local_wins" field.
func LocalWinsNotIn(vs ...int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNotIn(FieldLocalWins, vs...))
}

// LocalWinsGT applies the GT predicate on the "local_wins" field.
func LocalWinsGT(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGT(FieldLocalWins, v))
}

// LocalWinsGTE applies the GTE predicate on the "local_wins" field.
func LocalWinsGTE(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGTE(FieldLocalWins, v))
}

// LocalWinsLT applies the LT predicate on the "local_wins" field.
func LocalWinsLT(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLT(FieldLocalWins, v))
}

// LocalWinsLTE applies the LTE predicate on the "local_wins" field.
func LocalWinsLTE(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLTE(FieldLocalWins, v))
}

// LocalWinsIsNil applies the IsNil predicate on the "local_wins" field.
func LocalWinsIsNil() predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldIsNull(FieldLocalWins))
}

// LocalWinsNotNil applies the NotNil predicate on the "local_wins" field.
func LocalWinsNotNil() predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNotNull(FieldLocalWins))
}

// RenoteWinsEQ applies the EQ predicate on the "renote_wins" field.
func RenoteWinsEQ(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldEQ(FieldRenoteWins, v))
}

// RenoteWinsNEQ applies the NEQ predicate on the "renote_wins" field.
func RenoteWinsNEQ(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNEQ(FieldRenoteWins, v))
}

// RenoteWinsIn applies the In predicate on the "renote_wins" field.
func RenoteWinsIn(vs ...int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldIn(FieldRenoteWins, vs...))
}

// RenoteWinsNotIn applies the NotIn predicate on the "renote_wins" field.
func RenoteWinsNotIn(vs ...int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNotIn(FieldRenoteWins, vs...))
}

// RenoteWinsGT applies the GT predicate on the "renote_wins" field.
func RenoteWinsGT(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGT(FieldRenoteWins, v))
}

// RenoteWinsGTE applies the GTE predicate on the "renote_wins" field.
func RenoteWinsGTE(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldGTE(FieldRenoteWins, v))
}

// RenoteWinsLT applies the LT predicate on the "renote_wins" field.
func RenoteWinsLT(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLT(FieldRenoteWins, v))
}

// RenoteWinsLTE applies the LTE predicate on the "renote_wins" field.
func RenoteWinsLTE(v int64) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldLTE(FieldRenoteWins, v))
}

// RenoteWinsIsNil applies the IsNil predicate on the "renote_wins" field.
func RenoteWinsIsNil() predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldIsNull(FieldRenoteWins))
}

// RenoteWinsNotNil applies the NotNil predicate on the "renote_wins" field.
func RenoteWinsNotNil() predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.FieldNotNull(FieldRenoteWins))
}

// HasDynamicExecutionMetrics applies the HasEdge predicate on the "dynamic_execution_metrics" edge.
func HasDynamicExecutionMetrics() predicate.RaceStatistics {
	return predicate.RaceStatistics(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DynamicExecutionMetricsTable, DynamicExecutionMetricsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDynamicExecutionMetricsWith applies the HasEdge predicate on the "dynamic_execution_metrics" edge with a given conditions (other predicates).
func HasDynamicExecutionMetricsWith(preds ...predicate.DynamicExecutionMetrics) predicate.RaceStatistics {
	return predicate.RaceStatistics(func(s *sql.Selector) {
		step := newDynamicExecutionMetricsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RaceStatistics) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RaceStatistics) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RaceStatistics) predicate.RaceStatistics {
	return predicate.RaceStatistics(sql.NotPredicates(p))
}
