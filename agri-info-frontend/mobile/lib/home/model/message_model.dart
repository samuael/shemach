class Mes {}

class Message extends Mes {
  int id;
  String lang;
  String data;
  int createdAt;
  int createdBy;
  List<int> targets;
  bool seen = false;
  Message({
    required this.id,
    required this.lang,
    required this.data,
    required this.createdAt,
    required this.createdBy,
    required this.targets,
  });

  factory Message.fromJson(Map<String, dynamic> json) {
    return Message(
      id: json["id"] ?? 0,
      lang: json["lang"] ?? 'amh',
      data: json["data"] ?? '',
      createdAt: json["created_at"] ?? 0,
      createdBy: json["created_by"] ?? 0,
      targets: (json["targets"] ?? []).map<int>((e){return e as int;}).toList(),
    );
  }
}

class ProductUpdate extends Mes {
  int productID;
  double cost;

  ProductUpdate({required this.productID, required this.cost});

  factory ProductUpdate.fromJson(Map<String, dynamic> json) {
    return ProductUpdate(
      productID: json["id"] ?? -1,
      cost: json["cost"] ?? -1.0,
    );
  }
}

class MessageResponse {
  Mes body;
  int type;

  MessageResponse({required this.type, required this.body});

  factory MessageResponse.fromJson(Map<String, dynamic> json) {
    int dtype = (json["type"] ?? 0 as int);
    return MessageResponse(
      type: dtype,
      body: dtype == 0
          ? Message(
              id: 0,
              createdBy: 0,
              lang: "amh",
              data: "",
              createdAt: 0,
              targets: [-1])
          : (dtype == 1
              ? Message.fromJson(json["body"] as Map<String, dynamic>)
              : ProductUpdate.fromJson(json["body"] as Map<String, dynamic>)),
    );
  }
}
