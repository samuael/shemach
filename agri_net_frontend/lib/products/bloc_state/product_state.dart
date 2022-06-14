import "../../libs.dart";

class ProductState{}

class MyProductsLoadSuccess extends ProductState{
  List<ProductPost> posts;
  MyProductsLoadSuccess(this.posts);
}
class MyProductsLoadFailed extends ProductState{
  final int statusCode;
  final String msg;
  MyProductsLoadFailed(this.statusCode, this.msg );
}
class MyProductInit extends ProductState{}