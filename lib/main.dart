import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:flutter/services.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'reMousable',
      debugShowCheckedModeBanner: false,
      home: const MyHomePage(title: 'Flutter Demo Home Pasdfasdfage'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  final String title;

  const MyHomePage({Key? key, required this.title}) : super(key: key);

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  bool _hoveredSaveButton = false;

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration:
          BoxDecoration(border: Border.all(width: 1, color: Colors.blueAccent)),
      child: Scaffold(
        appBar: DraggebleAppBar(title: "reMousable"),
        backgroundColor: Colors.white,
        body: Container(
          padding: EdgeInsets.all(20),
          child: Align(
            alignment: Alignment.topCenter,
            child: Column(
              children: [
                Wrap(
                  children: [
                    Text(
                      "IP-ADDRESS",
                      style: TextStyle(
                        fontWeight: FontWeight.bold,
                        color: Colors.blueAccent,
                      ),
                    ),
                    SizedBox(
                      height: 25,
                    ),
                    CupertinoTextField(
                      placeholder: "192.168.0.110",
                    ),
                  ],
                ),
                SizedBox(
                  height: 20,
                ),
                Wrap(
                  children: [
                    Text(
                      "PASSWORD",
                      style: TextStyle(
                        fontWeight: FontWeight.bold,
                        color: Colors.blueAccent,
                      ),
                    ),
                    SizedBox(
                      height: 25,
                    ),
                    CupertinoTextField(
                      obscureText: true,
                      placeholder: "•••••••••••",
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
        floatingActionButton: FloatingActionButton(
          onPressed: () {
            setState(() {
              _hoveredSaveButton = true;
            });
          },
          tooltip: 'Update',
          elevation: 0,
          focusElevation: 0,
          hoverElevation: 0,
          disabledElevation: 0,
          highlightElevation: 0,
          splashColor: Colors.transparent,
          backgroundColor: Colors.blue.withAlpha(50),
          child: AnimatedContainer(
            duration: Duration(milliseconds: 350),
            curve: Curves.easeInOutExpo,
            decoration: BoxDecoration(
              borderRadius: BorderRadius.all(Radius.circular(50)),
              boxShadow: [
                BoxShadow(
                  color: Colors.black.withAlpha(40),
                  blurRadius: 10,
                  spreadRadius: 15,
                  offset: Offset(0, 15),
                )
              ],
            ),
            child: const Icon(
              Icons.save,
              color: Colors.blue,
            ),
          ),
        ),
      ),
    );
  }
}

class DraggebleAppBar extends StatelessWidget implements PreferredSizeWidget {
  static const platform_channel_draggable =
      MethodChannel('samples.go-flutter.dev/draggable');

  final String title;
  AppBar? appBar;

  DraggebleAppBar({Key? key, required this.title}) : super(key: key) {
    appBar = AppBar(
      automaticallyImplyLeading: false,
      titleSpacing: 0,
      backgroundColor: Colors.transparent,
      shadowColor: Colors.transparent,
      bottom: PreferredSize(
        child: Container(
          color: Colors.blueAccent,
          height: 1.0,
        ),
        preferredSize: Size.fromHeight(4.0),
      ),
      title: Row(
        children: [
          SizedBox(
            width: 10,
          ),
          Image.asset(
            "assets/icon.png",
          )
        ],
      ),
      actions: <Widget>[
        SizedBox(
          width: 5,
        ),
        IconButton(
          icon: const Icon(
            Icons.close,
            color: Colors.black,
          ),
          onPressed: () async =>
              await platform_channel_draggable.invokeMethod("onClose"),
        ),
        IconButton(
          icon: Text(
            "_",
            style: TextStyle(
              color: Colors.black,
              fontSize: 30,
              height: 0.2,
            ),
          ),
          onPressed: () async =>
              await platform_channel_draggable.invokeMethod("onMinimize"),
        ),
      ].reversed.toList(),
    );
  }

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
        child: MouseRegion(
          child: appBar,
          onHover: onHover,
          onExit: offHover,
        ),
        onPanStart: onPanStart,
        onPanUpdate: onPanUpdate);
  }

  @override
  Size get preferredSize => Size.fromHeight(kToolbarHeight);

  void onPanUpdate(DragUpdateDetails details) async {
    await platform_channel_draggable.invokeMethod('onPanUpdate');
  }

  void onPanStart(DragStartDetails details) async {
    await platform_channel_draggable.invokeMethod('onPanStart',
        {"dx": details.globalPosition.dx, "dy": details.globalPosition.dy});
  }

  void onHover(PointerEvent details) async {
    await platform_channel_draggable.invokeMethod('onHover');
  }

  void offHover(PointerEvent details) async {
    await platform_channel_draggable.invokeMethod('offHover');
  }
}
