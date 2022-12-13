package launch_manager

import (
	"context"

	robotv1alpha1 "github.com/robolaunch/robot-operator/api/v1alpha1"
	"github.com/robolaunch/robot-operator/internal"
	robotErr "github.com/robolaunch/robot-operator/internal/error"
	"github.com/robolaunch/robot-operator/internal/label"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/retry"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *LaunchManagerReconciler) reconcileGetInstance(ctx context.Context, meta types.NamespacedName) (*robotv1alpha1.LaunchManager, error) {
	instance := &robotv1alpha1.LaunchManager{}
	err := r.Get(ctx, meta, instance)
	if err != nil {
		return &robotv1alpha1.LaunchManager{}, err
	}

	return instance, nil
}

func (r *LaunchManagerReconciler) reconcileUpdateInstanceStatus(ctx context.Context, instance *robotv1alpha1.LaunchManager) error {
	return retry.RetryOnConflict(retry.DefaultRetry, func() error {
		instanceLV := &robotv1alpha1.LaunchManager{}
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

func (r *LaunchManagerReconciler) reconcileGetTargetRobot(ctx context.Context, instance *robotv1alpha1.LaunchManager) (*robotv1alpha1.Robot, error) {
	robot := &robotv1alpha1.Robot{}
	err := r.Get(ctx, types.NamespacedName{
		Namespace: instance.Namespace,
		Name:      label.GetTargetRobot(instance),
	}, robot)
	if err != nil {
		return nil, err
	}

	return robot, nil
}

func (r *LaunchManagerReconciler) reconcileCheckTargetRobot(ctx context.Context, instance *robotv1alpha1.LaunchManager) error {
	robot, err := r.reconcileGetTargetRobot(ctx, instance)
	if err != nil {
		return err
	}

	if robot.Status.AttachedLaunchObject.Reference.Kind == instance.Kind && robot.Status.AttachedLaunchObject.Reference.Name == instance.Name {
		instance.Status.Active = true
	} else {
		instance.Status.Active = false
	}

	return nil
}

func (r *LaunchManagerReconciler) reconcileGetCurrentBuildManager(ctx context.Context, instance *robotv1alpha1.LaunchManager) (*robotv1alpha1.BuildManager, error) {
	robot, err := r.reconcileGetTargetRobot(ctx, instance)
	if err != nil {
		return nil, err
	}

	buildManager := &robotv1alpha1.BuildManager{}
	err = r.Get(ctx, types.NamespacedName{
		Namespace: robot.Status.AttachedBuildObject.Reference.Namespace,
		Name:      robot.Status.AttachedBuildObject.Reference.Name,
	}, buildManager)
	if err != nil {
		return nil, err
	}

	return buildManager, nil
}

func (r *LaunchManagerReconciler) reconcileCheckOtherAttachedResources(ctx context.Context, instance *robotv1alpha1.LaunchManager) error {

	if instance.Status.Active {
		// Get attached launch manager objects for this robot
		requirements := []labels.Requirement{}
		newReq, err := labels.NewRequirement(internal.TARGET_ROBOT, selection.In, []string{label.GetTargetRobot(instance)})
		if err != nil {
			return err
		}
		requirements = append(requirements, *newReq)

		robotSelector := labels.NewSelector().Add(requirements...)

		launchManagerList := robotv1alpha1.LaunchManagerList{}
		err = r.List(ctx, &launchManagerList, &client.ListOptions{Namespace: instance.Namespace, LabelSelector: robotSelector})
		if err != nil {
			return err
		}

		for _, lm := range launchManagerList.Items {

			if lm.Name == instance.Name {
				continue
			}

			if lm.Status.Active == true {
				return &robotErr.RobotResourcesHasNotBeenReleasedError{
					ResourceKind:      instance.Kind,
					ResourceName:      instance.Name,
					ResourceNamespace: instance.Namespace,
				}
			}

			if lm.Status.Phase != robotv1alpha1.LaunchManagerPhaseInactive {
				return &robotErr.RobotResourcesHasNotBeenReleasedError{
					ResourceKind:      instance.Kind,
					ResourceName:      instance.Name,
					ResourceNamespace: instance.Namespace,
				}
			}
		}
	}

	return nil
}
