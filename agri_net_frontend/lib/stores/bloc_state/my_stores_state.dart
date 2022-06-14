import 'package:agri_net_frontend/libs.dart';

class StoreState {}

class MyStoresInit extends StoreState {}

class NewStoreCreatedState extends StoreState {
  Store store;
  NewStoreCreatedState({required this.store});
}

class FailedToCreateStoreState extends StoreState {
  String msg;
  FailedToCreateStoreState({required this.msg});
}
