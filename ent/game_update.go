// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lending-system/ent/game"
	"lending-system/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GameUpdate is the builder for updating Game entities.
type GameUpdate struct {
	config
	hooks    []Hook
	mutation *GameMutation
}

// Where appends a list predicates to the GameUpdate builder.
func (gu *GameUpdate) Where(ps ...predicate.Game) *GameUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// Mutation returns the GameMutation object of the builder.
func (gu *GameUpdate) Mutation() *GameMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GameUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GameUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GameUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GameUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GameUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(game.Table, game.Columns, sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GameUpdateOne is the builder for updating a single Game entity.
type GameUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameMutation
}

// Mutation returns the GameMutation object of the builder.
func (guo *GameUpdateOne) Mutation() *GameMutation {
	return guo.mutation
}

// Where appends a list predicates to the GameUpdate builder.
func (guo *GameUpdateOne) Where(ps ...predicate.Game) *GameUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GameUpdateOne) Select(field string, fields ...string) *GameUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Game entity.
func (guo *GameUpdateOne) Save(ctx context.Context) (*Game, error) {
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GameUpdateOne) SaveX(ctx context.Context) *Game {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GameUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GameUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GameUpdateOne) sqlSave(ctx context.Context) (_node *Game, err error) {
	_spec := sqlgraph.NewUpdateSpec(game.Table, game.Columns, sqlgraph.NewFieldSpec(game.FieldID, field.TypeInt))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Game.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, game.FieldID)
		for _, f := range fields {
			if !game.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != game.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &Game{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{game.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
