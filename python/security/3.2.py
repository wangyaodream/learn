import hashlib
from pathlib import Path


def store(path, data, key):
    data_path = Path(path)

    hash_path = data_path.with_suffix(".hash")

    hash_value = hashlib.blake2b(data, key=key).hexdigest()

    with data_path.open(mode='x'), hash_path.open(mode='x'):
        data_path.write_bytes(data)
        hash_path.write_text(hash_value)


def is_modified(path, key):
    data_path = Path(path)
    hash_path = data_path.with_suffix('.hash')

    data = data_path.read_bytes()
    original_hash_value = hash_path.read_text()
    hash_value = hashlib.blake2b(data, key=key).hexdigest()

    return original_hash_value != hash_value

