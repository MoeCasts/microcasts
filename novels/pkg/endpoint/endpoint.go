package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/moecasts/microcasts/novels/pkg/service"
)

// BrowseRequest collects the request parameters for the Browse method.
type BrowseRequest struct {
	Request interface{} `json:"request"`
}

// BrowseResponse collects the response parameters for the Browse method.
type BrowseResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeBrowseEndpoint returns an endpoint that invokes Browse on the service.
func MakeBrowseEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BrowseRequest)
		rs, err := s.Browse(ctx, req.Request)
		return BrowseResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r BrowseResponse) Failed() error {
	return r.Err
}

// ReadRequest collects the request parameters for the Read method.
type ReadRequest struct {
	Pk      interface{} `json:"pk"`
	Request interface{} `json:"request"`
}

// ReadResponse collects the response parameters for the Read method.
type ReadResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeReadEndpoint returns an endpoint that invokes Read on the service.
func MakeReadEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReadRequest)
		rs, err := s.Read(ctx, req.Pk, req.Request)
		return ReadResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ReadResponse) Failed() error {
	return r.Err
}

// AddRequest collects the request parameters for the Add method.
type AddRequest struct {
	Request interface{} `json:"request"`
}

// AddResponse collects the response parameters for the Add method.
type AddResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeAddEndpoint returns an endpoint that invokes Add on the service.
func MakeAddEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		rs, err := s.Add(ctx, req.Request)
		return AddResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r AddResponse) Failed() error {
	return r.Err
}

// EditRequest collects the request parameters for the Edit method.
type EditRequest struct {
	Pk      interface{} `json:"pk"`
	Request interface{} `json:"request"`
}

// EditResponse collects the response parameters for the Edit method.
type EditResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeEditEndpoint returns an endpoint that invokes Edit on the service.
func MakeEditEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EditRequest)
		rs, err := s.Edit(ctx, req.Pk, req.Request)
		return EditResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r EditResponse) Failed() error {
	return r.Err
}

// TrashRequest collects the request parameters for the Trash method.
type TrashRequest struct {
	Pk interface{} `json:"pk"`
}

// TrashResponse collects the response parameters for the Trash method.
type TrashResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeTrashEndpoint returns an endpoint that invokes Trash on the service.
func MakeTrashEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TrashRequest)
		rs, err := s.Trash(ctx, req.Pk)
		return TrashResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r TrashResponse) Failed() error {
	return r.Err
}

// RestoreRequest collects the request parameters for the Restore method.
type RestoreRequest struct {
	Pk interface{} `json:"pk"`
}

// RestoreResponse collects the response parameters for the Restore method.
type RestoreResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeRestoreEndpoint returns an endpoint that invokes Restore on the service.
func MakeRestoreEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RestoreRequest)
		rs, err := s.Restore(ctx, req.Pk)
		return RestoreResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r RestoreResponse) Failed() error {
	return r.Err
}

// DestroyRequest collects the request parameters for the Destroy method.
type DestroyRequest struct {
	Pk interface{} `json:"pk"`
}

// DestroyResponse collects the response parameters for the Destroy method.
type DestroyResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeDestroyEndpoint returns an endpoint that invokes Destroy on the service.
func MakeDestroyEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DestroyRequest)
		rs, err := s.Destroy(ctx, req.Pk)
		return DestroyResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DestroyResponse) Failed() error {
	return r.Err
}

// BatchAddRequest collects the request parameters for the BatchAdd method.
type BatchAddRequest struct {
	Request interface{} `json:"request"`
}

// BatchAddResponse collects the response parameters for the BatchAdd method.
type BatchAddResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeBatchAddEndpoint returns an endpoint that invokes BatchAdd on the service.
func MakeBatchAddEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BatchAddRequest)
		rs, err := s.BatchAdd(ctx, req.Request)
		return BatchAddResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r BatchAddResponse) Failed() error {
	return r.Err
}

// BatchEditRequest collects the request parameters for the BatchEdit method.
type BatchEditRequest struct {
	Request interface{} `json:"request"`
}

// BatchEditResponse collects the response parameters for the BatchEdit method.
type BatchEditResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeBatchEditEndpoint returns an endpoint that invokes BatchEdit on the service.
func MakeBatchEditEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BatchEditRequest)
		rs, err := s.BatchEdit(ctx, req.Request)
		return BatchEditResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r BatchEditResponse) Failed() error {
	return r.Err
}

// BatchTrashRequest collects the request parameters for the BatchTrash method.
type BatchTrashRequest struct {
	Pks interface{} `json:"pks"`
}

// BatchTrashResponse collects the response parameters for the BatchTrash method.
type BatchTrashResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeBatchTrashEndpoint returns an endpoint that invokes BatchTrash on the service.
func MakeBatchTrashEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BatchTrashRequest)
		rs, err := s.BatchTrash(ctx, req.Pks)
		return BatchTrashResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r BatchTrashResponse) Failed() error {
	return r.Err
}

// BatchRestoreRequest collects the request parameters for the BatchRestore method.
type BatchRestoreRequest struct {
	Pks interface{} `json:"pks"`
}

// BatchRestoreResponse collects the response parameters for the BatchRestore method.
type BatchRestoreResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeBatchRestoreEndpoint returns an endpoint that invokes BatchRestore on the service.
func MakeBatchRestoreEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BatchRestoreRequest)
		rs, err := s.BatchRestore(ctx, req.Pks)
		return BatchRestoreResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r BatchRestoreResponse) Failed() error {
	return r.Err
}

// BatchDestroyRequest collects the request parameters for the BatchDestroy method.
type BatchDestroyRequest struct {
	Pks interface{} `json:"pks"`
}

// BatchDestroyResponse collects the response parameters for the BatchDestroy method.
type BatchDestroyResponse struct {
	Rs  interface{} `json:"rs"`
	Err error       `json:"err"`
}

// MakeBatchDestroyEndpoint returns an endpoint that invokes BatchDestroy on the service.
func MakeBatchDestroyEndpoint(s service.NovelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(BatchDestroyRequest)
		rs, err := s.BatchDestroy(ctx, req.Pks)
		return BatchDestroyResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r BatchDestroyResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Browse implements Service. Primarily useful in a client.
func (e Endpoints) Browse(ctx context.Context, request interface{}) (rs interface{}, err error) {
	request0 := BrowseRequest{Request: request}
	response, err := e.BrowseEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response.(BrowseResponse).Rs, response.(BrowseResponse).Err
}

// Read implements Service. Primarily useful in a client.
func (e Endpoints) Read(ctx context.Context, pk interface{}, request interface{}) (rs interface{}, err error) {
	request0 := ReadRequest{
		Pk:      pk,
		Request: request,
	}
	response, err := e.ReadEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response.(ReadResponse).Rs, response.(ReadResponse).Err
}

// Add implements Service. Primarily useful in a client.
func (e Endpoints) Add(ctx context.Context, request interface{}) (rs interface{}, err error) {
	request0 := AddRequest{Request: request}
	response, err := e.AddEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response.(AddResponse).Rs, response.(AddResponse).Err
}

// Edit implements Service. Primarily useful in a client.
func (e Endpoints) Edit(ctx context.Context, pk interface{}, request interface{}) (rs interface{}, err error) {
	request0 := EditRequest{
		Pk:      pk,
		Request: request,
	}
	response, err := e.EditEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response.(EditResponse).Rs, response.(EditResponse).Err
}

// Trash implements Service. Primarily useful in a client.
func (e Endpoints) Trash(ctx context.Context, pk interface{}) (rs interface{}, err error) {
	request := TrashRequest{Pk: pk}
	response, err := e.TrashEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(TrashResponse).Rs, response.(TrashResponse).Err
}

// Restore implements Service. Primarily useful in a client.
func (e Endpoints) Restore(ctx context.Context, pk interface{}) (rs interface{}, err error) {
	request := RestoreRequest{Pk: pk}
	response, err := e.RestoreEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RestoreResponse).Rs, response.(RestoreResponse).Err
}

// Destroy implements Service. Primarily useful in a client.
func (e Endpoints) Destroy(ctx context.Context, pk interface{}) (rs interface{}, err error) {
	request := DestroyRequest{Pk: pk}
	response, err := e.DestroyEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DestroyResponse).Rs, response.(DestroyResponse).Err
}

// BatchAdd implements Service. Primarily useful in a client.
func (e Endpoints) BatchAdd(ctx context.Context, request interface{}) (rs interface{}, err error) {
	request0 := BatchAddRequest{Request: request}
	response, err := e.BatchAddEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response.(BatchAddResponse).Rs, response.(BatchAddResponse).Err
}

// BatchEdit implements Service. Primarily useful in a client.
func (e Endpoints) BatchEdit(ctx context.Context, request interface{}) (rs interface{}, err error) {
	request0 := BatchEditRequest{Request: request}
	response, err := e.BatchEditEndpoint(ctx, request0)
	if err != nil {
		return
	}
	return response.(BatchEditResponse).Rs, response.(BatchEditResponse).Err
}

// BatchTrash implements Service. Primarily useful in a client.
func (e Endpoints) BatchTrash(ctx context.Context, pks interface{}) (rs interface{}, err error) {
	request := BatchTrashRequest{Pks: pks}
	response, err := e.BatchTrashEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(BatchTrashResponse).Rs, response.(BatchTrashResponse).Err
}

// BatchRestore implements Service. Primarily useful in a client.
func (e Endpoints) BatchRestore(ctx context.Context, pks interface{}) (rs interface{}, err error) {
	request := BatchRestoreRequest{Pks: pks}
	response, err := e.BatchRestoreEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(BatchRestoreResponse).Rs, response.(BatchRestoreResponse).Err
}

// BatchDestroy implements Service. Primarily useful in a client.
func (e Endpoints) BatchDestroy(ctx context.Context, pks interface{}) (rs interface{}, err error) {
	request := BatchDestroyRequest{Pks: pks}
	response, err := e.BatchDestroyEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(BatchDestroyResponse).Rs, response.(BatchDestroyResponse).Err
}
