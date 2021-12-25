// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"github.com/giantswarm/graph/ent/cluster"
)

// ClusterCreate is the builder for creating a Cluster entity.
type ClusterCreate struct {
	config
	mutation *ClusterMutation
	hooks    []Hook
}

// Mutation returns the ClusterMutation object of the builder.
func (cc *ClusterCreate) Mutation() *ClusterMutation {
	return cc.mutation
}

// Save creates the Cluster in the database.
func (cc *ClusterCreate) Save(ctx context.Context) (*Cluster, error) {
	var (
		err  error
		node *Cluster
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClusterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.gremlinSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClusterCreate) SaveX(ctx context.Context) *Cluster {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClusterCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClusterCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClusterCreate) check() error {
	return nil
}

func (cc *ClusterCreate) gremlinSave(ctx context.Context) (*Cluster, error) {
	res := &gremlin.Response{}
	query, bindings := cc.gremlin().Query()
	if err := cc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	c := &Cluster{config: cc.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cc *ClusterCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		firstQueryPred *dsl.Traversal // constraint predicate.
		pred           *dsl.Traversal // constraint predicate.
		test           *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 0)
	createTraversal := func(constraints []*constraint, traversalFuncs []func(*dsl.Traversal)) *dsl.Traversal {
		var v *dsl.Traversal
		if len(constraints) > 0 {
			// We will use coalesce, therefore AddV will be child traversal, so we need __ here
			v = __.New().AddV(cluster.Label)
		} else {
			v = g.AddV(cluster.Label)
		}
		for _, tf := range traversalFuncs {
			tf(v)
		}
		return v
	}
	traversalFuncs := []func(*dsl.Traversal){}
	v := createTraversal(constraints, traversalFuncs)
	return v.ValueMap(true)
}

// ClusterCreateBulk is the builder for creating many Cluster entities in bulk.
type ClusterCreateBulk struct {
	config
	builders []*ClusterCreate
}
