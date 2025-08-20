import "../../libs.dart";

class ProductState{}

class ProductInit extends ProductState{}

class ProductLoadSuccess extends ProductState{
  List<ProductType> products ; 
  ProductLoadSuccess(this.products);
}

class ProductLoadFailure extends ProductState{
  ProductsResponse response ;
  ProductLoadFailure(this.response);
}