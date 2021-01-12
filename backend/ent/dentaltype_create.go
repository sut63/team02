// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/dentalappointment"
	"github.com/to63/app/ent/dentaltype"
)

// DentaltypeCreate is the builder for creating a Dentaltype entity.
type DentaltypeCreate struct {
	config
	mutation *DentaltypeMutation
	hooks    []Hook
}

// SetTypename sets the "typename" field.
func (dc *DentaltypeCreate) SetTypename(s string) *DentaltypeCreate {
	dc.mutation.SetTypename(s)
	return dc
}

// AddDentalappointmentIDs adds the "Dentalappointment" edge to the Dentalappointment entity by IDs.
func (dc *DentaltypeCreate) AddDentalappointmentIDs(ids ...int) *DentaltypeCreate {
	dc.mutation.AddDentalappointmentIDs(ids...)
	return dc
}

// AddDentalappointment adds the "Dentalappointment" edges to the Dentalappointment entity.
func (dc *DentaltypeCreate) AddDentalappointment(d ...*Dentalappointment) *DentaltypeCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDentalappointmentIDs(ids...)
}

// Mutation returns the DentaltypeMutation object of the builder.
func (dc *DentaltypeCreate) Mutation() *DentaltypeMutation {
	return dc.mutation
}

// Save creates the Dentaltype in the database.
func (dc *DentaltypeCreate) Save(ctx context.Context) (*Dentaltype, error) {
	var (
		err  error
		node *Dentaltype
	)
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DentaltypeMutation)
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
func (dc *DentaltypeCreate) SaveX(ctx context.Context) *Dentaltype {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (dc *DentaltypeCreate) check() error {
	if _, ok := dc.mutation.Typename(); !ok {
		return &ValidationError{Name: "typename", err: errors.New("ent: missing required field \"typename\"")}
	}
	return nil
}

func (dc *DentaltypeCreate) sqlSave(ctx context.Context) (*Dentaltype, error) {
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

func (dc *DentaltypeCreate) createSpec() (*Dentaltype, *sqlgraph.CreateSpec) {
	var (
		_node = &Dentaltype{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dentaltype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dentaltype.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Typename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dentaltype.FieldTypename,
		})
		_node.Typename = value
	}
	if nodes := dc.mutation.DentalappointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dentaltype.DentalappointmentTable,
			Columns: []string{dentaltype.DentalappointmentColumn},
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

// DentaltypeCreateBulk is the builder for creating many Dentaltype entities in bulk.
type DentaltypeCreateBulk struct {
	config
	builders []*DentaltypeCreate
}

// Save creates the Dentaltype entities in the database.
func (dcb *DentaltypeCreateBulk) Save(ctx context.Context) ([]*Dentaltype, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dentaltype, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DentaltypeMutation)
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
func (dcb *DentaltypeCreateBulk) SaveX(ctx context.Context) []*Dentaltype {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}