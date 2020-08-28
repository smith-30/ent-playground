// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeCars holds the string denoting the cars edge name in mutations.
	EdgeCars = "cars"

	// Table holds the table name of the user in the database.
	Table = "users"
	// CarsTable is the table the holds the cars relation/edge.
	CarsTable = "cars"
	// CarsInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarsInverseTable = "cars"
	// CarsColumn is the table column denoting the cars relation/edge.
	CarsColumn = "user_cars"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
}

var (
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
	// DefaultName holds the default value on creation for the name field.
	DefaultName string
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
