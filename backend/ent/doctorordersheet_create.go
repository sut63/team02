// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/checksymptoms"
	"github.com/to63/app/ent/doctorordersheet"
)

// DoctorOrderSheetCreate is the builder for creating a DoctorOrderSheet entity.
type DoctorOrderSheetCreate struct {
	config
	mutation *DoctorOrderSheetMutation
	hooks    []Hook
}

// SetDate sets the "date" field.
func (dosc *DoctorOrderSheetCreate) SetDate(t time.Time) *DoctorOrderSheetCreate {
	dosc.mutation.SetDate(t)
	return dosc
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (dosc *DoctorOrderSheetCreate) SetNillableDate(t *time.Time) *DoctorOrderSheetCreate {
	if t != nil {
		dosc.SetDate(*t)
	}
	return dosc
}

// SetTime sets the "time" field.
func (dosc *DoctorOrderSheetCreate) SetTime(s string) *DoctorOrderSheetCreate {
	dosc.mutation.SetTime(s)
	return dosc
}

// AddChecksymptomIDs adds the "Checksymptoms" edge to the Checksymptoms entity by IDs.
func (dosc *DoctorOrderSheetCreate) AddChecksymptomIDs(ids ...int) *DoctorOrderSheetCreate {
	dosc.mutation.AddChecksymptomIDs(ids...)
	return dosc
}

// AddChecksymptoms adds the "Checksymptoms" edges to the Checksymptoms entity.
func (dosc *DoctorOrderSheetCreate) AddChecksymptoms(c ...*Checksymptoms) *DoctorOrderSheetCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return dosc.AddChecksymptomIDs(ids...)
}

// Mutation returns the DoctorOrderSheetMutation object of the builder.
func (dosc *DoctorOrderSheetCreate) Mutation() *DoctorOrderSheetMutation {
	return dosc.mutation
}

// Save creates the DoctorOrderSheet in the database.
func (dosc *DoctorOrderSheetCreate) Save(ctx context.Context) (*DoctorOrderSheet, error) {
	var (
		err  error
		node *DoctorOrderSheet
	)
	dosc.defaults()
	if len(dosc.hooks) == 0 {
		if err = dosc.check(); err != nil {
			return nil, err
		}
		node, err = dosc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorOrderSheetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dosc.check(); err != nil {
				return nil, err
			}
			dosc.mutation = mutation
			node, err = dosc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dosc.hooks) - 1; i >= 0; i-- {
			mut = dosc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dosc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dosc *DoctorOrderSheetCreate) SaveX(ctx context.Context) *DoctorOrderSheet {
	v, err := dosc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (dosc *DoctorOrderSheetCreate) defaults() {
	if _, ok := dosc.mutation.Date(); !ok {
		v := doctorordersheet.DefaultDate()
		dosc.mutation.SetDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dosc *DoctorOrderSheetCreate) check() error {
	if _, ok := dosc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New("ent: missing required field \"date\"")}
	}
	if _, ok := dosc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New("ent: missing required field \"time\"")}
	}
	if v, ok := dosc.mutation.Time(); ok {
		if err := doctorordersheet.TimeValidator(v); err != nil {
			return &ValidationError{Name: "time", err: fmt.Errorf("ent: validator failed for field \"time\": %w", err)}
		}
	}
	return nil
}

func (dosc *DoctorOrderSheetCreate) sqlSave(ctx context.Context) (*DoctorOrderSheet, error) {
	_node, _spec := dosc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dosc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (dosc *DoctorOrderSheetCreate) createSpec() (*DoctorOrderSheet, *sqlgraph.CreateSpec) {
	var (
		_node = &DoctorOrderSheet{config: dosc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: doctorordersheet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctorordersheet.FieldID,
			},
		}
	)
	if value, ok := dosc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: doctorordersheet.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := dosc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctorordersheet.FieldTime,
		})
		_node.Time = value
	}
	if nodes := dosc.mutation.ChecksymptomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctorordersheet.ChecksymptomsTable,
			Columns: []string{doctorordersheet.ChecksymptomsColumn},
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

// DoctorOrderSheetCreateBulk is the builder for creating many DoctorOrderSheet entities in bulk.
type DoctorOrderSheetCreateBulk struct {
	config
	builders []*DoctorOrderSheetCreate
}

// Save creates the DoctorOrderSheet entities in the database.
func (doscb *DoctorOrderSheetCreateBulk) Save(ctx context.Context) ([]*DoctorOrderSheet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(doscb.builders))
	nodes := make([]*DoctorOrderSheet, len(doscb.builders))
	mutators := make([]Mutator, len(doscb.builders))
	for i := range doscb.builders {
		func(i int, root context.Context) {
			builder := doscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DoctorOrderSheetMutation)
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
					_, err = mutators[i+1].Mutate(root, doscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, doscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, doscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (doscb *DoctorOrderSheetCreateBulk) SaveX(ctx context.Context) []*DoctorOrderSheet {
	v, err := doscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}