import 'package:agri_net_frontend/products/widgets/productForm.dart';

import '../../libs.dart';

class ProductScreen extends StatefulWidget {
  static const String RouteName = "/products/screen";
  const ProductScreen({Key? key}) : super(key: key);
  @override
  State<ProductScreen> createState() => _ProductScreenState();
}

class _ProductScreenState extends State<ProductScreen> {
  bool productsList = true;
  PageController pageController = PageController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).canvasColor,
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
        centerTitle: true,
        actions: [
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
              // Navigator.of(context).pushNamed(ProductForm.RouteName);
            },
          ),
        ],
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          setState(() {
            productsList = !productsList;
          });
          if (productsList) {
            pageController.nextPage(
              duration: Duration(
                milliseconds: 500,
              ),
              curve: Curves.ease,
            );
          } else {
            pageController.previousPage(
              duration: Duration(
                milliseconds: 500,
              ),
              curve: Curves.ease,
            );
          }
        },
        child: Icon(
          productsList ? Icons.add : Icons.list,
          color: Colors.white,
        ),
      ),
      body: Container(
          child: PageView(
              controller: pageController,
              onPageChanged: (val) {
                setState(() {
                  productsList = !productsList;
                });
              },
              children: [
            ProductsList(),
            ProductForm(),
          ])),
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
                            "Unix id : " + p[index].unit_id.toString(),
                            style: TextStyle(
                                fontSize: 18, fontWeight: FontWeight.bold),
                          ),
                          Text(
                            "Product Type : " + p[index].name,
                            style: TextStyle(
                                fontSize: 18, fontWeight: FontWeight.bold),
                          ),
                          Text(
                            "Location :         " + p[index].production_area,
                            style: TextStyle(
                                fontSize: 18, fontWeight: FontWeight.bold),
                          ),
                          Text(
                            "Amounte :        " +
                                p[index].currentPrice.toString() +
                                "KG",
                            style: TextStyle(
                                fontSize: 18, fontWeight: FontWeight.bold),
                          ),
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
