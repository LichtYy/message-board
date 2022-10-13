import requests
import os
import urllib3

urllib3.disable_warnings()

url = "https://github.com/LichtYy/message-board/raw/ceceb39e63610af25de924cc39d3e0f7c400b067/toolbar.js"
my_headers = {
    'User-Agent':'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36 Edg/97.0.1072.69',
    'Connection': 'close'
}
os.environ['NO_PROXY'] = "https://github.com"
resp = requests.get(url, headers=my_headers, verify=False)
print(resp.text)

