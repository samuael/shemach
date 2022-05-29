import '../../libs.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class AuthBloc extends Bloc<AuthEvent, AuthState> {
  AuthRepository repository;
  
  AuthBloc(this.repository) : super(AuthInit());

  @override
  Stream<AuthState> mapEventToState(AuthEvent event) async* {
    if(event is AuthRegisterEvent){
      yield AuthRegistrationOnProgressState(fullname : event.fullname, phone:event.phone, unixTime :((DateTime.now()).millisecondsSinceEpoch/1000).round());
    }else if (event is AuthSubscriberAuthenticatedEvent){
      yield AuthSubscriberAuthenticated(subscriber : event.subscriber, token:event.token);
    }
  }

  Future<AuthenticationResponse> register( RegistrationInput input  ) async {
    return this.repository.register(input);
  }

  Future<SubscriberAuthenticationRespnse>  confirmRegistration(SubscriberConfirmation input) async{
    return this.repository.confirmRegistration(input);
  }
 Future<SubscriberAuthenticationRespnse>  confirmLogin(SubscriberConfirmation input) async{
    return this.repository.confirmLogin(input);
  }

  Future<AuthenticationResponse>  loginSubscriber(String phone) async {
    return this.repository.loginSubscriber(phone);
  }

  // Future<>

}
