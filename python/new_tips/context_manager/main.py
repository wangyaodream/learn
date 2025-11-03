import logging
from dataclasses import dataclass
from typing import Any

@dataclass
class AppContext:
    user_id: int
    logger: logging.Logger
    config: dict[str, Any]

