// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/pongsakorn-maker/entgo-playground/ent/playlist"
	"github.com/pongsakorn-maker/entgo-playground/ent/playlistvideo"
	"github.com/pongsakorn-maker/entgo-playground/ent/predicate"
	"github.com/pongsakorn-maker/entgo-playground/ent/resolution"
	"github.com/pongsakorn-maker/entgo-playground/ent/video"
)

// PlaylistVideoUpdate is the builder for updating PlaylistVideo entities.
type PlaylistVideoUpdate struct {
	config
	hooks      []Hook
	mutation   *PlaylistVideoMutation
	predicates []predicate.PlaylistVideo
}

// Where adds a new predicate for the builder.
func (pvu *PlaylistVideoUpdate) Where(ps ...predicate.PlaylistVideo) *PlaylistVideoUpdate {
	pvu.predicates = append(pvu.predicates, ps...)
	return pvu
}

// SetPlaylistVideoID sets the playlistVideo_id field.
func (pvu *PlaylistVideoUpdate) SetPlaylistVideoID(i int) *PlaylistVideoUpdate {
	pvu.mutation.ResetPlaylistVideoID()
	pvu.mutation.SetPlaylistVideoID(i)
	return pvu
}

// AddPlaylistVideoID adds i to playlistVideo_id.
func (pvu *PlaylistVideoUpdate) AddPlaylistVideoID(i int) *PlaylistVideoUpdate {
	pvu.mutation.AddPlaylistVideoID(i)
	return pvu
}

// SetVideoID sets the video edge to Video by id.
func (pvu *PlaylistVideoUpdate) SetVideoID(id int) *PlaylistVideoUpdate {
	pvu.mutation.SetVideoID(id)
	return pvu
}

// SetNillableVideoID sets the video edge to Video by id if the given value is not nil.
func (pvu *PlaylistVideoUpdate) SetNillableVideoID(id *int) *PlaylistVideoUpdate {
	if id != nil {
		pvu = pvu.SetVideoID(*id)
	}
	return pvu
}

// SetVideo sets the video edge to Video.
func (pvu *PlaylistVideoUpdate) SetVideo(v *Video) *PlaylistVideoUpdate {
	return pvu.SetVideoID(v.ID)
}

// SetPlaylistsID sets the playlists edge to Playlist by id.
func (pvu *PlaylistVideoUpdate) SetPlaylistsID(id int) *PlaylistVideoUpdate {
	pvu.mutation.SetPlaylistsID(id)
	return pvu
}

// SetNillablePlaylistsID sets the playlists edge to Playlist by id if the given value is not nil.
func (pvu *PlaylistVideoUpdate) SetNillablePlaylistsID(id *int) *PlaylistVideoUpdate {
	if id != nil {
		pvu = pvu.SetPlaylistsID(*id)
	}
	return pvu
}

// SetPlaylists sets the playlists edge to Playlist.
func (pvu *PlaylistVideoUpdate) SetPlaylists(p *Playlist) *PlaylistVideoUpdate {
	return pvu.SetPlaylistsID(p.ID)
}

// SetResolutionID sets the resolution edge to Resolution by id.
func (pvu *PlaylistVideoUpdate) SetResolutionID(id int) *PlaylistVideoUpdate {
	pvu.mutation.SetResolutionID(id)
	return pvu
}

// SetNillableResolutionID sets the resolution edge to Resolution by id if the given value is not nil.
func (pvu *PlaylistVideoUpdate) SetNillableResolutionID(id *int) *PlaylistVideoUpdate {
	if id != nil {
		pvu = pvu.SetResolutionID(*id)
	}
	return pvu
}

// SetResolution sets the resolution edge to Resolution.
func (pvu *PlaylistVideoUpdate) SetResolution(r *Resolution) *PlaylistVideoUpdate {
	return pvu.SetResolutionID(r.ID)
}

// Mutation returns the PlaylistVideoMutation object of the builder.
func (pvu *PlaylistVideoUpdate) Mutation() *PlaylistVideoMutation {
	return pvu.mutation
}

// ClearVideo clears the video edge to Video.
func (pvu *PlaylistVideoUpdate) ClearVideo() *PlaylistVideoUpdate {
	pvu.mutation.ClearVideo()
	return pvu
}

// ClearPlaylists clears the playlists edge to Playlist.
func (pvu *PlaylistVideoUpdate) ClearPlaylists() *PlaylistVideoUpdate {
	pvu.mutation.ClearPlaylists()
	return pvu
}

// ClearResolution clears the resolution edge to Resolution.
func (pvu *PlaylistVideoUpdate) ClearResolution() *PlaylistVideoUpdate {
	pvu.mutation.ClearResolution()
	return pvu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pvu *PlaylistVideoUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(pvu.hooks) == 0 {
		affected, err = pvu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistVideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pvu.mutation = mutation
			affected, err = pvu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pvu.hooks) - 1; i >= 0; i-- {
			mut = pvu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pvu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pvu *PlaylistVideoUpdate) SaveX(ctx context.Context) int {
	affected, err := pvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pvu *PlaylistVideoUpdate) Exec(ctx context.Context) error {
	_, err := pvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvu *PlaylistVideoUpdate) ExecX(ctx context.Context) {
	if err := pvu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pvu *PlaylistVideoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playlistvideo.Table,
			Columns: playlistvideo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playlistvideo.FieldID,
			},
		},
	}
	if ps := pvu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pvu.mutation.PlaylistVideoID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playlistvideo.FieldPlaylistVideoID,
		})
	}
	if value, ok := pvu.mutation.AddedPlaylistVideoID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playlistvideo.FieldPlaylistVideoID,
		})
	}
	if pvu.mutation.VideoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.VideoTable,
			Columns: []string{playlistvideo.VideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: video.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.VideoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.VideoTable,
			Columns: []string{playlistvideo.VideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pvu.mutation.PlaylistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.PlaylistsTable,
			Columns: []string{playlistvideo.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.PlaylistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.PlaylistsTable,
			Columns: []string{playlistvideo.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pvu.mutation.ResolutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.ResolutionTable,
			Columns: []string{playlistvideo.ResolutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resolution.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvu.mutation.ResolutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.ResolutionTable,
			Columns: []string{playlistvideo.ResolutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resolution.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playlistvideo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PlaylistVideoUpdateOne is the builder for updating a single PlaylistVideo entity.
type PlaylistVideoUpdateOne struct {
	config
	hooks    []Hook
	mutation *PlaylistVideoMutation
}

// SetPlaylistVideoID sets the playlistVideo_id field.
func (pvuo *PlaylistVideoUpdateOne) SetPlaylistVideoID(i int) *PlaylistVideoUpdateOne {
	pvuo.mutation.ResetPlaylistVideoID()
	pvuo.mutation.SetPlaylistVideoID(i)
	return pvuo
}

// AddPlaylistVideoID adds i to playlistVideo_id.
func (pvuo *PlaylistVideoUpdateOne) AddPlaylistVideoID(i int) *PlaylistVideoUpdateOne {
	pvuo.mutation.AddPlaylistVideoID(i)
	return pvuo
}

// SetVideoID sets the video edge to Video by id.
func (pvuo *PlaylistVideoUpdateOne) SetVideoID(id int) *PlaylistVideoUpdateOne {
	pvuo.mutation.SetVideoID(id)
	return pvuo
}

// SetNillableVideoID sets the video edge to Video by id if the given value is not nil.
func (pvuo *PlaylistVideoUpdateOne) SetNillableVideoID(id *int) *PlaylistVideoUpdateOne {
	if id != nil {
		pvuo = pvuo.SetVideoID(*id)
	}
	return pvuo
}

// SetVideo sets the video edge to Video.
func (pvuo *PlaylistVideoUpdateOne) SetVideo(v *Video) *PlaylistVideoUpdateOne {
	return pvuo.SetVideoID(v.ID)
}

// SetPlaylistsID sets the playlists edge to Playlist by id.
func (pvuo *PlaylistVideoUpdateOne) SetPlaylistsID(id int) *PlaylistVideoUpdateOne {
	pvuo.mutation.SetPlaylistsID(id)
	return pvuo
}

// SetNillablePlaylistsID sets the playlists edge to Playlist by id if the given value is not nil.
func (pvuo *PlaylistVideoUpdateOne) SetNillablePlaylistsID(id *int) *PlaylistVideoUpdateOne {
	if id != nil {
		pvuo = pvuo.SetPlaylistsID(*id)
	}
	return pvuo
}

// SetPlaylists sets the playlists edge to Playlist.
func (pvuo *PlaylistVideoUpdateOne) SetPlaylists(p *Playlist) *PlaylistVideoUpdateOne {
	return pvuo.SetPlaylistsID(p.ID)
}

// SetResolutionID sets the resolution edge to Resolution by id.
func (pvuo *PlaylistVideoUpdateOne) SetResolutionID(id int) *PlaylistVideoUpdateOne {
	pvuo.mutation.SetResolutionID(id)
	return pvuo
}

// SetNillableResolutionID sets the resolution edge to Resolution by id if the given value is not nil.
func (pvuo *PlaylistVideoUpdateOne) SetNillableResolutionID(id *int) *PlaylistVideoUpdateOne {
	if id != nil {
		pvuo = pvuo.SetResolutionID(*id)
	}
	return pvuo
}

// SetResolution sets the resolution edge to Resolution.
func (pvuo *PlaylistVideoUpdateOne) SetResolution(r *Resolution) *PlaylistVideoUpdateOne {
	return pvuo.SetResolutionID(r.ID)
}

// Mutation returns the PlaylistVideoMutation object of the builder.
func (pvuo *PlaylistVideoUpdateOne) Mutation() *PlaylistVideoMutation {
	return pvuo.mutation
}

// ClearVideo clears the video edge to Video.
func (pvuo *PlaylistVideoUpdateOne) ClearVideo() *PlaylistVideoUpdateOne {
	pvuo.mutation.ClearVideo()
	return pvuo
}

// ClearPlaylists clears the playlists edge to Playlist.
func (pvuo *PlaylistVideoUpdateOne) ClearPlaylists() *PlaylistVideoUpdateOne {
	pvuo.mutation.ClearPlaylists()
	return pvuo
}

// ClearResolution clears the resolution edge to Resolution.
func (pvuo *PlaylistVideoUpdateOne) ClearResolution() *PlaylistVideoUpdateOne {
	pvuo.mutation.ClearResolution()
	return pvuo
}

// Save executes the query and returns the updated entity.
func (pvuo *PlaylistVideoUpdateOne) Save(ctx context.Context) (*PlaylistVideo, error) {

	var (
		err  error
		node *PlaylistVideo
	)
	if len(pvuo.hooks) == 0 {
		node, err = pvuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlaylistVideoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pvuo.mutation = mutation
			node, err = pvuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pvuo.hooks) - 1; i >= 0; i-- {
			mut = pvuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pvuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pvuo *PlaylistVideoUpdateOne) SaveX(ctx context.Context) *PlaylistVideo {
	pv, err := pvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pv
}

// Exec executes the query on the entity.
func (pvuo *PlaylistVideoUpdateOne) Exec(ctx context.Context) error {
	_, err := pvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pvuo *PlaylistVideoUpdateOne) ExecX(ctx context.Context) {
	if err := pvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pvuo *PlaylistVideoUpdateOne) sqlSave(ctx context.Context) (pv *PlaylistVideo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playlistvideo.Table,
			Columns: playlistvideo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playlistvideo.FieldID,
			},
		},
	}
	id, ok := pvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PlaylistVideo.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := pvuo.mutation.PlaylistVideoID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playlistvideo.FieldPlaylistVideoID,
		})
	}
	if value, ok := pvuo.mutation.AddedPlaylistVideoID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: playlistvideo.FieldPlaylistVideoID,
		})
	}
	if pvuo.mutation.VideoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.VideoTable,
			Columns: []string{playlistvideo.VideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: video.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.VideoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.VideoTable,
			Columns: []string{playlistvideo.VideoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: video.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pvuo.mutation.PlaylistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.PlaylistsTable,
			Columns: []string{playlistvideo.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.PlaylistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.PlaylistsTable,
			Columns: []string{playlistvideo.PlaylistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: playlist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pvuo.mutation.ResolutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.ResolutionTable,
			Columns: []string{playlistvideo.ResolutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resolution.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pvuo.mutation.ResolutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playlistvideo.ResolutionTable,
			Columns: []string{playlistvideo.ResolutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resolution.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pv = &PlaylistVideo{config: pvuo.config}
	_spec.Assign = pv.assignValues
	_spec.ScanValues = pv.scanValues()
	if err = sqlgraph.UpdateNode(ctx, pvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{playlistvideo.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pv, nil
}
