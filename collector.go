package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

//Define the metrics we wish to expose
type MetricsCollector struct {
	domainTimesListedMetric         *prometheus.GaugeVec
	scrapesTotalMetric              prometheus.Counter
	scrapeErrorsTotalMetric         prometheus.Counter
	lastScrapeErrorMetric           prometheus.Gauge
	lastScrapeTimestampMetric       prometheus.Gauge
	lastScrapeDurationSecondsMetric prometheus.Gauge
}

/*
Result holds the individual IP lookup results for each bl search
*/
type Result struct {
	// Listed indicates whether or not the IP was on the bl
	Listed bool `json:"listed"`
	// bl lists sometimes add extra information as a TXT record
	// if any info is present, it will be stored here.
	Text string `json:"text"`
	// Error represents any error that was encountered (DNS timeout, host not
	// found, etc.) if any
	Error bool `json:"error"`
	// ErrorType is the type of error encountered if any
	ErrorType error `json:"error_type"`
}

//NewMetricsCollector TODO
func NewMetricsCollector() *MetricsCollector {
	domainTimesListedMetric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "dnsbl",
			Subsystem: "metrics",
			Name:      "listed_times",
			Help:      "Number of times the domain listed on blacklists",
		},
		[]string{"domain", "blacklist"},
	)
	scrapesTotalMetric := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "dnsbl",
			Subsystem: "metrics",
			Name:      "scrapes_total",
			Help:      "Total number of dnsbl metrics scrapes.",
		},
	)

	scrapeErrorsTotalMetric := prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "dnsbl",
			Subsystem: "metrics",
			Name:      "scrape_errors_total",
			Help:      "Total number of dnsbl metrics scrape errors.",
		},
	)

	lastScrapeErrorMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "dnsbl",
			Subsystem: "metrics",
			Name:      "last_scrape_error",
			Help:      "Whether the last metrics scrape from dnsbl resulted in an error (1 for error, 0 for success).",
		},
	)

	lastScrapeTimestampMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "dnsbl",
			Subsystem: "metrics",
			Name:      "last_scrape_timestamp",
			Help:      "Number of seconds since 1970 since last metrics scrape from dnsbl.",
		},
	)

	lastScrapeDurationSecondsMetric := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "dnsbl",
			Subsystem: "metrics",
			Name:      "last_scrape_duration_seconds",
			Help:      "Duration of the last metrics scrape from dnsbl.",
		},
	)
	metricsCollector := &MetricsCollector{
		domainTimesListedMetric:         domainTimesListedMetric,
		scrapesTotalMetric:              scrapesTotalMetric,
		scrapeErrorsTotalMetric:         scrapeErrorsTotalMetric,
		lastScrapeErrorMetric:           lastScrapeErrorMetric,
		lastScrapeTimestampMetric:       lastScrapeTimestampMetric,
		lastScrapeDurationSecondsMetric: lastScrapeDurationSecondsMetric,
	}
	return metricsCollector
}

// Describe TODO
func (c *MetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	c.domainTimesListedMetric.Describe(ch)
	c.scrapesTotalMetric.Describe(ch)
	c.scrapeErrorsTotalMetric.Describe(ch)
	c.lastScrapeErrorMetric.Describe(ch)
	c.lastScrapeTimestampMetric.Describe(ch)
	c.lastScrapeDurationSecondsMetric.Describe(ch)
}

// Collect TODO
func (c *MetricsCollector) Collect(ch chan<- prometheus.Metric) {
	var begun = time.Now()

	errorMetric := float64(0)
	if err := c.reportMetrics(ch); err != nil {
		errorMetric = float64(1)
		c.scrapeErrorsTotalMetric.Inc()
		log.Printf("Error while getting dnsbl metrics: %s", err)
	}
	c.scrapeErrorsTotalMetric.Collect(ch)

	c.scrapesTotalMetric.Inc()
	c.scrapesTotalMetric.Collect(ch)

	c.lastScrapeErrorMetric.Set(errorMetric)
	c.lastScrapeErrorMetric.Collect(ch)

	c.lastScrapeTimestampMetric.Set(float64(time.Now().Unix()))
	c.lastScrapeTimestampMetric.Collect(ch)

	c.lastScrapeDurationSecondsMetric.Set(time.Since(begun).Seconds())
	c.lastScrapeDurationSecondsMetric.Collect(ch)
}

func (c *MetricsCollector) reportMetrics(ch chan<- prometheus.Metric) error {
	var begun = time.Now()
	log.Printf("Started checking %s against %d blacklists", yamlConfig.Domains, len(yamlConfig.Lists))
	for _, domain := range yamlConfig.Domains {
		for _, bl := range yamlConfig.Lists {
			res := Result{}
			query(bl, domain, &res)
			if res.Listed {
				c.domainTimesListedMetric.With(prometheus.Labels{"domain": domain, "blacklist": bl}).Set(1)
			} else {
				c.domainTimesListedMetric.With(prometheus.Labels{"domain": domain, "blacklist": bl}).Set(0)
			}
		}
	}
	c.domainTimesListedMetric.Collect(ch)
	log.Printf("Finished checking %s against %d blacklists in %f seconds", yamlConfig.Domains, len(yamlConfig.Lists), time.Since(begun).Seconds())
	return nil
}

func query(bl string, domain string, r *Result) {
	r.Listed = false
	r.Text = ""
	lookup := fmt.Sprintf("%s.%s", domain, bl)

	res, err := net.LookupHost(lookup)
	if len(res) > 0 {
		r.Listed = true
		txt, _ := net.LookupTXT(lookup)
		if len(txt) > 0 {
			r.Text = txt[0]
		}

	}
	if err != nil {
		r.Error = true
		r.ErrorType = err
	}
	return
}
