package main

import (
        "fmt"
        "net/http"
	"flag"
//        "context"
//        "errors"
//        "flag"
//        "gopkg.in/yaml.v2"
//        "io/ioutil"
        "os"
//        "os/exec"
//        "regexp"
//        "time"

        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/common/log"
        "github.com/prometheus/common/version"
)


var (
// Command line parsing
	showVersion   = flag.Bool("version", false, "Print version information.")
	configFile    = flag.String("config.file", "helloworld-exporter.yml", "Hello World exporter configuration file.")
	listenAddress = flag.String("web.listen-address", ":8888", "The address to listen on for HTTP requests.")
	metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")

)

var (
	HelloWorld = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "HelloWorld",
			Help: "If you get this, means you have get the simplest exporter. Congrats!",
		})
)

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h3>METRICS</h3>")
}

func init() {
	prometheus.MustRegister(version.NewCollector("helloworld_exporter"))
	prometheus.MustRegister(HelloWorld)
}


func main() {
	flag.Parse()
	if  *showVersion {
		fmt.Fprintln(os.Stdout, version.Print("helloworld_exporter"))
		os.Exit(0)
	}
	HelloWorld.Set(1)
	log.Infoln("Starting helloworld_exporter", version.Info())

	http.Handle("/metrics", prometheus.Handler())

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
				<head><title>Hello World Exporter</title></head>
				<body>
	         		<h1>HelloWorld Exporter</h1>
				<p><a href="` + *metricsPath + `">Metrics</a></p>
		        	</body>
				</html>`))
	})


        log.Fatal(http.ListenAndServe(":8888", nil))

}
