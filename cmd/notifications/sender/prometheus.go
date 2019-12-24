/*
 * HomeWork-18: Prometheus monitoring
 * Created on 24.12.2019 17:07
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"net/http"

	"github.com/evakom/calendar/internal/loggers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type prometMonitor struct {
	listen            string
	messagesPerSecond prometheus.Gauge
	ch                chan float64
	logger            loggers.Logger
}

func newPrometheus(listen string) *prometMonitor {
	return &prometMonitor{
		listen: listen,
		messagesPerSecond: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "calendar_sender_messages_per_second",
			Help: "Messages per second sent to users",
		}),
		ch:     make(chan float64, 1),
		logger: loggers.GetLogger(),
	}
}

func (p *prometMonitor) start() {
	go func() {
		for g := range p.ch {
			p.messagesPerSecond.Set(g)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		p.logger.Info("Starting prometheus exporter at port: %s", p.listen)
		if err := http.ListenAndServe(":9102", nil); err != nil {
			p.logger.Error("Error start prometheus exporter:", err)
			return
		}
	}()
}

func (p *prometMonitor) stop() {
	close(p.ch)
	p.logger.Info("Stopped prometheus")
}