// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// CheckoutsColumns holds the columns for the "checkouts" table.
	CheckoutsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "checkoutdate", Type: field.TypeTime},
		{Name: "customer_checkout1", Type: field.TypeInt, Nullable: true},
		{Name: "employee_checkout2", Type: field.TypeInt, Nullable: true},
		{Name: "room_checkout3", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CheckoutsTable holds the schema information for the "checkouts" table.
	CheckoutsTable = &schema.Table{
		Name:       "checkouts",
		Columns:    CheckoutsColumns,
		PrimaryKey: []*schema.Column{CheckoutsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "checkouts_customers_checkout1",
				Columns: []*schema.Column{CheckoutsColumns[2]},

				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "checkouts_employees_checkout2",
				Columns: []*schema.Column{CheckoutsColumns[3]},

				RefColumns: []*schema.Column{EmployeesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "checkouts_rooms_checkout3",
				Columns: []*schema.Column{CheckoutsColumns[4]},

				RefColumns: []*schema.Column{RoomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
		{Name: "email", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:        "customers",
		Columns:     CustomersColumns,
		PrimaryKey:  []*schema.Column{CustomersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "employeename", Type: field.TypeString},
		{Name: "employeepassword", Type: field.TypeString},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:        "employees",
		Columns:     EmployeesColumns,
		PrimaryKey:  []*schema.Column{EmployeesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// RoomsColumns holds the columns for the "rooms" table.
	RoomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "roomnumber", Type: field.TypeString},
	}
	// RoomsTable holds the schema information for the "rooms" table.
	RoomsTable = &schema.Table{
		Name:        "rooms",
		Columns:     RoomsColumns,
		PrimaryKey:  []*schema.Column{RoomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CheckoutsTable,
		CustomersTable,
		EmployeesTable,
		RoomsTable,
	}
)

func init() {
	CheckoutsTable.ForeignKeys[0].RefTable = CustomersTable
	CheckoutsTable.ForeignKeys[1].RefTable = EmployeesTable
	CheckoutsTable.ForeignKeys[2].RefTable = RoomsTable
}