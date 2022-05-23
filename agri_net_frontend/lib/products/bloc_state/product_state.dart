import 'package:agri_net_frontend/products/products.dart';

class ProductState {}

// Retrieve

class GetProductListInItState extends ProductState {}

class GetProductListState extends ProductState {}

class GetProductListOnProgresState extends ProductState {}

class ProductListFetchedState extends ProductState {
  List<Product> products;
  ProductListFetchedState(this.products);
}

// POST

class PostNewProductInItState extends ProductState {}

class PostNewProductState extends ProductState {}

class NewProductPostedState extends ProductState {
  Product product;
  NewProductPostedState(this.product);
}
