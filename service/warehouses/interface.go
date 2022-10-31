// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package warehouses

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/retries"
)

// Access the history of queries through SQL warehouses.
//
// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type QueryHistoryService interface {

	// List
	//
	// List the history of queries through SQL warehouses.
	//
	// You can filter by user ID, warehouse ID, status, and time range.
	//
	// Use ListQueriesAll() to get all QueryInfo instances, which will iterate over every result page.
	ListQueries(ctx context.Context, request ListQueriesRequest) (*ListQueriesResponse, error)

	// ListQueriesAll calls ListQueries() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListQueriesAll(ctx context.Context, request ListQueriesRequest) ([]QueryInfo, error)
}

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type WarehousesService interface {

	// Create a warehouse
	//
	// Creates a new SQL warehouse.
	CreateWarehouse(ctx context.Context, request CreateWarehouseRequest) (*CreateWarehouseResponse, error)

	// CreateWarehouseAndWait calls CreateWarehouse() and waits to reach RUNNING state
	//
	// This method is generated by Databricks SDK Code Generator.
	CreateWarehouseAndWait(ctx context.Context, request CreateWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error)

	// Delete a warehouse
	//
	// Deletes a SQL warehouse.
	DeleteWarehouse(ctx context.Context, request DeleteWarehouseRequest) error

	// DeleteWarehouseAndWait calls DeleteWarehouse() and waits to reach DELETED state
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteWarehouseAndWait(ctx context.Context, request DeleteWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error)
	// DeleteWarehouseById calls DeleteWarehouse, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteWarehouseById(ctx context.Context, id string) error

	// DeleteWarehouseByIdAndWait calls DeleteWarehouseById and waits until GetWarehouseResponse is in desired state.
	//
	// This method is generated by Databricks SDK Code Generator.
	DeleteWarehouseByIdAndWait(ctx context.Context, id string, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error)

	// Update a warehouse
	//
	// Updates the configuration for a SQL warehouse.
	EditWarehouse(ctx context.Context, request EditWarehouseRequest) error

	// EditWarehouseAndWait calls EditWarehouse() and waits to reach RUNNING state
	//
	// This method is generated by Databricks SDK Code Generator.
	EditWarehouseAndWait(ctx context.Context, request EditWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error)

	// Get warehouse info
	//
	// Gets the information for a single SQL warehouse.
	GetWarehouse(ctx context.Context, request GetWarehouseRequest) (*GetWarehouseResponse, error)

	// GetWarehouseAndWait calls GetWarehouse() and waits to reach RUNNING state
	//
	// This method is generated by Databricks SDK Code Generator.
	GetWarehouseAndWait(ctx context.Context, request GetWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error)
	// GetWarehouseById calls GetWarehouse, but directly with primitive function arguments,
	// instead of constructing request instance.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetWarehouseById(ctx context.Context, id string) (*GetWarehouseResponse, error)

	// GetWarehouseByIdAndWait calls GetWarehouseById and waits until GetWarehouseResponse is in desired state.
	//
	// This method is generated by Databricks SDK Code Generator.
	GetWarehouseByIdAndWait(ctx context.Context, id string, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error)

	// Get a configuration
	//
	// Gets the workspace level configuration that is shared by all SQL
	// warehouses in a workspace.
	GetWorkspaceWarehouseConfig(ctx context.Context) (*GetWorkspaceWarehouseConfigResponse, error)

	// List warehouses
	//
	// Lists all SQL warehouses that a user has manager permissions on.
	//
	// Use ListWarehousesAll() to get all EndpointInfo instances
	ListWarehouses(ctx context.Context, request ListWarehousesRequest) (*ListWarehousesResponse, error)

	// ListWarehousesAll calls ListWarehouses() to retrieve all available results from the platform.
	//
	// This method is generated by Databricks SDK Code Generator.
	ListWarehousesAll(ctx context.Context, request ListWarehousesRequest) ([]EndpointInfo, error)

	// Set a configuration
	//
	// Sets the workspace level configuration that is shared by all SQL
	// warehouses in a workspace.
	SetWorkspaceWarehouseConfig(ctx context.Context, request SetWorkspaceWarehouseConfigRequest) error

	// Start a warehouse
	//
	// Starts a SQL warehouse.
	StartWarehouse(ctx context.Context, request StartWarehouseRequest) error

	// StartWarehouseAndWait calls StartWarehouse() and waits to reach RUNNING state
	//
	// This method is generated by Databricks SDK Code Generator.
	StartWarehouseAndWait(ctx context.Context, request StartWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error)

	// Stop a warehouse
	//
	// Stops a SQL warehouse.
	StopWarehouse(ctx context.Context, request StopWarehouseRequest) error

	// StopWarehouseAndWait calls StopWarehouse() and waits to reach STOPPED state
	//
	// This method is generated by Databricks SDK Code Generator.
	StopWarehouseAndWait(ctx context.Context, request StopWarehouseRequest, options ...retries.Option[GetWarehouseResponse]) (*GetWarehouseResponse, error)
}
