from pureml.schema import BackendSchema, StorageSchema


def set_backend(backend_url: str) -> None:
    """Set the backend url.

    Args:
        backend_url (str): The url of the backend.
    """
    backend: BackendSchema = BackendSchema().get_instance()
    backend.BASE_URL = backend_url


def set_storage(storage: str) -> None:
    """Set the storage key.

    Args:
        storage (str): The storage key to reference storage details added in secret.
    """
    storage: StorageSchema = StorageSchema().get_instance()
    storage.STORAGE = storage
