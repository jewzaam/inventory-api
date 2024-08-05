// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             (unknown)
// source: kessel/inventory/v1beta1/k8s_clusters_service.proto

package v1beta1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationKesselK8SClusterServiceCreateK8sCluster = "/kessel.inventory.v1beta1.KesselK8sClusterService/CreateK8sCluster"
const OperationKesselK8SClusterServiceDeleteK8sCluster = "/kessel.inventory.v1beta1.KesselK8sClusterService/DeleteK8sCluster"
const OperationKesselK8SClusterServiceUpdateK8sCluster = "/kessel.inventory.v1beta1.KesselK8sClusterService/UpdateK8sCluster"

type KesselK8SClusterServiceHTTPServer interface {
	CreateK8SCluster(context.Context, *CreateK8SClusterRequest) (*CreateK8SClusterResponse, error)
	DeleteK8SCluster(context.Context, *DeleteK8SClusterRequest) (*DeleteK8SClusterResponse, error)
	UpdateK8SCluster(context.Context, *UpdateK8SClusterRequest) (*UpdateK8SClusterResponse, error)
}

func RegisterKesselK8SClusterServiceHTTPServer(s *http.Server, srv KesselK8SClusterServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/inventory/v1beta1/k8sClusters", _KesselK8SClusterService_CreateK8SCluster0_HTTP_Handler(srv))
	r.PUT("/api/inventory/v1beta1/k8sClusters/{resource}", _KesselK8SClusterService_UpdateK8SCluster0_HTTP_Handler(srv))
	r.DELETE("/api/inventory/v1beta1/k8sClusters/{resource}", _KesselK8SClusterService_DeleteK8SCluster0_HTTP_Handler(srv))
}

func _KesselK8SClusterService_CreateK8SCluster0_HTTP_Handler(srv KesselK8SClusterServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateK8SClusterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKesselK8SClusterServiceCreateK8sCluster)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateK8SCluster(ctx, req.(*CreateK8SClusterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateK8SClusterResponse)
		return ctx.Result(200, reply)
	}
}

func _KesselK8SClusterService_UpdateK8SCluster0_HTTP_Handler(srv KesselK8SClusterServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateK8SClusterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKesselK8SClusterServiceUpdateK8sCluster)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateK8SCluster(ctx, req.(*UpdateK8SClusterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateK8SClusterResponse)
		return ctx.Result(200, reply)
	}
}

func _KesselK8SClusterService_DeleteK8SCluster0_HTTP_Handler(srv KesselK8SClusterServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteK8SClusterRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationKesselK8SClusterServiceDeleteK8sCluster)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteK8SCluster(ctx, req.(*DeleteK8SClusterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteK8SClusterResponse)
		return ctx.Result(200, reply)
	}
}

type KesselK8SClusterServiceHTTPClient interface {
	CreateK8SCluster(ctx context.Context, req *CreateK8SClusterRequest, opts ...http.CallOption) (rsp *CreateK8SClusterResponse, err error)
	DeleteK8SCluster(ctx context.Context, req *DeleteK8SClusterRequest, opts ...http.CallOption) (rsp *DeleteK8SClusterResponse, err error)
	UpdateK8SCluster(ctx context.Context, req *UpdateK8SClusterRequest, opts ...http.CallOption) (rsp *UpdateK8SClusterResponse, err error)
}

type KesselK8SClusterServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewKesselK8SClusterServiceHTTPClient(client *http.Client) KesselK8SClusterServiceHTTPClient {
	return &KesselK8SClusterServiceHTTPClientImpl{client}
}

func (c *KesselK8SClusterServiceHTTPClientImpl) CreateK8SCluster(ctx context.Context, in *CreateK8SClusterRequest, opts ...http.CallOption) (*CreateK8SClusterResponse, error) {
	var out CreateK8SClusterResponse
	pattern := "/api/inventory/v1beta1/k8sClusters"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKesselK8SClusterServiceCreateK8sCluster))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *KesselK8SClusterServiceHTTPClientImpl) DeleteK8SCluster(ctx context.Context, in *DeleteK8SClusterRequest, opts ...http.CallOption) (*DeleteK8SClusterResponse, error) {
	var out DeleteK8SClusterResponse
	pattern := "/api/inventory/v1beta1/k8sClusters/{resource}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationKesselK8SClusterServiceDeleteK8sCluster))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *KesselK8SClusterServiceHTTPClientImpl) UpdateK8SCluster(ctx context.Context, in *UpdateK8SClusterRequest, opts ...http.CallOption) (*UpdateK8SClusterResponse, error) {
	var out UpdateK8SClusterResponse
	pattern := "/api/inventory/v1beta1/k8sClusters/{resource}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationKesselK8SClusterServiceUpdateK8sCluster))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
