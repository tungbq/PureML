import typing
import numpy as np
import matplotlib.pyplot as plt
from pureml.components import metrics as pure_metrics
from pureml.components import params as pure_params
from pureml.components import figure as pure_figure
from pureml.utils.version_utils import parse_version_label

# def log(**kwargs):

#     if 'metrics' in kwargs.keys():
#         metrics.add(kwargs['metrics'])


#     if 'params' in kwargs.keys():
#         params.add(kwargs['params'])


def log(label: str, metrics=None, params=None, step=1, **kwargs):

    if metrics is not None:
        func_params = {}

        if label is not None:
            func_params["label"] = label

        func_params["metrics"] = metrics.copy()

        func_params["step"] = step

        pure_metrics.add(**func_params)

    if params is not None:
        func_params = {}

        if label is not None:
            func_params["label"] = label

        func_params["params"] = params.copy()

        func_params["step"] = step

        pure_params.add(**func_params)

    if "fig" in kwargs.keys():
        figure = kwargs["fig"]
        func_params = {}

        if model_name is not None:
            func_params["model_name"] = model_name
        if model_branch is not None:
            func_params["model_branch"] = model_branch
        if model_version is not None:
            func_params["model_version"] = model_version

        func_params["figure"] = figure.copy()
        # func_params['step']  = step

        pure_figure.add(**func_params)
