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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DepSpec defines the desired state of Dep
type DepSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Dep. Edit Dep_types.go to remove/update
	Foo    string `json:"foo,omitempty"`
	Detail string `json:"detail,omitempty"`
	Replicas *int32 `json:"replicas,omitempty"`
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
}

// DepStatus defines the observed state of Dep
type DepStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Created bool `json:"created,omitempty"`
}

// +kubebuilder:object:root=true

// Dep is the Schema for the deps API
type Dep struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DepSpec   `json:"spec,omitempty"`
	Status DepStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DepList contains a list of Dep
type DepList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dep `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dep{}, &DepList{})
}
