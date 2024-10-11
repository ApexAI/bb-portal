// Code generated by ent, DO NOT EDIT.

package garbagemetrics

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the garbagemetrics type in the database.
	Label = "garbage_metrics"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldGarbageCollected holds the string denoting the garbage_collected field in the database.
	FieldGarbageCollected = "garbage_collected"
	// EdgeMemoryMetrics holds the string denoting the memory_metrics edge name in mutations.
	EdgeMemoryMetrics = "memory_metrics"
	// Table holds the table name of the garbagemetrics in the database.
	Table = "garbage_metrics"
	// MemoryMetricsTable is the table that holds the memory_metrics relation/edge.
	MemoryMetricsTable = "garbage_metrics"
	// MemoryMetricsInverseTable is the table name for the MemoryMetrics entity.
	// It exists in this package in order to avoid circular dependency with the "memorymetrics" package.
	MemoryMetricsInverseTable = "memory_metrics"
	// MemoryMetricsColumn is the table column denoting the memory_metrics relation/edge.
	MemoryMetricsColumn = "memory_metrics_garbage_metrics"
)

// Columns holds all SQL columns for garbagemetrics fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldGarbageCollected,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "garbage_metrics"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"memory_metrics_garbage_metrics",
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

// OrderOption defines the ordering options for the GarbageMetrics queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByGarbageCollected orders the results by the garbage_collected field.
func ByGarbageCollected(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGarbageCollected, opts...).ToFunc()
}

// ByMemoryMetricsField orders the results by memory_metrics field.
func ByMemoryMetricsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMemoryMetricsStep(), sql.OrderByField(field, opts...))
	}
}
func newMemoryMetricsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MemoryMetricsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MemoryMetricsTable, MemoryMetricsColumn),
	)
}
