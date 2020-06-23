import 'package:flutter/material.dart';

import 'package:ui/app/navigation.dart';
import 'package:ui/app/settings.dart';
import 'package:ui/reminders/list.dart';
import 'package:ui/reminders/edit.dart';

class HomePage extends StatelessWidget {
  final String title;
  final GlobalKey<ReminderQueryState> _keyReminderQuery = GlobalKey();

  HomePage(this.title, {Key key}) : super(key: key);

  void _onAdd(BuildContext ctx) {
    Map<String, dynamic> reminder = Map<String, dynamic>();
    globalNavigator.push(ReminderEditor(reminder: reminder)).then((data) {
      onRefresh();
    }).catchError((error) {
      print('Caught error in _onAdd');
    });
  }

  void _onSettings(BuildContext ctx) {
    globalNavigator.push(Settings());
  }

  void onRefresh() {
    _keyReminderQuery.currentState.onRefresh();
  }

  @override
  Widget build(BuildContext ctx) {
    return Scaffold(
        appBar: AppBar(
          title: Text(title),
          actions: <Widget>[
            IconButton(
              icon: const Icon(Icons.add),
              tooltip: 'New Reminder',
              onPressed: () {
                _onAdd(ctx);
              },
            ),
            IconButton(
              icon: const Icon(Icons.refresh),
              tooltip: 'Refresh',
              onPressed: () {
                onRefresh();
              },
            ),
            IconButton(
              icon: const Icon(Icons.settings),
              tooltip: 'Settings',
              onPressed: () {
                _onSettings(ctx);
              },
            ),
          ],
        ),
        body: ReminderQuery(key: _keyReminderQuery));
  }
}
