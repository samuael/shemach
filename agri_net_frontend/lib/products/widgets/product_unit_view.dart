import 'package:flutter/src/widgets/framework.dart';

class ProductTypeItemSmallView extends StatelessWidget {
  const ProductTypeItemSmallView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      color : Theme.of(context).primaryColor, 
      width : 40, height:40, 
      child : Text("Some --------")
    );
  }
}