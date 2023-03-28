from pureml.components.model import register
from pureml.utils.pipeline import (
    add_model_to_config,
    load_metrics_from_config,
    load_params_from_config,
    load_figures_from_config,
)
from pureml import metrics, params, figure
from pureml.utils.version_utils import parse_version_label, generate_label
import functools


def model(label: str):
    def decorator(func):
        # print('Inside decorator')
        name, branch, _ = parse_version_label(label)

        # print('Adding model name: ', name, 'to config before invoking user function')
        add_model_to_config(name=name, branch=branch)

        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            # print("Inside wrapper")
            func_name = func.__name__
            func_description = func.__doc__

            func_output = func(*args, **kwargs)

            model_exists_in_remote, model_hash, model_version = register(
                model=func_output, label=label
            )

            if (
                model_exists_in_remote
            ):  # Only add the model to config if it is successfully pushed

                label_new = generate_label(name, branch, model_version)

                add_model_to_config(
                    name=name,
                    branch=branch,
                    hash=model_hash,
                    version=model_version,
                    func=func,
                )

                metric_values = load_metrics_from_config()
                if len(metric_values) != 0:
                    metrics.add(metrics=metric_values, label=label_new)

                param_values = load_params_from_config()
                if len(param_values) != 0:
                    params.add(params=param_values, label=label_new)

                figure_file_paths = load_figures_from_config()
                if len(figure_file_paths) != 0:
                    figure.add(file_paths=figure_file_paths, label=label_new)

            else:
                add_model_to_config(
                    name=name, branch=branch, hash=model_hash, func=func
                )

            return func_output

        # print("Outside  wrapper")

        return wrapper

    # print('Outside decorator')

    return decorator
