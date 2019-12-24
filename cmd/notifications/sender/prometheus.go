/*
 * HomeWork-18: Prometheus monitoring
 * Created on 24.12.2019 17:07
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"net/http"

	"github.com/evakom/calendar/internal/loggers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type prometheus struct {
	listen string
	logger loggers.Logger
}

func newPrometheus(listen string) *prometheus {
	return &prometheus{
		listen: listen,
		logger: loggers.GetLogger(),
	}
}

func (p *prometheus) start() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		p.logger.Info("Starting prometheus exporter at port: %s", p.listen)
		if err := http.ListenAndServe(":9102", nil); err != nil {
			p.logger.Error("Error start prometheus exporter:", err)
			return
		}
	}()
}
