package discovery_server

import (
	"context"
	"net"

	robotv1alpha1 "github.com/robolaunch/robot-operator/api/v1alpha1"
	robotErr "github.com/robolaunch/robot-operator/internal/error"
	"github.com/robolaunch/robot-operator/internal/resources"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *DiscoveryServerReconciler) reconcileUpdateConnectionInfo(ctx context.Context, instance *robotv1alpha1.DiscoveryServer) error {

	if instance.Status.Phase == robotv1alpha1.DiscoveryServerPhaseReady {
		dnsName := resources.GetDiscoveryServerDNS(*instance)

		ips, err := net.LookupIP(dnsName)
		if err != nil {
			return &robotErr.CannotResolveDiscoveryServerError{
				ResourceKind:      instance.Kind,
				ResourceName:      instance.Name,
				ResourceNamespace: instance.Namespace,
			}
		} else if len(ips) == 0 {
			return &robotErr.CannotResolveDiscoveryServerError{
				ResourceKind:      instance.Kind,
				ResourceName:      instance.Name,
				ResourceNamespace: instance.Namespace,
			}
		}

		instance.Status.ConnectionInfo.IP = ips[0].String()
		instance.Status.ConnectionInfo.ConfigMapName = instance.GetDiscoveryServerConfigMapMetadata().Name
	} else {
		instance.Status.ConnectionInfo = robotv1alpha1.ConnectionInfo{}
	}

	return nil
}

func (r *DiscoveryServerReconciler) reconcileCheckPod(ctx context.Context, instance *robotv1alpha1.DiscoveryServer) error {

	if instance.Spec.Attached {
		discoveryServerPodQuery := &corev1.Pod{}
		err := r.Get(ctx, *instance.GetDiscoveryServerPodMetadata(), discoveryServerPodQuery)
		if err != nil && errors.IsNotFound(err) {
			instance.Status.PodStatus.Created = false
		} else if err != nil {
			return err
		} else {

			if discoveryServerPodQuery.Spec.Hostname != instance.Spec.Hostname || discoveryServerPodQuery.Spec.Subdomain != instance.GetDiscoveryServerServiceMetadata().Name {
				err = r.Delete(ctx, discoveryServerPodQuery)
				if err != nil {
					return err
				}

				instance.Status.PodStatus.IP = ""
				instance.Status.PodStatus.Phase = ""
				instance.Status.Phase = ""

			} else {

				instance.Status.PodStatus.Created = true
				instance.Status.PodStatus.IP = discoveryServerPodQuery.Status.PodIP
				instance.Status.PodStatus.Phase = discoveryServerPodQuery.Status.Phase

			}

		}
	}

	return nil
}

func (r *DiscoveryServerReconciler) reconcileCheckService(ctx context.Context, instance *robotv1alpha1.DiscoveryServer) error {

	if instance.Spec.Attached {

		err := r.reconcileCleanupOwnedServices(ctx, instance)
		if err != nil {
			return err
		}

		discoveryServerServiceQuery := &corev1.Service{}
		err = r.Get(ctx, *instance.GetDiscoveryServerServiceMetadata(), discoveryServerServiceQuery)
		if err != nil && errors.IsNotFound(err) {
			instance.Status.ServiceStatus.Created = false
		} else if err != nil {
			return err
		} else {
			instance.Status.ServiceStatus.Created = true
		}

	}

	return nil
}

func (r *DiscoveryServerReconciler) reconcileCleanupOwnedServices(ctx context.Context, instance *robotv1alpha1.DiscoveryServer) error {

	// service cleanup
	services := &corev1.ServiceList{}
	err := r.List(ctx, services, &client.ListOptions{Namespace: instance.Namespace})
	if err != nil {
		return err
	}

	for _, svc := range services.Items {
		for _, owner := range svc.OwnerReferences {
			if owner.Name == instance.Name && svc.Name != instance.GetDiscoveryServerServiceMetadata().Name {
				err := r.Delete(ctx, &svc)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *DiscoveryServerReconciler) reconcileCheckConfigMap(ctx context.Context, instance *robotv1alpha1.DiscoveryServer) error {

	discoveryServerConfigMapQuery := &corev1.ConfigMap{}
	err := r.Get(ctx, *instance.GetDiscoveryServerConfigMapMetadata(), discoveryServerConfigMapQuery)
	if err != nil && errors.IsNotFound(err) {
		instance.Status.ConfigMapStatus.Created = false
	} else if err != nil {
		return err
	} else {

		if configuredIP, ok := discoveryServerConfigMapQuery.Labels["configuredIP"]; !ok {
			err := r.Delete(ctx, discoveryServerConfigMapQuery)
			if err != nil {
				return err
			}
		} else {
			if configuredIP != instance.Status.PodStatus.IP {
				err := r.Delete(ctx, discoveryServerConfigMapQuery)
				if err != nil {
					return err
				}
			}
		}

		instance.Status.ConfigMapStatus.Created = true
	}

	return nil
}