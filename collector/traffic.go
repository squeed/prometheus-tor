package collector

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yawning/bulb"
)

var (
	namespace       = "tor"
	TrafficReadDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "traffic", "read_bytes"),
		"Total traffic read",
		nil, nil,
	)
	TrafficWrittenDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "traffic", "written_bytes"),
		"Total traffic written",
		nil, nil,
	)
)

func ScrapeTraffic(c *bulb.Conn, ch chan<- prometheus.Metric) error {

	trafficRead, err := getInfoFloat(c, "traffic/read")
	if err != nil {
		return errors.Wrap(err, "could not scrape read_bytes")
	}

	trafficWritten, err := getInfoFloat(c, "traffic/written")
	if err != nil {
		return errors.Wrap(err, "could not scrape written_bytes")
	}

	fmt.Println("traffic read, written:", trafficRead, trafficWritten)

	ch <- prometheus.MustNewConstMetric(
		TrafficReadDesc, prometheus.CounterValue, float64(trafficRead),
	)

	ch <- prometheus.MustNewConstMetric(
		TrafficWrittenDesc, prometheus.CounterValue, float64(trafficWritten),
	)

	return nil
}
