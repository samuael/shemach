import "../../libs.dart";

class ProductPostsList extends StatefulWidget {
  @override
  State<ProductPostsList> createState() {
    return ProductPostsListState();
  }
}

class ProductPostsListState extends State<ProductPostsList> {
  @override
  Widget build(BuildContext context) {
    return Container(
      child: SingleChildScrollView(
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
          BlocBuilder<ProductsBloc, ProductState>(builder: (context, state) {
            if (state is ProductsLoadSuccess) {
              return Column(
                children: state.posts
                    .map<ProductPostItem>((e) => ProductPostItem(post: e))
                    .toList(),
              );
            } else {
              return Center(
                child: Column(
                  children: [
                    Text("No Product Instance found"),
                    IconButton(
                      icon: Icon(
                        Icons.airline_seat_recline_normal_rounded,
                        color: Colors.blue,
                      ),
                      onPressed: () {
                        context
                            .read<ProductTypeBloc>()
                            .add(ProductTypesLoadEvent());
                      },
                    )
                  ],
                ),
              );
            }
          }),
        ]),
      ),
    );
  }
}
