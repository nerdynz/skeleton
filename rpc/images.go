package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nerdynz/datastore"
	"github.com/nerdynz/fileupload"
	"github.com/nerdynz/flow"
	"github.com/nerdynz/security"
	"github.com/sirupsen/logrus"
)

func uploadImage(w http.ResponseWriter, req *http.Request, ctx *flow.Flow, store *datastore.Datastore) {
	_, fullpath, err := fileupload.FromRequestToFile(req, "./static/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := struct {
		Ulid string `json:"ulid"`
		Path string `json:"path"`
	}{
		Ulid: security.ULID(),
		Path: strings.TrimPrefix(fullpath, "."),
	}

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logrus.Info(js)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
