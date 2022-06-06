import '../../libs.dart';

class ProductForm extends StatefulWidget {
  static String RouteName = "products/post";
  const ProductForm({Key? key}) : super(key: key);

  @override
  State<ProductForm> createState() => _ProductFormState();
}

class _ProductFormState extends State<ProductForm> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: [
          Container(
            child: TextField(),
          )
        ],
      ),
    );
  }
}
