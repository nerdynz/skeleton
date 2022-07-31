package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/nerdynz/datastore"
	"github.com/nerdynz/security"
)

type Key struct {
	Store *datastore.Datastore
}

func (k *Key) GetLogin(authToken string) (*security.SessionInfo, error) {
	// k.Store.Logger.Info("geting user info", authToken)
	i := &security.SessionInfo{}
	blob, err := k.GetCacheValue(authToken, "")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(blob, i)
	// k.Store.Logger.Info("geting user info", i)
	return i, err
}

func (k *Key) SetLogin(authToken string, i *security.SessionInfo, duration time.Duration) error {
	// k.Store.Logger.Info("seting user info ", authToken, i)
	ubts, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return k.SetCacheValue("", authToken, ubts, duration) // we dont need to set a userkey for the login because this is our user key
}

func (k *Key) ExpireLoggedInUser(key string) error {
	dur := 1 * time.Second
	return k.SetCacheValue("", key, nil, dur) // cya
}

func (k *Key) DoLogin(notAuthorizedUser *security.NotAuthorizedUser) (*security.SessionUser, error) {
	loggedInUser := &security.SessionUser{}
	sql := k.Store.DB.
		Select("person_ulid as ulid, name, email, password, role, site_ulid").
		From("person")
	if notAuthorizedUser.ULID > "" {
		sql.Where("person_ulid = $1", notAuthorizedUser.ULID) // id doesn't need a password as we already know who they are
	} else if notAuthorizedUser.SiteULID != "" {
		sql.Where("email = $1 and password = $2 and site_ulid = $3", notAuthorizedUser.Email, notAuthorizedUser.Password, notAuthorizedUser.SiteULID)
	} else {
		sql.Where("email = $1 and password = $2", notAuthorizedUser.Email, notAuthorizedUser.Password)
	}

	sql.Limit(1)
	err := sql.QueryStruct(loggedInUser)
	if loggedInUser.ULID == "" {
		return nil, errors.New("Failed to login invalid user")
	}
	return loggedInUser, err
}

func (k *Key) SetCacheValue(userkey string, key string, value []byte, duration time.Duration) error {
	err := k.Store.Cache.SetBytes(userkey+key, value, duration)
	return err
}
func (k *Key) GetCacheValue(userkey string, key string) ([]byte, error) {
	val, err := k.Store.Cache.GetBytes(userkey + key)
	if err != nil {
		return nil, errors.New("value not set")
	}
	return val, err
}

func (k *Key) GetAuthToken(req *http.Request) (string, error) {
	// there are some basic checks built in so this is an extension
	authToken := req.Header.Get("authtoken")
	return authToken, nil
}

func (k *Key) GetSites(email string) ([]*security.Site, error) {
	sites := make([]*security.Site, 0)
	err := k.Store.DB.SQL(`
		select s.site_ulid, s.name from site s
		join person p on s.site_ulid = p.site_ulid
		where p.email = $1
	`, email).QueryStructs(&sites)
	if err != nil {
		return nil, err
	}
	return sites, nil
}
