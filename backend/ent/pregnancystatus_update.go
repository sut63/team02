// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/antenatalinformation"
	"github.com/to63/app/ent/predicate"
	"github.com/to63/app/ent/pregnancystatus"
)

// PregnancystatusUpdate is the builder for updating Pregnancystatus entities.
type PregnancystatusUpdate struct {
	config
	hooks    []Hook
	mutation *PregnancystatusMutation
}

// Where adds a new predicate for the PregnancystatusUpdate builder.
func (pu *PregnancystatusUpdate) Where(ps ...predicate.Pregnancystatus) *PregnancystatusUpdate {
	pu.mutation.predicates = append(pu.mutation.predicates, ps...)
	return pu
}

// SetPregnancystatus sets the "Pregnancystatus" field.
func (pu *PregnancystatusUpdate) SetPregnancystatus(s string) *PregnancystatusUpdate {
	pu.mutation.SetPregnancystatus(s)
	return pu
}

// AddAntenatalinformationIDs adds the "Antenatalinformation" edge to the Antenatalinformation entity by IDs.
func (pu *PregnancystatusUpdate) AddAntenatalinformationIDs(ids ...int) *PregnancystatusUpdate {
	pu.mutation.AddAntenatalinformationIDs(ids...)
	return pu
}

// AddAntenatalinformation adds the "Antenatalinformation" edges to the Antenatalinformation entity.
func (pu *PregnancystatusUpdate) AddAntenatalinformation(a ...*Antenatalinformation) *PregnancystatusUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.AddAntenatalinformationIDs(ids...)
}

// Mutation returns the PregnancystatusMutation object of the builder.
func (pu *PregnancystatusUpdate) Mutation() *PregnancystatusMutation {
	return pu.mutation
}

// ClearAntenatalinformation clears all "Antenatalinformation" edges to the Antenatalinformation entity.
func (pu *PregnancystatusUpdate) ClearAntenatalinformation() *PregnancystatusUpdate {
	pu.mutation.ClearAntenatalinformation()
	return pu
}

// RemoveAntenatalinformationIDs removes the "Antenatalinformation" edge to Antenatalinformation entities by IDs.
func (pu *PregnancystatusUpdate) RemoveAntenatalinformationIDs(ids ...int) *PregnancystatusUpdate {
	pu.mutation.RemoveAntenatalinformationIDs(ids...)
	return pu
}

// RemoveAntenatalinformation removes "Antenatalinformation" edges to Antenatalinformation entities.
func (pu *PregnancystatusUpdate) RemoveAntenatalinformation(a ...*Antenatalinformation) *PregnancystatusUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.RemoveAntenatalinformationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PregnancystatusUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PregnancystatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PregnancystatusUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PregnancystatusUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PregnancystatusUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PregnancystatusUpdate) check() error {
	if v, ok := pu.mutation.Pregnancystatus(); ok {
		if err := pregnancystatus.PregnancystatusValidator(v); err != nil {
			return &ValidationError{Name: "Pregnancystatus", err: fmt.Errorf("ent: validator failed for field \"Pregnancystatus\": %w", err)}
		}
	}
	return nil
}

func (pu *PregnancystatusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pregnancystatus.Table,
			Columns: pregnancystatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pregnancystatus.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Pregnancystatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pregnancystatus.FieldPregnancystatus,
		})
	}
	if pu.mutation.AntenatalinformationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pregnancystatus.AntenatalinformationTable,
			Columns: []string{pregnancystatus.AntenatalinformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: antenatalinformation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedAntenatalinformationIDs(); len(nodes) > 0 && !pu.mutation.AntenatalinformationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pregnancystatus.AntenatalinformationTable,
			Columns: []string{pregnancystatus.AntenatalinformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: antenatalinformation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AntenatalinformationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pregnancystatus.AntenatalinformationTable,
			Columns: []string{pregnancystatus.AntenatalinformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: antenatalinformation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pregnancystatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PregnancystatusUpdateOne is the builder for updating a single Pregnancystatus entity.
type PregnancystatusUpdateOne struct {
	config
	hooks    []Hook
	mutation *PregnancystatusMutation
}

// SetPregnancystatus sets the "Pregnancystatus" field.
func (puo *PregnancystatusUpdateOne) SetPregnancystatus(s string) *PregnancystatusUpdateOne {
	puo.mutation.SetPregnancystatus(s)
	return puo
}

// AddAntenatalinformationIDs adds the "Antenatalinformation" edge to the Antenatalinformation entity by IDs.
func (puo *PregnancystatusUpdateOne) AddAntenatalinformationIDs(ids ...int) *PregnancystatusUpdateOne {
	puo.mutation.AddAntenatalinformationIDs(ids...)
	return puo
}

// AddAntenatalinformation adds the "Antenatalinformation" edges to the Antenatalinformation entity.
func (puo *PregnancystatusUpdateOne) AddAntenatalinformation(a ...*Antenatalinformation) *PregnancystatusUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.AddAntenatalinformationIDs(ids...)
}

// Mutation returns the PregnancystatusMutation object of the builder.
func (puo *PregnancystatusUpdateOne) Mutation() *PregnancystatusMutation {
	return puo.mutation
}

// ClearAntenatalinformation clears all "Antenatalinformation" edges to the Antenatalinformation entity.
func (puo *PregnancystatusUpdateOne) ClearAntenatalinformation() *PregnancystatusUpdateOne {
	puo.mutation.ClearAntenatalinformation()
	return puo
}

// RemoveAntenatalinformationIDs removes the "Antenatalinformation" edge to Antenatalinformation entities by IDs.
func (puo *PregnancystatusUpdateOne) RemoveAntenatalinformationIDs(ids ...int) *PregnancystatusUpdateOne {
	puo.mutation.RemoveAntenatalinformationIDs(ids...)
	return puo
}

// RemoveAntenatalinformation removes "Antenatalinformation" edges to Antenatalinformation entities.
func (puo *PregnancystatusUpdateOne) RemoveAntenatalinformation(a ...*Antenatalinformation) *PregnancystatusUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.RemoveAntenatalinformationIDs(ids...)
}

// Save executes the query and returns the updated Pregnancystatus entity.
func (puo *PregnancystatusUpdateOne) Save(ctx context.Context) (*Pregnancystatus, error) {
	var (
		err  error
		node *Pregnancystatus
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PregnancystatusMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PregnancystatusUpdateOne) SaveX(ctx context.Context) *Pregnancystatus {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PregnancystatusUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PregnancystatusUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PregnancystatusUpdateOne) check() error {
	if v, ok := puo.mutation.Pregnancystatus(); ok {
		if err := pregnancystatus.PregnancystatusValidator(v); err != nil {
			return &ValidationError{Name: "Pregnancystatus", err: fmt.Errorf("ent: validator failed for field \"Pregnancystatus\": %w", err)}
		}
	}
	return nil
}

func (puo *PregnancystatusUpdateOne) sqlSave(ctx context.Context) (_node *Pregnancystatus, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pregnancystatus.Table,
			Columns: pregnancystatus.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pregnancystatus.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Pregnancystatus.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Pregnancystatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pregnancystatus.FieldPregnancystatus,
		})
	}
	if puo.mutation.AntenatalinformationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pregnancystatus.AntenatalinformationTable,
			Columns: []string{pregnancystatus.AntenatalinformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: antenatalinformation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedAntenatalinformationIDs(); len(nodes) > 0 && !puo.mutation.AntenatalinformationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pregnancystatus.AntenatalinformationTable,
			Columns: []string{pregnancystatus.AntenatalinformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: antenatalinformation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AntenatalinformationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pregnancystatus.AntenatalinformationTable,
			Columns: []string{pregnancystatus.AntenatalinformationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: antenatalinformation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Pregnancystatus{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pregnancystatus.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
