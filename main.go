package main

import (
	"fmt"
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"example.com/broker"
	"github.com/pivotal-cf/brokerapi"
)

func statusAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
func main() {
	logger := lager.NewLogger("worlds-simplest-service-broker")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))
	logger.RegisterSink(lager.NewWriterSink(os.Stderr, lager.ERROR))

	servicebroker := broker.NewSchedulerImpl(logger)

	brokerCredentials := brokerapi.BrokerCredentials{
		Username: "admin",
		Password: "Test1234",
	}

	brokerAPI := brokerapi.New(servicebroker, logger, brokerCredentials)

	http.HandleFunc("/health", statusAPI)
	http.Handle("/", brokerAPI)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Println("\n\nStarting OCF Scheduler Service broker on 0.0.0.0:" + port)
	logger.Fatal("http-listen", http.ListenAndServe("0.0.0.0:"+port, nil))
}
