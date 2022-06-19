import '../../libs.dart';

class MerchantsState {}

class MerchantsInItState extends MerchantsState {}

class MerchantsLoadedState extends MerchantsState {
  List<Merchant> merchants;
  MerchantsLoadedState({required this.merchants});
}

class MerchantsLoadingState extends MerchantsState {}

class MerchantsLoadingFailedState extends MerchantsState {
  int statusCode;
  String msg;
  MerchantsLoadingFailedState({required this.statusCode, required this.msg});
}
