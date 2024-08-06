class MerchantsEvent {}

class LoadMyMerchantsEvent extends MerchantsEvent {
  int adminID;
  LoadMyMerchantsEvent({required this.adminID});
}

class DeleteMerchantEvent extends MerchantsEvent {
  int userID;
  DeleteMerchantEvent({required this.userID});
}
