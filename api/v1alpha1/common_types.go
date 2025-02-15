/*
Copyright 2023. projectsveltos.io. All rights reserved.

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

import (
	"errors"

	corev1 "k8s.io/api/core/v1"
)

const (
	// PolicyTemplateAnnotation is the annotation that must be set on a policy when the
	// policy is a template and needs variable sustitution.
	PolicyTemplateAnnotation = "projectsveltos.io/template"
)

const (
	// DeployedBySveltosAnnotation is an annotation Sveltos adds to
	// EventSource/HealthCheck/Classifier/ResourceSummary instances deployed
	// by sveltos in managed clusters. Those resources, once deployed in a
	// managed cluster, are evaluated by sveltos services (sveltos-agent and
	// drift-detection-manager) running in the managed cluster
	DeployedBySveltosAnnotation = "projectsveltos.io/deployed-by-sveltos"
)

const (
	// ServiceAccountNameLabel can be set on various Sveltos resources (ClusterProfile/EventSource/...)
	// to indicate which admin (represented by a ServiceAccount) is creating it (service account name).
	// ServiceAccountNameLabel used along with RoleRequest is Sveltos solution for multi tenancy.
	ServiceAccountNameLabel = "projectsveltos.io/serviceaccount-name"

	// ServiceAccountNamespaceLabel can be set on various Sveltos resources (ClusterProfile/EventSource/...)
	// to indicate which admin (represented by a ServiceAccount) is creating it (service account namespace).
	// ServiceAccountNamespaceLabel used along with RoleRequest is Sveltos solution for multi tenancy.
	ServiceAccountNamespaceLabel = "projectsveltos.io/serviceaccount-namespace"
)

// ReferencedResourceKind is a string representation of allowed kind of resources
// that can be referenced in a ClusterProfile
type ReferencedResourceKind string

// Define the ReferencedResourceKind constants.
const (
	SecretReferencedResourceKind    ReferencedResourceKind = "Secret"
	ConfigMapReferencedResourceKind ReferencedResourceKind = "ConfigMap"
)

const (
	// ClusterProfileSecretType is the only accepted type of secret in resources.
	ClusterProfileSecretType corev1.SecretType = "addons.projectsveltos.io/cluster-profile"
)

var (
	// ErrSecretTypeNotSupported signals that a Secret is not supported.
	ErrSecretTypeNotSupported = errors.New("unsupported secret type")
)

type Selector string

// +kubebuilder:validation:Enum:=Provisioning;Provisioned;Failed;Removing;Removed
type SveltosFeatureStatus string

const (
	// SveltosStatusProvisioning indicates that sveltos feature is being
	// provisioned in the workload cluster
	SveltosStatusProvisioning = SveltosFeatureStatus("Provisioning")

	// SveltosStatusProvisioned indicates that sveltos has been
	// provisioned in the workload cluster
	SveltosStatusProvisioned = SveltosFeatureStatus("Provisioned")

	// SveltosStatusFailed indicates that configuring sveltos feature
	// in the workload cluster failed
	SveltosStatusFailed = SveltosFeatureStatus("Failed")

	// SveltosStatusRemoving indicates that sveltos feature is being
	// removed
	SveltosStatusRemoving = SveltosFeatureStatus("Removing")

	// SveltosStatusRemoved indicates that sveltos feature is removed
	SveltosStatusRemoved = SveltosFeatureStatus("Removed")
)

type ClusterInfo struct {
	// Cluster references the Cluster
	Cluster corev1.ObjectReference `json:"cluster"`

	// Hash represents the hash of the Classifier currently deployed
	// in the Cluster
	Hash []byte `json:"hash"`

	// Status represents the state of the feature in the workload cluster
	// +optional
	Status SveltosFeatureStatus `json:"status,omitempty"`

	// FailureMessage provides more information about the error.
	// +optional
	FailureMessage *string `json:"failureMessage,omitempty"`
}
