from django.db import models
from dataclasses import dataclass
import datetime
import utils

@dataclass
class User:
    user_id: str
    name: str

    @staticmethod
    def new(user_id: str = '', name: str = ''):
        if utils.is_empty(name):
            raise Exception('user name must not be empty')
        return User(user_id, name)

    def signup(self, name, password):
        return 

    def login(self, name, password):
        return