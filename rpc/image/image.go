package image

import (
	"context"
	"strings"
	"time"
)

func (h *ImageHelper) selectFrom() string {
	fields := make([]string, 0)
	for _, f := range h.fieldNames() {
		if strings.HasPrefix(f, "date_") {
			f = f + "::text"
		}
		fields = append(fields, f)
	}
	return `
	SELECT 
		` + strings.Join(fields, ", ") + `
	FROM 
		image
	`
}

func (h *ImageHelper) fieldNames() []string {
	return []string{image_ulid, image, original_height, original_width, top, left, scale, crop_height, crop_width, site_ulid, date_created, date_modified, is_processed}
}

func (h *ImageHelper) orderBy() string {
	return "date_created, date_modified"
}

func (h *ImageHelper) beforeSave(record *Image) (err error) {
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

func (h *ImageHelper) Save(ctx context.Context, record *Image) error {
	return h.save(ctx, record)
}

func (h *ImageHelper) SaveMany(ctx context.Context, records []*Image) error {
	for _, record := range records {
		err := h.save(ctx, record)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *ImageHelper) afterSave(record *Image) (err error) {
	return err
}

func (h *ImageHelper) validate(record *Image) (err error) {
	return nil
	// validationErrors := h.validator.Struct(record)
	//
	//	if validationErrors != nil {
	//		errMessage := ""
	//		for _, err := range err.(validator.ValidationErrors) {
	//			errMessage += err.Kind().String() + " validation Error on field "+err.Field()
	//		}
	//		if errMessage != "" {
	//			err = errors.New(errMessage)
	//		}
	//	}
	//
	// return err
}
