class TransactionInput {
  double price;
  int quantity;
  String description;
  int productid;
  int? storeid;

  TransactionInput(this.price, this.quantity, this.description, this.productid,
      this.storeid);

  Map<String, dynamic> toJson() {
    return {
      "price": this.price,
      "qty": this.quantity,
      "description": this.description,
      "product_id": this.productid,
      "requester_store_ref": this.storeid
    };
  }
}
