
import json
import os
from urllib.parse import urljoin

import numpy as np
import requests
from joblib import Parallel, delayed
from PIL import Image
from pureml.utils.constants import BASE_URL, PATH_FIGURE_DIR
from pureml.utils.pipeline import add_figures_to_config
from rich import print

from . import get_org_id, get_token


def save_images(figure):
    os.makedirs(PATH_FIGURE_DIR, exist_ok=True)
    figure_paths = {}
    for figure_key, figure_value in figure.items():
        save_name = os.path.join(PATH_FIGURE_DIR, '.'.join([figure_key, 'png']))

        canvas = figure_value.canvas
        canvas.draw()
        data = np.frombuffer(canvas.tostring_rgb(), dtype=np.uint8)
        rgb_array = data.reshape(canvas.get_width_height()[::-1] + (3,))

        data = Image.fromarray(rgb_array)        
        data.save(save_name)

        figure_paths[figure_key] = save_name
    
    return figure_paths



def post_figures(figure_paths, model_name: str, model_branch:str, model_version:str):
    user_token = get_token()
    org_id = get_org_id()
    
    # print('figure_paths', figure_paths)

    url = '/org/{}/model/{}/branch/{}/version/{}/log'.format(org_id, model_name, model_branch, model_version)
    url = urljoin(BASE_URL, url)

    headers = {
        'Authorization': 'Bearer {}'.format(user_token)
    }

    files = {}
    for file_name, file_path in figure_paths.items():
        
        if os.path.isfile(file_path):
            files[file_name] = open(file_path, 'rb')
        else:
            print('[bold red] figure', file_name,'doesnot exist at the given path')

    
    data = {'name_path_mapping' : figure_paths,
            'model_name': model_name,
            'model_version': model_branch,
            'model_version': model_version}

    data = json.dumps(data)

    # try:

    response = requests.post(url, data=data, files=files, headers=headers)
    

    if response.ok:
        print(f"[bold green]Figures have been registered!")

    else:
        print(f"[bold red]Figures have not been registered!")

    return response
    # except Exception as e:
    #     return


def add(figure: dict=None, model_name: str=None, model_branch:str=None, model_version:str='latest', file_paths:dict=None) -> str:    
    '''`add` function takes in the path of the figure, name of the figure and the model name and
    registers the figure
    
    Parameters
    ----------
    figure : dict
        Key is the figure name, value is the matplotlib figure object
    name : str
        The name of the figure.
    model_name : str
        The name of the model you want to add figure to.
    model_version: str
        The version of the model
    
    Returns
    -------
        The response is a JSON object
    
    '''
    # print('file_paths', file_paths)
    # print('figure', figure)

    if file_paths is None:
        file_paths = save_images(figure)
            # print('figre paths', figure_paths)
        add_figures_to_config(values=file_paths, model_name=model_name, model_branch=model_branch, model_version=model_version)


    if model_name is not None and model_version is not None and model_version is not None:
        response = post_figures(figure_paths=file_paths, model_name=model_name, model_branch=model_branch, model_version=model_version)
        
    
        # print(response.text)

    # return response.text


# def fetch(model_name: str, model_version:str='latest', name:str = ''):
#     '''It fetches the figure from the server and stores it in the local directory
    
#     Parameters
#     ----------
#     model_name : str
#         The name of the model you want to fetch the figure from.
#     model_version: str
#         The version of the model
#     name : str
#         The name of the figure to be fetched. If not specified, all figures will be fetched.
    
#     Returns
#     -------
#         The response text is being returned.
    
#     '''

#     user_token = get_token()
#     org_id = get_org_id()


#     def fetch_figure(figure_details: dict):

#         url = figure_details['location']
#         file_path_temp = figure_details['path']
#         file_name = file_path_temp.split(os.path.sep)[-1]
#         save_path = os.path.join(PATH_FIGURE_DIR, file_name)
#         print('save path', save_path)

#         name_fetched = figure_details['figure']


#         headers = {
#             'Content-Type': 'application/x-www-form-urlencoded',
#             'Authorization': 'Bearer {}'.format(user_token)
#         }
        
#         print('figure url', url)

#         # response = requests.get(url, headers=headers)
#         response = requests.get(url)

#         print(response.status_code)

#         if response.status_code == 200:
#             print('[bold green] figure {} has been fetched'.format(name_fetched))

#             save_dir = os.path.dirname(save_path)

#             os.makedirs(save_dir, exist_ok=True)

#             figure_bytes = response.content

#             open(save_path, 'wb').write(figure_bytes)


#             print('[bold green] figure {} has been stored at {}'.format(name_fetched, save_path))
            
#             return response.text
#         else:
#             print('[bold red] Unable to fetch the figure')

#             return response.text


#     figure_details = details(model_name=model_name, name=name, model_version=model_version)

#     if figure_details is None:
#         return

#     if type(figure_details) is dict:

#         res_text = fetch_figure(figure_details)

#     elif type(figure_details) is list:
#         res_text = Parallel(n_jobs=-1)(delayed(fetch_figure)(art_det) for art_det in figure_details)


#     return res_text
    


# def delete(name:str, model_name:str,  model_version:str='latest') -> str:
#     '''`delete()` deletes an figure from a model
    
#     Parameters
#     ----------
#     name : str
#         The name of the figure you want to delete.
#     model_name : str
#         The name of the model you want to delete the figure from
#     model_version: str
#         The version of the model
    
#     '''

#     user_token = get_token()
#     org_id = get_org_id()

#     url_path_1 = '{}/project/{}/model/{}/{}/figure/{}/delete'.format(org_id, project_id, model_name, model_version, name)
#     url = urljoin(BASE_URL, url_path_1)

    
#     headers = {
#         'Content-Type': 'application/x-www-form-urlencoded',
#         'Authorization': 'Bearer {}'.format(user_token)
#     }
    

#     # figure_details = details(model_name=model_name, figure=figure)

#     # if figure_details is None:
#     #     print('[bold red] Unable to find figure details')
#     #     return


#     response = requests.delete(url, headers=headers)


#     if response.status_code == 200:
#         print(f"[bold green]figure has been deleted")
        
#     else:
#         print(f"[bold red]Unable to delete figure")

#     return response.text


