#!/usr/bin/env python3

import os
import sys

import cgi
import cgitb
cgitb.enable()

import json

form = cgi.FieldStorage()
print('Content-Type: application/json\n')

required = ['name', 'topic', 'duration', 'description']

response = {}
values = ()

for item in required:
    if item not in form:
        response['success'] = False
        if response.get('missing') is None:
            response['missing'] = []
        response['missing'].append(item)
        continue
    if item is 'duration':
        values += (int(form.getfirst(item)),)
        continue
    values += (form.getfirst(item),)

if response.get('success') is False:
    json.dump(response, sys.stdout)
    sys.exit()

import sqlite3

db = sqlite3.connect('/srv/homebrew.db')
query = 'INSERT INTO talks (name, topic, duration, description) VALUES (?, ?, ?, ?)'

db.execute(query, values)
db.commit()
db.close()

response['success'] = True
response['values'] = values
json.dump(response, sys.stdout)
