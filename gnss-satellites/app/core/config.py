from pydantic_settings import BaseSettings, SettingsConfigDict
from pydantic import Field
from pathlib import Path
import os

class Configs(BaseSettings):
    ENV: str = Field(default="dev", env="ENV")
    API: str = "/api"
    API_V1_STR: str = "/api/v1"
    PROJECT_NAME: str = "measurement-services"

    PROJECT_ROOT: Path = Path(__file__).parent.parent.parent
    TLE_PATH: Path = PROJECT_ROOT / "app" / "services" / "tle" / "gps.tle"

    DATETIME_FORMAT: str = "%Y-%m-%dT%H:%M:%S"

    TLE_PATH: str = os.path.join(PROJECT_ROOT, "app", "services", "tle", "gps.tle")

    TLE_BUCKET: str = Field(default="ephemeris", env="TLE_BUCKET")
    S3_ENDPOINT_URL: str = Field(..., env="S3_ENDPOINT_URL")
    MINIO_ROOT_USER: str = Field(..., env="MINIO_ROOT_USER")
    MINIO_ROOT_PASSWORD: str = Field(..., env="MINIO_ROOT_PASSWORD")
    REGION_NAME: str = Field(default="us-east-1", env="REGION_NAME")

    model_config = SettingsConfigDict(
        env_file=".env" if Path(".env").exists() else None,
        env_file_encoding="utf-8",
        extra="ignore"
    )

configs = Configs()
