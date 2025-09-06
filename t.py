# Запустите это в отдельном терминале
from http.server import HTTPServer, BaseHTTPRequestHandler
import json

class TestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        response_data = {'test': 'data', 'delay': 0}
        response_json = json.dumps(response_data).encode()
        self.send_response(200)
        self.send_header('Content-Type', 'application/json')
        self.send_header('Content-Length', str(len(response_json)))
        self.end_headers()
        self.wfile.write(response_json)

server = HTTPServer(('localhost', 8000), TestHandler)
print('Serving on port 8000')
server.serve_forever()
