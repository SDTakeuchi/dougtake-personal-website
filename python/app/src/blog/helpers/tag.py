import requests
import json
from config import constants

def find_tags() -> dict:
    res = requests.get(constants.FIND_TAGS.url)
    return json.loads(res.json())
