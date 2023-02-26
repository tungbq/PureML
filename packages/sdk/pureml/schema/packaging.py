from pydantic import BaseModel

import typing
import os
from .paths import PathSchema


class DockerSchema(BaseModel):
    paths: PathSchema = PathSchema().get_instance()
    # PORT_DOCKER = 8005      #Same port as fastapi server
    PORT_HOST = 8000
    BASE_IMAGE_DOCKER = "python:3.8-slim"
    API_IP_DOCKER = "0.0.0.0"
    API_IP_HOST = "0.0.0.0"

    class Config:
        arbitrary_types_allowed = True


class FastAPISchema(BaseModel):
    paths: PathSchema = PathSchema().get_instance()
    PATH_FASTAPI_FILE = os.path.join(paths.PATH_PREDICT_DIR, "fastapi_server.py")
    PORT_FASTAPI = 8005  # Same port as docker server

    class Config:
        arbitrary_types_allowed = True
