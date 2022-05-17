package runner

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

type IPPoller interface {
	Poll() (net.IP, error)
}

type DNSUpdater interface {
	Update(ctx context.Context, name string, value string) error
}

type ipResolver interface {
	LookupIP(ctx context.Context, network, name string) ([]net.IP, error)
}

type Runner struct {
	poller       IPPoller
	resolver     ipResolver
	dnsUpdater   DNSUpdater
	recordName   string
	pollInterval time.Duration
}

func New(ipPoller IPPoller, dnsUpdater DNSUpdater, recordName string, pollInterval time.Duration) *Runner {
	return &Runner{
		poller:       ipPoller,
		resolver:     &net.Resolver{},
		dnsUpdater:   dnsUpdater,
		recordName:   recordName,
		pollInterval: pollInterval,
	}
}

func (r *Runner) Run(ctx context.Context) error {
	ticker := time.NewTicker(r.pollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			if err := r.run(ctx); err != nil {
				return fmt.Errorf("runner: %w", err)
			}
		}
	}
}

func (r *Runner) run(ctx context.Context) error {
	currentIP, err := r.poller.Poll()
	if err != nil {
		return fmt.Errorf("unable to poll current IP: %w", err)
	}
	log.Printf("Polled IP: %s\n", currentIP)
	ips, err := r.resolver.LookupIP(ctx, "ip", r.recordName)
	if err != nil {
		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
			log.Printf("DNS record not found for %s. Creating new record.", r.recordName)
		} else {
			return fmt.Errorf("unable to lookup IP for %s: %w", r.recordName, err)
		}
	}
	for _, ip := range ips {
		if ip.Equal(currentIP) {
			log.Printf("IP %s is already set for %s. Skipping.", currentIP, r.recordName)
			return nil
		}
	}
	err = r.dnsUpdater.Update(ctx, r.recordName, currentIP.String())
	if err != nil {
		return fmt.Errorf("unable to update DNS record for %s: %w", r.recordName, err)
	}
	log.Printf("Successfully updated DNS record for %s to %s\n", r.recordName, currentIP)
	return nil
}
