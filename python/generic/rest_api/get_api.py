import json

import requests


response = requests.get(
        'https://reststop.randomhouse.com/resources/authors?lastName=Grisham')

print(response.json())

