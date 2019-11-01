/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 22:08
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/models"
	"net/http"
	"time"
)

type handler struct {
	logger models.Logger
}

func newHandlers() *handler {
	return &handler{
		logger: models.Logger{}.GetLogger(),
	}
}

func (h handler) prepareRoutes() http.Handler {
	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/hello", h.helloHandler)
	siteMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqID := getRequestID(r.Context())
		h.logger.Error("RESPONSE ERROR, ID: %s, RETURN CODE: [%d]", reqID, http.StatusNotFound)
		http.NotFound(w, r)
	})
	siteHandler := h.loggerMiddleware(siteMux)
	siteHandler = h.panicMiddleware(siteHandler)
	//siteHandler = h.otherMiddleware(siteHandler)
	return siteHandler
}

func (h handler) panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.logger.Debug("Middleware 'panic' PASS")
		defer func() {
			if err := recover(); err != nil {
				h.logger.Error("recovered from panic: %s", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

//func (h handler) otherMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		h.logger.Warn("other Middleware")
//		next.ServeHTTP(w, r)
//	})
//}

func (h handler) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := assignRequestID(r.Context())
		r = r.WithContext(ctx)
		h.logger.Fields = requestFields(r, HOSTFIELD, METHODFIELD, URLFIELD, BROWSERFIELD, REMOTEFIELD, QUERYFIELD)
		h.logger.WithFields().Info("REQUEST START, ID: %s", getRequestID(ctx))
		start := time.Now()
		next.ServeHTTP(w, r)
		h.logger.Info("REQUEST END, ID: %s, TIME: [%s]", getRequestID(ctx), time.Since(start))
	})
}

func (h handler) helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "default name"
	}
	reqID := getRequestID(r.Context())
	h.logger.Info("RESPONSE, ID: %s, RETURN CODE: [%d]", reqID, http.StatusOK)
	if _, err := fmt.Fprint(w, "Hello, my name is ", name); err != nil {
		h.logger.Error("Error write to response writer!")
	}
}
