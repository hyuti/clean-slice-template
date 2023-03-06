// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/hyuti/clean-slice-template/services/service-1/ent/predicate"
	"github.com/hyuti/clean-slice-template/services/service-1/ent/product"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeProduct = "Product"
)

// ProductMutation represents an operation that mutates the Product nodes in the graph.
type ProductMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	price         *float64
	addprice      *float64
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Product, error)
	predicates    []predicate.Product
}

var _ ent.Mutation = (*ProductMutation)(nil)

// productOption allows management of the mutation configuration using functional options.
type productOption func(*ProductMutation)

// newProductMutation creates new mutation for the Product entity.
func newProductMutation(c config, op Op, opts ...productOption) *ProductMutation {
	m := &ProductMutation{
		config:        c,
		op:            op,
		typ:           TypeProduct,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withProductID sets the ID field of the mutation.
func withProductID(id uuid.UUID) productOption {
	return func(m *ProductMutation) {
		var (
			err   error
			once  sync.Once
			value *Product
		)
		m.oldValue = func(ctx context.Context) (*Product, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Product.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withProduct sets the old Product of the mutation.
func withProduct(node *Product) productOption {
	return func(m *ProductMutation) {
		m.oldValue = func(context.Context) (*Product, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ProductMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ProductMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Product entities.
func (m *ProductMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ProductMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ProductMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Product.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetPrice sets the "price" field.
func (m *ProductMutation) SetPrice(f float64) {
	m.price = &f
	m.addprice = nil
}

// Price returns the value of the "price" field in the mutation.
func (m *ProductMutation) Price() (r float64, exists bool) {
	v := m.price
	if v == nil {
		return
	}
	return *v, true
}

// OldPrice returns the old "price" field's value of the Product entity.
// If the Product object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProductMutation) OldPrice(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPrice is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPrice requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPrice: %w", err)
	}
	return oldValue.Price, nil
}

// AddPrice adds f to the "price" field.
func (m *ProductMutation) AddPrice(f float64) {
	if m.addprice != nil {
		*m.addprice += f
	} else {
		m.addprice = &f
	}
}

// AddedPrice returns the value that was added to the "price" field in this mutation.
func (m *ProductMutation) AddedPrice() (r float64, exists bool) {
	v := m.addprice
	if v == nil {
		return
	}
	return *v, true
}

// ResetPrice resets all changes to the "price" field.
func (m *ProductMutation) ResetPrice() {
	m.price = nil
	m.addprice = nil
}

// SetName sets the "name" field.
func (m *ProductMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ProductMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Product entity.
// If the Product object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ProductMutation) OldName(ctx context.Context) (v string, err error) {
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
func (m *ProductMutation) ResetName() {
	m.name = nil
}

// Where appends a list predicates to the ProductMutation builder.
func (m *ProductMutation) Where(ps ...predicate.Product) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ProductMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ProductMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Product, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ProductMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ProductMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Product).
func (m *ProductMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ProductMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.price != nil {
		fields = append(fields, product.FieldPrice)
	}
	if m.name != nil {
		fields = append(fields, product.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ProductMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case product.FieldPrice:
		return m.Price()
	case product.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ProductMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case product.FieldPrice:
		return m.OldPrice(ctx)
	case product.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Product field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProductMutation) SetField(name string, value ent.Value) error {
	switch name {
	case product.FieldPrice:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPrice(v)
		return nil
	case product.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Product field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ProductMutation) AddedFields() []string {
	var fields []string
	if m.addprice != nil {
		fields = append(fields, product.FieldPrice)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ProductMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case product.FieldPrice:
		return m.AddedPrice()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ProductMutation) AddField(name string, value ent.Value) error {
	switch name {
	case product.FieldPrice:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPrice(v)
		return nil
	}
	return fmt.Errorf("unknown Product numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ProductMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ProductMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ProductMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Product nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ProductMutation) ResetField(name string) error {
	switch name {
	case product.FieldPrice:
		m.ResetPrice()
		return nil
	case product.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Product field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ProductMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ProductMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ProductMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ProductMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ProductMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ProductMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ProductMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Product unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ProductMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Product edge %s", name)
}
