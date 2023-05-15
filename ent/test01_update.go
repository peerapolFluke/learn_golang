// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/pumy2517/ginent/ent/predicate"
	"github.com/pumy2517/ginent/ent/test01"
)

// Test01Update is the builder for updating Test01 entities.
type Test01Update struct {
	config
	hooks    []Hook
	mutation *Test01Mutation
}

// Where appends a list predicates to the Test01Update builder.
func (t *Test01Update) Where(ps ...predicate.Test01) *Test01Update {
	t.mutation.Where(ps...)
	return t
}

// SetText sets the "text" field.
func (t *Test01Update) SetText(s string) *Test01Update {
	t.mutation.SetText(s)
	return t
}

// SetStatus sets the "status" field.
func (t *Test01Update) SetStatus(value test01.Status) *Test01Update {
	t.mutation.SetStatus(value)
	return t
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (t *Test01Update) SetNillableStatus(value *test01.Status) *Test01Update {
	if value != nil {
		t.SetStatus(*value)
	}
	return t
}

// SetPriority sets the "priority" field.
func (t *Test01Update) SetPriority(i int) *Test01Update {
	t.mutation.ResetPriority()
	t.mutation.SetPriority(i)
	return t
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (t *Test01Update) SetNillablePriority(i *int) *Test01Update {
	if i != nil {
		t.SetPriority(*i)
	}
	return t
}

// AddPriority adds i to the "priority" field.
func (t *Test01Update) AddPriority(i int) *Test01Update {
	t.mutation.AddPriority(i)
	return t
}

// SetParentID sets the "parent" edge to the Test01 entity by ID.
func (t *Test01Update) SetParentID(id int) *Test01Update {
	t.mutation.SetParentID(id)
	return t
}

// SetNillableParentID sets the "parent" edge to the Test01 entity by ID if the given value is not nil.
func (t *Test01Update) SetNillableParentID(id *int) *Test01Update {
	if id != nil {
		t = t.SetParentID(*id)
	}
	return t
}

// SetParent sets the "parent" edge to the Test01 entity.
func (t *Test01Update) SetParent(v *Test01) *Test01Update {
	return t.SetParentID(v.ID)
}

// AddChildIDs adds the "children" edge to the Test01 entity by IDs.
func (t *Test01Update) AddChildIDs(ids ...int) *Test01Update {
	t.mutation.AddChildIDs(ids...)
	return t
}

// AddChildren adds the "children" edges to the Test01 entity.
func (t *Test01Update) AddChildren(v ...*Test01) *Test01Update {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return t.AddChildIDs(ids...)
}

// Mutation returns the Test01Mutation object of the builder.
func (t *Test01Update) Mutation() *Test01Mutation {
	return t.mutation
}

// ClearParent clears the "parent" edge to the Test01 entity.
func (t *Test01Update) ClearParent() *Test01Update {
	t.mutation.ClearParent()
	return t
}

// ClearChildren clears all "children" edges to the Test01 entity.
func (t *Test01Update) ClearChildren() *Test01Update {
	t.mutation.ClearChildren()
	return t
}

// RemoveChildIDs removes the "children" edge to Test01 entities by IDs.
func (t *Test01Update) RemoveChildIDs(ids ...int) *Test01Update {
	t.mutation.RemoveChildIDs(ids...)
	return t
}

// RemoveChildren removes "children" edges to Test01 entities.
func (t *Test01Update) RemoveChildren(v ...*Test01) *Test01Update {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return t.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (t *Test01Update) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, t.sqlSave, t.mutation, t.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (t *Test01Update) SaveX(ctx context.Context) int {
	affected, err := t.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (t *Test01Update) Exec(ctx context.Context) error {
	_, err := t.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (t *Test01Update) ExecX(ctx context.Context) {
	if err := t.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (t *Test01Update) check() error {
	if v, ok := t.mutation.Text(); ok {
		if err := test01.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Test01.text": %w`, err)}
		}
	}
	if v, ok := t.mutation.Status(); ok {
		if err := test01.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Test01.status": %w`, err)}
		}
	}
	return nil
}

func (t *Test01Update) sqlSave(ctx context.Context) (n int, err error) {
	if err := t.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(test01.Table, test01.Columns, sqlgraph.NewFieldSpec(test01.FieldID, field.TypeInt))
	if ps := t.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := t.mutation.Text(); ok {
		_spec.SetField(test01.FieldText, field.TypeString, value)
	}
	if value, ok := t.mutation.Status(); ok {
		_spec.SetField(test01.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := t.mutation.Priority(); ok {
		_spec.SetField(test01.FieldPriority, field.TypeInt, value)
	}
	if value, ok := t.mutation.AddedPriority(); ok {
		_spec.AddField(test01.FieldPriority, field.TypeInt, value)
	}
	if t.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if t.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := t.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !t.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, t.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{test01.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	t.mutation.done = true
	return n, nil
}

// Test01UpdateOne is the builder for updating a single Test01 entity.
type Test01UpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *Test01Mutation
}

// SetText sets the "text" field.
func (to *Test01UpdateOne) SetText(s string) *Test01UpdateOne {
	to.mutation.SetText(s)
	return to
}

// SetStatus sets the "status" field.
func (to *Test01UpdateOne) SetStatus(t test01.Status) *Test01UpdateOne {
	to.mutation.SetStatus(t)
	return to
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (to *Test01UpdateOne) SetNillableStatus(t *test01.Status) *Test01UpdateOne {
	if t != nil {
		to.SetStatus(*t)
	}
	return to
}

// SetPriority sets the "priority" field.
func (to *Test01UpdateOne) SetPriority(i int) *Test01UpdateOne {
	to.mutation.ResetPriority()
	to.mutation.SetPriority(i)
	return to
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (to *Test01UpdateOne) SetNillablePriority(i *int) *Test01UpdateOne {
	if i != nil {
		to.SetPriority(*i)
	}
	return to
}

// AddPriority adds i to the "priority" field.
func (to *Test01UpdateOne) AddPriority(i int) *Test01UpdateOne {
	to.mutation.AddPriority(i)
	return to
}

// SetParentID sets the "parent" edge to the Test01 entity by ID.
func (to *Test01UpdateOne) SetParentID(id int) *Test01UpdateOne {
	to.mutation.SetParentID(id)
	return to
}

// SetNillableParentID sets the "parent" edge to the Test01 entity by ID if the given value is not nil.
func (to *Test01UpdateOne) SetNillableParentID(id *int) *Test01UpdateOne {
	if id != nil {
		to = to.SetParentID(*id)
	}
	return to
}

// SetParent sets the "parent" edge to the Test01 entity.
func (to *Test01UpdateOne) SetParent(t *Test01) *Test01UpdateOne {
	return to.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Test01 entity by IDs.
func (to *Test01UpdateOne) AddChildIDs(ids ...int) *Test01UpdateOne {
	to.mutation.AddChildIDs(ids...)
	return to
}

// AddChildren adds the "children" edges to the Test01 entity.
func (to *Test01UpdateOne) AddChildren(t ...*Test01) *Test01UpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return to.AddChildIDs(ids...)
}

// Mutation returns the Test01Mutation object of the builder.
func (to *Test01UpdateOne) Mutation() *Test01Mutation {
	return to.mutation
}

// ClearParent clears the "parent" edge to the Test01 entity.
func (to *Test01UpdateOne) ClearParent() *Test01UpdateOne {
	to.mutation.ClearParent()
	return to
}

// ClearChildren clears all "children" edges to the Test01 entity.
func (to *Test01UpdateOne) ClearChildren() *Test01UpdateOne {
	to.mutation.ClearChildren()
	return to
}

// RemoveChildIDs removes the "children" edge to Test01 entities by IDs.
func (to *Test01UpdateOne) RemoveChildIDs(ids ...int) *Test01UpdateOne {
	to.mutation.RemoveChildIDs(ids...)
	return to
}

// RemoveChildren removes "children" edges to Test01 entities.
func (to *Test01UpdateOne) RemoveChildren(t ...*Test01) *Test01UpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return to.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the Test01Update builder.
func (to *Test01UpdateOne) Where(ps ...predicate.Test01) *Test01UpdateOne {
	to.mutation.Where(ps...)
	return to
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (to *Test01UpdateOne) Select(field string, fields ...string) *Test01UpdateOne {
	to.fields = append([]string{field}, fields...)
	return to
}

// Save executes the query and returns the updated Test01 entity.
func (to *Test01UpdateOne) Save(ctx context.Context) (*Test01, error) {
	return withHooks(ctx, to.sqlSave, to.mutation, to.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (to *Test01UpdateOne) SaveX(ctx context.Context) *Test01 {
	node, err := to.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (to *Test01UpdateOne) Exec(ctx context.Context) error {
	_, err := to.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (to *Test01UpdateOne) ExecX(ctx context.Context) {
	if err := to.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (to *Test01UpdateOne) check() error {
	if v, ok := to.mutation.Text(); ok {
		if err := test01.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Test01.text": %w`, err)}
		}
	}
	if v, ok := to.mutation.Status(); ok {
		if err := test01.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Test01.status": %w`, err)}
		}
	}
	return nil
}

func (to *Test01UpdateOne) sqlSave(ctx context.Context) (_node *Test01, err error) {
	if err := to.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(test01.Table, test01.Columns, sqlgraph.NewFieldSpec(test01.FieldID, field.TypeInt))
	id, ok := to.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Test01.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := to.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, test01.FieldID)
		for _, f := range fields {
			if !test01.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != test01.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := to.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := to.mutation.Text(); ok {
		_spec.SetField(test01.FieldText, field.TypeString, value)
	}
	if value, ok := to.mutation.Status(); ok {
		_spec.SetField(test01.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := to.mutation.Priority(); ok {
		_spec.SetField(test01.FieldPriority, field.TypeInt, value)
	}
	if value, ok := to.mutation.AddedPriority(); ok {
		_spec.AddField(test01.FieldPriority, field.TypeInt, value)
	}
	if to.mutation.ParentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := to.mutation.ParentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if to.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := to.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !to.mutation.ChildrenCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := to.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Test01{config: to.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, to.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{test01.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	to.mutation.done = true
	return _node, nil
}
