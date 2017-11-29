package client

import (
	"github.com/rancher/norman/types"
)

const (
	PersistentVolumeClaimType                      = "persistentVolumeClaim"
	PersistentVolumeClaimFieldAccessModes          = "accessModes"
	PersistentVolumeClaimFieldAnnotations          = "annotations"
	PersistentVolumeClaimFieldCreated              = "created"
	PersistentVolumeClaimFieldLabels               = "labels"
	PersistentVolumeClaimFieldName                 = "name"
	PersistentVolumeClaimFieldOwnerReferences      = "ownerReferences"
	PersistentVolumeClaimFieldRemoved              = "removed"
	PersistentVolumeClaimFieldResourcePath         = "resourcePath"
	PersistentVolumeClaimFieldResources            = "resources"
	PersistentVolumeClaimFieldSelector             = "selector"
	PersistentVolumeClaimFieldState                = "state"
	PersistentVolumeClaimFieldStatus               = "status"
	PersistentVolumeClaimFieldStorageClassName     = "storageClassName"
	PersistentVolumeClaimFieldTransitioning        = "transitioning"
	PersistentVolumeClaimFieldTransitioningMessage = "transitioningMessage"
	PersistentVolumeClaimFieldUuid                 = "uuid"
	PersistentVolumeClaimFieldVolumeName           = "volumeName"
)

type PersistentVolumeClaim struct {
	types.Resource
	AccessModes          []string                     `json:"accessModes,omitempty"`
	Annotations          map[string]string            `json:"annotations,omitempty"`
	Created              string                       `json:"created,omitempty"`
	Labels               map[string]string            `json:"labels,omitempty"`
	Name                 string                       `json:"name,omitempty"`
	OwnerReferences      []OwnerReference             `json:"ownerReferences,omitempty"`
	Removed              string                       `json:"removed,omitempty"`
	ResourcePath         string                       `json:"resourcePath,omitempty"`
	Resources            *ResourceRequirements        `json:"resources,omitempty"`
	Selector             *LabelSelector               `json:"selector,omitempty"`
	State                string                       `json:"state,omitempty"`
	Status               *PersistentVolumeClaimStatus `json:"status,omitempty"`
	StorageClassName     string                       `json:"storageClassName,omitempty"`
	Transitioning        string                       `json:"transitioning,omitempty"`
	TransitioningMessage string                       `json:"transitioningMessage,omitempty"`
	Uuid                 string                       `json:"uuid,omitempty"`
	VolumeName           string                       `json:"volumeName,omitempty"`
}
type PersistentVolumeClaimCollection struct {
	types.Collection
	Data   []PersistentVolumeClaim `json:"data,omitempty"`
	client *PersistentVolumeClaimClient
}

type PersistentVolumeClaimClient struct {
	apiClient *Client
}

type PersistentVolumeClaimOperations interface {
	List(opts *types.ListOpts) (*PersistentVolumeClaimCollection, error)
	Create(opts *PersistentVolumeClaim) (*PersistentVolumeClaim, error)
	Update(existing *PersistentVolumeClaim, updates interface{}) (*PersistentVolumeClaim, error)
	ByID(id string) (*PersistentVolumeClaim, error)
	Delete(container *PersistentVolumeClaim) error
}

func newPersistentVolumeClaimClient(apiClient *Client) *PersistentVolumeClaimClient {
	return &PersistentVolumeClaimClient{
		apiClient: apiClient,
	}
}

func (c *PersistentVolumeClaimClient) Create(container *PersistentVolumeClaim) (*PersistentVolumeClaim, error) {
	resp := &PersistentVolumeClaim{}
	err := c.apiClient.Ops.DoCreate(PersistentVolumeClaimType, container, resp)
	return resp, err
}

func (c *PersistentVolumeClaimClient) Update(existing *PersistentVolumeClaim, updates interface{}) (*PersistentVolumeClaim, error) {
	resp := &PersistentVolumeClaim{}
	err := c.apiClient.Ops.DoUpdate(PersistentVolumeClaimType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PersistentVolumeClaimClient) List(opts *types.ListOpts) (*PersistentVolumeClaimCollection, error) {
	resp := &PersistentVolumeClaimCollection{}
	err := c.apiClient.Ops.DoList(PersistentVolumeClaimType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PersistentVolumeClaimCollection) Next() (*PersistentVolumeClaimCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PersistentVolumeClaimCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PersistentVolumeClaimClient) ByID(id string) (*PersistentVolumeClaim, error) {
	resp := &PersistentVolumeClaim{}
	err := c.apiClient.Ops.DoByID(PersistentVolumeClaimType, id, resp)
	return resp, err
}

func (c *PersistentVolumeClaimClient) Delete(container *PersistentVolumeClaim) error {
	return c.apiClient.Ops.DoResourceDelete(PersistentVolumeClaimType, &container.Resource)
}
