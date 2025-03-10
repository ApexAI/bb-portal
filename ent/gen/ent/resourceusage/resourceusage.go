// Code generated by ent, DO NOT EDIT.

package resourceusage

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the resourceusage type in the database.
	Label = "resource_usage"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgeExecutionInfo holds the string denoting the execution_info edge name in mutations.
	EdgeExecutionInfo = "execution_info"
	// Table holds the table name of the resourceusage in the database.
	Table = "resource_usages"
	// ExecutionInfoTable is the table that holds the execution_info relation/edge.
	ExecutionInfoTable = "resource_usages"
	// ExecutionInfoInverseTable is the table name for the ExectionInfo entity.
	// It exists in this package in order to avoid circular dependency with the "exectioninfo" package.
	ExecutionInfoInverseTable = "exection_infos"
	// ExecutionInfoColumn is the table column denoting the execution_info relation/edge.
	ExecutionInfoColumn = "exection_info_resource_usage"
)

// Columns holds all SQL columns for resourceusage fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "resource_usages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"exection_info_resource_usage",
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

// OrderOption defines the ordering options for the ResourceUsage queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByExecutionInfoField orders the results by execution_info field.
func ByExecutionInfoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExecutionInfoStep(), sql.OrderByField(field, opts...))
	}
}
func newExecutionInfoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExecutionInfoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ExecutionInfoTable, ExecutionInfoColumn),
	)
}
