// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
)

// Cluster is the model entity for the Cluster schema.
type Cluster struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// FromResponse scans the gremlin response data into Cluster.
func (c *Cluster) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanc struct {
		ID int `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scanc); err != nil {
		return err
	}
	c.ID = scanc.ID
	return nil
}

// Update returns a builder for updating this Cluster.
// Note that you need to call Cluster.Unwrap() before calling this method if this Cluster
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cluster) Update() *ClusterUpdateOne {
	return (&ClusterClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Cluster entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cluster) Unwrap() *Cluster {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cluster is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cluster) String() string {
	var builder strings.Builder
	builder.WriteString("Cluster(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Clusters is a parsable slice of Cluster.
type Clusters []*Cluster

// FromResponse scans the gremlin response data into Clusters.
func (c *Clusters) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanc []struct {
		ID int `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scanc); err != nil {
		return err
	}
	for _, v := range scanc {
		*c = append(*c, &Cluster{
			ID: v.ID,
		})
	}
	return nil
}

func (c Clusters) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
