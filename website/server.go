/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 18:18
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package website implements http server control.
package website

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/models"
	"net/http"
	"path"
)

type pathResolver struct {
	handlers map[string]http.HandlerFunc
	logger   models.Logger
}

func newPathResolver() *pathResolver {
	return &pathResolver{
		handlers: make(map[string]http.HandlerFunc),
		logger:   models.Logger{}.GetLogger(),
	}
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.logger.Fields = requestFields(r)
	p.logger.WithFields().Info("REQUEST")
	check := r.Method + " " + r.URL.Path
	for pattern, handlerFunc := range p.handlers {
		ok, err := path.Match(pattern, check)
		if err != nil {
			p.logger.WithFields().Error("RESPONSE [%d]", http.StatusInternalServerError)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if ok {
			handlerFunc(w, r)
			return
		}
	}
	p.logger.WithFields().Error("RESPONSE [%d]", http.StatusNotFound)
	http.NotFound(w, r)
}

// StartWebsite inits routing and starts web listener.
func StartWebsite() {
	pr := newPathResolver()
	pr.Add("GET /hello", helloHandler)
	http.ListenAndServe(":8080", pr)
}

func requestFields(r *http.Request) models.Fields {
	fields := make(models.Fields)
	fields["host"] = r.Host
	fields["method"] = r.Method
	fields["url"] = r.URL
	//fields["browser"] = r.
	return fields
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "default name"
	}
	fmt.Fprint(w, "Hello, my name is ", name)
}
