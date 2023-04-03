from pydantic import BaseModel, root_validator
from .backend import BackendSchema
import typing
import os


class StorageSchema(BaseModel):

    STORAGE: str = "PUREML-STORAGE"
    backend: BackendSchema = BackendSchema().get_instance()

    class Config:
        arbitrary_types_allowed = True
