import 'dart:io';
import 'package:flutter/material.dart';

import 'package:ui/reminders/requests.dart';

class _SaveReminder extends StatefulWidget {
  final Map<String, dynamic> _reminder;
  final Function(String) onShowMessage;
  final Function(TupleReminder t) onSaved;

  _SaveReminder(this._reminder,
      {@required this.onShowMessage, @required this.onSaved, Key key})
      : super(key: key);

  @override
  _SaveReminderState createState() => _SaveReminderState();
}

class _SaveReminderState extends State<_SaveReminder> {
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
    String msg;
    if (t.httpStatus != HttpStatus.ok) {
      msg = 'Save error: ${t.reminder[fieldMessage]}';
    } else {
      String uuid = widget._reminder[fieldGuid];
      msg = 'Saved Reminder $uuid';
    }
    print(msg);

    print('_SaveReminder:handleTuple: calling widget.onShowMessage.');
    widget.onShowMessage(msg);

    print('_SaveReminder:handleTuple: calling widget.onSaved.');
    widget.onSaved(t);
  }

  @override
  void initState() {
    super.initState();

    _updateServer().then(handleTuple).catchError((error) {
      setState(() {
        print('_SaveReminder:initState: calling widget.onShowMessage.');
        widget.onShowMessage('Save error: ${error.toString()}');
      });
    });
  }

  @override
  Widget build(BuildContext ctx) {
    return Center(child: CircularProgressIndicator());
  }
}

class ReminderEditor extends StatefulWidget {
  final Map<String, dynamic> reminder;
  final Function(TupleReminder) onSaved;
  final Function(String) onShowMessage;

  ReminderEditor(
      {@required this.reminder,
      @required this.onSaved,
      @required this.onShowMessage,
      Key key})
      : super(key: key);

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

  /// Internal onSaved to copy the save results and pop the save dialog.
  void _onSaved(BuildContext ctx, TupleReminder t) {
    // Update the local reminder.
    String msg = 'ReminderEditorState: saving reminder to local copy.';
    print(msg);
    if (t != null)
      t.reminder.forEach((key, value) => widget.reminder[key] = value);

    // update the current display in case the parent doesn't pop.
    print('ReminderEditorState: setState.');
    setState(() {});

    // Pop the save process indicator dialog.
    print('ReminderEditorState: popping save process indicator dialog.');
    Navigator.of(ctx).pop(t);
  }

  void _save(BuildContext ctx) async {
    widget.reminder[fieldMessage] = _message;

    // Push a dialog to show a process indicator while saving.
    TupleReminder t = await showDialog<TupleReminder>(
        context: ctx,
        builder: (BuildContext ctx) {
          return _SaveReminder(widget.reminder, onSaved: (TupleReminder t) {
            _onSaved(ctx, t);
          }, onShowMessage: widget.onShowMessage);
        });

    // Notify the parent when the indicator dialog closes.
    print('ReminderEditorState: notifying parent of saved reminder.');
    widget.onSaved(t);
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
            Builder(builder: (BuildContext ctx) {
              return IconButton(
                  icon: const Icon(Icons.save),
                  tooltip: 'Save',
                  onPressed: () {
                    _save(ctx);
                  });
            })
          ],
        ),
        body: _getBody(ctx));
  }
}
