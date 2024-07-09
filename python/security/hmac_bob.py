import hashlib
import hmac
import json


hmac_sha256 = hmac.new(b'shared_key', digestmod=hashlib.sha256)
message = b'from Bob to Alice'
hmac_sha256.update(message)
hash_value = hmac_sha256.hexdigest()


authenticated_msg = {
        "message": list(message),
        'hash_value': hash_value
        }

outbound_msg_to_alice = json.dumps(authenticated_msg)
