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
	"github.com/to63/app/ent/bonedisease"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/predicate"
	"github.com/to63/app/ent/remedy"
)

// BonediseaseQuery is the builder for querying Bonedisease entities.
type BonediseaseQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Bonedisease
	// eager-loading edges.
	withRemedy    *RemedyQuery
	withPatient   *PatientQuery
	withPersonnel *PersonnelQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BonediseaseQuery builder.
func (bq *BonediseaseQuery) Where(ps ...predicate.Bonedisease) *BonediseaseQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BonediseaseQuery) Limit(limit int) *BonediseaseQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BonediseaseQuery) Offset(offset int) *BonediseaseQuery {
	bq.offset = &offset
	return bq
}

// Order adds an order step to the query.
func (bq *BonediseaseQuery) Order(o ...OrderFunc) *BonediseaseQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryRemedy chains the current query on the "remedy" edge.
func (bq *BonediseaseQuery) QueryRemedy() *RemedyQuery {
	query := &RemedyQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bonedisease.Table, bonedisease.FieldID, selector),
			sqlgraph.To(remedy.Table, remedy.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bonedisease.RemedyTable, bonedisease.RemedyColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPatient chains the current query on the "patient" edge.
func (bq *BonediseaseQuery) QueryPatient() *PatientQuery {
	query := &PatientQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bonedisease.Table, bonedisease.FieldID, selector),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bonedisease.PatientTable, bonedisease.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPersonnel chains the current query on the "personnel" edge.
func (bq *BonediseaseQuery) QueryPersonnel() *PersonnelQuery {
	query := &PersonnelQuery{config: bq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bonedisease.Table, bonedisease.FieldID, selector),
			sqlgraph.To(personnel.Table, personnel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bonedisease.PersonnelTable, bonedisease.PersonnelColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Bonedisease entity from the query.
// Returns a *NotFoundError when no Bonedisease was found.
func (bq *BonediseaseQuery) First(ctx context.Context) (*Bonedisease, error) {
	nodes, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bonedisease.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BonediseaseQuery) FirstX(ctx context.Context) *Bonedisease {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Bonedisease ID from the query.
// Returns a *NotFoundError when no Bonedisease ID was found.
func (bq *BonediseaseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bonedisease.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BonediseaseQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Bonedisease entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Bonedisease entity is not found.
// Returns a *NotFoundError when no Bonedisease entities are found.
func (bq *BonediseaseQuery) Only(ctx context.Context) (*Bonedisease, error) {
	nodes, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bonedisease.Label}
	default:
		return nil, &NotSingularError{bonedisease.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BonediseaseQuery) OnlyX(ctx context.Context) *Bonedisease {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Bonedisease ID in the query.
// Returns a *NotSingularError when exactly one Bonedisease ID is not found.
// Returns a *NotFoundError when no entities are found.
func (bq *BonediseaseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bonedisease.Label}
	default:
		err = &NotSingularError{bonedisease.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BonediseaseQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bonediseases.
func (bq *BonediseaseQuery) All(ctx context.Context) ([]*Bonedisease, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bq *BonediseaseQuery) AllX(ctx context.Context) []*Bonedisease {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Bonedisease IDs.
func (bq *BonediseaseQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bq.Select(bonedisease.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BonediseaseQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BonediseaseQuery) Count(ctx context.Context) (int, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BonediseaseQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BonediseaseQuery) Exist(ctx context.Context) (bool, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BonediseaseQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BonediseaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BonediseaseQuery) Clone() *BonediseaseQuery {
	if bq == nil {
		return nil
	}
	return &BonediseaseQuery{
		config:        bq.config,
		limit:         bq.limit,
		offset:        bq.offset,
		order:         append([]OrderFunc{}, bq.order...),
		predicates:    append([]predicate.Bonedisease{}, bq.predicates...),
		withRemedy:    bq.withRemedy.Clone(),
		withPatient:   bq.withPatient.Clone(),
		withPersonnel: bq.withPersonnel.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithRemedy tells the query-builder to eager-load the nodes that are connected to
// the "remedy" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BonediseaseQuery) WithRemedy(opts ...func(*RemedyQuery)) *BonediseaseQuery {
	query := &RemedyQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withRemedy = query
	return bq
}

// WithPatient tells the query-builder to eager-load the nodes that are connected to
// the "patient" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BonediseaseQuery) WithPatient(opts ...func(*PatientQuery)) *BonediseaseQuery {
	query := &PatientQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withPatient = query
	return bq
}

// WithPersonnel tells the query-builder to eager-load the nodes that are connected to
// the "personnel" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BonediseaseQuery) WithPersonnel(opts ...func(*PersonnelQuery)) *BonediseaseQuery {
	query := &PersonnelQuery{config: bq.config}
	for _, opt := range opts {
		opt(query)
	}
	bq.withPersonnel = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AddedTime time.Time `json:"addedTime,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Bonedisease.Query().
//		GroupBy(bonedisease.FieldAddedTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bq *BonediseaseQuery) GroupBy(field string, fields ...string) *BonediseaseGroupBy {
	group := &BonediseaseGroupBy{config: bq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AddedTime time.Time `json:"addedTime,omitempty"`
//	}
//
//	client.Bonedisease.Query().
//		Select(bonedisease.FieldAddedTime).
//		Scan(ctx, &v)
//
func (bq *BonediseaseQuery) Select(field string, fields ...string) *BonediseaseSelect {
	bq.fields = append([]string{field}, fields...)
	return &BonediseaseSelect{BonediseaseQuery: bq}
}

func (bq *BonediseaseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bq.fields {
		if !bonedisease.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BonediseaseQuery) sqlAll(ctx context.Context) ([]*Bonedisease, error) {
	var (
		nodes       = []*Bonedisease{}
		withFKs     = bq.withFKs
		_spec       = bq.querySpec()
		loadedTypes = [3]bool{
			bq.withRemedy != nil,
			bq.withPatient != nil,
			bq.withPersonnel != nil,
		}
	)
	if bq.withRemedy != nil || bq.withPatient != nil || bq.withPersonnel != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, bonedisease.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Bonedisease{config: bq.config}
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
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bq.withRemedy; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Bonedisease)
		for i := range nodes {
			if fk := nodes[i].remedy_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(remedy.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "remedy_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Remedy = n
			}
		}
	}

	if query := bq.withPatient; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Bonedisease)
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

	if query := bq.withPersonnel; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Bonedisease)
		for i := range nodes {
			if fk := nodes[i]._Personel_id; fk != nil {
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
				return nil, fmt.Errorf(`unexpected foreign-key "_Personel_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Personnel = n
			}
		}
	}

	return nodes, nil
}

func (bq *BonediseaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BonediseaseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (bq *BonediseaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bonedisease.Table,
			Columns: bonedisease.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bonedisease.FieldID,
			},
		},
		From:   bq.sql,
		Unique: true,
	}
	if fields := bq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bonedisease.FieldID)
		for i := range fields {
			if fields[i] != bonedisease.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, bonedisease.ValidColumn)
			}
		}
	}
	return _spec
}

func (bq *BonediseaseQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(bonedisease.Table)
	selector := builder.Select(t1.Columns(bonedisease.Columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(bonedisease.Columns...)...)
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector, bonedisease.ValidColumn)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BonediseaseGroupBy is the group-by builder for Bonedisease entities.
type BonediseaseGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BonediseaseGroupBy) Aggregate(fns ...AggregateFunc) *BonediseaseGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bgb *BonediseaseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bgb.path(ctx)
	if err != nil {
		return err
	}
	bgb.sql = query
	return bgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bgb *BonediseaseGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (bgb *BonediseaseGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BonediseaseGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bgb *BonediseaseGroupBy) StringsX(ctx context.Context) []string {
	v, err := bgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bgb *BonediseaseGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bonedisease.Label}
	default:
		err = fmt.Errorf("ent: BonediseaseGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bgb *BonediseaseGroupBy) StringX(ctx context.Context) string {
	v, err := bgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (bgb *BonediseaseGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BonediseaseGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bgb *BonediseaseGroupBy) IntsX(ctx context.Context) []int {
	v, err := bgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bgb *BonediseaseGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bonedisease.Label}
	default:
		err = fmt.Errorf("ent: BonediseaseGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bgb *BonediseaseGroupBy) IntX(ctx context.Context) int {
	v, err := bgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (bgb *BonediseaseGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BonediseaseGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bgb *BonediseaseGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bgb *BonediseaseGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bonedisease.Label}
	default:
		err = fmt.Errorf("ent: BonediseaseGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bgb *BonediseaseGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (bgb *BonediseaseGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bgb.fields) > 1 {
		return nil, errors.New("ent: BonediseaseGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bgb *BonediseaseGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bgb *BonediseaseGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bonedisease.Label}
	default:
		err = fmt.Errorf("ent: BonediseaseGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bgb *BonediseaseGroupBy) BoolX(ctx context.Context) bool {
	v, err := bgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bgb *BonediseaseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bgb.fields {
		if !bonedisease.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BonediseaseGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql
	columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
	columns = append(columns, bgb.fields...)
	for _, fn := range bgb.fns {
		columns = append(columns, fn(selector, bonedisease.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(bgb.fields...)
}

// BonediseaseSelect is the builder for selecting fields of Bonedisease entities.
type BonediseaseSelect struct {
	*BonediseaseQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BonediseaseSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	bs.sql = bs.BonediseaseQuery.sqlQuery()
	return bs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bs *BonediseaseSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (bs *BonediseaseSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BonediseaseSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bs *BonediseaseSelect) StringsX(ctx context.Context) []string {
	v, err := bs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (bs *BonediseaseSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bonedisease.Label}
	default:
		err = fmt.Errorf("ent: BonediseaseSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bs *BonediseaseSelect) StringX(ctx context.Context) string {
	v, err := bs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (bs *BonediseaseSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BonediseaseSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bs *BonediseaseSelect) IntsX(ctx context.Context) []int {
	v, err := bs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (bs *BonediseaseSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bonedisease.Label}
	default:
		err = fmt.Errorf("ent: BonediseaseSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bs *BonediseaseSelect) IntX(ctx context.Context) int {
	v, err := bs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (bs *BonediseaseSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BonediseaseSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bs *BonediseaseSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (bs *BonediseaseSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bonedisease.Label}
	default:
		err = fmt.Errorf("ent: BonediseaseSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bs *BonediseaseSelect) Float64X(ctx context.Context) float64 {
	v, err := bs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (bs *BonediseaseSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bs.fields) > 1 {
		return nil, errors.New("ent: BonediseaseSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bs *BonediseaseSelect) BoolsX(ctx context.Context) []bool {
	v, err := bs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (bs *BonediseaseSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bonedisease.Label}
	default:
		err = fmt.Errorf("ent: BonediseaseSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bs *BonediseaseSelect) BoolX(ctx context.Context) bool {
	v, err := bs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bs *BonediseaseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bs.sqlQuery().Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bs *BonediseaseSelect) sqlQuery() sql.Querier {
	selector := bs.sql
	selector.Select(selector.Columns(bs.fields...)...)
	return selector
}
