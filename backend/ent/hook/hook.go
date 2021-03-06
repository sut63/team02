// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/to63/app/ent"
)

// The AntenatalinformationFunc type is an adapter to allow the use of ordinary
// function as Antenatalinformation mutator.
type AntenatalinformationFunc func(context.Context, *ent.AntenatalinformationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AntenatalinformationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AntenatalinformationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AntenatalinformationMutation", m)
	}
	return f(ctx, mv)
}

// The BonediseaseFunc type is an adapter to allow the use of ordinary
// function as Bonedisease mutator.
type BonediseaseFunc func(context.Context, *ent.BonediseaseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BonediseaseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BonediseaseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BonediseaseMutation", m)
	}
	return f(ctx, mv)
}

// The ChecksymptomFunc type is an adapter to allow the use of ordinary
// function as Checksymptom mutator.
type ChecksymptomFunc func(context.Context, *ent.ChecksymptomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ChecksymptomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ChecksymptomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ChecksymptomMutation", m)
	}
	return f(ctx, mv)
}

// The DentalappointmentFunc type is an adapter to allow the use of ordinary
// function as Dentalappointment mutator.
type DentalappointmentFunc func(context.Context, *ent.DentalappointmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DentalappointmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DentalappointmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DentalappointmentMutation", m)
	}
	return f(ctx, mv)
}

// The DentalkindFunc type is an adapter to allow the use of ordinary
// function as Dentalkind mutator.
type DentalkindFunc func(context.Context, *ent.DentalkindMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DentalkindFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DentalkindMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DentalkindMutation", m)
	}
	return f(ctx, mv)
}

// The DepartmentFunc type is an adapter to allow the use of ordinary
// function as Department mutator.
type DepartmentFunc func(context.Context, *ent.DepartmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DepartmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DepartmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DepartmentMutation", m)
	}
	return f(ctx, mv)
}

// The DiseaseFunc type is an adapter to allow the use of ordinary
// function as Disease mutator.
type DiseaseFunc func(context.Context, *ent.DiseaseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DiseaseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DiseaseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DiseaseMutation", m)
	}
	return f(ctx, mv)
}

// The DoctorordersheetFunc type is an adapter to allow the use of ordinary
// function as Doctorordersheet mutator.
type DoctorordersheetFunc func(context.Context, *ent.DoctorordersheetMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DoctorordersheetFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DoctorordersheetMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DoctorordersheetMutation", m)
	}
	return f(ctx, mv)
}

// The GenderFunc type is an adapter to allow the use of ordinary
// function as Gender mutator.
type GenderFunc func(context.Context, *ent.GenderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GenderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GenderMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GenderMutation", m)
	}
	return f(ctx, mv)
}

// The PatientFunc type is an adapter to allow the use of ordinary
// function as Patient mutator.
type PatientFunc func(context.Context, *ent.PatientMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PatientFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PatientMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PatientMutation", m)
	}
	return f(ctx, mv)
}

// The PersonnelFunc type is an adapter to allow the use of ordinary
// function as Personnel mutator.
type PersonnelFunc func(context.Context, *ent.PersonnelMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PersonnelFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PersonnelMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PersonnelMutation", m)
	}
	return f(ctx, mv)
}

// The PhysicaltherapyrecordFunc type is an adapter to allow the use of ordinary
// function as Physicaltherapyrecord mutator.
type PhysicaltherapyrecordFunc func(context.Context, *ent.PhysicaltherapyrecordMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PhysicaltherapyrecordFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PhysicaltherapyrecordMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PhysicaltherapyrecordMutation", m)
	}
	return f(ctx, mv)
}

// The PhysicaltherapyroomFunc type is an adapter to allow the use of ordinary
// function as Physicaltherapyroom mutator.
type PhysicaltherapyroomFunc func(context.Context, *ent.PhysicaltherapyroomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PhysicaltherapyroomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PhysicaltherapyroomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PhysicaltherapyroomMutation", m)
	}
	return f(ctx, mv)
}

// The PregnancystatusFunc type is an adapter to allow the use of ordinary
// function as Pregnancystatus mutator.
type PregnancystatusFunc func(context.Context, *ent.PregnancystatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PregnancystatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PregnancystatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PregnancystatusMutation", m)
	}
	return f(ctx, mv)
}

// The RemedyFunc type is an adapter to allow the use of ordinary
// function as Remedy mutator.
type RemedyFunc func(context.Context, *ent.RemedyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RemedyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RemedyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RemedyMutation", m)
	}
	return f(ctx, mv)
}

// The RisksFunc type is an adapter to allow the use of ordinary
// function as Risks mutator.
type RisksFunc func(context.Context, *ent.RisksMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RisksFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RisksMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RisksMutation", m)
	}
	return f(ctx, mv)
}

// The StatusFunc type is an adapter to allow the use of ordinary
// function as Status mutator.
type StatusFunc func(context.Context, *ent.StatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StatusMutation", m)
	}
	return f(ctx, mv)
}

// The SurgeryappointmentFunc type is an adapter to allow the use of ordinary
// function as Surgeryappointment mutator.
type SurgeryappointmentFunc func(context.Context, *ent.SurgeryappointmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SurgeryappointmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SurgeryappointmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SurgeryappointmentMutation", m)
	}
	return f(ctx, mv)
}

// The SurgerytypeFunc type is an adapter to allow the use of ordinary
// function as Surgerytype mutator.
type SurgerytypeFunc func(context.Context, *ent.SurgerytypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SurgerytypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SurgerytypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SurgerytypeMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
