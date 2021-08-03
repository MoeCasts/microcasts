package grpc

import (
	"context"
	"errors"

	grpc "github.com/go-kit/kit/transport/grpc"
	endpoint "github.com/moecasts/microcasts/novels/pkg/endpoint"
	pb "github.com/moecasts/microcasts/novels/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeBrowseHandler creates the handler logic
func makeBrowseHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.BrowseEndpoint, decodeBrowseRequest, encodeBrowseResponse, options...)
}

// decodeBrowseResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Browse request.
func decodeBrowseRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.BrowseRequest)
	return endpoint.BrowseRequest{
		Request: req,
	}, nil
}

// encodeBrowseResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeBrowseResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.BrowseResponse)
	return &pb.BrowseReply{
		Data: res.Rs.([]*pb.Novel),
	}, nil
}
func (g *grpcServer) Browse(ctx context1.Context, req *pb.BrowseRequest) (*pb.BrowseReply, error) {
	_, rep, err := g.browse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.BrowseReply), nil
}

// makeReadHandler creates the handler logic
func makeReadHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.ReadEndpoint, decodeReadRequest, encodeReadResponse, options...)
}

// decodeReadResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Read request.
// TODO implement the decoder
func decodeReadRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeReadResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeReadResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}
func (g *grpcServer) Read(ctx context1.Context, req *pb.ReadRequest) (*pb.ReadReply, error) {
	_, rep, err := g.read.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ReadReply), nil
}

// makeAddHandler creates the handler logic
func makeAddHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddEndpoint, decodeAddRequest, encodeAddResponse, options...)
}

// decodeAddResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Add request.
// TODO implement the decoder
func decodeAddRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeAddResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAddResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}
func (g *grpcServer) Add(ctx context1.Context, req *pb.AddRequest) (*pb.AddReply, error) {
	_, rep, err := g.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddReply), nil
}

// makeEditHandler creates the handler logic
func makeEditHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.EditEndpoint, decodeEditRequest, encodeEditResponse, options...)
}

// decodeEditResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Edit request.
// TODO implement the decoder
func decodeEditRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeEditResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeEditResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}
func (g *grpcServer) Edit(ctx context1.Context, req *pb.EditRequest) (*pb.EditReply, error) {
	_, rep, err := g.edit.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.EditReply), nil
}

// makeTrashHandler creates the handler logic
func makeTrashHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.TrashEndpoint, decodeTrashRequest, encodeTrashResponse, options...)
}

// decodeTrashResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Trash request.
// TODO implement the decoder
func decodeTrashRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeTrashResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeTrashResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}
func (g *grpcServer) Trash(ctx context1.Context, req *pb.TrashRequest) (*pb.TrashReply, error) {
	_, rep, err := g.trash.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TrashReply), nil
}

// makeRestoreHandler creates the handler logic
func makeRestoreHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RestoreEndpoint, decodeRestoreRequest, encodeRestoreResponse, options...)
}

// decodeRestoreResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Restore request.
// TODO implement the decoder
func decodeRestoreRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeRestoreResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeRestoreResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}
func (g *grpcServer) Restore(ctx context1.Context, req *pb.RestoreRequest) (*pb.RestoreReply, error) {
	_, rep, err := g.restore.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RestoreReply), nil
}

// makeDestroyHandler creates the handler logic
func makeDestroyHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DestroyEndpoint, decodeDestroyRequest, encodeDestroyResponse, options...)
}

// decodeDestroyResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Destroy request.
// TODO implement the decoder
func decodeDestroyRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeDestroyResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDestroyResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}
func (g *grpcServer) Destroy(ctx context1.Context, req *pb.DestroyRequest) (*pb.DestroyReply, error) {
	_, rep, err := g.destroy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DestroyReply), nil
}

// makeBatchAddHandler creates the handler logic
func makeBatchAddHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.BatchAddEndpoint, decodeBatchAddRequest, encodeBatchAddResponse, options...)
}

// decodeBatchAddResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain BatchAdd request.
// TODO implement the decoder
func decodeBatchAddRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeBatchAddResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeBatchAddResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}
func (g *grpcServer) BatchAdd(ctx context1.Context, req *pb.BatchAddRequest) (*pb.BatchAddReply, error) {
	_, rep, err := g.batchAdd.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.BatchAddReply), nil
}

// makeBatchEditHandler creates the handler logic
func makeBatchEditHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.BatchEditEndpoint, decodeBatchEditRequest, encodeBatchEditResponse, options...)
}

// decodeBatchEditResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain BatchEdit request.
// TODO implement the decoder
func decodeBatchEditRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeBatchEditResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeBatchEditResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}
func (g *grpcServer) BatchEdit(ctx context1.Context, req *pb.BatchEditRequest) (*pb.BatchEditReply, error) {
	_, rep, err := g.batchEdit.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.BatchEditReply), nil
}

// makeBatchTrashHandler creates the handler logic
func makeBatchTrashHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.BatchTrashEndpoint, decodeBatchTrashRequest, encodeBatchTrashResponse, options...)
}

// decodeBatchTrashResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain BatchTrash request.
// TODO implement the decoder
func decodeBatchTrashRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeBatchTrashResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeBatchTrashResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}
func (g *grpcServer) BatchTrash(ctx context1.Context, req *pb.BatchTrashRequest) (*pb.BatchTrashReply, error) {
	_, rep, err := g.batchTrash.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.BatchTrashReply), nil
}

// makeBatchRestoreHandler creates the handler logic
func makeBatchRestoreHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.BatchRestoreEndpoint, decodeBatchRestoreRequest, encodeBatchRestoreResponse, options...)
}

// decodeBatchRestoreResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain BatchRestore request.
// TODO implement the decoder
func decodeBatchRestoreRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeBatchRestoreResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeBatchRestoreResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}
func (g *grpcServer) BatchRestore(ctx context1.Context, req *pb.BatchRestoreRequest) (*pb.BatchRestoreReply, error) {
	_, rep, err := g.batchRestore.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.BatchRestoreReply), nil
}

// makeBatchDestroyHandler creates the handler logic
func makeBatchDestroyHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.BatchDestroyEndpoint, decodeBatchDestroyRequest, encodeBatchDestroyResponse, options...)
}

// decodeBatchDestroyResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain BatchDestroy request.
// TODO implement the decoder
func decodeBatchDestroyRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Decoder is not impelemented")
}

// encodeBatchDestroyResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeBatchDestroyResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Novels' Encoder is not impelemented")
}
func (g *grpcServer) BatchDestroy(ctx context1.Context, req *pb.BatchDestroyRequest) (*pb.BatchDestroyReply, error) {
	_, rep, err := g.batchDestroy.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.BatchDestroyReply), nil
}
