import '../../libs.dart';

class Admin {
  int id;
  String fullname;
  String email;
  bool superadmin;
  String imgurl;
  DateTime? createdAt;

  Admin({
    required this.id,
    required this.fullname,
    required this.email,
    required this.superadmin,
    required this.imgurl,
    this.createdAt,
  });

  factory Admin.fromJson(Map<String, dynamic> json) {
    return Admin(
      id: int.parse("${json['id']}"),
      fullname: json["fullname"],
      email: json["email"],
      superadmin: json["superadmin"],
      imgurl: (json["imgurl"] ?? ''),
      createdAt: DateTime.parse(json["created_at"]),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      "id": this.id,
      "fullname": this.fullname,
      "email": this.email,
      "superadmin": this.superadmin,
      "imgurl": this.imgurl,
      "created_at": this.createdAt!,
    };
  }
}
