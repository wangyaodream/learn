from abc import ABC, abstractmethod

class Shape:

    @abstractmethod
    def area(self):
        pass

class Circle(Shape):

    def __init__(self, radius):
        self.radius = radius

    def area(self):
        return 3.14 * self.radius ** 2

    def __str__(self):
        return f"Square({self.radius})"

class Square(Shape):
    def __init__(self, side):
        self.side = side

    def area(self):
        return self.side ** 2

    def __repr__(self):
        return f"Square({self.side})"


class Triangle(Shape):
    def __init__(self, base, height):
        self.base = base
        self.height = height

    def area(self):
        return self.base * self.height * 0.5

    def __repr__(self):
        return f"Square({self.base}, {self.height})"
# circle是一个shape，但是它并不是square也不是triangle，仅仅是一个circle
circle = Circle(10)
square = Square(5)


shapes = [Circle(6), Square(7), Triangle(3, 4)]

for shape in shapes:
    print(shape)

    print("area: ", shape.area())
