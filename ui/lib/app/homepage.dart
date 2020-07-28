import 'package:flutter/material.dart';

import 'package:ui/app/settings.dart';
import 'package:ui/reminders/list.dart';
import 'package:ui/reminders/edit.dart';
import 'package:ui/reminders/requests.dart';

class HomePage extends StatelessWidget {
  final String title;
  final GlobalKey<ReminderQueryState> _keyReminderQuery = GlobalKey();

  HomePage(this.title, {Key key}) : super(key: key);

  void _onShowMessage(BuildContext ctx, String msg) {
    //Scaffold.of(ctx).showSnackBar(SnackBar(content: Text(msg)));
    print(msg);
  }

  void _onSaved(BuildContext ctx, TupleReminder t) {
    print('HomePage: popping ReminderEditor.');
    //_onShowMessage(ctx, 'Saved new reminder: ${t.reminder[fieldGuid]}.');
    Navigator.of(ctx).pop(t);
  }

  void _onAdd(BuildContext ctx) {
    Map<String, dynamic> reminder = Map<String, dynamic>();
    Navigator.of(ctx)
        .push(MaterialPageRoute(
            builder: (innerCtx) => ReminderEditor(
                reminder: reminder,
                onShowMessage: (String msg) {
                  _onShowMessage(innerCtx, msg);
                },
                onSaved: (TupleReminder t) {
                  _onSaved(innerCtx, t);
                })))
        .then((data) {
      onRefresh();
    }).catchError((error) {
      print('Caught error in _onAdd');
    });
  }

  void _onSettings(BuildContext ctx) {
    Navigator.of(ctx).push(MaterialPageRoute(
        builder: (context) => Settings(onSaved: (b) {
              Navigator.of(ctx).pop(b);
            })));
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
        body: Builder(
            builder: (BuildContext innerCtx) => ReminderQuery(
                key: _keyReminderQuery,
                onShowMessage: (String msg) {
                  _onShowMessage(innerCtx, msg);
                })));
  }
}
