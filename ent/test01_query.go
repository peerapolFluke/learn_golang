// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"ginent/ent/predicate"
	"ginent/ent/test01"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// Test01Query is the builder for querying Test01 entities.
type Test01Query struct {
	config
	ctx               *QueryContext
	order             []test01.OrderOption
	inters            []Interceptor
	predicates        []predicate.Test01
	withParent        *Test01Query
	withChildren      *Test01Query
	withFKs           bool
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*Test01) error
	withNamedChildren map[string]*Test01Query
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the Test01Query builder.
func (t *Test01Query) Where(ps ...predicate.Test01) *Test01Query {
	t.predicates = append(t.predicates, ps...)
	return t
}

// Limit the number of records to be returned by this query.
func (t *Test01Query) Limit(limit int) *Test01Query {
	t.ctx.Limit = &limit
	return t
}

// Offset to start from.
func (t *Test01Query) Offset(offset int) *Test01Query {
	t.ctx.Offset = &offset
	return t
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (t *Test01Query) Unique(unique bool) *Test01Query {
	t.ctx.Unique = &unique
	return t
}

// Order specifies how the records should be ordered.
func (t *Test01Query) Order(o ...test01.OrderOption) *Test01Query {
	t.order = append(t.order, o...)
	return t
}

// QueryParent chains the current query on the "parent" edge.
func (t *Test01Query) QueryParent() *Test01Query {
	query := (&Test01Client{config: t.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := t.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := t.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(test01.Table, test01.FieldID, selector),
			sqlgraph.To(test01.Table, test01.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, test01.ParentTable, test01.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(t.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (t *Test01Query) QueryChildren() *Test01Query {
	query := (&Test01Client{config: t.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := t.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := t.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(test01.Table, test01.FieldID, selector),
			sqlgraph.To(test01.Table, test01.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, test01.ChildrenTable, test01.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(t.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Test01 entity from the query.
// Returns a *NotFoundError when no Test01 was found.
func (t *Test01Query) First(ctx context.Context) (*Test01, error) {
	nodes, err := t.Limit(1).All(setContextOp(ctx, t.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{test01.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (t *Test01Query) FirstX(ctx context.Context) *Test01 {
	node, err := t.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Test01 ID from the query.
// Returns a *NotFoundError when no Test01 ID was found.
func (t *Test01Query) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = t.Limit(1).IDs(setContextOp(ctx, t.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{test01.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (t *Test01Query) FirstIDX(ctx context.Context) int {
	id, err := t.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Test01 entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Test01 entity is found.
// Returns a *NotFoundError when no Test01 entities are found.
func (t *Test01Query) Only(ctx context.Context) (*Test01, error) {
	nodes, err := t.Limit(2).All(setContextOp(ctx, t.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{test01.Label}
	default:
		return nil, &NotSingularError{test01.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (t *Test01Query) OnlyX(ctx context.Context) *Test01 {
	node, err := t.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Test01 ID in the query.
// Returns a *NotSingularError when more than one Test01 ID is found.
// Returns a *NotFoundError when no entities are found.
func (t *Test01Query) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = t.Limit(2).IDs(setContextOp(ctx, t.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{test01.Label}
	default:
		err = &NotSingularError{test01.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (t *Test01Query) OnlyIDX(ctx context.Context) int {
	id, err := t.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Test01s.
func (t *Test01Query) All(ctx context.Context) ([]*Test01, error) {
	ctx = setContextOp(ctx, t.ctx, "All")
	if err := t.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Test01, *Test01Query]()
	return withInterceptors[[]*Test01](ctx, t, qr, t.inters)
}

// AllX is like All, but panics if an error occurs.
func (t *Test01Query) AllX(ctx context.Context) []*Test01 {
	nodes, err := t.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Test01 IDs.
func (t *Test01Query) IDs(ctx context.Context) (ids []int, err error) {
	if t.ctx.Unique == nil && t.path != nil {
		t.Unique(true)
	}
	ctx = setContextOp(ctx, t.ctx, "IDs")
	if err = t.Select(test01.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (t *Test01Query) IDsX(ctx context.Context) []int {
	ids, err := t.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (t *Test01Query) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, t.ctx, "Count")
	if err := t.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, t, querierCount[*Test01Query](), t.inters)
}

// CountX is like Count, but panics if an error occurs.
func (t *Test01Query) CountX(ctx context.Context) int {
	count, err := t.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (t *Test01Query) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, t.ctx, "Exist")
	switch _, err := t.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (t *Test01Query) ExistX(ctx context.Context) bool {
	exist, err := t.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the Test01Query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (t *Test01Query) Clone() *Test01Query {
	if t == nil {
		return nil
	}
	return &Test01Query{
		config:       t.config,
		ctx:          t.ctx.Clone(),
		order:        append([]test01.OrderOption{}, t.order...),
		inters:       append([]Interceptor{}, t.inters...),
		predicates:   append([]predicate.Test01{}, t.predicates...),
		withParent:   t.withParent.Clone(),
		withChildren: t.withChildren.Clone(),
		// clone intermediate query.
		sql:  t.sql.Clone(),
		path: t.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (t *Test01Query) WithParent(opts ...func(*Test01Query)) *Test01Query {
	query := (&Test01Client{config: t.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	t.withParent = query
	return t
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (t *Test01Query) WithChildren(opts ...func(*Test01Query)) *Test01Query {
	query := (&Test01Client{config: t.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	t.withChildren = query
	return t
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Test01.Query().
//		GroupBy(test01.FieldText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (t *Test01Query) GroupBy(field string, fields ...string) *Test01GroupBy {
	t.ctx.Fields = append([]string{field}, fields...)
	grbuild := &Test01GroupBy{build: t}
	grbuild.flds = &t.ctx.Fields
	grbuild.label = test01.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.Test01.Query().
//		Select(test01.FieldText).
//		Scan(ctx, &v)
func (t *Test01Query) Select(fields ...string) *Test01Select {
	t.ctx.Fields = append(t.ctx.Fields, fields...)
	sbuild := &Test01Select{Test01Query: t}
	sbuild.label = test01.Label
	sbuild.flds, sbuild.scan = &t.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a Test01Select configured with the given aggregations.
func (t *Test01Query) Aggregate(fns ...AggregateFunc) *Test01Select {
	return t.Select().Aggregate(fns...)
}

func (t *Test01Query) prepareQuery(ctx context.Context) error {
	for _, inter := range t.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, t); err != nil {
				return err
			}
		}
	}
	for _, f := range t.ctx.Fields {
		if !test01.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if t.path != nil {
		prev, err := t.path(ctx)
		if err != nil {
			return err
		}
		t.sql = prev
	}
	return nil
}

func (t *Test01Query) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Test01, error) {
	var (
		nodes       = []*Test01{}
		withFKs     = t.withFKs
		_spec       = t.querySpec()
		loadedTypes = [2]bool{
			t.withParent != nil,
			t.withChildren != nil,
		}
	)
	if t.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, test01.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Test01).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Test01{config: t.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(t.modifiers) > 0 {
		_spec.Modifiers = t.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, t.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := t.withParent; query != nil {
		if err := t.loadParent(ctx, query, nodes, nil,
			func(n *Test01, e *Test01) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := t.withChildren; query != nil {
		if err := t.loadChildren(ctx, query, nodes,
			func(n *Test01) { n.Edges.Children = []*Test01{} },
			func(n *Test01, e *Test01) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range t.withNamedChildren {
		if err := t.loadChildren(ctx, query, nodes,
			func(n *Test01) { n.appendNamedChildren(name) },
			func(n *Test01, e *Test01) { n.appendNamedChildren(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range t.loadTotal {
		if err := t.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (t *Test01Query) loadParent(ctx context.Context, query *Test01Query, nodes []*Test01, init func(*Test01), assign func(*Test01, *Test01)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Test01)
	for i := range nodes {
		if nodes[i].test01_children == nil {
			continue
		}
		fk := *nodes[i].test01_children
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(test01.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "test01_children" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (t *Test01Query) loadChildren(ctx context.Context, query *Test01Query, nodes []*Test01, init func(*Test01), assign func(*Test01, *Test01)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Test01)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Test01(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(test01.ChildrenColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.test01_children
		if fk == nil {
			return fmt.Errorf(`foreign-key "test01_children" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "test01_children" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (t *Test01Query) sqlCount(ctx context.Context) (int, error) {
	_spec := t.querySpec()
	if len(t.modifiers) > 0 {
		_spec.Modifiers = t.modifiers
	}
	_spec.Node.Columns = t.ctx.Fields
	if len(t.ctx.Fields) > 0 {
		_spec.Unique = t.ctx.Unique != nil && *t.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, t.driver, _spec)
}

func (t *Test01Query) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(test01.Table, test01.Columns, sqlgraph.NewFieldSpec(test01.FieldID, field.TypeInt))
	_spec.From = t.sql
	if unique := t.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if t.path != nil {
		_spec.Unique = true
	}
	if fields := t.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, test01.FieldID)
		for i := range fields {
			if fields[i] != test01.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := t.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := t.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := t.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := t.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (t *Test01Query) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(t.driver.Dialect())
	t1 := builder.Table(test01.Table)
	columns := t.ctx.Fields
	if len(columns) == 0 {
		columns = test01.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if t.sql != nil {
		selector = t.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if t.ctx.Unique != nil && *t.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range t.predicates {
		p(selector)
	}
	for _, p := range t.order {
		p(selector)
	}
	if offset := t.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := t.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedChildren tells the query-builder to eager-load the nodes that are connected to the "children"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (t *Test01Query) WithNamedChildren(name string, opts ...func(*Test01Query)) *Test01Query {
	query := (&Test01Client{config: t.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if t.withNamedChildren == nil {
		t.withNamedChildren = make(map[string]*Test01Query)
	}
	t.withNamedChildren[name] = query
	return t
}

// Test01GroupBy is the group-by builder for Test01 entities.
type Test01GroupBy struct {
	selector
	build *Test01Query
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tb *Test01GroupBy) Aggregate(fns ...AggregateFunc) *Test01GroupBy {
	tb.fns = append(tb.fns, fns...)
	return tb
}

// Scan applies the selector query and scans the result into the given value.
func (tb *Test01GroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tb.build.ctx, "GroupBy")
	if err := tb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*Test01Query, *Test01GroupBy](ctx, tb.build, tb, tb.build.inters, v)
}

func (tb *Test01GroupBy) sqlScan(ctx context.Context, root *Test01Query, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tb.fns))
	for _, fn := range tb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tb.flds)+len(tb.fns))
		for _, f := range *tb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Test01Select is the builder for selecting fields of Test01 entities.
type Test01Select struct {
	*Test01Query
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (t *Test01Select) Aggregate(fns ...AggregateFunc) *Test01Select {
	t.fns = append(t.fns, fns...)
	return t
}

// Scan applies the selector query and scans the result into the given value.
func (t *Test01Select) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, t.ctx, "Select")
	if err := t.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*Test01Query, *Test01Select](ctx, t.Test01Query, t, t.inters, v)
}

func (t *Test01Select) sqlScan(ctx context.Context, root *Test01Query, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(t.fns))
	for _, fn := range t.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*t.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := t.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
