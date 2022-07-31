package person

import (
	"context"

	"github.com/nerdynz/datastore"
	base "github.com/nerdynz/skeleton/rpc/base"
	twirp "github.com/twitchtv/twirp"
)

type PersonServer struct {
	datastore    *datastore.Datastore
	personHelper *PersonHelper
}

func NewServer(datastore *datastore.Datastore) TwirpServer {
	return NewPersonRPCServer(&PersonServer{
		datastore:    datastore,
		personHelper: newPersonHelper(datastore),
	})
}

func (s *PersonServer) SavePerson(ctx context.Context, person *Person) (*Person, error) {
	siteUlid := ctx.Value("site_ulid").(string)
	err := s.personHelper.Save(siteUlid, person)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	return person, nil
}

func (s *PersonServer) LoadPerson(ctx context.Context, lookup *base.Lookup) (*Person, error) {
	siteUlid := ctx.Value("site_ulid").(string)
	person, err := s.personHelper.Load(siteUlid, lookup.Ulid)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	return person, err
}

func (s *PersonServer) PagedPeople(ctx context.Context, pagedInfo *base.PagedInfo) (*PeoplePaged, error) {
	siteUlid := ctx.Value("site_ulid").(string)
	pagedData, err := s.personHelper.PagedBy(siteUlid, pagedInfo.PageNumber, pagedInfo.Limit, pagedInfo.OrderBy, pagedInfo.Direction.String(), pagedInfo.Search)
	if err != nil {
		return nil, twirp.NewError(twirp.Malformed, err.Error())
	}
	return pagedData, nil
}

func (s *PersonServer) DeletePerson(ctx context.Context, lookup *base.Lookup) (*base.Deleted, error) {
	siteUlid := ctx.Value("site_ulid").(string)
	isDeleted, err := s.personHelper.Purge(siteUlid, lookup.Ulid)
	if err != nil {
		return nil, twirp.InternalError(err.Error())
	}
	return &base.Deleted{
		IsDeleted: isDeleted,
	}, nil
}
