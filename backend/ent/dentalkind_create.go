// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/dentalappointment"
	"github.com/to63/app/ent/dentalkind"
)

// DentalkindCreate is the builder for creating a Dentalkind entity.
type DentalkindCreate struct {
	config
	mutation *DentalkindMutation
	hooks    []Hook
}

// SetKindname sets the "kindname" field.
func (dc *DentalkindCreate) SetKindname(s string) *DentalkindCreate {
	dc.mutation.SetKindname(s)
	return dc
}

// AddDentalappointmentIDs adds the "Dentalappointment" edge to the Dentalappointment entity by IDs.
func (dc *DentalkindCreate) AddDentalappointmentIDs(ids ...int) *DentalkindCreate {
	dc.mutation.AddDentalappointmentIDs(ids...)
	return dc
}

// AddDentalappointment adds the "Dentalappointment" edges to the Dentalappointment entity.
func (dc *DentalkindCreate) AddDentalappointment(d ...*Dentalappointment) *DentalkindCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDentalappointmentIDs(ids...)
}

// Mutation returns the DentalkindMutation object of the builder.
func (dc *DentalkindCreate) Mutation() *DentalkindMutation {
	return dc.mutation
}

// Save creates the Dentalkind in the database.
func (dc *DentalkindCreate) Save(ctx context.Context) (*Dentalkind, error) {
	var (
		err  error
		node *Dentalkind
	)
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DentalkindMutation)
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
func (dc *DentalkindCreate) SaveX(ctx context.Context) *Dentalkind {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (dc *DentalkindCreate) check() error {
	if _, ok := dc.mutation.Kindname(); !ok {
		return &ValidationError{Name: "kindname", err: errors.New("ent: missing required field \"kindname\"")}
	}
	return nil
}

func (dc *DentalkindCreate) sqlSave(ctx context.Context) (*Dentalkind, error) {
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

func (dc *DentalkindCreate) createSpec() (*Dentalkind, *sqlgraph.CreateSpec) {
	var (
		_node = &Dentalkind{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dentalkind.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dentalkind.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Kindname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentalkind.FieldKindname,
		})
		_node.Kindname = value
	}
	if nodes := dc.mutation.DentalappointmentIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DentalkindCreateBulk is the builder for creating many Dentalkind entities in bulk.
type DentalkindCreateBulk struct {
	config
	builders []*DentalkindCreate
}

// Save creates the Dentalkind entities in the database.
func (dcb *DentalkindCreateBulk) Save(ctx context.Context) ([]*Dentalkind, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dentalkind, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DentalkindMutation)
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
func (dcb *DentalkindCreateBulk) SaveX(ctx context.Context) []*Dentalkind {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}