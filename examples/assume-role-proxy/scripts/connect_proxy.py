#!/usr/bin/env python3
"""Loopback-only CONNECT proxy. Logs destinations and byte counts, never headers/body."""
import argparse, datetime, http.server, json, select, socket, threading

parser=argparse.ArgumentParser()
parser.add_argument('--port',type=int,default=18080)
parser.add_argument('--log',required=True)
args=parser.parse_args()
lock=threading.Lock()
def record(**fields):
 with lock,open(args.log,'a') as f:
  f.write(json.dumps({'time':datetime.datetime.now(datetime.timezone.utc).isoformat(),**fields})+'\n')
class Handler(http.server.BaseHTTPRequestHandler):
 def log_message(self,*args): pass
 def do_CONNECT(self):
  if self.path!='sts.volcengineapi.com:443':
   self.send_error(403,'Only the STS endpoint is allowed'); return
  host,port=self.path.rsplit(':',1)
  try: upstream=socket.create_connection((host,int(port)),timeout=15)
  except OSError as exc:
   record(event='connect_failed',destination=self.path,error=type(exc).__name__)
   self.send_error(502); return
  record(event='connect',destination=self.path)
  self.send_response(200,'Connection Established'); self.end_headers()
  sent=received=0
  try:
   sockets=[self.connection,upstream]
   while True:
    ready,_,_=select.select(sockets,[],[],60)
    if not ready: break
    for source in ready:
     chunk=source.recv(65536)
     if not chunk: return
     target=upstream if source is self.connection else self.connection
     target.sendall(chunk)
     if source is self.connection: sent+=len(chunk)
     else: received+=len(chunk)
  except OSError: pass
  finally:
   upstream.close(); self.close_connection=True
   record(event='closed',destination=self.path,bytes_to_sts=sent,bytes_from_sts=received)
class Server(http.server.ThreadingHTTPServer):
 daemon_threads=True
record(event='start',listen=f'127.0.0.1:{args.port}')
Server(('127.0.0.1',args.port),Handler).serve_forever()
