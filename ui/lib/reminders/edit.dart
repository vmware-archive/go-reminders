import 'dart:io';
import 'package:flutter/material.dart';

import 'package:ui/app/navigation.dart';
import 'package:ui/reminders/requests.dart';

class _SaveReminder extends StatefulWidget {
  final Map<String, dynamic> _reminder;

  _SaveReminder(this._reminder, {Key key}) : super(key: key);

  @override
  _SaveReminderState createState() => _SaveReminderState();
}

class _SaveReminderState extends State<_SaveReminder> {
  String _msg;

  Future<TupleReminder> _updateServer() async {
    TupleReminder t;

    if (widget._reminder.containsKey(fieldGuid)) {
      t = await putReminder(
          widget._reminder[fieldGuid], widget._reminder[fieldMessage]);
    } else {
      t = await postReminder(widget._reminder[fieldMessage]);
    }

    if (t.httpStatus == HttpStatus.ok) {
      t.reminder.forEach((key, value) => widget._reminder[key] = value);
    }

    return t;
  }

  handleTuple(TupleReminder t) {
    if (t.httpStatus != HttpStatus.ok) {
      _msg = 'Save error: ${t.reminder[fieldMessage]}';
    } else {
      String uuid = widget._reminder[fieldGuid];
      _msg = 'Saved Reminder $uuid';
    }
    print(_msg);

    globalNavigator.pop(t);
    globalNavigator.showSnackBar(_msg);
  }

  @override
  void initState() {
    super.initState();

    _updateServer().then(handleTuple).catchError((error) {
      setState(() {
        _msg = 'Save error: ${error.toString()}';
      });
    });
  }

  @override
  Widget build(BuildContext ctx) {
    return Center(child: CircularProgressIndicator());
  }
}

class ReminderEditor extends StatefulWidget {
  ReminderEditor({@required this.reminder, Key key}) : super(key: key);

  final Map<String, dynamic> reminder;

  @override
  _ReminderEditorState createState() => _ReminderEditorState();
}

class _ReminderEditorState extends State<ReminderEditor> {
  String _message;
  final TextEditingController _tecMessage = TextEditingController();

  @override
  void initState() {
    super.initState();
    _message = widget.reminder.containsKey(fieldMessage)
        ? widget.reminder[fieldMessage]
        : '';
    _tecMessage.text = _message;
  }

  void dispose() {
    _tecMessage.dispose();
    super.dispose();
  }

  void _setMessage(String m) {
    setState(() {
      _message = m;
    });
  }

  void _save(BuildContext ctx) async {
    widget.reminder[fieldMessage] = _message;
    TupleReminder t = await showDialog<TupleReminder>(
        context: ctx,
        builder: (BuildContext ctx) {
          return _SaveReminder(widget.reminder);
        });
    globalNavigator.pop(t);
    if (t != null)
      t.reminder.forEach((key, value) => widget.reminder[key] = value);
    setState(() {});
  }

  Widget _getBody(BuildContext ctx) {
    return Center(
        child: Column(children: <Widget>[
      Text('Enter reminder text:', textAlign: TextAlign.left),
      Expanded(
          child: Container(
              padding: EdgeInsets.all(8.0),
              child: TextField(
                  controller: _tecMessage,
                  maxLines: null,
                  onChanged: _setMessage,
                  autocorrect: true,
                  enableSuggestions: true,
                  textAlign: TextAlign.left)))
    ]));
  }

  @override
  Widget build(BuildContext ctx) {
    return Scaffold(
        appBar: AppBar(
          title: Text("Edit Reminder"),
          actions: <Widget>[
            IconButton(
                icon: const Icon(Icons.save),
                tooltip: 'Save',
                onPressed: () async {
                  _save(ctx);
                })
          ],
        ),
        body: _getBody(ctx));
  }
}
