package main

import (
	"flag"

	"k8s.io/klog"

	"github.com/nevercase/k8s-controller-custom-resource/pkg/signals"
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()
	klog.Info("Starting Service")
	<-stopCh
	klog.Info("Shutdown Service")
}
