// Code generated by ent, DO NOT EDIT.

package ent

import (
	"attendr/watcher/ent/asuwatchedclass"
	"attendr/watcher/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ASUWatchedClassDelete is the builder for deleting a ASUWatchedClass entity.
type ASUWatchedClassDelete struct {
	config
	hooks    []Hook
	mutation *ASUWatchedClassMutation
}

// Where appends a list predicates to the ASUWatchedClassDelete builder.
func (awcd *ASUWatchedClassDelete) Where(ps ...predicate.ASUWatchedClass) *ASUWatchedClassDelete {
	awcd.mutation.Where(ps...)
	return awcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (awcd *ASUWatchedClassDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, awcd.sqlExec, awcd.mutation, awcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (awcd *ASUWatchedClassDelete) ExecX(ctx context.Context) int {
	n, err := awcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (awcd *ASUWatchedClassDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(asuwatchedclass.Table, sqlgraph.NewFieldSpec(asuwatchedclass.FieldID, field.TypeInt))
	if ps := awcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, awcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	awcd.mutation.done = true
	return affected, err
}

// ASUWatchedClassDeleteOne is the builder for deleting a single ASUWatchedClass entity.
type ASUWatchedClassDeleteOne struct {
	awcd *ASUWatchedClassDelete
}

// Where appends a list predicates to the ASUWatchedClassDelete builder.
func (awcdo *ASUWatchedClassDeleteOne) Where(ps ...predicate.ASUWatchedClass) *ASUWatchedClassDeleteOne {
	awcdo.awcd.mutation.Where(ps...)
	return awcdo
}

// Exec executes the deletion query.
func (awcdo *ASUWatchedClassDeleteOne) Exec(ctx context.Context) error {
	n, err := awcdo.awcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{asuwatchedclass.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (awcdo *ASUWatchedClassDeleteOne) ExecX(ctx context.Context) {
	if err := awcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
