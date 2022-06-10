import 'package:agri_net_frontend/libs.dart';

Future<ProductType?> selectProductType(
    BuildContext ctx, String text, Function callBack) async {
  ProductType? productType;
  TextEditingController searchController = TextEditingController(text: text);

  List<ProductType> filterProductType(ProductTypeLoadSuccess state) {
    print("Total ProdductTypes ${state.products.length}");
    if (searchController.text == "") {
      return state.products;
    }
    List<ProductType> productTypes = [];
    for (final p in state.products) {
      if (p.name.startsWith(text) ||
          p.name.startsWith(translate(lang, text)) ||
          p.name.startsWith(translate("eng", text)) ||
          p.name.startsWith(translate("amh", text)) ||
          p.name.startsWith(translate("oro", text)) ||
          p.name.startsWith(translate("tig", text))) {
        productTypes.add(p);
      }
    }
    return productTypes;
  }
  showDialog(
      context: ctx,
      builder: (context) {
        return AlertDialog(
          content: Column(
            children: [
              Container(
                padding: EdgeInsets.symmetric(
                  horizontal: 5,
                ),
                decoration: BoxDecoration(
                  border: Border.all(color: Theme.of(context).primaryColor),
                  borderRadius: BorderRadius.circular(10),
                ),
                child: TextField(
                  onChanged: (t) {
                    print(t);
                    if (t.length == 0) {
                      // Navigator.of(context).pop();
                      // callBack("");
                    }
                    callBack(t);
                  },
                  autofocus: true,
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                  ),
                  controller: searchController,
                  decoration: InputDecoration(
                    border: InputBorder.none,
                    suffixIcon: Icon(
                      Icons.search,
                      color: Theme.of(context).primaryColor,
                    ),
                  ),
                ),
              ),
              //
              BlocBuilder<ProductTypeBloc, ProductTypeState>(
                  builder: (context, state) {
                print(state.runtimeType.toString());
                return Container(
                  child: SingleChildScrollView(
                    child: Column(
                      children: (state is ProductTypeLoadSuccess)
                          ? filterProductType(state).map((p) {
                              return GestureDetector(
                                onTap: () {
                                  productType = p;
                                  Navigator.of(context).pop();
                                },
                                child: Container(
                                  padding: EdgeInsets.symmetric(
                                    horizontal: 10,
                                  ),
                                  margin: EdgeInsets.symmetric(
                                    vertical: 2,
                                  ),
                                  decoration: BoxDecoration(
                                      borderRadius: BorderRadius.circular(
                                        5,
                                      ),
                                      border: Border.all(
                                        color:
                                            Theme.of(context).primaryColorLight,
                                      )),
                                  child: Row(
                                      mainAxisAlignment:
                                          MainAxisAlignment.spaceBetween,
                                      children: [
                                        Column(children: [
                                          Text(
                                            p.name,
                                            style: TextStyle(
                                              fontStyle: FontStyle.italic,
                                              fontWeight: FontWeight.bold,
                                            ),
                                          ),
                                          Text(
                                            p.productionArea,
                                            style: TextStyle(
                                              color: Color.fromARGB(
                                                  137, 48, 47, 47),
                                            ),
                                          ),
                                        ]),
                                        Text(
                                            getProductunitByID(p.unitid)!.long),
                                        Icon(
                                          Icons.add,
                                          color: Theme.of(context).primaryColor,
                                        ),
                                      ]),
                                ),
                              );
                            }).toList()
                          : [
                              (state is ProductTypeLoadFailure)
                                  ? Center(
                                      child: Column(
                                          mainAxisAlignment:
                                              MainAxisAlignment.center,
                                          mainAxisSize: MainAxisSize.max,
                                          children: [
                                          GestureDetector(
                                              child: Icon(
                                                Icons.replay,
                                                size: 30,
                                                color: Colors.blue,
                                              ),
                                              onTap: () {
                                                context
                                                    .read<ProductTypeBloc>()
                                                    .add(
                                                        ProductTypesLoadEvent());
                                              }),
                                          Center(
                                            child: Text(
                                              translate(
                                                lang,
                                                " \tSorry!!\n Can't load the product Types\n${state.response.msg}",
                                              ),
                                              textAlign: TextAlign.center,
                                              style: TextStyle(
                                                fontWeight: FontWeight.bold,
                                                fontStyle: FontStyle.italic,
                                              ),
                                            ),
                                          ),
                                        ]))
                                  : Center(
                                      child: Text(
                                        translate(
                                          lang,
                                          " ${state.runtimeType.toString()}\tSorry!!\n No Product type found ",
                                        ),
                                        textAlign: TextAlign.center,
                                        style: TextStyle(
                                          fontWeight: FontWeight.bold,
                                          fontStyle: FontStyle.italic,
                                        ),
                                      ),
                                    ),
                            ],
                    ),
                  ),
                );
              }),
            ],
          ),
        );
      });
  if (productType != null) {
    print("The Product IS ${productType!.name}");
  } else {
    print("The Product is Null");
  }
  return productType;
}
