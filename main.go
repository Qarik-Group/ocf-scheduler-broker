package main

import (
	"fmt"
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"

	"github.com/starkandwayne/ocf-scheduler-broker/broker"
	"github.com/starkandwayne/ocf-scheduler-broker/util"
)

func statusAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
func main() {
	logger := lager.NewLogger("ocf-scheduler-broker")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	logger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.ERROR))

	servicebroker := broker.NewSchedulerImpl(logger)

	brokerCredentials := brokerapi.BrokerCredentials{
		Username: util.EnvOr("AUTH_USER", "admin"),
		Password: util.EnvOr("AUTH_PASSWORD", "Test1234"),
	}

	brokerAPI := brokerapi.New(servicebroker, logger, brokerCredentials)

	http.HandleFunc("/health", statusAPI)

	http.Handle("/", brokerAPI)

	port := util.EnvOr("PORT", "3000")

	fmt.Println("\n\nStarting OCF Scheduler Service broker on 0.0.0.0:" + port)
	logger.Fatal("http-listen", http.ListenAndServe("0.0.0.0:"+port, nil))
}
