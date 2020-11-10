/*


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

package controllers

import (
	"context"
	//"golang.org/x/tools/godoc/vfs"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	testbuildiov1 "test.io/api/v1"
)

// DepReconciler reconciles a Dep object
type DepReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=test-build.io.testbuilder.io,resources=deps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=test-build.io.testbuilder.io,resources=deps/status,verbs=get;update;patch

func (r *DepReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("dep", req.NamespacedName)

	// your logic here
	// init dep
	dep := testbuildiov1.Dep{}
	err := r.Get(ctx, req.NamespacedName, &dep)
	if err != nil {
		log.Error(err, "cannot found")
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
	}

	// setup finalizer
	testFinalizerName := "testfinailizer"
	if dep.ObjectMeta.DeletionTimestamp.IsZero(){
		if !containsString(dep.ObjectMeta.Finalizers, testFinalizerName){
			dep.ObjectMeta.Finalizers = append(dep.ObjectMeta.Finalizers, testFinalizerName)
			err := r.Update(ctx, &dep)
			if err != nil{
				return ctrl.Result{}, err
			}
		}

	} else {
		if containsString(dep.ObjectMeta.Finalizers, testFinalizerName){
			err := r.PreDelete(&dep)
			if err != nil{
				return ctrl.Result{}, err
			}
			dep.ObjectMeta.Finalizers = removeString(dep.ObjectMeta.Finalizers, testFinalizerName)
			err = r.Update(ctx, &dep)
			if err != nil{
				return ctrl.Result{}, err
			}
		}
		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}


	// change status
	if !dep.Status.Created {
		dep.Status.Created = true
		r.Update(ctx, &dep)
	}

	// deploy nginx
	deploynginxvar := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dep.Name + "nginx-deployment",
			Namespace: dep.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": dep.Name + "nginx"},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": dep.Name + "nginx"},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "nginx",
							Image: "nginx",
						},
					},
				},
			},
		},
	}

	if err := controllerutil.SetControllerReference(&dep, deploynginxvar, r.Scheme); err != nil {
		log.Error(err, "failed set reference")
		return reconcile.Result{}, err
	}

	err = r.Create(ctx, deploynginxvar)
	if err != nil {
		log.Error(err, "create failed")
		return reconcile.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *DepReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testbuildiov1.Dep{}).
		Complete(r)
}

func containsString(slice []string, s string) bool {
	for _, item := range slice{
		if item == s {
			return true
		}
	}
	return false
}

func (r *DepReconciler)PreDelete(newresource *testbuildiov1.Dep) error{
	// predelete logic

	return nil
}

func removeString(slice []string, s string)(result []string) {
	for _, item := range (slice) {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
