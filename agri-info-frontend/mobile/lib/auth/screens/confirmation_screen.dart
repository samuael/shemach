import 'dart:async';

import 'package:flutter/scheduler.dart';

import "../../libs.dart";

class ConfirmationScreen extends StatefulWidget {
  static const String RouteName = "/confirmation/screen";
  final String phone;
  const ConfirmationScreen(this.phone, {Key? key}) : super(key: key);

  @override
  State<ConfirmationScreen> createState() => _ConfirmationScreenState();
}

class _ConfirmationScreenState extends State<ConfirmationScreen> {
  _ConfirmationScreenState();

  late DateTime _initialTime;
  late Timer _timer;
  Duration duration = Duration.zero;
  @override
  void initState() {
    super.initState();
    _initialTime = DateTime.now();
    _timer = Timer.periodic(Duration(seconds: 1), (timer) {
      final now = DateTime.now();
      setState(() {
        this.duration = now.difference(_initialTime);
      });
      if (duration.inSeconds >= 180) {
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
        child: Column(
          // mainAxisAlignment : MainAxisAlignment.spaceAround,
          children: [
            Center(
              child: Container(
                height: MediaQuery.of(context).size.width * 0.7,
                width: MediaQuery.of(context).size.width * 0.7,
                margin: EdgeInsets.symmetric(
                  vertical: 10,
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
                          value: duration.inSeconds / 180,
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
              child: Text(
                "Confirmation code is sent to your phone number",
                style: TextStyle(
                  fontStyle: FontStyle.italic,
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
              padding : EdgeInsets.symmetric(horizontal:30, vertical:20,),
              child: Row(
                children : [
                  Expanded(flex: 1, child:Container(margin:EdgeInsets.symmetric(horizontal:5,),decoration:BoxDecoration(border:Border.all(color:Theme.of(context).primaryColor, )  ), width:40, height:40, child:TextField(cursorHeight: 25,cursorWidth: 3,style:TextStyle(fontWeight:FontWeight.bold,fontSize: 25,)),),),
                  Expanded(flex: 1, child:Container(margin:EdgeInsets.symmetric(horizontal:5,),decoration:BoxDecoration(border:Border.all(color:Theme.of(context).primaryColor, )  ), width:40, height:40, child:TextField(cursorHeight: 25,cursorWidth: 3,style:TextStyle(fontWeight:FontWeight.bold,fontSize: 25,)),),),
                  Expanded(flex: 1, child:Container(margin:EdgeInsets.symmetric(horizontal:5,),decoration:BoxDecoration(border:Border.all(color:Theme.of(context).primaryColor, )  ), width:40, height:40, child:TextField(cursorHeight: 25,cursorWidth: 3,style:TextStyle(fontWeight:FontWeight.bold,fontSize: 25,)),),),
                  Expanded(flex: 1, child:Container(margin:EdgeInsets.symmetric(horizontal:5,),decoration:BoxDecoration(border:Border.all(color:Theme.of(context).primaryColor, )  ), width:40, height:40, child:TextField(cursorHeight: 25,cursorWidth: 3,style:TextStyle(fontWeight:FontWeight.bold,fontSize: 25,)),),),
                  Expanded(flex: 1, child:Container(margin:EdgeInsets.symmetric(horizontal:5,),decoration:BoxDecoration(border:Border.all(color:Theme.of(context).primaryColor, )  ), width:40, height:40, child:TextField(cursorHeight: 25,cursorWidth: 3,style:TextStyle(fontWeight:FontWeight.bold,fontSize: 25,)),),),
                ]
              )
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Container(
                  margin: EdgeInsets.symmetric(vertical: 10, horizontal : 10),
                  child: Text(
                    "Don't recieve the code?",
                    style: TextStyle(
                      fontStyle: FontStyle.italic,
                      fontSize: 14,
                    ),
                  ),
                ),
                GestureDetector(
                  child: Container(
                    child: Text(
                      "RESEND",
                      style: TextStyle(
                        fontStyle: FontStyle.italic,
                        color: Colors.blue,
                        fontWeight: FontWeight.bold,
                        fontSize: 14,
                      ),
                    ),
                  ),
                ),
              ],
            ),
            ClipRRect(
              borderRadius: BorderRadius.circular(20),
              child: Container(
                color: Theme.of(context).primaryColor,
                padding: EdgeInsets.symmetric(horizontal: 105, vertical: 10),
                child: Text(
                  " Confirm ",
                  style: TextStyle(
                    fontStyle: FontStyle.italic,
                    color: Colors.white,
                    fontWeight: FontWeight.bold,
                    fontSize: 24,
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
