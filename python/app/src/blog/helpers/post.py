import requests
import json
from config import constants
import utils

def find_posts(post_id: int, tag_id: int, search_char: str, offset: int, limit: int) -> dict:
    if utils.is_empty(post_id):
        params = {
            tag_id: tag_id,
            search_char: search_char,
            offset: offset,
            limit: limit,
        }
        res = requests.get(constants.FIND_POSTS, params=params)
    else:
        res = requests.get(constants.FIND_POSTS + post_id)
    return json.loads(res.json())
