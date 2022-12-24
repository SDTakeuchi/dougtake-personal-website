from django.db import models
from dataclasses import dataclass
import datetime
import uuid

@dataclass
class Tag:
    id: int
    name: str