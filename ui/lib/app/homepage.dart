import 'package:flutter/material.dart';

import 'package:ui/app/settings.dart';
import 'package:ui/reminders/list.dart';
import 'package:ui/reminders/edit.dart';
import 'package:ui/reminders/requests.dart';

final GlobalKey<ReminderQueryState> _keyReminderQuery = GlobalKey();

class _NewReminderAction extends StatelessWidget {
  void _onAdd(BuildContext ctx) {
    Map<String, dynamic> reminder = Map<String, dynamic>();
    Navigator.of(ctx)
        .push(MaterialPageRoute(
            builder: (innerCtx) => ReminderEditor(
                reminder: reminder,
                onShowMessage: (String msg) {
                  Scaffold.of(ctx).showSnackBar(SnackBar(content: Text(msg)));
                },
                onSaved: (TupleReminder t) {
                  print('_NewReminder: popping ReminderEditor.');
                  Navigator.of(ctx).pop(t);
                })))
        .then((data) {
      _keyReminderQuery.currentState.onRefresh();
    }).catchError((error) {
      print('Caught error in _onAdd');
    });
  }

  @override
  Widget build(BuildContext ctx) {
    return IconButton(
        icon: const Icon(Icons.add),
        tooltip: 'New Reminder',
        onPressed: () {
          _onAdd(ctx);
        });
  }
}

class HomePage extends StatelessWidget {
  final String title;

  HomePage(this.title, {Key key}) : super(key: key);

  void _onShowMessage(BuildContext ctx, String msg) {
    ScaffoldState s;

    try {
      s = Scaffold.of(ctx);
    } catch (e) {
      print(
          'HomePage._onShowMessage: Scaffold not found for BuildContext $ctx.toString().');
      print(e.toString());
    }

    if (s == null) {
      print(msg);
    } else {
      Scaffold.of(ctx).showSnackBar(SnackBar(content: Text(msg)));
    }
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
            _NewReminderAction(),
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
