/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package relay_server

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"
	robotv1alpha1 "github.com/robolaunch/robot-operator/pkg/api/roboscale.io/v1alpha1"
)

// RelayServerReconciler reconciles a RelayServer object
type RelayServerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=robot.roboscale.io,resources=relayservers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=robot.roboscale.io,resources=relayservers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=robot.roboscale.io,resources=relayservers/finalizers,verbs=update

//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete

var logger logr.Logger

func (r *RelayServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger = log.FromContext(ctx)

	instance, err := r.reconcileGetInstance(ctx, req.NamespacedName)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	err = r.reconcileCheckStatus(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileUpdateInstanceStatus(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileCheckResources(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	err = r.reconcileUpdateInstanceStatus(ctx, instance)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *RelayServerReconciler) reconcileCheckStatus(ctx context.Context, instance *robotv1alpha1.RelayServer) error {

	switch instance.Status.PodStatus.Created {
	case true:

		switch instance.Status.PodStatus.Phase {
		case string(corev1.PodRunning):

			switch instance.Status.ServiceStatus.Resource.Created {
			case true:

				switch instance.Status.IngressStatus.Created {
				case true:

					instance.Status.Phase = robotv1alpha1.RelayServerPhaseReady

				case false:

					instance.Status.Phase = robotv1alpha1.RelayServerPhaseCreatingIngress
					err := r.createIngress(ctx, instance, instance.GetRelayServerIngressMetadata())
					if err != nil {
						return err
					}
					instance.Status.IngressStatus.Created = true

				}

			case false:

				instance.Status.Phase = robotv1alpha1.RelayServerPhaseCreatingService
				err := r.createService(ctx, instance, instance.GetRelayServerServiceMetadata())
				if err != nil {
					return err
				}
				instance.Status.ServiceStatus.Resource.Created = true

			}

		}

	case false:

		instance.Status.Phase = robotv1alpha1.RelayServerPhaseCreatingPod
		err := r.createPod(ctx, instance, instance.GetRelayServerPodMetadata())
		if err != nil {
			return err
		}
		instance.Status.PodStatus.Created = true

	}

	return nil
}

func (r *RelayServerReconciler) reconcileCheckResources(ctx context.Context, instance *robotv1alpha1.RelayServer) error {

	err := r.reconcileCheckPod(ctx, instance)
	if err != nil {
		return err
	}

	err = r.reconcileCheckService(ctx, instance)
	if err != nil {
		return err
	}

	err = r.reconcileCheckIngress(ctx, instance)
	if err != nil {
		return err
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RelayServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&robotv1alpha1.RelayServer{}).
		Owns(&corev1.Pod{}).
		Owns(&corev1.Service{}).
		Owns(&networkingv1.Ingress{}).
		Complete(r)
}