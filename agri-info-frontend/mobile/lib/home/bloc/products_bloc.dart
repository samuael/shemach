import '../../libs.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class ProductsBloc extends Bloc<ProductEvent, ProductState> {
  ProductsRepository repository;

  ProductsBloc(this.repository) : super(ProductInit());

  @override
  Stream<ProductState> mapEventToState(ProductEvent event) async* {
    if (event is ProductsLoadEvent){
      final response = await this.repository.loadProducts();
      if (response.statusCode == 200 ){
        yield ProductLoadSuccess(response.products);
      }else {
        yield ProductLoadFailure(response);
      }
    }
  }
}
