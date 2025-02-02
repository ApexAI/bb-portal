// Code generated by ent, DO NOT EDIT.

package namedsetoffiles

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the namedsetoffiles type in the database.
	Label = "named_set_of_files"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeOutputGroup holds the string denoting the output_group edge name in mutations.
	EdgeOutputGroup = "output_group"
	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// EdgeFileSets holds the string denoting the file_sets edge name in mutations.
	EdgeFileSets = "file_sets"
	// Table holds the table name of the namedsetoffiles in the database.
	Table = "named_set_of_files"
	// OutputGroupTable is the table that holds the output_group relation/edge.
	OutputGroupTable = "named_set_of_files"
	// OutputGroupInverseTable is the table name for the OutputGroup entity.
	// It exists in this package in order to avoid circular dependency with the "outputgroup" package.
	OutputGroupInverseTable = "output_groups"
	// OutputGroupColumn is the table column denoting the output_group relation/edge.
	OutputGroupColumn = "output_group_file_sets"
	// FilesTable is the table that holds the files relation/edge.
	FilesTable = "test_files"
	// FilesInverseTable is the table name for the TestFile entity.
	// It exists in this package in order to avoid circular dependency with the "testfile" package.
	FilesInverseTable = "test_files"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "named_set_of_files_files"
	// FileSetsTable is the table that holds the file_sets relation/edge.
	FileSetsTable = "named_set_of_files"
	// FileSetsColumn is the table column denoting the file_sets relation/edge.
	FileSetsColumn = "named_set_of_files_file_sets"
)

// Columns holds all SQL columns for namedsetoffiles fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "named_set_of_files"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"named_set_of_files_file_sets",
	"output_group_file_sets",
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

// OrderOption defines the ordering options for the NamedSetOfFiles queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOutputGroupField orders the results by output_group field.
func ByOutputGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOutputGroupStep(), sql.OrderByField(field, opts...))
	}
}

// ByFilesCount orders the results by files count.
func ByFilesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFilesStep(), opts...)
	}
}

// ByFiles orders the results by files terms.
func ByFiles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFilesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFileSetsField orders the results by file_sets field.
func ByFileSetsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFileSetsStep(), sql.OrderByField(field, opts...))
	}
}
func newOutputGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OutputGroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, OutputGroupTable, OutputGroupColumn),
	)
}
func newFilesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FilesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
	)
}
func newFileSetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, FileSetsTable, FileSetsColumn),
	)
}
