package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/kelseyhightower/envconfig"
	myr53 "github.com/larwef/my-platform/dyn-dns/internal/dns/route53"
	"github.com/larwef/my-platform/dyn-dns/internal/poller/ipapi"
	"github.com/larwef/my-platform/dyn-dns/internal/runner"
)

// Version injected at compile time.
var version = "No version provided"

type Config struct {
	AWSRegion    string        `envconfig:"DYNDNS_AWS_REGION" required:"true"`
	HostedZoneID string        `envconfig:"DYNDNS_HOSTED_ZONE_ID" required:"true"`
	RecordName   string        `envconfig:"DYNDNS_RECORD_NAME" required:"true"`
	PollInterval time.Duration `envconfig:"DYNDNS_POLL_INTERVAL" default:"5m"`
}

func main() {
	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err := realMain(ctx)
	done()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain(ctx context.Context) error {
	log.Printf("Starting dyn-dns %s\n", version)

	var conf Config
	if err := envconfig.Process("dyndns", &conf); err != nil {
		return err
	}

	awsConf, err := config.LoadDefaultConfig(ctx, config.WithRegion(conf.AWSRegion))
	if err != nil {
		return err
	}

	ipPoller := ipapi.New(nil)
	dnsUpdater := myr53.New(route53.NewFromConfig(awsConf), conf.HostedZoneID)
	runner := runner.New(ipPoller, dnsUpdater, conf.RecordName, conf.PollInterval)
	return runner.Run(ctx)
}
