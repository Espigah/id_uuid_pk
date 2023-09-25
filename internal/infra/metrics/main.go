package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	read  *prometheus.GaugeVec
	write *prometheus.GaugeVec
	count prometheus.Counter
}

func New(reg prometheus.Registerer, mode string) *Metrics {
	m := &Metrics{
		read: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "read_duration",
				Help: "Duration of read data.",
				ConstLabels: prometheus.Labels{
					"mode": mode,
				},
			},
			[]string{
				"id",
			},
		),
		write: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "write_duration",
			Help: "Duration of write data.",
			ConstLabels: prometheus.Labels{
				"mode": mode,
			},
		}, []string{"id"}),
		count: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "count_iteration",
			Help: "Counter.",
			ConstLabels: prometheus.Labels{
				"mode": mode,
			},
		}),
	}
	reg.MustRegister(m.read)
	reg.MustRegister(m.write)
	return m
}

func (m *Metrics) ReadDuration(id string, duration float64) {
	m.read.WithLabelValues(id).Set(duration)
}

func (m *Metrics) WriteDuration(id string, duration float64) {
	m.write.WithLabelValues(id).Set(duration)
}

func (m *Metrics) Count() {
	m.count.Inc()
}
