import 'package:mobile/auth/auth.dart';

class AuthState {}

class AuthAutenticatedState{
  Subscriber subscriber;
  AuthAutenticatedState(this.subscriber);
}

class AuthRegistrationOnProgressState {
  int unixTime;

  AuthRegistrationOnProgressState(this.unixTime);
}

class AuthOnConfirmationState {
  
}