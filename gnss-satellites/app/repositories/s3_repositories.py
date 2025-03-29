from typing import Callable
import boto3
from botocore.exceptions import NoCredentialsError, EndpointConnectionError

class S3Repository:
    def __init__(self, get_session: Callable[[], boto3.client]):
        self._get_session = get_session
    
    def check_s3_connection(self) -> bool:
        """Проверяет соединение с S3"""
        try:
            with self._get_session() as s3:
                s3.list_buckets()
                return True
        except NoCredentialsError:
            print("❌ Invalid credentials")
        except EndpointConnectionError:
            print("❌ Connection failed")
        except Exception as e:
            print(f"❌ S3 error: {e}")
        return False