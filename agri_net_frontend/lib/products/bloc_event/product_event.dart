import "../../libs.dart";

class ProductEvent {}

class LoadMyProductsEvent extends ProductEvent {
  LoadMyProductsEvent();
}

class AddNewProduct extends ProductEvent {
  ProductPost post;
  AddNewProduct(this.post);
}