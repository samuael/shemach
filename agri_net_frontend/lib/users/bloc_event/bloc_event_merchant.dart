class MerchantsEvent {}

class LoadMyMerchantsEvent extends MerchantsEvent {
  int adminID;
  LoadMyMerchantsEvent({required this.adminID});
}
