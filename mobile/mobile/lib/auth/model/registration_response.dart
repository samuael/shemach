
class AuthenticationResponse {
  String msg ;
  Map<String , dynamic> errors;
  int statusCode;

  AuthenticationResponse({required this.msg , required this.errors,required this.statusCode});

  factory AuthenticationResponse.fromJson(Map<String, dynamic> json) {
    return AuthenticationResponse(
      msg : json["msg"]??'',
      errors : (json["errors"]??{} as Map<String, dynamic>),
      statusCode : json["status_code"] ?? 999,
    );
  }
}

class RegistrationInput {
  String fullname;
  String phone;
  int role;
  String lang;
  RegistrationInput(
    {
      required this.fullname,
      required this.phone,
      required this.role,
      required this.lang,
    }
  );

  Map<String, dynamic> toJson(){
    return {
        "fullname": this.fullname,
        "phone" : this.phone,
        "lang" : this.lang,
        "role"  : this.role,
    };
  }

}



class SubscriberConfirmation {
  String phone;
  String confirmation;
  SubscriberConfirmation(this.phone, this.confirmation);

  Map<String, String> toJson(){
    return {
      "phone" : this.phone,
      "confirmation" : this.confirmation,  
    };
  }
}