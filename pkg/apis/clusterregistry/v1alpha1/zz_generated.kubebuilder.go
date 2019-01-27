/*
Copyright YEAR The Kubernetes Authors.

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
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: "clusterregistry.k8s.io", Version: "v1alpha1"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Cluster{},
		&ClusterList{},
		&ClusterCredentials{},
		&ClusterCredentialsList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterCredentialsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterCredentials `json:"items"`
}

// CRD Generation
func getFloat(f float64) *float64 {
	return &f
}

func getInt(i int64) *int64 {
	return &i
}

var (
	// Define CRDs for resources
	ClusterCRD = v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "clusters.clusterregistry.k8s.io",
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   "clusterregistry.k8s.io",
			Version: "v1alpha1",
			Names: v1beta1.CustomResourceDefinitionNames{
				Kind:   "Cluster",
				Plural: "clusters",
			},
			Scope: "Namespaced",
			Validation: &v1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &v1beta1.JSONSchemaProps{
					Properties: map[string]v1beta1.JSONSchemaProps{
						"apiVersion": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"kind": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"metadata": v1beta1.JSONSchemaProps{
							Type: "object",
						},
						"spec": v1beta1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1beta1.JSONSchemaProps{
								"authInfo": v1beta1.JSONSchemaProps{
									Type: "object",
									Properties: map[string]v1beta1.JSONSchemaProps{
										"controller": v1beta1.JSONSchemaProps{
											Type: "object",
											Properties: map[string]v1beta1.JSONSchemaProps{
												"kind": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"name": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"namespace": v1beta1.JSONSchemaProps{
													Type: "string",
												},
											},
										},
										"user": v1beta1.JSONSchemaProps{
											Type: "object",
											Properties: map[string]v1beta1.JSONSchemaProps{
												"kind": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"name": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"namespace": v1beta1.JSONSchemaProps{
													Type: "string",
												},
											},
										},
									},
								},
								"kubernetesApiEndpoints": v1beta1.JSONSchemaProps{
									Type: "object",
									Properties: map[string]v1beta1.JSONSchemaProps{
										"caBundle": v1beta1.JSONSchemaProps{
											Type:   "string",
											Format: "byte",
										},
										"serverEndpoints": v1beta1.JSONSchemaProps{
											Type: "array",
											Items: &v1beta1.JSONSchemaPropsOrArray{
												Schema: &v1beta1.JSONSchemaProps{
													Type: "object",
													Properties: map[string]v1beta1.JSONSchemaProps{
														"clientCIDR": v1beta1.JSONSchemaProps{
															Type: "string",
														},
														"serverAddress": v1beta1.JSONSchemaProps{
															Type: "string",
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"status": v1beta1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1beta1.JSONSchemaProps{
								"conditions": v1beta1.JSONSchemaProps{
									Type: "array",
									Items: &v1beta1.JSONSchemaPropsOrArray{
										Schema: &v1beta1.JSONSchemaProps{
											Type: "object",
											Properties: map[string]v1beta1.JSONSchemaProps{
												"lastHeartbeatTime": v1beta1.JSONSchemaProps{
													Type:   "string",
													Format: "date-time",
												},
												"lastTransitionTime": v1beta1.JSONSchemaProps{
													Type:   "string",
													Format: "date-time",
												},
												"message": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"reason": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"status": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"type": v1beta1.JSONSchemaProps{
													Type: "string",
												},
											},
											Required: []string{
												"type",
												"status",
											}},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	// Define CRDs for resources
	ClusterCredentialsCRD = v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "clustercredentials.clusterregistry.k8s.io",
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   "clusterregistry.k8s.io",
			Version: "v1alpha1",
			Names: v1beta1.CustomResourceDefinitionNames{
				Kind:   "ClusterCredentials",
				Plural: "clustercredentials",
			},
			Scope: "Namespaced",
			Validation: &v1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &v1beta1.JSONSchemaProps{
					Properties: map[string]v1beta1.JSONSchemaProps{
						"apiVersion": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"kind": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"metadata": v1beta1.JSONSchemaProps{
							Type: "object",
						},
						"spec": v1beta1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1beta1.JSONSchemaProps{
								"clusterRef": v1beta1.JSONSchemaProps{
									Type:       "object",
									Properties: map[string]v1beta1.JSONSchemaProps{},
								},
								"secretRef": v1beta1.JSONSchemaProps{
									Type:       "object",
									Properties: map[string]v1beta1.JSONSchemaProps{},
								},
							},
						},
						"status": v1beta1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1beta1.JSONSchemaProps{
								"availabilityZone": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"conditions": v1beta1.JSONSchemaProps{
									Type: "array",
									Items: &v1beta1.JSONSchemaPropsOrArray{
										Schema: &v1beta1.JSONSchemaProps{
											Type: "object",
											Properties: map[string]v1beta1.JSONSchemaProps{
												"lastHeartbeatTime": v1beta1.JSONSchemaProps{
													Type:   "string",
													Format: "date-time",
												},
												"lastProbeTime": v1beta1.JSONSchemaProps{
													Type:   "string",
													Format: "date-time",
												},
												"lastTransitionTime": v1beta1.JSONSchemaProps{
													Type:   "string",
													Format: "date-time",
												},
												"message": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"reason": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"status": v1beta1.JSONSchemaProps{
													Type: "string",
												},
												"type": v1beta1.JSONSchemaProps{
													Type: "string",
												},
											},
											Required: []string{
												"type",
												"status",
											}},
									},
								},
								"region": v1beta1.JSONSchemaProps{
									Type: "string",
								},
							},
						},
					},
				},
			},
			Subresources: &v1beta1.CustomResourceSubresources{
				Status: &v1beta1.CustomResourceSubresourceStatus{},
			},
		},
	}
)
