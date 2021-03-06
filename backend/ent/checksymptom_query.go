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
	"github.com/to63/app/ent/checksymptom"
	"github.com/to63/app/ent/disease"
	"github.com/to63/app/ent/doctorordersheet"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/predicate"
)

// ChecksymptomQuery is the builder for querying Checksymptom entities.
type ChecksymptomQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Checksymptom
	// eager-loading edges.
	withPatient          *PatientQuery
	withPersonnel        *PersonnelQuery
	withDoctorordersheet *DoctorordersheetQuery
	withDisease          *DiseaseQuery
	withFKs              bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ChecksymptomQuery builder.
func (cq *ChecksymptomQuery) Where(ps ...predicate.Checksymptom) *ChecksymptomQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *ChecksymptomQuery) Limit(limit int) *ChecksymptomQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *ChecksymptomQuery) Offset(offset int) *ChecksymptomQuery {
	cq.offset = &offset
	return cq
}

// Order adds an order step to the query.
func (cq *ChecksymptomQuery) Order(o ...OrderFunc) *ChecksymptomQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryPatient chains the current query on the "patient" edge.
func (cq *ChecksymptomQuery) QueryPatient() *PatientQuery {
	query := &PatientQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checksymptom.Table, checksymptom.FieldID, selector),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checksymptom.PatientTable, checksymptom.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPersonnel chains the current query on the "personnel" edge.
func (cq *ChecksymptomQuery) QueryPersonnel() *PersonnelQuery {
	query := &PersonnelQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checksymptom.Table, checksymptom.FieldID, selector),
			sqlgraph.To(personnel.Table, personnel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checksymptom.PersonnelTable, checksymptom.PersonnelColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDoctorordersheet chains the current query on the "doctorordersheet" edge.
func (cq *ChecksymptomQuery) QueryDoctorordersheet() *DoctorordersheetQuery {
	query := &DoctorordersheetQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checksymptom.Table, checksymptom.FieldID, selector),
			sqlgraph.To(doctorordersheet.Table, doctorordersheet.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checksymptom.DoctorordersheetTable, checksymptom.DoctorordersheetColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDisease chains the current query on the "disease" edge.
func (cq *ChecksymptomQuery) QueryDisease() *DiseaseQuery {
	query := &DiseaseQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery()
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(checksymptom.Table, checksymptom.FieldID, selector),
			sqlgraph.To(disease.Table, disease.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, checksymptom.DiseaseTable, checksymptom.DiseaseColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Checksymptom entity from the query.
// Returns a *NotFoundError when no Checksymptom was found.
func (cq *ChecksymptomQuery) First(ctx context.Context) (*Checksymptom, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{checksymptom.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ChecksymptomQuery) FirstX(ctx context.Context) *Checksymptom {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Checksymptom ID from the query.
// Returns a *NotFoundError when no Checksymptom ID was found.
func (cq *ChecksymptomQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{checksymptom.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ChecksymptomQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Checksymptom entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Checksymptom entity is not found.
// Returns a *NotFoundError when no Checksymptom entities are found.
func (cq *ChecksymptomQuery) Only(ctx context.Context) (*Checksymptom, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{checksymptom.Label}
	default:
		return nil, &NotSingularError{checksymptom.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ChecksymptomQuery) OnlyX(ctx context.Context) *Checksymptom {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Checksymptom ID in the query.
// Returns a *NotSingularError when exactly one Checksymptom ID is not found.
// Returns a *NotFoundError when no entities are found.
func (cq *ChecksymptomQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{checksymptom.Label}
	default:
		err = &NotSingularError{checksymptom.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ChecksymptomQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Checksymptoms.
func (cq *ChecksymptomQuery) All(ctx context.Context) ([]*Checksymptom, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *ChecksymptomQuery) AllX(ctx context.Context) []*Checksymptom {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Checksymptom IDs.
func (cq *ChecksymptomQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := cq.Select(checksymptom.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ChecksymptomQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ChecksymptomQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ChecksymptomQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ChecksymptomQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ChecksymptomQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ChecksymptomQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ChecksymptomQuery) Clone() *ChecksymptomQuery {
	if cq == nil {
		return nil
	}
	return &ChecksymptomQuery{
		config:               cq.config,
		limit:                cq.limit,
		offset:               cq.offset,
		order:                append([]OrderFunc{}, cq.order...),
		predicates:           append([]predicate.Checksymptom{}, cq.predicates...),
		withPatient:          cq.withPatient.Clone(),
		withPersonnel:        cq.withPersonnel.Clone(),
		withDoctorordersheet: cq.withDoctorordersheet.Clone(),
		withDisease:          cq.withDisease.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithPatient tells the query-builder to eager-load the nodes that are connected to
// the "patient" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChecksymptomQuery) WithPatient(opts ...func(*PatientQuery)) *ChecksymptomQuery {
	query := &PatientQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withPatient = query
	return cq
}

// WithPersonnel tells the query-builder to eager-load the nodes that are connected to
// the "personnel" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChecksymptomQuery) WithPersonnel(opts ...func(*PersonnelQuery)) *ChecksymptomQuery {
	query := &PersonnelQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withPersonnel = query
	return cq
}

// WithDoctorordersheet tells the query-builder to eager-load the nodes that are connected to
// the "doctorordersheet" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChecksymptomQuery) WithDoctorordersheet(opts ...func(*DoctorordersheetQuery)) *ChecksymptomQuery {
	query := &DoctorordersheetQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withDoctorordersheet = query
	return cq
}

// WithDisease tells the query-builder to eager-load the nodes that are connected to
// the "disease" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ChecksymptomQuery) WithDisease(opts ...func(*DiseaseQuery)) *ChecksymptomQuery {
	query := &DiseaseQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withDisease = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"date,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Checksymptom.Query().
//		GroupBy(checksymptom.FieldDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (cq *ChecksymptomQuery) GroupBy(field string, fields ...string) *ChecksymptomGroupBy {
	group := &ChecksymptomGroupBy{config: cq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"date,omitempty"`
//	}
//
//	client.Checksymptom.Query().
//		Select(checksymptom.FieldDate).
//		Scan(ctx, &v)
//
func (cq *ChecksymptomQuery) Select(field string, fields ...string) *ChecksymptomSelect {
	cq.fields = append([]string{field}, fields...)
	return &ChecksymptomSelect{ChecksymptomQuery: cq}
}

func (cq *ChecksymptomQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !checksymptom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ChecksymptomQuery) sqlAll(ctx context.Context) ([]*Checksymptom, error) {
	var (
		nodes       = []*Checksymptom{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [4]bool{
			cq.withPatient != nil,
			cq.withPersonnel != nil,
			cq.withDoctorordersheet != nil,
			cq.withDisease != nil,
		}
	)
	if cq.withPatient != nil || cq.withPersonnel != nil || cq.withDoctorordersheet != nil || cq.withDisease != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, checksymptom.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Checksymptom{config: cq.config}
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
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := cq.withPatient; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Checksymptom)
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

	if query := cq.withPersonnel; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Checksymptom)
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

	if query := cq.withDoctorordersheet; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Checksymptom)
		for i := range nodes {
			if fk := nodes[i]._Doctorordersheet_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(doctorordersheet.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "_Doctorordersheet_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Doctorordersheet = n
			}
		}
	}

	if query := cq.withDisease; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Checksymptom)
		for i := range nodes {
			if fk := nodes[i].disease_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(disease.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "disease_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Disease = n
			}
		}
	}

	return nodes, nil
}

func (cq *ChecksymptomQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ChecksymptomQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := cq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (cq *ChecksymptomQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   checksymptom.Table,
			Columns: checksymptom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: checksymptom.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, checksymptom.FieldID)
		for i := range fields {
			if fields[i] != checksymptom.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, checksymptom.ValidColumn)
			}
		}
	}
	return _spec
}

func (cq *ChecksymptomQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(checksymptom.Table)
	selector := builder.Select(t1.Columns(checksymptom.Columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(checksymptom.Columns...)...)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector, checksymptom.ValidColumn)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ChecksymptomGroupBy is the group-by builder for Checksymptom entities.
type ChecksymptomGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ChecksymptomGroupBy) Aggregate(fns ...AggregateFunc) *ChecksymptomGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *ChecksymptomGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cgb *ChecksymptomGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := cgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *ChecksymptomGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ChecksymptomGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cgb *ChecksymptomGroupBy) StringsX(ctx context.Context) []string {
	v, err := cgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *ChecksymptomGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checksymptom.Label}
	default:
		err = fmt.Errorf("ent: ChecksymptomGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cgb *ChecksymptomGroupBy) StringX(ctx context.Context) string {
	v, err := cgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *ChecksymptomGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ChecksymptomGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cgb *ChecksymptomGroupBy) IntsX(ctx context.Context) []int {
	v, err := cgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *ChecksymptomGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checksymptom.Label}
	default:
		err = fmt.Errorf("ent: ChecksymptomGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cgb *ChecksymptomGroupBy) IntX(ctx context.Context) int {
	v, err := cgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *ChecksymptomGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ChecksymptomGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cgb *ChecksymptomGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := cgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *ChecksymptomGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checksymptom.Label}
	default:
		err = fmt.Errorf("ent: ChecksymptomGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cgb *ChecksymptomGroupBy) Float64X(ctx context.Context) float64 {
	v, err := cgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (cgb *ChecksymptomGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(cgb.fields) > 1 {
		return nil, errors.New("ent: ChecksymptomGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := cgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cgb *ChecksymptomGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := cgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (cgb *ChecksymptomGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checksymptom.Label}
	default:
		err = fmt.Errorf("ent: ChecksymptomGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cgb *ChecksymptomGroupBy) BoolX(ctx context.Context) bool {
	v, err := cgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cgb *ChecksymptomGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range cgb.fields {
		if !checksymptom.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *ChecksymptomGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql
	columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
	columns = append(columns, cgb.fields...)
	for _, fn := range cgb.fns {
		columns = append(columns, fn(selector, checksymptom.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(cgb.fields...)
}

// ChecksymptomSelect is the builder for selecting fields of Checksymptom entities.
type ChecksymptomSelect struct {
	*ChecksymptomQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ChecksymptomSelect) Scan(ctx context.Context, v interface{}) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.ChecksymptomQuery.sqlQuery()
	return cs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (cs *ChecksymptomSelect) ScanX(ctx context.Context, v interface{}) {
	if err := cs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (cs *ChecksymptomSelect) Strings(ctx context.Context) ([]string, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ChecksymptomSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (cs *ChecksymptomSelect) StringsX(ctx context.Context) []string {
	v, err := cs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (cs *ChecksymptomSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = cs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checksymptom.Label}
	default:
		err = fmt.Errorf("ent: ChecksymptomSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (cs *ChecksymptomSelect) StringX(ctx context.Context) string {
	v, err := cs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (cs *ChecksymptomSelect) Ints(ctx context.Context) ([]int, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ChecksymptomSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (cs *ChecksymptomSelect) IntsX(ctx context.Context) []int {
	v, err := cs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (cs *ChecksymptomSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = cs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checksymptom.Label}
	default:
		err = fmt.Errorf("ent: ChecksymptomSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (cs *ChecksymptomSelect) IntX(ctx context.Context) int {
	v, err := cs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (cs *ChecksymptomSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ChecksymptomSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (cs *ChecksymptomSelect) Float64sX(ctx context.Context) []float64 {
	v, err := cs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (cs *ChecksymptomSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = cs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checksymptom.Label}
	default:
		err = fmt.Errorf("ent: ChecksymptomSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (cs *ChecksymptomSelect) Float64X(ctx context.Context) float64 {
	v, err := cs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (cs *ChecksymptomSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(cs.fields) > 1 {
		return nil, errors.New("ent: ChecksymptomSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := cs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (cs *ChecksymptomSelect) BoolsX(ctx context.Context) []bool {
	v, err := cs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (cs *ChecksymptomSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = cs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{checksymptom.Label}
	default:
		err = fmt.Errorf("ent: ChecksymptomSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (cs *ChecksymptomSelect) BoolX(ctx context.Context) bool {
	v, err := cs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cs *ChecksymptomSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := cs.sqlQuery().Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cs *ChecksymptomSelect) sqlQuery() sql.Querier {
	selector := cs.sql
	selector.Select(selector.Columns(cs.fields...)...)
	return selector
}
