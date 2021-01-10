// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/chanudomsr/app/ent/checkout"
	"github.com/chanudomsr/app/ent/employee"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// EmployeeCreate is the builder for creating a Employee entity.
type EmployeeCreate struct {
	config
	mutation *EmployeeMutation
	hooks    []Hook
}

// SetEMPLOYEENAME sets the EMPLOYEENAME field.
func (ec *EmployeeCreate) SetEMPLOYEENAME(s string) *EmployeeCreate {
	ec.mutation.SetEMPLOYEENAME(s)
	return ec
}

// SetEMPLOYEEPASSWORD sets the EMPLOYEEPASSWORD field.
func (ec *EmployeeCreate) SetEMPLOYEEPASSWORD(s string) *EmployeeCreate {
	ec.mutation.SetEMPLOYEEPASSWORD(s)
	return ec
}

// AddCheckout2IDs adds the checkout2 edge to Checkout by ids.
func (ec *EmployeeCreate) AddCheckout2IDs(ids ...int) *EmployeeCreate {
	ec.mutation.AddCheckout2IDs(ids...)
	return ec
}

// AddCheckout2 adds the checkout2 edges to Checkout.
func (ec *EmployeeCreate) AddCheckout2(c ...*Checkout) *EmployeeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ec.AddCheckout2IDs(ids...)
}

// Mutation returns the EmployeeMutation object of the builder.
func (ec *EmployeeCreate) Mutation() *EmployeeMutation {
	return ec.mutation
}

// Save creates the Employee in the database.
func (ec *EmployeeCreate) Save(ctx context.Context) (*Employee, error) {
	if _, ok := ec.mutation.EMPLOYEENAME(); !ok {
		return nil, &ValidationError{Name: "EMPLOYEENAME", err: errors.New("ent: missing required field \"EMPLOYEENAME\"")}
	}
	if _, ok := ec.mutation.EMPLOYEEPASSWORD(); !ok {
		return nil, &ValidationError{Name: "EMPLOYEEPASSWORD", err: errors.New("ent: missing required field \"EMPLOYEEPASSWORD\"")}
	}
	var (
		err  error
		node *Employee
	)
	if len(ec.hooks) == 0 {
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EmployeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ec.mutation = mutation
			node, err = ec.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EmployeeCreate) SaveX(ctx context.Context) *Employee {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ec *EmployeeCreate) sqlSave(ctx context.Context) (*Employee, error) {
	e, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	e.ID = int(id)
	return e, nil
}

func (ec *EmployeeCreate) createSpec() (*Employee, *sqlgraph.CreateSpec) {
	var (
		e     = &Employee{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: employee.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: employee.FieldID,
			},
		}
	)
	if value, ok := ec.mutation.EMPLOYEENAME(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEMPLOYEENAME,
		})
		e.EMPLOYEENAME = value
	}
	if value, ok := ec.mutation.EMPLOYEEPASSWORD(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: employee.FieldEMPLOYEEPASSWORD,
		})
		e.EMPLOYEEPASSWORD = value
	}
	if nodes := ec.mutation.Checkout2IDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   employee.Checkout2Table,
			Columns: []string{employee.Checkout2Column},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: checkout.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return e, _spec
}
