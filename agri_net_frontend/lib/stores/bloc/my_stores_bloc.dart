import '../../libs.dart';

class StoreBloc extends Bloc<StoreEvent, StoreState>
    implements Cubit<StoreState> {
  StoreRepo storeRepo;
  StoreBloc({required this.storeRepo}) : super(MyStoresInit());

  @override
  Stream<StoreState> mapEventToState(StoreEvent storeEvent) async* {
    if (storeEvent is MyStoresEvent) {
      var myStores = await storeRepo.myStores(storeEvent.ownerId);
      yield MyStoresState(myStores: myStores);
    }
  }

  Future<StoreState> createNewStore(
      CreateNewStoreEvent createNewStoreEvent) async {
    var newStoreRespo = await storeRepo.createMerchantStore(
        createNewStoreEvent.storeName,
        createNewStoreEvent.ownerID,
        createNewStoreEvent.kebele,
        createNewStoreEvent.woreda,
        createNewStoreEvent.city,
        createNewStoreEvent.unique_address,
        createNewStoreEvent.region,
        createNewStoreEvent.zone,
        createNewStoreEvent.latitude,
        createNewStoreEvent.longitude);
    if (newStoreRespo.newStore != null) {
      return NewStoreCreatedState(store: newStoreRespo.newStore!);
    }
    return FailedToCreateStoreState(msg: newStoreRespo.msg);
  }
}
