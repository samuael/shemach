import "../../libs.dart";

class BuyerAcceptTransactionAmendment extends StatefulWidget {
  const BuyerAcceptTransactionAmendment({Key? key}) : super(key: key);

  @override
  State<BuyerAcceptTransactionAmendment> createState() => _BuyerAcceptTransactionAmendmentState();
}

class _BuyerAcceptTransactionAmendmentState extends State<BuyerAcceptTransactionAmendment> {
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
              translate(lang, "Accept Amendment"),
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
