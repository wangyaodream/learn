import hashlib
import hmac
import json


authenticated_msg = json.loads(inbound_msg_from_bob)

message = bytes(authenticated_msg['message'])

hmac_sha256 = hmac.new(b'shared_key', digestmod=hashlib.sha256)
hmac_sha256.update(message)
hash_value = hmac_sha256.hexdigest()

if hash_value == authenticated_msg['hash_value']:
    print('trust message')
