import "../../libs.dart";

class AmendTransaction extends StatefulWidget {
  const AmendTransaction({Key? key}) : super(key: key);

  @override
  State<AmendTransaction> createState() => _AmendTransactionState();
}

class _AmendTransactionState extends State<AmendTransaction> {
  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {},
      child: Padding(
        padding: EdgeInsets.symmetric(
              vertical: 10,
              horizontal: 10,
            ),
        child: ClipRRect(
          borderRadius: BorderRadius.circular(
            5,
          ),
          child: Container(
            decoration: BoxDecoration(
              color: Theme.of(context).primaryColor,
            ),
            padding: EdgeInsets.symmetric(
              vertical: 10,
              horizontal: 10,
            ),
            child: Text(
              translate(lang, "Amend Transaction"),
              style: TextStyle(
                fontWeight: FontWeight.bold,
                color: Colors.white,
              ),
            ),
          ),
        ),
      ),
    );
  }
}
