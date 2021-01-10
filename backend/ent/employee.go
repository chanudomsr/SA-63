// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/chanudomsr/app/ent/employee"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Employee is the model entity for the Employee schema.
type Employee struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EMPLOYEENAME holds the value of the "EMPLOYEENAME" field.
	EMPLOYEENAME string `json:"EMPLOYEENAME,omitempty"`
	// EMPLOYEEPASSWORD holds the value of the "EMPLOYEEPASSWORD" field.
	EMPLOYEEPASSWORD string `json:"EMPLOYEEPASSWORD,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EmployeeQuery when eager-loading is set.
	Edges EmployeeEdges `json:"edges"`
}

// EmployeeEdges holds the relations/edges for other nodes in the graph.
type EmployeeEdges struct {
	// Checkout2 holds the value of the checkout2 edge.
	Checkout2 []*Checkout
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// Checkout2OrErr returns the Checkout2 value or an error if the edge
// was not loaded in eager-loading.
func (e EmployeeEdges) Checkout2OrErr() ([]*Checkout, error) {
	if e.loadedTypes[0] {
		return e.Checkout2, nil
	}
	return nil, &NotLoadedError{edge: "checkout2"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Employee) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // EMPLOYEENAME
		&sql.NullString{}, // EMPLOYEEPASSWORD
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Employee fields.
func (e *Employee) assignValues(values ...interface{}) error {
	if m, n := len(values), len(employee.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	e.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field EMPLOYEENAME", values[0])
	} else if value.Valid {
		e.EMPLOYEENAME = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field EMPLOYEEPASSWORD", values[1])
	} else if value.Valid {
		e.EMPLOYEEPASSWORD = value.String
	}
	return nil
}

// QueryCheckout2 queries the checkout2 edge of the Employee.
func (e *Employee) QueryCheckout2() *CheckoutQuery {
	return (&EmployeeClient{config: e.config}).QueryCheckout2(e)
}

// Update returns a builder for updating this Employee.
// Note that, you need to call Employee.Unwrap() before calling this method, if this Employee
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Employee) Update() *EmployeeUpdateOne {
	return (&EmployeeClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (e *Employee) Unwrap() *Employee {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Employee is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Employee) String() string {
	var builder strings.Builder
	builder.WriteString("Employee(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", EMPLOYEENAME=")
	builder.WriteString(e.EMPLOYEENAME)
	builder.WriteString(", EMPLOYEEPASSWORD=")
	builder.WriteString(e.EMPLOYEEPASSWORD)
	builder.WriteByte(')')
	return builder.String()
}

// Employees is a parsable slice of Employee.
type Employees []*Employee

func (e Employees) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
