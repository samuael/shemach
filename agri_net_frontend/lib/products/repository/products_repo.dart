import "../../libs.dart";

class ProductsRepo {
  final ProductProvider provider;

  ProductsRepo(this.provider);

  Future<ProductPostResponse>  createProductPost(ProductPostInput input) async {
    return this.provider.createProductPost(input);
  }

  Future<ProductsResponse> loadMyProductPosts() async {
    return this.provider.loadMyProductPosts();
  }

  Future<ProductsResponse>  loadProducts( int offset, int limit ) async {
    return this.provider.loadProducts(offset,limit);
  }

}