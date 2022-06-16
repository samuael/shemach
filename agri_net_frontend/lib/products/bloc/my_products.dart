import "../../libs.dart";

class MyProductsBloc extends Bloc<ProductEvent, ProductState> {
  final ProductsRepo repo;
  MyProductsBloc(this.repo) : super(MyProductInit()) {
    on<LoadMyProductsEvent>((event, emit) async {
      if (!(this.state is MyProductsLoadSuccess)){
        emit(MyProductsLoading());
      }
      final response = await this.repo.loadMyProductPosts();
      if (response.statusCode == 200) {
        if (this.state == MyProductsLoadSuccess) {
          final thestate = this.state;
          (thestate as MyProductsLoadSuccess).posts.addAll(response.posts);
          emit(thestate);
        } else {
          emit(MyProductsLoadSuccess(response.posts));
        }
      } else {
        emit(MyProductsLoadFailed(response.statusCode, response.msg));
      }
    });

    on<AddNewProduct>((event, emit) async {
      if (this.state is MyProductsLoadSuccess) {
        final thestate = this.state;
        (thestate as MyProductsLoadSuccess).posts.add(event.post);
        emit(thestate);
      } else {
        emit(MyProductsLoadSuccess([event.post]));
      }
    });
  }

  Future<ProductPostResponse> createProductPost(ProductPostInput input) async {
    return await this.repo.createProductPost(input);
  }
}
