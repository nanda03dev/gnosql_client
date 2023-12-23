// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: gnosql.proto

package proto

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

// GnoSQLServiceClient is the client API for GnoSQLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GnoSQLServiceClient interface {
	CreateNewDatabase(ctx context.Context, in *DatabaseCreateRequest, opts ...grpc.CallOption) (*DatabaseCreateResponse, error)
	DeleteDatabase(ctx context.Context, in *DatabaseDeleteRequest, opts ...grpc.CallOption) (*DatabaseDeleteResponse, error)
	GetAllDatabases(ctx context.Context, in *NoRequestBody, opts ...grpc.CallOption) (*DatabaseGetAllResponse, error)
	LoadToDisk(ctx context.Context, in *NoRequestBody, opts ...grpc.CallOption) (*LoadToDiskResponse, error)
	CreateNewCollection(ctx context.Context, in *CollectionCreateRequest, opts ...grpc.CallOption) (*CollectionCreateResponse, error)
	DeleteCollections(ctx context.Context, in *CollectionDeleteRequest, opts ...grpc.CallOption) (*CollectionDeleteResponse, error)
	GetAllCollections(ctx context.Context, in *CollectionGetAllRequest, opts ...grpc.CallOption) (*CollectionGetAllResponse, error)
	GetCollectionStats(ctx context.Context, in *CollectionStatsRequest, opts ...grpc.CallOption) (*CollectionStatsResponse, error)
	CreateDocument(ctx context.Context, in *DocumentCreateRequest, opts ...grpc.CallOption) (*DocumentCreateResponse, error)
	ReadDocument(ctx context.Context, in *DocumentReadRequest, opts ...grpc.CallOption) (*DocumentReadResponse, error)
	FilterDocument(ctx context.Context, in *DocumentFilterRequest, opts ...grpc.CallOption) (*DocumentFilterResponse, error)
	UpdateDocument(ctx context.Context, in *DocumentUpdateRequest, opts ...grpc.CallOption) (*DocumentUpdateResponse, error)
	DeleteDocument(ctx context.Context, in *DocumentDeleteRequest, opts ...grpc.CallOption) (*DocumentDeleteResponse, error)
	GetAllDocuments(ctx context.Context, in *DocumentGetAllRequest, opts ...grpc.CallOption) (*DocumentGetAllResponse, error)
}

type gnoSQLServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGnoSQLServiceClient(cc grpc.ClientConnInterface) GnoSQLServiceClient {
	return &gnoSQLServiceClient{cc}
}

func (c *gnoSQLServiceClient) CreateNewDatabase(ctx context.Context, in *DatabaseCreateRequest, opts ...grpc.CallOption) (*DatabaseCreateResponse, error) {
	out := new(DatabaseCreateResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/CreateNewDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) DeleteDatabase(ctx context.Context, in *DatabaseDeleteRequest, opts ...grpc.CallOption) (*DatabaseDeleteResponse, error) {
	out := new(DatabaseDeleteResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/DeleteDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) GetAllDatabases(ctx context.Context, in *NoRequestBody, opts ...grpc.CallOption) (*DatabaseGetAllResponse, error) {
	out := new(DatabaseGetAllResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/GetAllDatabases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) LoadToDisk(ctx context.Context, in *NoRequestBody, opts ...grpc.CallOption) (*LoadToDiskResponse, error) {
	out := new(LoadToDiskResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/LoadToDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) CreateNewCollection(ctx context.Context, in *CollectionCreateRequest, opts ...grpc.CallOption) (*CollectionCreateResponse, error) {
	out := new(CollectionCreateResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/CreateNewCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) DeleteCollections(ctx context.Context, in *CollectionDeleteRequest, opts ...grpc.CallOption) (*CollectionDeleteResponse, error) {
	out := new(CollectionDeleteResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/DeleteCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) GetAllCollections(ctx context.Context, in *CollectionGetAllRequest, opts ...grpc.CallOption) (*CollectionGetAllResponse, error) {
	out := new(CollectionGetAllResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/GetAllCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) GetCollectionStats(ctx context.Context, in *CollectionStatsRequest, opts ...grpc.CallOption) (*CollectionStatsResponse, error) {
	out := new(CollectionStatsResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/GetCollectionStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) CreateDocument(ctx context.Context, in *DocumentCreateRequest, opts ...grpc.CallOption) (*DocumentCreateResponse, error) {
	out := new(DocumentCreateResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/CreateDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) ReadDocument(ctx context.Context, in *DocumentReadRequest, opts ...grpc.CallOption) (*DocumentReadResponse, error) {
	out := new(DocumentReadResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/ReadDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) FilterDocument(ctx context.Context, in *DocumentFilterRequest, opts ...grpc.CallOption) (*DocumentFilterResponse, error) {
	out := new(DocumentFilterResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/FilterDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) UpdateDocument(ctx context.Context, in *DocumentUpdateRequest, opts ...grpc.CallOption) (*DocumentUpdateResponse, error) {
	out := new(DocumentUpdateResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/UpdateDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) DeleteDocument(ctx context.Context, in *DocumentDeleteRequest, opts ...grpc.CallOption) (*DocumentDeleteResponse, error) {
	out := new(DocumentDeleteResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/DeleteDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnoSQLServiceClient) GetAllDocuments(ctx context.Context, in *DocumentGetAllRequest, opts ...grpc.CallOption) (*DocumentGetAllResponse, error) {
	out := new(DocumentGetAllResponse)
	err := c.cc.Invoke(ctx, "/proto.GnoSQLService/GetAllDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GnoSQLServiceServer is the server API for GnoSQLService service.
// All implementations must embed UnimplementedGnoSQLServiceServer
// for forward compatibility
type GnoSQLServiceServer interface {
	CreateNewDatabase(context.Context, *DatabaseCreateRequest) (*DatabaseCreateResponse, error)
	DeleteDatabase(context.Context, *DatabaseDeleteRequest) (*DatabaseDeleteResponse, error)
	GetAllDatabases(context.Context, *NoRequestBody) (*DatabaseGetAllResponse, error)
	LoadToDisk(context.Context, *NoRequestBody) (*LoadToDiskResponse, error)
	CreateNewCollection(context.Context, *CollectionCreateRequest) (*CollectionCreateResponse, error)
	DeleteCollections(context.Context, *CollectionDeleteRequest) (*CollectionDeleteResponse, error)
	GetAllCollections(context.Context, *CollectionGetAllRequest) (*CollectionGetAllResponse, error)
	GetCollectionStats(context.Context, *CollectionStatsRequest) (*CollectionStatsResponse, error)
	CreateDocument(context.Context, *DocumentCreateRequest) (*DocumentCreateResponse, error)
	ReadDocument(context.Context, *DocumentReadRequest) (*DocumentReadResponse, error)
	FilterDocument(context.Context, *DocumentFilterRequest) (*DocumentFilterResponse, error)
	UpdateDocument(context.Context, *DocumentUpdateRequest) (*DocumentUpdateResponse, error)
	DeleteDocument(context.Context, *DocumentDeleteRequest) (*DocumentDeleteResponse, error)
	GetAllDocuments(context.Context, *DocumentGetAllRequest) (*DocumentGetAllResponse, error)
	mustEmbedUnimplementedGnoSQLServiceServer()
}

// UnimplementedGnoSQLServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGnoSQLServiceServer struct {
}

func (UnimplementedGnoSQLServiceServer) CreateNewDatabase(context.Context, *DatabaseCreateRequest) (*DatabaseCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewDatabase not implemented")
}
func (UnimplementedGnoSQLServiceServer) DeleteDatabase(context.Context, *DatabaseDeleteRequest) (*DatabaseDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDatabase not implemented")
}
func (UnimplementedGnoSQLServiceServer) GetAllDatabases(context.Context, *NoRequestBody) (*DatabaseGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDatabases not implemented")
}
func (UnimplementedGnoSQLServiceServer) LoadToDisk(context.Context, *NoRequestBody) (*LoadToDiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadToDisk not implemented")
}
func (UnimplementedGnoSQLServiceServer) CreateNewCollection(context.Context, *CollectionCreateRequest) (*CollectionCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewCollection not implemented")
}
func (UnimplementedGnoSQLServiceServer) DeleteCollections(context.Context, *CollectionDeleteRequest) (*CollectionDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollections not implemented")
}
func (UnimplementedGnoSQLServiceServer) GetAllCollections(context.Context, *CollectionGetAllRequest) (*CollectionGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCollections not implemented")
}
func (UnimplementedGnoSQLServiceServer) GetCollectionStats(context.Context, *CollectionStatsRequest) (*CollectionStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectionStats not implemented")
}
func (UnimplementedGnoSQLServiceServer) CreateDocument(context.Context, *DocumentCreateRequest) (*DocumentCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocument not implemented")
}
func (UnimplementedGnoSQLServiceServer) ReadDocument(context.Context, *DocumentReadRequest) (*DocumentReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadDocument not implemented")
}
func (UnimplementedGnoSQLServiceServer) FilterDocument(context.Context, *DocumentFilterRequest) (*DocumentFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterDocument not implemented")
}
func (UnimplementedGnoSQLServiceServer) UpdateDocument(context.Context, *DocumentUpdateRequest) (*DocumentUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDocument not implemented")
}
func (UnimplementedGnoSQLServiceServer) DeleteDocument(context.Context, *DocumentDeleteRequest) (*DocumentDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDocument not implemented")
}
func (UnimplementedGnoSQLServiceServer) GetAllDocuments(context.Context, *DocumentGetAllRequest) (*DocumentGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDocuments not implemented")
}
func (UnimplementedGnoSQLServiceServer) mustEmbedUnimplementedGnoSQLServiceServer() {}

// UnsafeGnoSQLServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GnoSQLServiceServer will
// result in compilation errors.
type UnsafeGnoSQLServiceServer interface {
	mustEmbedUnimplementedGnoSQLServiceServer()
}

func RegisterGnoSQLServiceServer(s grpc.ServiceRegistrar, srv GnoSQLServiceServer) {
	s.RegisterService(&GnoSQLService_ServiceDesc, srv)
}

func _GnoSQLService_CreateNewDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).CreateNewDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/CreateNewDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).CreateNewDatabase(ctx, req.(*DatabaseCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_DeleteDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatabaseDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).DeleteDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/DeleteDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).DeleteDatabase(ctx, req.(*DatabaseDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_GetAllDatabases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).GetAllDatabases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/GetAllDatabases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).GetAllDatabases(ctx, req.(*NoRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_LoadToDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoRequestBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).LoadToDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/LoadToDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).LoadToDisk(ctx, req.(*NoRequestBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_CreateNewCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).CreateNewCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/CreateNewCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).CreateNewCollection(ctx, req.(*CollectionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_DeleteCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).DeleteCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/DeleteCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).DeleteCollections(ctx, req.(*CollectionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_GetAllCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).GetAllCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/GetAllCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).GetAllCollections(ctx, req.(*CollectionGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_GetCollectionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).GetCollectionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/GetCollectionStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).GetCollectionStats(ctx, req.(*CollectionStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_CreateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).CreateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/CreateDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).CreateDocument(ctx, req.(*DocumentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_ReadDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).ReadDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/ReadDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).ReadDocument(ctx, req.(*DocumentReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_FilterDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).FilterDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/FilterDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).FilterDocument(ctx, req.(*DocumentFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_UpdateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).UpdateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/UpdateDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).UpdateDocument(ctx, req.(*DocumentUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_DeleteDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).DeleteDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/DeleteDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).DeleteDocument(ctx, req.(*DocumentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnoSQLService_GetAllDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnoSQLServiceServer).GetAllDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GnoSQLService/GetAllDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnoSQLServiceServer).GetAllDocuments(ctx, req.(*DocumentGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GnoSQLService_ServiceDesc is the grpc.ServiceDesc for GnoSQLService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GnoSQLService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GnoSQLService",
	HandlerType: (*GnoSQLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewDatabase",
			Handler:    _GnoSQLService_CreateNewDatabase_Handler,
		},
		{
			MethodName: "DeleteDatabase",
			Handler:    _GnoSQLService_DeleteDatabase_Handler,
		},
		{
			MethodName: "GetAllDatabases",
			Handler:    _GnoSQLService_GetAllDatabases_Handler,
		},
		{
			MethodName: "LoadToDisk",
			Handler:    _GnoSQLService_LoadToDisk_Handler,
		},
		{
			MethodName: "CreateNewCollection",
			Handler:    _GnoSQLService_CreateNewCollection_Handler,
		},
		{
			MethodName: "DeleteCollections",
			Handler:    _GnoSQLService_DeleteCollections_Handler,
		},
		{
			MethodName: "GetAllCollections",
			Handler:    _GnoSQLService_GetAllCollections_Handler,
		},
		{
			MethodName: "GetCollectionStats",
			Handler:    _GnoSQLService_GetCollectionStats_Handler,
		},
		{
			MethodName: "CreateDocument",
			Handler:    _GnoSQLService_CreateDocument_Handler,
		},
		{
			MethodName: "ReadDocument",
			Handler:    _GnoSQLService_ReadDocument_Handler,
		},
		{
			MethodName: "FilterDocument",
			Handler:    _GnoSQLService_FilterDocument_Handler,
		},
		{
			MethodName: "UpdateDocument",
			Handler:    _GnoSQLService_UpdateDocument_Handler,
		},
		{
			MethodName: "DeleteDocument",
			Handler:    _GnoSQLService_DeleteDocument_Handler,
		},
		{
			MethodName: "GetAllDocuments",
			Handler:    _GnoSQLService_GetAllDocuments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gnosql.proto",
}
