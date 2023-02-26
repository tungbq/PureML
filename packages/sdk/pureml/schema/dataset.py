from pydantic import BaseModel

from .paths import PathSchema
import os


class DatasetSchema(BaseModel):
    paths: PathSchema = PathSchema().get_instance()

    PATH_DATASET_README = os.path.join(paths.PATH_DATASET_DIR, "ReadME.md")
