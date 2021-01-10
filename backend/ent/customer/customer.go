// Code generated by entc, DO NOT EDIT.

package customer

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNAME holds the string denoting the name field in the database.
	FieldNAME = "name"
	// FieldAGE holds the string denoting the age field in the database.
	FieldAGE = "age"
	// FieldEMAIL holds the string denoting the email field in the database.
	FieldEMAIL = "email"
	// FieldUSERNAME holds the string denoting the username field in the database.
	FieldUSERNAME = "username"
	// FieldPASSWORD holds the string denoting the password field in the database.
	FieldPASSWORD = "password"

	// EdgeCheckout1 holds the string denoting the checkout1 edge name in mutations.
	EdgeCheckout1 = "checkout1"

	// Table holds the table name of the customer in the database.
	Table = "customers"
	// Checkout1Table is the table the holds the checkout1 relation/edge.
	Checkout1Table = "checkouts"
	// Checkout1InverseTable is the table name for the Checkout entity.
	// It exists in this package in order to avoid circular dependency with the "checkout" package.
	Checkout1InverseTable = "checkouts"
	// Checkout1Column is the table column denoting the checkout1 relation/edge.
	Checkout1Column = "customer_checkout1"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldNAME,
	FieldAGE,
	FieldEMAIL,
	FieldUSERNAME,
	FieldPASSWORD,
}