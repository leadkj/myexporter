package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"myexporter/collector"
	"net/http"
)

var addr = flag.String("listen-address",":8081","the address listen on for http requests")
var namespace = flag.String("metrics.namespace","LLHD","Prometheus metrices namespace")


func main() {
	flag.Parse()
	flag.Usage()
	metrics:=collector.NewMetrics(*namespace)
	registry:=prometheus.NewRegistry()
	registry.MustRegister(metrics)//要注册metrices,Metrice必须实现prometheus.Collector中的接口 Describe和Collect

	http.Handle("/metrics",promhttp.HandlerFor(registry,promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>A Prometheus Exporter</title></head>
			<body>
			<h1>A Prometheus Exporter</h1>
			<p><a href='/metrics'>Metrics</a></p>
			</body>
			</html>`))
	})
	log.Fatal(http.ListenAndServe(*addr,nil))

}
