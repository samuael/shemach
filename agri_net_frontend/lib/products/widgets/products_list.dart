import "../../libs.dart";

class ProductsList extends StatefulWidget {
  @override
  State<ProductsList> createState() {
    return ProductsListState();
  }
}

class ProductsListState extends State<ProductsList> {
  @override
  Widget build(BuildContext context) {
    return Container(
      child: Column(children: [
        Container(
          margin: EdgeInsets.symmetric(
            horizontal: 20,
            vertical: 20,
          ),
          padding: EdgeInsets.symmetric(
            horizontal: 25,
          ),
          decoration: BoxDecoration(
            border: Border.all(color: Theme.of(context).primaryColor),
            borderRadius: BorderRadius.circular(10),
          ),
          child: TextField(
            onChanged: (t) {
              setState(() {
                // widget.text = t;
              });
            },
            autofocus: true,
            style: TextStyle(
              fontWeight: FontWeight.bold,
            ),
            // controller: searchController,
            decoration: InputDecoration(
              border: InputBorder.none,
              suffixIcon: Icon(
                Icons.search,
                color: Theme.of(context).primaryColor,
              ),
            ),
          ),
        ),
        // --------------------------------------------
        BlocBuilder<MyProductsBloc, ProductState>(builder: (context, state) {
          if (state is MyProductsLoadSuccess) {
            return Column(
              children: state.posts
                  .map<ListTile>(
                    (e) => ListTile(
                      title: Text(
                        e.description,
                        style: TextStyle(
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                    ),
                  )
                  .toList(),
            );
          } else {
            return Center(
              child: Text("No Product Instance found"),
            );
          }
        }),
      ]),
    );
  }
}
