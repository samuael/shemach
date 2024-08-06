import 'package:mobile/libs.dart';

class ProductType {
  int id;
  String name;
  String productionArea;
  int unitid;
  double currentPrice;
  int createdBy;
  int createdAt;
  int lastUpdateTime;

  ProductType({
    required this.id,
    required this.name,
    required this.productionArea,
    required this.unitid,
    required this.currentPrice,
    required this.createdBy,
    required this.createdAt,
    required this.lastUpdateTime,
  });

  factory ProductType.fromJson(Map<String, dynamic> json) {
    return ProductType(
      id: json["id"] ?? 0,
      name: json["name"] ?? "",
      productionArea: json["production_area"] ?? '',
      unitid: json["unit_id"] ?? 1,
      currentPrice: double.parse("${json["current_price"]}", (el){return 0.0;}),
      createdBy: json["created_by"] ?? 0,
      createdAt: json["created_at"] ?? 0,
      lastUpdateTime: json["last_update_time"] ?? 0,
    );
  }

  static List<ProductType> fromListOfJSON(List<dynamic> jsons) {
    return jsons.map<ProductType>((a) {
      return ProductType.fromJson(a as Map<String, dynamic>);
    }).toList();
  }

  ProductUnit getProductUnit() {
      final punit = getProductunitByID(this.unitid);
      if (punit==null){
        return ProductUnit(category:"unknown",id:0, short:"unk", long : "unknown");
      }
      return punit;
  }
  
}

class ProductUnit {
  int id;
  String short;
  String long;
  String category;
  ProductUnit({
    required this.category,
    required this.id,
    required this.short,
    required this.long,
  });


}

final productUnits = [
  ProductUnit(category: "mass", id: 1, short: "K", long: "killo Gram"),
  ProductUnit(category: "mass", id: 2, short: "g", long: "gram"),
  ProductUnit(category: "mass", id: 3, short: "Kun", long: "kuntal"),
  ProductUnit(category: "mass", id: 4, short: "Ton", long: "ton"),
  ProductUnit(category: "volume", id: 5, short: "L", long: "litter"),
  ProductUnit(category: "volume", id: 6, short: "M3", long: "meter cube"),
  ProductUnit(category: "volume", id: 7, short: "Gal", long: "gallon"),
  ProductUnit(category: "item", id: 8, short: "SIT", long: "single item"),
  ProductUnit(category: "item", id: 9, short: "DZ", long: "dozen"),
  ProductUnit(category: "item", id: 10, short: "HDZ", long: "half dozen"),
  ProductUnit(category: "item", id: 11, short: "QDZ", long: "quarter dozen"),
  ProductUnit(category: "size", id: 12, short: "SM", long: "small"),
  ProductUnit(category: "size", id: 13, short: "LG", long: "large"),
  ProductUnit(category: "size", id: 14, short: "MD", long: "medium"),
  ProductUnit(category: "length", id: 14, short: "M", long: "meter"),
  ProductUnit(category: "length", id: 15, short: "KM", long: "killo meter"),
  ProductUnit(category: "length", id: 16, short: "cm", long: "centi meter"),
  ProductUnit(category: "length", id: 17, short: "Milli", long: "mile"),
];

ProductUnit? getProductunitByID(int id) {
  for (final a in productUnits) {
    if (a.id == id) {
      return a;
    }
  }
  return null;
}
