package access

import (
	"context"

	"github.com/nerdynz/datastore"
	"github.com/nerdynz/security"
	"github.com/twitchtv/twirp"
)

type AccessServer struct {
	*datastore.Datastore
	Key security.Key
}

func NewServer(datastore *datastore.Datastore, key security.Key) TwirpServer {
	s := &AccessServer{
		Datastore: datastore,
		Key:       key,
	}
	return NewAccessServer(s)
}

func (s *AccessServer) Login(ctx context.Context, req *UnauthorisedUser) (resp *SessionInfo, err error) {
	padlock := security.NewFromContext(ctx, s.Settings, s.Key)

	info, err := padlock.LoginToSiteReturningInfo(req.Email, req.Password, req.SiteUlid)
	if err != nil {
		return nil, err
	}

	sites := make([]*Site, 0)
	for _, site := range info.Sites {
		sites = append(sites, &Site{
			Name:     site.Name,
			SiteUlid: site.SiteULID,
		})
	}
	return &SessionInfo{
		Token:      info.Token,
		Expiration: info.Expiration.UTC().UnixNano(),
		User: &SessionUser{
			Username: info.User.Username,
			Email:    info.User.Email,
			Password: info.User.Password,
			Role:     info.User.Role,
			Picture:  info.User.Picture,
			Initials: info.User.Initials,
			SiteUlid: info.User.SiteULID,
			Ulid:     info.User.ULID,
		},
		Sites: sites,
	}, nil
}

func (s *AccessServer) ValidSites(ctx context.Context, req *SitesQuery) (*Sites, error) {
	padlock := security.NewFromContext(ctx, s.Settings, s.Key)

	email := req.Email
	if email == "" {
		return nil, twirp.NewError(twirp.InvalidArgument, "No email provided")
	}

	sites, err := padlock.Sites(email)
	if err != nil {
		return nil, twirp.NewError(twirp.InvalidArgument, "No sites for user for email "+email)
	}

	ss := make([]*Site, 0)
	for _, s := range sites {
		ss = append(ss, &Site{
			SiteUlid: s.SiteULID,
			Name:     s.Name,
		})
	}

	return &Sites{
		Sites: ss,
	}, nil
}

func (s *AccessServer) Logout(ctx context.Context, req *InvalidateUser) (*InvalidateSuccess, error) {
	padlock := security.NewFromContext(ctx, s.Settings, s.Key)
	success, _ := padlock.Logout()
	// if err != nil {
	s.Logger.Info("xxx", success)
	// 	return nil, twirp.NewError(twirp.Internal, err.Error())
	// }
	return &InvalidateSuccess{
		IsSuccess: true,
	}, nil
}
