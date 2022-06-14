import '../../libs.dart';
import "package:flutter_bloc/flutter_bloc.dart";

class SubscriptionProductItem extends StatefulWidget {
  final ProductType product;
  bool subscribed;
  SubscriptionProductItem(this.product, this.subscribed, {Key? key})
      : super(key: key);

  @override
  State<SubscriptionProductItem> createState() => _SubscriptionPageState();
}

class _SubscriptionPageState extends State<SubscriptionProductItem> {
  @override
  Widget build(BuildContext context) {
    return ClipRRect(
      borderRadius: BorderRadius.circular(20),
      child: Card(
        // elevation: 6,
        color: Theme.of(context).canvasColor,
        child: Container(
          alignment: Alignment.topLeft,
          width: MediaQuery.of(context).size.width,
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              Container(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Container(
                      padding: EdgeInsets.symmetric(vertical: 8),
                      margin: EdgeInsets.symmetric(horizontal: 120),
                      child: Text(translate(lang, widget.product.name),
                          style: TextStyle(
                              fontSize: 14, fontWeight: FontWeight.bold)),
                    ),
                    Row(
                      mainAxisSize: MainAxisSize.max,
                      crossAxisAlignment: CrossAxisAlignment.center,
                      children: [
                        Container(
                            padding: EdgeInsets.symmetric(horizontal: 30),
                            child: Text(translate(lang, "production area"),
                                style: TextStyle(
                                    fontSize: 13,
                                    fontWeight: FontWeight.bold))),
                        Text(
                          translate(lang, widget.product.productionArea),
                          style: TextStyle(fontSize: 13),
                        ),
                      ],
                    ),
                    Row(
                      children: [
                        Container(
                          padding: EdgeInsets.symmetric(vertical: 8),
                          margin: EdgeInsets.symmetric(horizontal: 30),
                          child: Row(
                            children: [
                              Text(translate(lang , "Price") + ":",
                                  style: TextStyle(
                                      fontSize: 13,
                                      fontWeight: FontWeight.bold)),
                              Container(
                                padding: EdgeInsets.symmetric(horizontal: 40),
                                child: Text(
                                  "${widget.product.currentPrice} " +
                                      translate(lang, "Birr")+" / "+ widget.product.getProductUnit().long,
                                  style: TextStyle(fontSize: 13),
                                ),
                              ),
                            ],
                          ),
                        ),
                      ],
                    ),
                    FlatButton(
                      onPressed: () async {
                        StatusAndMessage message;
                        if (widget.subscribed) {
                          message = await context
                              .read<ProductsBloc>()
                              .unSubscribeForProduct(widget.product.id);
                        } else {
                          message = await context
                              .read<ProductsBloc>()
                              .subscribeForProduct(widget.product.id);
                        }
                        if (message.statusCode == 200) {
                          if (widget.subscribed) {
                            context.read<AuthBloc>().add(
                                AuthRemoveSubscription(widget.product.id));
                          } else {
                            context
                                .read<AuthBloc>()
                                .add(AuthAddSubscription(widget.product.id));
                          }
                          widget.subscribed = !widget.subscribed;
                        }
                      },
                      child: Row(
                        mainAxisAlignment : MainAxisAlignment.end, 
                        children: [
                          Text(
                            widget.subscribed
                                ? translate(lang, "Unsubscribe")
                                : translate(lang, "Subscribe"),
                            style: TextStyle(
                              color: widget.subscribed
                                  ? Colors.red
                                  : Colors.green,
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                          // Icon(Icons.subsc)
                        ],
                      ),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
