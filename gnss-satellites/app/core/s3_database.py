from contextlib import contextmanager
from typing import Iterator
import boto3
from pydantic import BaseModel

class S3Config(BaseModel):
    endpoint_url: str = "http://localhost:9000"
    aws_access_key_id: str = "admin"
    aws_secret_access_key: str = "strongpassword"
    region_name: str = "us-east-1"

class S3Database:
    def __init__(self, config: S3Config = None):
        self.config = config or S3Config()
    
    def create_session(self):
        """Создает новую сессию S3 (без контекста)"""
        return boto3.client(
            's3',
            endpoint_url=self.config.endpoint_url,
            aws_access_key_id=self.config.aws_access_key_id,
            aws_secret_access_key=self.config.aws_secret_access_key,
            region_name=self.config.region_name
        )
    
    @contextmanager
    def session(self) -> Iterator[boto3.client]:
        """Контекстный менеджер для сессии"""
        session = self.create_session()
        try:
            yield session
        finally:
            session.close()