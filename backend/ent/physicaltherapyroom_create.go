// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/physicaltherapyrecord"
	"github.com/to63/app/ent/physicaltherapyroom"
)

// PhysicaltherapyroomCreate is the builder for creating a Physicaltherapyroom entity.
type PhysicaltherapyroomCreate struct {
	config
	mutation *PhysicaltherapyroomMutation
	hooks    []Hook
}

// SetPhysicalTherapyRoomName sets the "physical_therapy_room_name" field.
func (pc *PhysicaltherapyroomCreate) SetPhysicalTherapyRoomName(s string) *PhysicaltherapyroomCreate {
	pc.mutation.SetPhysicalTherapyRoomName(s)
	return pc
}

// SetPhysicaltherapyrecordID sets the "physicaltherapyrecord" edge to the Physicaltherapyrecord entity by ID.
func (pc *PhysicaltherapyroomCreate) SetPhysicaltherapyrecordID(id int) *PhysicaltherapyroomCreate {
	pc.mutation.SetPhysicaltherapyrecordID(id)
	return pc
}

// SetNillablePhysicaltherapyrecordID sets the "physicaltherapyrecord" edge to the Physicaltherapyrecord entity by ID if the given value is not nil.
func (pc *PhysicaltherapyroomCreate) SetNillablePhysicaltherapyrecordID(id *int) *PhysicaltherapyroomCreate {
	if id != nil {
		pc = pc.SetPhysicaltherapyrecordID(*id)
	}
	return pc
}

// SetPhysicaltherapyrecord sets the "physicaltherapyrecord" edge to the Physicaltherapyrecord entity.
func (pc *PhysicaltherapyroomCreate) SetPhysicaltherapyrecord(p *Physicaltherapyrecord) *PhysicaltherapyroomCreate {
	return pc.SetPhysicaltherapyrecordID(p.ID)
}

// Mutation returns the PhysicaltherapyroomMutation object of the builder.
func (pc *PhysicaltherapyroomCreate) Mutation() *PhysicaltherapyroomMutation {
	return pc.mutation
}

// Save creates the Physicaltherapyroom in the database.
func (pc *PhysicaltherapyroomCreate) Save(ctx context.Context) (*Physicaltherapyroom, error) {
	var (
		err  error
		node *Physicaltherapyroom
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PhysicaltherapyroomMutation)
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
func (pc *PhysicaltherapyroomCreate) SaveX(ctx context.Context) *Physicaltherapyroom {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (pc *PhysicaltherapyroomCreate) check() error {
	if _, ok := pc.mutation.PhysicalTherapyRoomName(); !ok {
		return &ValidationError{Name: "physical_therapy_room_name", err: errors.New("ent: missing required field \"physical_therapy_room_name\"")}
	}
	if v, ok := pc.mutation.PhysicalTherapyRoomName(); ok {
		if err := physicaltherapyroom.PhysicalTherapyRoomNameValidator(v); err != nil {
			return &ValidationError{Name: "physical_therapy_room_name", err: fmt.Errorf("ent: validator failed for field \"physical_therapy_room_name\": %w", err)}
		}
	}
	return nil
}

func (pc *PhysicaltherapyroomCreate) sqlSave(ctx context.Context) (*Physicaltherapyroom, error) {
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

func (pc *PhysicaltherapyroomCreate) createSpec() (*Physicaltherapyroom, *sqlgraph.CreateSpec) {
	var (
		_node = &Physicaltherapyroom{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: physicaltherapyroom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: physicaltherapyroom.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PhysicalTherapyRoomName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: physicaltherapyroom.FieldPhysicalTherapyRoomName,
		})
		_node.PhysicalTherapyRoomName = value
	}
	if nodes := pc.mutation.PhysicaltherapyrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   physicaltherapyroom.PhysicaltherapyrecordTable,
			Columns: []string{physicaltherapyroom.PhysicaltherapyrecordColumn},
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
	return _node, _spec
}

// PhysicaltherapyroomCreateBulk is the builder for creating many Physicaltherapyroom entities in bulk.
type PhysicaltherapyroomCreateBulk struct {
	config
	builders []*PhysicaltherapyroomCreate
}

// Save creates the Physicaltherapyroom entities in the database.
func (pcb *PhysicaltherapyroomCreateBulk) Save(ctx context.Context) ([]*Physicaltherapyroom, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Physicaltherapyroom, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PhysicaltherapyroomMutation)
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
func (pcb *PhysicaltherapyroomCreateBulk) SaveX(ctx context.Context) []*Physicaltherapyroom {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
