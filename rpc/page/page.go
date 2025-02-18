package page

import (

	//runner "github.com/nerdynz/dat/sqlx-runner"

	"github.com/jinzhu/copier"
	block "github.com/nerdynz/skeleton/rpc/block"
	img "github.com/nerdynz/skeleton/rpc/image"
)

type Metadata struct {
	MainNav []*NavLink `json:"mainNav"`
}

func (h *PageHelper) LoadMetadata() (*Metadata, error) {
	mainNav, err := h.LoadMainNav()
	if err != nil {
		return nil, err
	}
	return &Metadata{
		MainNav: mainNav,
	}, nil
}

func (h *PageHelper) LoadMainNav() ([]*NavLink, error) {
	var records []*NavLink
	err := h.datastore.DB.SQL("select slug, title from page").QueryStructs(&records)
	return records, err
}

func (pageHelper *PageHelper) LoadFullPage(siteUlid string, slug string) (*PageWithBlockWithImage, error) {
	page, err := pageHelper.One(siteUlid, " (page_ulid = $1 or slug = $1) and $SITEULID ", slug)
	if err != nil {
		return nil, err
	}

	pageBlockWithImage := &PageWithBlockWithImage{}
	err = copier.Copy(pageBlockWithImage, page)
	if err != nil {
		return nil, err
	}

	blockHelper := block.NewBlockHelper(pageHelper.datastore)

	records, err := blockHelper.Where(siteUlid, "$SITEULID and page_ulid = $1", pageBlockWithImage.PageUlid)
	if err != nil {
		return nil, err
	}

	allImageUlids := make([]string, 0)
	for _, rec := range records {
		if rec.ImageOneUlid != "" {
			allImageUlids = append(allImageUlids, rec.ImageOneUlid)
		}
		if rec.ImageTwoUlid != "" {
			allImageUlids = append(allImageUlids, rec.ImageTwoUlid)
		}
		if rec.ImageThreeUlid != "" {
			allImageUlids = append(allImageUlids, rec.ImageThreeUlid)
		}
		if rec.ImageFourUlid != "" {
			allImageUlids = append(allImageUlids, rec.ImageFourUlid)
		}
	}

	imageHelper := img.NewImageHelper(pageHelper.datastore)
	var images []*img.Image
	if len(allImageUlids) > 0 {
		images, err = imageHelper.Where(siteUlid, " $SITEULID and image_ulid in $1", allImageUlids)
		if err != nil {
			return nil, err
		}
	}

	blocksWithImages := make([]*block.BlockWithImage, 0)
	for _, rec := range records {
		bWI := &block.BlockWithImage{}
		copier.Copy(bWI, rec)
		for _, img := range images {
			if rec.ImageOneUlid == img.ImageUlid {
				bWI.ImageOne = img
			}
			if rec.ImageTwoUlid == img.ImageUlid {
				bWI.ImageTwo = img
			}
			if rec.ImageThreeUlid == img.ImageUlid {
				bWI.ImageThree = img
			}
			if rec.ImageFourUlid == img.ImageUlid {
				bWI.ImageFour = img
			}
		}

		blocksWithImages = append(blocksWithImages, bWI)
	}
	pageBlockWithImage.Blocks = blocksWithImages
	return pageBlockWithImage, nil
}
