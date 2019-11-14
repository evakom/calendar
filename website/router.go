/*
 * HomeWork-10: Calendar extending HTTP methods
 * Created on 14.11.2019 22:18
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"github.com/evakom/calendar/internal/loggers"
	"net/http"
	"time"
)

func (h handler) prepareRoutes() http.Handler {
	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/hello/", h.helloHandler)
	siteMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h.logger.WithFields(loggers.Fields{
			CodeField: http.StatusNotFound,
			IDField:   getRequestID(r.Context()),
		}).Error("RESPONSE")
		http.NotFound(w, r)
	})
	siteHandler := h.loggerMiddleware(siteMux)
	siteHandler = h.panicMiddleware(siteHandler)
	return siteHandler
}

func (h handler) panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.logger.Debug("Middleware 'panic' PASS")
		defer func() {
			if err := recover(); err != nil {
				h.logger.Error("recovered from panic: %s", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (h handler) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := assignRequestID(r.Context())
		r = r.WithContext(ctx)
		h.logger.WithFields(requestFields(
			r, IDField, HostField, MethodField, URLField,
			BrowserField, RemoteField, QueryField,
		)).Info("REQUEST START")
		start := time.Now()
		next.ServeHTTP(w, r)
		h.logger.WithFields(loggers.Fields{
			RespTimeField: time.Since(start),
			IDField:       getRequestID(ctx),
		}).Info("REQUEST END")
	})
}