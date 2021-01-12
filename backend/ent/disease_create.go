// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/checksymptoms"
	"github.com/to63/app/ent/disease"
)

// DiseaseCreate is the builder for creating a Disease entity.
type DiseaseCreate struct {
	config
	mutation *DiseaseMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dc *DiseaseCreate) SetName(s string) *DiseaseCreate {
	dc.mutation.SetName(s)
	return dc
}

// AddChecksymptomIDs adds the "Checksymptoms" edge to the Checksymptoms entity by IDs.
func (dc *DiseaseCreate) AddChecksymptomIDs(ids ...int) *DiseaseCreate {
	dc.mutation.AddChecksymptomIDs(ids...)
	return dc
}

// AddChecksymptoms adds the "Checksymptoms" edges to the Checksymptoms entity.
func (dc *DiseaseCreate) AddChecksymptoms(c ...*Checksymptoms) *DiseaseCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dc.AddChecksymptomIDs(ids...)
}

// Mutation returns the DiseaseMutation object of the builder.
func (dc *DiseaseCreate) Mutation() *DiseaseMutation {
	return dc.mutation
}

// Save creates the Disease in the database.
func (dc *DiseaseCreate) Save(ctx context.Context) (*Disease, error) {
	var (
		err  error
		node *Disease
	)
	if len(dc.hooks) == 0 {
		if err = dc.check(); err != nil {
			return nil, err
		}
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiseaseMutation)
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
func (dc *DiseaseCreate) SaveX(ctx context.Context) *Disease {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (dc *DiseaseCreate) check() error {
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := disease.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (dc *DiseaseCreate) sqlSave(ctx context.Context) (*Disease, error) {
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

func (dc *DiseaseCreate) createSpec() (*Disease, *sqlgraph.CreateSpec) {
	var (
		_node = &Disease{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: disease.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: disease.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: disease.FieldName,
		})
		_node.Name = value
	}
	if nodes := dc.mutation.ChecksymptomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   disease.ChecksymptomsTable,
			Columns: []string{disease.ChecksymptomsColumn},
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
	return _node, _spec
}

// DiseaseCreateBulk is the builder for creating many Disease entities in bulk.
type DiseaseCreateBulk struct {
	config
	builders []*DiseaseCreate
}

// Save creates the Disease entities in the database.
func (dcb *DiseaseCreateBulk) Save(ctx context.Context) ([]*Disease, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Disease, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiseaseMutation)
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
func (dcb *DiseaseCreateBulk) SaveX(ctx context.Context) []*Disease {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}