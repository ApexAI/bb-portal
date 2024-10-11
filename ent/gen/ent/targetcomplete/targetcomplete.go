// Code generated by ent, DO NOT EDIT.

package targetcomplete

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the targetcomplete type in the database.
	Label = "target_complete"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSuccess holds the string denoting the success field in the database.
	FieldSuccess = "success"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// FieldTargetKind holds the string denoting the target_kind field in the database.
	FieldTargetKind = "target_kind"
	// FieldEndTimeInMs holds the string denoting the end_time_in_ms field in the database.
	FieldEndTimeInMs = "end_time_in_ms"
	// FieldTestTimeoutSeconds holds the string denoting the test_timeout_seconds field in the database.
	FieldTestTimeoutSeconds = "test_timeout_seconds"
	// FieldTestTimeout holds the string denoting the test_timeout field in the database.
	FieldTestTimeout = "test_timeout"
	// FieldTestSize holds the string denoting the test_size field in the database.
	FieldTestSize = "test_size"
	// EdgeTargetPair holds the string denoting the target_pair edge name in mutations.
	EdgeTargetPair = "target_pair"
	// EdgeImportantOutput holds the string denoting the important_output edge name in mutations.
	EdgeImportantOutput = "important_output"
	// EdgeDirectoryOutput holds the string denoting the directory_output edge name in mutations.
	EdgeDirectoryOutput = "directory_output"
	// EdgeOutputGroup holds the string denoting the output_group edge name in mutations.
	EdgeOutputGroup = "output_group"
	// Table holds the table name of the targetcomplete in the database.
	Table = "target_completes"
	// TargetPairTable is the table that holds the target_pair relation/edge.
	TargetPairTable = "target_completes"
	// TargetPairInverseTable is the table name for the TargetPair entity.
	// It exists in this package in order to avoid circular dependency with the "targetpair" package.
	TargetPairInverseTable = "target_pairs"
	// TargetPairColumn is the table column denoting the target_pair relation/edge.
	TargetPairColumn = "target_pair_completion"
	// ImportantOutputTable is the table that holds the important_output relation/edge.
	ImportantOutputTable = "test_files"
	// ImportantOutputInverseTable is the table name for the TestFile entity.
	// It exists in this package in order to avoid circular dependency with the "testfile" package.
	ImportantOutputInverseTable = "test_files"
	// ImportantOutputColumn is the table column denoting the important_output relation/edge.
	ImportantOutputColumn = "target_complete_important_output"
	// DirectoryOutputTable is the table that holds the directory_output relation/edge.
	DirectoryOutputTable = "test_files"
	// DirectoryOutputInverseTable is the table name for the TestFile entity.
	// It exists in this package in order to avoid circular dependency with the "testfile" package.
	DirectoryOutputInverseTable = "test_files"
	// DirectoryOutputColumn is the table column denoting the directory_output relation/edge.
	DirectoryOutputColumn = "target_complete_directory_output"
	// OutputGroupTable is the table that holds the output_group relation/edge.
	OutputGroupTable = "output_groups"
	// OutputGroupInverseTable is the table name for the OutputGroup entity.
	// It exists in this package in order to avoid circular dependency with the "outputgroup" package.
	OutputGroupInverseTable = "output_groups"
	// OutputGroupColumn is the table column denoting the output_group relation/edge.
	OutputGroupColumn = "target_complete_output_group"
)

// Columns holds all SQL columns for targetcomplete fields.
var Columns = []string{
	FieldID,
	FieldSuccess,
	FieldTag,
	FieldTargetKind,
	FieldEndTimeInMs,
	FieldTestTimeoutSeconds,
	FieldTestTimeout,
	FieldTestSize,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "target_completes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"target_pair_completion",
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

// TestSize defines the type for the "test_size" enum field.
type TestSize string

// TestSize values.
const (
	TestSizeUNKNOWN  TestSize = "UNKNOWN"
	TestSizeSMALL    TestSize = "SMALL"
	TestSizeMEDIUM   TestSize = "MEDIUM"
	TestSizeLARGE    TestSize = "LARGE"
	TestSizeENORMOUS TestSize = "ENORMOUS"
)

func (ts TestSize) String() string {
	return string(ts)
}

// TestSizeValidator is a validator for the "test_size" field enum values. It is called by the builders before save.
func TestSizeValidator(ts TestSize) error {
	switch ts {
	case TestSizeUNKNOWN, TestSizeSMALL, TestSizeMEDIUM, TestSizeLARGE, TestSizeENORMOUS:
		return nil
	default:
		return fmt.Errorf("targetcomplete: invalid enum value for test_size field: %q", ts)
	}
}

// OrderOption defines the ordering options for the TargetComplete queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySuccess orders the results by the success field.
func BySuccess(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSuccess, opts...).ToFunc()
}

// ByTargetKind orders the results by the target_kind field.
func ByTargetKind(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetKind, opts...).ToFunc()
}

// ByEndTimeInMs orders the results by the end_time_in_ms field.
func ByEndTimeInMs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTimeInMs, opts...).ToFunc()
}

// ByTestTimeoutSeconds orders the results by the test_timeout_seconds field.
func ByTestTimeoutSeconds(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTestTimeoutSeconds, opts...).ToFunc()
}

// ByTestTimeout orders the results by the test_timeout field.
func ByTestTimeout(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTestTimeout, opts...).ToFunc()
}

// ByTestSize orders the results by the test_size field.
func ByTestSize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTestSize, opts...).ToFunc()
}

// ByTargetPairField orders the results by target_pair field.
func ByTargetPairField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTargetPairStep(), sql.OrderByField(field, opts...))
	}
}

// ByImportantOutputCount orders the results by important_output count.
func ByImportantOutputCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newImportantOutputStep(), opts...)
	}
}

// ByImportantOutput orders the results by important_output terms.
func ByImportantOutput(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newImportantOutputStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDirectoryOutputCount orders the results by directory_output count.
func ByDirectoryOutputCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDirectoryOutputStep(), opts...)
	}
}

// ByDirectoryOutput orders the results by directory_output terms.
func ByDirectoryOutput(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDirectoryOutputStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOutputGroupField orders the results by output_group field.
func ByOutputGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOutputGroupStep(), sql.OrderByField(field, opts...))
	}
}
func newTargetPairStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TargetPairInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, TargetPairTable, TargetPairColumn),
	)
}
func newImportantOutputStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImportantOutputInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ImportantOutputTable, ImportantOutputColumn),
	)
}
func newDirectoryOutputStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DirectoryOutputInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DirectoryOutputTable, DirectoryOutputColumn),
	)
}
func newOutputGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OutputGroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, OutputGroupTable, OutputGroupColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e TestSize) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *TestSize) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = TestSize(str)
	if err := TestSizeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid TestSize", str)
	}
	return nil
}
