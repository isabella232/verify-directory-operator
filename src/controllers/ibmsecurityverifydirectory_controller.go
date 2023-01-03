/* vi: set ts=4 sw=4 noexpandtab : */

/*
 * Copyright contributors to the IBM Security Verify Directory Operator project
 */

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"

	ibmv1 "github.com/ibm-security/verify-directory-operator/api/v1"
)

// IBMSecurityVerifyDirectoryReconciler reconciles a IBMSecurityVerifyDirectory object
type IBMSecurityVerifyDirectoryReconciler struct {
	client.Client
	Log logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=ibm.com,resources=ibmsecurityverifydirectories,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ibm.com,resources=ibmsecurityverifydirectories/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ibm.com,resources=ibmsecurityverifydirectories/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the IBMSecurityVerifyDirectory object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *IBMSecurityVerifyDirectoryReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	r.Log.Info("Entering the reconcile loop....")

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *IBMSecurityVerifyDirectoryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ibmv1.IBMSecurityVerifyDirectory{}).
		Complete(r)
}
