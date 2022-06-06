import "../../libs.dart";

class MessagesResponse {
  List<Message> messages;
  int statusCode;
  String msg;

  MessagesResponse({
    required this.statusCode,
    required this.messages,
    required this.msg,
  });

  factory MessagesResponse.fromJson(Map<String, dynamic> json) {
    return MessagesResponse(
      msg: json["msg"] ?? '',
      statusCode: json["status_code"] ?? 999,
      messages: (((json["messages"] ?? []).map<Map<String, dynamic>>((e) {
                return e as Map<String, dynamic>;
              }).toList() as List<Map<String, dynamic>>?) ??
              [])
          .map((e) {
        return Message.fromJson(e);
      }).toList(),
    );
  }
}

/*
 {
    "messages": [
        {
            "id": 24,
            "targets": [
                -1
            ],
            "lang": "all",
            "data": "Welcome to Agri-net systems",
            "created_by": 3,
            "created_at": 1650790664
        },
        {
            "id": 23,
            "targets": [
                -1
            ],
            "lang": "all",
            "data": "Welcome to Agri-net systems",
            "created_by": 3,
            "created_at": 1650790567
        },
        {
            "id": 22,
            "targets": [
                -1
            ],
            "lang": "all",
            "data": "Welcome to Agri-net systems",
            "created_by": 3,
            "created_at": 1650790259
        },
        {
            "id": 21,
            "targets": [
                -1
            ],
            "lang": "all",
            "data": "Welcome to Agri-net systems",
            "created_by": 2,
            "created_at": 1650290362
        },
        {
            "id": 20,
            "targets": [
                -1
            ],
            "lang": "all",
            "data": "Welcome to Agri-net systems",
            "created_by": 2,
            "created_at": 1650290360
        }
    ],
    "status_code": 200
}*/