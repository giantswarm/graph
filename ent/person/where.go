// Code generated by entc, DO NOT EDIT.

package person

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"github.com/giantswarm/graph/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.EQ(v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EQ(v))
	})
}

// IsGiantSwarmEmployee applies equality check predicate on the "isGiantSwarmEmployee" field. It's identical to IsGiantSwarmEmployeeEQ.
func IsGiantSwarmEmployee(v bool) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldIsGiantSwarmEmployee, p.EQ(v))
	})
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.EQ(v))
	})
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.NEQ(v))
	})
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Person {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.Within(v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Person {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.Without(v...))
	})
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.GT(v))
	})
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.GTE(v))
	})
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.LT(v))
	})
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.LTE(v))
	})
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.Containing(v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.StartingWith(v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmail, p.EndingWith(v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EQ(v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.NEQ(v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Person {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Within(v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Person {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Without(v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.GT(v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.GTE(v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.LT(v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.LTE(v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Containing(v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.StartingWith(v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EndingWith(v))
	})
}

// IsGiantSwarmEmployeeEQ applies the EQ predicate on the "isGiantSwarmEmployee" field.
func IsGiantSwarmEmployeeEQ(v bool) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldIsGiantSwarmEmployee, p.EQ(v))
	})
}

// IsGiantSwarmEmployeeNEQ applies the NEQ predicate on the "isGiantSwarmEmployee" field.
func IsGiantSwarmEmployeeNEQ(v bool) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.Has(Label, FieldIsGiantSwarmEmployee, p.NEQ(v))
	})
}

// HasGitHubAccount applies the HasEdge predicate on the "gitHubAccount" edge.
func HasGitHubAccount() predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		t.OutE(GitHubAccountLabel).OutV()
	})
}

// HasGitHubAccountWith applies the HasEdge predicate on the "gitHubAccount" edge with a given conditions (other predicates).
func HasGitHubAccountWith(preds ...predicate.GitHubUser) predicate.Person {
	return predicate.Person(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(GitHubAccountLabel).Where(tr).OutV()
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.And(trs...))
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.Or(trs...))
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Person) predicate.Person {
	return predicate.Person(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}