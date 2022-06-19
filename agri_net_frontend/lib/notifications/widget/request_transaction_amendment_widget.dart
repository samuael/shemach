import "../../libs.dart";

class RequestAmendment extends StatefulWidget {
  const RequestAmendment({Key? key}) : super(key: key);

  @override
  State<RequestAmendment> createState() => _RequestAmendmentState();
}

class _RequestAmendmentState extends State<RequestAmendment> {
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
              translate(lang, "Request Transaction Amendment"),
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
