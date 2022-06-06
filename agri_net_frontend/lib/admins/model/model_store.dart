class Store {
  int storeID;
  int ownerID;
  int addressId;
  int activeProducts;
  String storeName;
  int activeContracts;
  DateTime createdAt;
  int createdBy;

  Store(
      {required this.storeID,
      required this.ownerID,
      required this.activeProducts,
      required this.storeName,
      required this.activeContracts,
      required this.addressId,
      required this.createdAt,
      required this.createdBy});

  factory Store.fromJson(json) {
    return Store(
        storeID: json["store_id"],
        ownerID: json["owner_id"],
        activeProducts: json["active_products"],
        storeName: json["store_name"],
        activeContracts: json["active_contracts"],
        addressId: json["address_id"],
        createdAt: json["created_at"],
        createdBy: json["created_by"]);
  }
}
