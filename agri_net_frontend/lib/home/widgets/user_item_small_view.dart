import "../../libs.dart";
import "package:flutter/material.dart";

class UserSmallViewItem extends StatelessWidget {
  int userid;
  int storeid;
  UserSmallViewItem({this.storeid = 0, this.userid = 0});

  @override
  Widget build(BuildContext context) {
    if (context.watch<UsersBloc>().state is UsersLoadedState) {
      print((context.watch<UsersBloc>().state as UsersLoadedState).merchants);
    } else {
      print("The Users Status is not in the users loaded state ");
    }
    return Container(
      child: ((context.watch<UsersBloc>().state is UsersLoadedState) &&
              (((context.watch<UsersBloc>()).getUserByID(userid) ??
                      (context.watch<UsersBloc>())
                          .getMerchantByStoreID(storeid)) ==
                  null))
          ? Container(
              child: Column(children: [
                Center(
                  child: CircularProgressIndicator(),
                ),
              ]),
            )
          : Container(
              child: Column(children: [
                Text((context.watch<UsersBloc>().getUserByID(userid) ??
                        context
                            .watch<UsersBloc>()
                            .getMerchantByStoreID(storeid))!
                    .firstname),
                Container(
                  padding: EdgeInsets.symmetric(
                    horizontal: 20,
                    vertical: 10,
                  ),
                  margin: EdgeInsets.symmetric(
                    horizontal: 10,
                    vertical: 5,
                  ),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(10),
                    border: Border.all(
                      color: Theme.of(context).primaryColor,
                    ),
                  ),
                  child: Row(
                    children: [
                      Expanded(
                        flex: 1,
                        child: Container(
                          child: ClipRRect(
                            borderRadius: BorderRadius.circular(40),
                            child: (context
                                                .watch<UsersBloc>()
                                                .getUserByID(userid) ??
                                            context
                                                .watch<UsersBloc>()
                                                .getMerchantByStoreID(storeid))!
                                        .imgurl ==
                                    ""
                                ? Image.asset(
                                    'assets/images/logo.jpg',
                                    width: 80,
                                    height: 80,
                                  )
                                : Image.network(
                                    (context
                                                .watch<UsersBloc>()
                                                .getUserByID(userid) ??
                                            context
                                                .watch<UsersBloc>()
                                                .getMerchantByStoreID(storeid))!
                                        .imgurl,
                                  ),
                          ),
                        ),
                      ),
                      Expanded(
                        flex: 4,
                        child: Container(
                          child: Column(
                            children: [
                              Container(
                                child: Row(
                                  children: [
                                    Expanded(
                                      flex: 2,
                                      child: Text(
                                        translate(lang, "Fullname"),
                                        style: TextStyle(
                                          fontWeight: FontWeight.bold,
                                        ),
                                      ),
                                    ),
                                    Expanded(
                                      flex: 1,
                                      child: Text(" : "),
                                    ),
                                    Expanded(
                                      flex: 5,
                                      child: Text(
                                          " ${(context.watch<UsersBloc>().getUserByID(userid) ?? context.watch<UsersBloc>().getMerchantByStoreID(storeid))!.firstname} ${(context.watch<UsersBloc>().getUserByID(userid) ?? context.watch<UsersBloc>().getMerchantByStoreID(storeid))!.lastname} "),
                                    )
                                  ],
                                ),
                              ),
                              Container(
                                child: Row(
                                  children: [
                                    Expanded(
                                      flex: 2,
                                      child: Text(
                                        translate(lang, "Phone"),
                                        style: TextStyle(
                                          fontWeight: FontWeight.bold,
                                        ),
                                      ),
                                    ),
                                    Expanded(
                                      flex: 1,
                                      child: Text(" : "),
                                    ),
                                    Expanded(
                                      flex: 5,
                                      child: Text(
                                          "${(context.watch<UsersBloc>().getUserByID(userid) ?? context.watch<UsersBloc>().getMerchantByStoreID(storeid))!.phone}"),
                                    )
                                  ],
                                ),
                              ),
                              Container(
                                child: Row(
                                  children: [
                                    Expanded(
                                      flex: 2,
                                      child: Text(
                                        translate(lang, "Joined "),
                                        style: TextStyle(
                                          fontWeight: FontWeight.bold,
                                        ),
                                      ),
                                    ),
                                    Expanded(
                                      flex: 1,
                                      child: Text(" : "),
                                    ),
                                    Expanded(
                                      flex: 5,
                                      child: Text(
                                          "${UnixTime((((context.watch<UsersBloc>().getUserByID(userid) ?? context.watch<UsersBloc>().getMerchantByStoreID(storeid))!.createdAt ?? DateTime.now()).microsecondsSinceEpoch / 1000).round()).toString()}  ${translate(lang, " before ")}"),
                                    )
                                  ],
                                ),
                              ),
                            ],
                          ),
                        ),
                      ),
                    ],
                  ),
                )
              ]),
            ),
    );
  }
}
