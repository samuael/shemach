import '../../libs.dart';

// ID        uint64 `json:"id,omitempty"`
// Firstname string `json:"firstname"`
// Lastname  string `json:"lastname"`
// Phone     string `json:"phone,omitempty"`
// Email     string `json:"email"`
// Imgurl    string `json:"imgurl,omitempty"`
// CreatedAt uint64 `json:"created_at,omitempty"`
// Lang      string `json:"lang"`
// Password  string `json:"-"`

class User {
  int id;
  String firstname;
  String lastname;
  String phone;
  String email;
  String lang;
  DateTime? createdAt;
  String imgurl;

  User({
    required this.id,
    required this.firstname,
    required this.lastname,
    required this.phone,
    required this.email,
    required this.lang,
    this.createdAt,
    required this.imgurl,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: int.parse("${json['id']}"),
      firstname: json["firstname"],
      lastname: json["lastname"],
      phone: json["phone"] ?? '',
      email: json["email"],
      lang: json["lang"],
      imgurl: (json["imgurl"] ?? ''),
      createdAt:
          DateTime.fromMillisecondsSinceEpoch((json["created_at"] ?? 0) * 1000),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      "id": this.id,
      "firstname": this.firstname,
      "lastname": this.lastname,
      "email": this.email,
      "phone": this.phone,
      "imgurl": this.imgurl,
      "lang": this.lang,
      "created_at": (this.createdAt!.millisecondsSinceEpoch) / 1000,
    };
  }
}
