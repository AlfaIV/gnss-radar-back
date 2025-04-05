from contextlib import contextmanager
from typing import Iterator
import boto3
from app.core.config import Configs, configs

class S3Database:
    def __init__(self, config: Configs = None):
        self.config = config or configs
    
    def create_session(self):
        return boto3.client(
            's3',
            endpoint_url=self.config.S3_ENDPOINT_URL,
            aws_access_key_id=self.config.MINIO_ROOT_USER,
            aws_secret_access_key=self.config.MINIO_ROOT_PASSWORD,
            region_name=self.config.REGION_NAME
        )
    
    @contextmanager
    def session(self) -> Iterator[boto3.client]:
        session = self.create_session()
        try:
            yield session
        finally:
            session.close()