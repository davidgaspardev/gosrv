import sys
import http.client
import threading
import argparse

HOST = "localhost"
PORT = 8080

COLOR_RED = '\x1B[31m'
COLOR_GREEN = '\x1B[32m'
COLOR_RESET = '\x1B[0m'

result_status = 0

def make_request(method, path):
    global result_status 
    conn = http.client.HTTPConnection(f"{HOST}:{PORT}")

    # Send request
    conn.request(method, path, None)

    res = conn.getresponse()
    data = res.read()

    if res.status != 200:
        print(f'{COLOR_RED}[ ERROR ] Failed in the request: {data} {COLOR_RESET}')
        result_status = 1
    else:
        print(f'{COLOR_GREEN}[  OK  ] Successful request: {len(data)} bytes received {COLOR_RESET}')

def make_multiple_request(thread_max: int, method: str, path: str):
    requests = [threading.Thread(target=make_request, args=(method,path,)) for i in range(thread_max)]

    for request in requests:
        request.start()

    for request in requests:
        request.join()

def setup_arg():
    parser = argparse.ArgumentParser(description="Make multiple requests to the server")
    parser.add_argument("--method", type=str, help="HTTP method to be used", default="GET")
    parser.add_argument("--path", type=str, help="Path to be requested", default="/v1/hello")
    parser.add_argument("--max", type=int, help="Number of requests to be made", default=100)

    return parser.parse_args()

def main():
    args = setup_arg()

    make_multiple_request(args.max, args.method, args.path)

    sys.exit(result_status)

main()
