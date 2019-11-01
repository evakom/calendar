/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 01.11.2019 13:17
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"context"
	"github.com/evakom/calendar/internal/domain/models"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// Constants
const (
	HOSTFIELD    = "host"
	METHODFIELD  = "method"
	URLFIELD     = "url"
	BROWSERFIELD = "browser"
	REMOTEFIELD  = "remote"
	QUERYFIELD   = "query"
)

type contextKey string

const contextKeyRequestID contextKey = "requestID"

func requestFields(r *http.Request, args ...string) models.Fields {
	fields := make(models.Fields)
	for _, s := range args {
		switch s {
		case HOSTFIELD:
			fields[HOSTFIELD] = r.Host
			fallthrough
		case METHODFIELD:
			fields[METHODFIELD] = r.Method
			fallthrough
		case URLFIELD:
			fields[URLFIELD] = r.URL.Path
			fallthrough
		case BROWSERFIELD:
			fields[BROWSERFIELD] = r.Header.Get("User-Agent")
			fallthrough
		case REMOTEFIELD:
			fields[REMOTEFIELD] = r.RemoteAddr
			fallthrough
		case QUERYFIELD:
			fields[QUERYFIELD] = r.URL.RawQuery
		}
	}
	return fields
}

func assignRequestID(ctx context.Context) context.Context {
	reqID := uuid.NewV4()
	return context.WithValue(ctx, contextKeyRequestID, reqID.String())
}

func getRequestID(ctx context.Context) string {
	reqID := ctx.Value(contextKeyRequestID)
	if key, ok := reqID.(string); ok {
		return key
	}
	return ""
}
