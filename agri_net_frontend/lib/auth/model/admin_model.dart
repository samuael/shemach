import '../../libs.dart';
{
        "id": 2,
        "firstname": "Samuael",
        "lastname": "Adnew",
        "phone": "+251912131415",
        "email": "samuaeladadnew@outlook.com",
        "created_at": 1649737198,
        "lang": "amh",
        "password": "$2a$10$CCy3zvane0.Ngu8rgBPxNuymHBK5px9eRkDhcJBJBEbhLn8ZJ7CIW"
    }

class Admin {
  int id;
  String firstname;
  String lastname;
  String phone;
  String email;
  String lang;
  DateTime? createdAt;

  bool superadmin;
  String imgurl;
 
 

  Admin({
    required this.id,
    required this.firstname,
    required this.lastname,
    required this.phone,
    required this.email,
    required this.superadmin,
    required this.lang,
    this.createdAt,
    required this.imgurl,
  });

  factory Admin.fromJson(Map<String, dynamic> json) {
    return Admin(
      id: int.parse("${json['id']}"),
      firstname: json["firstname"],
      lastname: json["lastname"],
      phone: json["phone"],
      email: json["email"],
      superadmin: json["superadmin"],
      lang:json["lang"],
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
