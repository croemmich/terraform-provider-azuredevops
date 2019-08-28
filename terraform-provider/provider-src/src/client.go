package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/build"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/operations"
)

// AggregatedClient Aggregates all of the underlying clients into a single data
// type. Each client is ready to use and fully configured with the correct
// AzDO PAT/organization
type AggregatedClient struct {
	CoreClient       CoreClient
	BuildClient      *build.Client
	OperationsClient *operations.Client
}

type CoreClient interface {
	//RemoveProjectAvatar(ctx context.Context, args core.RemoveProjectAvatarArgs) error
	//SetProjectAvatar(ctx context.Context, args core.SetProjectAvatarArgs) error
	// CreateConnectedService(ctx context.Context, args core.CreateConnectedServiceArgs) (*WebApiConnectedService, error)
	// GetConnectedServiceDetails(ctx context.Context, args core.GetConnectedServiceDetailsArgs) (*WebApiConnectedServiceDetails, error)
	// GetConnectedServices(ctx context.Context, args core.GetConnectedServicesArgs) (*[]WebApiConnectedService, error)
	// GetTeamMembersWithExtendedProperties(ctx context.Context, args core.GetTeamMembersWithExtendedPropertiesArgs) (*[]webapi.TeamMember, error)
	//GetProcessById(ctx context.Context, args GetProcessByIdArgs) (*Process, error)
	GetProcesses(ctx context.Context, args core.GetProcessesArgs) (*[]core.Process, error)
	QueueCreateProject(ctx context.Context, args core.QueueCreateProjectArgs) (*operations.OperationReference, error)
	GetProjects(ctx context.Context, args core.GetProjectsArgs) (*core.GetProjectsResponseValue, error)
	
/**	{
		routeValues := make(map[string]string)
		if args.ProcessId == nil {
			return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ProcessId"}
		}
		routeValues["processId"] = (*args.ProcessId).String()
	
		locationId, _ := uuid.Parse("93878975-88c5-4e6a-8abb-7ddd77a8a7d8")
		resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue Process
		err = client.Client.UnmarshalBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the GetProcessById function
	type GetProcessByIdArgs struct {
		// (required) ID for a process.
		ProcessId *uuid.UUID
	}
	
	// Get a list of processes.
	func (client *Client) GetProcesses(ctx context.Context, args GetProcessesArgs) (*[]Process, error) {
		locationId, _ := uuid.Parse("93878975-88c5-4e6a-8abb-7ddd77a8a7d8")
		resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue []Process
		err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the GetProcesses function
	type GetProcessesArgs struct {
	}
	
	// Get project collection with the specified id or name.
	func (client *Client) GetProjectCollection(ctx context.Context, args GetProjectCollectionArgs) (*TeamProjectCollection, error) {
		routeValues := make(map[string]string)
		if args.CollectionId == nil || *args.CollectionId == "" {
			return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.CollectionId"}
		}
		routeValues["collectionId"] = *args.CollectionId
	
		locationId, _ := uuid.Parse("8031090f-ef1d-4af6-85fc-698cd75d42bf")
		resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue TeamProjectCollection
		err = client.Client.UnmarshalBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the GetProjectCollection function
	type GetProjectCollectionArgs struct {
		// (required)
		CollectionId *string
	}
	
	// Get project collection references for this application.
	func (client *Client) GetProjectCollections(ctx context.Context, args GetProjectCollectionsArgs) (*[]TeamProjectCollectionReference, error) {
		queryParams := url.Values{}
		if args.Top != nil {
			queryParams.Add("$top", strconv.Itoa(*args.Top))
		}
		if args.Skip != nil {
			queryParams.Add("$skip", strconv.Itoa(*args.Skip))
		}
		locationId, _ := uuid.Parse("8031090f-ef1d-4af6-85fc-698cd75d42bf")
		resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue []TeamProjectCollectionReference
		err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the GetProjectCollections function
	type GetProjectCollectionsArgs struct {
		// (optional)
		Top *int
		// (optional)
		Skip *int
	}
	
	// Get project with the specified id or name, optionally including capabilities.
	func (client *Client) GetProject(ctx context.Context, args GetProjectArgs) (*TeamProject, error) {
		routeValues := make(map[string]string)
		if args.ProjectId == nil || *args.ProjectId == "" {
			return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProjectId"}
		}
		routeValues["projectId"] = *args.ProjectId
	
		queryParams := url.Values{}
		if args.IncludeCapabilities != nil {
			queryParams.Add("includeCapabilities", strconv.FormatBool(*args.IncludeCapabilities))
		}
		if args.IncludeHistory != nil {
			queryParams.Add("includeHistory", strconv.FormatBool(*args.IncludeHistory))
		}
		locationId, _ := uuid.Parse("603fe2ac-9723-48b9-88ad-09305aa6c6e1")
		resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue TeamProject
		err = client.Client.UnmarshalBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the GetProject function
	type GetProjectArgs struct {
		// (required)
		ProjectId *string
		// (optional) Include capabilities (such as source control) in the team project result (default: false).
		IncludeCapabilities *bool
		// (optional) Search within renamed projects (that had such name in the past).
		IncludeHistory *bool
	}
	
	// Get all projects in the organization that the authenticated user has access to.
	func (client *Client) GetProjects(ctx context.Context, args GetProjectsArgs) (*GetProjectsResponseValue, error) {
		queryParams := url.Values{}
		if args.StateFilter != nil {
			queryParams.Add("stateFilter", string(*args.StateFilter))
		}
		if args.Top != nil {
			queryParams.Add("$top", strconv.Itoa(*args.Top))
		}
		if args.Skip != nil {
			queryParams.Add("$skip", strconv.Itoa(*args.Skip))
		}
		if args.ContinuationToken != nil {
			queryParams.Add("continuationToken", *args.ContinuationToken)
		}
		if args.GetDefaultTeamImageUrl != nil {
			queryParams.Add("getDefaultTeamImageUrl", strconv.FormatBool(*args.GetDefaultTeamImageUrl))
		}
		locationId, _ := uuid.Parse("603fe2ac-9723-48b9-88ad-09305aa6c6e1")
		resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue GetProjectsResponseValue
		responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
		err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
		return &responseValue, err
	}
	
	// Arguments for the GetProjects function
	type GetProjectsArgs struct {
		// (optional) Filter on team projects in a specific team project state (default: WellFormed).
		StateFilter *ProjectState
		// (optional)
		Top *int
		// (optional)
		Skip *int
		// (optional)
		ContinuationToken *string
		// (optional)
		GetDefaultTeamImageUrl *bool
	}
	
	// Return type for the GetProjects function
	type GetProjectsResponseValue struct {
		Value []TeamProjectReference
		// The continuation token to be used to get the next page of results.
		ContinuationToken string
	}
	
	// Queues a project to be created. Use the [GetOperation](../../operations/operations/get) to periodically check for create project status.
	func (client *Client) QueueCreateProject(ctx context.Context, args QueueCreateProjectArgs) (*operations.OperationReference, error) {
		if args.ProjectToCreate == nil {
			return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ProjectToCreate"}
		}
		body, marshalErr := json.Marshal(*args.ProjectToCreate)
		if marshalErr != nil {
			return nil, marshalErr
		}
		locationId, _ := uuid.Parse("603fe2ac-9723-48b9-88ad-09305aa6c6e1")
		resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue operations.OperationReference
		err = client.Client.UnmarshalBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the QueueCreateProject function
	type QueueCreateProjectArgs struct {
		// (required) The project to create.
		ProjectToCreate *TeamProject
	}
	
	// Queues a project to be deleted. Use the [GetOperation](../../operations/operations/get) to periodically check for delete project status.
	func (client *Client) QueueDeleteProject(ctx context.Context, args QueueDeleteProjectArgs) (*operations.OperationReference, error) {
		routeValues := make(map[string]string)
		if args.ProjectId == nil {
			return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ProjectId"}
		}
		routeValues["projectId"] = (*args.ProjectId).String()
	
		locationId, _ := uuid.Parse("603fe2ac-9723-48b9-88ad-09305aa6c6e1")
		resp, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue operations.OperationReference
		err = client.Client.UnmarshalBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the QueueDeleteProject function
	type QueueDeleteProjectArgs struct {
		// (required) The project id of the project to delete.
		ProjectId *uuid.UUID
	}
	
	// Update an existing project's name, abbreviation, description, or restore a project.
	func (client *Client) UpdateProject(ctx context.Context, args UpdateProjectArgs) (*operations.OperationReference, error) {
		if args.ProjectUpdate == nil {
			return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ProjectUpdate"}
		}
		routeValues := make(map[string]string)
		if args.ProjectId == nil {
			return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ProjectId"}
		}
		routeValues["projectId"] = (*args.ProjectId).String()
	
		body, marshalErr := json.Marshal(*args.ProjectUpdate)
		if marshalErr != nil {
			return nil, marshalErr
		}
		locationId, _ := uuid.Parse("603fe2ac-9723-48b9-88ad-09305aa6c6e1")
		resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue operations.OperationReference
		err = client.Client.UnmarshalBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the UpdateProject function
	type UpdateProjectArgs struct {
		// (required) The updates for the project. The state must be set to wellFormed to restore the project.
		ProjectUpdate *TeamProject
		// (required) The project id of the project to update.
		ProjectId *uuid.UUID
	}
	
	// [Preview API] Get a collection of team project properties.
	func (client *Client) GetProjectProperties(ctx context.Context, args GetProjectPropertiesArgs) (*[]ProjectProperty, error) {
		routeValues := make(map[string]string)
		if args.ProjectId == nil {
			return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ProjectId"}
		}
		routeValues["projectId"] = (*args.ProjectId).String()
	
		queryParams := url.Values{}
		if args.Keys != nil {
			listAsString := strings.Join((*args.Keys)[:], ",")
			queryParams.Add("keys", listAsString)
		}
		locationId, _ := uuid.Parse("4976a71a-4487-49aa-8aab-a1eda469037a")
		resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue []ProjectProperty
		err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the GetProjectProperties function
	type GetProjectPropertiesArgs struct {
		// (required) The team project ID.
		ProjectId *uuid.UUID
		// (optional) A comma-delimited string of team project property names. Wildcard characters ("?" and "*") are supported. If no key is specified, all properties will be returned.
		Keys *[]string
	}
	
	// [Preview API] Create, update, and delete team project properties.
	func (client *Client) SetProjectProperties(ctx context.Context, args SetProjectPropertiesArgs) error {
		if args.PatchDocument == nil {
			return &azuredevops.ArgumentNilError{ArgumentName: "args.PatchDocument"}
		}
		routeValues := make(map[string]string)
		if args.ProjectId == nil {
			return &azuredevops.ArgumentNilError{ArgumentName: "args.ProjectId"}
		}
		routeValues["projectId"] = (*args.ProjectId).String()
	
		body, marshalErr := json.Marshal(*args.PatchDocument)
		if marshalErr != nil {
			return marshalErr
		}
		locationId, _ := uuid.Parse("4976a71a-4487-49aa-8aab-a1eda469037a")
		_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
		if err != nil {
			return err
		}
	
		return nil
	}
	
	// Arguments for the SetProjectProperties function
	type SetProjectPropertiesArgs struct {
		// (required) The team project ID.
		ProjectId *uuid.UUID
		// (required) A JSON Patch document that represents an array of property operations. See RFC 6902 for more details on JSON Patch. The accepted operation verbs are Add and Remove, where Add is used for both creating and updating properties. The path consists of a forward slash and a property name.
		PatchDocument *[]webapi.JsonPatchOperation
	}
	
	// [Preview API]
	func (client *Client) CreateOrUpdateProxy(ctx context.Context, args CreateOrUpdateProxyArgs) (*Proxy, error) {
		if args.Proxy == nil {
			return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Proxy"}
		}
		body, marshalErr := json.Marshal(*args.Proxy)
		if marshalErr != nil {
			return nil, marshalErr
		}
		locationId, _ := uuid.Parse("ec1f4311-f2b4-4c15-b2b8-8990b80d2908")
		resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.2", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue Proxy
		err = client.Client.UnmarshalBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the CreateOrUpdateProxy function
	type CreateOrUpdateProxyArgs struct {
		// (required)
		Proxy *Proxy
	}
	
	// [Preview API]
	func (client *Client) DeleteProxy(ctx context.Context, args DeleteProxyArgs) error {
		queryParams := url.Values{}
		if args.ProxyUrl == nil {
			return &azuredevops.ArgumentNilError{ArgumentName: "proxyUrl"}
		}
		queryParams.Add("proxyUrl", *args.ProxyUrl)
		if args.Site != nil {
			queryParams.Add("site", *args.Site)
		}
		locationId, _ := uuid.Parse("ec1f4311-f2b4-4c15-b2b8-8990b80d2908")
		_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.2", nil, queryParams, nil, "", "application/json", nil)
		if err != nil {
			return err
		}
	
		return nil
	}
	
	// Arguments for the DeleteProxy function
	type DeleteProxyArgs struct {
		// (required)
		ProxyUrl *string
		// (optional)
		Site *string
	}
	
	// [Preview API]
	func (client *Client) GetProxies(ctx context.Context, args GetProxiesArgs) (*[]Proxy, error) {
		queryParams := url.Values{}
		if args.ProxyUrl != nil {
			queryParams.Add("proxyUrl", *args.ProxyUrl)
		}
		locationId, _ := uuid.Parse("ec1f4311-f2b4-4c15-b2b8-8990b80d2908")
		resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.2", nil, queryParams, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue []Proxy
		err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the GetProxies function
	type GetProxiesArgs struct {
		// (optional)
		ProxyUrl *string
	}
	
	// Create a team in a team project.
	func (client *Client) CreateTeam(ctx context.Context, args CreateTeamArgs) (*WebApiTeam, error) {
		if args.Team == nil {
			return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Team"}
		}
		routeValues := make(map[string]string)
		if args.ProjectId == nil || *args.ProjectId == "" {
			return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProjectId"}
		}
		routeValues["projectId"] = *args.ProjectId
	
		body, marshalErr := json.Marshal(*args.Team)
		if marshalErr != nil {
			return nil, marshalErr
		}
		locationId, _ := uuid.Parse("d30a3dd1-f8ba-442a-b86a-bd0c0c383e59")
		resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue WebApiTeam
		err = client.Client.UnmarshalBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the CreateTeam function
	type CreateTeamArgs struct {
		// (required) The team data used to create the team.
		Team *WebApiTeam
		// (required) The name or ID (GUID) of the team project in which to create the team.
		ProjectId *string
	}
	
	// Delete a team.
	func (client *Client) DeleteTeam(ctx context.Context, args DeleteTeamArgs) error {
		routeValues := make(map[string]string)
		if args.ProjectId == nil || *args.ProjectId == "" {
			return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProjectId"}
		}
		routeValues["projectId"] = *args.ProjectId
		if args.TeamId == nil || *args.TeamId == "" {
			return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.TeamId"}
		}
		routeValues["teamId"] = *args.TeamId
	
		locationId, _ := uuid.Parse("d30a3dd1-f8ba-442a-b86a-bd0c0c383e59")
		_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
		if err != nil {
			return err
		}
	
		return nil
	}
	
	// Arguments for the DeleteTeam function
	type DeleteTeamArgs struct {
		// (required) The name or ID (GUID) of the team project containing the team to delete.
		ProjectId *string
		// (required) The name or ID of the team to delete.
		TeamId *string
	}
	
	// Get a specific team.
	func (client *Client) GetTeam(ctx context.Context, args GetTeamArgs) (*WebApiTeam, error) {
		routeValues := make(map[string]string)
		if args.ProjectId == nil || *args.ProjectId == "" {
			return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProjectId"}
		}
		routeValues["projectId"] = *args.ProjectId
		if args.TeamId == nil || *args.TeamId == "" {
			return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.TeamId"}
		}
		routeValues["teamId"] = *args.TeamId
	
		queryParams := url.Values{}
		if args.ExpandIdentity != nil {
			queryParams.Add("$expandIdentity", strconv.FormatBool(*args.ExpandIdentity))
		}
		locationId, _ := uuid.Parse("d30a3dd1-f8ba-442a-b86a-bd0c0c383e59")
		resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue WebApiTeam
		err = client.Client.UnmarshalBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the GetTeam function
	type GetTeamArgs struct {
		// (required) The name or ID (GUID) of the team project containing the team.
		ProjectId *string
		// (required) The name or ID (GUID) of the team.
		TeamId *string
		// (optional) A value indicating whether or not to expand Identity information in the result WebApiTeam object.
		ExpandIdentity *bool
	}
	
	// Get a list of teams.
	func (client *Client) GetTeams(ctx context.Context, args GetTeamsArgs) (*[]WebApiTeam, error) {
		routeValues := make(map[string]string)
		if args.ProjectId == nil || *args.ProjectId == "" {
			return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProjectId"}
		}
		routeValues["projectId"] = *args.ProjectId
	
		queryParams := url.Values{}
		if args.Mine != nil {
			queryParams.Add("$mine", strconv.FormatBool(*args.Mine))
		}
		if args.Top != nil {
			queryParams.Add("$top", strconv.Itoa(*args.Top))
		}
		if args.Skip != nil {
			queryParams.Add("$skip", strconv.Itoa(*args.Skip))
		}
		if args.ExpandIdentity != nil {
			queryParams.Add("$expandIdentity", strconv.FormatBool(*args.ExpandIdentity))
		}
		locationId, _ := uuid.Parse("d30a3dd1-f8ba-442a-b86a-bd0c0c383e59")
		resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue []WebApiTeam
		err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the GetTeams function
	type GetTeamsArgs struct {
		// (required)
		ProjectId *string
		// (optional) If true return all the teams requesting user is member, otherwise return all the teams user has read access.
		Mine *bool
		// (optional) Maximum number of teams to return.
		Top *int
		// (optional) Number of teams to skip.
		Skip *int
		// (optional) A value indicating whether or not to expand Identity information in the result WebApiTeam object.
		ExpandIdentity *bool
	}
	
	// Update a team's name and/or description.
	func (client *Client) UpdateTeam(ctx context.Context, args UpdateTeamArgs) (*WebApiTeam, error) {
		if args.TeamData == nil {
			return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TeamData"}
		}
		routeValues := make(map[string]string)
		if args.ProjectId == nil || *args.ProjectId == "" {
			return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ProjectId"}
		}
		routeValues["projectId"] = *args.ProjectId
		if args.TeamId == nil || *args.TeamId == "" {
			return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.TeamId"}
		}
		routeValues["teamId"] = *args.TeamId
	
		body, marshalErr := json.Marshal(*args.TeamData)
		if marshalErr != nil {
			return nil, marshalErr
		}
		locationId, _ := uuid.Parse("d30a3dd1-f8ba-442a-b86a-bd0c0c383e59")
		resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue WebApiTeam
		err = client.Client.UnmarshalBody(resp, &responseValue)
		return &responseValue, err
	}
	
	// Arguments for the UpdateTeam function
	type UpdateTeamArgs struct {
		// (required)
		TeamData *WebApiTeam
		// (required) The name or ID (GUID) of the team project containing the team to update.
		ProjectId *string
		// (required) The name of ID of the team to update.
		TeamId *string
	}
	
	// [Preview API] Get a list of all teams.
	func (client *Client) GetAllTeams(ctx context.Context, args GetAllTeamsArgs) (*[]WebApiTeam, error) {
		queryParams := url.Values{}
		if args.Mine != nil {
			queryParams.Add("$mine", strconv.FormatBool(*args.Mine))
		}
		if args.Top != nil {
			queryParams.Add("$top", strconv.Itoa(*args.Top))
		}
		if args.Skip != nil {
			queryParams.Add("$skip", strconv.Itoa(*args.Skip))
		}
		if args.ExpandIdentity != nil {
			queryParams.Add("$expandIdentity", strconv.FormatBool(*args.ExpandIdentity))
		}
		locationId, _ := uuid.Parse("7a4d9ee9-3433-4347-b47a-7a80f1cf307e")
		resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.3", nil, queryParams, nil, "", "application/json", nil)
		if err != nil {
			return nil, err
		}
	
		var responseValue []WebApiTeam
		err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
		return &responseValue, err
	}**/
}

// Instantiates clients that talk to Azure DevOps
func getClients(ctx context.Context) (*AggregatedClient, error) {
	azdoOrg := os.Getenv("AZDO_ORGANIZATION")
	azdoPAT := os.Getenv("AZDO_PAT")

	if azdoOrg == "" {
		return nil, fmt.Errorf("organization is missing")
	}

	if azdoPAT == "" {
		return nil, fmt.Errorf("access token is missing")
	}

	organizationURL := "https://dev.azure.com/" + os.Getenv("AZDO_ORGANIZATION")
	connection := azuredevops.NewPatConnection(organizationURL, azdoPAT)

	// client for these APIs (includes CRUD for AzDO projects...):
	//	https://docs.microsoft.com/en-us/rest/api/azure/devops/core/?view=azure-devops-rest-5.1
	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		return nil, err
	}

	// client for these APIs (includes CRUD for AzDO build pipelines...):
	//	https://docs.microsoft.com/en-us/rest/api/azure/devops/build/?view=azure-devops-rest-5.1
	buildClient, err := build.NewClient(ctx, connection)
	if err != nil {
		return nil, err
	}

	// client for these APIs (monitor async operations...):
	//	https://docs.microsoft.com/en-us/rest/api/azure/devops/operations/operations?view=azure-devops-rest-5.1
	operationsClient := operations.NewClient(ctx, connection)

	aggregatedClient := &AggregatedClient{
		CoreClient:       coreClient,
		BuildClient:      buildClient,
		OperationsClient: operationsClient,
	}

	log.Printf("Created clients successfully!")
	return aggregatedClient, nil
}