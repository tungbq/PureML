from .dataset import DatasetSchema
from .paths import PathSchema
from .model import ModelSchema
from .prediction import PredictionSchema
from .storage import StorageSchema
from .packaging import DockerSchema, FastAPISchema


path_schema = PathSchema().get_instance()
model_schema = ModelSchema().get_instance()
