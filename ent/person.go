// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
	"github.com/giantswarm/graph/ent/githubuser"
)

// Person is the model entity for the Person schema.
type Person struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsGiantSwarmEmployee holds the value of the "isGiantSwarmEmployee" field.
	IsGiantSwarmEmployee bool `json:"isGiantSwarmEmployee,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonQuery when eager-loading is set.
	Edges PersonEdges `json:"edges"`
}

// PersonEdges holds the relations/edges for other nodes in the graph.
type PersonEdges struct {
	// GitHubAccount holds the value of the gitHubAccount edge.
	GitHubAccount *GitHubUser `json:"gitHubAccount,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GitHubAccountOrErr returns the GitHubAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) GitHubAccountOrErr() (*GitHubUser, error) {
	if e.loadedTypes[0] {
		if e.GitHubAccount == nil {
			// The edge gitHubAccount was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: githubuser.Label}
		}
		return e.GitHubAccount, nil
	}
	return nil, &NotLoadedError{edge: "gitHubAccount"}
}

// FromResponse scans the gremlin response data into Person.
func (pe *Person) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanpe struct {
		ID                   int    `json:"id,omitempty"`
		Email                string `json:"email,omitempty"`
		Name                 string `json:"name,omitempty"`
		IsGiantSwarmEmployee bool   `json:"is_giant_swarm_employee,omitempty"`
	}
	if err := vmap.Decode(&scanpe); err != nil {
		return err
	}
	pe.ID = scanpe.ID
	pe.Email = scanpe.Email
	pe.Name = scanpe.Name
	pe.IsGiantSwarmEmployee = scanpe.IsGiantSwarmEmployee
	return nil
}

// QueryGitHubAccount queries the "gitHubAccount" edge of the Person entity.
func (pe *Person) QueryGitHubAccount() *GitHubUserQuery {
	return (&PersonClient{config: pe.config}).QueryGitHubAccount(pe)
}

// Update returns a builder for updating this Person.
// Note that you need to call Person.Unwrap() before calling this method if this Person
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Person) Update() *PersonUpdateOne {
	return (&PersonClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the Person entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Person) Unwrap() *Person {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Person is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Person) String() string {
	var builder strings.Builder
	builder.WriteString("Person(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", email=")
	builder.WriteString(pe.Email)
	builder.WriteString(", name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", isGiantSwarmEmployee=")
	builder.WriteString(fmt.Sprintf("%v", pe.IsGiantSwarmEmployee))
	builder.WriteByte(')')
	return builder.String()
}

// Persons is a parsable slice of Person.
type Persons []*Person

// FromResponse scans the gremlin response data into Persons.
func (pe *Persons) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanpe []struct {
		ID                   int    `json:"id,omitempty"`
		Email                string `json:"email,omitempty"`
		Name                 string `json:"name,omitempty"`
		IsGiantSwarmEmployee bool   `json:"is_giant_swarm_employee,omitempty"`
	}
	if err := vmap.Decode(&scanpe); err != nil {
		return err
	}
	for _, v := range scanpe {
		*pe = append(*pe, &Person{
			ID:                   v.ID,
			Email:                v.Email,
			Name:                 v.Name,
			IsGiantSwarmEmployee: v.IsGiantSwarmEmployee,
		})
	}
	return nil
}

func (pe Persons) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
