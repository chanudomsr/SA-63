// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/chanudomsr/app/ent/customer"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NAME holds the value of the "NAME" field.
	NAME string `json:"NAME,omitempty"`
	// AGE holds the value of the "AGE" field.
	AGE int `json:"AGE,omitempty"`
	// EMAIL holds the value of the "EMAIL" field.
	EMAIL string `json:"EMAIL,omitempty"`
	// USERNAME holds the value of the "USERNAME" field.
	USERNAME string `json:"USERNAME,omitempty"`
	// PASSWORD holds the value of the "PASSWORD" field.
	PASSWORD string `json:"PASSWORD,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges CustomerEdges `json:"edges"`
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Checkout1 holds the value of the checkout1 edge.
	Checkout1 []*Checkout
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// Checkout1OrErr returns the Checkout1 value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) Checkout1OrErr() ([]*Checkout, error) {
	if e.loadedTypes[0] {
		return e.Checkout1, nil
	}
	return nil, &NotLoadedError{edge: "checkout1"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // NAME
		&sql.NullInt64{},  // AGE
		&sql.NullString{}, // EMAIL
		&sql.NullString{}, // USERNAME
		&sql.NullString{}, // PASSWORD
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(values ...interface{}) error {
	if m, n := len(values), len(customer.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field NAME", values[0])
	} else if value.Valid {
		c.NAME = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field AGE", values[1])
	} else if value.Valid {
		c.AGE = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field EMAIL", values[2])
	} else if value.Valid {
		c.EMAIL = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field USERNAME", values[3])
	} else if value.Valid {
		c.USERNAME = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PASSWORD", values[4])
	} else if value.Valid {
		c.PASSWORD = value.String
	}
	return nil
}

// QueryCheckout1 queries the checkout1 edge of the Customer.
func (c *Customer) QueryCheckout1() *CheckoutQuery {
	return (&CustomerClient{config: c.config}).QueryCheckout1(c)
}

// Update returns a builder for updating this Customer.
// Note that, you need to call Customer.Unwrap() before calling this method, if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", NAME=")
	builder.WriteString(c.NAME)
	builder.WriteString(", AGE=")
	builder.WriteString(fmt.Sprintf("%v", c.AGE))
	builder.WriteString(", EMAIL=")
	builder.WriteString(c.EMAIL)
	builder.WriteString(", USERNAME=")
	builder.WriteString(c.USERNAME)
	builder.WriteString(", PASSWORD=")
	builder.WriteString(c.PASSWORD)
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer

func (c Customers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}