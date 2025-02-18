package block

import (
	"context"

	"github.com/nerdynz/datastore"
	base "github.com/nerdynz/skeleton/rpc/base"
	twirp "github.com/twitchtv/twirp"
)

type BlockServer struct {
	datastore   *datastore.Datastore
	blockHelper *BlockHelper
}

func NewServer(datastore *datastore.Datastore) TwirpServer {
	return NewBlockRPCServer(&BlockServer{
		datastore:   datastore,
		blockHelper: newBlockHelper(datastore),
	})
}

func (s *BlockServer) SaveBlock(ctx context.Context, block *Block) (*Block, error) {
	siteUlid := ctx.Value("site_ulid").(string)
	err := s.blockHelper.Save(siteUlid, block)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	return block, nil
}

func (s *BlockServer) LoadBlock(ctx context.Context, lookup *base.Lookup) (*Block, error) {
	siteUlid := ctx.Value("site_ulid").(string)
	block, err := s.blockHelper.Load(siteUlid, lookup.Ulid)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	return block, err
}

func (s *BlockServer) PagedBlocks(ctx context.Context, pagedInfo *base.PagedInfo) (*BlocksPaged, error) {
	siteUlid := ctx.Value("site_ulid").(string)
	pagedData, err := s.blockHelper.PagedBy(siteUlid, pagedInfo.PageNumber, pagedInfo.Limit, pagedInfo.OrderBy, pagedInfo.Direction.String(), pagedInfo.Search)
	if err != nil {
		return nil, twirp.NewError(twirp.Malformed, err.Error())
	}
	return pagedData, nil
}

func (s *BlockServer) DeleteBlock(ctx context.Context, lookup *base.Lookup) (*base.Deleted, error) {
	siteUlid := ctx.Value("site_ulid").(string)
	isDeleted, err := s.blockHelper.Delete(siteUlid, lookup.Ulid)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	return &base.Deleted{
		IsDeleted: isDeleted,
	}, nil
}
