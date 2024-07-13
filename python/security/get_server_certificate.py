import ssl


address = ('trojancow.top', 443)
certificate = ssl.get_server_certificate(address)
print(certificate)
