// Code generated by ent, DO NOT EDIT.

package ent

import (
	"attendr/watcher/ent/asu_watched_class"
	"attendr/watcher/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

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

// SetTitle sets the "title" field.
func (awcu *ASUWatchedClassUpdate) SetTitle(s string) *ASUWatchedClassUpdate {
	awcu.mutation.SetTitle(s)
	return awcu
}

// SetInstructor sets the "instructor" field.
func (awcu *ASUWatchedClassUpdate) SetInstructor(s string) *ASUWatchedClassUpdate {
	awcu.mutation.SetInstructor(s)
	return awcu
}

// SetSubject sets the "subject" field.
func (awcu *ASUWatchedClassUpdate) SetSubject(s string) *ASUWatchedClassUpdate {
	awcu.mutation.SetSubject(s)
	return awcu
}

// SetSubjectNumber sets the "subject_number" field.
func (awcu *ASUWatchedClassUpdate) SetSubjectNumber(s string) *ASUWatchedClassUpdate {
	awcu.mutation.SetSubjectNumber(s)
	return awcu
}

// SetHasOpenSeats sets the "has_open_seats" field.
func (awcu *ASUWatchedClassUpdate) SetHasOpenSeats(b bool) *ASUWatchedClassUpdate {
	awcu.mutation.SetHasOpenSeats(b)
	return awcu
}

// SetNillableHasOpenSeats sets the "has_open_seats" field if the given value is not nil.
func (awcu *ASUWatchedClassUpdate) SetNillableHasOpenSeats(b *bool) *ASUWatchedClassUpdate {
	if b != nil {
		awcu.SetHasOpenSeats(*b)
	}
	return awcu
}

// SetTrackedAt sets the "tracked_at" field.
func (awcu *ASUWatchedClassUpdate) SetTrackedAt(t time.Time) *ASUWatchedClassUpdate {
	awcu.mutation.SetTrackedAt(t)
	return awcu
}

// SetNillableTrackedAt sets the "tracked_at" field if the given value is not nil.
func (awcu *ASUWatchedClassUpdate) SetNillableTrackedAt(t *time.Time) *ASUWatchedClassUpdate {
	if t != nil {
		awcu.SetTrackedAt(*t)
	}
	return awcu
}

// SetClassNumber sets the "class_number" field.
func (awcu *ASUWatchedClassUpdate) SetClassNumber(s string) *ASUWatchedClassUpdate {
	awcu.mutation.SetClassNumber(s)
	return awcu
}

// SetTerm sets the "term" field.
func (awcu *ASUWatchedClassUpdate) SetTerm(s string) *ASUWatchedClassUpdate {
	awcu.mutation.SetTerm(s)
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

func (awcu *ASUWatchedClassUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(asu_watched_class.Table, asu_watched_class.Columns, sqlgraph.NewFieldSpec(asu_watched_class.FieldID, field.TypeInt))
	if ps := awcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := awcu.mutation.Title(); ok {
		_spec.SetField(asu_watched_class.FieldTitle, field.TypeString, value)
	}
	if value, ok := awcu.mutation.Instructor(); ok {
		_spec.SetField(asu_watched_class.FieldInstructor, field.TypeString, value)
	}
	if value, ok := awcu.mutation.Subject(); ok {
		_spec.SetField(asu_watched_class.FieldSubject, field.TypeString, value)
	}
	if value, ok := awcu.mutation.SubjectNumber(); ok {
		_spec.SetField(asu_watched_class.FieldSubjectNumber, field.TypeString, value)
	}
	if value, ok := awcu.mutation.HasOpenSeats(); ok {
		_spec.SetField(asu_watched_class.FieldHasOpenSeats, field.TypeBool, value)
	}
	if value, ok := awcu.mutation.TrackedAt(); ok {
		_spec.SetField(asu_watched_class.FieldTrackedAt, field.TypeTime, value)
	}
	if value, ok := awcu.mutation.ClassNumber(); ok {
		_spec.SetField(asu_watched_class.FieldClassNumber, field.TypeString, value)
	}
	if value, ok := awcu.mutation.Term(); ok {
		_spec.SetField(asu_watched_class.FieldTerm, field.TypeString, value)
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

// SetTitle sets the "title" field.
func (awcuo *ASUWatchedClassUpdateOne) SetTitle(s string) *ASUWatchedClassUpdateOne {
	awcuo.mutation.SetTitle(s)
	return awcuo
}

// SetInstructor sets the "instructor" field.
func (awcuo *ASUWatchedClassUpdateOne) SetInstructor(s string) *ASUWatchedClassUpdateOne {
	awcuo.mutation.SetInstructor(s)
	return awcuo
}

// SetSubject sets the "subject" field.
func (awcuo *ASUWatchedClassUpdateOne) SetSubject(s string) *ASUWatchedClassUpdateOne {
	awcuo.mutation.SetSubject(s)
	return awcuo
}

// SetSubjectNumber sets the "subject_number" field.
func (awcuo *ASUWatchedClassUpdateOne) SetSubjectNumber(s string) *ASUWatchedClassUpdateOne {
	awcuo.mutation.SetSubjectNumber(s)
	return awcuo
}

// SetHasOpenSeats sets the "has_open_seats" field.
func (awcuo *ASUWatchedClassUpdateOne) SetHasOpenSeats(b bool) *ASUWatchedClassUpdateOne {
	awcuo.mutation.SetHasOpenSeats(b)
	return awcuo
}

// SetNillableHasOpenSeats sets the "has_open_seats" field if the given value is not nil.
func (awcuo *ASUWatchedClassUpdateOne) SetNillableHasOpenSeats(b *bool) *ASUWatchedClassUpdateOne {
	if b != nil {
		awcuo.SetHasOpenSeats(*b)
	}
	return awcuo
}

// SetTrackedAt sets the "tracked_at" field.
func (awcuo *ASUWatchedClassUpdateOne) SetTrackedAt(t time.Time) *ASUWatchedClassUpdateOne {
	awcuo.mutation.SetTrackedAt(t)
	return awcuo
}

// SetNillableTrackedAt sets the "tracked_at" field if the given value is not nil.
func (awcuo *ASUWatchedClassUpdateOne) SetNillableTrackedAt(t *time.Time) *ASUWatchedClassUpdateOne {
	if t != nil {
		awcuo.SetTrackedAt(*t)
	}
	return awcuo
}

// SetClassNumber sets the "class_number" field.
func (awcuo *ASUWatchedClassUpdateOne) SetClassNumber(s string) *ASUWatchedClassUpdateOne {
	awcuo.mutation.SetClassNumber(s)
	return awcuo
}

// SetTerm sets the "term" field.
func (awcuo *ASUWatchedClassUpdateOne) SetTerm(s string) *ASUWatchedClassUpdateOne {
	awcuo.mutation.SetTerm(s)
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

func (awcuo *ASUWatchedClassUpdateOne) sqlSave(ctx context.Context) (_node *ASU_Watched_Class, err error) {
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
	if value, ok := awcuo.mutation.Title(); ok {
		_spec.SetField(asu_watched_class.FieldTitle, field.TypeString, value)
	}
	if value, ok := awcuo.mutation.Instructor(); ok {
		_spec.SetField(asu_watched_class.FieldInstructor, field.TypeString, value)
	}
	if value, ok := awcuo.mutation.Subject(); ok {
		_spec.SetField(asu_watched_class.FieldSubject, field.TypeString, value)
	}
	if value, ok := awcuo.mutation.SubjectNumber(); ok {
		_spec.SetField(asu_watched_class.FieldSubjectNumber, field.TypeString, value)
	}
	if value, ok := awcuo.mutation.HasOpenSeats(); ok {
		_spec.SetField(asu_watched_class.FieldHasOpenSeats, field.TypeBool, value)
	}
	if value, ok := awcuo.mutation.TrackedAt(); ok {
		_spec.SetField(asu_watched_class.FieldTrackedAt, field.TypeTime, value)
	}
	if value, ok := awcuo.mutation.ClassNumber(); ok {
		_spec.SetField(asu_watched_class.FieldClassNumber, field.TypeString, value)
	}
	if value, ok := awcuo.mutation.Term(); ok {
		_spec.SetField(asu_watched_class.FieldTerm, field.TypeString, value)
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
