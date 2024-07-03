from cryptography.hazmat.primitives.asymmetric import rsa, padding
from cryptography.hazmat.primitives import serialization, hashes

# generate rsa keys
def generate_rsa_keys():
    private_key = rsa.generate_private_key(
        public_exponent=65537,
        key_size=2048
    )
    public_key = private_key.public_key()
    return private_key, public_key


def encrypt_message(public_key, message):
    encrypted_message = public_key.encrypt(
        message.encode('utf-8'),
        padding.OAEP(
            mgf=padding.MGF1(algorithm=hashes.SHA256()),
            algorithm=hashes.SHA256(),
            label=None
        )
    )

    return encrypted_message


def decrypt_message(private_key, encrypted_message):
    decrypted_message = private_key.decrypt(
        encrypted_message,
        padding.OAEP(
            mgf=padding.MGF1(algorithm=hashes.SHA256()),
            algorithm=hashes.SHA256(),
            label=None
        )
    )

    return decrypted_message.decode("utf-8")


def main():
    private_key, public_key = generate_rsa_keys()

    original_message = "Hello, world!"

    encrypted_message = encrypt_message(public_key, original_message)
    print("Encrypted message:", encrypted_message)

    decrypted_message = decrypt_message(private_key, encrypted_message)
    print("Decrypted message:", decrypted_message)



if __name__ == "__main__":
    main()

