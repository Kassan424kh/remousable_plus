import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

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
