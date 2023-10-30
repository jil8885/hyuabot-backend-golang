// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CommuteShuttleTimetable is an object representing the database table.
type CommuteShuttleTimetable struct {
	RouteName     string    `boil:"route_name" json:"routeName" toml:"routeName" yaml:"routeName"`
	StopName      string    `boil:"stop_name" json:"stopName" toml:"stopName" yaml:"stopName"`
	StopOrder     null.Int  `boil:"stop_order" json:"stopOrder,omitempty" toml:"stopOrder" yaml:"stopOrder,omitempty"`
	DepartureTime time.Time `boil:"departure_time" json:"departureTime" toml:"departureTime" yaml:"departureTime"`

	R *commuteShuttleTimetableR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L commuteShuttleTimetableL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CommuteShuttleTimetableColumns = struct {
	RouteName     string
	StopName      string
	StopOrder     string
	DepartureTime string
}{
	RouteName:     "route_name",
	StopName:      "stop_name",
	StopOrder:     "stop_order",
	DepartureTime: "departure_time",
}

var CommuteShuttleTimetableTableColumns = struct {
	RouteName     string
	StopName      string
	StopOrder     string
	DepartureTime string
}{
	RouteName:     "commute_shuttle_timetable.route_name",
	StopName:      "commute_shuttle_timetable.stop_name",
	StopOrder:     "commute_shuttle_timetable.stop_order",
	DepartureTime: "commute_shuttle_timetable.departure_time",
}

// Generated where

var CommuteShuttleTimetableWhere = struct {
	RouteName     whereHelperstring
	StopName      whereHelperstring
	StopOrder     whereHelpernull_Int
	DepartureTime whereHelpertime_Time
}{
	RouteName:     whereHelperstring{field: "\"commute_shuttle_timetable\".\"route_name\""},
	StopName:      whereHelperstring{field: "\"commute_shuttle_timetable\".\"stop_name\""},
	StopOrder:     whereHelpernull_Int{field: "\"commute_shuttle_timetable\".\"stop_order\""},
	DepartureTime: whereHelpertime_Time{field: "\"commute_shuttle_timetable\".\"departure_time\""},
}

// CommuteShuttleTimetableRels is where relationship names are stored.
var CommuteShuttleTimetableRels = struct {
	RouteNameCommuteShuttleRoute string
	StopNameCommuteShuttleStop   string
}{
	RouteNameCommuteShuttleRoute: "RouteNameCommuteShuttleRoute",
	StopNameCommuteShuttleStop:   "StopNameCommuteShuttleStop",
}

// commuteShuttleTimetableR is where relationships are stored.
type commuteShuttleTimetableR struct {
	RouteNameCommuteShuttleRoute *CommuteShuttleRoute `boil:"RouteNameCommuteShuttleRoute" json:"RouteNameCommuteShuttleRoute" toml:"RouteNameCommuteShuttleRoute" yaml:"RouteNameCommuteShuttleRoute"`
	StopNameCommuteShuttleStop   *CommuteShuttleStop  `boil:"StopNameCommuteShuttleStop" json:"StopNameCommuteShuttleStop" toml:"StopNameCommuteShuttleStop" yaml:"StopNameCommuteShuttleStop"`
}

// NewStruct creates a new relationship struct
func (*commuteShuttleTimetableR) NewStruct() *commuteShuttleTimetableR {
	return &commuteShuttleTimetableR{}
}

func (r *commuteShuttleTimetableR) GetRouteNameCommuteShuttleRoute() *CommuteShuttleRoute {
	if r == nil {
		return nil
	}
	return r.RouteNameCommuteShuttleRoute
}

func (r *commuteShuttleTimetableR) GetStopNameCommuteShuttleStop() *CommuteShuttleStop {
	if r == nil {
		return nil
	}
	return r.StopNameCommuteShuttleStop
}

// commuteShuttleTimetableL is where Load methods for each relationship are stored.
type commuteShuttleTimetableL struct{}

var (
	commuteShuttleTimetableAllColumns            = []string{"route_name", "stop_name", "stop_order", "departure_time"}
	commuteShuttleTimetableColumnsWithoutDefault = []string{"route_name", "stop_name", "departure_time"}
	commuteShuttleTimetableColumnsWithDefault    = []string{"stop_order"}
	commuteShuttleTimetablePrimaryKeyColumns     = []string{"route_name", "stop_name"}
	commuteShuttleTimetableGeneratedColumns      = []string{}
)

type (
	// CommuteShuttleTimetableSlice is an alias for a slice of pointers to CommuteShuttleTimetable.
	// This should almost always be used instead of []CommuteShuttleTimetable.
	CommuteShuttleTimetableSlice []*CommuteShuttleTimetable

	commuteShuttleTimetableQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	commuteShuttleTimetableType                 = reflect.TypeOf(&CommuteShuttleTimetable{})
	commuteShuttleTimetableMapping              = queries.MakeStructMapping(commuteShuttleTimetableType)
	commuteShuttleTimetablePrimaryKeyMapping, _ = queries.BindMapping(commuteShuttleTimetableType, commuteShuttleTimetableMapping, commuteShuttleTimetablePrimaryKeyColumns)
	commuteShuttleTimetableInsertCacheMut       sync.RWMutex
	commuteShuttleTimetableInsertCache          = make(map[string]insertCache)
	commuteShuttleTimetableUpdateCacheMut       sync.RWMutex
	commuteShuttleTimetableUpdateCache          = make(map[string]updateCache)
	commuteShuttleTimetableUpsertCacheMut       sync.RWMutex
	commuteShuttleTimetableUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single commuteShuttleTimetable record from the query.
func (q commuteShuttleTimetableQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CommuteShuttleTimetable, error) {
	o := &CommuteShuttleTimetable{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for commute_shuttle_timetable")
	}

	return o, nil
}

// All returns all CommuteShuttleTimetable records from the query.
func (q commuteShuttleTimetableQuery) All(ctx context.Context, exec boil.ContextExecutor) (CommuteShuttleTimetableSlice, error) {
	var o []*CommuteShuttleTimetable

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CommuteShuttleTimetable slice")
	}

	return o, nil
}

// Count returns the count of all CommuteShuttleTimetable records in the query.
func (q commuteShuttleTimetableQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count commute_shuttle_timetable rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q commuteShuttleTimetableQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if commute_shuttle_timetable exists")
	}

	return count > 0, nil
}

// RouteNameCommuteShuttleRoute pointed to by the foreign key.
func (o *CommuteShuttleTimetable) RouteNameCommuteShuttleRoute(mods ...qm.QueryMod) commuteShuttleRouteQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"route_name\" = ?", o.RouteName),
	}

	queryMods = append(queryMods, mods...)

	return CommuteShuttleRoutes(queryMods...)
}

// StopNameCommuteShuttleStop pointed to by the foreign key.
func (o *CommuteShuttleTimetable) StopNameCommuteShuttleStop(mods ...qm.QueryMod) commuteShuttleStopQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"stop_name\" = ?", o.StopName),
	}

	queryMods = append(queryMods, mods...)

	return CommuteShuttleStops(queryMods...)
}

// LoadRouteNameCommuteShuttleRoute allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (commuteShuttleTimetableL) LoadRouteNameCommuteShuttleRoute(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCommuteShuttleTimetable interface{}, mods queries.Applicator) error {
	var slice []*CommuteShuttleTimetable
	var object *CommuteShuttleTimetable

	if singular {
		var ok bool
		object, ok = maybeCommuteShuttleTimetable.(*CommuteShuttleTimetable)
		if !ok {
			object = new(CommuteShuttleTimetable)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCommuteShuttleTimetable)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCommuteShuttleTimetable))
			}
		}
	} else {
		s, ok := maybeCommuteShuttleTimetable.(*[]*CommuteShuttleTimetable)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCommuteShuttleTimetable)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCommuteShuttleTimetable))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &commuteShuttleTimetableR{}
		}
		args = append(args, object.RouteName)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &commuteShuttleTimetableR{}
			}

			for _, a := range args {
				if a == obj.RouteName {
					continue Outer
				}
			}

			args = append(args, obj.RouteName)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`commute_shuttle_route`),
		qm.WhereIn(`commute_shuttle_route.route_name in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CommuteShuttleRoute")
	}

	var resultSlice []*CommuteShuttleRoute
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CommuteShuttleRoute")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for commute_shuttle_route")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for commute_shuttle_route")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.RouteNameCommuteShuttleRoute = foreign
		if foreign.R == nil {
			foreign.R = &commuteShuttleRouteR{}
		}
		foreign.R.RouteNameCommuteShuttleTimetables = append(foreign.R.RouteNameCommuteShuttleTimetables, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RouteName == foreign.RouteName {
				local.R.RouteNameCommuteShuttleRoute = foreign
				if foreign.R == nil {
					foreign.R = &commuteShuttleRouteR{}
				}
				foreign.R.RouteNameCommuteShuttleTimetables = append(foreign.R.RouteNameCommuteShuttleTimetables, local)
				break
			}
		}
	}

	return nil
}

// LoadStopNameCommuteShuttleStop allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (commuteShuttleTimetableL) LoadStopNameCommuteShuttleStop(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCommuteShuttleTimetable interface{}, mods queries.Applicator) error {
	var slice []*CommuteShuttleTimetable
	var object *CommuteShuttleTimetable

	if singular {
		var ok bool
		object, ok = maybeCommuteShuttleTimetable.(*CommuteShuttleTimetable)
		if !ok {
			object = new(CommuteShuttleTimetable)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCommuteShuttleTimetable)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCommuteShuttleTimetable))
			}
		}
	} else {
		s, ok := maybeCommuteShuttleTimetable.(*[]*CommuteShuttleTimetable)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCommuteShuttleTimetable)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCommuteShuttleTimetable))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &commuteShuttleTimetableR{}
		}
		args = append(args, object.StopName)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &commuteShuttleTimetableR{}
			}

			for _, a := range args {
				if a == obj.StopName {
					continue Outer
				}
			}

			args = append(args, obj.StopName)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`commute_shuttle_stop`),
		qm.WhereIn(`commute_shuttle_stop.stop_name in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load CommuteShuttleStop")
	}

	var resultSlice []*CommuteShuttleStop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice CommuteShuttleStop")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for commute_shuttle_stop")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for commute_shuttle_stop")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.StopNameCommuteShuttleStop = foreign
		if foreign.R == nil {
			foreign.R = &commuteShuttleStopR{}
		}
		foreign.R.StopNameCommuteShuttleTimetables = append(foreign.R.StopNameCommuteShuttleTimetables, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StopName == foreign.StopName {
				local.R.StopNameCommuteShuttleStop = foreign
				if foreign.R == nil {
					foreign.R = &commuteShuttleStopR{}
				}
				foreign.R.StopNameCommuteShuttleTimetables = append(foreign.R.StopNameCommuteShuttleTimetables, local)
				break
			}
		}
	}

	return nil
}

// SetRouteNameCommuteShuttleRoute of the commuteShuttleTimetable to the related item.
// Sets o.R.RouteNameCommuteShuttleRoute to related.
// Adds o to related.R.RouteNameCommuteShuttleTimetables.
func (o *CommuteShuttleTimetable) SetRouteNameCommuteShuttleRoute(ctx context.Context, exec boil.ContextExecutor, insert bool, related *CommuteShuttleRoute) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"commute_shuttle_timetable\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"route_name"}),
		strmangle.WhereClause("\"", "\"", 2, commuteShuttleTimetablePrimaryKeyColumns),
	)
	values := []interface{}{related.RouteName, o.RouteName, o.StopName}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RouteName = related.RouteName
	if o.R == nil {
		o.R = &commuteShuttleTimetableR{
			RouteNameCommuteShuttleRoute: related,
		}
	} else {
		o.R.RouteNameCommuteShuttleRoute = related
	}

	if related.R == nil {
		related.R = &commuteShuttleRouteR{
			RouteNameCommuteShuttleTimetables: CommuteShuttleTimetableSlice{o},
		}
	} else {
		related.R.RouteNameCommuteShuttleTimetables = append(related.R.RouteNameCommuteShuttleTimetables, o)
	}

	return nil
}

// SetStopNameCommuteShuttleStop of the commuteShuttleTimetable to the related item.
// Sets o.R.StopNameCommuteShuttleStop to related.
// Adds o to related.R.StopNameCommuteShuttleTimetables.
func (o *CommuteShuttleTimetable) SetStopNameCommuteShuttleStop(ctx context.Context, exec boil.ContextExecutor, insert bool, related *CommuteShuttleStop) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"commute_shuttle_timetable\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"stop_name"}),
		strmangle.WhereClause("\"", "\"", 2, commuteShuttleTimetablePrimaryKeyColumns),
	)
	values := []interface{}{related.StopName, o.RouteName, o.StopName}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StopName = related.StopName
	if o.R == nil {
		o.R = &commuteShuttleTimetableR{
			StopNameCommuteShuttleStop: related,
		}
	} else {
		o.R.StopNameCommuteShuttleStop = related
	}

	if related.R == nil {
		related.R = &commuteShuttleStopR{
			StopNameCommuteShuttleTimetables: CommuteShuttleTimetableSlice{o},
		}
	} else {
		related.R.StopNameCommuteShuttleTimetables = append(related.R.StopNameCommuteShuttleTimetables, o)
	}

	return nil
}

// CommuteShuttleTimetables retrieves all the records using an executor.
func CommuteShuttleTimetables(mods ...qm.QueryMod) commuteShuttleTimetableQuery {
	mods = append(mods, qm.From("\"commute_shuttle_timetable\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"commute_shuttle_timetable\".*"})
	}

	return commuteShuttleTimetableQuery{q}
}

// FindCommuteShuttleTimetable retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCommuteShuttleTimetable(ctx context.Context, exec boil.ContextExecutor, routeName string, stopName string, selectCols ...string) (*CommuteShuttleTimetable, error) {
	commuteShuttleTimetableObj := &CommuteShuttleTimetable{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"commute_shuttle_timetable\" where \"route_name\"=$1 AND \"stop_name\"=$2", sel,
	)

	q := queries.Raw(query, routeName, stopName)

	err := q.Bind(ctx, exec, commuteShuttleTimetableObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from commute_shuttle_timetable")
	}

	return commuteShuttleTimetableObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CommuteShuttleTimetable) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no commute_shuttle_timetable provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(commuteShuttleTimetableColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	commuteShuttleTimetableInsertCacheMut.RLock()
	cache, cached := commuteShuttleTimetableInsertCache[key]
	commuteShuttleTimetableInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			commuteShuttleTimetableAllColumns,
			commuteShuttleTimetableColumnsWithDefault,
			commuteShuttleTimetableColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(commuteShuttleTimetableType, commuteShuttleTimetableMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(commuteShuttleTimetableType, commuteShuttleTimetableMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"commute_shuttle_timetable\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"commute_shuttle_timetable\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into commute_shuttle_timetable")
	}

	if !cached {
		commuteShuttleTimetableInsertCacheMut.Lock()
		commuteShuttleTimetableInsertCache[key] = cache
		commuteShuttleTimetableInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CommuteShuttleTimetable.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CommuteShuttleTimetable) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	commuteShuttleTimetableUpdateCacheMut.RLock()
	cache, cached := commuteShuttleTimetableUpdateCache[key]
	commuteShuttleTimetableUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			commuteShuttleTimetableAllColumns,
			commuteShuttleTimetablePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update commute_shuttle_timetable, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"commute_shuttle_timetable\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, commuteShuttleTimetablePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(commuteShuttleTimetableType, commuteShuttleTimetableMapping, append(wl, commuteShuttleTimetablePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update commute_shuttle_timetable row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for commute_shuttle_timetable")
	}

	if !cached {
		commuteShuttleTimetableUpdateCacheMut.Lock()
		commuteShuttleTimetableUpdateCache[key] = cache
		commuteShuttleTimetableUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q commuteShuttleTimetableQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for commute_shuttle_timetable")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for commute_shuttle_timetable")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CommuteShuttleTimetableSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commuteShuttleTimetablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"commute_shuttle_timetable\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, commuteShuttleTimetablePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in commuteShuttleTimetable slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all commuteShuttleTimetable")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CommuteShuttleTimetable) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no commute_shuttle_timetable provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(commuteShuttleTimetableColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	commuteShuttleTimetableUpsertCacheMut.RLock()
	cache, cached := commuteShuttleTimetableUpsertCache[key]
	commuteShuttleTimetableUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			commuteShuttleTimetableAllColumns,
			commuteShuttleTimetableColumnsWithDefault,
			commuteShuttleTimetableColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			commuteShuttleTimetableAllColumns,
			commuteShuttleTimetablePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert commute_shuttle_timetable, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(commuteShuttleTimetablePrimaryKeyColumns))
			copy(conflict, commuteShuttleTimetablePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"commute_shuttle_timetable\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(commuteShuttleTimetableType, commuteShuttleTimetableMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(commuteShuttleTimetableType, commuteShuttleTimetableMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert commute_shuttle_timetable")
	}

	if !cached {
		commuteShuttleTimetableUpsertCacheMut.Lock()
		commuteShuttleTimetableUpsertCache[key] = cache
		commuteShuttleTimetableUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CommuteShuttleTimetable record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CommuteShuttleTimetable) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CommuteShuttleTimetable provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), commuteShuttleTimetablePrimaryKeyMapping)
	sql := "DELETE FROM \"commute_shuttle_timetable\" WHERE \"route_name\"=$1 AND \"stop_name\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from commute_shuttle_timetable")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for commute_shuttle_timetable")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q commuteShuttleTimetableQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no commuteShuttleTimetableQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from commute_shuttle_timetable")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for commute_shuttle_timetable")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CommuteShuttleTimetableSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commuteShuttleTimetablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"commute_shuttle_timetable\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commuteShuttleTimetablePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from commuteShuttleTimetable slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for commute_shuttle_timetable")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CommuteShuttleTimetable) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCommuteShuttleTimetable(ctx, exec, o.RouteName, o.StopName)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommuteShuttleTimetableSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CommuteShuttleTimetableSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commuteShuttleTimetablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"commute_shuttle_timetable\".* FROM \"commute_shuttle_timetable\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commuteShuttleTimetablePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CommuteShuttleTimetableSlice")
	}

	*o = slice

	return nil
}

// CommuteShuttleTimetableExists checks if the CommuteShuttleTimetable row exists.
func CommuteShuttleTimetableExists(ctx context.Context, exec boil.ContextExecutor, routeName string, stopName string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"commute_shuttle_timetable\" where \"route_name\"=$1 AND \"stop_name\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, routeName, stopName)
	}
	row := exec.QueryRowContext(ctx, sql, routeName, stopName)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if commute_shuttle_timetable exists")
	}

	return exists, nil
}

// Exists checks if the CommuteShuttleTimetable row exists.
func (o *CommuteShuttleTimetable) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CommuteShuttleTimetableExists(ctx, exec, o.RouteName, o.StopName)
}
