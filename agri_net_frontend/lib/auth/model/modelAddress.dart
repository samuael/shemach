class Address {
  int ID;
  String Kebele;
  String Woreda;
  String City;
  String UniqueAddressName;
  String Region;
  String Zone;
  double Latitude;
  double Longitude;
  Address(
      {required this.ID,
      required this.Kebele,
      required this.City,
      required this.Latitude,
      required this.Longitude,
      required this.Region,
      required this.UniqueAddressName,
      required this.Woreda,
      required this.Zone});

  Address copyWith(
      {String? kebele,
      String? woreda,
      String? city,
      String? uniqueAddress,
      String? region,
      String? zone,
      double? latitude,
      double? longitude}) {
    return Address(
        ID: this.ID,
        Kebele: kebele ?? this.Kebele,
        City: city ?? this.City,
        Latitude: latitude ?? this.Latitude,
        Longitude: longitude ?? this.Longitude,
        Region: region ?? this.Region,
        UniqueAddressName: uniqueAddress ?? this.UniqueAddressName,
        Woreda: woreda ?? this.Woreda,
        Zone: zone ?? this.Zone);
  }

  factory Address.fromJson(Map<String, dynamic> json) {
    return Address(
        ID: json["id"] ?? 800,
        Kebele: json["kebele"] ?? '',
        City: json["city"] ?? '',
        Latitude: double.parse(json["latitude"].toString()),
        Longitude: double.parse(json["longitude"].toString()),
        Region: json["region"] ?? '',
        UniqueAddressName: json["unique_address"] ?? '',
        Woreda: json["woreda"] ?? '',
        Zone: json["zone"] ?? '');
  }
}
