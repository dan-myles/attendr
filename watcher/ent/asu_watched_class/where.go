// Code generated by ent, DO NOT EDIT.

package asu_watched_class

import (
	"attendr/watcher/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldTitle, v))
}

// Instructor applies equality check predicate on the "instructor" field. It's identical to InstructorEQ.
func Instructor(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldInstructor, v))
}

// Subject applies equality check predicate on the "subject" field. It's identical to SubjectEQ.
func Subject(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldSubject, v))
}

// SubjectNumber applies equality check predicate on the "subject_number" field. It's identical to SubjectNumberEQ.
func SubjectNumber(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldSubjectNumber, v))
}

// HasOpenSeats applies equality check predicate on the "has_open_seats" field. It's identical to HasOpenSeatsEQ.
func HasOpenSeats(v bool) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldHasOpenSeats, v))
}

// TrackedAt applies equality check predicate on the "tracked_at" field. It's identical to TrackedAtEQ.
func TrackedAt(v time.Time) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldTrackedAt, v))
}

// ClassNumber applies equality check predicate on the "class_number" field. It's identical to ClassNumberEQ.
func ClassNumber(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldClassNumber, v))
}

// Term applies equality check predicate on the "term" field. It's identical to TermEQ.
func Term(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldTerm, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContainsFold(FieldTitle, v))
}

// InstructorEQ applies the EQ predicate on the "instructor" field.
func InstructorEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldInstructor, v))
}

// InstructorNEQ applies the NEQ predicate on the "instructor" field.
func InstructorNEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNEQ(FieldInstructor, v))
}

// InstructorIn applies the In predicate on the "instructor" field.
func InstructorIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldIn(FieldInstructor, vs...))
}

// InstructorNotIn applies the NotIn predicate on the "instructor" field.
func InstructorNotIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNotIn(FieldInstructor, vs...))
}

// InstructorGT applies the GT predicate on the "instructor" field.
func InstructorGT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGT(FieldInstructor, v))
}

// InstructorGTE applies the GTE predicate on the "instructor" field.
func InstructorGTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGTE(FieldInstructor, v))
}

// InstructorLT applies the LT predicate on the "instructor" field.
func InstructorLT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLT(FieldInstructor, v))
}

// InstructorLTE applies the LTE predicate on the "instructor" field.
func InstructorLTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLTE(FieldInstructor, v))
}

// InstructorContains applies the Contains predicate on the "instructor" field.
func InstructorContains(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContains(FieldInstructor, v))
}

// InstructorHasPrefix applies the HasPrefix predicate on the "instructor" field.
func InstructorHasPrefix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasPrefix(FieldInstructor, v))
}

// InstructorHasSuffix applies the HasSuffix predicate on the "instructor" field.
func InstructorHasSuffix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasSuffix(FieldInstructor, v))
}

// InstructorEqualFold applies the EqualFold predicate on the "instructor" field.
func InstructorEqualFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEqualFold(FieldInstructor, v))
}

// InstructorContainsFold applies the ContainsFold predicate on the "instructor" field.
func InstructorContainsFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContainsFold(FieldInstructor, v))
}

// SubjectEQ applies the EQ predicate on the "subject" field.
func SubjectEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldSubject, v))
}

// SubjectNEQ applies the NEQ predicate on the "subject" field.
func SubjectNEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNEQ(FieldSubject, v))
}

// SubjectIn applies the In predicate on the "subject" field.
func SubjectIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldIn(FieldSubject, vs...))
}

// SubjectNotIn applies the NotIn predicate on the "subject" field.
func SubjectNotIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNotIn(FieldSubject, vs...))
}

// SubjectGT applies the GT predicate on the "subject" field.
func SubjectGT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGT(FieldSubject, v))
}

// SubjectGTE applies the GTE predicate on the "subject" field.
func SubjectGTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGTE(FieldSubject, v))
}

// SubjectLT applies the LT predicate on the "subject" field.
func SubjectLT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLT(FieldSubject, v))
}

// SubjectLTE applies the LTE predicate on the "subject" field.
func SubjectLTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLTE(FieldSubject, v))
}

// SubjectContains applies the Contains predicate on the "subject" field.
func SubjectContains(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContains(FieldSubject, v))
}

// SubjectHasPrefix applies the HasPrefix predicate on the "subject" field.
func SubjectHasPrefix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasPrefix(FieldSubject, v))
}

// SubjectHasSuffix applies the HasSuffix predicate on the "subject" field.
func SubjectHasSuffix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasSuffix(FieldSubject, v))
}

// SubjectEqualFold applies the EqualFold predicate on the "subject" field.
func SubjectEqualFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEqualFold(FieldSubject, v))
}

// SubjectContainsFold applies the ContainsFold predicate on the "subject" field.
func SubjectContainsFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContainsFold(FieldSubject, v))
}

// SubjectNumberEQ applies the EQ predicate on the "subject_number" field.
func SubjectNumberEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldSubjectNumber, v))
}

// SubjectNumberNEQ applies the NEQ predicate on the "subject_number" field.
func SubjectNumberNEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNEQ(FieldSubjectNumber, v))
}

// SubjectNumberIn applies the In predicate on the "subject_number" field.
func SubjectNumberIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldIn(FieldSubjectNumber, vs...))
}

// SubjectNumberNotIn applies the NotIn predicate on the "subject_number" field.
func SubjectNumberNotIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNotIn(FieldSubjectNumber, vs...))
}

// SubjectNumberGT applies the GT predicate on the "subject_number" field.
func SubjectNumberGT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGT(FieldSubjectNumber, v))
}

// SubjectNumberGTE applies the GTE predicate on the "subject_number" field.
func SubjectNumberGTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGTE(FieldSubjectNumber, v))
}

// SubjectNumberLT applies the LT predicate on the "subject_number" field.
func SubjectNumberLT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLT(FieldSubjectNumber, v))
}

// SubjectNumberLTE applies the LTE predicate on the "subject_number" field.
func SubjectNumberLTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLTE(FieldSubjectNumber, v))
}

// SubjectNumberContains applies the Contains predicate on the "subject_number" field.
func SubjectNumberContains(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContains(FieldSubjectNumber, v))
}

// SubjectNumberHasPrefix applies the HasPrefix predicate on the "subject_number" field.
func SubjectNumberHasPrefix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasPrefix(FieldSubjectNumber, v))
}

// SubjectNumberHasSuffix applies the HasSuffix predicate on the "subject_number" field.
func SubjectNumberHasSuffix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasSuffix(FieldSubjectNumber, v))
}

// SubjectNumberEqualFold applies the EqualFold predicate on the "subject_number" field.
func SubjectNumberEqualFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEqualFold(FieldSubjectNumber, v))
}

// SubjectNumberContainsFold applies the ContainsFold predicate on the "subject_number" field.
func SubjectNumberContainsFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContainsFold(FieldSubjectNumber, v))
}

// HasOpenSeatsEQ applies the EQ predicate on the "has_open_seats" field.
func HasOpenSeatsEQ(v bool) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldHasOpenSeats, v))
}

// HasOpenSeatsNEQ applies the NEQ predicate on the "has_open_seats" field.
func HasOpenSeatsNEQ(v bool) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNEQ(FieldHasOpenSeats, v))
}

// TrackedAtEQ applies the EQ predicate on the "tracked_at" field.
func TrackedAtEQ(v time.Time) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldTrackedAt, v))
}

// TrackedAtNEQ applies the NEQ predicate on the "tracked_at" field.
func TrackedAtNEQ(v time.Time) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNEQ(FieldTrackedAt, v))
}

// TrackedAtIn applies the In predicate on the "tracked_at" field.
func TrackedAtIn(vs ...time.Time) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldIn(FieldTrackedAt, vs...))
}

// TrackedAtNotIn applies the NotIn predicate on the "tracked_at" field.
func TrackedAtNotIn(vs ...time.Time) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNotIn(FieldTrackedAt, vs...))
}

// TrackedAtGT applies the GT predicate on the "tracked_at" field.
func TrackedAtGT(v time.Time) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGT(FieldTrackedAt, v))
}

// TrackedAtGTE applies the GTE predicate on the "tracked_at" field.
func TrackedAtGTE(v time.Time) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGTE(FieldTrackedAt, v))
}

// TrackedAtLT applies the LT predicate on the "tracked_at" field.
func TrackedAtLT(v time.Time) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLT(FieldTrackedAt, v))
}

// TrackedAtLTE applies the LTE predicate on the "tracked_at" field.
func TrackedAtLTE(v time.Time) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLTE(FieldTrackedAt, v))
}

// ClassNumberEQ applies the EQ predicate on the "class_number" field.
func ClassNumberEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldClassNumber, v))
}

// ClassNumberNEQ applies the NEQ predicate on the "class_number" field.
func ClassNumberNEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNEQ(FieldClassNumber, v))
}

// ClassNumberIn applies the In predicate on the "class_number" field.
func ClassNumberIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldIn(FieldClassNumber, vs...))
}

// ClassNumberNotIn applies the NotIn predicate on the "class_number" field.
func ClassNumberNotIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNotIn(FieldClassNumber, vs...))
}

// ClassNumberGT applies the GT predicate on the "class_number" field.
func ClassNumberGT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGT(FieldClassNumber, v))
}

// ClassNumberGTE applies the GTE predicate on the "class_number" field.
func ClassNumberGTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGTE(FieldClassNumber, v))
}

// ClassNumberLT applies the LT predicate on the "class_number" field.
func ClassNumberLT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLT(FieldClassNumber, v))
}

// ClassNumberLTE applies the LTE predicate on the "class_number" field.
func ClassNumberLTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLTE(FieldClassNumber, v))
}

// ClassNumberContains applies the Contains predicate on the "class_number" field.
func ClassNumberContains(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContains(FieldClassNumber, v))
}

// ClassNumberHasPrefix applies the HasPrefix predicate on the "class_number" field.
func ClassNumberHasPrefix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasPrefix(FieldClassNumber, v))
}

// ClassNumberHasSuffix applies the HasSuffix predicate on the "class_number" field.
func ClassNumberHasSuffix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasSuffix(FieldClassNumber, v))
}

// ClassNumberEqualFold applies the EqualFold predicate on the "class_number" field.
func ClassNumberEqualFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEqualFold(FieldClassNumber, v))
}

// ClassNumberContainsFold applies the ContainsFold predicate on the "class_number" field.
func ClassNumberContainsFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContainsFold(FieldClassNumber, v))
}

// TermEQ applies the EQ predicate on the "term" field.
func TermEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEQ(FieldTerm, v))
}

// TermNEQ applies the NEQ predicate on the "term" field.
func TermNEQ(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNEQ(FieldTerm, v))
}

// TermIn applies the In predicate on the "term" field.
func TermIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldIn(FieldTerm, vs...))
}

// TermNotIn applies the NotIn predicate on the "term" field.
func TermNotIn(vs ...string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldNotIn(FieldTerm, vs...))
}

// TermGT applies the GT predicate on the "term" field.
func TermGT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGT(FieldTerm, v))
}

// TermGTE applies the GTE predicate on the "term" field.
func TermGTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldGTE(FieldTerm, v))
}

// TermLT applies the LT predicate on the "term" field.
func TermLT(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLT(FieldTerm, v))
}

// TermLTE applies the LTE predicate on the "term" field.
func TermLTE(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldLTE(FieldTerm, v))
}

// TermContains applies the Contains predicate on the "term" field.
func TermContains(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContains(FieldTerm, v))
}

// TermHasPrefix applies the HasPrefix predicate on the "term" field.
func TermHasPrefix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasPrefix(FieldTerm, v))
}

// TermHasSuffix applies the HasSuffix predicate on the "term" field.
func TermHasSuffix(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldHasSuffix(FieldTerm, v))
}

// TermEqualFold applies the EqualFold predicate on the "term" field.
func TermEqualFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldEqualFold(FieldTerm, v))
}

// TermContainsFold applies the ContainsFold predicate on the "term" field.
func TermContainsFold(v string) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(sql.FieldContainsFold(FieldTerm, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ASU_Watched_Class) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ASU_Watched_Class) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ASU_Watched_Class) predicate.ASU_Watched_Class {
	return predicate.ASU_Watched_Class(func(s *sql.Selector) {
		p(s.Not())
	})
}
