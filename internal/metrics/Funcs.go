package metrics

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (m *Metric) CheckDB() {
	if err := m.PG.Sqlinstance.Ping(); err != nil {
		m.List.PGisAlive.Set(0)
	} else {
		m.List.PGisAlive.Set(1)
	}
}

func (m *Metric) MetricHandler(router *gin.Engine) {
	m.RegisterList()
	router.GET("/metrics", gin.WrapH(promhttp.HandlerFor(
		m.Registry,
		promhttp.HandlerOpts{},
	)))

	go func() {
		ticker := time.NewTicker(15 * time.Second)
		for range ticker.C {
			m.CheckDB()
		}
	}()
}
