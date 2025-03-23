from typing import List

from pydantic import BaseModel


class RadarPositionRequest(BaseModel):
    radar_x: float
    radar_y: float
    radar_z: float

class RadarPositionGeograthRequest(BaseModel):
    radar_latitude: float
    radar_longitude: float
    radar_height: float

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
    Ephemerises: List[Ephemeris]


class SatellitesTimeResponce(BaseModel):
    Satellites: List[SatelliteTime]