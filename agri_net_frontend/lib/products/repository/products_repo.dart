import "../../libs.dart";

class ProductsRepo {
  final ProductProvider provider;

  ProductsRepo(this.provider);

  Future<ProductPostResponse>  createProductPost(ProductPostInput input) async {
    return this.provider.createProductPost(input);
  }

}