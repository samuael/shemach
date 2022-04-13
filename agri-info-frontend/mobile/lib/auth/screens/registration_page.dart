import 'package:flutter/services.dart';

import "../../libs.dart";

class RegistrationScreen extends StatefulWidget {
  static const String RouteName = "/subscriber/registration/page";

  const RegistrationScreen({Key? key}) : super(key: key);

  @override
  State<RegistrationScreen> createState() => _RegistrationScreenState();
}

class _RegistrationScreenState extends State<RegistrationScreen> {
  TextEditingController phoneController = TextEditingController();

  bool roleFarmer = true;
  bool roleMerchant = false;
  bool roleConsumer = false;
  bool roleAll = false;

  int groupValue = 1;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        elevation: 0,
        centerTitle: true,
        title: Text(
          " Registration ",
          style: TextStyle(
            fontWeight: FontWeight.bold,
            color: Colors.white,
          ),
        ),
      ),
      body: SingleChildScrollView(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Container(
              padding: EdgeInsets.all(5),
              child: Text(
                " To Register, please provide your information ",
                style: TextStyle(
                  color: Theme.of(context).primaryColor,
                  fontWeight: FontWeight.bold,
                  fontFamily: "Moms TypeWriter",
                  fontSize: 14,
                  // fontSize: 18
                ),
              ),
            ),
            Container(
              padding: EdgeInsets.symmetric(
                horizontal: 20,
                vertical: 10,
              ),
              child: Text(
                "  ",
                style: TextStyle(
                  color: Colors.green,
                  fontWeight: FontWeight.bold,
                  fontFamily: "Elegant TypeWriter",
                  // fontSize: 18
                ),
              ),
            ),
            Container(
              margin: EdgeInsets.symmetric(
                vertical: 20,
              ),
              padding: EdgeInsets.symmetric(
                horizontal: 20,
              ),
              child: TextField(
                inputFormatters: [WhitelistingTextInputFormatter.digitsOnly],
                keyboardType: TextInputType.number,
                cursorColor: Theme.of(context).primaryColorLight,
                controller: phoneController,
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                ),
                decoration: InputDecoration(
                  labelText: "Full Name",
                  fillColor: Colors.lightBlue,
                  hoverColor: Colors.lightBlue,
                  suffixIcon: Icon(Icons.person),
                  border: OutlineInputBorder(
                    borderSide: BorderSide(
                      color: Colors.lightBlue,
                      style: BorderStyle.none,
                    ),
                  ),
                ),
                onChanged: (text) {
                  if (text.length > 9) {
                    phoneController.text = text.substring(1, 9);
                  }
                },
              ),
            ),
            Container(
              padding: EdgeInsets.symmetric(
                horizontal: 20,
              ),
              child: TextField(
                inputFormatters: [WhitelistingTextInputFormatter.digitsOnly],
                keyboardType: TextInputType.number,
                cursorColor: Theme.of(context).primaryColorLight,
                controller: phoneController,
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                ),
                decoration: InputDecoration(
                  labelText: "Phone",
                  fillColor: Colors.lightBlue,
                  hoverColor: Colors.lightBlue,
                  prefix: Container(
                    padding: EdgeInsets.symmetric(horizontal: 5),
                    child: Text(
                      "+251",
                      style: TextStyle(
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                  ),
                  suffixIcon: Icon(Icons.phone),
                  border: OutlineInputBorder(
                    borderSide: BorderSide(
                      color: Colors.lightBlue,
                      style: BorderStyle.none,
                    ),
                  ),
                ),
                onChanged: (text) {
                  if (text.length > 9) {
                    phoneController.text = text.substring(1, 9);
                  }
                },
              ),
            ),
            Container(
              padding: EdgeInsets.all(5),
              margin: EdgeInsets.only(
                left: 20,
              ),
              alignment: AlignmentGeometry.lerp(
                  Alignment.topLeft, Alignment.center, 0),
              child: Text(
                " Job Title ",
                style: TextStyle(
                  color: Theme.of(context).primaryColor,
                  fontWeight: FontWeight.bold,
                  fontFamily: "Moms TypeWriter",
                  fontSize: 17,
                  // fontSize: 18
                ),
              ),
            ),
            Container(
              padding: EdgeInsets.symmetric(
                horizontal: 40,
              ),
              child: Row(
                children: [
                  Expanded(
                    flex: 1,
                    child: Container(
                      child: Column(
                        children: [
                          Row(children: [
                            Checkbox(
                                value: roleFarmer,
                                onChanged: (values) {
                                  setState(() {
                                    this.roleFarmer = values!;
                                    if (!this.roleFarmer) {
                                      this.roleAll = false;
                                    }
                                    checkAllRolesIfApplicable();
                                  });
                                }),
                            Text("Farmer"),
                          ]),
                          Row(
                            children: [
                              Checkbox(
                                  value: roleMerchant,
                                  onChanged: (values) {
                                    setState(() {
                                      this.roleMerchant = values!;
                                      if (!this.roleMerchant) {
                                        this.roleAll = false;
                                      }
                                      checkAllRolesIfApplicable();
                                    });
                                  }),
                              Text("Merchant"),
                            ],
                          ),
                          Row(children: [
                            Checkbox(
                                value: roleConsumer,
                                onChanged: (values) {
                                  setState(() {
                                    this.roleConsumer = values!;
                                    if (!this.roleConsumer) {
                                      this.roleAll = false;
                                    }
                                    checkAllRolesIfApplicable();
                                  });
                                }),
                            Text("Consumer"),
                          ]),
                          Row(
                            children: [
                              Checkbox(
                                  value: this.roleAll,
                                  onChanged: (values) {
                                    setState(() {
                                      this.roleAll = values!;
                                      if (roleAll) {
                                        this.roleConsumer = roleAll;
                                        this.roleFarmer = roleAll;
                                        this.roleMerchant = roleAll;
                                      }
                                    });
                                  }),
                              Text("All"),
                            ],
                          ),
                        ],
                      ),
                    ),
                  ),
                  Expanded(
                    flex: 1,
                    child: Container(
                      child: Column(
                        mainAxisAlignment: MainAxisAlignment.end,
                        children: [
                          Row(
                            children: [
                              Text(
                                "Langauge",
                                style: TextStyle(
                                  fontWeight: FontWeight.bold,
                                  fontStyle: FontStyle.italic,
                                ),
                              ),
                            ],
                          ),
                          Row(
                            children: [
                              Radio(
                                  value: 1,
                                  groupValue: groupValue,
                                  onChanged: (val) {
                                    setState(() {
                                      this.groupValue = 1;
                                    });
                                  }),
                              Text("አማርኛ")
                            ],
                          ),
                          Row(
                            children: [
                              Radio(
                                  value: 2,
                                  groupValue: groupValue,
                                  onChanged: (val) {
                                    setState(() {
                                      this.groupValue = 2;
                                    });
                                  }),
                              Text("Oromiffa")
                            ],
                          ),
                          Row(
                            children: [
                              Radio(
                                  value: 3,
                                  groupValue: groupValue,
                                  onChanged: (val) {
                                    setState(() {
                                      this.groupValue = 3;
                                    });
                                  }),
                              Text("ትግርኛ")
                            ],
                          ),
                          Row(
                            children: [
                              Radio(
                                  value: 4,
                                  groupValue: groupValue,
                                  onChanged: (val) {
                                    setState(() {
                                      this.groupValue = 4;
                                    });
                                  }),
                              Text("English")
                            ],
                          ),
                        ],
                      ),
                    ),
                  ),
                ],
              ),
            ),
            Container(
              margin: EdgeInsets.symmetric(
                vertical: 10,
              ),
              padding: EdgeInsets.symmetric(
                horizontal: 30,
                vertical: 10,
              ),
              child: ElevatedButton.icon(
                style: ButtonStyle(
                  padding: MaterialStateProperty.all<EdgeInsets>(
                    EdgeInsets.symmetric(
                      vertical: 10,
                      horizontal: 40,
                    ),
                  ),
                ),
                onPressed: () async {
                  // checking the validity of input values
                  Navigator.of(context).pushNamed(
                    ConfirmationScreen.RouteName,
                    arguments: {
                      "phone": phoneController.text,
                    },
                  );
                },
                icon: Icon(Icons.app_registration),
                label: Text(
                  " Register ",
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

  checkAllRolesIfApplicable() {
    if (this.roleFarmer && this.roleMerchant && this.roleConsumer) {
      setState(() {
        this.roleAll = true;
      });
    }
  }
}
