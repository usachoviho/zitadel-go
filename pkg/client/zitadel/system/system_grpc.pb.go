// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package system

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SystemServiceClient is the client API for SystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemServiceClient interface {
	// Indicates if ZITADEL is running.
	// It respondes as soon as ZITADEL started
	Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error)
	// Returns a list of ZITADEL instances
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	// Returns the detail of an instance
	GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error)
	// Deprecated: Use CreateInstance instead
	// Creates a new instance with all needed setup data
	// This might take some time
	AddInstance(ctx context.Context, in *AddInstanceRequest, opts ...grpc.CallOption) (*AddInstanceResponse, error)
	// Updates name of an existing instance
	UpdateInstance(ctx context.Context, in *UpdateInstanceRequest, opts ...grpc.CallOption) (*UpdateInstanceResponse, error)
	// Creates a new instance with all needed setup data
	// This might take some time
	CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error)
	// Removes a instances
	// This might take some time
	RemoveInstance(ctx context.Context, in *RemoveInstanceRequest, opts ...grpc.CallOption) (*RemoveInstanceResponse, error)
	// Returns all instance members matching the request
	// all queries need to match (ANDed)
	ListIAMMembers(ctx context.Context, in *ListIAMMembersRequest, opts ...grpc.CallOption) (*ListIAMMembersResponse, error)
	// Checks if a domain exists
	ExistsDomain(ctx context.Context, in *ExistsDomainRequest, opts ...grpc.CallOption) (*ExistsDomainResponse, error)
	// Returns the custom domains of an instance
	ListDomains(ctx context.Context, in *ListDomainsRequest, opts ...grpc.CallOption) (*ListDomainsResponse, error)
	// Returns the domain of an instance
	AddDomain(ctx context.Context, in *AddDomainRequest, opts ...grpc.CallOption) (*AddDomainResponse, error)
	// Returns the domain of an instance
	RemoveDomain(ctx context.Context, in *RemoveDomainRequest, opts ...grpc.CallOption) (*RemoveDomainResponse, error)
	// Returns the domain of an instance
	SetPrimaryDomain(ctx context.Context, in *SetPrimaryDomainRequest, opts ...grpc.CallOption) (*SetPrimaryDomainResponse, error)
	// Returns all stored read models of ZITADEL
	// views are used for search optimisation and optimise request latencies
	// they represent the delta of the event happend on the objects
	ListViews(ctx context.Context, in *ListViewsRequest, opts ...grpc.CallOption) (*ListViewsResponse, error)
	// Truncates the delta of the change stream
	// be carefull with this function because ZITADEL has to
	// recompute the deltas after they got cleared.
	// Search requests will return wrong results until all deltas are recomputed
	ClearView(ctx context.Context, in *ClearViewRequest, opts ...grpc.CallOption) (*ClearViewResponse, error)
	// Returns event descriptions which cannot be processed.
	// It's possible that some events need some retries.
	// For example if the SMTP-API wasn't able to send an email at the first time
	ListFailedEvents(ctx context.Context, in *ListFailedEventsRequest, opts ...grpc.CallOption) (*ListFailedEventsResponse, error)
	// Deletes the event from failed events view.
	// the event is not removed from the change stream
	// This call is usefull if the system was able to process the event later.
	// e.g. if the second try of sending an email was successful. the first try produced a
	// failed event. You can find out if it worked on the `failure_count`
	RemoveFailedEvent(ctx context.Context, in *RemoveFailedEventRequest, opts ...grpc.CallOption) (*RemoveFailedEventResponse, error)
}

type systemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemServiceClient(cc grpc.ClientConnInterface) SystemServiceClient {
	return &systemServiceClient{cc}
}

func (c *systemServiceClient) Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error) {
	out := new(HealthzResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/ListInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceResponse, error) {
	out := new(GetInstanceResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/GetInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) AddInstance(ctx context.Context, in *AddInstanceRequest, opts ...grpc.CallOption) (*AddInstanceResponse, error) {
	out := new(AddInstanceResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/AddInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) UpdateInstance(ctx context.Context, in *UpdateInstanceRequest, opts ...grpc.CallOption) (*UpdateInstanceResponse, error) {
	out := new(UpdateInstanceResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/UpdateInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceResponse, error) {
	out := new(CreateInstanceResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/CreateInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) RemoveInstance(ctx context.Context, in *RemoveInstanceRequest, opts ...grpc.CallOption) (*RemoveInstanceResponse, error) {
	out := new(RemoveInstanceResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/RemoveInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ListIAMMembers(ctx context.Context, in *ListIAMMembersRequest, opts ...grpc.CallOption) (*ListIAMMembersResponse, error) {
	out := new(ListIAMMembersResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/ListIAMMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ExistsDomain(ctx context.Context, in *ExistsDomainRequest, opts ...grpc.CallOption) (*ExistsDomainResponse, error) {
	out := new(ExistsDomainResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/ExistsDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ListDomains(ctx context.Context, in *ListDomainsRequest, opts ...grpc.CallOption) (*ListDomainsResponse, error) {
	out := new(ListDomainsResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/ListDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) AddDomain(ctx context.Context, in *AddDomainRequest, opts ...grpc.CallOption) (*AddDomainResponse, error) {
	out := new(AddDomainResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/AddDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) RemoveDomain(ctx context.Context, in *RemoveDomainRequest, opts ...grpc.CallOption) (*RemoveDomainResponse, error) {
	out := new(RemoveDomainResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/RemoveDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) SetPrimaryDomain(ctx context.Context, in *SetPrimaryDomainRequest, opts ...grpc.CallOption) (*SetPrimaryDomainResponse, error) {
	out := new(SetPrimaryDomainResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/SetPrimaryDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ListViews(ctx context.Context, in *ListViewsRequest, opts ...grpc.CallOption) (*ListViewsResponse, error) {
	out := new(ListViewsResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/ListViews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ClearView(ctx context.Context, in *ClearViewRequest, opts ...grpc.CallOption) (*ClearViewResponse, error) {
	out := new(ClearViewResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/ClearView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ListFailedEvents(ctx context.Context, in *ListFailedEventsRequest, opts ...grpc.CallOption) (*ListFailedEventsResponse, error) {
	out := new(ListFailedEventsResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/ListFailedEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) RemoveFailedEvent(ctx context.Context, in *RemoveFailedEventRequest, opts ...grpc.CallOption) (*RemoveFailedEventResponse, error) {
	out := new(RemoveFailedEventResponse)
	err := c.cc.Invoke(ctx, "/zitadel.system.v1.SystemService/RemoveFailedEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServiceServer is the server API for SystemService service.
// All implementations must embed UnimplementedSystemServiceServer
// for forward compatibility
type SystemServiceServer interface {
	// Indicates if ZITADEL is running.
	// It respondes as soon as ZITADEL started
	Healthz(context.Context, *HealthzRequest) (*HealthzResponse, error)
	// Returns a list of ZITADEL instances
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	// Returns the detail of an instance
	GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceResponse, error)
	// Deprecated: Use CreateInstance instead
	// Creates a new instance with all needed setup data
	// This might take some time
	AddInstance(context.Context, *AddInstanceRequest) (*AddInstanceResponse, error)
	// Updates name of an existing instance
	UpdateInstance(context.Context, *UpdateInstanceRequest) (*UpdateInstanceResponse, error)
	// Creates a new instance with all needed setup data
	// This might take some time
	CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error)
	// Removes a instances
	// This might take some time
	RemoveInstance(context.Context, *RemoveInstanceRequest) (*RemoveInstanceResponse, error)
	// Returns all instance members matching the request
	// all queries need to match (ANDed)
	ListIAMMembers(context.Context, *ListIAMMembersRequest) (*ListIAMMembersResponse, error)
	// Checks if a domain exists
	ExistsDomain(context.Context, *ExistsDomainRequest) (*ExistsDomainResponse, error)
	// Returns the custom domains of an instance
	ListDomains(context.Context, *ListDomainsRequest) (*ListDomainsResponse, error)
	// Returns the domain of an instance
	AddDomain(context.Context, *AddDomainRequest) (*AddDomainResponse, error)
	// Returns the domain of an instance
	RemoveDomain(context.Context, *RemoveDomainRequest) (*RemoveDomainResponse, error)
	// Returns the domain of an instance
	SetPrimaryDomain(context.Context, *SetPrimaryDomainRequest) (*SetPrimaryDomainResponse, error)
	// Returns all stored read models of ZITADEL
	// views are used for search optimisation and optimise request latencies
	// they represent the delta of the event happend on the objects
	ListViews(context.Context, *ListViewsRequest) (*ListViewsResponse, error)
	// Truncates the delta of the change stream
	// be carefull with this function because ZITADEL has to
	// recompute the deltas after they got cleared.
	// Search requests will return wrong results until all deltas are recomputed
	ClearView(context.Context, *ClearViewRequest) (*ClearViewResponse, error)
	// Returns event descriptions which cannot be processed.
	// It's possible that some events need some retries.
	// For example if the SMTP-API wasn't able to send an email at the first time
	ListFailedEvents(context.Context, *ListFailedEventsRequest) (*ListFailedEventsResponse, error)
	// Deletes the event from failed events view.
	// the event is not removed from the change stream
	// This call is usefull if the system was able to process the event later.
	// e.g. if the second try of sending an email was successful. the first try produced a
	// failed event. You can find out if it worked on the `failure_count`
	RemoveFailedEvent(context.Context, *RemoveFailedEventRequest) (*RemoveFailedEventResponse, error)
	mustEmbedUnimplementedSystemServiceServer()
}

// UnimplementedSystemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemServiceServer struct {
}

func (UnimplementedSystemServiceServer) Healthz(context.Context, *HealthzRequest) (*HealthzResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthz not implemented")
}
func (UnimplementedSystemServiceServer) ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstances not implemented")
}
func (UnimplementedSystemServiceServer) GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstance not implemented")
}
func (UnimplementedSystemServiceServer) AddInstance(context.Context, *AddInstanceRequest) (*AddInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInstance not implemented")
}
func (UnimplementedSystemServiceServer) UpdateInstance(context.Context, *UpdateInstanceRequest) (*UpdateInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInstance not implemented")
}
func (UnimplementedSystemServiceServer) CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstance not implemented")
}
func (UnimplementedSystemServiceServer) RemoveInstance(context.Context, *RemoveInstanceRequest) (*RemoveInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveInstance not implemented")
}
func (UnimplementedSystemServiceServer) ListIAMMembers(context.Context, *ListIAMMembersRequest) (*ListIAMMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIAMMembers not implemented")
}
func (UnimplementedSystemServiceServer) ExistsDomain(context.Context, *ExistsDomainRequest) (*ExistsDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistsDomain not implemented")
}
func (UnimplementedSystemServiceServer) ListDomains(context.Context, *ListDomainsRequest) (*ListDomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomains not implemented")
}
func (UnimplementedSystemServiceServer) AddDomain(context.Context, *AddDomainRequest) (*AddDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDomain not implemented")
}
func (UnimplementedSystemServiceServer) RemoveDomain(context.Context, *RemoveDomainRequest) (*RemoveDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDomain not implemented")
}
func (UnimplementedSystemServiceServer) SetPrimaryDomain(context.Context, *SetPrimaryDomainRequest) (*SetPrimaryDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPrimaryDomain not implemented")
}
func (UnimplementedSystemServiceServer) ListViews(context.Context, *ListViewsRequest) (*ListViewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListViews not implemented")
}
func (UnimplementedSystemServiceServer) ClearView(context.Context, *ClearViewRequest) (*ClearViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearView not implemented")
}
func (UnimplementedSystemServiceServer) ListFailedEvents(context.Context, *ListFailedEventsRequest) (*ListFailedEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFailedEvents not implemented")
}
func (UnimplementedSystemServiceServer) RemoveFailedEvent(context.Context, *RemoveFailedEventRequest) (*RemoveFailedEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFailedEvent not implemented")
}
func (UnimplementedSystemServiceServer) mustEmbedUnimplementedSystemServiceServer() {}

// UnsafeSystemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemServiceServer will
// result in compilation errors.
type UnsafeSystemServiceServer interface {
	mustEmbedUnimplementedSystemServiceServer()
}

func RegisterSystemServiceServer(s grpc.ServiceRegistrar, srv SystemServiceServer) {
	s.RegisterService(&SystemService_ServiceDesc, srv)
}

func _SystemService_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).Healthz(ctx, req.(*HealthzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ListInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/ListInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ListInstances(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_GetInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).GetInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/GetInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).GetInstance(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_AddInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).AddInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/AddInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).AddInstance(ctx, req.(*AddInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_UpdateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).UpdateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/UpdateInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).UpdateInstance(ctx, req.(*UpdateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_CreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).CreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/CreateInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).CreateInstance(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_RemoveInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).RemoveInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/RemoveInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).RemoveInstance(ctx, req.(*RemoveInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ListIAMMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIAMMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ListIAMMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/ListIAMMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ListIAMMembers(ctx, req.(*ListIAMMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ExistsDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ExistsDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/ExistsDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ExistsDomain(ctx, req.(*ExistsDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ListDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ListDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/ListDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ListDomains(ctx, req.(*ListDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_AddDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).AddDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/AddDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).AddDomain(ctx, req.(*AddDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_RemoveDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).RemoveDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/RemoveDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).RemoveDomain(ctx, req.(*RemoveDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_SetPrimaryDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPrimaryDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).SetPrimaryDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/SetPrimaryDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).SetPrimaryDomain(ctx, req.(*SetPrimaryDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ListViews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListViewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ListViews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/ListViews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ListViews(ctx, req.(*ListViewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ClearView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ClearView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/ClearView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ClearView(ctx, req.(*ClearViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ListFailedEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFailedEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ListFailedEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/ListFailedEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ListFailedEvents(ctx, req.(*ListFailedEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_RemoveFailedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFailedEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).RemoveFailedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zitadel.system.v1.SystemService/RemoveFailedEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).RemoveFailedEvent(ctx, req.(*RemoveFailedEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemService_ServiceDesc is the grpc.ServiceDesc for SystemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zitadel.system.v1.SystemService",
	HandlerType: (*SystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthz",
			Handler:    _SystemService_Healthz_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _SystemService_ListInstances_Handler,
		},
		{
			MethodName: "GetInstance",
			Handler:    _SystemService_GetInstance_Handler,
		},
		{
			MethodName: "AddInstance",
			Handler:    _SystemService_AddInstance_Handler,
		},
		{
			MethodName: "UpdateInstance",
			Handler:    _SystemService_UpdateInstance_Handler,
		},
		{
			MethodName: "CreateInstance",
			Handler:    _SystemService_CreateInstance_Handler,
		},
		{
			MethodName: "RemoveInstance",
			Handler:    _SystemService_RemoveInstance_Handler,
		},
		{
			MethodName: "ListIAMMembers",
			Handler:    _SystemService_ListIAMMembers_Handler,
		},
		{
			MethodName: "ExistsDomain",
			Handler:    _SystemService_ExistsDomain_Handler,
		},
		{
			MethodName: "ListDomains",
			Handler:    _SystemService_ListDomains_Handler,
		},
		{
			MethodName: "AddDomain",
			Handler:    _SystemService_AddDomain_Handler,
		},
		{
			MethodName: "RemoveDomain",
			Handler:    _SystemService_RemoveDomain_Handler,
		},
		{
			MethodName: "SetPrimaryDomain",
			Handler:    _SystemService_SetPrimaryDomain_Handler,
		},
		{
			MethodName: "ListViews",
			Handler:    _SystemService_ListViews_Handler,
		},
		{
			MethodName: "ClearView",
			Handler:    _SystemService_ClearView_Handler,
		},
		{
			MethodName: "ListFailedEvents",
			Handler:    _SystemService_ListFailedEvents_Handler,
		},
		{
			MethodName: "RemoveFailedEvent",
			Handler:    _SystemService_RemoveFailedEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zitadel/system.proto",
}
