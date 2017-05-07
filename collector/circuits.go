package collector

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yawning/bulb"
)

var (
	NumCircuitsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "connection", "circuits"),
		"Count of currently open circuits",
		nil, nil,
	)
	NumStreamsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "connection", "streams"),
		"Count of currently open streams",
		nil, nil,
	)
	NumOrconnsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "connection", "orconns"),
		"Count of currently open ORConns",
		nil, nil,
	)
)

func ScrapeCircuits(c *bulb.Conn, ch chan<- prometheus.Metric) error {
	numCircuits, err := linesWithMatch(c, "circuit-status", " BUILT ")
	if err != nil {
		return errors.Wrapf(err, "could not scrape circuit-status")
	}

	numStreams, err := linesWithMatch(c, "stream-status", "SUCCEEDED")
	if err != nil {
		return errors.Wrapf(err, "could not scrape stream-status")
	}

	numOrconns, err := linesWithMatch(c, "orconn-status", " CONNECTED")
	if err != nil {
		return errors.Wrapf(err, "could not scrape orconn-status")
	}

	fmt.Println("circuits, streams, onions", numCircuits, numStreams, numOrconns)

	ch <- prometheus.MustNewConstMetric(
		NumCircuitsDesc, prometheus.GaugeValue, float64(numCircuits),
	)
	ch <- prometheus.MustNewConstMetric(
		NumStreamsDesc, prometheus.GaugeValue, float64(numStreams),
	)
	ch <- prometheus.MustNewConstMetric(
		NumOrconnsDesc, prometheus.GaugeValue, float64(numOrconns),
	)
	return nil
}

// Do a GETINFO %val, return the number of lines that match %match
func linesWithMatch(c *bulb.Conn, val string, match string) (int, error) {
	resp, err := c.Request("GETINFO " + val)
	if err != nil {
		return -1, errors.Wrap(err, "GETINFO circuit-status failed")
	}
	if len(resp.Data) < 2 {
		return 0, nil
	}

	ct := 0
	for _, line := range strings.Split(resp.Data[1], "\n") {
		if strings.Contains(line, match) {
			ct++
		}
	}

	return ct, nil
}
