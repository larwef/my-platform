package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	myr53 "github.com/larwef/my-platform/dyn-dns/internal/dns/route53"
	"github.com/larwef/my-platform/dyn-dns/internal/poller/ipapi"
)

// Version injected at compile time.
var version = "No version provided"

var (
	awsRegion    = ""
	hostedZoneID = ""
	record       = ""
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err := realMain(ctx)
	done()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain(ctx context.Context) error {
	p := ipapi.New(nil)
	res, err := p.Poll()
	if err != nil {
		return err
	}
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(awsRegion))
	if err != nil {
		return err
	}
	r53 := route53.NewFromConfig(cfg)
	dns := myr53.New(r53, hostedZoneID)
	if err := dns.Update(ctx, record, res.String()); err != nil {
		return err
	}
	fmt.Printf("Updated %s to point to %s\n", record, res.String())
	return nil
}
