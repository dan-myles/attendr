// Code generated by ent, DO NOT EDIT.

package ent

import (
	"attendr/watcher/ent/asu_watched_class"
	"attendr/watcher/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ASUWatchedClassUpdate is the builder for updating ASU_Watched_Class entities.
type ASUWatchedClassUpdate struct {
	config
	hooks    []Hook
	mutation *ASUWatchedClassMutation
}

// Where appends a list predicates to the ASUWatchedClassUpdate builder.
func (awcu *ASUWatchedClassUpdate) Where(ps ...predicate.ASU_Watched_Class) *ASUWatchedClassUpdate {
	awcu.mutation.Where(ps...)
	return awcu
}

// SetAge sets the "age" field.
func (awcu *ASUWatchedClassUpdate) SetAge(i int) *ASUWatchedClassUpdate {
	awcu.mutation.ResetAge()
	awcu.mutation.SetAge(i)
	return awcu
}

// AddAge adds i to the "age" field.
func (awcu *ASUWatchedClassUpdate) AddAge(i int) *ASUWatchedClassUpdate {
	awcu.mutation.AddAge(i)
	return awcu
}

// SetName sets the "name" field.
func (awcu *ASUWatchedClassUpdate) SetName(s string) *ASUWatchedClassUpdate {
	awcu.mutation.SetName(s)
	return awcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (awcu *ASUWatchedClassUpdate) SetNillableName(s *string) *ASUWatchedClassUpdate {
	if s != nil {
		awcu.SetName(*s)
	}
	return awcu
}

// Mutation returns the ASUWatchedClassMutation object of the builder.
func (awcu *ASUWatchedClassUpdate) Mutation() *ASUWatchedClassMutation {
	return awcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (awcu *ASUWatchedClassUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, awcu.sqlSave, awcu.mutation, awcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (awcu *ASUWatchedClassUpdate) SaveX(ctx context.Context) int {
	affected, err := awcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (awcu *ASUWatchedClassUpdate) Exec(ctx context.Context) error {
	_, err := awcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (awcu *ASUWatchedClassUpdate) ExecX(ctx context.Context) {
	if err := awcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (awcu *ASUWatchedClassUpdate) check() error {
	if v, ok := awcu.mutation.Age(); ok {
		if err := asu_watched_class.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "ASU_Watched_Class.age": %w`, err)}
		}
	}
	return nil
}

func (awcu *ASUWatchedClassUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := awcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(asu_watched_class.Table, asu_watched_class.Columns, sqlgraph.NewFieldSpec(asu_watched_class.FieldID, field.TypeInt))
	if ps := awcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := awcu.mutation.Age(); ok {
		_spec.SetField(asu_watched_class.FieldAge, field.TypeInt, value)
	}
	if value, ok := awcu.mutation.AddedAge(); ok {
		_spec.AddField(asu_watched_class.FieldAge, field.TypeInt, value)
	}
	if value, ok := awcu.mutation.Name(); ok {
		_spec.SetField(asu_watched_class.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, awcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asu_watched_class.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	awcu.mutation.done = true
	return n, nil
}

// ASUWatchedClassUpdateOne is the builder for updating a single ASU_Watched_Class entity.
type ASUWatchedClassUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ASUWatchedClassMutation
}

// SetAge sets the "age" field.
func (awcuo *ASUWatchedClassUpdateOne) SetAge(i int) *ASUWatchedClassUpdateOne {
	awcuo.mutation.ResetAge()
	awcuo.mutation.SetAge(i)
	return awcuo
}

// AddAge adds i to the "age" field.
func (awcuo *ASUWatchedClassUpdateOne) AddAge(i int) *ASUWatchedClassUpdateOne {
	awcuo.mutation.AddAge(i)
	return awcuo
}

// SetName sets the "name" field.
func (awcuo *ASUWatchedClassUpdateOne) SetName(s string) *ASUWatchedClassUpdateOne {
	awcuo.mutation.SetName(s)
	return awcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (awcuo *ASUWatchedClassUpdateOne) SetNillableName(s *string) *ASUWatchedClassUpdateOne {
	if s != nil {
		awcuo.SetName(*s)
	}
	return awcuo
}

// Mutation returns the ASUWatchedClassMutation object of the builder.
func (awcuo *ASUWatchedClassUpdateOne) Mutation() *ASUWatchedClassMutation {
	return awcuo.mutation
}

// Where appends a list predicates to the ASUWatchedClassUpdate builder.
func (awcuo *ASUWatchedClassUpdateOne) Where(ps ...predicate.ASU_Watched_Class) *ASUWatchedClassUpdateOne {
	awcuo.mutation.Where(ps...)
	return awcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (awcuo *ASUWatchedClassUpdateOne) Select(field string, fields ...string) *ASUWatchedClassUpdateOne {
	awcuo.fields = append([]string{field}, fields...)
	return awcuo
}

// Save executes the query and returns the updated ASU_Watched_Class entity.
func (awcuo *ASUWatchedClassUpdateOne) Save(ctx context.Context) (*ASU_Watched_Class, error) {
	return withHooks(ctx, awcuo.sqlSave, awcuo.mutation, awcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (awcuo *ASUWatchedClassUpdateOne) SaveX(ctx context.Context) *ASU_Watched_Class {
	node, err := awcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (awcuo *ASUWatchedClassUpdateOne) Exec(ctx context.Context) error {
	_, err := awcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (awcuo *ASUWatchedClassUpdateOne) ExecX(ctx context.Context) {
	if err := awcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (awcuo *ASUWatchedClassUpdateOne) check() error {
	if v, ok := awcuo.mutation.Age(); ok {
		if err := asu_watched_class.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "ASU_Watched_Class.age": %w`, err)}
		}
	}
	return nil
}

func (awcuo *ASUWatchedClassUpdateOne) sqlSave(ctx context.Context) (_node *ASU_Watched_Class, err error) {
	if err := awcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(asu_watched_class.Table, asu_watched_class.Columns, sqlgraph.NewFieldSpec(asu_watched_class.FieldID, field.TypeInt))
	id, ok := awcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ASU_Watched_Class.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := awcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asu_watched_class.FieldID)
		for _, f := range fields {
			if !asu_watched_class.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != asu_watched_class.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := awcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := awcuo.mutation.Age(); ok {
		_spec.SetField(asu_watched_class.FieldAge, field.TypeInt, value)
	}
	if value, ok := awcuo.mutation.AddedAge(); ok {
		_spec.AddField(asu_watched_class.FieldAge, field.TypeInt, value)
	}
	if value, ok := awcuo.mutation.Name(); ok {
		_spec.SetField(asu_watched_class.FieldName, field.TypeString, value)
	}
	_node = &ASU_Watched_Class{config: awcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, awcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asu_watched_class.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	awcuo.mutation.done = true
	return _node, nil
}
