// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/pongsakorn-maker/entgo-playground/ent/playlistvideo"
)

// PlaylistVideoCreate is the builder for creating a PlaylistVideo entity.
type PlaylistVideoCreate struct {
	config
	mutation *PlaylistVideoMutation
	hooks    []Hook
}

// SetPlaylistVideoID sets the PlaylistVideo_ID field.
func (pvc *PlaylistVideoCreate) SetPlaylistVideoID(i int) *PlaylistVideoCreate {
	pvc.mutation.SetPlaylistVideoID(i)
	return pvc
}

// Mutation returns the PlaylistVideoMutation object of the builder.
func (pvc *PlaylistVideoCreate) Mutation() *PlaylistVideoMutation {
	return pvc.mutation
}

// Save creates the PlaylistVideo in the database.
func (pvc *PlaylistVideoCreate) Save(ctx context.Context) (*PlaylistVideo, error) {
	if err := pvc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *PlaylistVideo
	)
	if len(pvc.hooks) == 0 {
		node, err = pvc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistVideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pvc.mutation = mutation
			node, err = pvc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pvc.hooks) - 1; i >= 0; i-- {
			mut = pvc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pvc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pvc *PlaylistVideoCreate) SaveX(ctx context.Context) *PlaylistVideo {
	v, err := pvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pvc *PlaylistVideoCreate) preSave() error {
	if _, ok := pvc.mutation.PlaylistVideoID(); !ok {
		return &ValidationError{Name: "PlaylistVideo_ID", err: errors.New("ent: missing required field \"PlaylistVideo_ID\"")}
	}
	return nil
}

func (pvc *PlaylistVideoCreate) sqlSave(ctx context.Context) (*PlaylistVideo, error) {
	pv, _spec := pvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pvc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pv.ID = int(id)
	return pv, nil
}

func (pvc *PlaylistVideoCreate) createSpec() (*PlaylistVideo, *sqlgraph.CreateSpec) {
	var (
		pv    = &PlaylistVideo{config: pvc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: playlistvideo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playlistvideo.FieldID,
			},
		}
	)
	if value, ok := pvc.mutation.PlaylistVideoID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playlistvideo.FieldPlaylistVideoID,
		})
		pv.PlaylistVideoID = value
	}
	return pv, _spec
}

// PlaylistVideoCreateBulk is the builder for creating a bulk of PlaylistVideo entities.
type PlaylistVideoCreateBulk struct {
	config
	builders []*PlaylistVideoCreate
}

// Save creates the PlaylistVideo entities in the database.
func (pvcb *PlaylistVideoCreateBulk) Save(ctx context.Context) ([]*PlaylistVideo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pvcb.builders))
	nodes := make([]*PlaylistVideo, len(pvcb.builders))
	mutators := make([]Mutator, len(pvcb.builders))
	for i := range pvcb.builders {
		func(i int, root context.Context) {
			builder := pvcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*PlaylistVideoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pvcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pvcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (pvcb *PlaylistVideoCreateBulk) SaveX(ctx context.Context) []*PlaylistVideo {
	v, err := pvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
