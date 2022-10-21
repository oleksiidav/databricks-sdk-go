// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package libraries

import (
	"context"
)

// Databricks Managed Libraries REST API
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type LibrariesService interface {

	// Get all statuses
	//
	// Get the status of all libraries on all clusters. A status will be
	// available for all libraries installed on this cluster via the API or the
	// libraries UI as well as libraries set to be installed on all clusters via
	// the libraries UI.
	AllClusterStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error)

	// Get status
	//
	// Get the status of libraries on a cluster. A status will be available for
	// all libraries installed on this cluster via the API or the libraries UI
	// as well as libraries set to be installed on all clusters via the
	// libraries UI. The order of returned libraries will be as follows.
	//
	// 1. Libraries set to be installed on this cluster will be returned first.
	// Within this group, the final order will be order in which the libraries
	// were added to the cluster.
	//
	// 2. Libraries set to be installed on all clusters are returned next.
	// Within this group there is no order guarantee.
	//
	// 3. Libraries that were previously requested on this cluster or on all
	// clusters, but now marked for removal. Within this group there is no order
	// guarantee.
	ClusterStatus(ctx context.Context, request ClusterStatusRequest) (*ClusterLibraryStatuses, error)

	// ClusterStatusByClusterId calls ClusterStatus, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	ClusterStatusByClusterId(ctx context.Context, clusterId string) (*ClusterLibraryStatuses, error)

	// Add a library
	//
	// Add libraries to be installed on a cluster. The installation is
	// asynchronous; it happens in the background after the completion of this
	// request.
	//
	// **Note**: The actual set of libraries to be installed on a cluster is the
	// union of the libraries specified via this method and the libraries set to
	// be installed on all clusters via the libraries UI.
	Install(ctx context.Context, request InstallLibraries) error

	// Uninstall libraries
	//
	// Set libraries to be uninstalled on a cluster. The libraries won't be
	// uninstalled until the cluster is restarted. Uninstalling libraries that
	// are not installed on the cluster will have no impact but is not an error.
	Uninstall(ctx context.Context, request UninstallLibraries) error

	UpdateLibraries(ctx context.Context, update UpdateLibraries) error

	WaitForLibrariesInstalled(ctx context.Context, wait Wait) (*ClusterLibraryStatuses, error)
}
