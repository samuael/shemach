import '../../libs.dart';

class Store {
  int storeID;
  int ownerID;
  Address address;
  int? activeProducts;
  String storeName;
  int? activeContracts;
  DateTime createdAt;
  int createdBy;

  Store(
      {required this.storeID,
      required this.ownerID,
      this.activeProducts,
      required this.storeName,
      this.activeContracts,
      required this.address,
      required this.createdAt,
      required this.createdBy});

  factory Store.fromJson(json) {
    return Store(
        storeID: json["store_id"],
        ownerID: json["owner_id"],
        activeProducts: json["active_products"],
        storeName: json["store_name"],
        activeContracts: json["active_contracts"],
        address: Address.fromJson(json["address"]),
        createdAt: json["created_at"],
        createdBy: json["created_by"]);
  }
}
