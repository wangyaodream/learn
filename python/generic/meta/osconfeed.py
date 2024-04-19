from urllib.request import urlopen
import warnings
import os
import json


URL = 'https://raw.githubusercontent.com/AllenDowney/fluent-python-notebooks/master/19-dyn-attr-prop/oscon/data/osconfeed.json'
JSON = 'data/osconfeed.json'


def load():
    if not os.path.exists(JSON):
        msg = 'downloading {} to {}'.format(URL, JSON)
        warnings.warn(msg)
        with urlopen(URL) as remote, open(JSON, 'wb') as local:
            local.write(remote.read())

    with open(JSON, 'r') as f:
        return json.load(f)


if __name__ == '__main__':
    feed = load()
    # 查看Schedule键中的四个记录集合
    print(sorted(feed['Schedule'].keys()))

    for key, val in sorted(feed['Schedule'].items()):
        # 这里的:3是为了输出结果对其
        print('{:3}: {}'.format(len(val), key))
