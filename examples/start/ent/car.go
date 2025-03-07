// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Car is the model entity for the Car schema.
type Car struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// RegisteredAt holds the value of the "registered_at" field.
	RegisteredAt time.Time `json:"registered_at,omitempty"`
}

// FromRows scans the sql response data into Car.
func (c *Car) FromRows(rows *sql.Rows) error {
	var vc struct {
		ID           int
		Model        sql.NullString
		RegisteredAt sql.NullTime
	}
	// the order here should be the same as in the `car.Columns`.
	if err := rows.Scan(
		&vc.ID,
		&vc.Model,
		&vc.RegisteredAt,
	); err != nil {
		return err
	}
	c.ID = vc.ID
	c.Model = vc.Model.String
	c.RegisteredAt = vc.RegisteredAt.Time
	return nil
}

// QueryOwner queries the owner edge of the Car.
func (c *Car) QueryOwner() *UserQuery {
	return (&CarClient{c.config}).QueryOwner(c)
}

// Update returns a builder for updating this Car.
// Note that, you need to call Car.Unwrap() before calling this method, if this Car
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Car) Update() *CarUpdateOne {
	return (&CarClient{c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Car) Unwrap() *Car {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Car is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Car) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("Car(")
	buf.WriteString(fmt.Sprintf("id=%v", c.ID))
	buf.WriteString(fmt.Sprintf(", model=%v", c.Model))
	buf.WriteString(fmt.Sprintf(", registered_at=%v", c.RegisteredAt))
	buf.WriteString(")")
	return buf.String()
}

// Cars is a parsable slice of Car.
type Cars []*Car

// FromRows scans the sql response data into Cars.
func (c *Cars) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vc := &Car{}
		if err := vc.FromRows(rows); err != nil {
			return err
		}
		*c = append(*c, vc)
	}
	return nil
}

func (c Cars) config(cfg config) {
	for i := range c {
		c[i].config = cfg
	}
}
