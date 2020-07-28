import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

const String ApiBase = '/api/reminders';

const String _sharedPreferencesTitle = 'Settings';
const String _keyApiHost = 'apiHost';
const String _keyApiPort = 'apiPort';
const String _defaultApiHost = 'go-reminders.corp.local';
int _defaultApiPort = 8080;

class AppSettings {
  SharedPreferences _prefs;

  getSharedPrefs() async =>
      _prefs == null ? _prefs = await SharedPreferences.getInstance() : _prefs;

  Future<String> getApiHost() async {
    SharedPreferences prefs = await getSharedPrefs();
    return prefs.getString(_keyApiHost) ?? _defaultApiHost;
  }

  Future<int> getApiPort() async {
    SharedPreferences prefs = await getSharedPrefs();
    return prefs.getInt(_keyApiPort) ?? _defaultApiPort;
  }
}

class Settings extends StatefulWidget {
  final Function(bool) onSaved;

  void _onSaved(bool b) {
    if (onSaved != null) {
      onSaved(b);
    }
  }

  Settings({Function(bool) this.onSaved, Key key}) : super(key: key);

  @override
  _SettingsState createState() => _SettingsState();
}

class _SettingsState extends State<Settings> {
  String _apiHost;
  int _apiPort;

  final TextEditingController tecHost = TextEditingController();
  final TextEditingController tecPort = TextEditingController();

  @override
  void initState() {
    super.initState();
    _initState();
  }

  void _initState() async {
    AppSettings prefs = AppSettings();
    _apiHost = await prefs.getApiHost();
    _apiPort = await prefs.getApiPort();

    setState(() {
      tecHost.text = _apiHost;
      tecPort.text = _apiPort.toString();

      tecHost.addListener(() {
        _apiHost = tecHost.text;
      });
    });
  }

  void dispose() {
    tecHost.dispose();
    tecPort.dispose();
    super.dispose();
  }

  void _save(BuildContext ctx) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    bool b = await prefs.setString('apiHost', _apiHost) &&
        await prefs.setInt('apiPort', _apiPort);
    widget._onSaved(b);
  }

  Future<void> _invalidPort(BuildContext ctx, String port) async {
    return showDialog<void>(
      context: ctx,
      barrierDismissible: false, // user must tap button!
      builder: (BuildContext ctx) {
        return AlertDialog(
          title: Text('Invalid Port Specification'),
          content: SingleChildScrollView(
              child:
                  Text('Port "$port" is not an integer between 1 and 65536.')),
          actions: <Widget>[
            FlatButton(
              child: Icon(Icons.check, color: Colors.red),
              onPressed: () {
                Navigator.of(ctx).pop();
              },
            ),
          ],
        );
      },
    );
  }

  void portChanged(BuildContext ctx, String text) {
    int port = int.parse(text, onError: (text) {
      _invalidPort(ctx, text);
      tecPort.text = _apiPort.toString();
      return _apiPort;
    });
    if (port < 0 || port > 65535) {
      _invalidPort(ctx, tecPort.text);
      tecPort.text = _apiPort.toString();
    } else {
      _apiPort = port;
    }
  }

  @override
  Widget build(BuildContext ctx) {
    return Scaffold(
        appBar: AppBar(title: Text(_sharedPreferencesTitle), actions: <Widget>[
          IconButton(
            icon: const Icon(Icons.save),
            tooltip: 'Save',
            onPressed: () {
              _save(ctx);
            },
          ),
        ]),
        body: Center(
            child: Column(children: <Widget>[
          TextField(
              controller: tecHost,
              decoration: InputDecoration(
                  labelText: 'Hostname',
                  hintText: 'Enter a valid IP address',
                  helperText: 'Address of go-reminders server'),
              autocorrect: false,
              enableSuggestions: true,
              textAlign: TextAlign.left,
              keyboardType: TextInputType.text),
          TextField(
              controller: tecPort,
              decoration: InputDecoration(
                  labelText: 'IP Port number',
                  hintText: 'Enter a valid IP port',
                  helperText: 'Port on which go-reminders server is listening'),
              onChanged: (text) {
                portChanged(ctx, text);
              },
              keyboardType: TextInputType.number,
              autocorrect: false,
              textAlign: TextAlign.left)
        ])));
  }
}
