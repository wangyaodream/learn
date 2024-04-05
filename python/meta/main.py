from collections import abc

from osconfeed import load


class FrozenJSON:
    """
    一个只读接口，使用属性表示法访问JSON类对象
    """

    def __init__(self, mapping):
        self.__data = dict(mapping)

    def __getattr__(self, name):
        if hasattr(self.__data, name):
            return getattr(self.__data, name)
        else:
            return FrozenJSON.build(self.__data[name])

    @classmethod
    def build(cls, obj):
        if isinstance(obj, abc.Mapping):
            return cls(obj)
        elif isinstance(obj, abc.MutableSequence):  # 在JSON中只有字典和列表是集合类型，上一个判断已经排除字典可能
            return [cls.build(item) for item in obj]
        else:
            return obj


if __name__ == '__main__':
    # 使用动态属性访问JSON类数据
    raw_feed = load()
    feed = FrozenJSON(raw_feed)
    print(len(feed.Schedule.speakers))
    # 通过feed对象来访问数据
    for key, value in sorted(feed.Schedule.items()):
        print('{:3} {}'.format(len(value), key))


