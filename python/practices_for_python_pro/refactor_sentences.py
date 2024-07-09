import re


def get_matches_for_pattern(pattern, string):
    matches = pattern.findall(string)
    return [match[0] for match in matches]


product_review = '''This is a fine milk,but the product
line apperas to be limited in available colors.I could only find white.'''

sentence_pattern = re.compile(r'(.*?\.)(\s|$)', re.DOTALL)
sentences = get_matches_for_pattern(sentence_pattern, product_review)


word_pattern = re.compile(r"([\w\-]+)([\s,.])?")
for sentence in sentences:
    words = get_matches_for_pattern(word_pattern, sentence)
    print(words)

