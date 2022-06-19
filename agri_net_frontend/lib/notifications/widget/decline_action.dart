import "../../libs.dart";

class DeclineTransaction extends StatefulWidget {
  const DeclineTransaction({Key? key}) : super(key: key);

  @override
  State<DeclineTransaction> createState() => _DeclineTransactionState();
}

class _DeclineTransactionState extends State<DeclineTransaction> {
  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {},
      child: ClipRRect(
        borderRadius: BorderRadius.circular(
          5,
        ),
        child: Container(
          decoration: BoxDecoration(
            color: Colors.red,
          ),
          padding: EdgeInsets.symmetric(
            vertical: 10,
            horizontal: 10,
          ),
          child: Text(
            translate(lang, "Decline"),
            style: TextStyle(
              fontWeight: FontWeight.bold,
              color: Colors.white,
            ),
          ),
        ),
      ),
    );
  }
}
