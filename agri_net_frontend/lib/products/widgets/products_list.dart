import "package:flutter/material.dart";


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
      child : Center(
        child : Text("Center"), 
      )

    );
  }
}