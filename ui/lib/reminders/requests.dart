import 'dart:async';
import 'dart:convert';
import 'dart:io';
import 'package:http/http.dart' as http;
import 'package:ui/app/settings.dart';

const Map<String, String> _headers = {'content-type': 'application/json'};

const String fieldMessage = 'message';
const String fieldGuid = 'guid';
const String fieldId = 'id';

class TupleReminder {
  int httpStatus;
  Map<String, dynamic> reminder;

  TupleReminder(this.httpStatus, {jsonEncoded: String}) {
    if (jsonEncoded == null || jsonEncoded.length > 0) {
      reminder = json.decode(jsonEncoded);
    }
  }

  TupleReminder.withError(int httpErrorStatus) {
    httpStatus = httpErrorStatus;
    reminder = null;
  }
}

// Get all existing reminders.
Future<List<Map<String, dynamic>>> getAllReminders() async {
  AppSettings prefs = AppSettings();
  String apiHost = await prefs.getApiHost();
  int apiPort = await prefs.getApiPort();

  http.Response res =
      await http.get(Uri.http('$apiHost:$apiPort', ApiBase), headers: _headers);
  if (res.statusCode == HttpStatus.ok) {
    var j = json.decode(res.body);
    List<Map<String, dynamic>> reminders = List<Map<String, dynamic>>();

    if (j is List<dynamic>) {
      reminders = List<Map<String, dynamic>>();
      for (int i = 0; i < j.length; i++) {
        Map<String, dynamic> r = j[i];
        reminders.add(r);
      }
    } else if (j is List<Map<String, dynamic>>) {
      reminders = j;
    } else {
      reminders = List<Map<String, dynamic>>();
    }

    return reminders;
  } else {
    return [
      {fieldMessage: "Error: HttpStatus=${res.statusCode}"}
    ];
  }
}

// Get an existing Reminder by int property property.
Future<TupleReminder> getById(int id) async {
  AppSettings prefs = AppSettings();
  String apiHost = await prefs.getApiHost();
  int apiPort = await prefs.getApiPort();

  http.Response res = await http.get(
      Uri.http('$apiHost:$apiPort', '$ApiBase/byid/$id'),
      headers: _headers);
  return TupleReminder(res.statusCode, jsonEncoded: res.body);
}

// Get an existing Reminder by Uuid.
Future<TupleReminder> getByUuid(String uuid) async {
  AppSettings prefs = AppSettings();
  String apiHost = await prefs.getApiHost();
  int apiPort = await prefs.getApiPort();

  http.Response res = await http
      .get(Uri.http('$apiHost:$apiPort', '$ApiBase/$uuid'), headers: _headers);
  return TupleReminder(res.statusCode, jsonEncoded: res.body);
}

// Create a Reminder, returns TupleReminder with:
//   Error message and http status code, or
//   new reminder and ok statuscode.
Future<TupleReminder> postReminder(String message) async {
  AppSettings prefs = AppSettings();
  String apiHost = await prefs.getApiHost();
  int apiPort = await prefs.getApiPort();

  Map<String, dynamic> reminder = {fieldMessage: message};
  String body = json.encode(reminder);
  http.Response res = await http.post(Uri.http('$apiHost:$apiPort', ApiBase),
      headers: _headers, body: body);
  return TupleReminder(res.statusCode, jsonEncoded: res.body);
}

// Delete a Reminder, returns TupleReminder with:
//   Error message and http status code, or
//   new reminder and ok statuscode.
Future<TupleReminder> deleteReminder(String uuid) async {
  AppSettings prefs = AppSettings();
  String apiHost = await prefs.getApiHost();
  int apiPort = await prefs.getApiPort();

  http.Response res = await http.delete(
      Uri.http('$apiHost:$apiPort', '$ApiBase/$uuid'),
      headers: _headers);
  return TupleReminder(res.statusCode, jsonEncoded: res.body);
}

// Update a Reminder, returns TupleReminder with:
//   Error message and http status code, or
//   updated reminder and ok statuscode.
Future<TupleReminder> putReminder(String uuid, String message) async {
  if (uuid == null || uuid.length < 0) {
    return TupleReminder(HttpStatus.badRequest,
        jsonEncoded: '{"error", "PUT request cannot be null or empty"}');
  }

  // Get the reminder freshly from storage, reset the message and update.
  TupleReminder tuple = await getByUuid(uuid);
  if (tuple.httpStatus != HttpStatus.ok) {
    return tuple;
  }

  AppSettings prefs = AppSettings();
  String apiHost = await prefs.getApiHost();
  int apiPort = await prefs.getApiPort();

  tuple.reminder[fieldMessage] = message;
  http.Response res = await http.put(
      Uri.http('$apiHost:$apiPort', '$ApiBase/$uuid'),
      headers: _headers,
      body: json.encode(tuple.reminder));

  return TupleReminder(res.statusCode, jsonEncoded: res.body);
}
