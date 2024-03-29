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

// Menu is an object representing the database table.
type Menu struct {
	RestaurantID int       `boil:"restaurant_id" json:"restaurantID" toml:"restaurantID" yaml:"restaurantID"`
	FeedDate     time.Time `boil:"feed_date" json:"feedDate" toml:"feedDate" yaml:"feedDate"`
	TimeType     string    `boil:"time_type" json:"timeType" toml:"timeType" yaml:"timeType"`
	MenuFood     string    `boil:"menu_food" json:"menuFood" toml:"menuFood" yaml:"menuFood"`
	MenuPrice    string    `boil:"menu_price" json:"menuPrice" toml:"menuPrice" yaml:"menuPrice"`

	R *menuR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L menuL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MenuColumns = struct {
	RestaurantID string
	FeedDate     string
	TimeType     string
	MenuFood     string
	MenuPrice    string
}{
	RestaurantID: "restaurant_id",
	FeedDate:     "feed_date",
	TimeType:     "time_type",
	MenuFood:     "menu_food",
	MenuPrice:    "menu_price",
}

var MenuTableColumns = struct {
	RestaurantID string
	FeedDate     string
	TimeType     string
	MenuFood     string
	MenuPrice    string
}{
	RestaurantID: "menu.restaurant_id",
	FeedDate:     "menu.feed_date",
	TimeType:     "menu.time_type",
	MenuFood:     "menu.menu_food",
	MenuPrice:    "menu.menu_price",
}

// Generated where

var MenuWhere = struct {
	RestaurantID whereHelperint
	FeedDate     whereHelpertime_Time
	TimeType     whereHelperstring
	MenuFood     whereHelperstring
	MenuPrice    whereHelperstring
}{
	RestaurantID: whereHelperint{field: "\"menu\".\"restaurant_id\""},
	FeedDate:     whereHelpertime_Time{field: "\"menu\".\"feed_date\""},
	TimeType:     whereHelperstring{field: "\"menu\".\"time_type\""},
	MenuFood:     whereHelperstring{field: "\"menu\".\"menu_food\""},
	MenuPrice:    whereHelperstring{field: "\"menu\".\"menu_price\""},
}

// MenuRels is where relationship names are stored.
var MenuRels = struct {
	Restaurant string
}{
	Restaurant: "Restaurant",
}

// menuR is where relationships are stored.
type menuR struct {
	Restaurant *Restaurant `boil:"Restaurant" json:"Restaurant" toml:"Restaurant" yaml:"Restaurant"`
}

// NewStruct creates a new relationship struct
func (*menuR) NewStruct() *menuR {
	return &menuR{}
}

func (r *menuR) GetRestaurant() *Restaurant {
	if r == nil {
		return nil
	}
	return r.Restaurant
}

// menuL is where Load methods for each relationship are stored.
type menuL struct{}

var (
	menuAllColumns            = []string{"restaurant_id", "feed_date", "time_type", "menu_food", "menu_price"}
	menuColumnsWithoutDefault = []string{"restaurant_id", "feed_date", "time_type", "menu_food", "menu_price"}
	menuColumnsWithDefault    = []string{}
	menuPrimaryKeyColumns     = []string{"restaurant_id", "feed_date", "time_type", "menu_food"}
	menuGeneratedColumns      = []string{}
)

type (
	// MenuSlice is an alias for a slice of pointers to Menu.
	// This should almost always be used instead of []Menu.
	MenuSlice []*Menu

	menuQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	menuType                 = reflect.TypeOf(&Menu{})
	menuMapping              = queries.MakeStructMapping(menuType)
	menuPrimaryKeyMapping, _ = queries.BindMapping(menuType, menuMapping, menuPrimaryKeyColumns)
	menuInsertCacheMut       sync.RWMutex
	menuInsertCache          = make(map[string]insertCache)
	menuUpdateCacheMut       sync.RWMutex
	menuUpdateCache          = make(map[string]updateCache)
	menuUpsertCacheMut       sync.RWMutex
	menuUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single menu record from the query.
func (q menuQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Menu, error) {
	o := &Menu{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for menu")
	}

	return o, nil
}

// All returns all Menu records from the query.
func (q menuQuery) All(ctx context.Context, exec boil.ContextExecutor) (MenuSlice, error) {
	var o []*Menu

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Menu slice")
	}

	return o, nil
}

// Count returns the count of all Menu records in the query.
func (q menuQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count menu rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q menuQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if menu exists")
	}

	return count > 0, nil
}

// Restaurant pointed to by the foreign key.
func (o *Menu) Restaurant(mods ...qm.QueryMod) restaurantQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"restaurant_id\" = ?", o.RestaurantID),
	}

	queryMods = append(queryMods, mods...)

	return Restaurants(queryMods...)
}

// LoadRestaurant allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (menuL) LoadRestaurant(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMenu interface{}, mods queries.Applicator) error {
	var slice []*Menu
	var object *Menu

	if singular {
		var ok bool
		object, ok = maybeMenu.(*Menu)
		if !ok {
			object = new(Menu)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMenu)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMenu))
			}
		}
	} else {
		s, ok := maybeMenu.(*[]*Menu)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMenu)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMenu))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &menuR{}
		}
		args = append(args, object.RestaurantID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &menuR{}
			}

			for _, a := range args {
				if a == obj.RestaurantID {
					continue Outer
				}
			}

			args = append(args, obj.RestaurantID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`restaurant`),
		qm.WhereIn(`restaurant.restaurant_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Restaurant")
	}

	var resultSlice []*Restaurant
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Restaurant")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for restaurant")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for restaurant")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Restaurant = foreign
		if foreign.R == nil {
			foreign.R = &restaurantR{}
		}
		foreign.R.Menus = append(foreign.R.Menus, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RestaurantID == foreign.RestaurantID {
				local.R.Restaurant = foreign
				if foreign.R == nil {
					foreign.R = &restaurantR{}
				}
				foreign.R.Menus = append(foreign.R.Menus, local)
				break
			}
		}
	}

	return nil
}

// SetRestaurant of the menu to the related item.
// Sets o.R.Restaurant to related.
// Adds o to related.R.Menus.
func (o *Menu) SetRestaurant(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Restaurant) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"menu\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"restaurant_id"}),
		strmangle.WhereClause("\"", "\"", 2, menuPrimaryKeyColumns),
	)
	values := []interface{}{related.RestaurantID, o.RestaurantID, o.FeedDate, o.TimeType, o.MenuFood}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RestaurantID = related.RestaurantID
	if o.R == nil {
		o.R = &menuR{
			Restaurant: related,
		}
	} else {
		o.R.Restaurant = related
	}

	if related.R == nil {
		related.R = &restaurantR{
			Menus: MenuSlice{o},
		}
	} else {
		related.R.Menus = append(related.R.Menus, o)
	}

	return nil
}

// Menus retrieves all the records using an executor.
func Menus(mods ...qm.QueryMod) menuQuery {
	mods = append(mods, qm.From("\"menu\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"menu\".*"})
	}

	return menuQuery{q}
}

// FindMenu retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMenu(ctx context.Context, exec boil.ContextExecutor, restaurantID int, feedDate time.Time, timeType string, menuFood string, selectCols ...string) (*Menu, error) {
	menuObj := &Menu{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"menu\" where \"restaurant_id\"=$1 AND \"feed_date\"=$2 AND \"time_type\"=$3 AND \"menu_food\"=$4", sel,
	)

	q := queries.Raw(query, restaurantID, feedDate, timeType, menuFood)

	err := q.Bind(ctx, exec, menuObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from menu")
	}

	return menuObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Menu) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no menu provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(menuColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	menuInsertCacheMut.RLock()
	cache, cached := menuInsertCache[key]
	menuInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			menuAllColumns,
			menuColumnsWithDefault,
			menuColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(menuType, menuMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(menuType, menuMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"menu\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"menu\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into menu")
	}

	if !cached {
		menuInsertCacheMut.Lock()
		menuInsertCache[key] = cache
		menuInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Menu.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Menu) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	menuUpdateCacheMut.RLock()
	cache, cached := menuUpdateCache[key]
	menuUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			menuAllColumns,
			menuPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update menu, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"menu\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, menuPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(menuType, menuMapping, append(wl, menuPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update menu row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for menu")
	}

	if !cached {
		menuUpdateCacheMut.Lock()
		menuUpdateCache[key] = cache
		menuUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q menuQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for menu")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for menu")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MenuSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"menu\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, menuPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in menu slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all menu")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Menu) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no menu provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(menuColumnsWithDefault, o)

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

	menuUpsertCacheMut.RLock()
	cache, cached := menuUpsertCache[key]
	menuUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			menuAllColumns,
			menuColumnsWithDefault,
			menuColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			menuAllColumns,
			menuPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert menu, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(menuPrimaryKeyColumns))
			copy(conflict, menuPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"menu\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(menuType, menuMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(menuType, menuMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert menu")
	}

	if !cached {
		menuUpsertCacheMut.Lock()
		menuUpsertCache[key] = cache
		menuUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Menu record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Menu) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Menu provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), menuPrimaryKeyMapping)
	sql := "DELETE FROM \"menu\" WHERE \"restaurant_id\"=$1 AND \"feed_date\"=$2 AND \"time_type\"=$3 AND \"menu_food\"=$4"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from menu")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for menu")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q menuQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no menuQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from menu")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for menu")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MenuSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"menu\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, menuPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from menu slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for menu")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Menu) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMenu(ctx, exec, o.RestaurantID, o.FeedDate, o.TimeType, o.MenuFood)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MenuSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MenuSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), menuPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"menu\".* FROM \"menu\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, menuPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MenuSlice")
	}

	*o = slice

	return nil
}

// MenuExists checks if the Menu row exists.
func MenuExists(ctx context.Context, exec boil.ContextExecutor, restaurantID int, feedDate time.Time, timeType string, menuFood string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"menu\" where \"restaurant_id\"=$1 AND \"feed_date\"=$2 AND \"time_type\"=$3 AND \"menu_food\"=$4 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, restaurantID, feedDate, timeType, menuFood)
	}
	row := exec.QueryRowContext(ctx, sql, restaurantID, feedDate, timeType, menuFood)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if menu exists")
	}

	return exists, nil
}

// Exists checks if the Menu row exists.
func (o *Menu) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return MenuExists(ctx, exec, o.RestaurantID, o.FeedDate, o.TimeType, o.MenuFood)
}
