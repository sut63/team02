// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/to63/app/ent/patient"
	"github.com/to63/app/ent/personnel"
	"github.com/to63/app/ent/predicate"
	"github.com/to63/app/ent/surgeryappointment"
	"github.com/to63/app/ent/surgerytype"
)

// SurgeryappointmentUpdate is the builder for updating Surgeryappointment entities.
type SurgeryappointmentUpdate struct {
	config
	hooks    []Hook
	mutation *SurgeryappointmentMutation
}

// Where adds a new predicate for the SurgeryappointmentUpdate builder.
func (su *SurgeryappointmentUpdate) Where(ps ...predicate.Surgeryappointment) *SurgeryappointmentUpdate {
	su.mutation.predicates = append(su.mutation.predicates, ps...)
	return su
}

// SetAppointTime sets the "appoint_time" field.
func (su *SurgeryappointmentUpdate) SetAppointTime(t time.Time) *SurgeryappointmentUpdate {
	su.mutation.SetAppointTime(t)
	return su
}

// SetPhone sets the "phone" field.
func (su *SurgeryappointmentUpdate) SetPhone(s string) *SurgeryappointmentUpdate {
	su.mutation.SetPhone(s)
	return su
}

// SetNote sets the "note" field.
func (su *SurgeryappointmentUpdate) SetNote(s string) *SurgeryappointmentUpdate {
	su.mutation.SetNote(s)
	return su
}

// SetAge sets the "age" field.
func (su *SurgeryappointmentUpdate) SetAge(i int) *SurgeryappointmentUpdate {
	su.mutation.ResetAge()
	su.mutation.SetAge(i)
	return su
}

// AddAge adds i to the "age" field.
func (su *SurgeryappointmentUpdate) AddAge(i int) *SurgeryappointmentUpdate {
	su.mutation.AddAge(i)
	return su
}

// SetPersonnelID sets the "Personnel" edge to the Personnel entity by ID.
func (su *SurgeryappointmentUpdate) SetPersonnelID(id int) *SurgeryappointmentUpdate {
	su.mutation.SetPersonnelID(id)
	return su
}

// SetNillablePersonnelID sets the "Personnel" edge to the Personnel entity by ID if the given value is not nil.
func (su *SurgeryappointmentUpdate) SetNillablePersonnelID(id *int) *SurgeryappointmentUpdate {
	if id != nil {
		su = su.SetPersonnelID(*id)
	}
	return su
}

// SetPersonnel sets the "Personnel" edge to the Personnel entity.
func (su *SurgeryappointmentUpdate) SetPersonnel(p *Personnel) *SurgeryappointmentUpdate {
	return su.SetPersonnelID(p.ID)
}

// SetPatientID sets the "Patient" edge to the Patient entity by ID.
func (su *SurgeryappointmentUpdate) SetPatientID(id int) *SurgeryappointmentUpdate {
	su.mutation.SetPatientID(id)
	return su
}

// SetNillablePatientID sets the "Patient" edge to the Patient entity by ID if the given value is not nil.
func (su *SurgeryappointmentUpdate) SetNillablePatientID(id *int) *SurgeryappointmentUpdate {
	if id != nil {
		su = su.SetPatientID(*id)
	}
	return su
}

// SetPatient sets the "Patient" edge to the Patient entity.
func (su *SurgeryappointmentUpdate) SetPatient(p *Patient) *SurgeryappointmentUpdate {
	return su.SetPatientID(p.ID)
}

// SetSurgerytypeID sets the "Surgerytype" edge to the Surgerytype entity by ID.
func (su *SurgeryappointmentUpdate) SetSurgerytypeID(id int) *SurgeryappointmentUpdate {
	su.mutation.SetSurgerytypeID(id)
	return su
}

// SetNillableSurgerytypeID sets the "Surgerytype" edge to the Surgerytype entity by ID if the given value is not nil.
func (su *SurgeryappointmentUpdate) SetNillableSurgerytypeID(id *int) *SurgeryappointmentUpdate {
	if id != nil {
		su = su.SetSurgerytypeID(*id)
	}
	return su
}

// SetSurgerytype sets the "Surgerytype" edge to the Surgerytype entity.
func (su *SurgeryappointmentUpdate) SetSurgerytype(s *Surgerytype) *SurgeryappointmentUpdate {
	return su.SetSurgerytypeID(s.ID)
}

// Mutation returns the SurgeryappointmentMutation object of the builder.
func (su *SurgeryappointmentUpdate) Mutation() *SurgeryappointmentMutation {
	return su.mutation
}

// ClearPersonnel clears the "Personnel" edge to the Personnel entity.
func (su *SurgeryappointmentUpdate) ClearPersonnel() *SurgeryappointmentUpdate {
	su.mutation.ClearPersonnel()
	return su
}

// ClearPatient clears the "Patient" edge to the Patient entity.
func (su *SurgeryappointmentUpdate) ClearPatient() *SurgeryappointmentUpdate {
	su.mutation.ClearPatient()
	return su
}

// ClearSurgerytype clears the "Surgerytype" edge to the Surgerytype entity.
func (su *SurgeryappointmentUpdate) ClearSurgerytype() *SurgeryappointmentUpdate {
	su.mutation.ClearSurgerytype()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SurgeryappointmentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SurgeryappointmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SurgeryappointmentUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SurgeryappointmentUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SurgeryappointmentUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SurgeryappointmentUpdate) check() error {
	if v, ok := su.mutation.Phone(); ok {
		if err := surgeryappointment.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf("ent: validator failed for field \"phone\": %w", err)}
		}
	}
	if v, ok := su.mutation.Note(); ok {
		if err := surgeryappointment.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	if v, ok := su.mutation.Age(); ok {
		if err := surgeryappointment.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	return nil
}

func (su *SurgeryappointmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   surgeryappointment.Table,
			Columns: surgeryappointment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: surgeryappointment.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.AppointTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: surgeryappointment.FieldAppointTime,
		})
	}
	if value, ok := su.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: surgeryappointment.FieldPhone,
		})
	}
	if value, ok := su.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: surgeryappointment.FieldNote,
		})
	}
	if value, ok := su.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: surgeryappointment.FieldAge,
		})
	}
	if value, ok := su.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: surgeryappointment.FieldAge,
		})
	}
	if su.mutation.PersonnelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.PersonnelTable,
			Columns: []string{surgeryappointment.PersonnelColumn},
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
	if nodes := su.mutation.PersonnelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.PersonnelTable,
			Columns: []string{surgeryappointment.PersonnelColumn},
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
	if su.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.PatientTable,
			Columns: []string{surgeryappointment.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.PatientTable,
			Columns: []string{surgeryappointment.PatientColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.SurgerytypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.SurgerytypeTable,
			Columns: []string{surgeryappointment.SurgerytypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surgerytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SurgerytypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.SurgerytypeTable,
			Columns: []string{surgeryappointment.SurgerytypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surgerytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surgeryappointment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SurgeryappointmentUpdateOne is the builder for updating a single Surgeryappointment entity.
type SurgeryappointmentUpdateOne struct {
	config
	hooks    []Hook
	mutation *SurgeryappointmentMutation
}

// SetAppointTime sets the "appoint_time" field.
func (suo *SurgeryappointmentUpdateOne) SetAppointTime(t time.Time) *SurgeryappointmentUpdateOne {
	suo.mutation.SetAppointTime(t)
	return suo
}

// SetPhone sets the "phone" field.
func (suo *SurgeryappointmentUpdateOne) SetPhone(s string) *SurgeryappointmentUpdateOne {
	suo.mutation.SetPhone(s)
	return suo
}

// SetNote sets the "note" field.
func (suo *SurgeryappointmentUpdateOne) SetNote(s string) *SurgeryappointmentUpdateOne {
	suo.mutation.SetNote(s)
	return suo
}

// SetAge sets the "age" field.
func (suo *SurgeryappointmentUpdateOne) SetAge(i int) *SurgeryappointmentUpdateOne {
	suo.mutation.ResetAge()
	suo.mutation.SetAge(i)
	return suo
}

// AddAge adds i to the "age" field.
func (suo *SurgeryappointmentUpdateOne) AddAge(i int) *SurgeryappointmentUpdateOne {
	suo.mutation.AddAge(i)
	return suo
}

// SetPersonnelID sets the "Personnel" edge to the Personnel entity by ID.
func (suo *SurgeryappointmentUpdateOne) SetPersonnelID(id int) *SurgeryappointmentUpdateOne {
	suo.mutation.SetPersonnelID(id)
	return suo
}

// SetNillablePersonnelID sets the "Personnel" edge to the Personnel entity by ID if the given value is not nil.
func (suo *SurgeryappointmentUpdateOne) SetNillablePersonnelID(id *int) *SurgeryappointmentUpdateOne {
	if id != nil {
		suo = suo.SetPersonnelID(*id)
	}
	return suo
}

// SetPersonnel sets the "Personnel" edge to the Personnel entity.
func (suo *SurgeryappointmentUpdateOne) SetPersonnel(p *Personnel) *SurgeryappointmentUpdateOne {
	return suo.SetPersonnelID(p.ID)
}

// SetPatientID sets the "Patient" edge to the Patient entity by ID.
func (suo *SurgeryappointmentUpdateOne) SetPatientID(id int) *SurgeryappointmentUpdateOne {
	suo.mutation.SetPatientID(id)
	return suo
}

// SetNillablePatientID sets the "Patient" edge to the Patient entity by ID if the given value is not nil.
func (suo *SurgeryappointmentUpdateOne) SetNillablePatientID(id *int) *SurgeryappointmentUpdateOne {
	if id != nil {
		suo = suo.SetPatientID(*id)
	}
	return suo
}

// SetPatient sets the "Patient" edge to the Patient entity.
func (suo *SurgeryappointmentUpdateOne) SetPatient(p *Patient) *SurgeryappointmentUpdateOne {
	return suo.SetPatientID(p.ID)
}

// SetSurgerytypeID sets the "Surgerytype" edge to the Surgerytype entity by ID.
func (suo *SurgeryappointmentUpdateOne) SetSurgerytypeID(id int) *SurgeryappointmentUpdateOne {
	suo.mutation.SetSurgerytypeID(id)
	return suo
}

// SetNillableSurgerytypeID sets the "Surgerytype" edge to the Surgerytype entity by ID if the given value is not nil.
func (suo *SurgeryappointmentUpdateOne) SetNillableSurgerytypeID(id *int) *SurgeryappointmentUpdateOne {
	if id != nil {
		suo = suo.SetSurgerytypeID(*id)
	}
	return suo
}

// SetSurgerytype sets the "Surgerytype" edge to the Surgerytype entity.
func (suo *SurgeryappointmentUpdateOne) SetSurgerytype(s *Surgerytype) *SurgeryappointmentUpdateOne {
	return suo.SetSurgerytypeID(s.ID)
}

// Mutation returns the SurgeryappointmentMutation object of the builder.
func (suo *SurgeryappointmentUpdateOne) Mutation() *SurgeryappointmentMutation {
	return suo.mutation
}

// ClearPersonnel clears the "Personnel" edge to the Personnel entity.
func (suo *SurgeryappointmentUpdateOne) ClearPersonnel() *SurgeryappointmentUpdateOne {
	suo.mutation.ClearPersonnel()
	return suo
}

// ClearPatient clears the "Patient" edge to the Patient entity.
func (suo *SurgeryappointmentUpdateOne) ClearPatient() *SurgeryappointmentUpdateOne {
	suo.mutation.ClearPatient()
	return suo
}

// ClearSurgerytype clears the "Surgerytype" edge to the Surgerytype entity.
func (suo *SurgeryappointmentUpdateOne) ClearSurgerytype() *SurgeryappointmentUpdateOne {
	suo.mutation.ClearSurgerytype()
	return suo
}

// Save executes the query and returns the updated Surgeryappointment entity.
func (suo *SurgeryappointmentUpdateOne) Save(ctx context.Context) (*Surgeryappointment, error) {
	var (
		err  error
		node *Surgeryappointment
	)
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SurgeryappointmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SurgeryappointmentUpdateOne) SaveX(ctx context.Context) *Surgeryappointment {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SurgeryappointmentUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SurgeryappointmentUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SurgeryappointmentUpdateOne) check() error {
	if v, ok := suo.mutation.Phone(); ok {
		if err := surgeryappointment.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf("ent: validator failed for field \"phone\": %w", err)}
		}
	}
	if v, ok := suo.mutation.Note(); ok {
		if err := surgeryappointment.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	if v, ok := suo.mutation.Age(); ok {
		if err := surgeryappointment.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	return nil
}

func (suo *SurgeryappointmentUpdateOne) sqlSave(ctx context.Context) (_node *Surgeryappointment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   surgeryappointment.Table,
			Columns: surgeryappointment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: surgeryappointment.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Surgeryappointment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.AppointTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: surgeryappointment.FieldAppointTime,
		})
	}
	if value, ok := suo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: surgeryappointment.FieldPhone,
		})
	}
	if value, ok := suo.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: surgeryappointment.FieldNote,
		})
	}
	if value, ok := suo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: surgeryappointment.FieldAge,
		})
	}
	if value, ok := suo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: surgeryappointment.FieldAge,
		})
	}
	if suo.mutation.PersonnelCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.PersonnelTable,
			Columns: []string{surgeryappointment.PersonnelColumn},
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
	if nodes := suo.mutation.PersonnelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.PersonnelTable,
			Columns: []string{surgeryappointment.PersonnelColumn},
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
	if suo.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.PatientTable,
			Columns: []string{surgeryappointment.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.PatientTable,
			Columns: []string{surgeryappointment.PatientColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.SurgerytypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.SurgerytypeTable,
			Columns: []string{surgeryappointment.SurgerytypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surgerytype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SurgerytypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   surgeryappointment.SurgerytypeTable,
			Columns: []string{surgeryappointment.SurgerytypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: surgerytype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Surgeryappointment{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{surgeryappointment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
