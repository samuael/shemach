import "../../libs.dart";

class AuthEvent{}


class AuthLoginEvent extends AuthEvent{}

class AuthRegisterEvent extends AuthEvent{
  String fullname;
  String phone;
  String otp;
  AuthRegisterEvent({required this.fullname, required this.phone, required this.otp});
}

class  AuthConfirmEvent extends AuthEvent{}

class AuthSubscriberAuthenticatedEvent extends AuthEvent{
  Subscriber subscriber;
  String token;
  
  AuthSubscriberAuthenticatedEvent({
    required this.subscriber, 
    required this.token});
}