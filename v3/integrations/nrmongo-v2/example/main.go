// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"os"
	"time"

	"github.com/newrelic/go-agent/v3/integrations/nrmongo-v2"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Basic Mongo App"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		panic(err)
	}
	app.WaitForConnection(10 * time.Second)

	// If you have another CommandMonitor, you can pass it to NewCommandMonitor and it will get called along
	// with the NR monitor
	nrCmdMonitor := nrmongo.NewCommandMonitor(nil)
	ctx := context.Background()

	// nrCmdMonitor must be added after any other monitors are added, as previous options get overwritten.
	// This example assumes Mongo is running locally on port 27017
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017").SetMonitor(nrCmdMonitor))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	txn := app.StartTransaction("Mongo txn")
	// Make sure to add the newrelic.Transaction to the context
	nrCtx := newrelic.NewContext(context.Background(), txn)
	collection := client.Database("testing").Collection("numbers")
	_, err = collection.InsertOne(nrCtx, bson.M{"name": "exampleName", "value": "exampleValue"})
	if err != nil {
		panic(err)
	}
	txn.End()
	app.Shutdown(10 * time.Second)

}
