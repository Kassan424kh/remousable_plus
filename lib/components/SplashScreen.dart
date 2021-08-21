import 'package:flutter/material.dart';

import 'DraggableWidget.dart';
import 'dart:math' as math;

class Splashscreen extends StatefulWidget {
  @override
  _SplashscreenState createState() => _SplashscreenState();
}

class _SplashscreenState extends State<Splashscreen>
    with TickerProviderStateMixin {
  late double _animate = 1.0;
  late AnimationController _controller;

  @override
  void initState() {
    super.initState();

    _controller = AnimationController(
        vsync: this,
        lowerBound: 1.0,
        upperBound: 2.0,
        duration: Duration(seconds: 10));
    _controller.addListener(() {
      setState(() {
        _animate = _controller.value;
      });
    });

    _controller.forward();
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

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
              child: Transform.scale(
                scale: _animate,
                child: Transform.rotate(
                  angle: (((_animate - 1) * 100) / 100 * 20) * math.pi / 180,
                  child: Container(
                    color: Colors.white,
                    height: _size.height * 1,
                    width: _size.height * 1,
                    child: Image.asset(
                      "assets/fabric.png",
                      fit: BoxFit.fitHeight,
                    ),
                  ),
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
            /* Material(
              type: MaterialType.transparency,
              child: Align(
                alignment: Alignment.center,
                heightFactor: _size.height,
                widthFactor: _size.width,
                child: Container(
                  color: Colors.transparent,
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text("|"),
                      Text("|"),
                    ],
                  ),
                ),
              ),
            ),
            Material(
              type: MaterialType.transparency,
              child: Align(
                alignment: Alignment.center,
                heightFactor: _size.height,
                widthFactor: _size.width,
                child: Container(
                  color: Colors.transparent,
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text("-"),
                      Text("-"),
                    ],
                  ),
                ),
              ),
            ), */
          ],
        ),
      ),
    );
  }
}
