// Code generated by ent, DO NOT EDIT.

package cumulativemetrics

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the cumulativemetrics type in the database.
	Label = "cumulative_metrics"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumAnalyses holds the string denoting the num_analyses field in the database.
	FieldNumAnalyses = "num_analyses"
	// FieldNumBuilds holds the string denoting the num_builds field in the database.
	FieldNumBuilds = "num_builds"
	// EdgeMetrics holds the string denoting the metrics edge name in mutations.
	EdgeMetrics = "metrics"
	// Table holds the table name of the cumulativemetrics in the database.
	Table = "cumulative_metrics"
	// MetricsTable is the table that holds the metrics relation/edge.
	MetricsTable = "cumulative_metrics"
	// MetricsInverseTable is the table name for the Metrics entity.
	// It exists in this package in order to avoid circular dependency with the "metrics" package.
	MetricsInverseTable = "metrics"
	// MetricsColumn is the table column denoting the metrics relation/edge.
	MetricsColumn = "metrics_cumulative_metrics"
)

// Columns holds all SQL columns for cumulativemetrics fields.
var Columns = []string{
	FieldID,
	FieldNumAnalyses,
	FieldNumBuilds,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cumulative_metrics"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"metrics_cumulative_metrics",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the CumulativeMetrics queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNumAnalyses orders the results by the num_analyses field.
func ByNumAnalyses(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumAnalyses, opts...).ToFunc()
}

// ByNumBuilds orders the results by the num_builds field.
func ByNumBuilds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumBuilds, opts...).ToFunc()
}

// ByMetricsField orders the results by metrics field.
func ByMetricsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMetricsStep(), sql.OrderByField(field, opts...))
	}
}
func newMetricsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MetricsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, MetricsTable, MetricsColumn),
	)
}
