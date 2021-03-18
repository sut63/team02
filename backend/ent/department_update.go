// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/department"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/predicate"
)

// DepartmentUpdate is the builder for updating Department entities.
type DepartmentUpdate struct {
	config
	hooks    []Hook
	mutation *DepartmentMutation
}

// Where adds a new predicate for the DepartmentUpdate builder.
func (du *DepartmentUpdate) Where(ps ...predicate.Department) *DepartmentUpdate {
	du.mutation.predicates = append(du.mutation.predicates, ps...)
	return du
}

// SetDepartment sets the "department" field.
func (du *DepartmentUpdate) SetDepartment(s string) *DepartmentUpdate {
	du.mutation.SetDepartment(s)
	return du
}

// AddPersonnelIDs adds the "Personnel" edge to the Personnel entity by IDs.
func (du *DepartmentUpdate) AddPersonnelIDs(ids ...int) *DepartmentUpdate {
	du.mutation.AddPersonnelIDs(ids...)
	return du
}

// AddPersonnel adds the "Personnel" edges to the Personnel entity.
func (du *DepartmentUpdate) AddPersonnel(p ...*Personnel) *DepartmentUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.AddPersonnelIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (du *DepartmentUpdate) Mutation() *DepartmentMutation {
	return du.mutation
}

// ClearPersonnel clears all "Personnel" edges to the Personnel entity.
func (du *DepartmentUpdate) ClearPersonnel() *DepartmentUpdate {
	du.mutation.ClearPersonnel()
	return du
}

// RemovePersonnelIDs removes the "Personnel" edge to Personnel entities by IDs.
func (du *DepartmentUpdate) RemovePersonnelIDs(ids ...int) *DepartmentUpdate {
	du.mutation.RemovePersonnelIDs(ids...)
	return du
}

// RemovePersonnel removes "Personnel" edges to Personnel entities.
func (du *DepartmentUpdate) RemovePersonnel(p ...*Personnel) *DepartmentUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.RemovePersonnelIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DepartmentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
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
func (du *DepartmentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DepartmentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DepartmentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DepartmentUpdate) check() error {
	if v, ok := du.mutation.Department(); ok {
		if err := department.DepartmentValidator(v); err != nil {
			return &ValidationError{Name: "department", err: fmt.Errorf("ent: validator failed for field \"department\": %w", err)}
		}
	}
	return nil
}

func (du *DepartmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: department.FieldID,
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
	if value, ok := du.mutation.Department(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldDepartment,
		})
	}
	if du.mutation.PersonnelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.PersonnelTable,
			Columns: []string{department.PersonnelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: personnel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedPersonnelIDs(); len(nodes) > 0 && !du.mutation.PersonnelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.PersonnelTable,
			Columns: []string{department.PersonnelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: personnel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.PersonnelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.PersonnelTable,
			Columns: []string{department.PersonnelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: personnel.FieldID,
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
			err = &NotFoundError{department.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DepartmentUpdateOne is the builder for updating a single Department entity.
type DepartmentUpdateOne struct {
	config
	hooks    []Hook
	mutation *DepartmentMutation
}

// SetDepartment sets the "department" field.
func (duo *DepartmentUpdateOne) SetDepartment(s string) *DepartmentUpdateOne {
	duo.mutation.SetDepartment(s)
	return duo
}

// AddPersonnelIDs adds the "Personnel" edge to the Personnel entity by IDs.
func (duo *DepartmentUpdateOne) AddPersonnelIDs(ids ...int) *DepartmentUpdateOne {
	duo.mutation.AddPersonnelIDs(ids...)
	return duo
}

// AddPersonnel adds the "Personnel" edges to the Personnel entity.
func (duo *DepartmentUpdateOne) AddPersonnel(p ...*Personnel) *DepartmentUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.AddPersonnelIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (duo *DepartmentUpdateOne) Mutation() *DepartmentMutation {
	return duo.mutation
}

// ClearPersonnel clears all "Personnel" edges to the Personnel entity.
func (duo *DepartmentUpdateOne) ClearPersonnel() *DepartmentUpdateOne {
	duo.mutation.ClearPersonnel()
	return duo
}

// RemovePersonnelIDs removes the "Personnel" edge to Personnel entities by IDs.
func (duo *DepartmentUpdateOne) RemovePersonnelIDs(ids ...int) *DepartmentUpdateOne {
	duo.mutation.RemovePersonnelIDs(ids...)
	return duo
}

// RemovePersonnel removes "Personnel" edges to Personnel entities.
func (duo *DepartmentUpdateOne) RemovePersonnel(p ...*Personnel) *DepartmentUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.RemovePersonnelIDs(ids...)
}

// Save executes the query and returns the updated Department entity.
func (duo *DepartmentUpdateOne) Save(ctx context.Context) (*Department, error) {
	var (
		err  error
		node *Department
	)
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
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
func (duo *DepartmentUpdateOne) SaveX(ctx context.Context) *Department {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DepartmentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DepartmentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DepartmentUpdateOne) check() error {
	if v, ok := duo.mutation.Department(); ok {
		if err := department.DepartmentValidator(v); err != nil {
			return &ValidationError{Name: "department", err: fmt.Errorf("ent: validator failed for field \"department\": %w", err)}
		}
	}
	return nil
}

func (duo *DepartmentUpdateOne) sqlSave(ctx context.Context) (_node *Department, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: department.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Department.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.Department(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldDepartment,
		})
	}
	if duo.mutation.PersonnelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.PersonnelTable,
			Columns: []string{department.PersonnelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: personnel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedPersonnelIDs(); len(nodes) > 0 && !duo.mutation.PersonnelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.PersonnelTable,
			Columns: []string{department.PersonnelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: personnel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.PersonnelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.PersonnelTable,
			Columns: []string{department.PersonnelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: personnel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Department{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{department.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}