// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/pumy2517/ginent/ent/test01"
)

// Test01Create is the builder for creating a Test01 entity.
type Test01Create struct {
	config
	mutation *Test01Mutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (t *Test01Create) SetText(s string) *Test01Create {
	t.mutation.SetText(s)
	return t
}

// SetCreatedAt sets the "created_at" field.
func (t *Test01Create) SetCreatedAt(value time.Time) *Test01Create {
	t.mutation.SetCreatedAt(value)
	return t
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (t *Test01Create) SetNillableCreatedAt(value *time.Time) *Test01Create {
	if value != nil {
		t.SetCreatedAt(*value)
	}
	return t
}

// SetUpdateAt sets the "update_at" field.
func (t *Test01Create) SetUpdateAt(value time.Time) *Test01Create {
	t.mutation.SetUpdateAt(value)
	return t
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (t *Test01Create) SetNillableUpdateAt(value *time.Time) *Test01Create {
	if value != nil {
		t.SetUpdateAt(*value)
	}
	return t
}

// SetStatus sets the "status" field.
func (t *Test01Create) SetStatus(value test01.Status) *Test01Create {
	t.mutation.SetStatus(value)
	return t
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (t *Test01Create) SetNillableStatus(value *test01.Status) *Test01Create {
	if value != nil {
		t.SetStatus(*value)
	}
	return t
}

// SetPriority sets the "priority" field.
func (t *Test01Create) SetPriority(i int) *Test01Create {
	t.mutation.SetPriority(i)
	return t
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (t *Test01Create) SetNillablePriority(i *int) *Test01Create {
	if i != nil {
		t.SetPriority(*i)
	}
	return t
}

// SetParentID sets the "parent" edge to the Test01 entity by ID.
func (t *Test01Create) SetParentID(id int) *Test01Create {
	t.mutation.SetParentID(id)
	return t
}

// SetNillableParentID sets the "parent" edge to the Test01 entity by ID if the given value is not nil.
func (t *Test01Create) SetNillableParentID(id *int) *Test01Create {
	if id != nil {
		t = t.SetParentID(*id)
	}
	return t
}

// SetParent sets the "parent" edge to the Test01 entity.
func (t *Test01Create) SetParent(v *Test01) *Test01Create {
	return t.SetParentID(v.ID)
}

// AddChildIDs adds the "children" edge to the Test01 entity by IDs.
func (t *Test01Create) AddChildIDs(ids ...int) *Test01Create {
	t.mutation.AddChildIDs(ids...)
	return t
}

// AddChildren adds the "children" edges to the Test01 entity.
func (t *Test01Create) AddChildren(v ...*Test01) *Test01Create {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return t.AddChildIDs(ids...)
}

// Mutation returns the Test01Mutation object of the builder.
func (t *Test01Create) Mutation() *Test01Mutation {
	return t.mutation
}

// Save creates the Test01 in the database.
func (t *Test01Create) Save(ctx context.Context) (*Test01, error) {
	t.defaults()
	return withHooks(ctx, t.sqlSave, t.mutation, t.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (t *Test01Create) SaveX(ctx context.Context) *Test01 {
	v, err := t.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (t *Test01Create) Exec(ctx context.Context) error {
	_, err := t.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (t *Test01Create) ExecX(ctx context.Context) {
	if err := t.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (t *Test01Create) defaults() {
	if _, ok := t.mutation.CreatedAt(); !ok {
		v := test01.DefaultCreatedAt()
		t.mutation.SetCreatedAt(v)
	}
	if _, ok := t.mutation.UpdateAt(); !ok {
		v := test01.DefaultUpdateAt()
		t.mutation.SetUpdateAt(v)
	}
	if _, ok := t.mutation.Status(); !ok {
		v := test01.DefaultStatus
		t.mutation.SetStatus(v)
	}
	if _, ok := t.mutation.Priority(); !ok {
		v := test01.DefaultPriority
		t.mutation.SetPriority(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (t *Test01Create) check() error {
	if _, ok := t.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Test01.text"`)}
	}
	if v, ok := t.mutation.Text(); ok {
		if err := test01.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Test01.text": %w`, err)}
		}
	}
	if _, ok := t.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Test01.created_at"`)}
	}
	if _, ok := t.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "Test01.update_at"`)}
	}
	if _, ok := t.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Test01.status"`)}
	}
	if v, ok := t.mutation.Status(); ok {
		if err := test01.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Test01.status": %w`, err)}
		}
	}
	if _, ok := t.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "Test01.priority"`)}
	}
	return nil
}

func (t *Test01Create) sqlSave(ctx context.Context) (*Test01, error) {
	if err := t.check(); err != nil {
		return nil, err
	}
	_node, _spec := t.createSpec()
	if err := sqlgraph.CreateNode(ctx, t.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	t.mutation.id = &_node.ID
	t.mutation.done = true
	return _node, nil
}

func (t *Test01Create) createSpec() (*Test01, *sqlgraph.CreateSpec) {
	var (
		_node = &Test01{config: t.config}
		_spec = sqlgraph.NewCreateSpec(test01.Table, sqlgraph.NewFieldSpec(test01.FieldID, field.TypeInt))
	)
	if value, ok := t.mutation.Text(); ok {
		_spec.SetField(test01.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := t.mutation.CreatedAt(); ok {
		_spec.SetField(test01.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := t.mutation.UpdateAt(); ok {
		_spec.SetField(test01.FieldUpdateAt, field.TypeTime, value)
		_node.UpdateAt = value
	}
	if value, ok := t.mutation.Status(); ok {
		_spec.SetField(test01.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := t.mutation.Priority(); ok {
		_spec.SetField(test01.FieldPriority, field.TypeInt, value)
		_node.Priority = value
	}
	if nodes := t.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   test01.ParentTable,
			Columns: []string{test01.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(test01.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.test01_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := t.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test01.ChildrenTable,
			Columns: []string{test01.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(test01.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// Test01CreateBulk is the builder for creating many Test01 entities in bulk.
type Test01CreateBulk struct {
	config
	builders []*Test01Create
}

// Save creates the Test01 entities in the database.
func (tb *Test01CreateBulk) Save(ctx context.Context) ([]*Test01, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tb.builders))
	nodes := make([]*Test01, len(tb.builders))
	mutators := make([]Mutator, len(tb.builders))
	for i := range tb.builders {
		func(i int, root context.Context) {
			builder := tb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*Test01Mutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, tb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tb *Test01CreateBulk) SaveX(ctx context.Context) []*Test01 {
	v, err := tb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tb *Test01CreateBulk) Exec(ctx context.Context) error {
	_, err := tb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tb *Test01CreateBulk) ExecX(ctx context.Context) {
	if err := tb.Exec(ctx); err != nil {
		panic(err)
	}
}
