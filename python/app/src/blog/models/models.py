from django.db import models
from dataclasses import dataclass
import datetime
import uuid

@dataclass
class User:
    id: uuid.UUID
    name: str

@dataclass
class Post:
    id: int
    title: str
    tag_ids: list[int]
    body: str
    comment_ids: list[int]
    posted_by: User
    created_at: datetime.datetime
    updated_at: datetime.datetime

@dataclass
class Comment:
    id: int
    body: str
    created_at: datetime.datetime
    updated_at: datetime.datetime

@dataclass
class Tag:
    id: int
    name: str