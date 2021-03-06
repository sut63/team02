// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/dentalappointment"
	"github.com/to63/app/ent/dentalkind"
	"github.com/to63/app/ent/predicate"
)

// DentalkindUpdate is the builder for updating Dentalkind entities.
type DentalkindUpdate struct {
	config
	hooks    []Hook
	mutation *DentalkindMutation
}

// Where adds a new predicate for the DentalkindUpdate builder.
func (du *DentalkindUpdate) Where(ps ...predicate.Dentalkind) *DentalkindUpdate {
	du.mutation.predicates = append(du.mutation.predicates, ps...)
	return du
}

// SetKindname sets the "kindname" field.
func (du *DentalkindUpdate) SetKindname(s string) *DentalkindUpdate {
	du.mutation.SetKindname(s)
	return du
}

// AddDentalappointmentIDs adds the "Dentalappointment" edge to the Dentalappointment entity by IDs.
func (du *DentalkindUpdate) AddDentalappointmentIDs(ids ...int) *DentalkindUpdate {
	du.mutation.AddDentalappointmentIDs(ids...)
	return du
}

// AddDentalappointment adds the "Dentalappointment" edges to the Dentalappointment entity.
func (du *DentalkindUpdate) AddDentalappointment(d ...*Dentalappointment) *DentalkindUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDentalappointmentIDs(ids...)
}

// Mutation returns the DentalkindMutation object of the builder.
func (du *DentalkindUpdate) Mutation() *DentalkindMutation {
	return du.mutation
}

// ClearDentalappointment clears all "Dentalappointment" edges to the Dentalappointment entity.
func (du *DentalkindUpdate) ClearDentalappointment() *DentalkindUpdate {
	du.mutation.ClearDentalappointment()
	return du
}

// RemoveDentalappointmentIDs removes the "Dentalappointment" edge to Dentalappointment entities by IDs.
func (du *DentalkindUpdate) RemoveDentalappointmentIDs(ids ...int) *DentalkindUpdate {
	du.mutation.RemoveDentalappointmentIDs(ids...)
	return du
}

// RemoveDentalappointment removes "Dentalappointment" edges to Dentalappointment entities.
func (du *DentalkindUpdate) RemoveDentalappointment(d ...*Dentalappointment) *DentalkindUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDentalappointmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DentalkindUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DentalkindMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DentalkindUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DentalkindUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DentalkindUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DentalkindUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dentalkind.Table,
			Columns: dentalkind.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dentalkind.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Kindname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentalkind.FieldKindname,
		})
	}
	if du.mutation.DentalappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentalkind.DentalappointmentTable,
			Columns: []string{dentalkind.DentalappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalappointment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDentalappointmentIDs(); len(nodes) > 0 && !du.mutation.DentalappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentalkind.DentalappointmentTable,
			Columns: []string{dentalkind.DentalappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalappointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DentalappointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentalkind.DentalappointmentTable,
			Columns: []string{dentalkind.DentalappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalappointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dentalkind.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DentalkindUpdateOne is the builder for updating a single Dentalkind entity.
type DentalkindUpdateOne struct {
	config
	hooks    []Hook
	mutation *DentalkindMutation
}

// SetKindname sets the "kindname" field.
func (duo *DentalkindUpdateOne) SetKindname(s string) *DentalkindUpdateOne {
	duo.mutation.SetKindname(s)
	return duo
}

// AddDentalappointmentIDs adds the "Dentalappointment" edge to the Dentalappointment entity by IDs.
func (duo *DentalkindUpdateOne) AddDentalappointmentIDs(ids ...int) *DentalkindUpdateOne {
	duo.mutation.AddDentalappointmentIDs(ids...)
	return duo
}

// AddDentalappointment adds the "Dentalappointment" edges to the Dentalappointment entity.
func (duo *DentalkindUpdateOne) AddDentalappointment(d ...*Dentalappointment) *DentalkindUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDentalappointmentIDs(ids...)
}

// Mutation returns the DentalkindMutation object of the builder.
func (duo *DentalkindUpdateOne) Mutation() *DentalkindMutation {
	return duo.mutation
}

// ClearDentalappointment clears all "Dentalappointment" edges to the Dentalappointment entity.
func (duo *DentalkindUpdateOne) ClearDentalappointment() *DentalkindUpdateOne {
	duo.mutation.ClearDentalappointment()
	return duo
}

// RemoveDentalappointmentIDs removes the "Dentalappointment" edge to Dentalappointment entities by IDs.
func (duo *DentalkindUpdateOne) RemoveDentalappointmentIDs(ids ...int) *DentalkindUpdateOne {
	duo.mutation.RemoveDentalappointmentIDs(ids...)
	return duo
}

// RemoveDentalappointment removes "Dentalappointment" edges to Dentalappointment entities.
func (duo *DentalkindUpdateOne) RemoveDentalappointment(d ...*Dentalappointment) *DentalkindUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDentalappointmentIDs(ids...)
}

// Save executes the query and returns the updated Dentalkind entity.
func (duo *DentalkindUpdateOne) Save(ctx context.Context) (*Dentalkind, error) {
	var (
		err  error
		node *Dentalkind
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DentalkindMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DentalkindUpdateOne) SaveX(ctx context.Context) *Dentalkind {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DentalkindUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DentalkindUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DentalkindUpdateOne) sqlSave(ctx context.Context) (_node *Dentalkind, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dentalkind.Table,
			Columns: dentalkind.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dentalkind.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Dentalkind.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Kindname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentalkind.FieldKindname,
		})
	}
	if duo.mutation.DentalappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentalkind.DentalappointmentTable,
			Columns: []string{dentalkind.DentalappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalappointment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDentalappointmentIDs(); len(nodes) > 0 && !duo.mutation.DentalappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentalkind.DentalappointmentTable,
			Columns: []string{dentalkind.DentalappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalappointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DentalappointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentalkind.DentalappointmentTable,
			Columns: []string{dentalkind.DentalappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalappointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Dentalkind{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dentalkind.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
