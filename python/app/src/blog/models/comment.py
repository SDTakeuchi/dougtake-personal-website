from django.db import models
from dataclasses import dataclass
import datetime
import uuid

@dataclass
class Comment:
    id: int
    body: str
    created_at: datetime.datetime
    updated_at: datetime.datetime
