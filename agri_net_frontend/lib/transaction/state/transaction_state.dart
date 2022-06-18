import "../../libs.dart";

class TransactionState {}

class TransactionLoading extends TransactionState {}

class TransactionInit extends TransactionState {}

class TransactionsLoaded extends TransactionState {
  List<Transaction> transactions;
  TransactionsLoaded(
    this.transactions,
  );
}

class TransactionsLoadingFailed extends TransactionState {
  int statusCode;
  String? msg;

  TransactionsLoadingFailed(this.statusCode, this.msg);
}
