import "../../libs.dart";

class ProductState{}

class MyProductsLoadSuccess extends ProductState{
  List<ProductPost> posts;
  MyProductsLoadSuccess(this.posts);
}
class MyProductsLoadFailed extends ProductState{}
class MyProductInit extends ProductState{}