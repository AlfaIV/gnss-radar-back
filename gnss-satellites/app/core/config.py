import os

from dotenv import load_dotenv
from pydantic_settings import BaseSettings

load_dotenv()

ENV: str = ""


class Configs(BaseSettings):
    ENV: str = os.getenv("ENV", "dev")
    API: str = "/api"
    API_V1_STR: str = "/api/v1"
    PROJECT_NAME: str = "measurement-services"

    PROJECT_ROOT: str = os.path.dirname(
        os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    )

    DATETIME_FORMAT: str = "%Y-%m-%dT%H:%M:%S"
    DATE_FORMAT: str = "%Y-%m-%d"

    TLE_PATH: str = os.path.join(PROJECT_ROOT, "app", "services", "tle", "gps.tle")

    TLE_BUCKET:str = 'ephemeris'
    S3_ENDPOINT_URL: str = "http://localhost:9000"
    MINIO_ROOT_USER: str = "admin"
    MINIO_ROOT_PASSWORD: str = "strongpassword"
    REGION_NAME: str = "us-east-1"

    class Config:
        case_sensitive = True

configs = Configs()

if ENV == "prod":
    pass
