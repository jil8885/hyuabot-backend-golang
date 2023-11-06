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

// NoticeCategory is an object representing the database table.
type NoticeCategory struct {
	CategoryID   int    `boil:"category_id" json:"categoryID" toml:"categoryID" yaml:"categoryID"`
	CategoryName string `boil:"category_name" json:"categoryName" toml:"categoryName" yaml:"categoryName"`

	R *noticeCategoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L noticeCategoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NoticeCategoryColumns = struct {
	CategoryID   string
	CategoryName string
}{
	CategoryID:   "category_id",
	CategoryName: "category_name",
}

var NoticeCategoryTableColumns = struct {
	CategoryID   string
	CategoryName string
}{
	CategoryID:   "notice_category.category_id",
	CategoryName: "notice_category.category_name",
}

// Generated where

var NoticeCategoryWhere = struct {
	CategoryID   whereHelperint
	CategoryName whereHelperstring
}{
	CategoryID:   whereHelperint{field: "\"notice_category\".\"category_id\""},
	CategoryName: whereHelperstring{field: "\"notice_category\".\"category_name\""},
}

// NoticeCategoryRels is where relationship names are stored.
var NoticeCategoryRels = struct {
	CategoryNotices string
}{
	CategoryNotices: "CategoryNotices",
}

// noticeCategoryR is where relationships are stored.
type noticeCategoryR struct {
	CategoryNotices NoticeSlice `boil:"CategoryNotices" json:"CategoryNotices" toml:"CategoryNotices" yaml:"CategoryNotices"`
}

// NewStruct creates a new relationship struct
func (*noticeCategoryR) NewStruct() *noticeCategoryR {
	return &noticeCategoryR{}
}

func (r *noticeCategoryR) GetCategoryNotices() NoticeSlice {
	if r == nil {
		return nil
	}
	return r.CategoryNotices
}

// noticeCategoryL is where Load methods for each relationship are stored.
type noticeCategoryL struct{}

var (
	noticeCategoryAllColumns            = []string{"category_id", "category_name"}
	noticeCategoryColumnsWithoutDefault = []string{"category_id", "category_name"}
	noticeCategoryColumnsWithDefault    = []string{}
	noticeCategoryPrimaryKeyColumns     = []string{"category_id"}
	noticeCategoryGeneratedColumns      = []string{}
)

type (
	// NoticeCategorySlice is an alias for a slice of pointers to NoticeCategory.
	// This should almost always be used instead of []NoticeCategory.
	NoticeCategorySlice []*NoticeCategory

	noticeCategoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	noticeCategoryType                 = reflect.TypeOf(&NoticeCategory{})
	noticeCategoryMapping              = queries.MakeStructMapping(noticeCategoryType)
	noticeCategoryPrimaryKeyMapping, _ = queries.BindMapping(noticeCategoryType, noticeCategoryMapping, noticeCategoryPrimaryKeyColumns)
	noticeCategoryInsertCacheMut       sync.RWMutex
	noticeCategoryInsertCache          = make(map[string]insertCache)
	noticeCategoryUpdateCacheMut       sync.RWMutex
	noticeCategoryUpdateCache          = make(map[string]updateCache)
	noticeCategoryUpsertCacheMut       sync.RWMutex
	noticeCategoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single noticeCategory record from the query.
func (q noticeCategoryQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NoticeCategory, error) {
	o := &NoticeCategory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for notice_category")
	}

	return o, nil
}

// All returns all NoticeCategory records from the query.
func (q noticeCategoryQuery) All(ctx context.Context, exec boil.ContextExecutor) (NoticeCategorySlice, error) {
	var o []*NoticeCategory

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NoticeCategory slice")
	}

	return o, nil
}

// Count returns the count of all NoticeCategory records in the query.
func (q noticeCategoryQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count notice_category rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q noticeCategoryQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if notice_category exists")
	}

	return count > 0, nil
}

// CategoryNotices retrieves all the notice's Notices with an executor via category_id column.
func (o *NoticeCategory) CategoryNotices(mods ...qm.QueryMod) noticeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"notices\".\"category_id\"=?", o.CategoryID),
	)

	return Notices(queryMods...)
}

// LoadCategoryNotices allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (noticeCategoryL) LoadCategoryNotices(ctx context.Context, e boil.ContextExecutor, singular bool, maybeNoticeCategory interface{}, mods queries.Applicator) error {
	var slice []*NoticeCategory
	var object *NoticeCategory

	if singular {
		var ok bool
		object, ok = maybeNoticeCategory.(*NoticeCategory)
		if !ok {
			object = new(NoticeCategory)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeNoticeCategory)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeNoticeCategory))
			}
		}
	} else {
		s, ok := maybeNoticeCategory.(*[]*NoticeCategory)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeNoticeCategory)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeNoticeCategory))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &noticeCategoryR{}
		}
		args = append(args, object.CategoryID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &noticeCategoryR{}
			}

			for _, a := range args {
				if a == obj.CategoryID {
					continue Outer
				}
			}

			args = append(args, obj.CategoryID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`notices`),
		qm.WhereIn(`notices.category_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load notices")
	}

	var resultSlice []*Notice
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice notices")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on notices")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for notices")
	}

	if singular {
		object.R.CategoryNotices = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &noticeR{}
			}
			foreign.R.Category = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.CategoryID == foreign.CategoryID {
				local.R.CategoryNotices = append(local.R.CategoryNotices, foreign)
				if foreign.R == nil {
					foreign.R = &noticeR{}
				}
				foreign.R.Category = local
				break
			}
		}
	}

	return nil
}

// AddCategoryNotices adds the given related objects to the existing relationships
// of the notice_category, optionally inserting them as new records.
// Appends related to o.R.CategoryNotices.
// Sets related.R.Category appropriately.
func (o *NoticeCategory) AddCategoryNotices(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Notice) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CategoryID = o.CategoryID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"notices\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"category_id"}),
				strmangle.WhereClause("\"", "\"", 2, noticePrimaryKeyColumns),
			)
			values := []interface{}{o.CategoryID, rel.NoticeID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CategoryID = o.CategoryID
		}
	}

	if o.R == nil {
		o.R = &noticeCategoryR{
			CategoryNotices: related,
		}
	} else {
		o.R.CategoryNotices = append(o.R.CategoryNotices, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &noticeR{
				Category: o,
			}
		} else {
			rel.R.Category = o
		}
	}
	return nil
}

// NoticeCategories retrieves all the records using an executor.
func NoticeCategories(mods ...qm.QueryMod) noticeCategoryQuery {
	mods = append(mods, qm.From("\"notice_category\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"notice_category\".*"})
	}

	return noticeCategoryQuery{q}
}

// FindNoticeCategory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNoticeCategory(ctx context.Context, exec boil.ContextExecutor, categoryID int, selectCols ...string) (*NoticeCategory, error) {
	noticeCategoryObj := &NoticeCategory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"notice_category\" where \"category_id\"=$1", sel,
	)

	q := queries.Raw(query, categoryID)

	err := q.Bind(ctx, exec, noticeCategoryObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from notice_category")
	}

	return noticeCategoryObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NoticeCategory) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no notice_category provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(noticeCategoryColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	noticeCategoryInsertCacheMut.RLock()
	cache, cached := noticeCategoryInsertCache[key]
	noticeCategoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			noticeCategoryAllColumns,
			noticeCategoryColumnsWithDefault,
			noticeCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(noticeCategoryType, noticeCategoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(noticeCategoryType, noticeCategoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"notice_category\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"notice_category\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into notice_category")
	}

	if !cached {
		noticeCategoryInsertCacheMut.Lock()
		noticeCategoryInsertCache[key] = cache
		noticeCategoryInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the NoticeCategory.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NoticeCategory) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	noticeCategoryUpdateCacheMut.RLock()
	cache, cached := noticeCategoryUpdateCache[key]
	noticeCategoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			noticeCategoryAllColumns,
			noticeCategoryPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update notice_category, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"notice_category\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, noticeCategoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(noticeCategoryType, noticeCategoryMapping, append(wl, noticeCategoryPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update notice_category row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for notice_category")
	}

	if !cached {
		noticeCategoryUpdateCacheMut.Lock()
		noticeCategoryUpdateCache[key] = cache
		noticeCategoryUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q noticeCategoryQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for notice_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for notice_category")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NoticeCategorySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), noticeCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"notice_category\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, noticeCategoryPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in noticeCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all noticeCategory")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *NoticeCategory) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no notice_category provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(noticeCategoryColumnsWithDefault, o)

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

	noticeCategoryUpsertCacheMut.RLock()
	cache, cached := noticeCategoryUpsertCache[key]
	noticeCategoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			noticeCategoryAllColumns,
			noticeCategoryColumnsWithDefault,
			noticeCategoryColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			noticeCategoryAllColumns,
			noticeCategoryPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert notice_category, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(noticeCategoryPrimaryKeyColumns))
			copy(conflict, noticeCategoryPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"notice_category\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(noticeCategoryType, noticeCategoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(noticeCategoryType, noticeCategoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert notice_category")
	}

	if !cached {
		noticeCategoryUpsertCacheMut.Lock()
		noticeCategoryUpsertCache[key] = cache
		noticeCategoryUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single NoticeCategory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NoticeCategory) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NoticeCategory provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), noticeCategoryPrimaryKeyMapping)
	sql := "DELETE FROM \"notice_category\" WHERE \"category_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from notice_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for notice_category")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q noticeCategoryQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no noticeCategoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from notice_category")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for notice_category")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NoticeCategorySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), noticeCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"notice_category\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, noticeCategoryPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from noticeCategory slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for notice_category")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *NoticeCategory) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNoticeCategory(ctx, exec, o.CategoryID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NoticeCategorySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NoticeCategorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), noticeCategoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"notice_category\".* FROM \"notice_category\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, noticeCategoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NoticeCategorySlice")
	}

	*o = slice

	return nil
}

// NoticeCategoryExists checks if the NoticeCategory row exists.
func NoticeCategoryExists(ctx context.Context, exec boil.ContextExecutor, categoryID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"notice_category\" where \"category_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, categoryID)
	}
	row := exec.QueryRowContext(ctx, sql, categoryID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if notice_category exists")
	}

	return exists, nil
}

// Exists checks if the NoticeCategory row exists.
func (o *NoticeCategory) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return NoticeCategoryExists(ctx, exec, o.CategoryID)
}