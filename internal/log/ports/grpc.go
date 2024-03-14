package ports

import (
	"context"
	"distributed_log/internal/common/api/protobuf"
	"distributed_log/internal/common/api/protobuf/v4"
	server "distributed_log/internal/common/server/grpc"
	"distributed_log/internal/event"
	"log/slog"
)

type GrpcServer struct {
	Server *server.GrpcServer
	index_proto.UnimplementedLogServer
	handler *event.Handler
}

func NewGrpcServer(handler *event.Handler, port int) (*GrpcServer, error) {
	ser := GrpcServer{}

	newServer, err := server.NewServer(port, server.Register(&ser))
	if err != nil {
		return nil, err
	}

	ser.Server = newServer
	ser.handler = handler

	return &ser, nil
}

func (s *GrpcServer) Serve() {
	s.Server.Serve()
}
func (s *GrpcServer) GracefulStop() {
	s.Server.GracefulStop()
}

func (s *GrpcServer) Produce(ctx context.Context, req *index_proto.ProduceRequest) (*index_proto.ProduceResponse, error) {

	slog.Info("(s *GrpcServer) Produce(ctx context.Context, req *v4.ProduceRequest) (*v4.ProduceResponse, error)")
	off, err := s.handler.EventBus.Publish(&event.Event{Type: string(event.EventInsert), Data: req})
	if err != nil {
		return nil, err
	}

	//offset, err := s.CommitLog.Append(req.Record)
	//if err != nil {
	//	return nil, err
	//}
	return &index_proto.ProduceResponse{Offset: off.Offset}, nil
	//return &v4.ProduceResponse{Offset: 777}, nil
}

func (s *GrpcServer) Consume(ctx context.Context, req *index_proto.ConsumeRequest) (*index_proto.ConsumeResponse, error) {
	off, err := s.handler.EventBus.Publish(&event.Event{Type: string(event.EventGet), Data: req})
	if err != nil {
		return nil, err
	}
	//slog.Info("Consume(ctx context.Context, req *v4.ConsumeRequest) (*v4.ConsumeResponse, error) ")
	//record, err := s.CommitLog.Read(req.Offset)
	//if err != nil {
	//	return nil, err
	//}

	//return &v4.ConsumeResponse{Record: record}, nil
	//return nil, nil
	return &index_proto.ConsumeResponse{Record: off}, nil
}

func (s *GrpcServer) ProduceStream(stream index_proto.Log_ProduceStreamServer) error {
	for {
		req, err := stream.Recv()
		//slog.Info("ProduceStream(stream v4.Log_ProduceStreamServer) error ")
		if err != nil {
			return err
		}
		res, err := s.Produce(stream.Context(), req)
		if err != nil {
			return err
		}
		if err = stream.Send(res); err != nil {
			return err
		}
	}
}

func (s *GrpcServer) ConsumeStream(req *index_proto.ConsumeRequest, stream index_proto.Log_ConsumeStreamServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			res, err := s.Consume(stream.Context(), req)
			//slog.Info("ConsumeStream(req *v4.ConsumeRequest, stream v4.Log_ConsumeStreamServer) error")
			switch err.(type) {
			case nil:
				//slog.Info("consume stream case nil")
				//continue
			case protobuf.ErrOffsetOutOfRange:
				continue
			default:
				return err
			}
			if err = stream.Send(res); err != nil {
				slog.Info("if err = stream.Send(res); err != nil {", res)
				return err
			}
			req.Offset++
		}
	}
}
