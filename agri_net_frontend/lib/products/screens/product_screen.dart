import '../../libs.dart';

class ProductScreen extends StatefulWidget {
  static const String RouteName = "products";
  const ProductScreen({Key? key}) : super(key: key);
  @override
  State<ProductScreen> createState() => _ProductScreenState();
}

class _ProductScreenState extends State<ProductScreen> {
  final List<Product> productList = products;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Padding(
        padding: const EdgeInsets.fromLTRB(20, 10, 10, 20),
        child: Center(
          child: Container(
            child: _buildProductList(),
          ),
        ),
      ),
    );
  }

  Widget _buildProductList() {
    return Container(
      child: _buildProductCard(context, products),
    );
  }

  // Widget _buildProductList() {
  //   return Container(
  //     child: BlocBuilder<ProductBloc, ProductState>(
  //       builder: (context, state) {
  //         if (state is GetProductListInItEvent) {
  //           return _buildLoading();
  //         } else if (state is ProductListFetchedState) {
  //           return _buildProductCard(context, state.products);
  //         }
  //         return (Container());
  //       },
  //     ),
  //   );
  // }

  Widget _buildProductCard(BuildContext context, List<Product> p) {
    return ListView.builder(
        itemCount: p.length,
        itemBuilder: (context, index) {
          return Material(
            elevation: 50,
            child: Padding(
              padding: const EdgeInsets.all(8.0),
              child: Container(
                // height: MediaQuery.of(context).size.height * 1 / 2 - 70,
                // width: MediaQuery.of(context).size.width * -20,
                child: Card(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          Flexible(
                            child: ClipRRect(
                              borderRadius: BorderRadius.circular(20),
                              child: Image.asset(
                                "assets/images/bekolo1.jpg",
                                height:
                                    MediaQuery.of(context).size.height * 0.2,
                                width:
                                    MediaQuery.of(context).size.height * 0.4 -
                                        40,
                                fit: BoxFit.cover,
                              ),
                            ),
                          ),
                          Divider(
                            thickness: 0.1,
                            color: Theme.of(context).canvasColor,
                          ),
                          Column(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              ClipRRect(
                                borderRadius: BorderRadius.circular(20),
                                child: Image.asset(
                                  "assets/images/bekolo1.jpg",
                                  height:
                                      MediaQuery.of(context).size.height * 0.1 -
                                          2.5,
                                  width:
                                      MediaQuery.of(context).size.height * 0.2 +
                                          15,
                                  fit: BoxFit.cover,
                                ),
                              ),
                              ClipRRect(
                                borderRadius: BorderRadius.circular(20),
                                child: Image.asset(
                                  "assets/images/bekolo1.jpg",
                                  height:
                                      MediaQuery.of(context).size.height * 0.1 -
                                          2.5,
                                  width:
                                      MediaQuery.of(context).size.height * 0.2 +
                                          15,
                                  fit: BoxFit.cover,
                                ),
                              )
                            ],
                          )
                        ],
                      ),
                      Container(
                        child: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              "Product Type : " + p[index].productName,
                              style: TextStyle(
                                  fontSize: 18, fontWeight: FontWeight.bold),
                            ),
                            Text(
                              "Location :         " + p[index].location,
                              style: TextStyle(
                                  fontSize: 18, fontWeight: FontWeight.bold),
                            ),
                            Text(
                              "Amounte :        " +
                                  p[index].amounte.toString() +
                                  "KG",
                              style: TextStyle(
                                  fontSize: 18, fontWeight: FontWeight.bold),
                            ),
                            Text(
                              "Price :               " +
                                  p[index].amounte.toString() +
                                  "Birr",
                              style: TextStyle(
                                  fontSize: 18, fontWeight: FontWeight.bold),
                            )
                          ],
                        ),
                      )
                    ],
                  ),
                ),
              ),
            ),
          );
        });
  }

  Widget _buildLoading() => Center(child: CircularProgressIndicator());
}
