import sys
import http.client
import threading

HOST = "localhost"
PORT = 8080

COLOR_RED = '\x1B[31m'
COLOR_GREEN = '\x1B[32m'
COLOR_RESET = '\x1B[0m'

def make_request():
    conn = http.client.HTTPConnection(f"{HOST}:{PORT}")

    # HTTP setup
    method = "GET"
    path = "/v1/hello"
    headers = {
        'Accept': "application/json"
    }
    
    # Send request
    conn.request(method, path, None, headers)

    res = conn.getresponse()
    data = res.read()
    
    if res.status != 200:
        print(f'{COLOR_RED}[ ERROR ] Failed in the request: {data} {COLOR_RESET}')
    else:
        print(f'{COLOR_GREEN}[  OK  ] Successful request: {len(data)} bytes received {COLOR_RESET}')

def make_multiple_request(thread_max):
    requests = [threading.Thread(target=make_request) for i in range(thread_max)]

    for request in requests:
        request.start()
    
    for request in requests:
        request.join()

def main():
    max = 100
    if len(sys.argv) > 2 and str(sys.argv[1]) == "--max":
        max = int(sys.argv[2])
    
    make_multiple_request(max)

main()