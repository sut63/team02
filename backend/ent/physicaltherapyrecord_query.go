// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/physicaltherapyrecord"
	"github.com/to63/app/ent/physicaltherapyroom"
	"github.com/to63/app/ent/predicate"
	"github.com/to63/app/ent/status"
)

// PhysicaltherapyrecordQuery is the builder for querying Physicaltherapyrecord entities.
type PhysicaltherapyrecordQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Physicaltherapyrecord
	// eager-loading edges.
	withPersonnel           *PersonnelQuery
	withPatient             *PatientQuery
	withPhysicaltherapyroom *PhysicaltherapyroomQuery
	withStatus              *StatusQuery
	withFKs                 bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PhysicaltherapyrecordQuery builder.
func (pq *PhysicaltherapyrecordQuery) Where(ps ...predicate.Physicaltherapyrecord) *PhysicaltherapyrecordQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PhysicaltherapyrecordQuery) Limit(limit int) *PhysicaltherapyrecordQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PhysicaltherapyrecordQuery) Offset(offset int) *PhysicaltherapyrecordQuery {
	pq.offset = &offset
	return pq
}

// Order adds an order step to the query.
func (pq *PhysicaltherapyrecordQuery) Order(o ...OrderFunc) *PhysicaltherapyrecordQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryPersonnel chains the current query on the "personnel" edge.
func (pq *PhysicaltherapyrecordQuery) QueryPersonnel() *PersonnelQuery {
	query := &PersonnelQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(physicaltherapyrecord.Table, physicaltherapyrecord.FieldID, selector),
			sqlgraph.To(personnel.Table, personnel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, physicaltherapyrecord.PersonnelTable, physicaltherapyrecord.PersonnelColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPatient chains the current query on the "patient" edge.
func (pq *PhysicaltherapyrecordQuery) QueryPatient() *PatientQuery {
	query := &PatientQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(physicaltherapyrecord.Table, physicaltherapyrecord.FieldID, selector),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, physicaltherapyrecord.PatientTable, physicaltherapyrecord.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPhysicaltherapyroom chains the current query on the "physicaltherapyroom" edge.
func (pq *PhysicaltherapyrecordQuery) QueryPhysicaltherapyroom() *PhysicaltherapyroomQuery {
	query := &PhysicaltherapyroomQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(physicaltherapyrecord.Table, physicaltherapyrecord.FieldID, selector),
			sqlgraph.To(physicaltherapyroom.Table, physicaltherapyroom.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, physicaltherapyrecord.PhysicaltherapyroomTable, physicaltherapyrecord.PhysicaltherapyroomColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStatus chains the current query on the "status" edge.
func (pq *PhysicaltherapyrecordQuery) QueryStatus() *StatusQuery {
	query := &StatusQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(physicaltherapyrecord.Table, physicaltherapyrecord.FieldID, selector),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, physicaltherapyrecord.StatusTable, physicaltherapyrecord.StatusColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Physicaltherapyrecord entity from the query.
// Returns a *NotFoundError when no Physicaltherapyrecord was found.
func (pq *PhysicaltherapyrecordQuery) First(ctx context.Context) (*Physicaltherapyrecord, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{physicaltherapyrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PhysicaltherapyrecordQuery) FirstX(ctx context.Context) *Physicaltherapyrecord {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Physicaltherapyrecord ID from the query.
// Returns a *NotFoundError when no Physicaltherapyrecord ID was found.
func (pq *PhysicaltherapyrecordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{physicaltherapyrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PhysicaltherapyrecordQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Physicaltherapyrecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Physicaltherapyrecord entity is not found.
// Returns a *NotFoundError when no Physicaltherapyrecord entities are found.
func (pq *PhysicaltherapyrecordQuery) Only(ctx context.Context) (*Physicaltherapyrecord, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{physicaltherapyrecord.Label}
	default:
		return nil, &NotSingularError{physicaltherapyrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PhysicaltherapyrecordQuery) OnlyX(ctx context.Context) *Physicaltherapyrecord {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Physicaltherapyrecord ID in the query.
// Returns a *NotSingularError when exactly one Physicaltherapyrecord ID is not found.
// Returns a *NotFoundError when no entities are found.
func (pq *PhysicaltherapyrecordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{physicaltherapyrecord.Label}
	default:
		err = &NotSingularError{physicaltherapyrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PhysicaltherapyrecordQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Physicaltherapyrecords.
func (pq *PhysicaltherapyrecordQuery) All(ctx context.Context) ([]*Physicaltherapyrecord, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PhysicaltherapyrecordQuery) AllX(ctx context.Context) []*Physicaltherapyrecord {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Physicaltherapyrecord IDs.
func (pq *PhysicaltherapyrecordQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(physicaltherapyrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PhysicaltherapyrecordQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PhysicaltherapyrecordQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PhysicaltherapyrecordQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PhysicaltherapyrecordQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PhysicaltherapyrecordQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PhysicaltherapyrecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PhysicaltherapyrecordQuery) Clone() *PhysicaltherapyrecordQuery {
	if pq == nil {
		return nil
	}
	return &PhysicaltherapyrecordQuery{
		config:                  pq.config,
		limit:                   pq.limit,
		offset:                  pq.offset,
		order:                   append([]OrderFunc{}, pq.order...),
		predicates:              append([]predicate.Physicaltherapyrecord{}, pq.predicates...),
		withPersonnel:           pq.withPersonnel.Clone(),
		withPatient:             pq.withPatient.Clone(),
		withPhysicaltherapyroom: pq.withPhysicaltherapyroom.Clone(),
		withStatus:              pq.withStatus.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithPersonnel tells the query-builder to eager-load the nodes that are connected to
// the "personnel" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PhysicaltherapyrecordQuery) WithPersonnel(opts ...func(*PersonnelQuery)) *PhysicaltherapyrecordQuery {
	query := &PersonnelQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withPersonnel = query
	return pq
}

// WithPatient tells the query-builder to eager-load the nodes that are connected to
// the "patient" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PhysicaltherapyrecordQuery) WithPatient(opts ...func(*PatientQuery)) *PhysicaltherapyrecordQuery {
	query := &PatientQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withPatient = query
	return pq
}

// WithPhysicaltherapyroom tells the query-builder to eager-load the nodes that are connected to
// the "physicaltherapyroom" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PhysicaltherapyrecordQuery) WithPhysicaltherapyroom(opts ...func(*PhysicaltherapyroomQuery)) *PhysicaltherapyrecordQuery {
	query := &PhysicaltherapyroomQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withPhysicaltherapyroom = query
	return pq
}

// WithStatus tells the query-builder to eager-load the nodes that are connected to
// the "status" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PhysicaltherapyrecordQuery) WithStatus(opts ...func(*StatusQuery)) *PhysicaltherapyrecordQuery {
	query := &StatusQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withStatus = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Price int `json:"price,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Physicaltherapyrecord.Query().
//		GroupBy(physicaltherapyrecord.FieldPrice).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PhysicaltherapyrecordQuery) GroupBy(field string, fields ...string) *PhysicaltherapyrecordGroupBy {
	group := &PhysicaltherapyrecordGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Price int `json:"price,omitempty"`
//	}
//
//	client.Physicaltherapyrecord.Query().
//		Select(physicaltherapyrecord.FieldPrice).
//		Scan(ctx, &v)
//
func (pq *PhysicaltherapyrecordQuery) Select(field string, fields ...string) *PhysicaltherapyrecordSelect {
	pq.fields = append([]string{field}, fields...)
	return &PhysicaltherapyrecordSelect{PhysicaltherapyrecordQuery: pq}
}

func (pq *PhysicaltherapyrecordQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !physicaltherapyrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PhysicaltherapyrecordQuery) sqlAll(ctx context.Context) ([]*Physicaltherapyrecord, error) {
	var (
		nodes       = []*Physicaltherapyrecord{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [4]bool{
			pq.withPersonnel != nil,
			pq.withPatient != nil,
			pq.withPhysicaltherapyroom != nil,
			pq.withStatus != nil,
		}
	)
	if pq.withPersonnel != nil || pq.withPatient != nil || pq.withPhysicaltherapyroom != nil || pq.withStatus != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, physicaltherapyrecord.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Physicaltherapyrecord{config: pq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withPersonnel; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Physicaltherapyrecord)
		for i := range nodes {
			if fk := nodes[i]._Personnel_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(personnel.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "_Personnel_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Personnel = n
			}
		}
	}

	if query := pq.withPatient; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Physicaltherapyrecord)
		for i := range nodes {
			if fk := nodes[i]._Patient_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(patient.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "_Patient_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Patient = n
			}
		}
	}

	if query := pq.withPhysicaltherapyroom; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Physicaltherapyrecord)
		for i := range nodes {
			if fk := nodes[i].physicaltherapyroomname_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(physicaltherapyroom.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "physicaltherapyroomname_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Physicaltherapyroom = n
			}
		}
	}

	if query := pq.withStatus; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Physicaltherapyrecord)
		for i := range nodes {
			if fk := nodes[i].statusname_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(status.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "statusname_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Status = n
			}
		}
	}

	return nodes, nil
}

func (pq *PhysicaltherapyrecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PhysicaltherapyrecordQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pq *PhysicaltherapyrecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   physicaltherapyrecord.Table,
			Columns: physicaltherapyrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: physicaltherapyrecord.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, physicaltherapyrecord.FieldID)
		for i := range fields {
			if fields[i] != physicaltherapyrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, physicaltherapyrecord.ValidColumn)
			}
		}
	}
	return _spec
}

func (pq *PhysicaltherapyrecordQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(physicaltherapyrecord.Table)
	selector := builder.Select(t1.Columns(physicaltherapyrecord.Columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(physicaltherapyrecord.Columns...)...)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector, physicaltherapyrecord.ValidColumn)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PhysicaltherapyrecordGroupBy is the group-by builder for Physicaltherapyrecord entities.
type PhysicaltherapyrecordGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PhysicaltherapyrecordGroupBy) Aggregate(fns ...AggregateFunc) *PhysicaltherapyrecordGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *PhysicaltherapyrecordGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *PhysicaltherapyrecordGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PhysicaltherapyrecordGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PhysicaltherapyrecordGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *PhysicaltherapyrecordGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PhysicaltherapyrecordGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{physicaltherapyrecord.Label}
	default:
		err = fmt.Errorf("ent: PhysicaltherapyrecordGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgb *PhysicaltherapyrecordGroupBy) StringX(ctx context.Context) string {
	v, err := pgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PhysicaltherapyrecordGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PhysicaltherapyrecordGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *PhysicaltherapyrecordGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PhysicaltherapyrecordGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{physicaltherapyrecord.Label}
	default:
		err = fmt.Errorf("ent: PhysicaltherapyrecordGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgb *PhysicaltherapyrecordGroupBy) IntX(ctx context.Context) int {
	v, err := pgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PhysicaltherapyrecordGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PhysicaltherapyrecordGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *PhysicaltherapyrecordGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PhysicaltherapyrecordGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{physicaltherapyrecord.Label}
	default:
		err = fmt.Errorf("ent: PhysicaltherapyrecordGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgb *PhysicaltherapyrecordGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *PhysicaltherapyrecordGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: PhysicaltherapyrecordGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *PhysicaltherapyrecordGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *PhysicaltherapyrecordGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{physicaltherapyrecord.Label}
	default:
		err = fmt.Errorf("ent: PhysicaltherapyrecordGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgb *PhysicaltherapyrecordGroupBy) BoolX(ctx context.Context) bool {
	v, err := pgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *PhysicaltherapyrecordGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pgb.fields {
		if !physicaltherapyrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PhysicaltherapyrecordGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql
	columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
	columns = append(columns, pgb.fields...)
	for _, fn := range pgb.fns {
		columns = append(columns, fn(selector, physicaltherapyrecord.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(pgb.fields...)
}

// PhysicaltherapyrecordSelect is the builder for selecting fields of Physicaltherapyrecord entities.
type PhysicaltherapyrecordSelect struct {
	*PhysicaltherapyrecordQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PhysicaltherapyrecordSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.PhysicaltherapyrecordQuery.sqlQuery()
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *PhysicaltherapyrecordSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ps *PhysicaltherapyrecordSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PhysicaltherapyrecordSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *PhysicaltherapyrecordSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ps *PhysicaltherapyrecordSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{physicaltherapyrecord.Label}
	default:
		err = fmt.Errorf("ent: PhysicaltherapyrecordSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ps *PhysicaltherapyrecordSelect) StringX(ctx context.Context) string {
	v, err := ps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ps *PhysicaltherapyrecordSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PhysicaltherapyrecordSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *PhysicaltherapyrecordSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ps *PhysicaltherapyrecordSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{physicaltherapyrecord.Label}
	default:
		err = fmt.Errorf("ent: PhysicaltherapyrecordSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ps *PhysicaltherapyrecordSelect) IntX(ctx context.Context) int {
	v, err := ps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ps *PhysicaltherapyrecordSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PhysicaltherapyrecordSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *PhysicaltherapyrecordSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ps *PhysicaltherapyrecordSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{physicaltherapyrecord.Label}
	default:
		err = fmt.Errorf("ent: PhysicaltherapyrecordSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ps *PhysicaltherapyrecordSelect) Float64X(ctx context.Context) float64 {
	v, err := ps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ps *PhysicaltherapyrecordSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: PhysicaltherapyrecordSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *PhysicaltherapyrecordSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ps *PhysicaltherapyrecordSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{physicaltherapyrecord.Label}
	default:
		err = fmt.Errorf("ent: PhysicaltherapyrecordSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ps *PhysicaltherapyrecordSelect) BoolX(ctx context.Context) bool {
	v, err := ps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *PhysicaltherapyrecordSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sqlQuery().Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ps *PhysicaltherapyrecordSelect) sqlQuery() sql.Querier {
	selector := ps.sql
	selector.Select(selector.Columns(ps.fields...)...)
	return selector
}
