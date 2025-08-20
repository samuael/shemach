import 'package:flutter/services.dart';
import "package:flutter_bloc/flutter_bloc.dart";
import "../../libs.dart";

class RegistrationScreen extends StatefulWidget {
  static const String RouteName = "/subscriber/registration/page";

  const RegistrationScreen({Key? key}) : super(key: key);

  @override
  State<RegistrationScreen> createState() => _RegistrationScreenState();
}

class _RegistrationScreenState extends State<RegistrationScreen> {
  TextEditingController phoneController = TextEditingController();
  TextEditingController fullnameController = TextEditingController();

  bool roleFarmer = false;
  bool roleMerchant = true;
  bool roleConsumer = false;
  bool roleAll = false;

  bool onProgress = false;

  int groupValue = 1;

  Color messageColor = Colors.green;
  String message = "";

  @override
  Widget build(BuildContext context) {
    lang = groupValue == 1
        ? "amh"
        : (groupValue == 2 ? "oro" : (groupValue == 3 ? "tig" : "eng"));
    return Scaffold(
      appBar: AppBar(
        elevation: 0,
        centerTitle: true,
        title: Text(
          translate(lang, " Registration "),
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
            ClipRRect(
              borderRadius: BorderRadius.circular(30),
              child: Container(
                padding: EdgeInsets.symmetric(
                  horizontal: 20,
                  vertical: 10,
                ),
                margin: EdgeInsets.symmetric(
                  horizontal: 20,
                  vertical: 10,
                ),
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(
                    30,
                  ),
                  // border: Border()
                ),
                child: onProgress
                    ? LinearProgressIndicator(
                        color: Theme.of(context).primaryColor,
                        minHeight: 5,
                      )
                    : Text(
                        message,
                        style: TextStyle(
                          color: messageColor,
                          fontWeight: FontWeight.bold,
                          fontFamily: "Elegant TypeWriter",
                        ),
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
                cursorColor: Theme.of(context).primaryColorLight,
                controller: fullnameController,
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                ),
                decoration: InputDecoration(
                  labelText: translate(lang, "Full Name"),
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
                // inputFormatters: [WhitelistingTextInputFormatter.digitsOnly],
                keyboardType: TextInputType.number,
                cursorColor: Theme.of(context).primaryColorLight,
                controller: phoneController,
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                ),
                decoration: InputDecoration(
                  labelText: translate(lang, "Phone"),
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
                translate(lang, " Job Title "),
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
                            Text(translate(lang, "Farmer")),
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
                              Text(translate(lang, "Merchant")),
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
                            Text(translate(lang, "Consumer")),
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
                              Text(translate(lang, "All")),
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
                                translate(lang, "Langauge"),
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
              margin: EdgeInsets.only(
                top: 10,
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
                  if (onProgress) {
                    return;
                  }
                  lang = groupValue == 1
                      ? "amh"
                      : (groupValue == 2
                          ? "oro"
                          : (groupValue == 3 ? "tig" : "eng"));
                  if (fullnameController.text != "" &&
                      (phoneController.text != "" &&
                          phoneController.text.length == 9) &&
                      (roleFarmer || roleMerchant || roleConsumer || roleAll)) {
                    setState(() {
                      onProgress = true;
                    });

                    final registrationResult =
                        await context.read<AuthBloc>().register(
                              RegistrationInput(
                                fullname: fullnameController.text,
                                phone: "+251${phoneController.text.trim()}",
                                role: 1,
                                lang: lang,
                              ),
                            );
                    if (registrationResult.statusCode == 200 ||
                        registrationResult.statusCode == 201) {
                      setState(() {
                        message = registrationResult.msg;
                        messageColor = Colors.green;
                        onProgress = false;
                      });
                      Navigator.of(context)
                          .pushNamed(ConfirmationScreen.RouteName, arguments: {
                        "fullname": fullnameController.text,
                        "phone": phoneController.text,
                        "islogin": false,
                      });
                    } else {
                      setState(() {
                        message = registrationResult.msg;
                        messageColor = Colors.red;
                      });
                    }
                  } else if (fullnameController.text == "") {
                    setState(() {
                      message = translate(lang, "please enter your full name");
                      messageColor = Colors.red;
                    });
                  } else if (phoneController.text == "") {
                    setState(() {
                      message =
                          translate(lang, "please enter your phone number");
                      messageColor = Colors.red;
                    });
                  } else {
                    setState(() {
                      message = translate(lang, "please select a role");
                      messageColor = Colors.red;
                    });
                  }
                },
                icon: Icon(Icons.app_registration),
                label: Text(
                  translate(lang, " Register "),
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
            ),
            ElevatedButton(
              onPressed: () {
                Navigator.of(context)
                    .pushNamed(ConfirmationScreen.RouteName, arguments: {
                  "fullname": fullnameController.text,
                  "phone": phoneController.text,
                  "islogin": false,
                });
              },
              child: Text(
                "Confirmation",
                style: TextStyle(
                  color: Colors.blue,
                  fontWeight: FontWeight.bold,
                  fontStyle: FontStyle.italic,
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
