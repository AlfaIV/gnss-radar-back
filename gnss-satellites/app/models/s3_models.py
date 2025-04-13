from pydantic import BaseModel

class S3_tle_model(BaseModel):
    tle_file: bytes | str
    tle_name: str
    