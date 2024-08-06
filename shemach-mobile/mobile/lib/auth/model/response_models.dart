import '../../libs.dart';


class ResponseSubscription {
  String msg;
  Map<String , String> errors;
  int statusCode;

  ResponseSubscription({required this.msg , required this.errors, required this.statusCode});

  factory ResponseSubscription.fromJson(Map<String , dynamic> json){
    return ResponseSubscription(
      msg :"${json['msg']??''}",
      errors: json['errors']??{},
      statusCode : int.parse("${json["status_code"]??999}"),
    );
  }
}


class ConfirmedSubscription{
  int statusCode;
  Subscriber subscriber;
  String token;
  String msg;

  ConfirmedSubscription({required this.statusCode,required this.subscriber,required this.token,required this.msg});

  // factory ConfirmedSubscription.fromJson(Map<String , dynamic> json){
  //     return ConfirmedSubscription(statusCode: json['status_code'], subscriber: /*Subscriber.fromJson(json['subscriber'])*/, token: json['token']??'', msg: json['msg']??'');
  // }

}