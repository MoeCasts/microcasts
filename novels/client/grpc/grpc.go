package grpc

import (
	"context"
	"errors"

	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	endpoint1 "github.com/moecasts/microcasts/novels/pkg/endpoint"
	pb "github.com/moecasts/microcasts/novels/pkg/grpc/pb"
	service "github.com/moecasts/microcasts/novels/pkg/service"
	grpc "google.golang.org/grpc"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.NovelsService, error) {
	var browseEndpoint endpoint.Endpoint
	{
		browseEndpoint = grpc1.NewClient(conn, "pb.Novels", "Browse", encodeBrowseRequest, decodeBrowseResponse, pb.BrowseReply{}, options["Browse"]...).Endpoint()
	}

	var readEndpoint endpoint.Endpoint
	{
		readEndpoint = grpc1.NewClient(conn, "pb.Novels", "Read", encodeReadRequest, decodeReadResponse, pb.ReadReply{}, options["Read"]...).Endpoint()
	}

	var addEndpoint endpoint.Endpoint
	{
		addEndpoint = grpc1.NewClient(conn, "pb.Novels", "Add", encodeAddRequest, decodeAddResponse, pb.AddReply{}, options["Add"]...).Endpoint()
	}

	var editEndpoint endpoint.Endpoint
	{
		editEndpoint = grpc1.NewClient(conn, "pb.Novels", "Edit", encodeEditRequest, decodeEditResponse, pb.EditReply{}, options["Edit"]...).Endpoint()
	}

	var trashEndpoint endpoint.Endpoint
	{
		trashEndpoint = grpc1.NewClient(conn, "pb.Novels", "Trash", encodeTrashRequest, decodeTrashResponse, pb.TrashReply{}, options["Trash"]...).Endpoint()
	}

	var restoreEndpoint endpoint.Endpoint
	{
		restoreEndpoint = grpc1.NewClient(conn, "pb.Novels", "Restore", encodeRestoreRequest, decodeRestoreResponse, pb.RestoreReply{}, options["Restore"]...).Endpoint()
	}

	var destroyEndpoint endpoint.Endpoint
	{
		destroyEndpoint = grpc1.NewClient(conn, "pb.Novels", "Destroy", encodeDestroyRequest, decodeDestroyResponse, pb.DestroyReply{}, options["Destroy"]...).Endpoint()
	}

	var batchAddEndpoint endpoint.Endpoint
	{
		batchAddEndpoint = grpc1.NewClient(conn, "pb.Novels", "BatchAdd", encodeBatchAddRequest, decodeBatchAddResponse, pb.BatchAddReply{}, options["BatchAdd"]...).Endpoint()
	}

	var batchEditEndpoint endpoint.Endpoint
	{
		batchEditEndpoint = grpc1.NewClient(conn, "pb.Novels", "BatchEdit", encodeBatchEditRequest, decodeBatchEditResponse, pb.BatchEditReply{}, options["BatchEdit"]...).Endpoint()
	}

	var batchTrashEndpoint endpoint.Endpoint
	{
		batchTrashEndpoint = grpc1.NewClient(conn, "pb.Novels", "BatchTrash", encodeBatchTrashRequest, decodeBatchTrashResponse, pb.BatchTrashReply{}, options["BatchTrash"]...).Endpoint()
	}

	var batchRestoreEndpoint endpoint.Endpoint
	{
		batchRestoreEndpoint = grpc1.NewClient(conn, "pb.Novels", "BatchRestore", encodeBatchRestoreRequest, decodeBatchRestoreResponse, pb.BatchRestoreReply{}, options["BatchRestore"]...).Endpoint()
	}

	var batchDestroyEndpoint endpoint.Endpoint
	{
		batchDestroyEndpoint = grpc1.NewClient(conn, "pb.Novels", "BatchDestroy", encodeBatchDestroyRequest, decodeBatchDestroyResponse, pb.BatchDestroyReply{}, options["BatchDestroy"]...).Endpoint()
	}

	return endpoint1.Endpoints{
		AddEndpoint:          addEndpoint,
		BatchAddEndpoint:     batchAddEndpoint,
		BatchDestroyEndpoint: batchDestroyEndpoint,
		BatchEditEndpoint:    batchEditEndpoint,
		BatchRestoreEndpoint: batchRestoreEndpoint,
		BatchTrashEndpoint:   batchTrashEndpoint,
		BrowseEndpoint:       browseEndpoint,
		DestroyEndpoint:      destroyEndpoint,
		EditEndpoint:         editEndpoint,
		ReadEndpoint:         readEndpoint,
		RestoreEndpoint:      restoreEndpoint,
		TrashEndpoint:        trashEndpoint,
	}, nil
}

// encodeBrowseRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Browse request to a gRPC request.
func encodeBrowseRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint1.BrowseRequest)
	return req.Request.(*pb.BrowseRequest), nil
}

// decodeBrowseResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeBrowseResponse(_ context.Context, reply interface{}) (interface{}, error) {
	rep := reply.(*pb.BrowseReply)
	return endpoint1.BrowseResponse{Rs: rep.Data}, nil
}

// encodeReadRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Read request to a gRPC request.
func encodeReadRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}

// decodeReadResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeReadResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeAddRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Add request to a gRPC request.
func encodeAddRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}

// decodeAddResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeAddResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeEditRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Edit request to a gRPC request.
func encodeEditRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}

// decodeEditResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeEditResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeTrashRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Trash request to a gRPC request.
func encodeTrashRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}

// decodeTrashResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeTrashResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeRestoreRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Restore request to a gRPC request.
func encodeRestoreRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}

// decodeRestoreResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeRestoreResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeDestroyRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain Destroy request to a gRPC request.
func encodeDestroyRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}

// decodeDestroyResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeDestroyResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeBatchAddRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain BatchAdd request to a gRPC request.
func encodeBatchAddRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}

// decodeBatchAddResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeBatchAddResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeBatchEditRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain BatchEdit request to a gRPC request.
func encodeBatchEditRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}

// decodeBatchEditResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeBatchEditResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeBatchTrashRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain BatchTrash request to a gRPC request.
func encodeBatchTrashRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}

// decodeBatchTrashResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeBatchTrashResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeBatchRestoreRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain BatchRestore request to a gRPC request.
func encodeBatchRestoreRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}

// decodeBatchRestoreResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeBatchRestoreResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeBatchDestroyRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain BatchDestroy request to a gRPC request.
func encodeBatchDestroyRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}

// decodeBatchDestroyResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeBatchDestroyResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}
