import '../../libs.dart';

class ProductRepo {
  ProductProvider? provider;
  ProductRepo({required this.provider});

  Future<ProductResponse> createProduct(int pid, String productName,
      String location, double amounte, double price) {
    return this
        .provider!
        .createProduct(pid, productName, location, amounte, price);
  }

  Future<Product> getProductList() {
    return this.provider!.getProducts();
  }

  // getALLProduct() {}
  // getProductById(int pid) {}
  // updateProduct(int pid) {}
  // deleteProduct(int pid) {}
}
