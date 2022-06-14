import "../../libs.dart";

class ProductTypeState{}

class ProductTypeInit extends ProductTypeState{}

class ProductTypeLoadSuccess extends ProductTypeState{
  List<ProductType> products ; 
  ProductTypeLoadSuccess(this.products);
}

class ProductTypeLoadFailure extends ProductTypeState{
  ProductTypesResponse response ;
  ProductTypeLoadFailure(this.response);
}