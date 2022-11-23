// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package deployment

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just CredentialConfigurations API methods
type credentialConfigurationsImpl struct {
	client *client.DatabricksClient
}

func (a *credentialConfigurationsImpl) CreateCredentialConfig(ctx context.Context, request CreateCredentialRequest) (*Credential, error) {
	var credential Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", a.client.Config.AccountID)
	err := a.client.Post(ctx, path, request, &credential)
	return &credential, err
}

func (a *credentialConfigurationsImpl) DeleteCredentialConfig(ctx context.Context, request DeleteCredentialConfigRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", a.client.Config.AccountID, request.CredentialsId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *credentialConfigurationsImpl) GetCredentialConfig(ctx context.Context, request GetCredentialConfigRequest) (*Credential, error) {
	var credential Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", a.client.Config.AccountID, request.CredentialsId)
	err := a.client.Get(ctx, path, request, &credential)
	return &credential, err
}

func (a *credentialConfigurationsImpl) ListCredentials(ctx context.Context) ([]Credential, error) {
	var credentialList []Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, nil, &credentialList)
	return credentialList, err
}

// unexported type that holds implementations of just KeyConfigurations API methods
type keyConfigurationsImpl struct {
	client *client.DatabricksClient
}

func (a *keyConfigurationsImpl) CreateKeyConfig(ctx context.Context, request CreateCustomerManagedKeyRequest) (*CustomerManagedKey, error) {
	var customerManagedKey CustomerManagedKey
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys", a.client.Config.AccountID)
	err := a.client.Post(ctx, path, request, &customerManagedKey)
	return &customerManagedKey, err
}

func (a *keyConfigurationsImpl) DeleteKeyConfig(ctx context.Context, request DeleteKeyConfigRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys/%v", a.client.Config.AccountID, request.CustomerManagedKeyId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *keyConfigurationsImpl) GetKeyConfig(ctx context.Context, request GetKeyConfigRequest) (*CustomerManagedKey, error) {
	var customerManagedKey CustomerManagedKey
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys/%v", a.client.Config.AccountID, request.CustomerManagedKeyId)
	err := a.client.Get(ctx, path, request, &customerManagedKey)
	return &customerManagedKey, err
}

func (a *keyConfigurationsImpl) GetKeyWorkspaceHistory(ctx context.Context) (*ListWorkspaceEncryptionKeyRecordsResponse, error) {
	var listWorkspaceEncryptionKeyRecordsResponse ListWorkspaceEncryptionKeyRecordsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-key-history", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, nil, &listWorkspaceEncryptionKeyRecordsResponse)
	return &listWorkspaceEncryptionKeyRecordsResponse, err
}

func (a *keyConfigurationsImpl) ListKeyConfigs(ctx context.Context) ([]CustomerManagedKey, error) {
	var customerManagedKeyList []CustomerManagedKey
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, nil, &customerManagedKeyList)
	return customerManagedKeyList, err
}

// unexported type that holds implementations of just NetworkConfigurations API methods
type networkConfigurationsImpl struct {
	client *client.DatabricksClient
}

func (a *networkConfigurationsImpl) CreateNetworkConfig(ctx context.Context, request CreateNetworkRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", a.client.Config.AccountID)
	err := a.client.Post(ctx, path, request, &network)
	return &network, err
}

func (a *networkConfigurationsImpl) DeleteNetworkConfig(ctx context.Context, request DeleteNetworkConfigRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", a.client.Config.AccountID, request.NetworkId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *networkConfigurationsImpl) GetNetworkConfig(ctx context.Context, request GetNetworkConfigRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", a.client.Config.AccountID, request.NetworkId)
	err := a.client.Get(ctx, path, request, &network)
	return &network, err
}

func (a *networkConfigurationsImpl) ListNetworkConfigs(ctx context.Context) ([]Network, error) {
	var networkList []Network
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, nil, &networkList)
	return networkList, err
}

// unexported type that holds implementations of just PrivateAccessSettings API methods
type privateAccessSettingsImpl struct {
	client *client.DatabricksClient
}

func (a *privateAccessSettingsImpl) CreatePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", a.client.Config.AccountID)
	err := a.client.Post(ctx, path, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

func (a *privateAccessSettingsImpl) DeletePrivateAccessSettings(ctx context.Context, request DeletePrivateAccessSettingsRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.Config.AccountID, request.PrivateAccessSettingsId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *privateAccessSettingsImpl) GetPrivateAccessSettings(ctx context.Context, request GetPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.Config.AccountID, request.PrivateAccessSettingsId)
	err := a.client.Get(ctx, path, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

func (a *privateAccessSettingsImpl) ListPrivateAccessSettings(ctx context.Context) ([]PrivateAccessSettings, error) {
	var privateAccessSettingsList []PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, nil, &privateAccessSettingsList)
	return privateAccessSettingsList, err
}

func (a *privateAccessSettingsImpl) ReplacePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.Config.AccountID, request.PrivateAccessSettingsId)
	err := a.client.Put(ctx, path, request)
	return err
}

// unexported type that holds implementations of just StorageConfigurations API methods
type storageConfigurationsImpl struct {
	client *client.DatabricksClient
}

func (a *storageConfigurationsImpl) CreateStorageConfig(ctx context.Context, request CreateStorageConfigurationRequest) (*StorageConfiguration, error) {
	var storageConfiguration StorageConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations", a.client.Config.AccountID)
	err := a.client.Post(ctx, path, request, &storageConfiguration)
	return &storageConfiguration, err
}

func (a *storageConfigurationsImpl) DeleteStorageConfig(ctx context.Context, request DeleteStorageConfigRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations/%v", a.client.Config.AccountID, request.StorageConfigurationId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *storageConfigurationsImpl) GetStorageConfig(ctx context.Context, request GetStorageConfigRequest) (*StorageConfiguration, error) {
	var storageConfiguration StorageConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations/%v", a.client.Config.AccountID, request.StorageConfigurationId)
	err := a.client.Get(ctx, path, request, &storageConfiguration)
	return &storageConfiguration, err
}

func (a *storageConfigurationsImpl) ListStorageConfigs(ctx context.Context) ([]StorageConfiguration, error) {
	var storageConfigurationList []StorageConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, nil, &storageConfigurationList)
	return storageConfigurationList, err
}

// unexported type that holds implementations of just VpcEndpoints API methods
type vpcEndpointsImpl struct {
	client *client.DatabricksClient
}

func (a *vpcEndpointsImpl) CreateVpcEndpoint(ctx context.Context, request CreateVpcEndpointRequest) (*VpcEndpoint, error) {
	var vpcEndpoint VpcEndpoint
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints", a.client.Config.AccountID)
	err := a.client.Post(ctx, path, request, &vpcEndpoint)
	return &vpcEndpoint, err
}

func (a *vpcEndpointsImpl) DeleteVpcEndpoint(ctx context.Context, request DeleteVpcEndpointRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints/%v", a.client.Config.AccountID, request.VpcEndpointId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *vpcEndpointsImpl) GetVpcEndpoint(ctx context.Context, request GetVpcEndpointRequest) (*VpcEndpoint, error) {
	var vpcEndpoint VpcEndpoint
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints/%v", a.client.Config.AccountID, request.VpcEndpointId)
	err := a.client.Get(ctx, path, request, &vpcEndpoint)
	return &vpcEndpoint, err
}

func (a *vpcEndpointsImpl) ListVpcEndpoints(ctx context.Context) ([]VpcEndpoint, error) {
	var vpcEndpointList []VpcEndpoint
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, nil, &vpcEndpointList)
	return vpcEndpointList, err
}

// unexported type that holds implementations of just Workspaces API methods
type workspacesImpl struct {
	client *client.DatabricksClient
}

func (a *workspacesImpl) CreateWorkspace(ctx context.Context, request CreateWorkspaceRequest) (*Workspace, error) {
	var workspace Workspace
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces", a.client.Config.AccountID)
	err := a.client.Post(ctx, path, request, &workspace)
	return &workspace, err
}

func (a *workspacesImpl) DeleteWorkspace(ctx context.Context, request DeleteWorkspaceRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.Config.AccountID, request.WorkspaceId)
	err := a.client.Delete(ctx, path, request)
	return err
}

func (a *workspacesImpl) GetAllWorkspaces(ctx context.Context) ([]Workspace, error) {
	var workspaceList []Workspace
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces", a.client.Config.AccountID)
	err := a.client.Get(ctx, path, nil, &workspaceList)
	return workspaceList, err
}

func (a *workspacesImpl) GetWorkspace(ctx context.Context, request GetWorkspaceRequest) (*Workspace, error) {
	var workspace Workspace
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.Config.AccountID, request.WorkspaceId)
	err := a.client.Get(ctx, path, request, &workspace)
	return &workspace, err
}

func (a *workspacesImpl) GetWorkspaceKeyHistory(ctx context.Context, request GetWorkspaceKeyHistoryRequest) (*ListWorkspaceEncryptionKeyRecordsResponse, error) {
	var listWorkspaceEncryptionKeyRecordsResponse ListWorkspaceEncryptionKeyRecordsResponse
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v/customer-managed-key-history", a.client.Config.AccountID, request.WorkspaceId)
	err := a.client.Get(ctx, path, request, &listWorkspaceEncryptionKeyRecordsResponse)
	return &listWorkspaceEncryptionKeyRecordsResponse, err
}

func (a *workspacesImpl) UpdateWorkspace(ctx context.Context, request UpdateWorkspaceRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.Config.AccountID, request.WorkspaceId)
	err := a.client.Patch(ctx, path, request)
	return err
}