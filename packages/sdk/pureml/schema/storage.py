from pydantic import BaseModel, root_validator
from .singleton import Singleton_BaseModel
import typing
import os


class StorageSchema(Singleton_BaseModel):

    STORAGE: str = "PUREML-STORAGE"

    class Config:
        arbitrary_types_allowed = True
