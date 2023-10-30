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

// CommuteShuttleRoute is an object representing the database table.
type CommuteShuttleRoute struct {
	RouteName               string      `boil:"route_name" json:"routeName" toml:"routeName" yaml:"routeName"`
	RouteDescriptionKorean  null.String `boil:"route_description_korean" json:"routeDescriptionKorean,omitempty" toml:"routeDescriptionKorean" yaml:"routeDescriptionKorean,omitempty"`
	RouteDescriptionEnglish null.String `boil:"route_description_english" json:"routeDescriptionEnglish,omitempty" toml:"routeDescriptionEnglish" yaml:"routeDescriptionEnglish,omitempty"`

	R *commuteShuttleRouteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L commuteShuttleRouteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CommuteShuttleRouteColumns = struct {
	RouteName               string
	RouteDescriptionKorean  string
	RouteDescriptionEnglish string
}{
	RouteName:               "route_name",
	RouteDescriptionKorean:  "route_description_korean",
	RouteDescriptionEnglish: "route_description_english",
}

var CommuteShuttleRouteTableColumns = struct {
	RouteName               string
	RouteDescriptionKorean  string
	RouteDescriptionEnglish string
}{
	RouteName:               "commute_shuttle_route.route_name",
	RouteDescriptionKorean:  "commute_shuttle_route.route_description_korean",
	RouteDescriptionEnglish: "commute_shuttle_route.route_description_english",
}

// Generated where

var CommuteShuttleRouteWhere = struct {
	RouteName               whereHelperstring
	RouteDescriptionKorean  whereHelpernull_String
	RouteDescriptionEnglish whereHelpernull_String
}{
	RouteName:               whereHelperstring{field: "\"commute_shuttle_route\".\"route_name\""},
	RouteDescriptionKorean:  whereHelpernull_String{field: "\"commute_shuttle_route\".\"route_description_korean\""},
	RouteDescriptionEnglish: whereHelpernull_String{field: "\"commute_shuttle_route\".\"route_description_english\""},
}

// CommuteShuttleRouteRels is where relationship names are stored.
var CommuteShuttleRouteRels = struct {
	RouteNameCommuteShuttleTimetables string
}{
	RouteNameCommuteShuttleTimetables: "RouteNameCommuteShuttleTimetables",
}

// commuteShuttleRouteR is where relationships are stored.
type commuteShuttleRouteR struct {
	RouteNameCommuteShuttleTimetables CommuteShuttleTimetableSlice `boil:"RouteNameCommuteShuttleTimetables" json:"RouteNameCommuteShuttleTimetables" toml:"RouteNameCommuteShuttleTimetables" yaml:"RouteNameCommuteShuttleTimetables"`
}

// NewStruct creates a new relationship struct
func (*commuteShuttleRouteR) NewStruct() *commuteShuttleRouteR {
	return &commuteShuttleRouteR{}
}

func (r *commuteShuttleRouteR) GetRouteNameCommuteShuttleTimetables() CommuteShuttleTimetableSlice {
	if r == nil {
		return nil
	}
	return r.RouteNameCommuteShuttleTimetables
}

// commuteShuttleRouteL is where Load methods for each relationship are stored.
type commuteShuttleRouteL struct{}

var (
	commuteShuttleRouteAllColumns            = []string{"route_name", "route_description_korean", "route_description_english"}
	commuteShuttleRouteColumnsWithoutDefault = []string{"route_name"}
	commuteShuttleRouteColumnsWithDefault    = []string{"route_description_korean", "route_description_english"}
	commuteShuttleRoutePrimaryKeyColumns     = []string{"route_name"}
	commuteShuttleRouteGeneratedColumns      = []string{}
)

type (
	// CommuteShuttleRouteSlice is an alias for a slice of pointers to CommuteShuttleRoute.
	// This should almost always be used instead of []CommuteShuttleRoute.
	CommuteShuttleRouteSlice []*CommuteShuttleRoute

	commuteShuttleRouteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	commuteShuttleRouteType                 = reflect.TypeOf(&CommuteShuttleRoute{})
	commuteShuttleRouteMapping              = queries.MakeStructMapping(commuteShuttleRouteType)
	commuteShuttleRoutePrimaryKeyMapping, _ = queries.BindMapping(commuteShuttleRouteType, commuteShuttleRouteMapping, commuteShuttleRoutePrimaryKeyColumns)
	commuteShuttleRouteInsertCacheMut       sync.RWMutex
	commuteShuttleRouteInsertCache          = make(map[string]insertCache)
	commuteShuttleRouteUpdateCacheMut       sync.RWMutex
	commuteShuttleRouteUpdateCache          = make(map[string]updateCache)
	commuteShuttleRouteUpsertCacheMut       sync.RWMutex
	commuteShuttleRouteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single commuteShuttleRoute record from the query.
func (q commuteShuttleRouteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CommuteShuttleRoute, error) {
	o := &CommuteShuttleRoute{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for commute_shuttle_route")
	}

	return o, nil
}

// All returns all CommuteShuttleRoute records from the query.
func (q commuteShuttleRouteQuery) All(ctx context.Context, exec boil.ContextExecutor) (CommuteShuttleRouteSlice, error) {
	var o []*CommuteShuttleRoute

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CommuteShuttleRoute slice")
	}

	return o, nil
}

// Count returns the count of all CommuteShuttleRoute records in the query.
func (q commuteShuttleRouteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count commute_shuttle_route rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q commuteShuttleRouteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if commute_shuttle_route exists")
	}

	return count > 0, nil
}

// RouteNameCommuteShuttleTimetables retrieves all the commute_shuttle_timetable's CommuteShuttleTimetables with an executor via route_name column.
func (o *CommuteShuttleRoute) RouteNameCommuteShuttleTimetables(mods ...qm.QueryMod) commuteShuttleTimetableQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"commute_shuttle_timetable\".\"route_name\"=?", o.RouteName),
	)

	return CommuteShuttleTimetables(queryMods...)
}

// LoadRouteNameCommuteShuttleTimetables allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (commuteShuttleRouteL) LoadRouteNameCommuteShuttleTimetables(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCommuteShuttleRoute interface{}, mods queries.Applicator) error {
	var slice []*CommuteShuttleRoute
	var object *CommuteShuttleRoute

	if singular {
		var ok bool
		object, ok = maybeCommuteShuttleRoute.(*CommuteShuttleRoute)
		if !ok {
			object = new(CommuteShuttleRoute)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCommuteShuttleRoute)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCommuteShuttleRoute))
			}
		}
	} else {
		s, ok := maybeCommuteShuttleRoute.(*[]*CommuteShuttleRoute)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCommuteShuttleRoute)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCommuteShuttleRoute))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &commuteShuttleRouteR{}
		}
		args = append(args, object.RouteName)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &commuteShuttleRouteR{}
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
		qm.From(`commute_shuttle_timetable`),
		qm.WhereIn(`commute_shuttle_timetable.route_name in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load commute_shuttle_timetable")
	}

	var resultSlice []*CommuteShuttleTimetable
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice commute_shuttle_timetable")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on commute_shuttle_timetable")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for commute_shuttle_timetable")
	}

	if singular {
		object.R.RouteNameCommuteShuttleTimetables = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &commuteShuttleTimetableR{}
			}
			foreign.R.RouteNameCommuteShuttleRoute = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.RouteName == foreign.RouteName {
				local.R.RouteNameCommuteShuttleTimetables = append(local.R.RouteNameCommuteShuttleTimetables, foreign)
				if foreign.R == nil {
					foreign.R = &commuteShuttleTimetableR{}
				}
				foreign.R.RouteNameCommuteShuttleRoute = local
				break
			}
		}
	}

	return nil
}

// AddRouteNameCommuteShuttleTimetables adds the given related objects to the existing relationships
// of the commute_shuttle_route, optionally inserting them as new records.
// Appends related to o.R.RouteNameCommuteShuttleTimetables.
// Sets related.R.RouteNameCommuteShuttleRoute appropriately.
func (o *CommuteShuttleRoute) AddRouteNameCommuteShuttleTimetables(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*CommuteShuttleTimetable) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RouteName = o.RouteName
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"commute_shuttle_timetable\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"route_name"}),
				strmangle.WhereClause("\"", "\"", 2, commuteShuttleTimetablePrimaryKeyColumns),
			)
			values := []interface{}{o.RouteName, rel.RouteName, rel.StopName}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.RouteName = o.RouteName
		}
	}

	if o.R == nil {
		o.R = &commuteShuttleRouteR{
			RouteNameCommuteShuttleTimetables: related,
		}
	} else {
		o.R.RouteNameCommuteShuttleTimetables = append(o.R.RouteNameCommuteShuttleTimetables, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &commuteShuttleTimetableR{
				RouteNameCommuteShuttleRoute: o,
			}
		} else {
			rel.R.RouteNameCommuteShuttleRoute = o
		}
	}
	return nil
}

// CommuteShuttleRoutes retrieves all the records using an executor.
func CommuteShuttleRoutes(mods ...qm.QueryMod) commuteShuttleRouteQuery {
	mods = append(mods, qm.From("\"commute_shuttle_route\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"commute_shuttle_route\".*"})
	}

	return commuteShuttleRouteQuery{q}
}

// FindCommuteShuttleRoute retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCommuteShuttleRoute(ctx context.Context, exec boil.ContextExecutor, routeName string, selectCols ...string) (*CommuteShuttleRoute, error) {
	commuteShuttleRouteObj := &CommuteShuttleRoute{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"commute_shuttle_route\" where \"route_name\"=$1", sel,
	)

	q := queries.Raw(query, routeName)

	err := q.Bind(ctx, exec, commuteShuttleRouteObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from commute_shuttle_route")
	}

	return commuteShuttleRouteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CommuteShuttleRoute) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no commute_shuttle_route provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(commuteShuttleRouteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	commuteShuttleRouteInsertCacheMut.RLock()
	cache, cached := commuteShuttleRouteInsertCache[key]
	commuteShuttleRouteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			commuteShuttleRouteAllColumns,
			commuteShuttleRouteColumnsWithDefault,
			commuteShuttleRouteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(commuteShuttleRouteType, commuteShuttleRouteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(commuteShuttleRouteType, commuteShuttleRouteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"commute_shuttle_route\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"commute_shuttle_route\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into commute_shuttle_route")
	}

	if !cached {
		commuteShuttleRouteInsertCacheMut.Lock()
		commuteShuttleRouteInsertCache[key] = cache
		commuteShuttleRouteInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the CommuteShuttleRoute.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CommuteShuttleRoute) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	commuteShuttleRouteUpdateCacheMut.RLock()
	cache, cached := commuteShuttleRouteUpdateCache[key]
	commuteShuttleRouteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			commuteShuttleRouteAllColumns,
			commuteShuttleRoutePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update commute_shuttle_route, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"commute_shuttle_route\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, commuteShuttleRoutePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(commuteShuttleRouteType, commuteShuttleRouteMapping, append(wl, commuteShuttleRoutePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update commute_shuttle_route row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for commute_shuttle_route")
	}

	if !cached {
		commuteShuttleRouteUpdateCacheMut.Lock()
		commuteShuttleRouteUpdateCache[key] = cache
		commuteShuttleRouteUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q commuteShuttleRouteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for commute_shuttle_route")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for commute_shuttle_route")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CommuteShuttleRouteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commuteShuttleRoutePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"commute_shuttle_route\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, commuteShuttleRoutePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in commuteShuttleRoute slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all commuteShuttleRoute")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CommuteShuttleRoute) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no commute_shuttle_route provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(commuteShuttleRouteColumnsWithDefault, o)

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

	commuteShuttleRouteUpsertCacheMut.RLock()
	cache, cached := commuteShuttleRouteUpsertCache[key]
	commuteShuttleRouteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			commuteShuttleRouteAllColumns,
			commuteShuttleRouteColumnsWithDefault,
			commuteShuttleRouteColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			commuteShuttleRouteAllColumns,
			commuteShuttleRoutePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert commute_shuttle_route, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(commuteShuttleRoutePrimaryKeyColumns))
			copy(conflict, commuteShuttleRoutePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"commute_shuttle_route\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(commuteShuttleRouteType, commuteShuttleRouteMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(commuteShuttleRouteType, commuteShuttleRouteMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert commute_shuttle_route")
	}

	if !cached {
		commuteShuttleRouteUpsertCacheMut.Lock()
		commuteShuttleRouteUpsertCache[key] = cache
		commuteShuttleRouteUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single CommuteShuttleRoute record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CommuteShuttleRoute) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CommuteShuttleRoute provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), commuteShuttleRoutePrimaryKeyMapping)
	sql := "DELETE FROM \"commute_shuttle_route\" WHERE \"route_name\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from commute_shuttle_route")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for commute_shuttle_route")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q commuteShuttleRouteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no commuteShuttleRouteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from commute_shuttle_route")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for commute_shuttle_route")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CommuteShuttleRouteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commuteShuttleRoutePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"commute_shuttle_route\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commuteShuttleRoutePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from commuteShuttleRoute slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for commute_shuttle_route")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CommuteShuttleRoute) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCommuteShuttleRoute(ctx, exec, o.RouteName)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommuteShuttleRouteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CommuteShuttleRouteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commuteShuttleRoutePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"commute_shuttle_route\".* FROM \"commute_shuttle_route\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commuteShuttleRoutePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CommuteShuttleRouteSlice")
	}

	*o = slice

	return nil
}

// CommuteShuttleRouteExists checks if the CommuteShuttleRoute row exists.
func CommuteShuttleRouteExists(ctx context.Context, exec boil.ContextExecutor, routeName string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"commute_shuttle_route\" where \"route_name\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, routeName)
	}
	row := exec.QueryRowContext(ctx, sql, routeName)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if commute_shuttle_route exists")
	}

	return exists, nil
}

// Exists checks if the CommuteShuttleRoute row exists.
func (o *CommuteShuttleRoute) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CommuteShuttleRouteExists(ctx, exec, o.RouteName)
}
