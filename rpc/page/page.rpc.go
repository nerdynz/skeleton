package page

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/nerdynz/datastore"
	base "github.com/nerdynz/skeleton/rpc/base"
	block "github.com/nerdynz/skeleton/rpc/block"
	"github.com/nerdynz/skeleton/rpc/image"
	"github.com/sirupsen/logrus"
	twirp "github.com/twitchtv/twirp"
)

type PageServer struct {
	datastore   *datastore.Datastore
	pageHelper  *PageHelper
	blockHelper *block.BlockHelper
	imageHelper *image.ImageHelper
}

func NewServer(datastore *datastore.Datastore) *PageServer {
	return &PageServer{
		datastore:   datastore,
		pageHelper:  newPageHelper(datastore),
		blockHelper: block.NewBlockHelper(datastore),
		imageHelper: image.NewImageHelper(datastore),
	}
}

func NewRpcServer(datastore *datastore.Datastore) TwirpServer {
	return NewPageRPCServer(NewServer(datastore))
}

func (s *PageServer) SavePage(ctx context.Context, page *Page) (*Page, error) {
	siteUlid := "01EHZXH0YBCM8Q8PEFDZB8K3WW" //ctx.Value("site_ulid").(string)
	err := s.pageHelper.Save(siteUlid, page)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	return page, nil
}

func (s *PageServer) LoadPage(ctx context.Context, lookup *base.Lookup) (*Page, error) {
	siteUlid := "01EHZXH0YBCM8Q8PEFDZB8K3WW" //ctx.Value("site_ulid").(string)
	page, err := s.pageHelper.Load(siteUlid, lookup.Ulid)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	return page, err
}

func (s *PageServer) SaveFullPage(ctx context.Context, pageWithBlockWithImage *PageWithBlockWithImage) (*PageWithBlockWithImage, error) {
	siteUlid := "01EHZXH0YBCM8Q8PEFDZB8K3WW" //ctx.Value("site_ulid").(string)

	logrus.Info("site_ulid", siteUlid)

	page := s.pageHelper.New(siteUlid)
	err := copier.Copy(page, pageWithBlockWithImage)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	err = s.pageHelper.Save(siteUlid, page)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}

	blockUlids := make([]string, 0)
	for _, b := range pageWithBlockWithImage.Blocks {
		blockUlids = append(blockUlids, b.BlockUlid)
	}

	_, err = s.datastore.DB.SQL("delete from block where page_ulid = $1 and block_ulid not in $2", pageWithBlockWithImage.PageUlid, blockUlids).Exec()
	if err != nil {
		return nil, err
	}

	for _, b := range pageWithBlockWithImage.Blocks {
		newBlock := &block.Block{}
		err := copier.Copy(newBlock, b)
		if err != nil {
			return nil, twirp.InternalError(err.Error())
		}
		logrus.Info("newBlock", newBlock)

		if b.ImageOne != nil && b.ImageOne.ImageUlid != "" {
			err = s.imageHelper.Save(siteUlid, b.ImageOne)
			if err != nil {
				return nil, twirp.InternalError(err.Error())
			}
			newBlock.ImageOneUlid = b.ImageOne.ImageUlid
		}
		err = s.blockHelper.Save(siteUlid, newBlock)
		if err != nil {
			return nil, twirp.InternalError(err.Error())
		}
	}

	return pageWithBlockWithImage, nil
}

func (s *PageServer) ServeFullPageJson(w http.ResponseWriter, r *http.Request) {
	siteUlid := r.URL.Query().Get("site")
	pageUlid := r.URL.Query().Get("page")

	p, err := s.loadFullPage(siteUlid, pageUlid)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func (s *PageServer) LoadFullPage(ctx context.Context, lookup *base.Lookup) (*PageWithBlockWithImage, error) {
	siteUlid := "01EHZXH0YBCM8Q8PEFDZB8K3WW" //ctx.Value("site_ulid").(string)
	page, err := s.loadFullPage(siteUlid, lookup.Ulid)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	return page, nil
}

func (s *PageServer) loadFullPage(siteUlid string, pageUlid string) (*PageWithBlockWithImage, error) {
	page, err := s.pageHelper.One(siteUlid, " (page_ulid = $1 or slug = $1) and $SITEULID ", pageUlid)
	if err != nil {
		return nil, err
	}

	pageBlockWithImage := &PageWithBlockWithImage{}
	err = copier.Copy(pageBlockWithImage, page)
	if err != nil {
		return nil, err
	}

	records, err := s.blockHelper.Where(siteUlid, "$SITEULID and page_ulid = $1", pageBlockWithImage.PageUlid)
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

	var images []*image.Image
	if len(allImageUlids) > 0 {
		images, err = s.imageHelper.Where(siteUlid, " $SITEULID and image_ulid in $1", allImageUlids)
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

func (s *PageServer) PagedPages(ctx context.Context, pagedInfo *base.PagedInfo) (*PagesPaged, error) {
	siteUlid := "01EHZXH0YBCM8Q8PEFDZB8K3WW" //ctx.Value("site_ulid").(string)
	pagedData, err := s.pageHelper.PagedBy(siteUlid, pagedInfo.PageNumber, pagedInfo.Limit, pagedInfo.OrderBy, pagedInfo.Direction.String(), pagedInfo.Search)
	if err != nil {
		return nil, twirp.NewError(twirp.Malformed, err.Error())
	}
	return pagedData, nil
}

func (s *PageServer) DeletePage(ctx context.Context, lookup *base.Lookup) (*base.Deleted, error) {
	siteUlid := "01EHZXH0YBCM8Q8PEFDZB8K3WW" //ctx.Value("site_ulid").(string)
	isDeleted, err := s.pageHelper.Purge(siteUlid, lookup.Ulid)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	return &base.Deleted{
		IsDeleted: isDeleted,
	}, nil
}
