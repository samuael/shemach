import "../../libs.dart";

class ProductsRepository {
  ProductsProvider provider;

  ProductsRepository(this.provider);

  Future<ProductsResponse> loadProducts() async {
    return this.provider.loadProducts();
  }
}
