import json
from typing import Any, Callable
from functools import wraps


type Data = dict[str, Any]
type ExportFn = Callable[[Data], None]

exports: dict[str, ExportFn] = {}

def register_exporter(format: str) -> Callable[[ExportFn], ExportFn]:
    """
    注册器的目的是将format注册到exports这个字典中
    """
    def decorator(fn: ExportFn) -> ExportFn:
        @wraps(fn)
        def wrapper(data: Data) -> None:
            return fn(data)
        # register
        exports[format] = wrapper
        return wrapper

    return decorator


@register_exporter("pdf")
def export_pdf(data: Data) -> None:
    print("Exporting data to PDF:", data)

@register_exporter("csv")
def export_csv(data: Data) -> None:
    print("Exporting data to CSV:", data)

@register_exporter("json")
def export_json(data: Data) -> None:
    print(json.dumps(data, indent=2))


@register_exporter("foo")
def export_bar(data: Data) -> None:
    print("Exporting foo data:", data)


def export_data(data: Data, format: str) -> None:
    exporter = exports.get(format)
    if exporter is None:
        raise ValueError(f"No exporter found for format: {format}")
    exporter(data) 

def main() -> None:
    sample_data: Data = {"name": "Alice", "age": 30}

    export_data(sample_data, "json")
    export_data(sample_data, "pdf")
    export_data(sample_data, "foo")


if __name__ == "__main__":
    main()

