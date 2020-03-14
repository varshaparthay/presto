package main

import (
	"context"
	"fmt"
	"net/url"
	"github.com/lyft/flyteplugins-private/go/tasks/plugins/presto"
)

func main() {

	u, _ := url.Parse("https://prestoproxy-staging-internal.lyft.net:443")
	
	cfg := config.Config{
		Environment:               *u,
		Workers:                   1,
		DestinationClusterConfigs: nil,
		AwsS3ShardFormatter:       "s3://lyft-modelbuilder/{}/",
		AwsS3ShardCount:           2,
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
	ctx = context.WithValue(ctx, client3.PrincipalContextKey, "varshap@lyft.com")

	NewPrestoClient()

	cli := client3.NewPrestoClient(&cfg)

	x, _ := cli.ExecuteCommand(ctx, `SELECT *
		FROM hive.city.fact_airport_sessions
		WHERE ds = '2019-07-21'
		LIMIT 10`, 3, extraArgs)

	fmt.Print(x.(client3.PrestoExecuteResponse))
}
