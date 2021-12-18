// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"github.com/giantswarm/graph/ent/person"
	"github.com/giantswarm/graph/ent/predicate"
)

// PersonDelete is the builder for deleting a Person entity.
type PersonDelete struct {
	config
	hooks    []Hook
	mutation *PersonMutation
}

// Where appends a list predicates to the PersonDelete builder.
func (pd *PersonDelete) Where(ps ...predicate.Person) *PersonDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PersonDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pd.hooks) == 0 {
		affected, err = pd.gremlinExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pd.mutation = mutation
			affected, err = pd.gremlinExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pd.hooks) - 1; i >= 0; i-- {
			if pd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PersonDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PersonDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := pd.gremlin().Query()
	if err := pd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (pd *PersonDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(person.Label)
	for _, p := range pd.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// PersonDeleteOne is the builder for deleting a single Person entity.
type PersonDeleteOne struct {
	pd *PersonDelete
}

// Exec executes the deletion query.
func (pdo *PersonDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{person.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PersonDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}