import '../../libs.dart';

class UserProfile {
  int id;
  String firstname;
  String lastname;
  String phone;
  String email;
  String lang;
  DateTime? createdAt;
  String imgurl;

  UserProfile({
    required this.id,
    required this.firstname,
    required this.lastname,
    required this.phone,
    required this.email,
    required this.lang,
    this.createdAt,
    required this.imgurl,
  });

  factory UserProfile.fromJson(Map<String, dynamic> json) {
    return UserProfile(
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
      "created_at": this.createdAt!,
      "lang": this.lang,
      "created_at": (this.createdAt!.millisecondsSinceEpoch) / 1000,
    };
  }
}
