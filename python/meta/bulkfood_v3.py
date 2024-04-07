class Quantity:

    def __init__(self, storage_name):
        self.storage_name = storage_name

    def __set__(self, instance, value):
        if value > 0:
            instance.__dict__[self.storage_name] = value
        else:
            raise ValueError("value must be > 0")


class LineItem:
    # 绑定描述符，由于Quantity类实现了__set__方法，所以它是一个描述符
    # weight = Quantity('weight')
    # price = Quantity('price')
    weight = Quantity()
    price = Quantity()

    def __init__(self, description, weight, price):
        self.description = description
        self.weight = weight
        self.price = price

    def subtotal(self):
        return self.weight * self.price


if __name__ == '__main__':
    # 此时会raise错误,weight不能为0，这个过程由Quantity类的__set__实现
    n = LineItem("test", 0, 100.0)
    print(n.subtotal())
