package gosatellite

import (
	"context"
	"fmt"
	"net/http"
)

const filtersPath = basePath + "/filters"

type filterPermission struct {
	ID           *int    `json:"id"`
	Name         *string `json:"name"`
	ResourceType *string `json:"resource_type"`
}

type filterRole struct {
	Description *string `json:"description"`
	ID          *int    `json:"id"`
	Name        *string `json:"name"`
	Origin      *string `json:"origin"`
}

// Filter defines model for a Filter.
type Filter struct {
	CreatedAt     *string             `json:"created_at"`
	ID            *int                `json:"id"`
	Locations     *[]genericReference `json:"locations"`
	Organizations *[]genericReference `json:"organizations"`
	Override      *bool               `json:"override?"`
	Permissions   *[]filterPermission `json:"permissions"`
	ResourceType  *string             `json:"resource_type"`
	Role          *filterRole         `json:"role"`
	Search        *string             `json:"search"`
	Unlimited     *bool               `json:"unlimited?"`
	UpdatedAt     *string             `json:"updated_at"`
}

// FilterCreate defines model for the body of the creation of a filter.
type FilterCreate struct {
	Filter struct {
		RoleID          *int    `json:"role_id"`
		Search          *string `json:"search,omitempty"`
		Override        *bool   `json:"override,omitempty"`
		PermissionIDs   *[]int  `json:"permission_ids,omitempty"`
		OrganizationIDs *[]int  `json:"organization_ids,omitempty"`
		LocationIDs     *[]int  `json:"location_ids,omitempty"`
	} `json:"filter"`
}

// FilterUpdate defines model for the body of the update of a filter.
type FilterUpdate struct {
	Filter struct {
		RoleID          *int    `json:"role_id,omitempty"`
		Search          *string `json:"search,omitempty"`
		Override        *bool   `json:"override,omitempty"`
		PermissionIDs   *[]int  `json:"permission_ids,omitempty"`
		OrganizationIDs *[]int  `json:"organization_ids,omitempty"`
		LocationIDs     *[]int  `json:"location_ids,omitempty"`
	} `json:"filter"`
}

// Filters is an interface for interacting with
// Red Hat Satellite role filters
type Filters interface {
	Create(ctx context.Context, filterCreate FilterCreate) (*Filter, *http.Response, error)
	Delete(ctx context.Context, filterID int) (*http.Response, error)
	Get(ctx context.Context, filterID int) (*Filter, *http.Response, error)
	Update(ctx context.Context, filterID int, filterUpdate FilterUpdate) (*Filter, *http.Response, error)
}

// FiltersOp handles communication with the Filter related methods of the
// Red Hat Satellite REST API
type FiltersOp struct {
	client *Client
}

// Create a new filter
func (s *FiltersOp) Create(ctx context.Context, filterCreate FilterCreate) (*Filter, *http.Response, error) {
	path := filtersPath

	if filterCreate.Filter.RoleID == nil {
		return nil, nil, NewArgError("filterCreate.RoleIDs", "cannot be empty")
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, filterCreate)
	if err != nil {
		return nil, nil, err
	}
	filter := new(Filter)
	resp, err := s.client.Do(ctx, req, filter)
	if err != nil {
		return nil, resp, err
	}

	return filter, resp, err
}

// Delete a single filter by its ID
func (s *FiltersOp) Delete(ctx context.Context, filterID int) (*http.Response, error) {
	path := fmt.Sprintf("%s/%d", filtersPath, filterID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Get a single filter by its ID
func (s *FiltersOp) Get(ctx context.Context, filterID int) (*Filter, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", filtersPath, filterID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	filter := new(Filter)
	resp, err := s.client.Do(ctx, req, filter)
	if err != nil {
		return nil, resp, err
	}

	return filter, resp, err
}

// Update a filter
func (s *FiltersOp) Update(ctx context.Context, filterID int, filterUpdate FilterUpdate) (*Filter, *http.Response, error) {
	path := fmt.Sprintf("%s/%d", filtersPath, filterID)

	req, err := s.client.NewRequest(ctx, http.MethodPut, path, filterUpdate)
	if err != nil {
		return nil, nil, err
	}
	filter := new(Filter)
	resp, err := s.client.Do(ctx, req, filter)
	if err != nil {
		return nil, resp, err
	}

	return filter, resp, err
}
