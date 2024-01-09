import base64
import json


f = open("records.txt","r")
w = open("records_parsed.txt","w")
parsed = json.loads(f.read())

for d in parsed:
    line = base64.b64decode(d['Data']).decode()
    message = json.loads(line)['message']
    w.write(message+"\n")

print("parsing complete")
f.close()
w.close()