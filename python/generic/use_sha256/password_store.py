import hashlib
import os


def hash_password(password):
    salt = os.urandom(16)  # 生成随机盐
    hashed_password = hashlib.sha256(salt + password.encode('utf-8')).hexdigest()
    return salt, hashed_password


def verify_password(stored_password, salt, provided_password):
    return stored_password == hashlib.sha256(salt + provided_password.encode("utf-8")).hexdigest()


password = "mypassword"
salt, hashed_password = hash_password(password)
print(f"Salt: {salt}")
print(f"Hashed Password: {hashed_password}")


# verify password
is_vaild = verify_password(hashed_password, salt, "mypassword")
print(is_vaild)
