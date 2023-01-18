from pureml.components.dataset import register
from pureml.utils.pipeline import add_dataset_to_config
from pureml.pipeline.data.create_pipeline import create_pipeline

def dataset(name:str, parent:str=None, upload=False):

    def decorator(func):

        #Add dataset name to config here if it is being used by any of the pipeline components.
        # add_dataset_to_config(name=name, parent=parent)
        
        def wrapper(*args, **kwargs):
            
            func_output = func(*args, **kwargs)

            if not upload:
                dataset = None
            else:
                dataset = func_output


            add_dataset_to_config(name=name, parent=parent, func=func)

            pipeline = create_pipeline()


            dataset_exists_in_remote, dataset_hash, dataset_version = register(dataset=dataset, name=name, pipeline=pipeline)

            #Uncomment this if there any components that depend on dataset version, or dataset hash
            # if dataset_exists_in_remote:
            #     add_dataset_to_config(name=name, hash=dataset_hash, version=dataset_version, parent=parent, func=func)

            return func_output

        return wrapper

        
    return decorator