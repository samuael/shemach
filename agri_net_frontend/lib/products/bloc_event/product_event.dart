import '../../libs.dart';

class ProductEvent {}

class GetProductListInItEvent extends ProductEvent {}

class GetProductListEvent extends ProductEvent {}

class GetProductListOnProdressEvent extends ProductEvent {}

class ProductListFetchedEvent extends ProductEvent {
  List<Product> products;
  ProductListFetchedEvent(this.products);
}

// POST EVENTS

class PostNewProductInItEvent extends ProductEvent {}

class PostNewProductEvent extends ProductEvent {
  int pid;
  String productName;
  String location;
  double amounte;
  double price;

  PostNewProductEvent(
      this.pid, this.productName, this.location, this.amounte, this.price);
}

class NewProductPostedEvent extends ProductEvent {
  Product product;
  NewProductPostedEvent(this.product);
}
