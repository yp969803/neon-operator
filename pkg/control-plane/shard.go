package controlplane

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "github.com/stateless-pg/stateless-pg/pkg/api/v1alpha1"
	"github.com/stateless-pg/stateless-pg/pkg/operator"
)

func (o *Operator) createOrUpdateShard(ctx context.Context, tenant *corev1alpha1.Tenant) error {
    for cnt := range tenant.Spec.ShardParameters.Count {
		tenantId := fmt.Sprintf("%s-%s-%s-%s", tenant.Spec.NeonClusterRef.Namespace,tenant.Spec.NeonClusterRef.Name, tenant.Namespace, tenant.Name )
		shardName := fmt.Sprintf("%s-%d", tenantId, cnt)
		shard := &corev1alpha1.TenantShard{
			ObjectMeta: metav1.ObjectMeta{
				Name:      shardName,
				Namespace: tenant.Namespace,
			},
			Spec: corev1alpha1.TenantShardSpec{
				ShardId: corev1alpha1.TenantShardId{
					TenantId:    tenantId,
					ShardNumber: cnt,
					ShardCount:  tenant.Spec.ShardParameters.Count,
				},
				Identity: corev1alpha1.ShardIdentity{
					Number: cnt,
					Count: tenant.Spec.ShardParameters.Count,
				},
				Policy: tenant.Spec.PlacementPolicy,
				Config: &tenant.Spec.Config,
			},
		}
		
		operator.UpdateObject(shard,
			operator.WithOwner(tenant),
			operator.WithLabels(map[string]string{
				"tenant": tenant.Name,
			}),
		)
	}
	return nil
}
