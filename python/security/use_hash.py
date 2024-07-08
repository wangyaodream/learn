import hashlib
from hashlib import sha256


named = hashlib.sha256()
generic = hashlib.new('sha256')


message = b'message'
hash_function = sha256(message)


# 以字节形式返回散列值
print(hash_function.digest())

# 以字符串形式返回散列值
print(hash_function.hexdigest())

# 不使用消息构造的散列函数
named.update(message)

print(named.hexdigest())
