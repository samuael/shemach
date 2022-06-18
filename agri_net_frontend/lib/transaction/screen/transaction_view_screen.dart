import '../../libs.dart';

class TransactionDetailScreen extends StatefulWidget {
  static const String RouteName = "/transaction/detail/screen";
  Transaction transaction;
  TransactionDetailScreen(this.transaction, {Key? key}) : super(key: key);

  @override
  State<TransactionDetailScreen> createState() =>
      _TransactionDetailScreenState();
}

class _TransactionDetailScreenState extends State<TransactionDetailScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        elevation: 0,
        title: Text("Transaction Detail"),
        centerTitle: true,
        backgroundColor: Theme.of(context).primaryColor,
        foregroundColor: Colors.white,
      ),
      body: Container(
        padding: EdgeInsets.symmetric(
          vertical: 10,
        ),
        child: Column(
          children: [
            TransactionItem(widget.transaction),
            UserSmallViewItem(
              userid:
                  (this.widget.transaction.requesterId == StaticDataStore.ID)
                      ? this.widget.transaction.sellerId
                      : this.widget.transaction.requesterId,
            ),
          ],
        ),
      ),
    );
  }
}
