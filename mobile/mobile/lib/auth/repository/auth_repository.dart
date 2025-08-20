import "../../libs.dart";

class AuthRepository{

  AuthDataProvider provider;
  AuthRepository(this.provider);

  Future<AuthenticationResponse> register(RegistrationInput input) async {
    return this.provider.register(input);
  }

  Future<SubscriberAuthenticationRespnse>  confirmRegistration(SubscriberConfirmation input) async{
    return this.provider.confirmRegistration(input);
  }
  Future<AuthenticationResponse>  loginSubscriber(String phone) async {
    return this.provider.loginSubscriber(phone);
  }

  Future<SubscriberAuthenticationRespnse>  confirmLogin(SubscriberConfirmation input) async{
    return this.provider.confirmLogin(input);
  }

}