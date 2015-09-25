/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//This is a utility for prometheus pushing functionality.
package e2e

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// Prometheus stuff: Setup metrics.
var podGauge = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "e2e_density_pod_count",
		Help: "The number of pods, broken out by state",
	},
	[]string{"state"},
)

// Turn this to true after we register.
var prom_registered = false

// Reusable function for pushing metrics to prometheus.  Handles initialization and so on.
func promPushRunningPending(running, pending, waiting, inactive, terminating, unknown int) error {
	if testContext.PrometheusPushGateway == "" {
		return nil
	} else {
		// Register metrics if necessary
		if !prom_registered && testContext.PrometheusPushGateway != "" {
			prometheus.Register(podGauge)
			prom_registered = true
		}
		// Update metric values
		podGauge.WithLabelValues("running").Set(float64(running))
		podGauge.WithLabelValues("pending").Set(float64(pending))
		podGauge.WithLabelValues("waiting").Set(float64(waiting))
		podGauge.WithLabelValues("inactive").Set(float64(inactive))
		podGauge.WithLabelValues("terminating").Set(float64(terminating))
		podGauge.WithLabelValues("unknown").Set(float64(unknown))

		// Push them to the push gateway.  This will be scraped by prometheus
		// provided you launch it with the pushgateway as an endpoint.
		if err := prometheus.Push(
			"e2e",
			"none",
			testContext.PrometheusPushGateway, //i.e. "127.0.0.1:9091"
		); err != nil {
			fmt.Println("failed at pushing to pushgateway ", err)
			return err
		}
	}
	return nil
}
