import json
import requests

def lambda_handler(event, context):
    url = event["url"]
    method = event["method"]
    
    if method == "POST":
        payload = event["payload"]
        response = requests.post(url, data = payload)
    else:
        response = requests.get(url)
        
    print(response.text)
    
    return {
        'statusCode': 200,
        'body': json.loads(response.text)
    }