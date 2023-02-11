import requests
import json
from config import constants

def create_comment(post_id: int, body: str):
    payload = {
        post_id: post_id,
        body: body,
    }
    res = requests.get(constants.CREATE_COMMENT.url, json=payload)
    return json.loads(res.json())
