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

// Campus is an object representing the database table.
type Campus struct {
	CampusID   int    `boil:"campus_id" json:"campusID" toml:"campusID" yaml:"campusID"`
	CampusName string `boil:"campus_name" json:"campusName" toml:"campusName" yaml:"campusName"`

	R *campusR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L campusL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CampusColumns = struct {
	CampusID   string
	CampusName string
}{
	CampusID:   "campus_id",
	CampusName: "campus_name",
}

var CampusTableColumns = struct {
	CampusID   string
	CampusName string
}{
	CampusID:   "campus.campus_id",
	CampusName: "campus.campus_name",
}

// Generated where

var CampusWhere = struct {
	CampusID   whereHelperint
	CampusName whereHelperstring
}{
	CampusID:   whereHelperint{field: "\"campus\".\"campus_id\""},
	CampusName: whereHelperstring{field: "\"campus\".\"campus_name\""},
}

// CampusRels is where relationship names are stored.
var CampusRels = struct {
	ReadingRooms string
	Restaurants  string
}{
	ReadingRooms: "ReadingRooms",
	Restaurants:  "Restaurants",
}

// campusR is where relationships are stored.
type campusR struct {
	ReadingRooms ReadingRoomSlice `boil:"ReadingRooms" json:"ReadingRooms" toml:"ReadingRooms" yaml:"ReadingRooms"`
	Restaurants  RestaurantSlice  `boil:"Restaurants" json:"Restaurants" toml:"Restaurants" yaml:"Restaurants"`
}

// NewStruct creates a new relationship struct
func (*campusR) NewStruct() *campusR {
	return &campusR{}
}

func (r *campusR) GetReadingRooms() ReadingRoomSlice {
	if r == nil {
		return nil
	}
	return r.ReadingRooms
}

func (r *campusR) GetRestaurants() RestaurantSlice {
	if r == nil {
		return nil
	}
	return r.Restaurants
}

// campusL is where Load methods for each relationship are stored.
type campusL struct{}

var (
	campusAllColumns            = []string{"campus_id", "campus_name"}
	campusColumnsWithoutDefault = []string{"campus_id", "campus_name"}
	campusColumnsWithDefault    = []string{}
	campusPrimaryKeyColumns     = []string{"campus_id"}
	campusGeneratedColumns      = []string{}
)

type (
	// CampusSlice is an alias for a slice of pointers to Campus.
	// This should almost always be used instead of []Campus.
	CampusSlice []*Campus

	campusQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	campusType                 = reflect.TypeOf(&Campus{})
	campusMapping              = queries.MakeStructMapping(campusType)
	campusPrimaryKeyMapping, _ = queries.BindMapping(campusType, campusMapping, campusPrimaryKeyColumns)
	campusInsertCacheMut       sync.RWMutex
	campusInsertCache          = make(map[string]insertCache)
	campusUpdateCacheMut       sync.RWMutex
	campusUpdateCache          = make(map[string]updateCache)
	campusUpsertCacheMut       sync.RWMutex
	campusUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single campus record from the query.
func (q campusQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Campus, error) {
	o := &Campus{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for campus")
	}

	return o, nil
}

// All returns all Campus records from the query.
func (q campusQuery) All(ctx context.Context, exec boil.ContextExecutor) (CampusSlice, error) {
	var o []*Campus

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Campus slice")
	}

	return o, nil
}

// Count returns the count of all Campus records in the query.
func (q campusQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count campus rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q campusQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if campus exists")
	}

	return count > 0, nil
}

// ReadingRooms retrieves all the reading_room's ReadingRooms with an executor.
func (o *Campus) ReadingRooms(mods ...qm.QueryMod) readingRoomQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"reading_room\".\"campus_id\"=?", o.CampusID),
	)

	return ReadingRooms(queryMods...)
}

// Restaurants retrieves all the restaurant's Restaurants with an executor.
func (o *Campus) Restaurants(mods ...qm.QueryMod) restaurantQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"restaurant\".\"campus_id\"=?", o.CampusID),
	)

	return Restaurants(queryMods...)
}

// LoadReadingRooms allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (campusL) LoadReadingRooms(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCampus interface{}, mods queries.Applicator) error {
	var slice []*Campus
	var object *Campus

	if singular {
		var ok bool
		object, ok = maybeCampus.(*Campus)
		if !ok {
			object = new(Campus)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCampus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCampus))
			}
		}
	} else {
		s, ok := maybeCampus.(*[]*Campus)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCampus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCampus))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &campusR{}
		}
		args = append(args, object.CampusID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &campusR{}
			}

			for _, a := range args {
				if a == obj.CampusID {
					continue Outer
				}
			}

			args = append(args, obj.CampusID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`reading_room`),
		qm.WhereIn(`reading_room.campus_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load reading_room")
	}

	var resultSlice []*ReadingRoom
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice reading_room")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on reading_room")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for reading_room")
	}

	if singular {
		object.R.ReadingRooms = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &readingRoomR{}
			}
			foreign.R.Campus = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CampusID == foreign.CampusID {
				local.R.ReadingRooms = append(local.R.ReadingRooms, foreign)
				if foreign.R == nil {
					foreign.R = &readingRoomR{}
				}
				foreign.R.Campus = local
				break
			}
		}
	}

	return nil
}

// LoadRestaurants allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (campusL) LoadRestaurants(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCampus interface{}, mods queries.Applicator) error {
	var slice []*Campus
	var object *Campus

	if singular {
		var ok bool
		object, ok = maybeCampus.(*Campus)
		if !ok {
			object = new(Campus)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCampus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCampus))
			}
		}
	} else {
		s, ok := maybeCampus.(*[]*Campus)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCampus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCampus))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &campusR{}
		}
		args = append(args, object.CampusID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &campusR{}
			}

			for _, a := range args {
				if a == obj.CampusID {
					continue Outer
				}
			}

			args = append(args, obj.CampusID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`restaurant`),
		qm.WhereIn(`restaurant.campus_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load restaurant")
	}

	var resultSlice []*Restaurant
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice restaurant")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on restaurant")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for restaurant")
	}

	if singular {
		object.R.Restaurants = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &restaurantR{}
			}
			foreign.R.Campus = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CampusID == foreign.CampusID {
				local.R.Restaurants = append(local.R.Restaurants, foreign)
				if foreign.R == nil {
					foreign.R = &restaurantR{}
				}
				foreign.R.Campus = local
				break
			}
		}
	}

	return nil
}

// AddReadingRooms adds the given related objects to the existing relationships
// of the campus, optionally inserting them as new records.
// Appends related to o.R.ReadingRooms.
// Sets related.R.Campus appropriately.
func (o *Campus) AddReadingRooms(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ReadingRoom) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CampusID = o.CampusID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"reading_room\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"campus_id"}),
				strmangle.WhereClause("\"", "\"", 2, readingRoomPrimaryKeyColumns),
			)
			values := []interface{}{o.CampusID, rel.RoomID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CampusID = o.CampusID
		}
	}

	if o.R == nil {
		o.R = &campusR{
			ReadingRooms: related,
		}
	} else {
		o.R.ReadingRooms = append(o.R.ReadingRooms, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &readingRoomR{
				Campus: o,
			}
		} else {
			rel.R.Campus = o
		}
	}
	return nil
}

// AddRestaurants adds the given related objects to the existing relationships
// of the campus, optionally inserting them as new records.
// Appends related to o.R.Restaurants.
// Sets related.R.Campus appropriately.
func (o *Campus) AddRestaurants(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Restaurant) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CampusID = o.CampusID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"restaurant\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"campus_id"}),
				strmangle.WhereClause("\"", "\"", 2, restaurantPrimaryKeyColumns),
			)
			values := []interface{}{o.CampusID, rel.RestaurantID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CampusID = o.CampusID
		}
	}

	if o.R == nil {
		o.R = &campusR{
			Restaurants: related,
		}
	} else {
		o.R.Restaurants = append(o.R.Restaurants, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &restaurantR{
				Campus: o,
			}
		} else {
			rel.R.Campus = o
		}
	}
	return nil
}

// Campuses retrieves all the records using an executor.
func Campuses(mods ...qm.QueryMod) campusQuery {
	mods = append(mods, qm.From("\"campus\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"campus\".*"})
	}

	return campusQuery{q}
}

// FindCampus retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCampus(ctx context.Context, exec boil.ContextExecutor, campusID int, selectCols ...string) (*Campus, error) {
	campusObj := &Campus{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"campus\" where \"campus_id\"=$1", sel,
	)

	q := queries.Raw(query, campusID)

	err := q.Bind(ctx, exec, campusObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from campus")
	}

	return campusObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Campus) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no campus provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(campusColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	campusInsertCacheMut.RLock()
	cache, cached := campusInsertCache[key]
	campusInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			campusAllColumns,
			campusColumnsWithDefault,
			campusColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(campusType, campusMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(campusType, campusMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"campus\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"campus\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into campus")
	}

	if !cached {
		campusInsertCacheMut.Lock()
		campusInsertCache[key] = cache
		campusInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Campus.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Campus) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	campusUpdateCacheMut.RLock()
	cache, cached := campusUpdateCache[key]
	campusUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			campusAllColumns,
			campusPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update campus, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"campus\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, campusPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(campusType, campusMapping, append(wl, campusPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update campus row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for campus")
	}

	if !cached {
		campusUpdateCacheMut.Lock()
		campusUpdateCache[key] = cache
		campusUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q campusQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for campus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for campus")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CampusSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), campusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"campus\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, campusPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in campus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all campus")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Campus) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no campus provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(campusColumnsWithDefault, o)

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

	campusUpsertCacheMut.RLock()
	cache, cached := campusUpsertCache[key]
	campusUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			campusAllColumns,
			campusColumnsWithDefault,
			campusColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			campusAllColumns,
			campusPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert campus, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(campusPrimaryKeyColumns))
			copy(conflict, campusPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"campus\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(campusType, campusMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(campusType, campusMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert campus")
	}

	if !cached {
		campusUpsertCacheMut.Lock()
		campusUpsertCache[key] = cache
		campusUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Campus record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Campus) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Campus provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), campusPrimaryKeyMapping)
	sql := "DELETE FROM \"campus\" WHERE \"campus_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from campus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for campus")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q campusQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no campusQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from campus")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for campus")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CampusSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), campusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"campus\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, campusPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from campus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for campus")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Campus) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCampus(ctx, exec, o.CampusID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CampusSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CampusSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), campusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"campus\".* FROM \"campus\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, campusPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CampusSlice")
	}

	*o = slice

	return nil
}

// CampusExists checks if the Campus row exists.
func CampusExists(ctx context.Context, exec boil.ContextExecutor, campusID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"campus\" where \"campus_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, campusID)
	}
	row := exec.QueryRowContext(ctx, sql, campusID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if campus exists")
	}

	return exists, nil
}

// Exists checks if the Campus row exists.
func (o *Campus) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return CampusExists(ctx, exec, o.CampusID)
}