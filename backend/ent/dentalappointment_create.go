// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/dentalappointment"
	"github.com/to63/app/ent/dentalkind"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
)

// DentalappointmentCreate is the builder for creating a Dentalappointment entity.
type DentalappointmentCreate struct {
	config
	mutation *DentalappointmentMutation
	hooks    []Hook
}

// SetAppointtime sets the "appointtime" field.
func (dc *DentalappointmentCreate) SetAppointtime(t time.Time) *DentalappointmentCreate {
	dc.mutation.SetAppointtime(t)
	return dc
}

// SetAmount sets the "amount" field.
func (dc *DentalappointmentCreate) SetAmount(i int) *DentalappointmentCreate {
	dc.mutation.SetAmount(i)
	return dc
}

// SetPrice sets the "price" field.
func (dc *DentalappointmentCreate) SetPrice(i int) *DentalappointmentCreate {
	dc.mutation.SetPrice(i)
	return dc
}

// SetNote sets the "note" field.
func (dc *DentalappointmentCreate) SetNote(s string) *DentalappointmentCreate {
	dc.mutation.SetNote(s)
	return dc
}

// SetPersonnelID sets the "Personnel" edge to the Personnel entity by ID.
func (dc *DentalappointmentCreate) SetPersonnelID(id int) *DentalappointmentCreate {
	dc.mutation.SetPersonnelID(id)
	return dc
}

// SetNillablePersonnelID sets the "Personnel" edge to the Personnel entity by ID if the given value is not nil.
func (dc *DentalappointmentCreate) SetNillablePersonnelID(id *int) *DentalappointmentCreate {
	if id != nil {
		dc = dc.SetPersonnelID(*id)
	}
	return dc
}

// SetPersonnel sets the "Personnel" edge to the Personnel entity.
func (dc *DentalappointmentCreate) SetPersonnel(p *Personnel) *DentalappointmentCreate {
	return dc.SetPersonnelID(p.ID)
}

// SetPatientID sets the "Patient" edge to the Patient entity by ID.
func (dc *DentalappointmentCreate) SetPatientID(id int) *DentalappointmentCreate {
	dc.mutation.SetPatientID(id)
	return dc
}

// SetNillablePatientID sets the "Patient" edge to the Patient entity by ID if the given value is not nil.
func (dc *DentalappointmentCreate) SetNillablePatientID(id *int) *DentalappointmentCreate {
	if id != nil {
		dc = dc.SetPatientID(*id)
	}
	return dc
}

// SetPatient sets the "Patient" edge to the Patient entity.
func (dc *DentalappointmentCreate) SetPatient(p *Patient) *DentalappointmentCreate {
	return dc.SetPatientID(p.ID)
}

// SetDentalkindID sets the "Dentalkind" edge to the Dentalkind entity by ID.
func (dc *DentalappointmentCreate) SetDentalkindID(id int) *DentalappointmentCreate {
	dc.mutation.SetDentalkindID(id)
	return dc
}

// SetNillableDentalkindID sets the "Dentalkind" edge to the Dentalkind entity by ID if the given value is not nil.
func (dc *DentalappointmentCreate) SetNillableDentalkindID(id *int) *DentalappointmentCreate {
	if id != nil {
		dc = dc.SetDentalkindID(*id)
	}
	return dc
}

// SetDentalkind sets the "Dentalkind" edge to the Dentalkind entity.
func (dc *DentalappointmentCreate) SetDentalkind(d *Dentalkind) *DentalappointmentCreate {
	return dc.SetDentalkindID(d.ID)
}

// Mutation returns the DentalappointmentMutation object of the builder.
func (dc *DentalappointmentCreate) Mutation() *DentalappointmentMutation {
	return dc.mutation
}

// Save creates the Dentalappointment in the database.
func (dc *DentalappointmentCreate) Save(ctx context.Context) (*Dentalappointment, error) {
	var (
		err  error
		node *Dentalappointment
	)
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DentalappointmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dc.check(); err != nil {
				return nil, err
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DentalappointmentCreate) SaveX(ctx context.Context) *Dentalappointment {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (dc *DentalappointmentCreate) check() error {
	if _, ok := dc.mutation.Appointtime(); !ok {
		return &ValidationError{Name: "appointtime", err: errors.New("ent: missing required field \"appointtime\"")}
	}
	if _, ok := dc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New("ent: missing required field \"amount\"")}
	}
	if v, ok := dc.mutation.Amount(); ok {
		if err := dentalappointment.AmountValidator(v); err != nil {
			return &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	if _, ok := dc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if v, ok := dc.mutation.Price(); ok {
		if err := dentalappointment.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf("ent: validator failed for field \"price\": %w", err)}
		}
	}
	if _, ok := dc.mutation.Note(); !ok {
		return &ValidationError{Name: "note", err: errors.New("ent: missing required field \"note\"")}
	}
	if v, ok := dc.mutation.Note(); ok {
		if err := dentalappointment.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	return nil
}

func (dc *DentalappointmentCreate) sqlSave(ctx context.Context) (*Dentalappointment, error) {
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dc *DentalappointmentCreate) createSpec() (*Dentalappointment, *sqlgraph.CreateSpec) {
	var (
		_node = &Dentalappointment{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dentalappointment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dentalappointment.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Appointtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dentalappointment.FieldAppointtime,
		})
		_node.Appointtime = value
	}
	if value, ok := dc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dentalappointment.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := dc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dentalappointment.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := dc.mutation.Note(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentalappointment.FieldNote,
		})
		_node.Note = value
	}
	if nodes := dc.mutation.PersonnelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalappointment.PersonnelTable,
			Columns: []string{dentalappointment.PersonnelColumn},
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
	if nodes := dc.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalappointment.PatientTable,
			Columns: []string{dentalappointment.PatientColumn},
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
	if nodes := dc.mutation.DentalkindIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dentalappointment.DentalkindTable,
			Columns: []string{dentalappointment.DentalkindColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalkind.FieldID,
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

// DentalappointmentCreateBulk is the builder for creating many Dentalappointment entities in bulk.
type DentalappointmentCreateBulk struct {
	config
	builders []*DentalappointmentCreate
}

// Save creates the Dentalappointment entities in the database.
func (dcb *DentalappointmentCreateBulk) Save(ctx context.Context) ([]*Dentalappointment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dentalappointment, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DentalappointmentMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DentalappointmentCreateBulk) SaveX(ctx context.Context) []*Dentalappointment {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
