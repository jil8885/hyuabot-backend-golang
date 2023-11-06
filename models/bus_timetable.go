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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// BusTimetable is an object representing the database table.
type BusTimetable struct {
	RouteID       int       `boil:"route_id" json:"routeID" toml:"routeID" yaml:"routeID"`
	StartStopID   int       `boil:"start_stop_id" json:"startStopID" toml:"startStopID" yaml:"startStopID"`
	DepartureTime time.Time `boil:"departure_time" json:"departureTime" toml:"departureTime" yaml:"departureTime"`
	Weekday       string    `boil:"weekday" json:"weekday" toml:"weekday" yaml:"weekday"`

	R *busTimetableR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L busTimetableL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BusTimetableColumns = struct {
	RouteID       string
	StartStopID   string
	DepartureTime string
	Weekday       string
}{
	RouteID:       "route_id",
	StartStopID:   "start_stop_id",
	DepartureTime: "departure_time",
	Weekday:       "weekday",
}

var BusTimetableTableColumns = struct {
	RouteID       string
	StartStopID   string
	DepartureTime string
	Weekday       string
}{
	RouteID:       "bus_timetable.route_id",
	StartStopID:   "bus_timetable.start_stop_id",
	DepartureTime: "bus_timetable.departure_time",
	Weekday:       "bus_timetable.weekday",
}

// Generated where

var BusTimetableWhere = struct {
	RouteID       whereHelperint
	StartStopID   whereHelperint
	DepartureTime whereHelpertime_Time
	Weekday       whereHelperstring
}{
	RouteID:       whereHelperint{field: "\"bus_timetable\".\"route_id\""},
	StartStopID:   whereHelperint{field: "\"bus_timetable\".\"start_stop_id\""},
	DepartureTime: whereHelpertime_Time{field: "\"bus_timetable\".\"departure_time\""},
	Weekday:       whereHelperstring{field: "\"bus_timetable\".\"weekday\""},
}

// BusTimetableRels is where relationship names are stored.
var BusTimetableRels = struct {
	Route     string
	StartStop string
}{
	Route:     "Route",
	StartStop: "StartStop",
}

// busTimetableR is where relationships are stored.
type busTimetableR struct {
	Route     *BusRoute `boil:"Route" json:"Route" toml:"Route" yaml:"Route"`
	StartStop *BusStop  `boil:"StartStop" json:"StartStop" toml:"StartStop" yaml:"StartStop"`
}

// NewStruct creates a new relationship struct
func (*busTimetableR) NewStruct() *busTimetableR {
	return &busTimetableR{}
}

func (r *busTimetableR) GetRoute() *BusRoute {
	if r == nil {
		return nil
	}
	return r.Route
}

func (r *busTimetableR) GetStartStop() *BusStop {
	if r == nil {
		return nil
	}
	return r.StartStop
}

// busTimetableL is where Load methods for each relationship are stored.
type busTimetableL struct{}

var (
	busTimetableAllColumns            = []string{"route_id", "start_stop_id", "departure_time", "weekday"}
	busTimetableColumnsWithoutDefault = []string{"route_id", "start_stop_id", "departure_time", "weekday"}
	busTimetableColumnsWithDefault    = []string{}
	busTimetablePrimaryKeyColumns     = []string{"route_id", "start_stop_id", "departure_time", "weekday"}
	busTimetableGeneratedColumns      = []string{}
)

type (
	// BusTimetableSlice is an alias for a slice of pointers to BusTimetable.
	// This should almost always be used instead of []BusTimetable.
	BusTimetableSlice []*BusTimetable

	busTimetableQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	busTimetableType                 = reflect.TypeOf(&BusTimetable{})
	busTimetableMapping              = queries.MakeStructMapping(busTimetableType)
	busTimetablePrimaryKeyMapping, _ = queries.BindMapping(busTimetableType, busTimetableMapping, busTimetablePrimaryKeyColumns)
	busTimetableInsertCacheMut       sync.RWMutex
	busTimetableInsertCache          = make(map[string]insertCache)
	busTimetableUpdateCacheMut       sync.RWMutex
	busTimetableUpdateCache          = make(map[string]updateCache)
	busTimetableUpsertCacheMut       sync.RWMutex
	busTimetableUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single busTimetable record from the query.
func (q busTimetableQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BusTimetable, error) {
	o := &BusTimetable{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bus_timetable")
	}

	return o, nil
}

// All returns all BusTimetable records from the query.
func (q busTimetableQuery) All(ctx context.Context, exec boil.ContextExecutor) (BusTimetableSlice, error) {
	var o []*BusTimetable

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BusTimetable slice")
	}

	return o, nil
}

// Count returns the count of all BusTimetable records in the query.
func (q busTimetableQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bus_timetable rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q busTimetableQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bus_timetable exists")
	}

	return count > 0, nil
}

// Route pointed to by the foreign key.
func (o *BusTimetable) Route(mods ...qm.QueryMod) busRouteQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"route_id\" = ?", o.RouteID),
	}

	queryMods = append(queryMods, mods...)

	return BusRoutes(queryMods...)
}

// StartStop pointed to by the foreign key.
func (o *BusTimetable) StartStop(mods ...qm.QueryMod) busStopQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"stop_id\" = ?", o.StartStopID),
	}

	queryMods = append(queryMods, mods...)

	return BusStops(queryMods...)
}

// LoadRoute allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (busTimetableL) LoadRoute(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBusTimetable interface{}, mods queries.Applicator) error {
	var slice []*BusTimetable
	var object *BusTimetable

	if singular {
		var ok bool
		object, ok = maybeBusTimetable.(*BusTimetable)
		if !ok {
			object = new(BusTimetable)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBusTimetable)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBusTimetable))
			}
		}
	} else {
		s, ok := maybeBusTimetable.(*[]*BusTimetable)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBusTimetable)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBusTimetable))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &busTimetableR{}
		}
		args = append(args, object.RouteID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &busTimetableR{}
			}

			for _, a := range args {
				if a == obj.RouteID {
					continue Outer
				}
			}

			args = append(args, obj.RouteID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`bus_route`),
		qm.WhereIn(`bus_route.route_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load BusRoute")
	}

	var resultSlice []*BusRoute
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice BusRoute")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for bus_route")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for bus_route")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Route = foreign
		if foreign.R == nil {
			foreign.R = &busRouteR{}
		}
		foreign.R.RouteBusTimetables = append(foreign.R.RouteBusTimetables, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RouteID == foreign.RouteID {
				local.R.Route = foreign
				if foreign.R == nil {
					foreign.R = &busRouteR{}
				}
				foreign.R.RouteBusTimetables = append(foreign.R.RouteBusTimetables, local)
				break
			}
		}
	}

	return nil
}

// LoadStartStop allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (busTimetableL) LoadStartStop(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBusTimetable interface{}, mods queries.Applicator) error {
	var slice []*BusTimetable
	var object *BusTimetable

	if singular {
		var ok bool
		object, ok = maybeBusTimetable.(*BusTimetable)
		if !ok {
			object = new(BusTimetable)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBusTimetable)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBusTimetable))
			}
		}
	} else {
		s, ok := maybeBusTimetable.(*[]*BusTimetable)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBusTimetable)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBusTimetable))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &busTimetableR{}
		}
		args = append(args, object.StartStopID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &busTimetableR{}
			}

			for _, a := range args {
				if a == obj.StartStopID {
					continue Outer
				}
			}

			args = append(args, obj.StartStopID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`bus_stop`),
		qm.WhereIn(`bus_stop.stop_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load BusStop")
	}

	var resultSlice []*BusStop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice BusStop")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for bus_stop")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for bus_stop")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.StartStop = foreign
		if foreign.R == nil {
			foreign.R = &busStopR{}
		}
		foreign.R.StartStopBusTimetables = append(foreign.R.StartStopBusTimetables, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StartStopID == foreign.StopID {
				local.R.StartStop = foreign
				if foreign.R == nil {
					foreign.R = &busStopR{}
				}
				foreign.R.StartStopBusTimetables = append(foreign.R.StartStopBusTimetables, local)
				break
			}
		}
	}

	return nil
}

// SetRoute of the busTimetable to the related item.
// Sets o.R.Route to related.
// Adds o to related.R.RouteBusTimetables.
func (o *BusTimetable) SetRoute(ctx context.Context, exec boil.ContextExecutor, insert bool, related *BusRoute) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"bus_timetable\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"route_id"}),
		strmangle.WhereClause("\"", "\"", 2, busTimetablePrimaryKeyColumns),
	)
	values := []interface{}{related.RouteID, o.RouteID, o.StartStopID, o.DepartureTime, o.Weekday}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RouteID = related.RouteID
	if o.R == nil {
		o.R = &busTimetableR{
			Route: related,
		}
	} else {
		o.R.Route = related
	}

	if related.R == nil {
		related.R = &busRouteR{
			RouteBusTimetables: BusTimetableSlice{o},
		}
	} else {
		related.R.RouteBusTimetables = append(related.R.RouteBusTimetables, o)
	}

	return nil
}

// SetStartStop of the busTimetable to the related item.
// Sets o.R.StartStop to related.
// Adds o to related.R.StartStopBusTimetables.
func (o *BusTimetable) SetStartStop(ctx context.Context, exec boil.ContextExecutor, insert bool, related *BusStop) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"bus_timetable\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"start_stop_id"}),
		strmangle.WhereClause("\"", "\"", 2, busTimetablePrimaryKeyColumns),
	)
	values := []interface{}{related.StopID, o.RouteID, o.StartStopID, o.DepartureTime, o.Weekday}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StartStopID = related.StopID
	if o.R == nil {
		o.R = &busTimetableR{
			StartStop: related,
		}
	} else {
		o.R.StartStop = related
	}

	if related.R == nil {
		related.R = &busStopR{
			StartStopBusTimetables: BusTimetableSlice{o},
		}
	} else {
		related.R.StartStopBusTimetables = append(related.R.StartStopBusTimetables, o)
	}

	return nil
}

// BusTimetables retrieves all the records using an executor.
func BusTimetables(mods ...qm.QueryMod) busTimetableQuery {
	mods = append(mods, qm.From("\"bus_timetable\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"bus_timetable\".*"})
	}

	return busTimetableQuery{q}
}

// FindBusTimetable retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBusTimetable(ctx context.Context, exec boil.ContextExecutor, routeID int, startStopID int, departureTime time.Time, weekday string, selectCols ...string) (*BusTimetable, error) {
	busTimetableObj := &BusTimetable{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"bus_timetable\" where \"route_id\"=$1 AND \"start_stop_id\"=$2 AND \"departure_time\"=$3 AND \"weekday\"=$4", sel,
	)

	q := queries.Raw(query, routeID, startStopID, departureTime, weekday)

	err := q.Bind(ctx, exec, busTimetableObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bus_timetable")
	}

	return busTimetableObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BusTimetable) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bus_timetable provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(busTimetableColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	busTimetableInsertCacheMut.RLock()
	cache, cached := busTimetableInsertCache[key]
	busTimetableInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			busTimetableAllColumns,
			busTimetableColumnsWithDefault,
			busTimetableColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(busTimetableType, busTimetableMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(busTimetableType, busTimetableMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"bus_timetable\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"bus_timetable\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into bus_timetable")
	}

	if !cached {
		busTimetableInsertCacheMut.Lock()
		busTimetableInsertCache[key] = cache
		busTimetableInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the BusTimetable.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BusTimetable) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	busTimetableUpdateCacheMut.RLock()
	cache, cached := busTimetableUpdateCache[key]
	busTimetableUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			busTimetableAllColumns,
			busTimetablePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bus_timetable, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"bus_timetable\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, busTimetablePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(busTimetableType, busTimetableMapping, append(wl, busTimetablePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update bus_timetable row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bus_timetable")
	}

	if !cached {
		busTimetableUpdateCacheMut.Lock()
		busTimetableUpdateCache[key] = cache
		busTimetableUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q busTimetableQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bus_timetable")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bus_timetable")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BusTimetableSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), busTimetablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"bus_timetable\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, busTimetablePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in busTimetable slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all busTimetable")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BusTimetable) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bus_timetable provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(busTimetableColumnsWithDefault, o)

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

	busTimetableUpsertCacheMut.RLock()
	cache, cached := busTimetableUpsertCache[key]
	busTimetableUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			busTimetableAllColumns,
			busTimetableColumnsWithDefault,
			busTimetableColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			busTimetableAllColumns,
			busTimetablePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert bus_timetable, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(busTimetablePrimaryKeyColumns))
			copy(conflict, busTimetablePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"bus_timetable\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(busTimetableType, busTimetableMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(busTimetableType, busTimetableMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert bus_timetable")
	}

	if !cached {
		busTimetableUpsertCacheMut.Lock()
		busTimetableUpsertCache[key] = cache
		busTimetableUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single BusTimetable record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BusTimetable) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BusTimetable provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), busTimetablePrimaryKeyMapping)
	sql := "DELETE FROM \"bus_timetable\" WHERE \"route_id\"=$1 AND \"start_stop_id\"=$2 AND \"departure_time\"=$3 AND \"weekday\"=$4"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bus_timetable")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bus_timetable")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q busTimetableQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no busTimetableQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bus_timetable")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bus_timetable")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BusTimetableSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), busTimetablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"bus_timetable\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, busTimetablePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from busTimetable slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bus_timetable")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BusTimetable) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBusTimetable(ctx, exec, o.RouteID, o.StartStopID, o.DepartureTime, o.Weekday)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BusTimetableSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BusTimetableSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), busTimetablePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"bus_timetable\".* FROM \"bus_timetable\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, busTimetablePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BusTimetableSlice")
	}

	*o = slice

	return nil
}

// BusTimetableExists checks if the BusTimetable row exists.
func BusTimetableExists(ctx context.Context, exec boil.ContextExecutor, routeID int, startStopID int, departureTime time.Time, weekday string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"bus_timetable\" where \"route_id\"=$1 AND \"start_stop_id\"=$2 AND \"departure_time\"=$3 AND \"weekday\"=$4 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, routeID, startStopID, departureTime, weekday)
	}
	row := exec.QueryRowContext(ctx, sql, routeID, startStopID, departureTime, weekday)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bus_timetable exists")
	}

	return exists, nil
}

// Exists checks if the BusTimetable row exists.
func (o *BusTimetable) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BusTimetableExists(ctx, exec, o.RouteID, o.StartStopID, o.DepartureTime, o.Weekday)
}