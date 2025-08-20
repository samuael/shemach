import '../../libs.dart';

class ProductItem extends StatefulWidget {
  final ProductType product;
  ProductItem(this.product, {Key? key}) : super(key: key);

  @override
  State<ProductItem> createState() => _ProductItemState();
}

class _ProductItemState extends State<ProductItem> {
  @override
  Widget build(BuildContext context) {
    return ClipRRect(
      borderRadius: BorderRadius.circular(20),
      child: Card(
        // elevation: 6,
        color: Colors.green.shade200,
        child: Container(
          color: Colors.white,
          // width: 350,
          child: Row(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisSize: MainAxisSize.max,
            children: [
              Expanded(
                flex: 25,
                child: Column(
                    mainAxisSize: MainAxisSize.min,
                    mainAxisAlignment: MainAxisAlignment.start,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Container(
                        width: 400,
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(10),
                          border: Border.all(color: Colors.black26),
                        ),
                        padding: EdgeInsets.only(
                          top: 1,
                          bottom: 1,
                          left: 10,
                        ),
                        margin: EdgeInsets.symmetric(horizontal: 40),
                        child: Text(
                          widget.product.name,
                          style: TextStyle(
                            fontSize: 18,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                      ),
                      Row(
                        mainAxisAlignment: MainAxisAlignment.start,
                        children: [
                          Container(
                              padding: EdgeInsets.symmetric(horizontal: 40),
                              child: Text(translate(lang, "የምርት ቦታ") + ":",
                                  style: TextStyle(
                                      fontSize: 13,
                                      fontWeight: FontWeight.bold))),
                          Text(
                            widget.product.productionArea,
                            style: TextStyle(fontSize: 13),
                          ),
                        ],
                      ),
                      Container(
                        padding:
                            EdgeInsets.symmetric(horizontal: 40, vertical: 5),
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.start,
                          children: [
                            Text(translate(lang, "price") + " :",
                                style: TextStyle(
                                    fontSize: 13, fontWeight: FontWeight.bold)),
                            Container(
                              padding: EdgeInsets.symmetric(horizontal: 40),
                              child: Text(
                                "${widget.product.currentPrice} /${widget.product.getProductUnit().long}",
                                style: TextStyle(fontSize: 13),
                              ),
                            ),
                          ],
                        ),
                      ),
                      Container(
                        padding:
                            EdgeInsets.symmetric(vertical: 5, horizontal: 40),
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.start,
                          children: [
                            Text("ከ ",
                                style: TextStyle(
                                    fontSize: 13, fontWeight: FontWeight.bold)),
                            Text(
                              UnixTime(widget.product.lastUpdateTime)
                                  .toString()+" ",
                              style: TextStyle(
                                fontSize: 13,
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                            Text(
                              "በፊት",
                              style: TextStyle(fontSize: 13),
                            ),
                          ],
                        ),
                      ),
                    ]),
              ),
              Expanded(
                flex: 1,
                child: Container(
                    padding: EdgeInsets.symmetric(vertical: 40),
                    child: Icon(Icons.arrow_right_sharp)),
              )
            ],
          ),
        ),
      ),
    );
  }
}
