package route53

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsroute53 "github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

type route53iface interface {
	ChangeResourceRecordSets(ctx context.Context, params *awsroute53.ChangeResourceRecordSetsInput, optFns ...func(*awsroute53.Options)) (*awsroute53.ChangeResourceRecordSetsOutput, error)
}

type Route53 struct {
	client       route53iface
	hostedZoneID string
}

func New(client route53iface, hostedZoneID string) *Route53 {
	return &Route53{
		client:       client,
		hostedZoneID: hostedZoneID,
	}
}

func (r *Route53) Update(ctx context.Context, name string, value string) error {
	_, err := r.client.ChangeResourceRecordSets(ctx, &awsroute53.ChangeResourceRecordSetsInput{
		ChangeBatch: &types.ChangeBatch{
			Changes: []types.Change{
				{
					Action: types.ChangeActionUpsert,
					ResourceRecordSet: &types.ResourceRecordSet{
						Name: &name,
						Type: types.RRTypeA,
						ResourceRecords: []types.ResourceRecord{
							{
								Value: &value,
							},
						},
						TTL: aws.Int64(60),
					},
				},
			},
		},
		HostedZoneId: &r.hostedZoneID,
	})
	if err != nil {
		return fmt.Errorf("route53: unable to update record: %w", err)
	}
	return nil
}
