import '../../libs.dart';

class ProductForm extends StatefulWidget {
  static String RouteName = "products/post";
  const ProductForm({Key? key}) : super(key: key);

  @override
  State<ProductForm> createState() => _ProductFormState();
}

class _ProductFormState extends State<ProductForm> {
  ProductType? type;
  bool negotiablePrice = true;

  bool isPosting = false;

  TextEditingController searchController = TextEditingController();
  TextEditingController quantityController = TextEditingController();
  TextEditingController descriptionController = TextEditingController();
  TextEditingController priceController = TextEditingController();

  setMyText(ProductType? produ, String text) {
    setState(() {
      this.type = produ;
      searchController.text = text;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        child: Column(
          children: [
            Container(
              decoration: BoxDecoration(
                border: Border.all(color: Theme.of(context).primaryColor),
                borderRadius: BorderRadius.circular(10),
              ),
              padding: EdgeInsets.symmetric(
                horizontal: 40,
                vertical: 30,
              ),
              margin: EdgeInsets.symmetric(
                vertical: 10,
                horizontal: 10,
              ),
              width: MediaQuery.of(context).size.width * 0.8,
              height: 30,
              child: Text(
                "I am starting an startup. Fot that, I need an Android app. Its functionality will be like freelance or fiverr. Rest details in ib. I have a very low budget. Skills: Android | iPhone | Java | Mobile App Development | PHP",
              ),
            ),
            Container(
              padding: EdgeInsets.symmetric(
                horizontal: 10,
              ),
              child: Column(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Container(
                    padding: EdgeInsets.symmetric(
                      vertical: 3,
                      horizontal: 10,
                    ),
                    child: Row(
                      children: [
                        Expanded(
                          flex: 2,
                          child: Container(
                            width: 100,
                            child: Text(
                              translate(lang, " Product Type "),
                              style: TextStyle(
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                          ),
                        ),
                        Expanded(
                          flex: 5,
                          child: type == null
                              ? Container(
                                  padding: EdgeInsets.symmetric(
                                    horizontal: 5,
                                  ),
                                  decoration: BoxDecoration(
                                    border: Border.all(
                                        color: Theme.of(context).primaryColor),
                                    borderRadius: BorderRadius.circular(10),
                                  ),
                                  child: TextField(
                                    controller: searchController,
                                    onChanged: (String text) async {
                                      final state =
                                          context.read<ProductTypeBloc>().state;
                                      List<ProductType> products = [];
                                      if (state is ProductTypeLoadSuccess) {
                                        products = state.products;
                                      }
                                      print("Products Length: ${products}");
                                      // type = await selectProductType(context,
                                      //     text, state, products, setMyText);
                                      Navigator.of(context).pushNamed(
                                          ProductTypeSelectionScreen.RouteName,
                                          arguments: {
                                            "state": state,
                                            "products": products,
                                            "callback": setMyText,
                                            "text": text,
                                          });
                                    },
                                    decoration: InputDecoration(
                                      border: InputBorder.none,
                                      suffixIcon: Icon(
                                        Icons.search,
                                        color: Theme.of(context).primaryColor,
                                      ),
                                    ),
                                  ),
                                )
                              : Container(
                                  decoration: BoxDecoration(
                                    border: Border.all(
                                      color: Theme.of(context).primaryColor,
                                    ),
                                    borderRadius: BorderRadius.circular(5),
                                  ),
                                  child: Row(
                                    mainAxisSize: MainAxisSize.min,
                                    mainAxisAlignment: MainAxisAlignment.center,
                                    children: [
                                      Text(
                                        type!.name,
                                        textAlign: TextAlign.center,
                                        style: TextStyle(
                                          fontWeight: FontWeight.bold,
                                          fontStyle: FontStyle.italic,
                                        ),
                                      ),
                                      IconButton(
                                        onPressed: () {
                                          setState(() {
                                            type = null;
                                            searchController.text = "";
                                          });
                                        },
                                        icon: Icon(
                                          Icons.cancel,
                                          color: Theme.of(context).primaryColor,
                                        ),
                                      ),
                                    ],
                                  ),
                                ),
                        )
                      ],
                    ),
                  ),
                  Container(
                    padding: EdgeInsets.symmetric(
                      vertical: 3,
                      horizontal: 10,
                    ),
                    child: Row(
                      children: [
                        Expanded(
                          flex: 2,
                          child: Container(
                            width: 100,
                            child: Text(
                              translate(lang, "Quantity "),
                              style: TextStyle(
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                          ),
                        ),
                        Expanded(
                          flex: 5,
                          child: Container(
                            padding: EdgeInsets.symmetric(
                              horizontal: 5,
                            ),
                            decoration: BoxDecoration(
                              border: Border.all(
                                  color: Theme.of(context).primaryColor),
                              borderRadius: BorderRadius.circular(10),
                            ),
                            child: TextField(
                              controller: quantityController,
                              textAlignVertical: TextAlignVertical.center,
                              textAlign: TextAlign.center,
                              keyboardType: TextInputType.number,
                              decoration: InputDecoration(
                                border: InputBorder.none,
                                suffixIcon: Icon(
                                  Icons.money,
                                  color: Theme.of(context).primaryColor,
                                ),
                              ),
                            ),
                          ),
                        )
                      ],
                    ),
                  ),
                  Container(
                    padding: EdgeInsets.symmetric(
                      vertical: 3,
                      horizontal: 10,
                    ),
                    child: Row(
                      children: [
                        Expanded(
                          flex: 2,
                          child: Container(
                            width: 100,
                            child: Text(
                              translate(lang, "Description "),
                              style: TextStyle(
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                          ),
                        ),
                        Expanded(
                          flex: 5,
                          child: Container(
                            padding: EdgeInsets.symmetric(
                              horizontal: 5,
                            ),
                            decoration: BoxDecoration(
                              border: Border.all(
                                  color: Theme.of(context).primaryColor),
                              borderRadius: BorderRadius.circular(10),
                            ),
                            child: TextField(
                              maxLength: 500,
                              maxLines: null,
                              controller: descriptionController,
                              textAlignVertical: TextAlignVertical.center,
                              textAlign: TextAlign.center,
                              keyboardType: TextInputType.multiline,
                              decoration: InputDecoration(
                                border: InputBorder.none,
                                suffixIcon: Icon(
                                  Icons.description,
                                  color: Theme.of(context).primaryColor,
                                ),
                              ),
                            ),
                          ),
                        )
                      ],
                    ),
                  ),
                  Container(
                    padding: EdgeInsets.symmetric(
                      vertical: 3,
                      horizontal: 10,
                    ),
                    child: Row(
                      children: [
                        Expanded(
                          flex: 2,
                          child: Container(
                            width: 100,
                            child: Text(
                              translate(lang, "Price "),
                              style: TextStyle(
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                          ),
                        ),
                        Expanded(
                          flex: 5,
                          child: Container(
                            padding: EdgeInsets.symmetric(
                              horizontal: 5,
                            ),
                            decoration: BoxDecoration(
                              border: Border.all(
                                  color: Theme.of(context).primaryColor),
                              borderRadius: BorderRadius.circular(10),
                            ),
                            child: TextField(
                              controller: priceController,
                              textAlignVertical: TextAlignVertical.center,
                              textAlign: TextAlign.center,
                              keyboardType: TextInputType.number,
                              decoration: InputDecoration(
                                border: InputBorder.none,
                                suffixIcon: Icon(
                                  Icons.money,
                                  color: Theme.of(context).primaryColor,
                                ),
                              ),
                            ),
                          ),
                        )
                      ],
                    ),
                  ),
                  Container(
                    padding: EdgeInsets.symmetric(
                      vertical: 3,
                      horizontal: 10,
                    ),
                    child: Row(
                      children: [
                        Expanded(
                          flex: 2,
                          child: Container(
                            width: 100,
                            child: Text(
                              translate(lang, "Negotiable \n Price"),
                              style: TextStyle(
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                          ),
                        ),
                        Expanded(
                          flex: 5,
                          child: Container(
                            padding: EdgeInsets.symmetric(
                              horizontal: 5,
                            ),
                            decoration: BoxDecoration(
                              border: Border.all(
                                  color: Theme.of(context).primaryColor),
                              borderRadius: BorderRadius.circular(10),
                            ),
                            child: Row(
                              children: [
                                Expanded(
                                  flex: 2,
                                  child: Text(translate(lang, "Yes")),
                                ),
                                Expanded(
                                  flex: 4,
                                  child: Radio(
                                    groupValue: negotiablePrice,
                                    value: true,
                                    onChanged: (resu) {
                                      setState(() {
                                        negotiablePrice = true;
                                      });
                                    },
                                  ),
                                ),
                                Expanded(
                                  flex: 1,
                                  child: SizedBox(
                                    child: Container(
                                      child: Text("|"),
                                    ),
                                  ),
                                ),
                                Expanded(
                                  flex: 2,
                                  child: Text(translate(lang, "No")),
                                ),
                                Expanded(
                                  flex: 4,
                                  child: Radio(
                                    groupValue: negotiablePrice,
                                    value: false,
                                    onChanged: (resu) {
                                      setState(() {
                                        negotiablePrice = false;
                                      });
                                    },
                                  ),
                                )
                              ],
                            ),
                          ),
                        ),
                      ],
                    ),
                  ),
                  SizedBox(
                    height: 20,
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      GestureDetector(
                        onTap: () async {
                          isPosting = false;
                          if (type != null &&
                              quantityController.text != "" &&
                              descriptionController.text != "" &&
                              quantityController.text != "" &&
                              priceController.text != "") {
                            final response = await context
                                .read<MyProductsBloc>()
                                .createProductPost(
                                  ProductPostInput(
                                    typeID: type!.id,
                                    sellingPrice:
                                        double.parse(priceController.text),
                                    storeid: 1,
                                    description: descriptionController.text,
                                    quantity:
                                        int.parse(quantityController.text),
                                    negotiablePrice: negotiablePrice,
                                  ),
                                );
                                print("\n\n\n\n");
                                print(response.msg );
                                print("\n\n\n\n");

                            if (response.statusCode == 200 ||
                                response.statusCode == 201) {
                              context
                                  .read<MyProductsBloc>()
                                  .add(AddNewProduct(response.crop!));
                            }
                          }
                        },
                        child: ClipRRect(
                          borderRadius: BorderRadius.circular(
                            5,
                          ),
                          child: Container(
                            color: Theme.of(context).primaryColor,
                            padding: EdgeInsets.symmetric(
                              horizontal: 40,
                              vertical: 10,
                            ),
                            child: Text(
                              translate(lang, "Post"),
                              style: TextStyle(
                                fontWeight: FontWeight.bold,
                                color: Colors.white,
                              ),
                            ),
                          ),
                        ),
                      )
                    ],
                  )
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
