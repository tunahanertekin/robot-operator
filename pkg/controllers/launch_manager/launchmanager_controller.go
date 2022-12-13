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

package launch_manager

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"
	robotv1alpha1 "github.com/robolaunch/robot-operator/api/v1alpha1"
)

// LaunchManagerReconciler reconciles a LaunchManager object
type LaunchManagerReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	DynamicClient dynamic.Interface
}

//+kubebuilder:rbac:groups=robot.roboscale.io,resources=launchmanagers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=robot.roboscale.io,resources=launchmanagers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=robot.roboscale.io,resources=launchmanagers/finalizers,verbs=update

var logger logr.Logger

func (r *LaunchManagerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
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

func (r *LaunchManagerReconciler) reconcileCheckStatus(ctx context.Context, instance *robotv1alpha1.LaunchManager) error {

	switch !instance.Status.Active {
	case true:

		switch instance.Status.LaunchPodStatus.Created {
		case true:

			switch instance.Status.LaunchPodStatus.Phase {
			case v1.PodRunning:

				instance.Status.Phase = robotv1alpha1.LaunchManagerPhaseReady

			}

		case false:

			instance.Status.Phase = robotv1alpha1.LaunchManagerPhaseCreatingPod
			err := r.createLaunchPod(ctx, instance)
			if err != nil {
				return err
			}
			instance.Status.Phase = robotv1alpha1.LaunchManagerPhaseLaunching

		}

	case false:

		// delete resources

	}

	return nil
}

func (r *LaunchManagerReconciler) reconcileCheckResources(ctx context.Context, instance *robotv1alpha1.LaunchManager) error {

	err := r.reconcileCheckLaunchPod(ctx, instance)
	if err != nil {
		return err
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LaunchManagerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&robotv1alpha1.LaunchManager{}).
		Complete(r)
}
