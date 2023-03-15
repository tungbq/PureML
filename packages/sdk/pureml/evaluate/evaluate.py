from pydantic import BaseModel
from pureml import BasePredictor
from .evaluator import Evaluate
from pureml.components import dataset
from typing import Any


class Test(BaseModel):
    task_type: str
    model_label: str
    dataset_label: str
    predictor: BasePredictor
    evaluator: Evaluate = None
    dataset: Any = None

    class Config:
        validate_assignment = True
        arbitrary_types_allowed = True

    def load_dataset(self):
        self.dataset = dataset.fetch(self.dataset_label)

    def load_predictor(self):
        self.predictor = 0
