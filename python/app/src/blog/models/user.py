from django.db import models
from dataclasses import dataclass
import datetime
import uuid

@dataclass
class User:
    id: uuid.UUID
    name: str

    def signup(self, name, password):
        return 

    def login(self, name, password):
        return