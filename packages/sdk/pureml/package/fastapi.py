import os
import shutil
import uvicorn
from fastapi import FastAPI
from pureml.schema import FastAPISchema, PredictionSchema
from pureml.utils.package import process_input, process_output
from pureml.utils.version_utils import parse_version_label


prediction_schema = PredictionSchema()
fastapi_schema = FastAPISchema()


def get_predict_file(predict_path):

    # os.makedirs(PATH_PREDICT_DIR, exist_ok=True)

    if predict_path is None:
        predict_path = prediction_schema.PATH_PREDICT_USER
        print("Taking the default predict.py file path: ", predict_path)
    else:
        print("Taking the predict.py file path: ", predict_path)

    if os.path.exists(predict_path):
        shutil.copy(predict_path, prediction_schema.PATH_PREDICT)
    else:
        raise Exception(predict_path, "doesnot exists!!!")


def get_requirements_file(requirements_path):

    # os.makedirs(prediction_schema.paths.PATH_PREDICT_DIR, exist_ok=True)

    if requirements_path is None:
        requirements_path = prediction_schema.PATH_PREDICT_REQUIREMENTS_USER
        print("Taking the default requirements.txt file path: ", requirements_path)
    else:
        print("Taking the requirements.txt file path: ", requirements_path)

    if os.path.exists(requirements_path):
        shutil.copy(requirements_path, prediction_schema.PATH_PREDICT_REQUIREMENTS)
    else:
        raise Exception(requirements_path, "doesnot exists!!!")


def generate_api(label: str):
    name, branch, version = parse_version_label(label=label)

    api = """
from fastapi import FastAPI, Depends, Request, UploadFile, File
import uvicorn
import pureml
from predict import Predictor
import os
from dotenv import load_dotenv
import json
import shutil
from pureml.utils.package import process_input, process_output
from typing import Union, Optional
from pureml.utils.prediction import predict_request_with_json, predict_request_with_file

load_dotenv()

org_id = os.getenv('ORG_ID')
access_token = os.getenv('ACCESS_TOKEN')

pureml.login(org_id=org_id, access_token=access_token)

predictor = Predictor()
predictor.load_models()

# input_type, input_shape = process_input(input=predictor.input)
# output_type, output_shape = process_output(output=predictor.output)


# Create the app
app = FastAPI()     


@app.post('/predict')
async def predict(request:Request = None, file: Optional[UploadFile] = File(None)):
    if request is None and file is None:
        print('Error in data input format')
        predictions = json.dumps(None)
        return predictions
    

    if request is not None and file is None:
        print("Processing request")
        predictions = await predict_request_with_json(request=request, predictor=predictor)


    if file is not None:
        print("Processing uploaded file")
        predictions = await predict_request_with_file(file=file, predictor=predictor)
        
    return predictions

if __name__ == '__main__':
    uvicorn.run(app, host='{HOST}', port={PORT})""".format(
        HOST=fastapi_schema.API_IP_HOST, PORT=fastapi_schema.PORT_FASTAPI
    )

    return api


def create_fastapi_file(
    label,
    predict_path,
    requirements_path,
):
    fastapi_schema = FastAPISchema()

    get_predict_file(predict_path)

    get_requirements_file(requirements_path)

    api = generate_api(label=label)

    with open(fastapi_schema.PATH_FASTAPI_FILE, "w") as api_writer:
        api_writer.write(api)

    api_writer.close()

    print("FastAPI server files are created")


#     print("""
#           API sucessfully created. To run your API, please run the following command
# --> !python <api_name>
#           """)


def run(label, predict_path, requirements_path):

    create_fastapi_file(
        label, predict_path=predict_path, requirements_path=requirements_path
    )

    run_command = "python {api_path}".format(api_path=fastapi_schema.PATH_FASTAPI_FILE)

    os.system(run_command)

    # app = FastAPI()

    # uvicorn.run(app, host=fastapi_schema.API_IP_HOST, port=fastapi_schema.PORT_FASTAPI)
