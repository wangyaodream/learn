import common
# from common import spam

"""
当在common包的__init__.py文件中使用from .spam import Spam时，下面这句是生效的。
也就是在import common时 同时也是执行了制定的导入操作才可以直接使用Spam，而Spam是在common.spam下的。
"""
# a = spam.Spam()
a = common.Spam()


