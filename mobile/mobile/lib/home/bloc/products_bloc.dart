import '../../libs.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class ProductsBloc extends Bloc<ProductEvent, ProductState> {
  ProductsRepository repository;

  ProductsBloc(this.repository) : super(ProductInit()) {
    on<ProductsLoadEvent>((event, emit) async {
      final response = await this.repository.loadProducts();
      if (response.statusCode == 200) {
        emit(ProductLoadSuccess(response.products));
      } else if (response.statusCode == 1000) {
        emit(this.state);
      } else {
        emit(ProductLoadFailure(response));
      }
    });
    on<SearchProductEvent>((event, emit) async {
      final response = await this.repository.searchProducts(event.text);
      if (response.statusCode == 200) {
        emit(ProductInit());
        emit(ProductLoadSuccess(response.products));
      } else {
        emit(this.state);
      }
    });
    on<ProductUpdateEvent>((event, emit){
      print("Event Called");
      if(this.state is ProductLoadSuccess){
        for (final pr in (this.state as ProductLoadSuccess).products ){
          if(pr.id == event.update.productID){
            pr.currentPrice= event.update.cost;
            print("Updating....");
            emit(ProductLoadSuccess((this.state as ProductLoadSuccess).products));
          }
        }
      }
    });
    
  }

  // subscribeForProduct ... 
  Future<StatusAndMessage>  subscribeForProduct(int productID) async {
    return this.repository.subscribeForProduct(productID);
  }
  Future<StatusAndMessage>  unSubscribeForProduct(int productID) async {
    return this.repository.unSubscribeForProduct(productID);
  }
}
