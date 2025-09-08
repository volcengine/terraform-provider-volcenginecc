package generic

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	fwvalidators "github.com/volcengine/terraform-provider-volcenginecc/internal/validators"
)

func TestTranslateForUpdate(t *testing.T) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AttachmentCount
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "策略绑定的身份数量。",
		//	  "type": "integer"
		//	}
		"attachment_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "策略绑定的身份数量。",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Category
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "系统预设策略所属的分类，通常为服务代码，对于自定义策略该字段不会返回值。",
		//	  "type": "string"
		//	}
		//"category": schema.StringAttribute{ /*START ATTRIBUTE*/
		//	Description: "系统预设策略所属的分类，通常为服务代码，对于自定义策略该字段不会返回值。",
		//	Optional:    true,
		//	Computed:    true,
		//	PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
		//		stringplanmodifier.UseStateForUnknown(),
		//	}, /*END PLAN MODIFIERS*/
		//}, /*END ATTRIBUTE*/
		// Property: CreatedTime
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "策略创建时间。",
		//	  "type": "string"
		//	}
		"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "策略创建时间。",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "策略描述，长度不超过128。",
		//	  "maxLength": 128,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "策略描述，长度不超过128。",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(128),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IsServiceRolePolicy
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "是否是服务关联角色的策略，0代表否，1代表是。",
		//	  "type": "integer"
		//	}
		"is_service_role_policy": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "是否是服务关联角色的策略，0代表否，1代表是。",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyDocument
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "策略语法内容，例如：{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"iam:\",\"tag:\"],\"Resource\":[\"*\"]}]}",
		//	  "maxLength": 6144,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "策略语法内容，例如：{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"iam:\",\"tag:\"],\"Resource\":[\"*\"]}]}",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 6144),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyName
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "策略名，长度1~64，支持英文、数字和+=,.@-_符号。",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "策略名，长度1~64，支持英文、数字和+=,.@-_符号。",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyRoles
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "策略绑定的角色列表。",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "properties": {
		//	      "CreatedTime": {
		//	        "description": "策略绑定时间。",
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "策略描述。",
		//	        "type": "string"
		//	      },
		//	      "DisplayName": {
		//	        "description": "显示名称。",
		//	        "type": "string"
		//	      },
		//	      "Id": {
		//	        "description": "唯一标识。",
		//	        "type": "integer"
		//	      },
		//	      "Name": {
		//	        "description": "对应用户、角色、用户组的名称。",
		//	        "type": "string"
		//	      },
		//	      "PolicyScope": {
		//	        "description": "策略绑定的项目列表。",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "properties": {
		//	            "CreatedTime": {
		//	              "description": "项目授权时间。",
		//	              "type": "string"
		//	            },
		//	            "PolicyScopeType": {
		//	              "description": "授权类型。Global代表全局授权，Project代表按项目授权。",
		//	              "enum": [
		//	                "Global",
		//	                "Project"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "ProjectDisplayName": {
		//	              "description": "项目显示名。",
		//	              "type": "string"
		//	            },
		//	            "ProjectName": {
		//	              "description": "项目名。",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "ProjectName"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      }
		//	    },
		//	    "required": [
		//	      "Name"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"policy_roles": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CreatedTime
					"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "策略绑定时间。",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "策略描述。",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: DisplayName
					"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "显示名称。",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Id
					"id": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "唯一标识。",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "对应用户、角色、用户组的名称。",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: PolicyScope
					"policy_scope": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CreatedTime
								//"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
								//	Description: "项目授权时间。",
								//	Optional:    true,
								//	Computed:    true,
								//	PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								//		stringplanmodifier.UseStateForUnknown(),
								//	}, /*END PLAN MODIFIERS*/
								//}, /*END ATTRIBUTE*/
								// Property: PolicyScopeType
								"policy_scope_type": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "授权类型。Global代表全局授权，Project代表按项目授权。",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.OneOf(
											"Global",
											"Project",
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: ProjectDisplayName
								"project_display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "项目显示名。",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: ProjectName
								"project_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "项目名。",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										fwvalidators.NotNullString(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "策略绑定的项目列表。",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
							setplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "策略绑定的角色列表。",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyTrn
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "策略的TRN。",
		//	  "type": "string"
		//	}
		"policy_trn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "策略的TRN。",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyType
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "策略类型。System代表系统预设策略，Custom代表自定义策略。",
		//	  "enum": [
		//	    "System",
		//	    "Custom"
		//	  ],
		//	  "type": "string"
		//	}
		"policy_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "策略类型。System代表系统预设策略，Custom代表自定义策略。",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"System",
					"Custom",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyUserGroups
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "策略绑定的用户组列表。",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "properties": {
		//	      "CreatedTime": {
		//	        "description": "策略绑定时间。",
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "策略描述。",
		//	        "type": "string"
		//	      },
		//	      "DisplayName": {
		//	        "description": "显示名称。",
		//	        "type": "string"
		//	      },
		//	      "Id": {
		//	        "description": "唯一标识。",
		//	        "type": "integer"
		//	      },
		//	      "Name": {
		//	        "description": "对应用户、角色、用户组的名称。",
		//	        "type": "string"
		//	      },
		//	      "PolicyScope": {
		//	        "description": "策略绑定的项目列表。",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "properties": {
		//	            "CreatedTime": {
		//	              "description": "项目授权时间。",
		//	              "type": "string"
		//	            },
		//	            "PolicyScopeType": {
		//	              "description": "授权类型。Global代表全局授权，Project代表按项目授权。",
		//	              "enum": [
		//	                "Global",
		//	                "Project"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "ProjectDisplayName": {
		//	              "description": "项目显示名。",
		//	              "type": "string"
		//	            },
		//	            "ProjectName": {
		//	              "description": "项目名。",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "ProjectName"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      }
		//	    },
		//	    "required": [
		//	      "Name"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		//"policy_user_groups": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
		//	NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
		//		Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
		//			// Property: CreatedTime
		//			"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
		//				Description: "策略绑定时间。",
		//				Optional:    true,
		//				Computed:    true,
		//				PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
		//					stringplanmodifier.UseStateForUnknown(),
		//				}, /*END PLAN MODIFIERS*/
		//			}, /*END ATTRIBUTE*/
		//			// Property: Description
		//			"description": schema.StringAttribute{ /*START ATTRIBUTE*/
		//				Description: "策略描述。",
		//				Optional:    true,
		//				Computed:    true,
		//				PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
		//					stringplanmodifier.UseStateForUnknown(),
		//				}, /*END PLAN MODIFIERS*/
		//			}, /*END ATTRIBUTE*/
		//			// Property: DisplayName
		//			"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
		//				Description: "显示名称。",
		//				Optional:    true,
		//				Computed:    true,
		//				PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
		//					stringplanmodifier.UseStateForUnknown(),
		//				}, /*END PLAN MODIFIERS*/
		//			}, /*END ATTRIBUTE*/
		//			// Property: Id
		//			"id": schema.Int64Attribute{ /*START ATTRIBUTE*/
		//				Description: "唯一标识。",
		//				Optional:    true,
		//				Computed:    true,
		//				PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
		//					int64planmodifier.UseStateForUnknown(),
		//				}, /*END PLAN MODIFIERS*/
		//			}, /*END ATTRIBUTE*/
		//			// Property: Name
		//			"name": schema.StringAttribute{ /*START ATTRIBUTE*/
		//				Description: "对应用户、角色、用户组的名称。",
		//				Optional:    true,
		//				Computed:    true,
		//				Validators: []validator.String{ /*START VALIDATORS*/
		//					fwvalidators.NotNullString(),
		//				}, /*END VALIDATORS*/
		//				PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
		//					stringplanmodifier.UseStateForUnknown(),
		//				}, /*END PLAN MODIFIERS*/
		//			}, /*END ATTRIBUTE*/
		//			// Property: PolicyScope
		//			"policy_scope": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
		//				NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
		//					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
		//						// Property: CreatedTime
		//						"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
		//							Description: "项目授权时间。",
		//							Optional:    true,
		//							Computed:    true,
		//							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
		//								stringplanmodifier.UseStateForUnknown(),
		//							}, /*END PLAN MODIFIERS*/
		//						}, /*END ATTRIBUTE*/
		//						// Property: PolicyScopeType
		//						"policy_scope_type": schema.StringAttribute{ /*START ATTRIBUTE*/
		//							Description: "授权类型。Global代表全局授权，Project代表按项目授权。",
		//							Optional:    true,
		//							Computed:    true,
		//							Validators: []validator.String{ /*START VALIDATORS*/
		//								stringvalidator.OneOf(
		//									"Global",
		//									"Project",
		//								),
		//							}, /*END VALIDATORS*/
		//							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
		//								stringplanmodifier.UseStateForUnknown(),
		//							}, /*END PLAN MODIFIERS*/
		//						}, /*END ATTRIBUTE*/
		//						// Property: ProjectDisplayName
		//						"project_display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
		//							Description: "项目显示名。",
		//							Optional:    true,
		//							Computed:    true,
		//							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
		//								stringplanmodifier.UseStateForUnknown(),
		//							}, /*END PLAN MODIFIERS*/
		//						}, /*END ATTRIBUTE*/
		//						// Property: ProjectName
		//						"project_name": schema.StringAttribute{ /*START ATTRIBUTE*/
		//							Description: "项目名。",
		//							Optional:    true,
		//							Computed:    true,
		//							Validators: []validator.String{ /*START VALIDATORS*/
		//								fwvalidators.NotNullString(),
		//							}, /*END VALIDATORS*/
		//							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
		//								stringplanmodifier.UseStateForUnknown(),
		//							}, /*END PLAN MODIFIERS*/
		//						}, /*END ATTRIBUTE*/
		//					}, /*END SCHEMA*/
		//				}, /*END NESTED OBJECT*/
		//				Description: "策略绑定的项目列表。",
		//				Optional:    true,
		//				Computed:    true,
		//				PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
		//					setplanmodifier.UseStateForUnknown(),
		//				}, /*END PLAN MODIFIERS*/
		//			}, /*END ATTRIBUTE*/
		//		}, /*END SCHEMA*/
		//	}, /*END NESTED OBJECT*/
		//	Description: "策略绑定的用户组列表。",
		//	Optional:    true,
		//	Computed:    true,
		//	PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
		//		setplanmodifier.UseStateForUnknown(),
		//	}, /*END PLAN MODIFIERS*/
		//}, /*END ATTRIBUTE*/
		// Property: PolicyUsers
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "策略绑定的用户列表。",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "properties": {
		//	      "CreatedTime": {
		//	        "description": "策略绑定时间。",
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "description": "策略描述。",
		//	        "type": "string"
		//	      },
		//	      "DisplayName": {
		//	        "description": "显示名称。",
		//	        "type": "string"
		//	      },
		//	      "Id": {
		//	        "description": "唯一标识。",
		//	        "type": "integer"
		//	      },
		//	      "Name": {
		//	        "description": "对应用户、角色、用户组的名称。",
		//	        "type": "string"
		//	      },
		//	      "PolicyScope": {
		//	        "description": "策略绑定的项目列表。",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "properties": {
		//	            "CreatedTime": {
		//	              "description": "项目授权时间。",
		//	              "type": "string"
		//	            },
		//	            "PolicyScopeType": {
		//	              "description": "授权类型。Global代表全局授权，Project代表按项目授权。",
		//	              "enum": [
		//	                "Global",
		//	                "Project"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "ProjectDisplayName": {
		//	              "description": "项目显示名。",
		//	              "type": "string"
		//	            },
		//	            "ProjectName": {
		//	              "description": "项目名。",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "ProjectName"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      }
		//	    },
		//	    "required": [
		//	      "Name"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"policy_users": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CreatedTime
					//"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					//	Description: "策略绑定时间。",
					//	Optional:    true,
					//	Computed:    true,
					//	PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
					//		stringplanmodifier.UseStateForUnknown(),
					//	}, /*END PLAN MODIFIERS*/
					//}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "策略描述。",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: DisplayName
					"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "显示名称。",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Id
					"id": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "唯一标识。",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "对应用户、角色、用户组的名称。",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: PolicyScope
					"policy_scope": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CreatedTime
								"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "项目授权时间。",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: PolicyScopeType
								"policy_scope_type": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "授权类型。Global代表全局授权，Project代表按项目授权。",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.OneOf(
											"Global",
											"Project",
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: ProjectDisplayName
								"project_display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "项目显示名。",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: ProjectName
								"project_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "项目名。",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										fwvalidators.NotNullString(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "策略绑定的项目列表。",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
							setplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "策略绑定的用户列表。",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedTime
		// Cloud Control resource type schema:
		//
		//	{
		//	  "description": "策略更新时间。",
		//	  "type": "string"
		//	}
		"updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "策略更新时间。",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to Cloud Control primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "策略是对权限的一种描述，IAM用户、用户组或角色均需通过关联策略来赋予权限。当系统预设策略不能满足要求时，您可以创建自定义策略，对权限进行细粒度的定义。",
		Version:     1,
		Attributes:  attributes,
	}

	json := `
{
    "AttachmentCount": 7,
    "Category": "Custom",
    "CreatedTime": "20250815T030714Z",
    "Description": "iam_policy description",
    "IsServiceRolePolicy": 0,
    "PolicyDocument": "{\"Statement\":[{\"Action\":[\"clb:DescribeAclAttributes\",\"clb:DescribeHealthCheckLogProjectAttributes\",\"clb:DescribeListenerAttributes\",\"clb:DescribeListenerHealth\",\"clb:DescribeLoadBalancerAttributes\",\"clb:DescribeLoadBalancerHealth\",\"clb:DescribeLoadBalancersBilling\",\"clb:DescribeNLBListenerAttributes\",\"clb:DescribeNLBListenerCertificates\",\"clb:DescribeNLBListenerHealth\",\"clb:DescribeNLBListeners\",\"clb:DescribeNLBServerGroupAttributes\",\"clb:DescribeNLBServerGroups\",\"clb:DescribeNLBZones\",\"clb:DescribeNetworkLoadBalancerAttributes\",\"clb:DescribeNetworkLoadBalancers\",\"clb:DescribeServerGroupAttributes\",\"clb:DescribeZones\",\"clb:ListTagsForNLBResources\",\"clb:DescribeAcls\",\"clb:DescribeCertificates\",\"clb:DescribeHealthCheckLogTopicAttributes\",\"clb:DescribeListeners\",\"clb:DescribeLoadBalancerSpecs\",\"clb:DescribeLoadBalancers\",\"clb:DescribeRules\",\"clb:DescribeServerGroups\",\"clb:ListTagsForResources\",\"clb:TagNLBResources\",\"clb:TagResources\",\"clb:UntagNLBResources\",\"clb:UntagResources\"],\"Effect\":\"Allow\",\"Resource\":[\"*\"]}]}",
    "PolicyName": "ccapi-policy-1001",
    "PolicyRoles":
    [
        {
            "CreatedTime": "20250815T030720Z",
            "Description": "该角色允许容器服务访问您账号下的其他服务资源",
            "DisplayName": "容器服务服务角色",
            "Id": 42399074,
            "Name": "VKEAutoscalerRole",
            "PolicyScope":
            [
                {
                    "CreatedTime": "20250815T030720Z",
                    "PolicyScopeType": "Project",
                    "ProjectDisplayName": "默认项目",
                    "ProjectName": "default"
                }
            ]
        },
        {
            "CreatedTime": "20250815T030719Z",
            "Description": "该角色允许容器服务访问您账号下的其他服务资源",
            "DisplayName": "容器服务服务角色",
            "Id": 42399071,
            "Name": "VKEImageRole",
            "PolicyScope":
            [
                {
                    "CreatedTime": "20250815T030719Z",
                    "PolicyScopeType": "Project",
                    "ProjectDisplayName": "pln-demo",
                    "ProjectName": "demo"
                },
                {
                    "CreatedTime": "20250815T030719Z",
                    "PolicyScopeType": "Project",
                    "ProjectDisplayName": "默认项目",
                    "ProjectName": "default"
                }
            ]
        },
        {
            "CreatedTime": "20250815T030719Z",
            "Description": "测试terraformtofu role",
            "DisplayName": "测试角色111",
            "Id": 40527308,
            "Name": "terraformtofu_iam_role-1",
            "PolicyScope":
            [
                {
                    "CreatedTime": "20250815T030719Z",
                    "PolicyScopeType": "Project",
                    "ProjectDisplayName": "默认项目",
                    "ProjectName": "default"
                }
            ]
        }
    ],
    "PolicyTrn": "trn:iam::2107436677:policy/ccapi-policy-1001",
    "PolicyType": "Custom",
    "PolicyUserGroups":
    [
        {
            "CreatedTime": "20250815T030720Z",
            "Description": "",
            "DisplayName": "",
            "Id": 10550146,
            "Name": "ccapi-test",
            "PolicyScope":
            [
                {
                    "CreatedTime": "20250815T030720Z",
                    "PolicyScopeType": "Project",
                    "ProjectDisplayName": "默认项目",
                    "ProjectName": "default"
                }
            ]
        },
        {
            "CreatedTime": "20250815T030720Z",
            "Description": "",
            "DisplayName": "CCAPI研发小组-勿动",
            "Id": 10129016,
            "Name": "ccapi-dev",
            "PolicyScope":
            [
                {
                    "CreatedTime": "20250815T030720Z",
                    "PolicyScopeType": "Project",
                    "ProjectDisplayName": "默认项目",
                    "ProjectName": "default"
                },
                {
                    "CreatedTime": "20250815T030720Z",
                    "PolicyScopeType": "Project",
                    "ProjectDisplayName": "",
                    "ProjectName": "ceshi"
                }
            ]
        }
    ],
    "PolicyUsers":
    [
        {
            "CreatedTime": "20250815T030719Z",
            "Description": "",
            "DisplayName": "",
            "Id": 53288663,
            "Name": "wangwuqiang.iidu",
            "PolicyScope":
            [
                {
                    "CreatedTime": "20250815T030719Z",
                    "PolicyScopeType": "Project",
                    "ProjectDisplayName": "默认项目",
                    "ProjectName": "default"
                },
                {
                    "CreatedTime": "20250815T030719Z",
                    "PolicyScopeType": "Project",
                    "ProjectDisplayName": "",
                    "ProjectName": "ceshi"
                }
            ]
        },
        {
            "CreatedTime": "20250815T030719Z",
            "Description": "",
            "DisplayName": "",
            "Id": 51513418,
            "Name": "wangshuang.ws1",
            "PolicyScope":
            [
                {
                    "CreatedTime": "20250815T030719Z",
                    "PolicyScopeType": "Project",
                    "ProjectDisplayName": "默认项目",
                    "ProjectName": "default"
                },
                {
                    "CreatedTime": "20250815T030719Z",
                    "PolicyScopeType": "Project",
                    "ProjectDisplayName": "",
                    "ProjectName": "ceshi"
                }
            ]
        }
    ],
    "UpdatedTime": "20250815T030714Z"
}
`
	tfToCC := map[string]string{
		"attachment_count":       "AttachmentCount",
		"category":               "Category",
		"created_time":           "CreatedTime",
		"description":            "Description",
		"display_name":           "DisplayName",
		"id":                     "Id",
		"is_service_role_policy": "IsServiceRolePolicy",
		"name":                   "Name",
		"policy_document":        "PolicyDocument",
		"policy_name":            "PolicyName",
		"policy_roles":           "PolicyRoles",
		"policy_scope":           "PolicyScope",
		"policy_scope_type":      "PolicyScopeType",
		"policy_trn":             "PolicyTrn",
		"policy_type":            "PolicyType",
		"policy_user_groups":     "PolicyUserGroups",
		"policy_users":           "PolicyUsers",
		"project_display_name":   "ProjectDisplayName",
		"project_name":           "ProjectName",
		"updated_time":           "UpdatedTime",
	}
	ccToTf := make(map[string]string)
	for key, value := range tfToCC {
		ccToTf[value] = key
	}
	result, _ := translateForUpdate(json, schema.Attributes, ccToTf)
	fmt.Printf(result)
}
