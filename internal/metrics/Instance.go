package metrics

import (
	database "v2/internal/Database"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metric struct {
	List     *MetricList
	Registry *prometheus.Registry
	PG       database.DBInstance
}

func SetupMetrics(pg database.DBInstance) *Metric {
	return &Metric{
		List:     SetupMetricList(),
		Registry: prometheus.NewRegistry(),
		PG:       pg,
	}
}

type MetricList struct {
	PGisAlive prometheus.Gauge
}

func (m *Metric) RegisterList() {
	m.Registry.MustRegister(m.List.PGisAlive)
}

func SetupMetricList() *MetricList {
	return &MetricList{
		PGisAlive: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "db_up",
			Help: "Checking status of db is alive or not",
		}),
	}
}
