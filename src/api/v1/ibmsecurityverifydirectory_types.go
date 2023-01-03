/* vi: set ts=4 sw=4 noexpandtab : */

/*
 * Copyright contributors to the IBM Security Verify Directory Operator project
 */

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IBMSecurityVerifyDirectorySpec defines the desired state of IBMSecurityVerifyDirectory
type IBMSecurityVerifyDirectorySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The name of the image which will be used in the deployment.
	// Cannot be updated.
	Image string `json:"image"`

	// Foo is an example field of IBMSecurityVerifyDirectory. Edit ibmsecurityverifydirectory_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// IBMSecurityVerifyDirectoryStatus defines the observed state of IBMSecurityVerifyDirectory
type IBMSecurityVerifyDirectoryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// IBMSecurityVerifyDirectory is the Schema for the ibmsecurityverifydirectories API
type IBMSecurityVerifyDirectory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IBMSecurityVerifyDirectorySpec   `json:"spec,omitempty"`
	Status IBMSecurityVerifyDirectoryStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IBMSecurityVerifyDirectoryList contains a list of IBMSecurityVerifyDirectory
type IBMSecurityVerifyDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IBMSecurityVerifyDirectory `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IBMSecurityVerifyDirectory{}, &IBMSecurityVerifyDirectoryList{})
}
