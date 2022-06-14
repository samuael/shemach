import "../../libs.dart";

class MyProductsBloc extends Bloc<ProductEvent, ProductState> {
  final ProductsRepo repo;
  MyProductsBloc(this.repo) : super(MyProductInit()) {
    on<LoadMyProductsEvent>((event, emit) async {
      final response = await repo.createProductPost(event.input);
      if (response.statusCode == 1000) {
        emit(this.state);
      } else if (response.statusCode == 200 || response.statusCode == 201) {
        if (this.state is MyProductsLoadSuccess) {
          final thestate = this.state;
          emit(MyProductInit());
          (thestate as MyProductsLoadSuccess).posts.add(response.crop!);
          emit(thestate);
        } else {
          emit(MyProductsLoadSuccess([response.crop!]));
        }
      } else {
        emit(MyProductsLoadFailed());
      }
    });

    on<AddNewProduct>((event , emit) async {
      if (this.state is MyProductsLoadSuccess){
        final thestate = this.state;
        (thestate as MyProductsLoadSuccess).posts.add(event.post);
        emit(thestate);
      }else {
        emit(MyProductsLoadSuccess([event.post]));
      }
    });
    
  }

  Future<ProductPostResponse> createProductPost(ProductPostInput input) async {
    return await this.repo.createProductPost(input);
  }
}
