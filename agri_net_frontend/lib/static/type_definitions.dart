import "package:flutter/material.dart";

enum AdminActionType {
  registration,
  approval,
}

enum PaymentStatus {
  PaymentStatusNotSent,
  PaymentStatusSent,
  PaymentStatusSeen,
  PaymentStatusAccepted,
  PaymentStatusRejected,
}

List<String> paymentStatusMessages = [
  "Created",
  "Sent",
  "Seen",
  "Accepted",
  "Rejected",
];

List<Color>  paymentStatusColors = [
    Colors.grey,
    Colors.yellow,
    Colors.blue,
    Colors.green,
    Colors.red,
];

List<IconData>  paymentStatusIcons = [
    Icons.create,
    Icons.send,
    Icons.double_arrow,
    Icons.check,
    Icons.wrong_location_outlined,
];