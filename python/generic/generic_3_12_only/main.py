from typing import Generic, TypeAlias, TypeVar


T = TypeVar("T")

# A non-generic type alias
IntOrStr = int | str

# A generic type alias
ListOrSet: TypeAlias = list[T] | set[T]


class Box(Generic[T]):
    def __init__(self, item: T):
        self.item = item

    def get_item(self) -> T:
        return self.item

    def set_item(self, new_item: T) -> None:
        self.item = new_item
        

def process_number(numbers: list[int]) -> list[int]:
    return [number + 1 for number in numbers]


def process_elements(elements: list[T]) -> list[T]:
    return [element for index, element in enumerate(elements) if index % 2 == 1]

def main() -> None:
    my_list = [1,2,3,4,5]
    processed = process_elements(my_list)
    print(processed)

    my_str_list = list("abcdef")
    processed = process_elements(my_str_list)
    print(processed)

    # For Integer 这里会让lsp报错，但是解释器本身并不会报错
    int_box = Box(123)
    int_item = int_box.get_item()
    int_box.set_item("3456")
    int_item = int_box.get_item()
    print(int_item)

    # For strings
    str_box = Box("hello, world")
    str_item = str_box.get_item()
    print(str_item)

    # For list
    list_box = Box([1,2,3,4,5])
    list_item = list_box.get_item()
    print(list_item)


if __name__ == "__main__":
    main()


