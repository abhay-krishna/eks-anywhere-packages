// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PackageControllerSpec defines the desired state of PackageController
type PackageControllerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Optional log level for packages controller
	// +optional
	LogLevel *int32 `json:"logLevel,omitempty"`

	// UpgradeCheckInterval is the time between upgrade checks.
	//
	// The format is that of time's ParseDuration.
	UpgradeCheckInterval metav1.Duration `json:"upgradeCheckInterval,omitempty"`
}

// PackageControllerStatus defines the observed state of PackageController
type PackageControllerStatus struct { // INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PackageController is the Schema for the packagecontrollers API
type PackageController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PackageControllerSpec   `json:"spec,omitempty"`
	Status PackageControllerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PackageControllerList contains a list of PackageController
type PackageControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PackageController `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PackageController{}, &PackageControllerList{})
}
