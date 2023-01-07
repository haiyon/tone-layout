// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sample/internal/data/ent/greeter"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GreeterCreate is the builder for creating a Greeter entity.
type GreeterCreate struct {
	config
	mutation *GreeterMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (gc *GreeterCreate) SetName(s string) *GreeterCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetCreatedAt sets the "created_at" field.
func (gc *GreeterCreate) SetCreatedAt(t time.Time) *GreeterCreate {
	gc.mutation.SetCreatedAt(t)
	return gc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gc *GreeterCreate) SetNillableCreatedAt(t *time.Time) *GreeterCreate {
	if t != nil {
		gc.SetCreatedAt(*t)
	}
	return gc
}

// SetUpdatedAt sets the "updated_at" field.
func (gc *GreeterCreate) SetUpdatedAt(t time.Time) *GreeterCreate {
	gc.mutation.SetUpdatedAt(t)
	return gc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gc *GreeterCreate) SetNillableUpdatedAt(t *time.Time) *GreeterCreate {
	if t != nil {
		gc.SetUpdatedAt(*t)
	}
	return gc
}

// SetDeletedAt sets the "deleted_at" field.
func (gc *GreeterCreate) SetDeletedAt(t time.Time) *GreeterCreate {
	gc.mutation.SetDeletedAt(t)
	return gc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gc *GreeterCreate) SetNillableDeletedAt(t *time.Time) *GreeterCreate {
	if t != nil {
		gc.SetDeletedAt(*t)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GreeterCreate) SetID(s string) *GreeterCreate {
	gc.mutation.SetID(s)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GreeterCreate) SetNillableID(s *string) *GreeterCreate {
	if s != nil {
		gc.SetID(*s)
	}
	return gc
}

// Mutation returns the GreeterMutation object of the builder.
func (gc *GreeterCreate) Mutation() *GreeterMutation {
	return gc.mutation
}

// Save creates the Greeter in the database.
func (gc *GreeterCreate) Save(ctx context.Context) (*Greeter, error) {
	var (
		err  error
		node *Greeter
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GreeterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			if node, err = gc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			if gc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (gc *GreeterCreate) SaveX(ctx context.Context) *Greeter {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GreeterCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GreeterCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GreeterCreate) defaults() {
	if _, ok := gc.mutation.CreatedAt(); !ok {
		v := greeter.DefaultCreatedAt()
		gc.mutation.SetCreatedAt(v)
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		v := greeter.DefaultUpdatedAt()
		gc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gc.mutation.ID(); !ok {
		v := greeter.DefaultID()
		gc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GreeterCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Greeter.name"`)}
	}
	if _, ok := gc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Greeter.created_at"`)}
	}
	if _, ok := gc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Greeter.updated_at"`)}
	}
	return nil
}

func (gc *GreeterCreate) sqlSave(ctx context.Context) (*Greeter, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Greeter.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (gc *GreeterCreate) createSpec() (*Greeter, *sqlgraph.CreateSpec) {
	var (
		_node = &Greeter{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: greeter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: greeter.FieldID,
			},
		}
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.SetField(greeter.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gc.mutation.CreatedAt(); ok {
		_spec.SetField(greeter.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gc.mutation.UpdatedAt(); ok {
		_spec.SetField(greeter.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gc.mutation.DeletedAt(); ok {
		_spec.SetField(greeter.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	return _node, _spec
}

// GreeterCreateBulk is the builder for creating many Greeter entities in bulk.
type GreeterCreateBulk struct {
	config
	builders []*GreeterCreate
}

// Save creates the Greeter entities in the database.
func (gcb *GreeterCreateBulk) Save(ctx context.Context) ([]*Greeter, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Greeter, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GreeterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GreeterCreateBulk) SaveX(ctx context.Context) []*Greeter {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GreeterCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GreeterCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}