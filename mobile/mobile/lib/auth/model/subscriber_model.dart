class Subscriber{
  int id;
  String fullname;
  List<int > subscriptions;
  String lang;
  int role;
  String phone;
  Subscriber({required this.id ,required  this.fullname , required this.lang , required  this.subscriptions ,required  this.role ,required  this.phone });


  factory Subscriber.fromJson(Map<String , dynamic> json ){
    return Subscriber(
      id : json["id"]??0,
      fullname : json["fullname"]??'',
      lang : json["lang"]??"amh",
      subscriptions: (json['subscriptions']??[] as List<dynamic>).map<int>((a){return a as int;}).toList(),
      role : json['role']??-1,
      phone : json['phone']??'', 
    );
  }
}


class SubscriberAuthenticationRespnse {
  int statusCode;
  Subscriber? subscriber;
  String token;
  String msg;

  SubscriberAuthenticationRespnse({required this.statusCode , required this.subscriber,
   required this.token ,required this.msg,
   });


  factory SubscriberAuthenticationRespnse.fromJson(Map<String, dynamic> json){
    return SubscriberAuthenticationRespnse(
      statusCode : json['status_code']??999,
      subscriber: Subscriber.fromJson(json["subscriber"]??{}),
      token : json["token"]??'',
      msg : json['msg']??'',   
    );
  }

}