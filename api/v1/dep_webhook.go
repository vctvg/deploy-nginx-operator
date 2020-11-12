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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var deplog = logf.Log.WithName("dep-resource")

func (r *Dep) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-test-build-io-testbuilder-io-v1-dep,mutating=true,failurePolicy=fail,groups=test-build.io.testbuilder.io,resources=deps,verbs=create;update,versions=v1,name=mdep.kb.io

var _ webhook.Defaulter = &Dep{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Dep) Default() {
	deplog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	// set default replicas to 2
	if r.Spec.Replicas == nil{
		defaultReplicas := int32(2)
		r.Spec.Replicas = &defaultReplicas
	}

	// add default selector label
	LableMap := make(map[string]string, 1)
	LableMap["Dep"] = r.Name
	r.Spec.Selector = &metav1.LabelSelector{
		MatchLabels: LableMap,
	}
	
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-test-build-io-testbuilder-io-v1-dep,mutating=false,failurePolicy=fail,groups=test-build.io.testbuilder.io,resources=deps,versions=v1,name=vdep.kb.io

var _ webhook.Validator = &Dep{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Dep) ValidateCreate() error {
	deplog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Dep) ValidateUpdate(old runtime.Object) error {
	deplog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Dep) ValidateDelete() error {
	deplog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
