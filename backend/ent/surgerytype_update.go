// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/predicate"
	"github.com/to63/app/ent/surgeryappointment"
	"github.com/to63/app/ent/surgerytype"
)

// SurgerytypeUpdate is the builder for updating Surgerytype entities.
type SurgerytypeUpdate struct {
	config
	hooks    []Hook
	mutation *SurgerytypeMutation
}

// Where adds a new predicate for the SurgerytypeUpdate builder.
func (su *SurgerytypeUpdate) Where(ps ...predicate.Surgerytype) *SurgerytypeUpdate {
	su.mutation.predicates = append(su.mutation.predicates, ps...)
	return su
}

// SetTypename sets the "typename" field.
func (su *SurgerytypeUpdate) SetTypename(s string) *SurgerytypeUpdate {
	su.mutation.SetTypename(s)
	return su
}

// AddSurgeryappointmentIDs adds the "Surgeryappointment" edge to the Surgeryappointment entity by IDs.
func (su *SurgerytypeUpdate) AddSurgeryappointmentIDs(ids ...int) *SurgerytypeUpdate {
	su.mutation.AddSurgeryappointmentIDs(ids...)
	return su
}

// AddSurgeryappointment adds the "Surgeryappointment" edges to the Surgeryappointment entity.
func (su *SurgerytypeUpdate) AddSurgeryappointment(s ...*Surgeryappointment) *SurgerytypeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddSurgeryappointmentIDs(ids...)
}

// Mutation returns the SurgerytypeMutation object of the builder.
func (su *SurgerytypeUpdate) Mutation() *SurgerytypeMutation {
	return su.mutation
}

// ClearSurgeryappointment clears all "Surgeryappointment" edges to the Surgeryappointment entity.
func (su *SurgerytypeUpdate) ClearSurgeryappointment() *SurgerytypeUpdate {
	su.mutation.ClearSurgeryappointment()
	return su
}

// RemoveSurgeryappointmentIDs removes the "Surgeryappointment" edge to Surgeryappointment entities by IDs.
func (su *SurgerytypeUpdate) RemoveSurgeryappointmentIDs(ids ...int) *SurgerytypeUpdate {
	su.mutation.RemoveSurgeryappointmentIDs(ids...)
	return su
}

// RemoveSurgeryappointment removes "Surgeryappointment" edges to Surgeryappointment entities.
func (su *SurgerytypeUpdate) RemoveSurgeryappointment(s ...*Surgeryappointment) *SurgerytypeUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveSurgeryappointmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SurgerytypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SurgerytypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SurgerytypeUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SurgerytypeUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SurgerytypeUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SurgerytypeUpdate) check() error {
	if v, ok := su.mutation.Typename(); ok {
		if err := surgerytype.TypenameValidator(v); err != nil {
			return &ValidationError{Name: "typename", err: fmt.Errorf("ent: validator failed for field \"typename\": %w", err)}
		}
	}
	return nil
}

func (su *SurgerytypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   surgerytype.Table,
			Columns: surgerytype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: surgerytype.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Typename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: surgerytype.FieldTypename,
		})
	}
	if su.mutation.SurgeryappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surgerytype.SurgeryappointmentTable,
			Columns: []string{surgerytype.SurgeryappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surgeryappointment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedSurgeryappointmentIDs(); len(nodes) > 0 && !su.mutation.SurgeryappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surgerytype.SurgeryappointmentTable,
			Columns: []string{surgerytype.SurgeryappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surgeryappointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SurgeryappointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surgerytype.SurgeryappointmentTable,
			Columns: []string{surgerytype.SurgeryappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surgeryappointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surgerytype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SurgerytypeUpdateOne is the builder for updating a single Surgerytype entity.
type SurgerytypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *SurgerytypeMutation
}

// SetTypename sets the "typename" field.
func (suo *SurgerytypeUpdateOne) SetTypename(s string) *SurgerytypeUpdateOne {
	suo.mutation.SetTypename(s)
	return suo
}

// AddSurgeryappointmentIDs adds the "Surgeryappointment" edge to the Surgeryappointment entity by IDs.
func (suo *SurgerytypeUpdateOne) AddSurgeryappointmentIDs(ids ...int) *SurgerytypeUpdateOne {
	suo.mutation.AddSurgeryappointmentIDs(ids...)
	return suo
}

// AddSurgeryappointment adds the "Surgeryappointment" edges to the Surgeryappointment entity.
func (suo *SurgerytypeUpdateOne) AddSurgeryappointment(s ...*Surgeryappointment) *SurgerytypeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddSurgeryappointmentIDs(ids...)
}

// Mutation returns the SurgerytypeMutation object of the builder.
func (suo *SurgerytypeUpdateOne) Mutation() *SurgerytypeMutation {
	return suo.mutation
}

// ClearSurgeryappointment clears all "Surgeryappointment" edges to the Surgeryappointment entity.
func (suo *SurgerytypeUpdateOne) ClearSurgeryappointment() *SurgerytypeUpdateOne {
	suo.mutation.ClearSurgeryappointment()
	return suo
}

// RemoveSurgeryappointmentIDs removes the "Surgeryappointment" edge to Surgeryappointment entities by IDs.
func (suo *SurgerytypeUpdateOne) RemoveSurgeryappointmentIDs(ids ...int) *SurgerytypeUpdateOne {
	suo.mutation.RemoveSurgeryappointmentIDs(ids...)
	return suo
}

// RemoveSurgeryappointment removes "Surgeryappointment" edges to Surgeryappointment entities.
func (suo *SurgerytypeUpdateOne) RemoveSurgeryappointment(s ...*Surgeryappointment) *SurgerytypeUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveSurgeryappointmentIDs(ids...)
}

// Save executes the query and returns the updated Surgerytype entity.
func (suo *SurgerytypeUpdateOne) Save(ctx context.Context) (*Surgerytype, error) {
	var (
		err  error
		node *Surgerytype
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SurgerytypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SurgerytypeUpdateOne) SaveX(ctx context.Context) *Surgerytype {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SurgerytypeUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SurgerytypeUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SurgerytypeUpdateOne) check() error {
	if v, ok := suo.mutation.Typename(); ok {
		if err := surgerytype.TypenameValidator(v); err != nil {
			return &ValidationError{Name: "typename", err: fmt.Errorf("ent: validator failed for field \"typename\": %w", err)}
		}
	}
	return nil
}

func (suo *SurgerytypeUpdateOne) sqlSave(ctx context.Context) (_node *Surgerytype, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   surgerytype.Table,
			Columns: surgerytype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: surgerytype.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Surgerytype.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Typename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: surgerytype.FieldTypename,
		})
	}
	if suo.mutation.SurgeryappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surgerytype.SurgeryappointmentTable,
			Columns: []string{surgerytype.SurgeryappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surgeryappointment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedSurgeryappointmentIDs(); len(nodes) > 0 && !suo.mutation.SurgeryappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surgerytype.SurgeryappointmentTable,
			Columns: []string{surgerytype.SurgeryappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surgeryappointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SurgeryappointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   surgerytype.SurgeryappointmentTable,
			Columns: []string{surgerytype.SurgeryappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surgeryappointment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Surgerytype{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surgerytype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
