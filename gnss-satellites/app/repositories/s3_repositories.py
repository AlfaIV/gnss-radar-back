from typing import Callable
from app.core.config import configs
import boto3
from botocore.exceptions import NoCredentialsError, EndpointConnectionError, ClientError
from app.models.s3_models import S3_tle_model
from fastapi import HTTPException, status

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
    
    def get_tle(self, tle_name: str) -> S3_tle_model:
        # tle_name = 'gps_tle'
        encoding = 'utf-8'
        try:
            with self._get_session() as s3:
                response = s3.get_object(
                    Bucket = configs.TLE_BUCKET,
                    Key = tle_name
                )
                file_content = response['Body'].read()  
                text_content = file_content.decode(encoding)

                return S3_tle_model(
                    tle_file = text_content,
                    tle_name = tle_name,
                )
        except ClientError as e:
            raise HTTPException(
                status_code=status.HTTP_204_NO_CONTENT,
                detail=f"Error deceptions: {e}",
                headers={"X-Error": "Custom header", "Error-type": "Get TLE from S3"},
            )
        