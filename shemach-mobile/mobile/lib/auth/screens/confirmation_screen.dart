import 'dart:async';
import 'package:flutter/services.dart';
import "package:flutter_bloc/flutter_bloc.dart";
import 'package:mobile/auth/auth.dart';
import "../../libs.dart";

class ConfirmationScreen extends StatefulWidget {
  static const String RouteName = "/confirmation/screen";
  final String phone;
  final String fullname;
  bool islogin = false;
  ConfirmationScreen(this.phone, this.fullname,
      {Key? key, this.islogin = false})
      : super(key: key);

  @override
  State<ConfirmationScreen> createState() => _ConfirmationScreenState();
}

class _ConfirmationScreenState extends State<ConfirmationScreen> {
  _ConfirmationScreenState();

  // ---emails controller
  TextEditingController controller1 = TextEditingController();
  TextEditingController controller2 = TextEditingController();
  TextEditingController controller3 = TextEditingController();
  TextEditingController controller4 = TextEditingController();
  TextEditingController controller5 = TextEditingController();

  late DateTime _initialTime;
  late Timer _timer;
  Duration duration = Duration.zero;
  String message = "Confirmation code is sent to your phone number";
  Color messageColor = Colors.black;

  bool onProgress = false;
  @override
  void initState() {
    super.initState();
    _initialTime = DateTime.now();
    _timer = Timer.periodic(Duration(seconds: 1), (timer) {
      final now = DateTime.now();
      setState(() {
        this.duration = now.difference(_initialTime);
      });
      if (duration.inSeconds >= 60 * 3) {
        _initialTime = DateTime.now();
      }
    });
  }

  @override
  void dispose() {
    super.dispose();
    _timer.cancel();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        child: SingleChildScrollView(
          child: Column(
            // mainAxisAlignment : MainAxisAlignment.spaceAround,
            children: [
              Positioned(
                  top: 0,
                  left: 0,
                  child: IconButton(
                      onPressed: () {},
                      icon: Icon(
                        Icons.arrow_left,
                        color: Theme.of(context).primaryColor,
                      ))),
              Center(
                child: Container(
                  height: MediaQuery.of(context).size.width * 0.7,
                  width: MediaQuery.of(context).size.width * 0.7,
                  margin: EdgeInsets.symmetric(
                    vertical: 5,
                  ),
                  child: Center(
                    child: Stack(
                      children: [
                        Container(
                          height: MediaQuery.of(context).size.width * 0.2,
                          width: MediaQuery.of(context).size.width * 0.2,
                          child: Center(
                            child: Text(
                              "${duration.inMinutes}:${duration.inSeconds % 60}",
                              style: TextStyle(
                                fontWeight: FontWeight.bold,
                                fontSize: 20,
                              ),
                            ),
                          ),
                        ),
                        Container(
                          child: CircularProgressIndicator(
                            value: duration.inSeconds / (60 * 3),
                            strokeWidth: 5,
                            // valueColor: Theme.of(context).primaryColor,
                          ),
                          height: MediaQuery.of(context).size.width * 0.2,
                          width: MediaQuery.of(context).size.width * 0.2,
                        ),
                      ],
                    ),
                  ),
                ),
              ),
              Container(
                padding: EdgeInsets.all(5),
                decoration: BoxDecoration(
                  border: Border.all(
                    color: Colors.black26,
                  ),
                  borderRadius: BorderRadius.circular(10),
                ),
                child: Text(
                  translate(lang, message),
                  style: TextStyle(
                    fontStyle: FontStyle.italic,
                    color: messageColor,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
              Container(
                child: Text(
                  "+251-${widget.phone}",
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                    fontStyle: FontStyle.italic,
                  ),
                ),
              ),
              Container(
                  padding: EdgeInsets.symmetric(
                    horizontal: 30,
                    vertical: 20,
                  ),
                  child: Row(children: [
                    Expanded(
                      flex: 1,
                      child: Container(
                        margin: EdgeInsets.symmetric(
                          horizontal: 5,
                        ),
                        decoration: BoxDecoration(
                            border: Border.all(
                          color: Theme.of(context).primaryColor,
                        )),
                        width: 40,
                        height: 40,
                        child: TextField(
                            textAlignVertical: TextAlignVertical.center,
                            textAlign: TextAlign.center,
                            keyboardType: TextInputType.number,
                            maxLength: 5,
                            controller: controller1,
                            onChanged: (val) {
                              setState((){
                              if (val.length >= 1) {
                                controller1.text = val[0];
                                if (val.length >= 2) {
                                  controller2.text = val[1];
                                  if (val.length >= 3) {
                                    controller3.text = val[2];
                                    if (val.length >= 4) {
                                      controller4.text = val[3];
                                      if (val.length >= 5) {
                                        controller5.text = val[4];
                                      }
                                    }
                                  }
                                }
                              }});
                            },
                            cursorHeight: 25,
                            cursorWidth: 3,
                            // maxLengthEnforced: true,
                            decoration: InputDecoration(
                              counter: Offstage(),
                            ),
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                              fontSize: 25,
                            )),
                      ),
                    ),
                    Expanded(
                      flex: 1,
                      child: Container(
                        margin: EdgeInsets.symmetric(
                          horizontal: 5,
                        ),
                        decoration: BoxDecoration(
                            border: Border.all(
                          color: Theme.of(context).primaryColor,
                        )),
                        width: 40,
                        height: 40,
                        child: TextField(
                            inputFormatters: [
                              // WhitelistingTextInputFormatter.digitsOnly,/
                              LengthLimitingTextInputFormatter(1),
                            ],
                            decoration: InputDecoration(
                              counter: Offstage(),
                            ),
                            keyboardType: TextInputType.number,
                            textAlignVertical: TextAlignVertical.center,
                            textAlign: TextAlign.center,
                            maxLength: 1,
                            controller: controller2,
                            cursorHeight: 25,
                            cursorWidth: 3,
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                              fontSize: 25,
                            )),
                      ),
                    ),
                    Expanded(
                      flex: 1,
                      child: Container(
                        margin: EdgeInsets.symmetric(
                          horizontal: 5,
                        ),
                        decoration: BoxDecoration(
                            border: Border.all(
                          color: Theme.of(context).primaryColor,
                        )),
                        width: 40,
                        height: 40,
                        child: TextField(
                            inputFormatters: [
                              // WhitelistingTextInputFormatter.digitsOnly,
                              LengthLimitingTextInputFormatter(1),
                            ],
                            decoration: InputDecoration(
                              counter: Offstage(),
                            ),
                            keyboardType: TextInputType.number,
                            textAlignVertical: TextAlignVertical.center,
                            textAlign: TextAlign.center,
                            maxLength: 1,
                            controller: controller3,
                            cursorHeight: 25,
                            cursorWidth: 3,
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                              fontSize: 25,
                            )),
                      ),
                    ),
                    Expanded(
                      flex: 1,
                      child: Container(
                        margin: EdgeInsets.symmetric(
                          horizontal: 5,
                        ),
                        decoration: BoxDecoration(
                            border: Border.all(
                          color: Theme.of(context).primaryColor,
                        )),
                        width: 40,
                        height: 40,
                        child: TextField(
                            inputFormatters: [
                              // WhitelistingTextInputFormatter.digitsOnly,
                              LengthLimitingTextInputFormatter(1),
                            ],
                            decoration: InputDecoration(
                              counter: Offstage(),
                            ),
                            keyboardType: TextInputType.number,
                            textAlignVertical: TextAlignVertical.center,
                            textAlign: TextAlign.center,
                            maxLength: 1,
                            controller: controller4,
                            cursorHeight: 25,
                            cursorWidth: 3,
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                              fontSize: 25,
                            )),
                      ),
                    ),
                    Expanded(
                      flex: 1,
                      child: Container(
                        margin: EdgeInsets.symmetric(
                          horizontal: 5,
                        ),
                        decoration: BoxDecoration(
                            border: Border.all(
                          color: Theme.of(context).primaryColor,
                        )),
                        width: 40,
                        height: 40,
                        child: TextField(
                            inputFormatters: [
                              // WhitelistingTextInputFormatter.digitsOnly,
                              LengthLimitingTextInputFormatter(1),
                            ],
                            decoration: InputDecoration(
                                counter: Offstage(), counterText: ""),
                            textAlignVertical: TextAlignVertical.center,
                            textAlign: TextAlign.center,
                            keyboardType: TextInputType.number,
                            maxLength: 1,
                            controller: controller5,
                            cursorHeight: 25,
                            cursorWidth: 3,
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                              fontSize: 25,
                            )),
                      ),
                    ),
                  ])),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Container(
                    margin: EdgeInsets.symmetric(vertical: 10, horizontal: 10),
                    child: Text(
                      translate(lang, "Don't recieve the code?"),
                      style: TextStyle(
                        fontStyle: FontStyle.italic,
                        fontSize: 14,
                      ),
                    ),
                  ),
                ],
              ),
              GestureDetector(
                child: ClipRRect(
                  borderRadius: BorderRadius.circular(20),
                  child: Container(
                    color: Theme.of(context).primaryColor,
                    padding:
                        EdgeInsets.symmetric(horizontal: 105, vertical: 10),
                    child: onProgress
                        ? CircularProgressIndicator(
                            color: Colors.white,
                            strokeWidth: 6,
                          )
                        : Text(
                            translate(lang, " Confirm "),
                            style: TextStyle(
                              fontStyle: FontStyle.italic,
                              color: Colors.white,
                              fontWeight: FontWeight.bold,
                              fontSize: 24,
                            ),
                          ),
                  ),
                ),
                onTap: () async {
                  if (onProgress) {
                    return;
                  }
                  try {
                    if (int.parse(controller1.text) >= 0 &&
                        int.parse(controller2.text) >= 0 &&
                        int.parse(controller3.text) >= 0 &&
                        int.parse(controller4.text) >= 0 &&
                        int.parse(controller5.text) >= 0) {
                      final String confirmation =
                          "${int.parse(controller1.text)}${int.parse(controller2.text)}${int.parse(controller3.text)}${int.parse(controller4.text)}${int.parse(controller5.text)}";
                      final phone = "+251${this.widget.phone}";
                      setState(() {
                        onProgress = true;
                      });
                      SubscriberAuthenticationRespnse result;
                      if (widget.islogin) {
                        result = await context.read<AuthBloc>().confirmLogin(
                            SubscriberConfirmation(phone, confirmation));
                      } else {
                        result = await context
                            .read<AuthBloc>()
                            .confirmRegistration(
                                SubscriberConfirmation(phone, confirmation));
                      }
                      if (result.statusCode == 200) {
                        setState(() {
                          this.message = result.msg;
                          messageColor = Colors.green;
                          onProgress = true;
                        });
                        setState(() {
                          onProgress = false;
                        });
                        context.read<AuthBloc>().add(
                            AuthSubscriberAuthenticatedEvent(
                                subscriber: result.subscriber!,
                                token: result.token));
                        Navigator.of(context).pushNamedAndRemoveUntil(
                            HomeScreen.RouteName, (route) => false);

                        setState(() {
                          onProgress = false;
                        });
                      } else {
                        setState(() {
                          this.message = result.msg;
                          messageColor = Colors.red;
                          onProgress = false;
                        });
                      }
                    }
                  } catch (e, a) {
                    setState(() {
                      onProgress = false;
                    });
                  }
                },
              ),
            ],
          ),
        ),
      ),
    );
  }
}
