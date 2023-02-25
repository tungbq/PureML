from pydantic import BaseModel, Field
import typing
from abc import ABC, abstractmethod


class BasePredictor(BaseModel, ABC):
    model_details: typing.Union(typing.List(str), typing.List(typing.List(str)))
    model: typing.Any = None
    requirements_py: typing.List(str) = None
    requirements_sys: typing.List(str) = None

    class Config:
        arbitrary_types_allowed = True

    @abstractmethod
    def predict(self, **kwargs: typing.Any):
        pass

    @abstractmethod
    def load_requirements_py(self):
        pass

    @abstractmethod
    def load_requirements_sys(self):
        pass

    @abstractmethod
    def load_models(self):
        pass

    @abstractmethod
    def load_resources(self):
        pass
