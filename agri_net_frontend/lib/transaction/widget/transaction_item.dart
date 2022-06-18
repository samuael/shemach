import "../../libs.dart";
import "package:flutter/material.dart";

class TransactionItem extends StatefulWidget {
  Transaction transaction;
  TransactionItem(this.transaction, {Key? key}) : super(key: key);

  @override
  State<TransactionItem> createState() => _TransactionItemState();
}

class _TransactionItemState extends State<TransactionItem> {
  @override
  Widget build(BuildContext context) {
    return Container(
      child: Center(
        child: Column(
          children: [
            context
                        .watch<ProductsBloc>()
                        .getProductPostByID(widget.transaction.productId) !=
                    null
                ? ProductItemSmallView(context
                    .watch<ProductsBloc>()
                    .getProductPostByID(widget.transaction.productId)!)
                : Container(child: CircularProgressIndicator()),
            Container(
              width: MediaQuery.of(context).size.width * 0.88,
              padding: EdgeInsets.symmetric(
                horizontal: 10,
              ),
              decoration: BoxDecoration(
                border: Border.all(
                  color: Theme.of(context).primaryColorLight,
                ),
                borderRadius: BorderRadius.circular(5),
              ),
              // margin: EdgeInsets.symmetric(
              //   horizontal: 5,
              // ),
              child: SingleChildScrollView(
                child: Column(mainAxisSize: MainAxisSize.min, children: [
                  Container(
                    child: Row(
                      children: [
                        Expanded(
                          flex: 2,
                          child: Text(
                            translate(lang, "Quantity"),
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                        ),
                        Expanded(
                          flex: 1,
                          child: Text(" : "),
                        ),
                        Expanded(
                          flex: 5,
                          child: Text(
                              "${widget.transaction.quantity}   ${context.watch<ProductsBloc>().getProductPostByID(widget.transaction.productId) != null ? ((context.watch<ProductTypeBloc>().state is ProductTypeLoadSuccess) ? ((context.watch<ProductTypeBloc>().state as ProductTypeLoadSuccess).getProductTypeByID(context.watch<ProductsBloc>().getProductPostByID(widget.transaction.productId)!.typeId)!.getProductUnit().long) : "") : ""}" +
                                  "  "
                                      "${context.watch<ProductsBloc>().getProductPostByID(widget.transaction.productId) != null ? ((context.watch<ProductTypeBloc>().state is ProductTypeLoadSuccess) ? ((context.watch<ProductTypeBloc>().state as ProductTypeLoadSuccess).getProductTypeByID(context.watch<ProductsBloc>().getProductPostByID(widget.transaction.productId)!.typeId)!.name) : "") : ""}"),
                        )
                      ],
                    ),
                  ),
                  Container(
                    child: Row(
                      children: [
                        Expanded(
                          flex: 2,
                          child: Text(
                            translate(lang, "Price"),
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                        ),
                        Expanded(
                          flex: 1,
                          child: Text(" : "),
                        ),
                        Expanded(
                          flex: 5,
                          child: Text(
                              "${widget.transaction.price}  ${translate(lang, " Birr ")} "),
                        )
                      ],
                    ),
                  ),
                  Container(
                    child: Row(
                      children: [
                        Expanded(
                          flex: 2,
                          child: Text(
                            translate(lang, "Created At"),
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                        ),
                        Expanded(
                          flex: 1,
                          child: Text(" : "),
                        ),
                        Expanded(
                          flex: 5,
                          child: Text(
                              "${UnixTime(widget.transaction.createdAt).toString()}  ${translate(lang, " before ")}"),
                        )
                      ],
                    ),
                  ),
                  Container(
                    width: 200,
                    padding: EdgeInsets.symmetric(
                      horizontal: 8,
                      vertical: 3,
                    ),
                    decoration: BoxDecoration(
                      borderRadius: BorderRadius.circular(5),
                      border: Border.all(
                        color: Colors.white,
                      ),
                      color:
                          StaticDataStore.ID == widget.transaction.requesterId
                              ? Theme.of(context).primaryColor
                              : Colors.red,
                    ),
                    child: Text(
                      StaticDataStore.ID == widget.transaction.requesterId
                          ? "Requested by you"
                          : "Requested Other",
                      style: TextStyle(
                        color: Colors.white,
                      ),
                    ),
                  )
                ]),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
