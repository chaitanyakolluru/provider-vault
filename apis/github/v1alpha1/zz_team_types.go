/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TeamObservation struct {

	// Auth backend to which team mapping will be congigured.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Policies to be assigned to this team.
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// GitHub team name in "slugified" format.
	Team *string `json:"team,omitempty" tf:"team,omitempty"`
}

type TeamParameters struct {

	// Auth backend to which team mapping will be congigured.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Policies to be assigned to this team.
	// +kubebuilder:validation:Optional
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// GitHub team name in "slugified" format.
	// +kubebuilder:validation:Optional
	Team *string `json:"team,omitempty" tf:"team,omitempty"`
}

// TeamSpec defines the desired state of Team
type TeamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TeamParameters `json:"forProvider"`
}

// TeamStatus defines the observed state of Team.
type TeamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TeamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Team is the Schema for the Teams API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type Team struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.team)",message="team is a required parameter"
	Spec   TeamSpec   `json:"spec"`
	Status TeamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamList contains a list of Teams
type TeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Team `json:"items"`
}

// Repository type metadata.
var (
	Team_Kind             = "Team"
	Team_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Team_Kind}.String()
	Team_KindAPIVersion   = Team_Kind + "." + CRDGroupVersion.String()
	Team_GroupVersionKind = CRDGroupVersion.WithKind(Team_Kind)
)

func init() {
	SchemeBuilder.Register(&Team{}, &TeamList{})
}