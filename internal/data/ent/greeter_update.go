// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sample/internal/data/ent/greeter"
	"sample/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GreeterUpdate is the builder for updating Greeter entities.
type GreeterUpdate struct {
	config
	hooks    []Hook
	mutation *GreeterMutation
}

// Where appends a list predicates to the GreeterUpdate builder.
func (gu *GreeterUpdate) Where(ps ...predicate.Greeter) *GreeterUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GreeterUpdate) SetName(s string) *GreeterUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetCreatedAt sets the "created_at" field.
func (gu *GreeterUpdate) SetCreatedAt(t time.Time) *GreeterUpdate {
	gu.mutation.SetCreatedAt(t)
	return gu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gu *GreeterUpdate) SetNillableCreatedAt(t *time.Time) *GreeterUpdate {
	if t != nil {
		gu.SetCreatedAt(*t)
	}
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GreeterUpdate) SetUpdatedAt(t time.Time) *GreeterUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetDeletedAt sets the "deleted_at" field.
func (gu *GreeterUpdate) SetDeletedAt(t time.Time) *GreeterUpdate {
	gu.mutation.SetDeletedAt(t)
	return gu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gu *GreeterUpdate) SetNillableDeletedAt(t *time.Time) *GreeterUpdate {
	if t != nil {
		gu.SetDeletedAt(*t)
	}
	return gu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gu *GreeterUpdate) ClearDeletedAt() *GreeterUpdate {
	gu.mutation.ClearDeletedAt()
	return gu
}

// Mutation returns the GreeterMutation object of the builder.
func (gu *GreeterUpdate) Mutation() *GreeterMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GreeterUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	gu.defaults()
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GreeterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GreeterUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GreeterUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GreeterUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GreeterUpdate) defaults() {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		v := greeter.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
}

func (gu *GreeterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   greeter.Table,
			Columns: greeter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: greeter.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(greeter.FieldName, field.TypeString, value)
	}
	if value, ok := gu.mutation.CreatedAt(); ok {
		_spec.SetField(greeter.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(greeter.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gu.mutation.DeletedAt(); ok {
		_spec.SetField(greeter.FieldDeletedAt, field.TypeTime, value)
	}
	if gu.mutation.DeletedAtCleared() {
		_spec.ClearField(greeter.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{greeter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GreeterUpdateOne is the builder for updating a single Greeter entity.
type GreeterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GreeterMutation
}

// SetName sets the "name" field.
func (guo *GreeterUpdateOne) SetName(s string) *GreeterUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetCreatedAt sets the "created_at" field.
func (guo *GreeterUpdateOne) SetCreatedAt(t time.Time) *GreeterUpdateOne {
	guo.mutation.SetCreatedAt(t)
	return guo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (guo *GreeterUpdateOne) SetNillableCreatedAt(t *time.Time) *GreeterUpdateOne {
	if t != nil {
		guo.SetCreatedAt(*t)
	}
	return guo
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GreeterUpdateOne) SetUpdatedAt(t time.Time) *GreeterUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetDeletedAt sets the "deleted_at" field.
func (guo *GreeterUpdateOne) SetDeletedAt(t time.Time) *GreeterUpdateOne {
	guo.mutation.SetDeletedAt(t)
	return guo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guo *GreeterUpdateOne) SetNillableDeletedAt(t *time.Time) *GreeterUpdateOne {
	if t != nil {
		guo.SetDeletedAt(*t)
	}
	return guo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (guo *GreeterUpdateOne) ClearDeletedAt() *GreeterUpdateOne {
	guo.mutation.ClearDeletedAt()
	return guo
}

// Mutation returns the GreeterMutation object of the builder.
func (guo *GreeterUpdateOne) Mutation() *GreeterMutation {
	return guo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GreeterUpdateOne) Select(field string, fields ...string) *GreeterUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Greeter entity.
func (guo *GreeterUpdateOne) Save(ctx context.Context) (*Greeter, error) {
	var (
		err  error
		node *Greeter
	)
	guo.defaults()
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GreeterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, guo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Greeter)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GreeterMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GreeterUpdateOne) SaveX(ctx context.Context) *Greeter {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GreeterUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GreeterUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GreeterUpdateOne) defaults() {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		v := greeter.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
}

func (guo *GreeterUpdateOne) sqlSave(ctx context.Context) (_node *Greeter, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   greeter.Table,
			Columns: greeter.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: greeter.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Greeter.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, greeter.FieldID)
		for _, f := range fields {
			if !greeter.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != greeter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(greeter.FieldName, field.TypeString, value)
	}
	if value, ok := guo.mutation.CreatedAt(); ok {
		_spec.SetField(greeter.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(greeter.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := guo.mutation.DeletedAt(); ok {
		_spec.SetField(greeter.FieldDeletedAt, field.TypeTime, value)
	}
	if guo.mutation.DeletedAtCleared() {
		_spec.ClearField(greeter.FieldDeletedAt, field.TypeTime)
	}
	_node = &Greeter{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{greeter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}