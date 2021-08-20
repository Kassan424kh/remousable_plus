import 'dart:async';

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

  late TextEditingController _ipAddressController;
  late TextEditingController _passwordController;
  late bool _showSplashscreen = true;
  late bool _removeSplashscreen = false;

  static const platform_channel_draggable =
      MethodChannel('samples.go-flutter.dev/draggable');
  @override
  void initState() {
    _ipAddressController = TextEditingController(text: "192.168.0.1");
    _passwordController = TextEditingController(text: "Hello ");

    super.initState();
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();

    new Timer(const Duration(seconds: 3), () {
      setState(() {
        _showSplashscreen = false;
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    double _timeDilation = 5.0; // 1.0 means normal animation speed.
    Size _size = MediaQuery.of(context).size;

    return Container(
      decoration:
          BoxDecoration(border: Border.all(width: 1, color: Colors.blueAccent)),
      height: _size.height,
      width: _size.width,
      child: Stack(
        children: [
          Scaffold(
              appBar: DraggebleWidget(
                widget: AppBar(
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
                      Hero(
                        tag: "icon",
                        child: Image.asset(
                          "assets/icon.png",
                        ),
                      ),
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
                      onPressed: () async => await platform_channel_draggable
                          .invokeMethod("onClose"),
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
                      onPressed: () async => await platform_channel_draggable
                          .invokeMethod("onMinimize"),
                    ),
                  ].reversed.toList(),
                ),
              ),
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
                            controller: _ipAddressController,
                            placeholder: "192.168.0.110",
                          ),
                        ],
                      ),
                      SizedBox(
                        height: 15,
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
                            controller: _passwordController,
                            obscureText: true,
                            placeholder: "•••••••••••",
                            onChanged: (value) {
                              print(value);

                              _passwordController.text =
                                  value.split("").toList().reversed.join();
                            },
                          ),
                        ],
                      ),
                      SizedBox(
                        height: 15,
                      ),
                      Text(
                        "In your reMarkable tablet got to:\nSettings > About > Copyrights and licenses > General information (scroll down)",
                        style: TextStyle(
                          fontSize: 13,
                          color: Colors.black54,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              floatingActionButton: MouseRegion(
                child: FloatingActionButton(
                  onPressed: () {
                    setState(() {
                      _hoveredSaveButton = true;
                    });
                  },
                  elevation: 0,
                  focusElevation: 0,
                  hoverElevation: 0,
                  disabledElevation: 0,
                  highlightElevation: 0,
                  splashColor: Colors.transparent,
                  backgroundColor: Colors.blue.withAlpha(50),
                  child: AnimatedContainer(
                    duration: Duration(milliseconds: 250),
                    curve: Curves.easeInOut,
                    decoration: BoxDecoration(
                      borderRadius: BorderRadius.all(Radius.circular(50)),
                      boxShadow: [
                        BoxShadow(
                          color: Colors.black
                              .withAlpha(_hoveredSaveButton ? 30 : 25),
                          blurRadius: _hoveredSaveButton ? 9 : 10,
                          spreadRadius: 15,
                          offset: Offset(0, _hoveredSaveButton ? 10 : 20),
                        )
                      ],
                    ),
                    child: const Icon(
                      Icons.save,
                      color: Colors.blue,
                    ),
                  ),
                ),
                onHover: (x) {
                  setState(() {
                    _hoveredSaveButton = true;
                  });
                },
                onExit: (x) {
                  setState(() {
                    _hoveredSaveButton = false;
                  });
                },
              )),
          !_removeSplashscreen
              ? Positioned(
                  top: 0,
                  child: AnimatedOpacity(
                    opacity: _showSplashscreen ? 1 : 0,
                    duration: Duration(milliseconds: 250),
                    onEnd: () {
                      setState(() {
                        _removeSplashscreen = true;
                      });
                    },
                    child: Splashscreen(),
                  ),
                )
              : Container()
        ],
      ),
    );
  }
}

class Splashscreen extends StatelessWidget {
  Widget build(BuildContext context) {
    Size _size = MediaQuery.of(context).size;

    return DraggebleWidget(
      widget: Container(
        height: _size.height,
        width: _size.width,
        color: Colors.white,
        child: Stack(
          children: [
            Align(
              alignment: Alignment.center,
              heightFactor: _size.height,
              widthFactor: _size.width,
              child: Container(
                color: Colors.white,
                height: _size.height * 1.5,
                width: _size.height * 1.5,
                child: Image.asset(
                  "assets/fabric.png",
                  fit: BoxFit.cover,
                ),
              ),
            ),
            Align(
              alignment: Alignment.center,
              child: Material(
                type: MaterialType.transparency,
                child: Container(
                  color: Colors.transparent,
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.center,
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Container(
                        child: Image.asset(
                          "assets/icon_original.png",
                          height: _size.height / 3,
                        ),
                      ),
                      Text(
                        "reMousable",
                        style: TextStyle(
                          fontSize: 20,
                          color: Colors.black87,
                        ),
                      )
                    ],
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

class DraggebleWidget extends StatelessWidget implements PreferredSizeWidget {
  static const platform_channel_draggable =
      MethodChannel('samples.go-flutter.dev/draggable');

  final Widget? widget;

  DraggebleWidget({Key? key, required this.widget}) : super(key: key) {}

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
        child: MouseRegion(
          child: widget,
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
