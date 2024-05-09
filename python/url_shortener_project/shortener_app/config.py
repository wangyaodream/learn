from functools import lru_cache

from pydantic_settings import BaseSettings



class Settings(BaseSettings):
    env_name: str = "Local"
    base_url: str = "http://192.168.240.10:8000"
    db_url: str = "sqlite:///shortener.db"

    # 读取.env文件
    class Config:
        env_file = ".env"


@lru_cache
def get_settings() -> Settings:
    settings = Settings()
    print(f"Loading setting for: {settings.env_name}")
    return settings
