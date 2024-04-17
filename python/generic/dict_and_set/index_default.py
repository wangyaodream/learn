import sys
import re
import collections

WORD_RE = re.compile(r'\w+')

# 使用collections.defaultdict来给未查找到的键赋值,这里的list是一个构造方法
# 它将作为default_factory来创建一个defaultdict
index = collections.defaultdict(list)
with open(sys.argv[1], encoding="utf-8") as fp:
    for line_no, line in enumerate(fp, 1):
        for match in WORD_RE.finditer(line):
            word = match.group()
            column_no = match.start() + 1
            location = (line_no, column_no)
            index[word].append(location)

for word in sorted(index, key=str.upper):
    print(word, index[word])

