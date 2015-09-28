// Copyright (c) 2015 RightScale Inc, All Rights Reserved

package demo

import (
	"net/http"

	"github.com/rightscale/gojiutil"
	"github.com/zenazn/goji/web"
	"gopkg.in/inconshreveable/log15.v2"
)

// simple string->string map for demo purposes
var settings = make(map[string]string)

// helper function to get the context logger, used in putSetting and deleteSetting,
// see alse an alternative way to use the context logger in getSetting
func log(c web.C) log15.Logger { return c.Env[gojiutil.ContextLog].(log15.Logger) }

func indexSetting(c web.C, rw http.ResponseWriter, r *http.Request) {
	gojiutil.WriteJSON(c, rw, 200, settings)
}

// getSetting retrieves a setting from the settings map
func getSetting(c web.C, rw http.ResponseWriter, r *http.Request) {
	log := c.Env[gojiutil.ContextLog].(log15.Logger) // get the context logger
	key := c.URLParams["key"]
	value := settings[key]
	if key == "" || value == "" {
		gojiutil.Errorf(c, rw, 404, `settings key '%s' not found`, key)
		return
	}
	log.Info("settings", "op", "get", "key", key, "value", value)
	gojiutil.WriteString(rw, 200, value)
}

func putSetting(c web.C, rw http.ResponseWriter, r *http.Request) {
	key := c.URLParams["key"]
	if key == "" {
		gojiutil.ErrorString(c, rw, 413, `settings key missing`)
		return
	}
	value := r.Form.Get("value")
	if key == "" {
		gojiutil.ErrorString(c, rw, 413, `value query string param missing`)
		return
	}
	log(c).Info("settings", "op", "put", "key", key, "value", value)
	settings[key] = value
}

func deleteSetting(c web.C, rw http.ResponseWriter, r *http.Request) {
	key := c.URLParams["key"]
	log(c).Info("settings", "op", "delete", "key", key)
	delete(settings, key)
	rw.WriteHeader(201)
}
