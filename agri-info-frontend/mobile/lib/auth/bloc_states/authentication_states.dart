import 'package:mobile/auth/auth.dart';

class AuthState {}

class AuthRegistrationOnProgressState extends AuthState{
  String fullname;
  String phone;
  int unixTime;

  AuthRegistrationOnProgressState({required this.fullname, required this.phone , required this.unixTime});
}

// class AuthAutenticatedState{
//   Subscriber subscriber;
//   AuthAutenticatedState(this.subscriber);
// }

class AuthOnConfirmationState {}

// AuthInit 
class AuthInit extends AuthState{}


class AuthSubscriberAuthenticated extends AuthState{
  Subscriber subscriber;
  String token;
  
  AuthSubscriberAuthenticated({
    required this.subscriber, 
    required this.token});
}