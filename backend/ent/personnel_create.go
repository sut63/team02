// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/bonedisease"
	"github.com/to63/app/ent/checksymptoms"
	"github.com/to63/app/ent/dentalappointment"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/physicaltherapyrecord"
)

// PersonnelCreate is the builder for creating a Personnel entity.
type PersonnelCreate struct {
	config
	mutation *PersonnelMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PersonnelCreate) SetName(s string) *PersonnelCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDepartment sets the "department" field.
func (pc *PersonnelCreate) SetDepartment(s string) *PersonnelCreate {
	pc.mutation.SetDepartment(s)
	return pc
}

// SetUser sets the "user" field.
func (pc *PersonnelCreate) SetUser(s string) *PersonnelCreate {
	pc.mutation.SetUser(s)
	return pc
}

// SetPassword sets the "password" field.
func (pc *PersonnelCreate) SetPassword(s string) *PersonnelCreate {
	pc.mutation.SetPassword(s)
	return pc
}

// AddPhysicaltherapyrecordIDs adds the "physicaltherapyrecord" edge to the Physicaltherapyrecord entity by IDs.
func (pc *PersonnelCreate) AddPhysicaltherapyrecordIDs(ids ...int) *PersonnelCreate {
	pc.mutation.AddPhysicaltherapyrecordIDs(ids...)
	return pc
}

// AddPhysicaltherapyrecord adds the "physicaltherapyrecord" edges to the Physicaltherapyrecord entity.
func (pc *PersonnelCreate) AddPhysicaltherapyrecord(p ...*Physicaltherapyrecord) *PersonnelCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPhysicaltherapyrecordIDs(ids...)
}

// AddBonediseaseIDs adds the "Bonedisease" edge to the Bonedisease entity by IDs.
func (pc *PersonnelCreate) AddBonediseaseIDs(ids ...int) *PersonnelCreate {
	pc.mutation.AddBonediseaseIDs(ids...)
	return pc
}

// AddBonedisease adds the "Bonedisease" edges to the Bonedisease entity.
func (pc *PersonnelCreate) AddBonedisease(b ...*Bonedisease) *PersonnelCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pc.AddBonediseaseIDs(ids...)
}

// AddChecksymptomIDs adds the "Checksymptoms" edge to the Checksymptoms entity by IDs.
func (pc *PersonnelCreate) AddChecksymptomIDs(ids ...int) *PersonnelCreate {
	pc.mutation.AddChecksymptomIDs(ids...)
	return pc
}

// AddChecksymptoms adds the "Checksymptoms" edges to the Checksymptoms entity.
func (pc *PersonnelCreate) AddChecksymptoms(c ...*Checksymptoms) *PersonnelCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddChecksymptomIDs(ids...)
}

// AddDentalappointmentIDs adds the "Dentalappointment" edge to the Dentalappointment entity by IDs.
func (pc *PersonnelCreate) AddDentalappointmentIDs(ids ...int) *PersonnelCreate {
	pc.mutation.AddDentalappointmentIDs(ids...)
	return pc
}

// AddDentalappointment adds the "Dentalappointment" edges to the Dentalappointment entity.
func (pc *PersonnelCreate) AddDentalappointment(d ...*Dentalappointment) *PersonnelCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddDentalappointmentIDs(ids...)
}

// Mutation returns the PersonnelMutation object of the builder.
func (pc *PersonnelCreate) Mutation() *PersonnelMutation {
	return pc.mutation
}

// Save creates the Personnel in the database.
func (pc *PersonnelCreate) Save(ctx context.Context) (*Personnel, error) {
	var (
		err  error
		node *Personnel
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonnelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PersonnelCreate) SaveX(ctx context.Context) *Personnel {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (pc *PersonnelCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := personnel.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Department(); !ok {
		return &ValidationError{Name: "department", err: errors.New("ent: missing required field \"department\"")}
	}
	if v, ok := pc.mutation.Department(); ok {
		if err := personnel.DepartmentValidator(v); err != nil {
			return &ValidationError{Name: "department", err: fmt.Errorf("ent: validator failed for field \"department\": %w", err)}
		}
	}
	if _, ok := pc.mutation.User(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required field \"user\"")}
	}
	if v, ok := pc.mutation.User(); ok {
		if err := personnel.UserValidator(v); err != nil {
			return &ValidationError{Name: "user", err: fmt.Errorf("ent: validator failed for field \"user\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if v, ok := pc.mutation.Password(); ok {
		if err := personnel.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	return nil
}

func (pc *PersonnelCreate) sqlSave(ctx context.Context) (*Personnel, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PersonnelCreate) createSpec() (*Personnel, *sqlgraph.CreateSpec) {
	var (
		_node = &Personnel{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: personnel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: personnel.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personnel.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.Department(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personnel.FieldDepartment,
		})
		_node.Department = value
	}
	if value, ok := pc.mutation.User(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personnel.FieldUser,
		})
		_node.User = value
	}
	if value, ok := pc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: personnel.FieldPassword,
		})
		_node.Password = value
	}
	if nodes := pc.mutation.PhysicaltherapyrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   personnel.PhysicaltherapyrecordTable,
			Columns: []string{personnel.PhysicaltherapyrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: physicaltherapyrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BonediseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   personnel.BonediseaseTable,
			Columns: []string{personnel.BonediseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bonedisease.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ChecksymptomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   personnel.ChecksymptomsTable,
			Columns: []string{personnel.ChecksymptomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checksymptoms.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.DentalappointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   personnel.DentalappointmentTable,
			Columns: []string{personnel.DentalappointmentColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PersonnelCreateBulk is the builder for creating many Personnel entities in bulk.
type PersonnelCreateBulk struct {
	config
	builders []*PersonnelCreate
}

// Save creates the Personnel entities in the database.
func (pcb *PersonnelCreateBulk) Save(ctx context.Context) ([]*Personnel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Personnel, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonnelMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PersonnelCreateBulk) SaveX(ctx context.Context) []*Personnel {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
