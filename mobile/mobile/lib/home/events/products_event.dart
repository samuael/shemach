import "../../libs.dart";

class ProductEvent{}

class ProductsLoadEvent extends ProductEvent{}

class SearchProductEvent extends ProductEvent{
  String text;
  SearchProductEvent(this.text);
}

class ProductUpdateEvent extends ProductEvent{
  ProductUpdate update;
  ProductUpdateEvent(this.update);
}