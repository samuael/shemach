import '../../libs.dart';

class ProductBloc extends Bloc<ProductEvent, ProductState>
    implements Cubit<ProductState> {
  final ProductRepo repo;
  ProductBloc({required this.repo}) : super(GetProductListInItState());

  Stream<ProductState> mapEventToState(ProductEvent event) async* {
    if (event is GetProductListInItEvent) {
      yield GetProductListInItState();
    }
    if (event is GetProductListEvent) {
      yield (GetProductListState());
    }
    if (event is GetProductListOnProdressEvent) {
      yield (GetProductListOnProgresState());
    }
    if (event is ProductListFetchedEvent) {
      yield (ProductListFetchedState(event.products));
    }

    if (event is PostNewProductInItEvent) {
      yield PostNewProductInItState();
    }
    if (event is PostNewProductEvent) {
      yield (PostNewProductState());
    }
    if (event is NewProductPostedEvent) {
      yield (NewProductPostedState(event.product));
    }
  }

  Future<ProductState?> getProductList(GetProductListEvent event) async {
    this.mapEventToState(GetProductListEvent());
    final productList = await repo.getProductList();
    this.mapEventToState(GetProductListOnProdressEvent());
    // if (productList.product != null) {
    //   this.mapEventToState(ProductListFetchedEvent(productList.product!));
    // }
  }

  Future<ProductState?> createNewProduct(PostNewProductEvent event) async {
    this.mapEventToState(PostNewProductInItEvent());
    final productState = await repo.createProduct(event.pid, event.productName,
        event.location, event.amounte, event.price);
    if (productState != null) {
      var newProductState = NewProductPostedEvent(productState.product!);
      this.mapEventToState(NewProductPostedEvent(newProductState.product));
    }
  }
}
