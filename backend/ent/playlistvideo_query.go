// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/pongsakorn-maker/entgo-playground/ent/playlist"
	"github.com/pongsakorn-maker/entgo-playground/ent/playlistvideo"
	"github.com/pongsakorn-maker/entgo-playground/ent/predicate"
	"github.com/pongsakorn-maker/entgo-playground/ent/resolution"
	"github.com/pongsakorn-maker/entgo-playground/ent/video"
)

// PlaylistVideoQuery is the builder for querying PlaylistVideo entities.
type PlaylistVideoQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.PlaylistVideo
	// eager-loading edges.
	withVideo      *VideoQuery
	withPlaylists  *PlaylistQuery
	withResolution *ResolutionQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (pvq *PlaylistVideoQuery) Where(ps ...predicate.PlaylistVideo) *PlaylistVideoQuery {
	pvq.predicates = append(pvq.predicates, ps...)
	return pvq
}

// Limit adds a limit step to the query.
func (pvq *PlaylistVideoQuery) Limit(limit int) *PlaylistVideoQuery {
	pvq.limit = &limit
	return pvq
}

// Offset adds an offset step to the query.
func (pvq *PlaylistVideoQuery) Offset(offset int) *PlaylistVideoQuery {
	pvq.offset = &offset
	return pvq
}

// Order adds an order step to the query.
func (pvq *PlaylistVideoQuery) Order(o ...OrderFunc) *PlaylistVideoQuery {
	pvq.order = append(pvq.order, o...)
	return pvq
}

// QueryVideo chains the current query on the video edge.
func (pvq *PlaylistVideoQuery) QueryVideo() *VideoQuery {
	query := &VideoQuery{config: pvq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playlistvideo.Table, playlistvideo.FieldID, pvq.sqlQuery()),
			sqlgraph.To(video.Table, video.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, playlistvideo.VideoTable, playlistvideo.VideoColumn),
		)
		fromU = sqlgraph.SetNeighbors(pvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPlaylists chains the current query on the playlists edge.
func (pvq *PlaylistVideoQuery) QueryPlaylists() *PlaylistQuery {
	query := &PlaylistQuery{config: pvq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playlistvideo.Table, playlistvideo.FieldID, pvq.sqlQuery()),
			sqlgraph.To(playlist.Table, playlist.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, playlistvideo.PlaylistsTable, playlistvideo.PlaylistsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryResolution chains the current query on the resolution edge.
func (pvq *PlaylistVideoQuery) QueryResolution() *ResolutionQuery {
	query := &ResolutionQuery{config: pvq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(playlistvideo.Table, playlistvideo.FieldID, pvq.sqlQuery()),
			sqlgraph.To(resolution.Table, resolution.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, playlistvideo.ResolutionTable, playlistvideo.ResolutionColumn),
		)
		fromU = sqlgraph.SetNeighbors(pvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PlaylistVideo entity in the query. Returns *NotFoundError when no playlistvideo was found.
func (pvq *PlaylistVideoQuery) First(ctx context.Context) (*PlaylistVideo, error) {
	pvs, err := pvq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(pvs) == 0 {
		return nil, &NotFoundError{playlistvideo.Label}
	}
	return pvs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pvq *PlaylistVideoQuery) FirstX(ctx context.Context) *PlaylistVideo {
	pv, err := pvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return pv
}

// FirstID returns the first PlaylistVideo id in the query. Returns *NotFoundError when no id was found.
func (pvq *PlaylistVideoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pvq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{playlistvideo.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (pvq *PlaylistVideoQuery) FirstXID(ctx context.Context) int {
	id, err := pvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only PlaylistVideo entity in the query, returns an error if not exactly one entity was returned.
func (pvq *PlaylistVideoQuery) Only(ctx context.Context) (*PlaylistVideo, error) {
	pvs, err := pvq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(pvs) {
	case 1:
		return pvs[0], nil
	case 0:
		return nil, &NotFoundError{playlistvideo.Label}
	default:
		return nil, &NotSingularError{playlistvideo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pvq *PlaylistVideoQuery) OnlyX(ctx context.Context) *PlaylistVideo {
	pv, err := pvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return pv
}

// OnlyID returns the only PlaylistVideo id in the query, returns an error if not exactly one id was returned.
func (pvq *PlaylistVideoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pvq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{playlistvideo.Label}
	default:
		err = &NotSingularError{playlistvideo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pvq *PlaylistVideoQuery) OnlyIDX(ctx context.Context) int {
	id, err := pvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PlaylistVideos.
func (pvq *PlaylistVideoQuery) All(ctx context.Context) ([]*PlaylistVideo, error) {
	if err := pvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pvq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pvq *PlaylistVideoQuery) AllX(ctx context.Context) []*PlaylistVideo {
	pvs, err := pvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return pvs
}

// IDs executes the query and returns a list of PlaylistVideo ids.
func (pvq *PlaylistVideoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pvq.Select(playlistvideo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pvq *PlaylistVideoQuery) IDsX(ctx context.Context) []int {
	ids, err := pvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pvq *PlaylistVideoQuery) Count(ctx context.Context) (int, error) {
	if err := pvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pvq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pvq *PlaylistVideoQuery) CountX(ctx context.Context) int {
	count, err := pvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pvq *PlaylistVideoQuery) Exist(ctx context.Context) (bool, error) {
	if err := pvq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pvq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pvq *PlaylistVideoQuery) ExistX(ctx context.Context) bool {
	exist, err := pvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pvq *PlaylistVideoQuery) Clone() *PlaylistVideoQuery {
	return &PlaylistVideoQuery{
		config:     pvq.config,
		limit:      pvq.limit,
		offset:     pvq.offset,
		order:      append([]OrderFunc{}, pvq.order...),
		unique:     append([]string{}, pvq.unique...),
		predicates: append([]predicate.PlaylistVideo{}, pvq.predicates...),
		// clone intermediate query.
		sql:  pvq.sql.Clone(),
		path: pvq.path,
	}
}

//  WithVideo tells the query-builder to eager-loads the nodes that are connected to
// the "video" edge. The optional arguments used to configure the query builder of the edge.
func (pvq *PlaylistVideoQuery) WithVideo(opts ...func(*VideoQuery)) *PlaylistVideoQuery {
	query := &VideoQuery{config: pvq.config}
	for _, opt := range opts {
		opt(query)
	}
	pvq.withVideo = query
	return pvq
}

//  WithPlaylists tells the query-builder to eager-loads the nodes that are connected to
// the "playlists" edge. The optional arguments used to configure the query builder of the edge.
func (pvq *PlaylistVideoQuery) WithPlaylists(opts ...func(*PlaylistQuery)) *PlaylistVideoQuery {
	query := &PlaylistQuery{config: pvq.config}
	for _, opt := range opts {
		opt(query)
	}
	pvq.withPlaylists = query
	return pvq
}

//  WithResolution tells the query-builder to eager-loads the nodes that are connected to
// the "resolution" edge. The optional arguments used to configure the query builder of the edge.
func (pvq *PlaylistVideoQuery) WithResolution(opts ...func(*ResolutionQuery)) *PlaylistVideoQuery {
	query := &ResolutionQuery{config: pvq.config}
	for _, opt := range opts {
		opt(query)
	}
	pvq.withResolution = query
	return pvq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PlaylistVideoID int `json:"PlaylistVideo_ID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PlaylistVideo.Query().
//		GroupBy(playlistvideo.FieldPlaylistVideoID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pvq *PlaylistVideoQuery) GroupBy(field string, fields ...string) *PlaylistVideoGroupBy {
	group := &PlaylistVideoGroupBy{config: pvq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pvq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		PlaylistVideoID int `json:"PlaylistVideo_ID,omitempty"`
//	}
//
//	client.PlaylistVideo.Query().
//		Select(playlistvideo.FieldPlaylistVideoID).
//		Scan(ctx, &v)
//
func (pvq *PlaylistVideoQuery) Select(field string, fields ...string) *PlaylistVideoSelect {
	selector := &PlaylistVideoSelect{config: pvq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pvq.sqlQuery(), nil
	}
	return selector
}

func (pvq *PlaylistVideoQuery) prepareQuery(ctx context.Context) error {
	if pvq.path != nil {
		prev, err := pvq.path(ctx)
		if err != nil {
			return err
		}
		pvq.sql = prev
	}
	return nil
}

func (pvq *PlaylistVideoQuery) sqlAll(ctx context.Context) ([]*PlaylistVideo, error) {
	var (
		nodes       = []*PlaylistVideo{}
		withFKs     = pvq.withFKs
		_spec       = pvq.querySpec()
		loadedTypes = [3]bool{
			pvq.withVideo != nil,
			pvq.withPlaylists != nil,
			pvq.withResolution != nil,
		}
	)
	if pvq.withVideo != nil || pvq.withPlaylists != nil || pvq.withResolution != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, playlistvideo.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &PlaylistVideo{config: pvq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, pvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pvq.withVideo; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PlaylistVideo)
		for i := range nodes {
			if fk := nodes[i].video_playlist_videos; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(video.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "video_playlist_videos" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Video = n
			}
		}
	}

	if query := pvq.withPlaylists; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PlaylistVideo)
		for i := range nodes {
			if fk := nodes[i].playlist_playlist_videos; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(playlist.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "playlist_playlist_videos" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Playlists = n
			}
		}
	}

	if query := pvq.withResolution; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*PlaylistVideo)
		for i := range nodes {
			if fk := nodes[i].resolution_playlist_videos; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(resolution.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "resolution_playlist_videos" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Resolution = n
			}
		}
	}

	return nodes, nil
}

func (pvq *PlaylistVideoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pvq.querySpec()
	return sqlgraph.CountNodes(ctx, pvq.driver, _spec)
}

func (pvq *PlaylistVideoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pvq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (pvq *PlaylistVideoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   playlistvideo.Table,
			Columns: playlistvideo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: playlistvideo.FieldID,
			},
		},
		From:   pvq.sql,
		Unique: true,
	}
	if ps := pvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pvq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pvq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pvq *PlaylistVideoQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(pvq.driver.Dialect())
	t1 := builder.Table(playlistvideo.Table)
	selector := builder.Select(t1.Columns(playlistvideo.Columns...)...).From(t1)
	if pvq.sql != nil {
		selector = pvq.sql
		selector.Select(selector.Columns(playlistvideo.Columns...)...)
	}
	for _, p := range pvq.predicates {
		p(selector)
	}
	for _, p := range pvq.order {
		p(selector)
	}
	if offset := pvq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pvq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PlaylistVideoGroupBy is the builder for group-by PlaylistVideo entities.
type PlaylistVideoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pvgb *PlaylistVideoGroupBy) Aggregate(fns ...AggregateFunc) *PlaylistVideoGroupBy {
	pvgb.fns = append(pvgb.fns, fns...)
	return pvgb
}

// Scan applies the group-by query and scan the result into the given value.
func (pvgb *PlaylistVideoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pvgb.path(ctx)
	if err != nil {
		return err
	}
	pvgb.sql = query
	return pvgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pvgb *PlaylistVideoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pvgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (pvgb *PlaylistVideoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pvgb.fields) > 1 {
		return nil, errors.New("ent: PlaylistVideoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pvgb *PlaylistVideoGroupBy) StringsX(ctx context.Context) []string {
	v, err := pvgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (pvgb *PlaylistVideoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pvgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playlistvideo.Label}
	default:
		err = fmt.Errorf("ent: PlaylistVideoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pvgb *PlaylistVideoGroupBy) StringX(ctx context.Context) string {
	v, err := pvgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (pvgb *PlaylistVideoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pvgb.fields) > 1 {
		return nil, errors.New("ent: PlaylistVideoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pvgb *PlaylistVideoGroupBy) IntsX(ctx context.Context) []int {
	v, err := pvgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (pvgb *PlaylistVideoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pvgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playlistvideo.Label}
	default:
		err = fmt.Errorf("ent: PlaylistVideoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pvgb *PlaylistVideoGroupBy) IntX(ctx context.Context) int {
	v, err := pvgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (pvgb *PlaylistVideoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pvgb.fields) > 1 {
		return nil, errors.New("ent: PlaylistVideoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pvgb *PlaylistVideoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pvgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (pvgb *PlaylistVideoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pvgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playlistvideo.Label}
	default:
		err = fmt.Errorf("ent: PlaylistVideoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pvgb *PlaylistVideoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pvgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (pvgb *PlaylistVideoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pvgb.fields) > 1 {
		return nil, errors.New("ent: PlaylistVideoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pvgb *PlaylistVideoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pvgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (pvgb *PlaylistVideoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pvgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playlistvideo.Label}
	default:
		err = fmt.Errorf("ent: PlaylistVideoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pvgb *PlaylistVideoGroupBy) BoolX(ctx context.Context) bool {
	v, err := pvgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pvgb *PlaylistVideoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pvgb.sqlQuery().Query()
	if err := pvgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pvgb *PlaylistVideoGroupBy) sqlQuery() *sql.Selector {
	selector := pvgb.sql
	columns := make([]string, 0, len(pvgb.fields)+len(pvgb.fns))
	columns = append(columns, pvgb.fields...)
	for _, fn := range pvgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(pvgb.fields...)
}

// PlaylistVideoSelect is the builder for select fields of PlaylistVideo entities.
type PlaylistVideoSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (pvs *PlaylistVideoSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := pvs.path(ctx)
	if err != nil {
		return err
	}
	pvs.sql = query
	return pvs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pvs *PlaylistVideoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := pvs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (pvs *PlaylistVideoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(pvs.fields) > 1 {
		return nil, errors.New("ent: PlaylistVideoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := pvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pvs *PlaylistVideoSelect) StringsX(ctx context.Context) []string {
	v, err := pvs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (pvs *PlaylistVideoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pvs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playlistvideo.Label}
	default:
		err = fmt.Errorf("ent: PlaylistVideoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pvs *PlaylistVideoSelect) StringX(ctx context.Context) string {
	v, err := pvs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (pvs *PlaylistVideoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(pvs.fields) > 1 {
		return nil, errors.New("ent: PlaylistVideoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := pvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pvs *PlaylistVideoSelect) IntsX(ctx context.Context) []int {
	v, err := pvs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (pvs *PlaylistVideoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pvs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playlistvideo.Label}
	default:
		err = fmt.Errorf("ent: PlaylistVideoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pvs *PlaylistVideoSelect) IntX(ctx context.Context) int {
	v, err := pvs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (pvs *PlaylistVideoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(pvs.fields) > 1 {
		return nil, errors.New("ent: PlaylistVideoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := pvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pvs *PlaylistVideoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := pvs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (pvs *PlaylistVideoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pvs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playlistvideo.Label}
	default:
		err = fmt.Errorf("ent: PlaylistVideoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pvs *PlaylistVideoSelect) Float64X(ctx context.Context) float64 {
	v, err := pvs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (pvs *PlaylistVideoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(pvs.fields) > 1 {
		return nil, errors.New("ent: PlaylistVideoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := pvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pvs *PlaylistVideoSelect) BoolsX(ctx context.Context) []bool {
	v, err := pvs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (pvs *PlaylistVideoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pvs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{playlistvideo.Label}
	default:
		err = fmt.Errorf("ent: PlaylistVideoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pvs *PlaylistVideoSelect) BoolX(ctx context.Context) bool {
	v, err := pvs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pvs *PlaylistVideoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pvs.sqlQuery().Query()
	if err := pvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pvs *PlaylistVideoSelect) sqlQuery() sql.Querier {
	selector := pvs.sql
	selector.Select(selector.Columns(pvs.fields...)...)
	return selector
}
