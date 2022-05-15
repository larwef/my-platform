package route53

import (
	"context"
	"errors"
	"fmt"
	"testing"

	awsroute53 "github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/stretchr/testify/assert"
)

type route53Mock struct {
	changeResourceRecordSets func(ctx context.Context, params *awsroute53.ChangeResourceRecordSetsInput, optFns ...func(*awsroute53.Options)) (*awsroute53.ChangeResourceRecordSetsOutput, error)
}

func (r *route53Mock) ChangeResourceRecordSets(ctx context.Context, params *awsroute53.ChangeResourceRecordSetsInput, optFns ...func(*awsroute53.Options)) (*awsroute53.ChangeResourceRecordSetsOutput, error) {
	return r.changeResourceRecordSets(ctx, params, optFns...)
}

func TestRoute53_Update(t *testing.T) {
	type fields struct {
		client       route53iface
		hostedZoneID string
	}
	type args struct {
		ctx  context.Context
		name string
		ip   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Test successful",
			fields: fields{
				client: &route53Mock{
					changeResourceRecordSets: func(ctx context.Context, params *awsroute53.ChangeResourceRecordSetsInput, optFns ...func(*awsroute53.Options)) (*awsroute53.ChangeResourceRecordSetsOutput, error) {
						return nil, nil
					},
				},
				hostedZoneID: "hostedZoneID",
			},
			args: args{
				ctx:  context.Background(),
				name: "foo.example.com",
				ip:   "1.2.3.4",
			},
			wantErr: nil,
		},
		{
			name: "Test error",
			fields: fields{
				client: &route53Mock{
					changeResourceRecordSets: func(ctx context.Context, params *awsroute53.ChangeResourceRecordSetsInput, optFns ...func(*awsroute53.Options)) (*awsroute53.ChangeResourceRecordSetsOutput, error) {
						return nil, errors.New("some error")
					},
				},
				hostedZoneID: "hostedZoneID",
			},
			args: args{
				ctx:  context.Background(),
				name: "foo.example.com",
				ip:   "1.2.3.4",
			},
			wantErr: fmt.Errorf("route53: unable to update record: %w", errors.New("some error")),
		},
	}
	for _, test := range tests {
		dns := New(test.fields.client, test.fields.hostedZoneID)
		err := dns.Update(test.args.ctx, test.args.name, test.args.ip)
		assert.Equal(t, test.wantErr, err)
	}
}
