// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/bonedisease"
	"github.com/to63/app/ent/checksymptoms"
	"github.com/to63/app/ent/dentalappointment"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/physicaltherapyrecord"
	"github.com/to63/app/ent/predicate"
)

// PatientUpdate is the builder for updating Patient entities.
type PatientUpdate struct {
	config
	hooks    []Hook
	mutation *PatientMutation
}

// Where adds a new predicate for the PatientUpdate builder.
func (pu *PatientUpdate) Where(ps ...predicate.Patient) *PatientUpdate {
	pu.mutation.predicates = append(pu.mutation.predicates, ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PatientUpdate) SetName(s string) *PatientUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetBirthday sets the "birthday" field.
func (pu *PatientUpdate) SetBirthday(s string) *PatientUpdate {
	pu.mutation.SetBirthday(s)
	return pu
}

// SetGender sets the "gender" field.
func (pu *PatientUpdate) SetGender(s string) *PatientUpdate {
	pu.mutation.SetGender(s)
	return pu
}

// AddPhysicaltherapyrecordIDs adds the "physicaltherapyrecord" edge to the Physicaltherapyrecord entity by IDs.
func (pu *PatientUpdate) AddPhysicaltherapyrecordIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddPhysicaltherapyrecordIDs(ids...)
	return pu
}

// AddPhysicaltherapyrecord adds the "physicaltherapyrecord" edges to the Physicaltherapyrecord entity.
func (pu *PatientUpdate) AddPhysicaltherapyrecord(p ...*Physicaltherapyrecord) *PatientUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPhysicaltherapyrecordIDs(ids...)
}

// AddBonediseaseIDs adds the "Bonedisease" edge to the Bonedisease entity by IDs.
func (pu *PatientUpdate) AddBonediseaseIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddBonediseaseIDs(ids...)
	return pu
}

// AddBonedisease adds the "Bonedisease" edges to the Bonedisease entity.
func (pu *PatientUpdate) AddBonedisease(b ...*Bonedisease) *PatientUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.AddBonediseaseIDs(ids...)
}

// AddChecksymptomIDs adds the "Checksymptoms" edge to the Checksymptoms entity by IDs.
func (pu *PatientUpdate) AddChecksymptomIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddChecksymptomIDs(ids...)
	return pu
}

// AddChecksymptoms adds the "Checksymptoms" edges to the Checksymptoms entity.
func (pu *PatientUpdate) AddChecksymptoms(c ...*Checksymptoms) *PatientUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddChecksymptomIDs(ids...)
}

// AddDentalappointmentIDs adds the "Dentalappointment" edge to the Dentalappointment entity by IDs.
func (pu *PatientUpdate) AddDentalappointmentIDs(ids ...int) *PatientUpdate {
	pu.mutation.AddDentalappointmentIDs(ids...)
	return pu
}

// AddDentalappointment adds the "Dentalappointment" edges to the Dentalappointment entity.
func (pu *PatientUpdate) AddDentalappointment(d ...*Dentalappointment) *PatientUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.AddDentalappointmentIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (pu *PatientUpdate) Mutation() *PatientMutation {
	return pu.mutation
}

// ClearPhysicaltherapyrecord clears all "physicaltherapyrecord" edges to the Physicaltherapyrecord entity.
func (pu *PatientUpdate) ClearPhysicaltherapyrecord() *PatientUpdate {
	pu.mutation.ClearPhysicaltherapyrecord()
	return pu
}

// RemovePhysicaltherapyrecordIDs removes the "physicaltherapyrecord" edge to Physicaltherapyrecord entities by IDs.
func (pu *PatientUpdate) RemovePhysicaltherapyrecordIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemovePhysicaltherapyrecordIDs(ids...)
	return pu
}

// RemovePhysicaltherapyrecord removes "physicaltherapyrecord" edges to Physicaltherapyrecord entities.
func (pu *PatientUpdate) RemovePhysicaltherapyrecord(p ...*Physicaltherapyrecord) *PatientUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePhysicaltherapyrecordIDs(ids...)
}

// ClearBonedisease clears all "Bonedisease" edges to the Bonedisease entity.
func (pu *PatientUpdate) ClearBonedisease() *PatientUpdate {
	pu.mutation.ClearBonedisease()
	return pu
}

// RemoveBonediseaseIDs removes the "Bonedisease" edge to Bonedisease entities by IDs.
func (pu *PatientUpdate) RemoveBonediseaseIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemoveBonediseaseIDs(ids...)
	return pu
}

// RemoveBonedisease removes "Bonedisease" edges to Bonedisease entities.
func (pu *PatientUpdate) RemoveBonedisease(b ...*Bonedisease) *PatientUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pu.RemoveBonediseaseIDs(ids...)
}

// ClearChecksymptoms clears all "Checksymptoms" edges to the Checksymptoms entity.
func (pu *PatientUpdate) ClearChecksymptoms() *PatientUpdate {
	pu.mutation.ClearChecksymptoms()
	return pu
}

// RemoveChecksymptomIDs removes the "Checksymptoms" edge to Checksymptoms entities by IDs.
func (pu *PatientUpdate) RemoveChecksymptomIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemoveChecksymptomIDs(ids...)
	return pu
}

// RemoveChecksymptoms removes "Checksymptoms" edges to Checksymptoms entities.
func (pu *PatientUpdate) RemoveChecksymptoms(c ...*Checksymptoms) *PatientUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveChecksymptomIDs(ids...)
}

// ClearDentalappointment clears all "Dentalappointment" edges to the Dentalappointment entity.
func (pu *PatientUpdate) ClearDentalappointment() *PatientUpdate {
	pu.mutation.ClearDentalappointment()
	return pu
}

// RemoveDentalappointmentIDs removes the "Dentalappointment" edge to Dentalappointment entities by IDs.
func (pu *PatientUpdate) RemoveDentalappointmentIDs(ids ...int) *PatientUpdate {
	pu.mutation.RemoveDentalappointmentIDs(ids...)
	return pu
}

// RemoveDentalappointment removes "Dentalappointment" edges to Dentalappointment entities.
func (pu *PatientUpdate) RemoveDentalappointment(d ...*Dentalappointment) *PatientUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.RemoveDentalappointmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PatientUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*PatientMutation)
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
func (pu *PatientUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PatientUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PatientUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PatientUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := patient.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Birthday(); ok {
		if err := patient.BirthdayValidator(v); err != nil {
			return &ValidationError{Name: "birthday", err: fmt.Errorf("ent: validator failed for field \"birthday\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Gender(); ok {
		if err := patient.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	return nil
}

func (pu *PatientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
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
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldName,
		})
	}
	if value, ok := pu.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldBirthday,
		})
	}
	if value, ok := pu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldGender,
		})
	}
	if pu.mutation.PhysicaltherapyrecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PhysicaltherapyrecordTable,
			Columns: []string{patient.PhysicaltherapyrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: physicaltherapyrecord.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedPhysicaltherapyrecordIDs(); len(nodes) > 0 && !pu.mutation.PhysicaltherapyrecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PhysicaltherapyrecordTable,
			Columns: []string{patient.PhysicaltherapyrecordColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PhysicaltherapyrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PhysicaltherapyrecordTable,
			Columns: []string{patient.PhysicaltherapyrecordColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.BonediseaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.BonediseaseTable,
			Columns: []string{patient.BonediseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bonedisease.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedBonediseaseIDs(); len(nodes) > 0 && !pu.mutation.BonediseaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.BonediseaseTable,
			Columns: []string{patient.BonediseaseColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.BonediseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.BonediseaseTable,
			Columns: []string{patient.BonediseaseColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ChecksymptomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.ChecksymptomsTable,
			Columns: []string{patient.ChecksymptomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checksymptoms.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedChecksymptomsIDs(); len(nodes) > 0 && !pu.mutation.ChecksymptomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.ChecksymptomsTable,
			Columns: []string{patient.ChecksymptomsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ChecksymptomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.ChecksymptomsTable,
			Columns: []string{patient.ChecksymptomsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.DentalappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.DentalappointmentTable,
			Columns: []string{patient.DentalappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalappointment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedDentalappointmentIDs(); len(nodes) > 0 && !pu.mutation.DentalappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.DentalappointmentTable,
			Columns: []string{patient.DentalappointmentColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DentalappointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.DentalappointmentTable,
			Columns: []string{patient.DentalappointmentColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PatientUpdateOne is the builder for updating a single Patient entity.
type PatientUpdateOne struct {
	config
	hooks    []Hook
	mutation *PatientMutation
}

// SetName sets the "name" field.
func (puo *PatientUpdateOne) SetName(s string) *PatientUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetBirthday sets the "birthday" field.
func (puo *PatientUpdateOne) SetBirthday(s string) *PatientUpdateOne {
	puo.mutation.SetBirthday(s)
	return puo
}

// SetGender sets the "gender" field.
func (puo *PatientUpdateOne) SetGender(s string) *PatientUpdateOne {
	puo.mutation.SetGender(s)
	return puo
}

// AddPhysicaltherapyrecordIDs adds the "physicaltherapyrecord" edge to the Physicaltherapyrecord entity by IDs.
func (puo *PatientUpdateOne) AddPhysicaltherapyrecordIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddPhysicaltherapyrecordIDs(ids...)
	return puo
}

// AddPhysicaltherapyrecord adds the "physicaltherapyrecord" edges to the Physicaltherapyrecord entity.
func (puo *PatientUpdateOne) AddPhysicaltherapyrecord(p ...*Physicaltherapyrecord) *PatientUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPhysicaltherapyrecordIDs(ids...)
}

// AddBonediseaseIDs adds the "Bonedisease" edge to the Bonedisease entity by IDs.
func (puo *PatientUpdateOne) AddBonediseaseIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddBonediseaseIDs(ids...)
	return puo
}

// AddBonedisease adds the "Bonedisease" edges to the Bonedisease entity.
func (puo *PatientUpdateOne) AddBonedisease(b ...*Bonedisease) *PatientUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.AddBonediseaseIDs(ids...)
}

// AddChecksymptomIDs adds the "Checksymptoms" edge to the Checksymptoms entity by IDs.
func (puo *PatientUpdateOne) AddChecksymptomIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddChecksymptomIDs(ids...)
	return puo
}

// AddChecksymptoms adds the "Checksymptoms" edges to the Checksymptoms entity.
func (puo *PatientUpdateOne) AddChecksymptoms(c ...*Checksymptoms) *PatientUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddChecksymptomIDs(ids...)
}

// AddDentalappointmentIDs adds the "Dentalappointment" edge to the Dentalappointment entity by IDs.
func (puo *PatientUpdateOne) AddDentalappointmentIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.AddDentalappointmentIDs(ids...)
	return puo
}

// AddDentalappointment adds the "Dentalappointment" edges to the Dentalappointment entity.
func (puo *PatientUpdateOne) AddDentalappointment(d ...*Dentalappointment) *PatientUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.AddDentalappointmentIDs(ids...)
}

// Mutation returns the PatientMutation object of the builder.
func (puo *PatientUpdateOne) Mutation() *PatientMutation {
	return puo.mutation
}

// ClearPhysicaltherapyrecord clears all "physicaltherapyrecord" edges to the Physicaltherapyrecord entity.
func (puo *PatientUpdateOne) ClearPhysicaltherapyrecord() *PatientUpdateOne {
	puo.mutation.ClearPhysicaltherapyrecord()
	return puo
}

// RemovePhysicaltherapyrecordIDs removes the "physicaltherapyrecord" edge to Physicaltherapyrecord entities by IDs.
func (puo *PatientUpdateOne) RemovePhysicaltherapyrecordIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemovePhysicaltherapyrecordIDs(ids...)
	return puo
}

// RemovePhysicaltherapyrecord removes "physicaltherapyrecord" edges to Physicaltherapyrecord entities.
func (puo *PatientUpdateOne) RemovePhysicaltherapyrecord(p ...*Physicaltherapyrecord) *PatientUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePhysicaltherapyrecordIDs(ids...)
}

// ClearBonedisease clears all "Bonedisease" edges to the Bonedisease entity.
func (puo *PatientUpdateOne) ClearBonedisease() *PatientUpdateOne {
	puo.mutation.ClearBonedisease()
	return puo
}

// RemoveBonediseaseIDs removes the "Bonedisease" edge to Bonedisease entities by IDs.
func (puo *PatientUpdateOne) RemoveBonediseaseIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemoveBonediseaseIDs(ids...)
	return puo
}

// RemoveBonedisease removes "Bonedisease" edges to Bonedisease entities.
func (puo *PatientUpdateOne) RemoveBonedisease(b ...*Bonedisease) *PatientUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return puo.RemoveBonediseaseIDs(ids...)
}

// ClearChecksymptoms clears all "Checksymptoms" edges to the Checksymptoms entity.
func (puo *PatientUpdateOne) ClearChecksymptoms() *PatientUpdateOne {
	puo.mutation.ClearChecksymptoms()
	return puo
}

// RemoveChecksymptomIDs removes the "Checksymptoms" edge to Checksymptoms entities by IDs.
func (puo *PatientUpdateOne) RemoveChecksymptomIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemoveChecksymptomIDs(ids...)
	return puo
}

// RemoveChecksymptoms removes "Checksymptoms" edges to Checksymptoms entities.
func (puo *PatientUpdateOne) RemoveChecksymptoms(c ...*Checksymptoms) *PatientUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveChecksymptomIDs(ids...)
}

// ClearDentalappointment clears all "Dentalappointment" edges to the Dentalappointment entity.
func (puo *PatientUpdateOne) ClearDentalappointment() *PatientUpdateOne {
	puo.mutation.ClearDentalappointment()
	return puo
}

// RemoveDentalappointmentIDs removes the "Dentalappointment" edge to Dentalappointment entities by IDs.
func (puo *PatientUpdateOne) RemoveDentalappointmentIDs(ids ...int) *PatientUpdateOne {
	puo.mutation.RemoveDentalappointmentIDs(ids...)
	return puo
}

// RemoveDentalappointment removes "Dentalappointment" edges to Dentalappointment entities.
func (puo *PatientUpdateOne) RemoveDentalappointment(d ...*Dentalappointment) *PatientUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.RemoveDentalappointmentIDs(ids...)
}

// Save executes the query and returns the updated Patient entity.
func (puo *PatientUpdateOne) Save(ctx context.Context) (*Patient, error) {
	var (
		err  error
		node *Patient
	)
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientMutation)
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
func (puo *PatientUpdateOne) SaveX(ctx context.Context) *Patient {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PatientUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PatientUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PatientUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := patient.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Birthday(); ok {
		if err := patient.BirthdayValidator(v); err != nil {
			return &ValidationError{Name: "birthday", err: fmt.Errorf("ent: validator failed for field \"birthday\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Gender(); ok {
		if err := patient.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf("ent: validator failed for field \"gender\": %w", err)}
		}
	}
	return nil
}

func (puo *PatientUpdateOne) sqlSave(ctx context.Context) (_node *Patient, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patient.Table,
			Columns: patient.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patient.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Patient.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldName,
		})
	}
	if value, ok := puo.mutation.Birthday(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldBirthday,
		})
	}
	if value, ok := puo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patient.FieldGender,
		})
	}
	if puo.mutation.PhysicaltherapyrecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PhysicaltherapyrecordTable,
			Columns: []string{patient.PhysicaltherapyrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: physicaltherapyrecord.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedPhysicaltherapyrecordIDs(); len(nodes) > 0 && !puo.mutation.PhysicaltherapyrecordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PhysicaltherapyrecordTable,
			Columns: []string{patient.PhysicaltherapyrecordColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PhysicaltherapyrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.PhysicaltherapyrecordTable,
			Columns: []string{patient.PhysicaltherapyrecordColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.BonediseaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.BonediseaseTable,
			Columns: []string{patient.BonediseaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bonedisease.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedBonediseaseIDs(); len(nodes) > 0 && !puo.mutation.BonediseaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.BonediseaseTable,
			Columns: []string{patient.BonediseaseColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.BonediseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.BonediseaseTable,
			Columns: []string{patient.BonediseaseColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ChecksymptomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.ChecksymptomsTable,
			Columns: []string{patient.ChecksymptomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checksymptoms.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedChecksymptomsIDs(); len(nodes) > 0 && !puo.mutation.ChecksymptomsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.ChecksymptomsTable,
			Columns: []string{patient.ChecksymptomsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ChecksymptomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.ChecksymptomsTable,
			Columns: []string{patient.ChecksymptomsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.DentalappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.DentalappointmentTable,
			Columns: []string{patient.DentalappointmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dentalappointment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedDentalappointmentIDs(); len(nodes) > 0 && !puo.mutation.DentalappointmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.DentalappointmentTable,
			Columns: []string{patient.DentalappointmentColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DentalappointmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patient.DentalappointmentTable,
			Columns: []string{patient.DentalappointmentColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Patient{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patient.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
