
def spam():
    pass

def grok():
    pass

def foo():
    # 没有放在__all__中，所以不会隐性的导入
    print("I'm function foo")

blah = 42


__all__ = ["spam", "grok"]
