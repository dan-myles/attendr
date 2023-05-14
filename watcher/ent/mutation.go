// Code generated by ent, DO NOT EDIT.

package ent

import (
	"attendr/watcher/ent/asu_watched_class"
	"attendr/watcher/ent/predicate"
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeASUWatchedClass = "ASU_Watched_Class"
)

// ASUWatchedClassMutation represents an operation that mutates the ASU_Watched_Class nodes in the graph.
type ASUWatchedClassMutation struct {
	config
	op            Op
	typ           string
	id            *int
	age           *int
	addage        *int
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*ASU_Watched_Class, error)
	predicates    []predicate.ASU_Watched_Class
}

var _ ent.Mutation = (*ASUWatchedClassMutation)(nil)

// asuWatchedClassOption allows management of the mutation configuration using functional options.
type asuWatchedClassOption func(*ASUWatchedClassMutation)

// newASUWatchedClassMutation creates new mutation for the ASU_Watched_Class entity.
func newASUWatchedClassMutation(c config, op Op, opts ...asuWatchedClassOption) *ASUWatchedClassMutation {
	m := &ASUWatchedClassMutation{
		config:        c,
		op:            op,
		typ:           TypeASUWatchedClass,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withASU_Watched_ClassID sets the ID field of the mutation.
func withASU_Watched_ClassID(id int) asuWatchedClassOption {
	return func(m *ASUWatchedClassMutation) {
		var (
			err   error
			once  sync.Once
			value *ASU_Watched_Class
		)
		m.oldValue = func(ctx context.Context) (*ASU_Watched_Class, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ASU_Watched_Class.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withASU_Watched_Class sets the old ASU_Watched_Class of the mutation.
func withASU_Watched_Class(node *ASU_Watched_Class) asuWatchedClassOption {
	return func(m *ASUWatchedClassMutation) {
		m.oldValue = func(context.Context) (*ASU_Watched_Class, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ASUWatchedClassMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ASUWatchedClassMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ASUWatchedClassMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ASUWatchedClassMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ASU_Watched_Class.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetAge sets the "age" field.
func (m *ASUWatchedClassMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *ASUWatchedClassMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the ASU_Watched_Class entity.
// If the ASU_Watched_Class object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ASUWatchedClassMutation) OldAge(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *ASUWatchedClassMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *ASUWatchedClassMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *ASUWatchedClassMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetName sets the "name" field.
func (m *ASUWatchedClassMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ASUWatchedClassMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the ASU_Watched_Class entity.
// If the ASU_Watched_Class object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ASUWatchedClassMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ASUWatchedClassMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the ASUWatchedClassMutation builder.
func (m *ASUWatchedClassMutation) Where(ps ...predicate.ASU_Watched_Class) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ASUWatchedClassMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ASUWatchedClassMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.ASU_Watched_Class, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ASUWatchedClassMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ASUWatchedClassMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (ASU_Watched_Class).
func (m *ASUWatchedClassMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ASUWatchedClassMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.age != nil {
		fields = append(fields, asu_watched_class.FieldAge)
	}
	if m.name != nil {
		fields = append(fields, asu_watched_class.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ASUWatchedClassMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case asu_watched_class.FieldAge:
		return m.Age()
	case asu_watched_class.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ASUWatchedClassMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case asu_watched_class.FieldAge:
		return m.OldAge(ctx)
	case asu_watched_class.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown ASU_Watched_Class field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ASUWatchedClassMutation) SetField(name string, value ent.Value) error {
	switch name {
	case asu_watched_class.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case asu_watched_class.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown ASU_Watched_Class field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ASUWatchedClassMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, asu_watched_class.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ASUWatchedClassMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case asu_watched_class.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ASUWatchedClassMutation) AddField(name string, value ent.Value) error {
	switch name {
	case asu_watched_class.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown ASU_Watched_Class numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ASUWatchedClassMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ASUWatchedClassMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ASUWatchedClassMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ASU_Watched_Class nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ASUWatchedClassMutation) ResetField(name string) error {
	switch name {
	case asu_watched_class.FieldAge:
		m.ResetAge()
		return nil
	case asu_watched_class.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown ASU_Watched_Class field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ASUWatchedClassMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ASUWatchedClassMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ASUWatchedClassMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ASUWatchedClassMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ASUWatchedClassMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ASUWatchedClassMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ASUWatchedClassMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ASU_Watched_Class unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ASUWatchedClassMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ASU_Watched_Class edge %s", name)
}
