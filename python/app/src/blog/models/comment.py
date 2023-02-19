from django.db import models
from dataclasses import dataclass
import datetime
import uuid
import utils

@dataclass
class Comment:
    comment_id: int
    body: str
    created_at: datetime.datetime
    updated_at: datetime.datetime

    @staticmethod
    def new(
            comment_id: int = 0,
            body: str = '',
            created_at: datetime.datetime = datetime.datetime.now(),
            updated_at: datetime.datetime = datetime.datetime.now(),
        ):
        if utils.is_empty(body):
            raise Exception('body cannot be empty')

        if updated_at < created_at:
            raise Exception('date updated at is ahead of date created')

        return Comment(
            comment_id=comment_id if comment_id != 0 else None,
            body=body,
            created_at=created_at,
            updated_at=updated_at,
        )

