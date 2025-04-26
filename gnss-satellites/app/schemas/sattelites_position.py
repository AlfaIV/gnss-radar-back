from typing import List, Optional
from pydantic import BaseModel, field_validator
import math

class RadarPositionGeograthRequest(BaseModel):
    satellites_name: List[str] | None = None
    radar_latitude: float
    radar_longitude: float
    radar_height: float # В километрах
    inspection_time: str # 2025-04-17T15:00:00
    tle_file: str

class SatellitesTimeRequest(BaseModel):
    satellites_name: List[str] | None = None
    radar_latitude: float
    radar_longitude: float
    radar_height: float # В километрах
    begin_time: str # Время в UTC
    end_time: str # Время в UTC
    tle_file: str

class SatellitePosition(BaseModel):
    Group: str
    Name: str
    Azimuth: float
    Elevation: float
    Range: float
    
    @field_validator('Group')
    def check_group(cls, v):
        if v is None:
            print(f"Value {v} is None")
            return "Unidentified"
        return v
    
    @field_validator('Azimuth', 'Elevation', 'Range')
    def check_float_values(cls, v):
        if math.isnan(v) or math.isinf(v):
            # raise ValueError(f"Value {v} is not JSON compliant")
            print(f"Value {v} is not JSON compliant")
            return 0
        return v


class Ephemeris(BaseModel):
    Group: str
    Name: str
    Longitude: float
    Latitude: float
    Height: float


class VisionTime(BaseModel):
    startDatetime: str
    endDatetime: str

class SatelliteTime(BaseModel):
    name: Optional[str] = "Unidentified"
    group: str
    intervals: List[VisionTime]

    @field_validator('group', 'name', mode='before')
    def validate_fields(cls, v):
        if v is None:
            return "Unidentified"
        return v


class SatellitesPositionResponce(BaseModel):
    Satellites: List[SatellitePosition]
    # Ephemerises: List[Ephemeris]


class SatellitesTimeResponce(BaseModel):
    Satellites: List[SatelliteTime]