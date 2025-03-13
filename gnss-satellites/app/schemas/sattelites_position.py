from typing import List

from pydantic import BaseModel


class RadarPositionRequest(BaseModel):
    radar_x: int
    radar_y: int
    radar_z: int


class SatellitePosition(BaseModel):
    Group: str
    Name: str
    Azimuth: float
    Range: float


class Ephemeris(BaseModel):
    Group: str
    Name: str
    Longitude: float
    Latitude: float
    Height: float


class SatellitesPositionResponce(BaseModel):
    Satellites: List[SatellitePosition]
    Ephemerises: List[Ephemeris]
