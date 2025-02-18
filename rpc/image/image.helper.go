package image

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

var imageHelperGlobal *ImageHelper

func (h *ImageHelper) beforeSave(record *Image) (err error) {
	if record.DateCreated == "" {
		record.DateCreated = time.Now().Format(time.RFC3339)
	}
	record.DateModified = time.Now().Format(time.RFC3339)

	record.IsProcessed = false // just constantly reprocess the image if it changes

	validationErr := h.validate(record)
	if validationErr != nil {
		return validationErr
	}
	return err
}

func (h *ImageHelper) afterSave(record *Image) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type ImageHelper struct {
	datastore  *datastore.Datastore
	fieldNames []string
	orderBy    string
}

func NewImageHelper(datastore *datastore.Datastore) *ImageHelper {
	if imageHelperGlobal == nil {
		imageHelperGlobal = newImageHelper(datastore)
	}
	return imageHelperGlobal
}

func newImageHelper(datastore *datastore.Datastore) *ImageHelper {
	helper := &ImageHelper{}
	helper.datastore = datastore

	// Fields
	fieldnames := []string{"image_ulid", "image", "original_height", "original_width", "top", "left", "scale", "crop_height", "crop_width", "site_ulid", "date_created", "date_modified", "is_processed"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames

	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *ImageHelper) New(siteUlid string) *Image {
	record := &Image{}
	// check DateCreated
	record.DateCreated = time.Now().Format(time.RFC3339)
	record.SiteUlid = siteUlid
	return record
}

func (h *ImageHelper) FromRequest(siteUlid string, req *http.Request) (*Image, error) {
	record := h.New(siteUlid)

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(record)
	if err != nil {
		return nil, err
	}

	if record.SiteUlid != siteUlid {
		return nil, errors.New("Image update failed. Site ULID Mismatch")
	}
	record.SiteUlid = siteUlid
	return record, nil
}

func (h *ImageHelper) Load(siteUlid string, ulid string) (*Image, error) {
	record, err := h.One(siteUlid, "image_ulid = $1 and $SITEULID", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *ImageHelper) All(siteUlid string) ([]*Image, error) {
	var records []*Image
	err := h.datastore.DB.Select(h.fieldNames...).
		From("image").
		Where("site_ulid = $1", siteUlid).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *ImageHelper) Where(siteUlid string, sql string, args ...interface{}) ([]*Image, error) {
	var records []*Image
	sql, args, err := datastore.AppendSiteULID(siteUlid, sql, args...)
	if err != nil {
		return nil, err
	}

	fields := make([]string, 0)
	for _, f := range h.fieldNames {
		fields = append(fields, `"`+f+`"`)
	}

	err = h.datastore.DB.Select(fields...).
		From("image").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *ImageHelper) SQL(siteUlid string, sql string, args ...interface{}) ([]*Image, error) {
	if !strings.Contains(sql, "$SITEULID") {
		return nil, errors.New("No $SITEULID placeholder defined")
	}
	var records []*Image
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

func (h *ImageHelper) One(siteUlid string, sql string, args ...interface{}) (*Image, error) {
	var record Image
	sql, args, err := datastore.AppendSiteULID(siteUlid, sql, args...)

	if err != nil {
		return nil, err
	}
	err = h.datastore.DB.Select(h.fieldNames...).
		From("image").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *ImageHelper) Paged(siteUlid string, pageNum int32, itemsPerPage int32) (*ImagesPaged, error) {
	pd, err := h.PagedBy(siteUlid, pageNum, itemsPerPage, "date_created", "", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *ImageHelper) PagedBy(siteUlid string, pageNum int32, itemsPerPage int32, orderByFieldName string, direction string, searchText string) (*ImagesPaged, error) {
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

	var records []*Image
	err := h.datastore.DB.Select(h.fieldNames...).
		From("image").
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
	h.datastore.DB.SQL(`select count(image_ulid) from image where site_ulid = $1`, siteUlid).QueryStruct(&count)
	paged := &base.PagedInfo{
		PageNumber: pageNum,
		OrderBy:    orderByFieldName,
		Direction:  dirEnum,
		Limit:      itemsPerPage,
		Total:      count,
		Search:     searchText,
	}
	return &ImagesPaged{
		PagedInfo: paged,
		Records:   records,
	}, nil
}

func (h *ImageHelper) Save(siteUlid string, record *Image) error {
	return h.save(siteUlid, record)
}

func (h *ImageHelper) SaveMany(siteUlid string, records []*Image) error {
	for _, record := range records {
		err := h.save(siteUlid, record)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *ImageHelper) save(siteUlid string, record *Image) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteUlid != siteUlid {
		return errors.New("*Image update failed. SITEULID Mismatch")
	}
	cols := []string{"image", "original_height", "original_width", "top", "left", "scale", "crop_height", "crop_width", "site_ulid", "date_created", "date_modified", "is_processed"}
	colsSafe := make([]string, 0)
	for _, f := range cols {
		colsSafe = append(colsSafe, `"`+f+`"`)
	}
	vals := []interface{}{record.Image, record.OriginalHeight, record.OriginalWidth, record.Top, record.Left, record.Scale, record.CropHeight, record.CropWidth, record.SiteUlid, record.DateCreated, record.DateModified, record.IsProcessed}
	err = h.datastore.DB.Upsert("image").
		Columns(colsSafe...).
		Values(vals...).
		Where("image_ulid = $1", record.ImageUlid).
		Returning("image_ulid").
		QueryStruct(record)

	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *ImageHelper) Delete(siteUlid string, recordULID string) (bool, error) {
	result, err := h.datastore.DB.
		Update("image").
		Set("date_deleted", time.Now().Format(time.RFC3339)).
		Where("site_ulid=$1 and image_ulid=$2", siteUlid, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *ImageHelper) Purge(siteUlid string, recordULID string) (bool, error) {
	result, err := h.datastore.DB.
		DeleteFrom("image").
		Where("site_ulid=$1 and image_ulid=$2", siteUlid, recordULID).
		Exec()
	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *ImageHelper) validate(record *Image) (err error) {
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
