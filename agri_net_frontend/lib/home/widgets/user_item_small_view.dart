import "../../libs.dart";
import "package:flutter/material.dart";

class UserSmallViewItem extends StatelessWidget {
  int userid;
  int storeid;
  UserSmallViewItem({this.storeid = 0, this.userid = 0});

  @override
  Widget build(BuildContext context) {
    if (context.watch<UsersBloc>().state is UsersLoadedState) {
      print("\n\n\n\\n\n\n\n\n\n");
      print((context.watch<UsersBloc>().state as UsersLoadedState).merchants);
      print("\n\n\n\\n\n\n\n\n\n");
    }
    return Container(
        child: ((context.watch<UsersBloc>().state is UsersLoadedState) &&
                ((context.watch<UsersBloc>()).getUserByID(userid) ??
                        (context.watch<UsersBloc>())
                            .getMerchantByStoreID(storeid)) ==
                    null)
            ? Center(
                child: Text(
                    "Loaded ${(context.watch<UsersBloc>().state is UsersLoadedState && (context.watch<UsersBloc>()).getUserByID(userid) != null) ? (context.watch<UsersBloc>().getMerchantByStoreID(storeid) ?? context.watch<UsersBloc>().getUserByID(userid))!.firstname : ""}"))
            : Center(
                child: Text(
                    "Loaded ${(context.watch<UsersBloc>().state is UsersLoadedState && (context.watch<UsersBloc>()).getUserByID(userid) != null) ? (context.watch<UsersBloc>().getMerchantByStoreID(storeid) ?? context.watch<UsersBloc>().getUserByID(userid))!.firstname : ""}")));
  }
}
