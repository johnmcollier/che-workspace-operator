// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/workspace/v1alpha1.Workspace":               schema_pkg_apis_workspace_v1alpha1_Workspace(ref),
		"./pkg/apis/workspace/v1alpha1.WorkspaceExposure":       schema_pkg_apis_workspace_v1alpha1_WorkspaceExposure(ref),
		"./pkg/apis/workspace/v1alpha1.WorkspaceExposureSpec":   schema_pkg_apis_workspace_v1alpha1_WorkspaceExposureSpec(ref),
		"./pkg/apis/workspace/v1alpha1.WorkspaceExposureStatus": schema_pkg_apis_workspace_v1alpha1_WorkspaceExposureStatus(ref),
		"./pkg/apis/workspace/v1alpha1.WorkspaceSpec":           schema_pkg_apis_workspace_v1alpha1_WorkspaceSpec(ref),
	}
}

func schema_pkg_apis_workspace_v1alpha1_Workspace(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Workspace is the Schema for the workspaces API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Desired state of the workspace",
							Ref:         ref("./pkg/apis/workspace/v1alpha1.WorkspaceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Observed state of the workspace",
							Ref:         ref("./pkg/apis/workspace/v1alpha1.WorkspaceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/workspace/v1alpha1.WorkspaceSpec", "./pkg/apis/workspace/v1alpha1.WorkspaceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_workspace_v1alpha1_WorkspaceExposure(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WorkspaceExposure is the Schema for the workspaceexposures API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/workspace/v1alpha1.WorkspaceExposureSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/workspace/v1alpha1.WorkspaceExposureStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/workspace/v1alpha1.WorkspaceExposureSpec", "./pkg/apis/workspace/v1alpha1.WorkspaceExposureStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_workspace_v1alpha1_WorkspaceExposureSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WorkspaceExposureSpec defines the desired state of WorkspaceExposure",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"exposureClass": {
						SchemaProps: spec.SchemaProps{
							Description: "Class of the exposure: this drives which Workspace exposer controller will manage this exposure",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"exposed": {
						SchemaProps: spec.SchemaProps{
							Description: "Should the workspace be exposed ?",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"ingressGlobalDomain": {
						SchemaProps: spec.SchemaProps{
							Description: "ingress global domain (corresponds to the Openshift route suffix)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"workspacePodSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "Selector that shoud be used by created services to point to the workspace Pod",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"services": {
						SchemaProps: spec.SchemaProps{
							Description: "Services by machine name",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/workspace/v1alpha1.ServiceDescription"),
									},
								},
							},
						},
					},
				},
				Required: []string{"exposureClass", "exposed", "ingressGlobalDomain", "workspacePodSelector", "services"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/workspace/v1alpha1.ServiceDescription"},
	}
}

func schema_pkg_apis_workspace_v1alpha1_WorkspaceExposureStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WorkspaceExposureStatus defines the observed state of WorkspaceExposure",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Workspace Exposure status",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"exposedEndpoints": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Ref: ref("./pkg/apis/workspace/v1alpha1.ExposedEndpoint"),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/workspace/v1alpha1.ExposedEndpoint"},
	}
}

func schema_pkg_apis_workspace_v1alpha1_WorkspaceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WorkspaceSpec defines the desired state of Workspace",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"started": {
						SchemaProps: spec.SchemaProps{
							Description: "Whether the workspace should be started or stopped",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"exposureClass": {
						SchemaProps: spec.SchemaProps{
							Description: "Exposure class the defines how the workspace will be exposed toon the external network",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"devfile": {
						SchemaProps: spec.SchemaProps{
							Description: "Workspace Structure defined in the Devfile format syntax. For more details see the Che 7 documentation: https://www.eclipse.org/che/docs/che-7/making-a-workspace-portable-using-a-devfile/",
							Ref:         ref("./pkg/apis/workspace/v1alpha1.DevFileSpec"),
						},
					},
				},
				Required: []string{"started", "devfile"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/workspace/v1alpha1.DevFileSpec"},
	}
}
