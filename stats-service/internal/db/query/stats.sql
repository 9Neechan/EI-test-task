syntax = "proto3";

package stats;

option go_package = "proto/stats";

service StatsService {
    // Фиксируем новый вызов сервиса (POST /call)
    rpc PostCall(PostCallRequest) returns (PostCallResponse);

    // Получаем статистику вызовов (GET /calls)
    rpc GetCalls(GetCallsRequest) returns (GetCallsResponse);
}

// Запрос на создание/обновление вызова
message PostCallRequest {
    int64 user_id = 1;
    int64 service_id = 2;
}

message PostCallResponse {
    int64 user_id = 1;
    int64 service_id = 2;
    int64 count = 3;
    string created_at = 4;
}

// Запрос на получение статистики с фильтрами и пагинацией
message GetCallsRequest {
    optional int64 user_id = 1;    // Фильтр по user_id
    optional int64 service_id = 2; // Фильтр по service_id
    int32 limit = 3;               // Количество записей на страницу
    int32 offset = 4;              // Смещение для пагинации
}

// Ответ со списком вызовов
message GetCallsResponse {
    repeated CallEntry calls = 1;
}

// Структура одной записи статистики
message CallEntry {
    int64 user_id = 1;
    string user_name = 2;
    int64 service_id = 3;
    string service_name = 4;
    int64 count = 5;
}
