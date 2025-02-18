package page

import (
	"encoding/json"
	"errors"
	"net/http"
	"sort"
	"strings"
	"time"

	//runner "github.com/nerdynz/dat/sqlx-runner"
	"github.com/nerdynz/datastore"
	"github.com/nerdynz/skeleton/rpc/base"
)

var pageHelperGlobal *PageHelper

func (h *PageHelper) beforeSave(record *Page) (err error) {
	if record.DateCreated == "" {
		record.DateCreated = time.Now().Format(time.RFC3339)
	}
	record.DateModified = time.Now().Format(time.RFC3339)

	validationErr := h.validate(record)
	if validationErr != nil {
		return validationErr
	}
	return err
}

func (h *PageHelper) afterSave(record *Page) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type PageHelper struct {
	datastore  *datastore.Datastore
	fieldNames []string
	orderBy    string
}

func NewPageHelper(datastore *datastore.Datastore) *PageHelper {
	if pageHelperGlobal == nil {
		pageHelperGlobal = newPageHelper(datastore)
	}
	return pageHelperGlobal
}

func newPageHelper(datastore *datastore.Datastore) *PageHelper {
	helper := &PageHelper{}
	helper.datastore = datastore

	// Fields
	fieldnames := []string{"page_ulid", "title", "summary", "date_created", "date_modified", "site_ulid", "slug"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames

	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *PageHelper) New(siteUlid string) *Page {
	record := &Page{}
	// check DateCreated
	record.DateCreated = time.Now().Format(time.RFC3339)
	record.SiteUlid = siteUlid
	return record
}

func (h *PageHelper) FromRequest(siteUlid string, req *http.Request) (*Page, error) {
	record := h.New(siteUlid)

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(record)
	if err != nil {
		return nil, err
	}

	if record.SiteUlid != siteUlid {
		return nil, errors.New("Page update failed. Site ULID Mismatch")
	}
	record.SiteUlid = siteUlid
	return record, nil
}

func (h *PageHelper) Load(siteUlid string, ulid string) (*Page, error) {
	record, err := h.One(siteUlid, "page_ulid = $1 and $SITEULID", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *PageHelper) All(siteUlid string) ([]*Page, error) {
	var records []*Page
	err := h.datastore.DB.Select(h.fieldNames...).
		From("page").
		Where("site_ulid = $1", siteUlid).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *PageHelper) Where(siteUlid string, sql string, args ...interface{}) ([]*Page, error) {
	var records []*Page
	sql, args, err := datastore.AppendSiteULID(siteUlid, sql, args...)
	if err != nil {
		return nil, err
	}
	err = h.datastore.DB.Select(h.fieldNames...).
		From("page").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *PageHelper) SQL(siteUlid string, sql string, args ...interface{}) ([]*Page, error) {
	var records []*Page
	sql, args, err := datastore.AppendSiteULID(siteUlid, sql, args...)
	if err != nil {
		return nil, err
	}
	err = h.datastore.DB.SQL(sql, args...).
		QueryStructs(&records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *PageHelper) One(siteUlid string, sql string, args ...interface{}) (*Page, error) {
	var record Page
	sql, args, err := datastore.AppendSiteULID(siteUlid, sql, args...)

	if err != nil {
		return nil, err
	}
	err = h.datastore.DB.Select(h.fieldNames...).
		From("page").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *PageHelper) Paged(siteUlid string, pageNum int32, itemsPerPage int32) (*PagesPaged, error) {
	pd, err := h.PagedBy(siteUlid, pageNum, itemsPerPage, "date_created", "", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *PageHelper) PagedBy(siteUlid string, pageNum int32, itemsPerPage int32, orderByFieldName string, direction string, searchText string) (*PagesPaged, error) {
	if orderByFieldName == "" || orderByFieldName == "default" {
		// we only want the first field name
		orderByFieldName = strings.Split(h.orderBy, ",")[0]
		orderByFieldName = strings.Trim(orderByFieldName, " ")
	}
	i := sort.SearchStrings(h.fieldNames, orderByFieldName)
	// check the orderby exists within the fields as this could be an easy sql injection hole.
	if !(i < len(h.fieldNames) && h.fieldNames[i] == orderByFieldName) { // NOT
		return nil, errors.New("field name [" + orderByFieldName + "]  isn't a valid field name")
	}

	direction = strings.ToLower(direction)
	if direction == "none" {
		direction = ""
	}

	if !(direction == "asc" || direction == "desc" || direction == "") {
		return nil, errors.New("direction isn't valid")
	}

	whereClause := "site_ulid = $1"
	if searchText != "" {
		whereClause += " and tsv @@ to_tsquery($2)"
	}

	var records []*Page
	err := h.datastore.DB.Select(h.fieldNames...).
		From("page").
		Where(whereClause, siteUlid, datastore.FormatSearch(searchText)).
		OrderBy(orderByFieldName + " " + direction).
		Offset(uint64((pageNum - 1) * itemsPerPage)).
		Limit(uint64(itemsPerPage)).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	dirEnum := base.PagedInfo_NONE
	switch strings.ToUpper(direction) {
	case base.PagedInfo_ASC.String():
		dirEnum = base.PagedInfo_ASC
	case base.PagedInfo_DESC.String():
		dirEnum = base.PagedInfo_DESC
	}

	count := int32(0)
	h.datastore.DB.SQL(`select count(page_ulid) from page where site_ulid = $1`, siteUlid).QueryStruct(&count)
	paged := &base.PagedInfo{
		PageNumber: pageNum,
		OrderBy:    orderByFieldName,
		Direction:  dirEnum,
		Limit:      itemsPerPage,
		Total:      count,
		Search:     searchText,
	}
	return &PagesPaged{
		PagedInfo: paged,
		Records:   records,
	}, nil
}

func (h *PageHelper) Save(siteUlid string, record *Page) error {
	return h.save(siteUlid, record)
}

func (h *PageHelper) SaveMany(siteUlid string, records []*Page) error {
	for _, record := range records {
		err := h.save(siteUlid, record)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *PageHelper) save(siteUlid string, record *Page) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteUlid != siteUlid {
		return errors.New("*Page update failed. SITEULID Mismatch")
	}
	cols := []string{"title", "summary", "date_created", "date_modified", "site_ulid", "slug"}
	vals := []interface{}{record.Title, record.Summary, record.DateCreated, record.DateModified, record.SiteUlid, record.Slug}
	err = h.datastore.DB.Upsert("page").
		Columns(cols...).
		Values(vals...).
		Where("page_ulid = $1", record.PageUlid).
		Returning("page_ulid").
		QueryStruct(record)

	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *PageHelper) Delete(siteUlid string, recordULID string) (bool, error) {
	result, err := h.datastore.DB.
		Update("page").
		Set("date_deleted", time.Now().Format(time.RFC3339)).
		Where("site_ulid=$1 and page_ulid=$2", siteUlid, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *PageHelper) Purge(siteUlid string, recordULID string) (bool, error) {
	result, err := h.datastore.DB.
		DeleteFrom("page").
		Where("site_ulid=$1 and page_ulid=$2", siteUlid, recordULID).
		Exec()
	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *PageHelper) validate(record *Page) (err error) {
	return nil
	//	validationErrors := h.validator.Struct(record)
	//	if validationErrors != nil {
	//		errMessage := ""
	//		for _, err := range err.(validator.ValidationErrors) {
	//			errMessage += err.Kind().String() + " validation Error on field "+err.Field()
	//		}
	//		if errMessage != "" {
	//			err = errors.New(errMessage)
	//		}
	//	}
	//	return err
}

//func (page *Page) SaveBlocks(siteUlid string) error {
//	return BlockHelper().SaveMany(siteUlid, page.Blocks)
//}
//
//func (page *Page) LoadBlocks(siteUlid string) error {
//	return page.LoadBlocksWhere(siteUlid, "page_ulid = $1 $SITEULID", page.PageULID)
//}
//
//func (page *Page) LoadBlocksWhere(siteUlid string, sql string, args ...interface{}) error {
//	children, err := BlockHelper().Where(siteUlid, sql, args...)
//	if err != nil {
//		return err
//	}
//	page.Blocks = children
//	return nil
//}
