/*
Copyright 2020 The Flux CD contributors.

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

package v1alpha1

// CrossNamespaceObjectReference contains enough information to let you locate the
// typed referenced object at cluster level
type CrossNamespaceObjectReference struct {
	// API version of the referent, defaults to 'apps/v1'
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind of the referent
	// +required
	Kind string `json:"kind"`

	// Name of the referent
	// +required
	Name string `json:"name"`

	// Namespace of the referent
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// CrossNamespaceSourceReference contains enough information to let you locate the
// typed referenced object at cluster level
type CrossNamespaceSourceReference struct {
	// API version of the referent
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind of the referent
	// +kubebuilder:validation:Enum=GitRepository
	// +required
	Kind string `json:"kind"`

	// Name of the referent
	// +required
	Name string `json:"name"`

	// Namespace of the referent, defaults to the Kustomization namespace
	// +optional
	Namespace string `json:"namespace,omitempty"`
}
