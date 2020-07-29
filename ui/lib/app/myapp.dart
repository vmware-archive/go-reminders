import 'package:flutter/material.dart';

import 'package:ui/app/homepage.dart';

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        title: 'Go-Reminders',
        theme: ThemeData(
          primarySwatch: Colors.blue,
        ),
        home: HomePage('Go-Reminders'));
  }
}
