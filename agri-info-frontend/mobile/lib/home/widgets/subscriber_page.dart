import 'package:flutter/material.dart';
import "package:flutter_bloc/flutter_bloc.dart";
import '../../libs.dart';

class SubscriptionPage extends StatefulWidget {
  static const String RouteName = "/auth_screen2";

  SubscriptionPage({Key? key}) : super(key: key);

  @override
  State createState() {
    return SubscriptionPageState();
  }
}

class SubscriptionPageState extends State<SubscriptionPage> {
  bool right = false;
  bool searching = false;
  double opacity = 0.7;

  TextEditingController controller = TextEditingController();

  @override
  Widget build(BuildContext context) {
    final productsProvider = BlocProvider.of<ProductsBloc>(context);
    return SingleChildScrollView(
        child: Column(
          children: [
            SingleChildScrollView(
              child: Column(
                children: [
                  Container(
                    child: BlocBuilder<ProductsBloc, ProductState>(
                      builder: (ctx, state) {
                        return (state is ProductLoadSuccess)
                            ? Column(
                                children: state.products.map<Widget>((e) {
                                  bool subscribed = false;
                                  final subscriptions =
                                      (context.watch<AuthBloc>().state
                                              as AuthSubscriberAuthenticated)
                                          .subscriber
                                          .subscriptions;
                                  for (final s in subscriptions) {
                                    if (s == e.id) {
                                      subscribed = true;
                                    }
                                  }
                                  return SubscriptionProductItem(
                                      e, subscribed);
                                }).toList(),
                              )
                            : Center(
                                child: Column(
                                  mainAxisAlignment: MainAxisAlignment.center,
                                  children: [
                                    Text(
                                      (state is ProductLoadFailure)
                                          ? state.response.msg
                                          : "No product available",
                                      style: TextStyle(
                                        fontWeight: FontWeight.bold,
                                        fontStyle: FontStyle.italic,
                                      ),
                                    ),
                                    IconButton(
                                      icon: Icon(Icons.refresh_sharp),
                                      onPressed: () {
                                        context
                                            .read<ProductsBloc>()
                                            .add(ProductsLoadEvent());
                                      },
                                    )
                                  ],
                                ),
                              );
                      },
                    ),
                  ),
                ],
              ),
            ),
          ],
        ),
      );
  }
}
