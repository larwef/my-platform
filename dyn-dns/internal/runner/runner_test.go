package runner

import (
	"context"
	"errors"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type ipPollerMock struct {
	poll func() (net.IP, error)
}

func (i *ipPollerMock) Poll() (net.IP, error) {
	return i.poll()
}

type resolverMock struct {
	lookupIP func(ctx context.Context, network, name string) ([]net.IP, error)
}

func (r *resolverMock) LookupIP(ctx context.Context, network, name string) ([]net.IP, error) {
	return r.lookupIP(ctx, network, name)
}

type dnsUpdaterMock struct {
	update func(ctx context.Context, name string, value string) error
}

func (d *dnsUpdaterMock) Update(ctx context.Context, name string, value string) error {
	return d.update(ctx, name, value)
}

func TestRunner_run(t *testing.T) {
	type fields struct {
		poller       IPPoller
		resolver     ipResolver
		dnsUpdater   DNSUpdater
		recordName   string
		pollInterval time.Duration
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "First poll successfull",
			fields: fields{
				poller: &ipPollerMock{
					poll: func() (net.IP, error) { return net.ParseIP("1.2.3.4"), nil },
				},
				resolver: &resolverMock{
					lookupIP: func(ctx context.Context, network, name string) ([]net.IP, error) {
						assert.Equal(t, "foo.bar.com", name)
						return nil, &net.DNSError{IsNotFound: true}
					},
				},
				dnsUpdater: &dnsUpdaterMock{
					update: func(ctx context.Context, name string, value string) error { return nil },
				},
				recordName: "foo.bar.com",
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: nil,
		},
		{
			name: "First already set",
			fields: fields{
				poller: &ipPollerMock{
					poll: func() (net.IP, error) { return net.ParseIP("1.2.3.4"), nil },
				},
				resolver: &resolverMock{
					lookupIP: func(ctx context.Context, network, name string) ([]net.IP, error) {
						assert.Equal(t, "foo.bar.com", name)
						return []net.IP{net.ParseIP("1.2.3.4")}, nil
					},
				},
				dnsUpdater: &dnsUpdaterMock{
					update: func(ctx context.Context, name string, value string) error {
						return errors.New("update function should not be called in this test")
					},
				},
				recordName: "foo.bar.com",
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: nil,
		},
		{
			name: "Error from poller",
			fields: fields{
				poller: &ipPollerMock{
					poll: func() (net.IP, error) { return nil, errors.New("poller error") },
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: fmt.Errorf("unable to poll current IP: %w", errors.New("poller error")),
		},
		{
			name: "Error from lookup",
			fields: fields{
				poller: &ipPollerMock{
					poll: func() (net.IP, error) { return net.ParseIP("1.2.3.4"), nil },
				},
				resolver: &resolverMock{
					lookupIP: func(ctx context.Context, network, name string) ([]net.IP, error) {
						return nil, errors.New("lookup error")
					},
				},
				recordName: "foo.bar.com",
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: fmt.Errorf("unable to lookup IP for foo.bar.com: %w", errors.New("lookup error")),
		},
		{
			name: "Error from update",
			fields: fields{
				poller: &ipPollerMock{
					poll: func() (net.IP, error) { return net.ParseIP("1.2.3.4"), nil },
				},
				resolver: &resolverMock{
					lookupIP: func(ctx context.Context, network, name string) ([]net.IP, error) {
						assert.Equal(t, "foo.bar.com", name)
						return []net.IP{net.ParseIP("4.3.2.1")}, nil
					},
				},
				dnsUpdater: &dnsUpdaterMock{
					update: func(ctx context.Context, name string, value string) error { return errors.New("update error") },
				},
				recordName: "foo.bar.com",
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: fmt.Errorf("unable to update DNS record for foo.bar.com: %w", errors.New("update error")),
		},
	}

	for _, test := range tests {
		r := &Runner{
			poller:       test.fields.poller,
			resolver:     test.fields.resolver,
			dnsUpdater:   test.fields.dnsUpdater,
			recordName:   test.fields.recordName,
			pollInterval: test.fields.pollInterval,
		}
		err := r.run(test.args.ctx)
		assert.Equal(t, test.wantErr, err)
	}
}
