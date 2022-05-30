import "../../libs.dart";
import "package:flutter/material.dart";
import "package:flutter_bloc/flutter_bloc.dart";

class Products extends StatefulWidget {
  const Products({Key? key}) : super(key: key);

  @override
  State<Products> createState() => _ProductsState();
}

class _ProductsState extends State<Products> {
  @override
  Widget build(BuildContext context) {
    // final state = (context.read<AuthBloc>()).state;
    final products = context.read<ProductsBloc>().state;
    if (!(products is ProductLoadSuccess)) {
      context.read<ProductsBloc>().add(ProductsLoadEvent());
    }
    return Column(
      mainAxisSize: MainAxisSize.max,
      mainAxisAlignment : MainAxisAlignment.center, 
      children: [
        Container(
          child: BlocBuilder<ProductsBloc, ProductState>(
            builder: (ctx, state) {
              return (state is ProductLoadSuccess)
                  ? Column(
                      children: state.products
                          .map<ProductItem>((e) => ProductItem(e))
                          .toList(),
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
    );
  }
}
