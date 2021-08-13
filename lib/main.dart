import 'package:flutter/material.dart';
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
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      _counter++;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: DraggebleAppBar(title: "reMousable"),
      backgroundColor: Colors.white,
      body: const Center(
        child: Text("reMousable"),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
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
      backgroundColor: Colors.transparent,
      shadowColor: Colors.transparent,
      bottom: PreferredSize(
          child: Container(
            color: Colors.blue,
            height: 1.0,
          ),
          preferredSize: Size.fromHeight(4.0)),
      title: Text(title, style: TextStyle(color: Colors.black)),
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
        child: appBar, onPanStart: onPanStart, onPanUpdate: onPanUpdate);
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
}
