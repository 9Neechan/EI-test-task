package grpc

import (
    "context"
    "log"

    pb "github.com/9Neechan/EI-test-task/api/pb" // Подключаем сгенерированный gRPC-код
    "google.golang.org/grpc"
)

func NewStatsServiceClient(conn *grpc.ClientConn) pb.StatsServiceClient {
    return pb.NewStatsServiceClient(conn)
}

func CallStatsService(client pb.StatsServiceClient, userID, serviceID int64) {
    req := &pb.PostCallRequest{
        UserId:    userID,
        ServiceId: serviceID,
    }
    
    resp, err := client.PostCall(context.Background(), req)
    if err != nil {
        log.Fatalf("Ошибка вызова gRPC: %v", err)
    }
    log.Printf("Ответ сервера: %v", resp.Success)
}
