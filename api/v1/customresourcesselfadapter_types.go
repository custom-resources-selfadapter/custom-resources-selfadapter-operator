/*
Copyright 2025 The Custom Resources Self-Adapter Developers.

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
	autoscaling "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourcesSelfadapterConfig struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CustomResourcesSelfadapterSpec defines the desired state of CustomResourcesSelfadapter.
type CustomResourcesSelfadapterSpec struct {
	// The image of the Custom Pod Autoscaler
	Template PodTemplateSpec `json:"template"`
	// ScaleTargetRef defining what the Custom Pod Autoscaler should manage
	ScaleTargetRef autoscaling.CrossVersionObjectReference `json:"scaleTargetRef"`
	// Configuration options to be delivered as environment variables to the container
	Config                    []ResourcesSelfadapterConfig `json:"config,omitempty"`
	ProvisionRole             *bool                        `json:"provisionRole,omitempty"`
	ProvisionRoleBinding      *bool                        `json:"provisionRoleBinding,omitempty"`
	ProvisionServiceAccount   *bool                        `json:"provisionServiceAccount,omitempty"`
	ProvisionPod              *bool                        `json:"provisionPod,omitempty"`
	RoleRequiresMetricsServer *bool                        `json:"roleRequiresMetricsServer,omitempty"`
	RoleRequiresArgoRollouts  *bool                        `json:"roleRequiresArgoRollouts,omitempty"`
}

// CustomResourcesSelfadapterStatus defines the observed state of CustomResourcesSelfadapter.
type CustomResourcesSelfadapterStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName:crs

// CustomResourcesSelfadapter is the Schema for the customresourcesselfadapters API.
type CustomResourcesSelfadapter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomResourcesSelfadapterSpec   `json:"spec,omitempty"`
	Status CustomResourcesSelfadapterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomResourcesSelfadapterList contains a list of CustomResourcesSelfadapter.
type CustomResourcesSelfadapterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomResourcesSelfadapter `json:"items"`
}

type PodTemplateSpec struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	ObjectMeta PodMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the pod.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec PodSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`
}

// +kubebuilder:pruning:PreserveUnknownFields
type PodMeta metav1.ObjectMeta

type PodSpec corev1.PodSpec

func init() {
	SchemeBuilder.Register(&CustomResourcesSelfadapter{}, &CustomResourcesSelfadapterList{})
}
