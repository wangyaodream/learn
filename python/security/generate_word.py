from pathlib import Path
import secrets


# 将字典文件加载到内存中
words = Path('path/to/word').read_text().splitlines()

passphrase = ' '.join(secrets.choice(words) for i in range(4))

print(passphrase)
