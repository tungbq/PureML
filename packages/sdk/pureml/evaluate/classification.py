from .metrics import Accuracy, Precision, Recall, F1, ROC_AUC


class Classification():
    def __init__(self):
        self.task_type = 'classification'

        self.kwargs = None

        self.references = None
        self.predictions = None
        self.prediction_scores = None

        self.metrics = [Accuracy(), Precision(), Recall(), F1()]#, ROC_AUC()]
        self.scores = {}

    
    def compute(self):
        
        for m in self.metrics:
            #Adding  prediction scores to kwargs. It will be utilized my metrics needing it(roc_auc).
            self.kwargs['prediction_scores'] = self.prediction_scores

            score = m.compute(references=self.references, predictions=self.predictions, **self.kwargs)

            self.scores.update(score)

        return self.scores
    
