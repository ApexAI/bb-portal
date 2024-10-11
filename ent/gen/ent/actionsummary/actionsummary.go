// Code generated by ent, DO NOT EDIT.

package actionsummary

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the actionsummary type in the database.
	Label = "action_summary"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActionsCreated holds the string denoting the actions_created field in the database.
	FieldActionsCreated = "actions_created"
	// FieldActionsCreatedNotIncludingAspects holds the string denoting the actions_created_not_including_aspects field in the database.
	FieldActionsCreatedNotIncludingAspects = "actions_created_not_including_aspects"
	// FieldActionsExecuted holds the string denoting the actions_executed field in the database.
	FieldActionsExecuted = "actions_executed"
	// FieldRemoteCacheHits holds the string denoting the remote_cache_hits field in the database.
	FieldRemoteCacheHits = "remote_cache_hits"
	// EdgeMetrics holds the string denoting the metrics edge name in mutations.
	EdgeMetrics = "metrics"
	// EdgeActionData holds the string denoting the action_data edge name in mutations.
	EdgeActionData = "action_data"
	// EdgeRunnerCount holds the string denoting the runner_count edge name in mutations.
	EdgeRunnerCount = "runner_count"
	// EdgeActionCacheStatistics holds the string denoting the action_cache_statistics edge name in mutations.
	EdgeActionCacheStatistics = "action_cache_statistics"
	// Table holds the table name of the actionsummary in the database.
	Table = "action_summaries"
	// MetricsTable is the table that holds the metrics relation/edge.
	MetricsTable = "action_summaries"
	// MetricsInverseTable is the table name for the Metrics entity.
	// It exists in this package in order to avoid circular dependency with the "metrics" package.
	MetricsInverseTable = "metrics"
	// MetricsColumn is the table column denoting the metrics relation/edge.
	MetricsColumn = "metrics_action_summary"
	// ActionDataTable is the table that holds the action_data relation/edge.
	ActionDataTable = "action_data"
	// ActionDataInverseTable is the table name for the ActionData entity.
	// It exists in this package in order to avoid circular dependency with the "actiondata" package.
	ActionDataInverseTable = "action_data"
	// ActionDataColumn is the table column denoting the action_data relation/edge.
	ActionDataColumn = "action_summary_action_data"
	// RunnerCountTable is the table that holds the runner_count relation/edge.
	RunnerCountTable = "runner_counts"
	// RunnerCountInverseTable is the table name for the RunnerCount entity.
	// It exists in this package in order to avoid circular dependency with the "runnercount" package.
	RunnerCountInverseTable = "runner_counts"
	// RunnerCountColumn is the table column denoting the runner_count relation/edge.
	RunnerCountColumn = "action_summary_runner_count"
	// ActionCacheStatisticsTable is the table that holds the action_cache_statistics relation/edge.
	ActionCacheStatisticsTable = "action_cache_statistics"
	// ActionCacheStatisticsInverseTable is the table name for the ActionCacheStatistics entity.
	// It exists in this package in order to avoid circular dependency with the "actioncachestatistics" package.
	ActionCacheStatisticsInverseTable = "action_cache_statistics"
	// ActionCacheStatisticsColumn is the table column denoting the action_cache_statistics relation/edge.
	ActionCacheStatisticsColumn = "action_summary_action_cache_statistics"
)

// Columns holds all SQL columns for actionsummary fields.
var Columns = []string{
	FieldID,
	FieldActionsCreated,
	FieldActionsCreatedNotIncludingAspects,
	FieldActionsExecuted,
	FieldRemoteCacheHits,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "action_summaries"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"metrics_action_summary",
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

// OrderOption defines the ordering options for the ActionSummary queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByActionsCreated orders the results by the actions_created field.
func ByActionsCreated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActionsCreated, opts...).ToFunc()
}

// ByActionsCreatedNotIncludingAspects orders the results by the actions_created_not_including_aspects field.
func ByActionsCreatedNotIncludingAspects(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActionsCreatedNotIncludingAspects, opts...).ToFunc()
}

// ByActionsExecuted orders the results by the actions_executed field.
func ByActionsExecuted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActionsExecuted, opts...).ToFunc()
}

// ByRemoteCacheHits orders the results by the remote_cache_hits field.
func ByRemoteCacheHits(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemoteCacheHits, opts...).ToFunc()
}

// ByMetricsField orders the results by metrics field.
func ByMetricsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMetricsStep(), sql.OrderByField(field, opts...))
	}
}

// ByActionDataCount orders the results by action_data count.
func ByActionDataCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newActionDataStep(), opts...)
	}
}

// ByActionData orders the results by action_data terms.
func ByActionData(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActionDataStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRunnerCountCount orders the results by runner_count count.
func ByRunnerCountCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRunnerCountStep(), opts...)
	}
}

// ByRunnerCount orders the results by runner_count terms.
func ByRunnerCount(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRunnerCountStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByActionCacheStatisticsField orders the results by action_cache_statistics field.
func ByActionCacheStatisticsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActionCacheStatisticsStep(), sql.OrderByField(field, opts...))
	}
}
func newMetricsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MetricsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, MetricsTable, MetricsColumn),
	)
}
func newActionDataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActionDataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ActionDataTable, ActionDataColumn),
	)
}
func newRunnerCountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RunnerCountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RunnerCountTable, RunnerCountColumn),
	)
}
func newActionCacheStatisticsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActionCacheStatisticsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ActionCacheStatisticsTable, ActionCacheStatisticsColumn),
	)
}
