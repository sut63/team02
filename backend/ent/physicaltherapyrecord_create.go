// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/physicaltherapyrecord"
	"github.com/to63/app/ent/physicaltherapyroom"
	"github.com/to63/app/ent/status"
)

// PhysicaltherapyrecordCreate is the builder for creating a Physicaltherapyrecord entity.
type PhysicaltherapyrecordCreate struct {
	config
	mutation *PhysicaltherapyrecordMutation
	hooks    []Hook
}

// SetPrice sets the "price" field.
func (pc *PhysicaltherapyrecordCreate) SetPrice(i int) *PhysicaltherapyrecordCreate {
	pc.mutation.SetPrice(i)
	return pc
}

// SetPhone sets the "phone" field.
func (pc *PhysicaltherapyrecordCreate) SetPhone(s string) *PhysicaltherapyrecordCreate {
	pc.mutation.SetPhone(s)
	return pc
}

// SetNote sets the "note" field.
func (pc *PhysicaltherapyrecordCreate) SetNote(s string) *PhysicaltherapyrecordCreate {
	pc.mutation.SetNote(s)
	return pc
}

// SetAppointtime sets the "appointtime" field.
func (pc *PhysicaltherapyrecordCreate) SetAppointtime(t time.Time) *PhysicaltherapyrecordCreate {
	pc.mutation.SetAppointtime(t)
	return pc
}

// SetPersonnelID sets the "personnel" edge to the Personnel entity by ID.
func (pc *PhysicaltherapyrecordCreate) SetPersonnelID(id int) *PhysicaltherapyrecordCreate {
	pc.mutation.SetPersonnelID(id)
	return pc
}

// SetNillablePersonnelID sets the "personnel" edge to the Personnel entity by ID if the given value is not nil.
func (pc *PhysicaltherapyrecordCreate) SetNillablePersonnelID(id *int) *PhysicaltherapyrecordCreate {
	if id != nil {
		pc = pc.SetPersonnelID(*id)
	}
	return pc
}

// SetPersonnel sets the "personnel" edge to the Personnel entity.
func (pc *PhysicaltherapyrecordCreate) SetPersonnel(p *Personnel) *PhysicaltherapyrecordCreate {
	return pc.SetPersonnelID(p.ID)
}

// SetPatientID sets the "patient" edge to the Patient entity by ID.
func (pc *PhysicaltherapyrecordCreate) SetPatientID(id int) *PhysicaltherapyrecordCreate {
	pc.mutation.SetPatientID(id)
	return pc
}

// SetNillablePatientID sets the "patient" edge to the Patient entity by ID if the given value is not nil.
func (pc *PhysicaltherapyrecordCreate) SetNillablePatientID(id *int) *PhysicaltherapyrecordCreate {
	if id != nil {
		pc = pc.SetPatientID(*id)
	}
	return pc
}

// SetPatient sets the "patient" edge to the Patient entity.
func (pc *PhysicaltherapyrecordCreate) SetPatient(p *Patient) *PhysicaltherapyrecordCreate {
	return pc.SetPatientID(p.ID)
}

// SetPhysicaltherapyroomID sets the "physicaltherapyroom" edge to the Physicaltherapyroom entity by ID.
func (pc *PhysicaltherapyrecordCreate) SetPhysicaltherapyroomID(id int) *PhysicaltherapyrecordCreate {
	pc.mutation.SetPhysicaltherapyroomID(id)
	return pc
}

// SetNillablePhysicaltherapyroomID sets the "physicaltherapyroom" edge to the Physicaltherapyroom entity by ID if the given value is not nil.
func (pc *PhysicaltherapyrecordCreate) SetNillablePhysicaltherapyroomID(id *int) *PhysicaltherapyrecordCreate {
	if id != nil {
		pc = pc.SetPhysicaltherapyroomID(*id)
	}
	return pc
}

// SetPhysicaltherapyroom sets the "physicaltherapyroom" edge to the Physicaltherapyroom entity.
func (pc *PhysicaltherapyrecordCreate) SetPhysicaltherapyroom(p *Physicaltherapyroom) *PhysicaltherapyrecordCreate {
	return pc.SetPhysicaltherapyroomID(p.ID)
}

// SetStatusID sets the "status" edge to the Status entity by ID.
func (pc *PhysicaltherapyrecordCreate) SetStatusID(id int) *PhysicaltherapyrecordCreate {
	pc.mutation.SetStatusID(id)
	return pc
}

// SetNillableStatusID sets the "status" edge to the Status entity by ID if the given value is not nil.
func (pc *PhysicaltherapyrecordCreate) SetNillableStatusID(id *int) *PhysicaltherapyrecordCreate {
	if id != nil {
		pc = pc.SetStatusID(*id)
	}
	return pc
}

// SetStatus sets the "status" edge to the Status entity.
func (pc *PhysicaltherapyrecordCreate) SetStatus(s *Status) *PhysicaltherapyrecordCreate {
	return pc.SetStatusID(s.ID)
}

// Mutation returns the PhysicaltherapyrecordMutation object of the builder.
func (pc *PhysicaltherapyrecordCreate) Mutation() *PhysicaltherapyrecordMutation {
	return pc.mutation
}

// Save creates the Physicaltherapyrecord in the database.
func (pc *PhysicaltherapyrecordCreate) Save(ctx context.Context) (*Physicaltherapyrecord, error) {
	var (
		err  error
		node *Physicaltherapyrecord
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PhysicaltherapyrecordMutation)
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
func (pc *PhysicaltherapyrecordCreate) SaveX(ctx context.Context) *Physicaltherapyrecord {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (pc *PhysicaltherapyrecordCreate) check() error {
	if _, ok := pc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if v, ok := pc.mutation.Price(); ok {
		if err := physicaltherapyrecord.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf("ent: validator failed for field \"price\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New("ent: missing required field \"phone\"")}
	}
	if v, ok := pc.mutation.Phone(); ok {
		if err := physicaltherapyrecord.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf("ent: validator failed for field \"phone\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Note(); !ok {
		return &ValidationError{Name: "note", err: errors.New("ent: missing required field \"note\"")}
	}
	if v, ok := pc.mutation.Note(); ok {
		if err := physicaltherapyrecord.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Appointtime(); !ok {
		return &ValidationError{Name: "appointtime", err: errors.New("ent: missing required field \"appointtime\"")}
	}
	return nil
}

func (pc *PhysicaltherapyrecordCreate) sqlSave(ctx context.Context) (*Physicaltherapyrecord, error) {
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

func (pc *PhysicaltherapyrecordCreate) createSpec() (*Physicaltherapyrecord, *sqlgraph.CreateSpec) {
	var (
		_node = &Physicaltherapyrecord{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: physicaltherapyrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: physicaltherapyrecord.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: physicaltherapyrecord.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := pc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: physicaltherapyrecord.FieldPhone,
		})
		_node.Phone = value
	}
	if value, ok := pc.mutation.Note(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: physicaltherapyrecord.FieldNote,
		})
		_node.Note = value
	}
	if value, ok := pc.mutation.Appointtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: physicaltherapyrecord.FieldAppointtime,
		})
		_node.Appointtime = value
	}
	if nodes := pc.mutation.PersonnelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   physicaltherapyrecord.PersonnelTable,
			Columns: []string{physicaltherapyrecord.PersonnelColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   physicaltherapyrecord.PatientTable,
			Columns: []string{physicaltherapyrecord.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PhysicaltherapyroomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   physicaltherapyrecord.PhysicaltherapyroomTable,
			Columns: []string{physicaltherapyrecord.PhysicaltherapyroomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: physicaltherapyroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.StatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   physicaltherapyrecord.StatusTable,
			Columns: []string{physicaltherapyrecord.StatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: status.FieldID,
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

// PhysicaltherapyrecordCreateBulk is the builder for creating many Physicaltherapyrecord entities in bulk.
type PhysicaltherapyrecordCreateBulk struct {
	config
	builders []*PhysicaltherapyrecordCreate
}

// Save creates the Physicaltherapyrecord entities in the database.
func (pcb *PhysicaltherapyrecordCreateBulk) Save(ctx context.Context) ([]*Physicaltherapyrecord, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Physicaltherapyrecord, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PhysicaltherapyrecordMutation)
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
func (pcb *PhysicaltherapyrecordCreateBulk) SaveX(ctx context.Context) []*Physicaltherapyrecord {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
