package main

import (
	"context"
	"fmt"
	"net/url"
	presto_client "github.com/lyft/flyteplugins-private/go/tasks/presto/client"
	"github.com/lyft/flyteplugins-private/go/tasks/presto/config"
	
)

func main() {

	u, _ := url.Parse("https://prestoproxy-staging-internal.lyft.net:443")

	cfg := config.Config{
		Environment:               *u,
		Workers:                   1,
		AwsS3ShardFormatter:       "s3://lyft-modelbuilder/{}/",
		AwsS3ShardCount:           2,
	}
	type PrestoExecuteArgs struct {
		EnvironmentURL         url.URL
		RoutingGroup           string
		Catalog                string
		Schema                 string
		Source                 string
		AwsS3ShardFormatter    string
		AwsS3ShardStringLength int
	}
	 extraArgs := PrestoExecuteArgs{
		EnvironmentURL: *u,
		RoutingGroup:   "etl",
		Catalog:        "hive",
		Schema:         "city",
		Source:         "flyte-test-varshap",
	}
	fmt.Print("Hello world")

	ctx := context.Background()
	client := presto_client.NewPrestoClient(ctx, &cfg) // Is this correct? 

	x, _ := client.ExecuteCommand(ctx, `SELECT *
		FROM hive.city.fact_airport_sessions
		WHERE ds = '2019-07-21'
		LIMIT 10`, extraArgs)

	fmt.Print(x.(client.PrestoExecuteResponse))
}
