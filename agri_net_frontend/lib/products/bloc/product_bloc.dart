import '../../libs.dart';

class ProductsBloc extends Bloc<ProductEvent, ProductState>
    implements Cubit<ProductState> {
  final ProductsRepo repo;
  ProductsBloc({required this.repo}) : super(ProductInit()) {
    on<LoadProductsEvent>((event, emit) async {
      int offset = 0;
      int limit = 0;
      if (this.state is ProductsLoadSuccess) {
        offset = (this.state as ProductsLoadSuccess).posts.length;
        limit = offset + 20;
      }
      final response = await this.repo.loadProducts(offset, limit);
      if (response.statusCode == 200 || response.statusCode == 201) {
        if (this.state is ProductsLoadSuccess) {
          final thestate = this.state;
          (thestate as ProductsLoadSuccess).posts.addAll(response.posts);
          emit(thestate);
        } else {
          emit(ProductsLoadSuccess(response.posts));
        }
      } else {
        if (!(this.state is ProductsLoadSuccess)) {
          emit(
            ProductLoadFailed(
              statusCode: response.statusCode,
              msg: response.msg,
            ),
          );
        }
      }
    });
  }

  Future<ProductPostResponse> createProductPost(ProductPostInput input) async {
    return await this.repo.createProductPost(input);
  }
}
