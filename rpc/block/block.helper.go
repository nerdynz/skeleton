package block

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

var blockHelperGlobal *BlockHelper

func (h *BlockHelper) beforeSave(record *Block) (err error) {
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

func (h *BlockHelper) afterSave(record *Block) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type BlockHelper struct {
	datastore  *datastore.Datastore
	fieldNames []string
	orderBy    string
}

func NewBlockHelper(datastore *datastore.Datastore) *BlockHelper {
	if blockHelperGlobal == nil {
		blockHelperGlobal = newBlockHelper(datastore)
	}
	return blockHelperGlobal
}

func newBlockHelper(datastore *datastore.Datastore) *BlockHelper {
	helper := &BlockHelper{}
	helper.datastore = datastore

	// Fields
	fieldnames := []string{"block_ulid", "title", "kind", "content_one_html", "content_two_html", "content_three_html", "content_four_html", "image_one_ulid", "image_two_ulid", "image_three_ulid", "image_four_ulid", "date_created", "date_modified", "page_ulid", "site_ulid", "sort_position"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames

	helper.orderBy = "sort_position, date_created, date_modified"
	return helper
}

func (h *BlockHelper) New(siteUlid string) *Block {
	record := &Block{}
	// check DateCreated
	record.DateCreated = time.Now().Format(time.RFC3339)
	record.SiteUlid = siteUlid
	return record
}

func (h *BlockHelper) FromRequest(siteUlid string, req *http.Request) (*Block, error) {
	record := h.New(siteUlid)

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(record)
	if err != nil {
		return nil, err
	}

	if record.SiteUlid != siteUlid {
		return nil, errors.New("Block update failed. Site ULID Mismatch")
	}
	record.SiteUlid = siteUlid
	return record, nil
}

func (h *BlockHelper) Load(siteUlid string, ulid string) (*Block, error) {
	record, err := h.One(siteUlid, "block_ulid = $1 and $SITEULID", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *BlockHelper) All(siteUlid string) ([]*Block, error) {
	var records []*Block
	err := h.datastore.DB.Select(h.fieldNames...).
		From("block").
		Where("site_ulid = $1", siteUlid).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *BlockHelper) Where(siteUlid string, sql string, args ...interface{}) ([]*Block, error) {
	var records []*Block
	sql, args, err := datastore.AppendSiteULID(siteUlid, sql, args...)
	if err != nil {
		return nil, err
	}
	err = h.datastore.DB.Select(h.fieldNames...).
		From("block").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *BlockHelper) SQL(siteUlid string, sql string, args ...interface{}) ([]*Block, error) {
	if !strings.Contains(sql, "$SITEULID") {
		return nil, errors.New("No $SITEULID placeholder defined")
	}
	var records []*Block
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

func (h *BlockHelper) One(siteUlid string, sql string, args ...interface{}) (*Block, error) {
	var record Block
	sql, args, err := datastore.AppendSiteULID(siteUlid, sql, args...)

	if err != nil {
		return nil, err
	}
	err = h.datastore.DB.Select(h.fieldNames...).
		From("block").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *BlockHelper) Paged(siteUlid string, pageNum int32, itemsPerPage int32) (*BlocksPaged, error) {
	pd, err := h.PagedBy(siteUlid, pageNum, itemsPerPage, "date_created", "", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *BlockHelper) PagedBy(siteUlid string, pageNum int32, itemsPerPage int32, orderByFieldName string, direction string, searchText string) (*BlocksPaged, error) {
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

	var records []*Block
	err := h.datastore.DB.Select(h.fieldNames...).
		From("block").
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
	h.datastore.DB.SQL(`select count(block_ulid) from block where site_ulid = $1`, siteUlid).QueryStruct(&count)
	paged := &base.PagedInfo{
		PageNumber: pageNum,
		OrderBy:    orderByFieldName,
		Direction:  dirEnum,
		Limit:      itemsPerPage,
		Total:      count,
		Search:     searchText,
	}
	return &BlocksPaged{
		PagedInfo: paged,
		Records:   records,
	}, nil
}

func (h *BlockHelper) Save(siteUlid string, record *Block) error {
	return h.save(siteUlid, record)
}

func (h *BlockHelper) SaveMany(siteUlid string, records []*Block) error {
	for _, record := range records {
		err := h.save(siteUlid, record)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *BlockHelper) save(siteUlid string, record *Block) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteUlid != siteUlid {
		return errors.New("*Block update failed. SITEULID Mismatch")
	}
	cols := []string{"title", "kind", "content_one_html", "content_two_html", "content_three_html", "content_four_html", "image_one_ulid", "image_two_ulid", "image_three_ulid", "image_four_ulid", "date_created", "date_modified", "page_ulid", "site_ulid", "sort_position"}
	vals := []interface{}{record.Title, record.Kind, record.ContentOneHtml, record.ContentTwoHtml, record.ContentThreeHtml, record.ContentFourHtml, record.ImageOneUlid, record.ImageTwoUlid, record.ImageThreeUlid, record.ImageFourUlid, record.DateCreated, record.DateModified, record.PageUlid, record.SiteUlid, record.SortPosition}
	err = h.datastore.DB.Upsert("block").
		Columns(cols...).
		Values(vals...).
		Where("block_ulid = $1", record.BlockUlid).
		Returning("block_ulid").
		QueryStruct(record)

	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *BlockHelper) Delete(siteUlid string, recordULID string) (bool, error) {
	result, err := h.datastore.DB.
		Update("block").
		Set("date_deleted", time.Now().Format(time.RFC3339)).
		Where("site_ulid=$1 and block_ulid=$2", siteUlid, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *BlockHelper) Purge(siteUlid string, recordULID string) (bool, error) {
	result, err := h.datastore.DB.
		DeleteFrom("block").
		Where("site_ulid=$1 and block_ulid=$2", siteUlid, recordULID).
		Exec()
	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *BlockHelper) validate(record *Block) (err error) {
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
