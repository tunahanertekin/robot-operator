package robot

import (
	"context"

	robotv1alpha1 "github.com/robolaunch/robot-operator/api/v1alpha1"
	"github.com/robolaunch/robot-operator/internal/node"
	"github.com/robolaunch/robot-operator/internal/resources"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *RobotReconciler) createPVC(ctx context.Context, instance *robotv1alpha1.Robot, pvcNamespacedName *types.NamespacedName) error {

	pvc := resources.GetPersistentVolumeClaim(instance, pvcNamespacedName)

	err := ctrl.SetControllerReference(instance, pvc, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, pvc)
	if err != nil {
		return err
	}

	logger.Info("STATUS: PVC " + pvc.Name + " is created.")
	return nil
}

func (r *RobotReconciler) createDiscoveryServer(ctx context.Context, instance *robotv1alpha1.Robot, dsNamespacedName *types.NamespacedName) error {

	discoveryServer := resources.GetDiscoveryServer(instance, dsNamespacedName)

	err := ctrl.SetControllerReference(instance, discoveryServer, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, discoveryServer)
	if err != nil {
		return err
	}

	logger.Info("STATUS: Discovery server " + discoveryServer.Name + " is created.")
	return nil
}

func (r *RobotReconciler) createJob(ctx context.Context, instance *robotv1alpha1.Robot, jobNamespacedName *types.NamespacedName) error {

	activeNode, err := r.reconcileCheckNode(ctx, instance)
	if err != nil {
		return err
	}

	job := resources.GetLoaderJob(instance, jobNamespacedName, node.HasGPU(*activeNode))

	err = ctrl.SetControllerReference(instance, job, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, job)
	if err != nil {
		return err
	}

	logger.Info("STATUS: Job " + job.Name + " is created.")
	return nil
}

func (r *RobotReconciler) createROSBridge(ctx context.Context, instance *robotv1alpha1.Robot, bridgeNamespacedName *types.NamespacedName) error {

	rosBridge := resources.GetROSBridge(instance, bridgeNamespacedName)

	err := ctrl.SetControllerReference(instance, rosBridge, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, rosBridge)
	if err != nil {
		return err
	}

	logger.Info("STATUS: ROS bridge " + rosBridge.Name + " is created.")
	return nil
}

func (r *RobotReconciler) createRobotDevSuite(ctx context.Context, instance *robotv1alpha1.Robot, rdsNamespacedName *types.NamespacedName) error {

	robotDevSuite := resources.GetRobotDevSuite(instance, rdsNamespacedName)

	err := ctrl.SetControllerReference(instance, robotDevSuite, r.Scheme)
	if err != nil {
		return err
	}

	err = r.Create(ctx, robotDevSuite)
	if err != nil {
		return err
	}

	logger.Info("STATUS: Robot dev suite " + robotDevSuite.Name + " is created.")
	return nil
}