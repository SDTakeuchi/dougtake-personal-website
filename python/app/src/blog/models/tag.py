from django.db import models
from dataclasses import dataclass
import utils

@dataclass
class Tag:
    tag_id: int
    name: str

    @staticmethod
    def new(tag_id: int = 0, name: str = ''):
        if utils.is_empty(name):
            raise Exception('Tag name must not be empty')
        return Tag(
            tag_id=tag_id if tag_id != 0 else None,
            name=name,
        )