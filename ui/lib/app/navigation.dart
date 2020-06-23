import 'package:flutter/material.dart';

final _GlobalNavigator globalNavigator = _GlobalNavigator();

class _GlobalNavigator {
  final GlobalKey<NavigatorState> _key = new GlobalKey<NavigatorState>();

  Future<dynamic> pushNamed(String route) {
    return _key.currentState.pushNamed(route);
  }

  Future<dynamic> push(Widget w) {
    return _key.currentState.push(MaterialPageRoute(builder: (context) => w));
  }

  void showSnackBar(String msg) {
    //return _key.currentState.push(MaterialPageRoute(builder: (context) => w));
    Scaffold.of(_key.currentState.context)
        .showSnackBar(SnackBar(content: Text(msg)));
  }

  void pop(dynamic d) {
    if (d == null) {
      print('GlobalNavigator: popping with nothing (because d is null).');
      _key.currentState.pop();
    } else {
      print('GlobalNavigator: popping with data.');
      _key.currentState.pop(d);
    }
  }

  GlobalKey<NavigatorState> key() => _key;
}

class GlobalNavigator {
  final GlobalKey<NavigatorState> key = new GlobalKey<NavigatorState>();

  Future<dynamic> pushNamed(String route) {
    return key.currentState.pushNamed(route);
  }

  Future<dynamic> push(Widget w) {
    return key.currentState.push(MaterialPageRoute(builder: (context) => w));
  }

  void pop(dynamic d) {
    if (d == null) {
      print('pop: data is null, popping with nothing');
      key.currentState.pop();
    } else {
      print('pop: data is null, popping with data');
      key.currentState.pop();
    }
  }
}
