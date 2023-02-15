from .classification import Classification
from .regression import Regression
from pydantic import BaseModel, root_validator
import typing

class Evaluate(BaseModel):
    # def __init__(self, task_type):
    task_type:str

    kwargs:typing.Any = None
    scores:dict = {}
    task_evaluator:typing.Any = None

    # print(self.task_evaluator)


    
    class Config:
        validate_assignment = True
        arbitrary_types_allowed = True



    @root_validator(pre=True)
    def _set_fields(cls, values: dict) -> dict:

        task_type = values['task_type']
        
        if task_type is not None:

            if task_type == 'classification':
                task_evaluator = Classification()
            elif task_type == 'regression':
                task_evaluator = Regression()
            else:
                task_evaluator = None
            
            values['task_evaluator'] = task_evaluator
        
        return values
    


    def prepare_task_evaluator(self):
        print('inside prepare evaluator')
        if self.task_type == 'classification':
            self.task_evaluator = Classification()
        elif self.task_type == 'regression':
            self.task_evaluator = Regression()
        else:
            self.task_evaluator = None
        

        print(self.task_evaluator)



    def compute(self, references, predictions, prediction_scores=None, **kwargs):

        print(self.task_evaluator)
        if self.task_evaluator is not None:
            print('task evaluator is not none')
            self.task_evaluator.kwargs = kwargs
            self.task_evaluator.references = references
            self.task_evaluator.predictions = predictions
            self.task_evaluator.prediction_scores = prediction_scores

            self.scores = self.task_evaluator.compute()

        return self.scores


def evaluator(task_type, **kwargs):
    eval = Evaluate(task_type=task_type)
    
    return eval


