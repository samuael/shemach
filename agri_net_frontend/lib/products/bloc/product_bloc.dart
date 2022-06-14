import '../../libs.dart';

class ProductBloc extends Bloc<ProductEvent, ProductState>
    implements Cubit<ProductState> {
  final ProductsRepo repo;
  ProductBloc({required this.repo}) : super(MyProductInit()){
    on<AddNewProduct>((event , emit) async {
      if (this.state is MyProductsLoadSuccess){
        final thestate = this.state;
        (thestate as MyProductsLoadSuccess).posts.add(event.post);
        emit(thestate);
      }else {
        emit(MyProductsLoadSuccess([event.post]));
      }
    });

    on<LoadMyProductsEvent>((event , emit)async{
      // final response = await this.repo.loadMyPosts();
    });
    
  }


  Future<ProductPostResponse> createProductPost( ProductPostInput input ) async {
    return await this.repo.createProductPost(input);
  } 

}
