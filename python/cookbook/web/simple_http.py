from urllib import request, parse


url = 'http://www.baidu.com/'


parms = {
    'name1': 'value1',
    'name2': 'value2'
}


r = request.urlopen(url)
print(r.read().decode('utf-8'))