// Code generated by ent, DO NOT EDIT.

package asu_watched_class

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the asu_watched_class type in the database.
	Label = "asu_watched_class"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldInstructor holds the string denoting the instructor field in the database.
	FieldInstructor = "instructor"
	// FieldSubject holds the string denoting the subject field in the database.
	FieldSubject = "subject"
	// FieldSubjectNumber holds the string denoting the subject_number field in the database.
	FieldSubjectNumber = "subject_number"
	// FieldHasOpenSeats holds the string denoting the has_open_seats field in the database.
	FieldHasOpenSeats = "has_open_seats"
	// FieldTrackedAt holds the string denoting the tracked_at field in the database.
	FieldTrackedAt = "tracked_at"
	// FieldClassNumber holds the string denoting the class_number field in the database.
	FieldClassNumber = "class_number"
	// FieldTerm holds the string denoting the term field in the database.
	FieldTerm = "term"
	// Table holds the table name of the asu_watched_class in the database.
	Table = "asu_watched_classes"
)

// Columns holds all SQL columns for asu_watched_class fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldInstructor,
	FieldSubject,
	FieldSubjectNumber,
	FieldHasOpenSeats,
	FieldTrackedAt,
	FieldClassNumber,
	FieldTerm,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultHasOpenSeats holds the default value on creation for the "has_open_seats" field.
	DefaultHasOpenSeats bool
	// DefaultTrackedAt holds the default value on creation for the "tracked_at" field.
	DefaultTrackedAt func() time.Time
)

// OrderOption defines the ordering options for the ASU_Watched_Class queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByInstructor orders the results by the instructor field.
func ByInstructor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInstructor, opts...).ToFunc()
}

// BySubject orders the results by the subject field.
func BySubject(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubject, opts...).ToFunc()
}

// BySubjectNumber orders the results by the subject_number field.
func BySubjectNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubjectNumber, opts...).ToFunc()
}

// ByHasOpenSeats orders the results by the has_open_seats field.
func ByHasOpenSeats(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasOpenSeats, opts...).ToFunc()
}

// ByTrackedAt orders the results by the tracked_at field.
func ByTrackedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrackedAt, opts...).ToFunc()
}

// ByClassNumber orders the results by the class_number field.
func ByClassNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClassNumber, opts...).ToFunc()
}

// ByTerm orders the results by the term field.
func ByTerm(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTerm, opts...).ToFunc()
}
