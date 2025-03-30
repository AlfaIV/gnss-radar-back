from typing import List

from pydantic import BaseModel


class RadarPositionGeograthRequest(BaseModel):
    satellites_name: list[str]
    radar_latitude: float
    radar_longitude: float
    radar_height: float # В километрах
    inspection_time: int # Время в UTC

class SatellitesTimeRequest(BaseModel):
    satellites_name: list[str]
    radar_latitude: float
    radar_longitude: float
    radar_height: float # В километрах
    begin_time: int # Время в UTC
    end_time: int # Время в UTC

class SatellitePosition(BaseModel):
    Group: str
    Name: str
    Azimuth: float
    Elevation: float
    Range: float


class Ephemeris(BaseModel):
    Group: str
    Name: str
    Longitude: float
    Latitude: float
    Height: float


class VisionTime(BaseModel):
    Begin: float
    End: float

class SatelliteTime(BaseModel):
    Group: str
    Name: str
    Time: List[VisionTime]


class SatellitesPositionResponce(BaseModel):
    Satellites: List[SatellitePosition]
    # Ephemerises: List[Ephemeris]


class SatellitesTimeResponce(BaseModel):
    Satellites: List[SatelliteTime]