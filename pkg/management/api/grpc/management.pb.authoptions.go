// Code generated by protoc-gen-authmethod. DO NOT EDIT.

package grpc

import (
	"google.golang.org/grpc"

	utils_auth "github.com/caos/zitadel/internal/auth"
	utils_grpc "github.com/caos/zitadel/internal/grpc"
)

/**
 * ManagementService
 */

var ManagementService_AuthMethods = utils_auth.AuthMethodMapping{

	"/caos.zitadel.management.api.v1.ManagementService/GetUserByID": utils_auth.AuthOption{
		Permission: "user.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetUserByEmailGlobal": utils_auth.AuthOption{
		Permission: "user.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SearchUsers": utils_auth.AuthOption{
		Permission: "user.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/IsUserUnique": utils_auth.AuthOption{
		Permission: "user.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/CreateUser": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeactivateUser": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ReactivateUser": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/LockUser": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UnlockUser": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeleteUser": utils_auth.AuthOption{
		Permission: "user.delete",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UserChanges": utils_auth.AuthOption{
		Permission: "user.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ApplicationChanges": utils_auth.AuthOption{
		Permission: "project.app.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/OrgChanges": utils_auth.AuthOption{
		Permission: "org.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ProjectChanges": utils_auth.AuthOption{
		Permission: "project.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetUserProfile": utils_auth.AuthOption{
		Permission: "user.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdateUserProfile": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetUserEmail": utils_auth.AuthOption{
		Permission: "user.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ChangeUserEmail": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ResendEmailVerificationMail": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetUserPhone": utils_auth.AuthOption{
		Permission: "user.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ChangeUserPhone": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ResendPhoneVerificationCode": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetUserAddress": utils_auth.AuthOption{
		Permission: "user.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdateUserAddress": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetUserMfas": utils_auth.AuthOption{
		Permission: "user.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SendSetPasswordNotification": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SetInitialPassword": utils_auth.AuthOption{
		Permission: "user.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetPasswordComplexityPolicy": utils_auth.AuthOption{
		Permission: "policy.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/CreatePasswordComplexityPolicy": utils_auth.AuthOption{
		Permission: "policy.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdatePasswordComplexityPolicy": utils_auth.AuthOption{
		Permission: "policy.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeletePasswordComplexityPolicy": utils_auth.AuthOption{
		Permission: "policy.delete",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetPasswordAgePolicy": utils_auth.AuthOption{
		Permission: "policy.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/CreatePasswordAgePolicy": utils_auth.AuthOption{
		Permission: "policy.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdatePasswordAgePolicy": utils_auth.AuthOption{
		Permission: "policy.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeletePasswordAgePolicy": utils_auth.AuthOption{
		Permission: "policy.delete",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetPasswordLockoutPolicy": utils_auth.AuthOption{
		Permission: "policy.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/CreatePasswordLockoutPolicy": utils_auth.AuthOption{
		Permission: "policy.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdatePasswordLockoutPolicy": utils_auth.AuthOption{
		Permission: "policy.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeletePasswordLockoutPolicy": utils_auth.AuthOption{
		Permission: "policy.delete",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetOrgByID": utils_auth.AuthOption{
		Permission: "org.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetOrgByDomainGlobal": utils_auth.AuthOption{
		Permission: "org.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeactivateOrg": utils_auth.AuthOption{
		Permission: "org.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ReactivateOrg": utils_auth.AuthOption{
		Permission: "org.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetOrgMemberRoles": utils_auth.AuthOption{
		Permission: "org.member.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/AddOrgMember": utils_auth.AuthOption{
		Permission: "org.member.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ChangeOrgMember": utils_auth.AuthOption{
		Permission: "org.member.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/RemoveOrgMember": utils_auth.AuthOption{
		Permission: "org.member.delete",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SearchOrgMembers": utils_auth.AuthOption{
		Permission: "org.member.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SearchProjects": utils_auth.AuthOption{
		Permission: "project.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ProjectByID": utils_auth.AuthOption{
		Permission: "project.read",
		CheckParam: "Id",
	},

	"/caos.zitadel.management.api.v1.ManagementService/CreateProject": utils_auth.AuthOption{
		Permission: "project.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdateProject": utils_auth.AuthOption{
		Permission: "project.write",
		CheckParam: "Id",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeactivateProject": utils_auth.AuthOption{
		Permission: "project.write",
		CheckParam: "Id",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ReactivateProject": utils_auth.AuthOption{
		Permission: "project.write",
		CheckParam: "Id",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetGrantedProjectGrantByID": utils_auth.AuthOption{
		Permission: "project.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetProjectMemberRoles": utils_auth.AuthOption{
		Permission: "project.member.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SearchProjectMembers": utils_auth.AuthOption{
		Permission: "project.member.read",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/AddProjectMember": utils_auth.AuthOption{
		Permission: "project.member.write",
		CheckParam: "Id",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ChangeProjectMember": utils_auth.AuthOption{
		Permission: "project.member.write",
		CheckParam: "Id",
	},

	"/caos.zitadel.management.api.v1.ManagementService/RemoveProjectMember": utils_auth.AuthOption{
		Permission: "project.member.delete",
		CheckParam: "Id",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SearchProjectRoles": utils_auth.AuthOption{
		Permission: "project.role.read",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/AddProjectRole": utils_auth.AuthOption{
		Permission: "project.role.write",
		CheckParam: "Id",
	},

	"/caos.zitadel.management.api.v1.ManagementService/RemoveProjectRole": utils_auth.AuthOption{
		Permission: "project.role.delete",
		CheckParam: "Id",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SearchApplications": utils_auth.AuthOption{
		Permission: "project.app.read",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ApplicationByID": utils_auth.AuthOption{
		Permission: "project.app.read",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/CreateOIDCApplication": utils_auth.AuthOption{
		Permission: "project.app.write",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdateApplication": utils_auth.AuthOption{
		Permission: "project.app.write",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeactivateApplication": utils_auth.AuthOption{
		Permission: "project.app.write",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ReactivateApplication": utils_auth.AuthOption{
		Permission: "project.app.write",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdateApplicationOIDCConfig": utils_auth.AuthOption{
		Permission: "project.app.write",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/RegenerateOIDCClientSecret": utils_auth.AuthOption{
		Permission: "project.app.write",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SearchProjectGrants": utils_auth.AuthOption{
		Permission: "project.grant.read",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ProjectGrantByID": utils_auth.AuthOption{
		Permission: "project.grant.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/CreateProjectGrant": utils_auth.AuthOption{
		Permission: "project.grant.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdateProjectGrant": utils_auth.AuthOption{
		Permission: "project.grant.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeactivateProjectGrant": utils_auth.AuthOption{
		Permission: "project.grant.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ReactivateProjectGrant": utils_auth.AuthOption{
		Permission: "project.grant.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/GetProjectGrantMemberRoles": utils_auth.AuthOption{
		Permission: "project.grant.member.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SearchProjectGrantMembers": utils_auth.AuthOption{
		Permission: "project.grant.member.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/AddProjectGrantMember": utils_auth.AuthOption{
		Permission: "project.grant.member.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ChangeProjectGrantMember": utils_auth.AuthOption{
		Permission: "project.grant.member.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/RemoveProjectGrantMember": utils_auth.AuthOption{
		Permission: "project.grant.member.delete",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SearchUserGrants": utils_auth.AuthOption{
		Permission: "user.grant.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UserGrantByID": utils_auth.AuthOption{
		Permission: "user.grant.read",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/CreateUserGrant": utils_auth.AuthOption{
		Permission: "user.grant.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdateUserGrant": utils_auth.AuthOption{
		Permission: "user.grant.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeactivateUserGrant": utils_auth.AuthOption{
		Permission: "user.grant.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ReactivateUserGrant": utils_auth.AuthOption{
		Permission: "user.grant.write",
		CheckParam: "",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SearchProjectUserGrants": utils_auth.AuthOption{
		Permission: "project.user.grant.read",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ProjectUserGrantByID": utils_auth.AuthOption{
		Permission: "project.user.grant.read",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/CreateProjectUserGrant": utils_auth.AuthOption{
		Permission: "project.user.grant.write",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdateProjectUserGrant": utils_auth.AuthOption{
		Permission: "project.user.grant.write",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeactivateProjectUserGrant": utils_auth.AuthOption{
		Permission: "project.user.grant.write",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ReactivateProjectUserGrant": utils_auth.AuthOption{
		Permission: "project.user.grant.write",
		CheckParam: "ProjectId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/SearchProjectGrantUserGrants": utils_auth.AuthOption{
		Permission: "project.grant.user.grant.read",
		CheckParam: "ProjectGrantId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ProjectGrantUserGrantByID": utils_auth.AuthOption{
		Permission: "project.grant.user.grant.read",
		CheckParam: "ProjectGrantId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/CreateProjectGrantUserGrant": utils_auth.AuthOption{
		Permission: "project.grant.user.grant.write",
		CheckParam: "ProjectGrantId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/UpdateProjectGrantUserGrant": utils_auth.AuthOption{
		Permission: "project.grant.user.grant.write",
		CheckParam: "ProjectGrantId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/DeactivateProjectGrantUserGrant": utils_auth.AuthOption{
		Permission: "project.grant.user.grant.write",
		CheckParam: "ProjectGrantId",
	},

	"/caos.zitadel.management.api.v1.ManagementService/ReactivateProjectGrantUserGrant": utils_auth.AuthOption{
		Permission: "project.grant.user.grant.write",
		CheckParam: "ProjectGrantId",
	},
}

func ManagementService_Authorization_Interceptor(verifier utils_auth.TokenVerifier, authConf *utils_auth.AuthConfig) grpc.UnaryServerInterceptor {
	return utils_grpc.AuthorizationInterceptor(verifier, authConf, ManagementService_AuthMethods)
}
