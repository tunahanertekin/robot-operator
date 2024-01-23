package ros_bridge

import (
	"context"

	robotErr "github.com/robolaunch/robot-operator/internal/error"
	"github.com/robolaunch/robot-operator/internal/label"
	robotv1alpha1 "github.com/robolaunch/robot-operator/pkg/api/roboscale.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *ROSBridgeReconciler) reconcileGetInstance(ctx context.Context, meta types.NamespacedName) (*robotv1alpha1.ROSBridge, error) {
	instance := &robotv1alpha1.ROSBridge{}
	err := r.Get(ctx, meta, instance)
	if err != nil {
		return &robotv1alpha1.ROSBridge{}, err
	}

	return instance, nil
}

func (r *ROSBridgeReconciler) reconcileUpdateInstanceStatus(ctx context.Context, instance *robotv1alpha1.ROSBridge) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		instanceLV := &robotv1alpha1.ROSBridge{}
		err := r.Get(ctx, types.NamespacedName{
			Name:      instance.Name,
			Namespace: instance.Namespace,
		}, instanceLV)

		if err == nil {
			instance.ResourceVersion = instanceLV.ResourceVersion
		}

		err1 := r.Status().Update(ctx, instance)
		return err1
	})
}

func (r *ROSBridgeReconciler) reconcileGetOwner(ctx context.Context, instance *robotv1alpha1.ROSBridge) (*robotv1alpha1.Robot, error) {
	robot := &robotv1alpha1.Robot{}
	err := r.Get(ctx, *instance.GetOwnerMetadata(), robot)
	if err != nil {
		return nil, err
	}

	return robot, nil
}

func (r *ROSBridgeReconciler) reconcileCheckNode(ctx context.Context, instance *robotv1alpha1.Robot) (*corev1.Node, error) {

	tenancyMap := label.GetTenancyMap(instance)

	requirements := []labels.Requirement{}
	for k, v := range tenancyMap {
		newReq, err := labels.NewRequirement(k, selection.In, []string{v})
		if err != nil {
			return nil, err
		}
		requirements = append(requirements, *newReq)
	}

	nodeSelector := labels.NewSelector().Add(requirements...)

	nodes := &corev1.NodeList{}
	err := r.List(ctx, nodes, &client.ListOptions{
		LabelSelector: nodeSelector,
	})
	if err != nil {
		return nil, err
	}

	if len(nodes.Items) == 0 {
		return nil, &robotErr.NodeNotFoundError{
			ResourceKind:      instance.Kind,
			ResourceName:      instance.Name,
			ResourceNamespace: instance.Namespace,
		}
	} else if len(nodes.Items) > 1 {
		return nil, &robotErr.MultipleNodeFoundError{
			ResourceKind:      instance.Kind,
			ResourceName:      instance.Name,
			ResourceNamespace: instance.Namespace,
		}
	}

	instance.Status.NodeName = nodes.Items[0].Name

	return &nodes.Items[0], nil
}
