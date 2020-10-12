package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v2"
)

// YamlConfig parse
type YamlConfig struct {
	Domains []string `yaml:"dnsbl_check_domains"`
	Lists   []string `yaml:"dnsbl_lists"`
}

var (
	addr       = flag.String("listen-address", ":8881", "The address to listen on for HTTP requests.")
	yamlConfig = YamlConfig{}
)

func main() {
	configFile := flag.String("config", "default.yml", "config file")
	flag.Parse()
	yamlFile, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Printf("Error reading YAML file: %s\n", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		log.Printf("Error parsing YAML file: %s\n", err)
	}
	metricsCollector := NewMetricsCollector()
	prometheus.MustRegister(metricsCollector)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(*addr, nil))
}
