from django.db import models
from dataclasses import dataclass
import datetime
import utils

@dataclass
class Post:
    id: int
    title: str
    tag_ids: list[int]
    body: str
    comment_ids: list[int]
    user_id: int
    created_at: datetime.datetime
    updated_at: datetime.datetime

    @staticmethod
    def new(id: int = 0,
            title: str = '',
            tag_ids: list[int] = [],
            body: str = '',
            comment_ids: list[int] = [],
            posted_by: int = 0,
            created_at: datetime.datetime = datetime.datetime.now(),
            updated_at: datetime.datetime = datetime.datetime.now()):

        if not isinstance(id, int) or \
                not isinstance(title, str) or \
                not isinstance(tag_ids, list) or \
                not isinstance(body, str) or \
                not isinstance(comment_ids, list) or \
                not isinstance(posted_by, int) or \
                not isinstance(created_at, datetime.datetime) or \
                not isinstance(updated_at, datetime.datetime):
            raise Exception('Post.new() received mis-typed input')

        if updated_at > created_at:
            raise Exception('Post.new() error: updated time is ahead of updated time')

        for tag_id in tag_ids:
            if not isinstance(tag_id, int):
                raise Exception('Post.new() received mis-typed input')

        for comment_id in comment_ids:
            if not isinstance(comment_id, int):
                raise Exception('Post.new() received mis-typed input')

        if utils.is_empty(title):
            raise Exception('Post must have title')
        
        if utils.is_empty(tag_ids):
            raise Exception('Post must be connected with at least one tag')

        if utils.is_empty(body):
            raise Exception('Post must have body')
        
        if utils.is_empty(posted_by):
            raise Exception('Post must have its author')

        return Post(
            id=id,
            title=title,
            tag_ids=tag_ids,
            body=body,
            comment_ids=comment_ids,
            user_id=posted_by,
            created_at=created_at,
            updated_at=updated_at
        )