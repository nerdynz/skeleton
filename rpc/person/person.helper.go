package person

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

var personHelperGlobal *PersonHelper

func (h *PersonHelper) beforeSave(record *Person) (err error) {
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

func (h *PersonHelper) afterSave(record *Person) (err error) {
	return err
}

// GENERATED CODE - Leave the below code alone
type PersonHelper struct {
	datastore  *datastore.Datastore
	fieldNames []string
	orderBy    string
}

//func PersonHelper(datastore  *datastore.Datastore) *PersonHelper {
//	if personHelperGlobal == nil {
//		personHelperGlobal = newPersonHelper(datastore)
//	}
//	return personHelperGlobal
//}

func newPersonHelper(datastore *datastore.Datastore) *PersonHelper {
	helper := &PersonHelper{}
	helper.datastore = datastore

	// Fields
	fieldnames := []string{"person_ulid", "site_ulid", "date_created", "date_modified", "name", "username", "email", "phone", "role", "initials", "password"}
	sort.Strings(fieldnames) // sort it makes searching it work correctly
	helper.fieldNames = fieldnames

	helper.orderBy = "date_created, date_modified"
	return helper
}

func (h *PersonHelper) New(siteUlid string) *Person {
	record := &Person{}
	// check DateCreated
	record.DateCreated = time.Now().Format(time.RFC3339)
	record.SiteUlid = siteUlid
	return record
}

func (h *PersonHelper) FromRequest(siteUlid string, req *http.Request) (*Person, error) {
	record := h.New(siteUlid)

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(record)
	if err != nil {
		return nil, err
	}

	if record.SiteUlid != siteUlid {
		return nil, errors.New("Person update failed. Site ULID Mismatch")
	}
	record.SiteUlid = siteUlid
	return record, nil
}

func (h *PersonHelper) Load(siteUlid string, ulid string) (*Person, error) {
	record, err := h.One(siteUlid, "person_ulid = $1 and $SITEULID", ulid)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (h *PersonHelper) All(siteUlid string) ([]*Person, error) {
	var records []*Person
	err := h.datastore.DB.Select(h.fieldNames...).
		From("person").
		Where("site_ulid = $1", siteUlid).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (h *PersonHelper) Where(siteUlid string, sql string, args ...interface{}) ([]*Person, error) {
	var records []*Person
	sql, args, err := datastore.AppendSiteULID(siteUlid, sql, args...)
	if err != nil {
		return nil, err
	}
	err = h.datastore.DB.Select(h.fieldNames...).
		From("person").
		Where(sql, args...).
		OrderBy(h.orderBy).
		QueryStructs(&records)

	if err != nil {
		return nil, err
	}
	return records, nil
}

func (h *PersonHelper) SQL(siteUlid string, sql string, args ...interface{}) ([]*Person, error) {
	if !strings.Contains(sql, "$SITEULID") {
		return nil, errors.New("No $SITEULID placeholder defined")
	}
	var records []*Person
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

func (h *PersonHelper) One(siteUlid string, sql string, args ...interface{}) (*Person, error) {
	var record Person
	sql, args, err := datastore.AppendSiteULID(siteUlid, sql, args...)

	if err != nil {
		return nil, err
	}
	err = h.datastore.DB.Select(h.fieldNames...).
		From("person").
		Where(sql, args...).
		OrderBy(h.orderBy).
		Limit(1).
		QueryStruct(&record)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (h *PersonHelper) Paged(siteUlid string, pageNum int32, itemsPerPage int32) (*PeoplePaged, error) {
	pd, err := h.PagedBy(siteUlid, pageNum, itemsPerPage, "date_created", "", "") // date_created should be the most consistant because it doesn't change
	if err != nil {
		return nil, err
	}
	return pd, nil
}

func (h *PersonHelper) PagedBy(siteUlid string, pageNum int32, itemsPerPage int32, orderByFieldName string, direction string, searchText string) (*PeoplePaged, error) {
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

	var records []*Person
	err := h.datastore.DB.Select(h.fieldNames...).
		From("person").
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
	h.datastore.DB.SQL(`select count(person_ulid) from person where site_ulid = $1`, siteUlid).QueryStruct(&count)
	paged := &base.PagedInfo{
		PageNumber: pageNum,
		OrderBy:    orderByFieldName,
		Direction:  dirEnum,
		Limit:      itemsPerPage,
		Total:      count,
		Search:     searchText,
	}
	return &PeoplePaged{
		PagedInfo: paged,
		Records:   records,
	}, nil
}

func (h *PersonHelper) Save(siteUlid string, record *Person) error {
	return h.save(siteUlid, record)
}

func (h *PersonHelper) SaveMany(siteUlid string, records []*Person) error {
	for _, record := range records {
		err := h.save(siteUlid, record)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *PersonHelper) save(siteUlid string, record *Person) error {
	err := h.beforeSave(record)
	if err != nil {
		return err
	}

	if record.SiteUlid != siteUlid {
		return errors.New("*Person update failed. SITEULID Mismatch")
	}
	cols := []string{"site_ulid", "date_created", "date_modified", "name", "username", "email", "phone", "role", "initials", "password"}
	vals := []interface{}{record.SiteUlid, record.DateCreated, record.DateModified, record.Name, record.Username, record.Email, record.Phone, record.Role, record.Initials, record.Password}
	err = h.datastore.DB.Upsert("person").
		Columns(cols...).
		Values(vals...).
		Where("person_ulid = $1", record.PersonUlid).
		Returning("person_ulid").
		QueryStruct(record)

	if err != nil {
		return err
	}
	err = h.afterSave(record)
	return err
}

func (h *PersonHelper) Delete(siteUlid string, recordULID string) (bool, error) {
	result, err := h.datastore.DB.
		Update("person").
		Set("date_deleted", time.Now().Format(time.RFC3339)).
		Where("site_ulid=$1 and person_ulid=$2", siteUlid, recordULID).
		Exec()

	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *PersonHelper) Purge(siteUlid string, recordULID string) (bool, error) {
	result, err := h.datastore.DB.
		DeleteFrom("person").
		Where("site_ulid=$1 and person_ulid=$2", siteUlid, recordULID).
		Exec()
	if err != nil {
		return false, err
	}

	return (result.RowsAffected > 0), nil
}

func (h *PersonHelper) validate(record *Person) (err error) {
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
