import 'package:agri_net_frontend/products/screens/productForm.dart';

import '../../libs.dart';

class ProductScreen extends StatefulWidget {
  static const String RouteName = "/products/screen";
  const ProductScreen({Key? key}) : super(key: key);
  @override
  State<ProductScreen> createState() => _ProductScreenState();
}

class _ProductScreenState extends State<ProductScreen> {
  final List<Product> productList = products;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).canvasColor,
        elevation: 5,
        toolbarHeight: MediaQuery.of(context).size.height / 13,
        leading: IconButton(
            color: Colors.black,
            onPressed: () {
              Navigator.pop(context);
            },
            icon: BackButton()),
        title: Text(
          "Products",
          style: TextStyle(
              fontSize: 18, fontWeight: FontWeight.bold, color: Colors.black),
        ),
      ),
      body: Padding(
        padding: const EdgeInsets.fromLTRB(20, 10, 10, 20),
        child: Center(
          child: Column(
            children: [
              InkWell(
                child: Container(
                  decoration:
                      BoxDecoration(borderRadius: BorderRadius.circular(10)),
                  child: Text(
                    "New",
                    style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                  ),
                ),
                onTap: () {
                  context.read<ProductBloc>().add(PostNewProductInItEvent());
                  Navigator.of(context).pushNamed(ProductFormScreen.RouteName);
                },
              ),
              _buildProductList()
            ],
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

  Widget _buildProductCard(BuildContext context, List<Product> p) {
    return ListView.builder(
        itemCount: p.length,
        itemBuilder: (context, index) {
          return Padding(
            padding: const EdgeInsets.all(8.0),
            child: Container(
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
                              height: MediaQuery.of(context).size.height * 0.2,
                              width:
                                  MediaQuery.of(context).size.height * 0.4 - 40,
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
          );
        });
  }

  Widget _buildLoading() => Center(child: CircularProgressIndicator());
}
