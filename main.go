//https://www.youtube.com/watch?v=51WetEt_G4c

package main

import (
	"context"

	"github.com/uber-go/tally"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()

	if err != nil {
		panic(err)
	}

	logger.Info("Zap logger created")

	scope := tally.NoopScope

	serviceClient, err := client.NewClient(client.Options{
		HostPort:     client.DefaultHostPort,
		MetricsScope: scope,
	})

	if err != nil {
		logger.Fatal("Unable to start worker", zap.Error(err))
	}

	worker := worker.New(serviceClient, "worker_tutorial_go_skd", worker.Options{})

	worker.RegisterWorkflow(MyWorkflow)
	worker.RegisterActivity(MyActivity)

	err = worker.Start()
	if err != nil {
		logger.Fatal("Unable to start worker", zap.Error(err))
	}

	select {}
}

func MyActivity(ctx context.Context) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity MyActivity called")
	return nil
}

func MyWorkflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow MyWorkflow started")
	return nil
}
