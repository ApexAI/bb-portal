// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/buildbarn/bb-portal/ent/gen/ent/bazelinvocation"
	"github.com/buildbarn/bb-portal/ent/gen/ent/sourcecontrol"
)

// SourceControl is the model entity for the SourceControl schema.
type SourceControl struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RepoURL holds the value of the "repo_url" field.
	RepoURL string `json:"repo_url,omitempty"`
	// Branch holds the value of the "branch" field.
	Branch string `json:"branch,omitempty"`
	// CommitSha holds the value of the "commit_sha" field.
	CommitSha string `json:"commit_sha,omitempty"`
	// Actor holds the value of the "actor" field.
	Actor string `json:"actor,omitempty"`
	// Refs holds the value of the "refs" field.
	Refs string `json:"refs,omitempty"`
	// RunID holds the value of the "run_id" field.
	RunID string `json:"run_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SourceControlQuery when eager-loading is set.
	Edges                           SourceControlEdges `json:"edges"`
	bazel_invocation_source_control *int
	selectValues                    sql.SelectValues
}

// SourceControlEdges holds the relations/edges for other nodes in the graph.
type SourceControlEdges struct {
	// BazelInvocation holds the value of the bazel_invocation edge.
	BazelInvocation *BazelInvocation `json:"bazel_invocation,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// BazelInvocationOrErr returns the BazelInvocation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SourceControlEdges) BazelInvocationOrErr() (*BazelInvocation, error) {
	if e.BazelInvocation != nil {
		return e.BazelInvocation, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: bazelinvocation.Label}
	}
	return nil, &NotLoadedError{edge: "bazel_invocation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SourceControl) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sourcecontrol.FieldID:
			values[i] = new(sql.NullInt64)
		case sourcecontrol.FieldRepoURL, sourcecontrol.FieldBranch, sourcecontrol.FieldCommitSha, sourcecontrol.FieldActor, sourcecontrol.FieldRefs, sourcecontrol.FieldRunID:
			values[i] = new(sql.NullString)
		case sourcecontrol.ForeignKeys[0]: // bazel_invocation_source_control
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SourceControl fields.
func (sc *SourceControl) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sourcecontrol.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sc.ID = int(value.Int64)
		case sourcecontrol.FieldRepoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field repo_url", values[i])
			} else if value.Valid {
				sc.RepoURL = value.String
			}
		case sourcecontrol.FieldBranch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field branch", values[i])
			} else if value.Valid {
				sc.Branch = value.String
			}
		case sourcecontrol.FieldCommitSha:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commit_sha", values[i])
			} else if value.Valid {
				sc.CommitSha = value.String
			}
		case sourcecontrol.FieldActor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field actor", values[i])
			} else if value.Valid {
				sc.Actor = value.String
			}
		case sourcecontrol.FieldRefs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refs", values[i])
			} else if value.Valid {
				sc.Refs = value.String
			}
		case sourcecontrol.FieldRunID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field run_id", values[i])
			} else if value.Valid {
				sc.RunID = value.String
			}
		case sourcecontrol.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field bazel_invocation_source_control", value)
			} else if value.Valid {
				sc.bazel_invocation_source_control = new(int)
				*sc.bazel_invocation_source_control = int(value.Int64)
			}
		default:
			sc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SourceControl.
// This includes values selected through modifiers, order, etc.
func (sc *SourceControl) Value(name string) (ent.Value, error) {
	return sc.selectValues.Get(name)
}

// QueryBazelInvocation queries the "bazel_invocation" edge of the SourceControl entity.
func (sc *SourceControl) QueryBazelInvocation() *BazelInvocationQuery {
	return NewSourceControlClient(sc.config).QueryBazelInvocation(sc)
}

// Update returns a builder for updating this SourceControl.
// Note that you need to call SourceControl.Unwrap() before calling this method if this SourceControl
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *SourceControl) Update() *SourceControlUpdateOne {
	return NewSourceControlClient(sc.config).UpdateOne(sc)
}

// Unwrap unwraps the SourceControl entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *SourceControl) Unwrap() *SourceControl {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: SourceControl is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *SourceControl) String() string {
	var builder strings.Builder
	builder.WriteString("SourceControl(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("repo_url=")
	builder.WriteString(sc.RepoURL)
	builder.WriteString(", ")
	builder.WriteString("branch=")
	builder.WriteString(sc.Branch)
	builder.WriteString(", ")
	builder.WriteString("commit_sha=")
	builder.WriteString(sc.CommitSha)
	builder.WriteString(", ")
	builder.WriteString("actor=")
	builder.WriteString(sc.Actor)
	builder.WriteString(", ")
	builder.WriteString("refs=")
	builder.WriteString(sc.Refs)
	builder.WriteString(", ")
	builder.WriteString("run_id=")
	builder.WriteString(sc.RunID)
	builder.WriteByte(')')
	return builder.String()
}

// SourceControls is a parsable slice of SourceControl.
type SourceControls []*SourceControl