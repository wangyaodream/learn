import hashlib
import time


def generate_unique_id(data):
    current_time = str(time.time()).encode('utf-8')
    data_bytes = data.encode('utf-8')
    unique_id = hashlib.sha256(data_bytes + current_time).hexdigest()
    return unique_id


data = 'user123'
unique_id = generate_unique_id(data)
print(unique_id)
