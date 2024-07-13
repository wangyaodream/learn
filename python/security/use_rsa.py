import os
import json
from cryptography.hazmat.backends import default_backend
from cryptography.hazmat.primitives import serialization, hashes
from cryptography.hazmat.primitives.asymmetric import rsa, padding


def generate_keys():
    # 生成密钥
    private_key = rsa.generate_private_key(
            public_exponent=65537,
            key_size=3072,
            backend=default_backend()
            )

    # 提取公钥
    public_key = private_key.public_key()


    # 密钥对序列化
    private_bytes = private_key.private_bytes(
        encoding=serialization.Encoding.PEM,
        format=serialization.PrivateFormat.PKCS8,
        encryption_algorithm=serialization.NoEncryption(),
    )

    with open('private_key.pem', 'xb') as private_file:
        private_file.write(private_bytes)


    public_bytes = public_key.public_bytes(
            encoding=serialization.Encoding.PEM,
            format=serialization.PublicFormat.SubjectPublicKeyInfo,
    )

    with open("public_key.pem", 'xb') as public_file:
        public_file.write(public_bytes)


def my_encrypt(msg):
    # 反序列化
    with open("private_key.pem", 'rb') as private_file:
        loaded_private_key = serialization.load_pem_private_key(
            private_file.read(),
            password=None,
            backend=default_backend()
        )
    with open("public_key.pem", 'rb') as public_file:
        loaded_public_key = serialization.load_pem_public_key(
            public_file.read(),
            backend=default_backend()
        )

    padding_config = padding.OAEP(
            mgf=padding.MGF1(algorithm=hashes.SHA256()),
            algorithm=hashes.SHA256(),
            label=None
    )

    plaintext = msg

    ciphertext = loaded_public_key.encrypt(
            plaintext=plaintext,
            padding=padding_config,
    )

    decrypted_by_private_key = loaded_private_key.decrypt(
        ciphertext=ciphertext,
        padding=padding_config,
    )

    return decrypted_by_private_key

def generate_signature(msg, private_key):
    padding_config = padding.PSS(
        mgf=padding.MGF1(hashes.SHA256()),
        salt_length=padding.PSS.MAX_LENGTH,
    )

    private_key = private_key
    signature = private_key.sign(
            msg,
            padding_config,
            hashes.SHA256()
    )

    signed_msg = {
            'message': list(msg),
            'signature': signature
            }

    outbound_msg_to_alice = json.dumps(signed_msg)

    return outbound_msg_to_alice


if __name__ == "__main__":
    if not os.path.exists("private_key.pem"):
        generate_keys()

    message = b'hello world'

    result = my_encrypt(message)

    print(result)
    

