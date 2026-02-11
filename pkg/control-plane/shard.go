package controlplane

import (
	"context"
	"fmt"

	corev1alpha1 "github.com/stateless-pg/stateless-pg/pkg/api/v1alpha1"
)

func (o *Operator) createOrUpdateShard(ctx context.Context, tenant *corev1alpha1.Tenant) error {
    for _, cnt := range tenant.Spec.ShardParameters.Count {
		tenantId := fmt.Sprintf("%s-%s-%s-%s", tenant.Spec.NeonClusterRef.Namespace,tenant.Spec.NeonClusterRef.Name, tenant.Namespace, tenant.Name )
		shardName := fmt.Sprintf("%s-%d", tenantId, cnt)
		shard := &corev1alpha1.TenantShard{
			ObjectMeta: metav1.ObjectMeta{
				Name:      shardName,
				Namespace: tenant.Namespace,
				Labels: map[string]string{
					"tenant": tenant.Name,
				},
			},
			Spec: corev1alpha1.TenantShardSpec{
				
			},
		}
	}
}
