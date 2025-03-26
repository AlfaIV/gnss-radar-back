import os
from datetime import datetime,timedelta

import numpy as np
from astropy import units as u
from astropy.coordinates import (
    ITRS,
    TEME,
    AltAz,
    CartesianDifferential,
    CartesianRepresentation,
    Distance,
    EarthLocation,
)
from astropy.time import Time
from sgp4.api import Satrec, SGP4_ERRORS

from app.core.config import configs
from app.entities.sattellites import TLE
from app.schemas.sattelites_position import (
    RadarPositionRequest,
    SatellitesPositionResponce,
    SatellitesTimeResponce,
    RadarPositionGeograthRequest,
    VisionTime,
    SatelliteTime,
    SatellitePosition,
)


class SatellitesPositions:
    def __init__(self):
        tle_file = configs.TLE_PATH

        self.TLE_array = []
        with open(tle_file, "r") as file:
            for line in file:
                words = line.split()
                if words[0] == "1":
                    self.TLE_array[-1].line1 = line
                elif words[0] == "2":
                    self.TLE_array[-1].line2 = line
                else:
                    self.TLE_array.append(TLE(line))

        self.satellites = []
        for tle in self.TLE_array:
            self.satellites.append(
                {"satrec": Satrec.twoline2rv(tle.line1, tle.line2), "tle": tle}
            )

    def get_sattelites_positions(
        self, radar: RadarPositionGeograthRequest
    ) -> SatellitesPositionResponce:
        current_time = Time.now()

        elevation_mask = 0

        satellite_positions = []

        for satellite in self.satellites:
            sattelite_props = self.get_sattelite_positions(
                current_time, satellite["satrec"], radar
            )
            name = satellite["tle"].name.strip()
            parts = name.split()
            grouping = parts[0]
            satellite_name = " ".join(parts[1:])

            if sattelite_props["Elevation"] > elevation_mask:
                satellite_positions.append(
                    SatellitePosition(
                        Group = grouping,
                        Name = satellite_name,
                        Azimuth = sattelite_props["Azimuth"],
                        Elevation = sattelite_props["Elevation"],
                        Range = sattelite_props["Range"],
                    )
                )

        return SatellitesPositionResponce(
            Satellites = satellite_positions,
        )

    def get_sattelite_positions(
        self, current_time: datetime, satellite: object, observer: RadarPositionGeograthRequest
    ) -> SatellitePosition:
        error_code, teme_p, teme_v = satellite.sgp4(current_time.jd1, current_time.jd2)
        if error_code != 0:
            raise RuntimeError(SGP4_ERRORS[error_code])

        teme_p = CartesianRepresentation(teme_p * u.km)
        teme_v = CartesianDifferential(teme_v * u.km / u.s)
        teme = TEME(teme_p.with_differentials(teme_v), obstime=current_time)

        itrs_geo = teme.transform_to(ITRS(obstime=current_time))
        location = itrs_geo.earth_location
        geo = location.geodetic

        observer_location = EarthLocation.from_geodetic(
            observer.radar_longitude * u.deg, observer.radar_latitude * u.deg, observer.radar_height * u.m
        )
        observer_itrs = observer_location.get_itrs(obstime=current_time)

        altaz_frame = AltAz(obstime=current_time, location=observer_location)
        satellite_altaz = itrs_geo.transform_to(altaz_frame)

        azimuth = satellite_altaz.az
        elevation = satellite_altaz.alt
        distance_value = np.sqrt(
            (itrs_geo.x - observer_itrs.x) ** 2
            + (itrs_geo.y - observer_itrs.y) ** 2
            + (itrs_geo.z - observer_itrs.z) ** 2
        )
        distance = Distance(value=distance_value, unit=u.m)

        return {
            "Azimuth": round(azimuth.degree, 2),
            "Range": round(distance.km, 2),
            "Elevation": round(elevation.degree, 2),
            "Longitude": round(geo.lon.degree, 2),
            "Latitude": round(geo.lat.degree, 2),
            "Height": round(geo.height.to(u.km).value, 2),
        }
    

    def get_sattelites_times_vison(
        self, radar: RadarPositionGeograthRequest
    ) -> SatellitesTimeResponce:
        current_time = datetime.now()
        radar_position = {
            "radar_latitude": radar.radar_latitude,
            "radar_longitude": radar.radar_longitude,
            "radar_height": radar.radar_height,
        }

        satellite_time = []

        for satellite in self.satellites:
            visibility_times = self.visibility_times(
                current_time, satellite["satrec"], radar_position
            )
            name = satellite["tle"].name.strip()
            parts = name.split()
            grouping = parts[0]
            satellite_name = " ".join(parts[1:])
            satellite_time.append(
                SatelliteTime(
                    Group=grouping,
                    Name=satellite_name,
                    Time=visibility_times,
                )
            )

        return SatellitesTimeResponce(Satellites=satellite_time)


    def visibility_times(
        self,
        day: datetime,
        satellite: object,
        observer: RadarPositionRequest,
        min_elevation: float = 0.0,  
        max_elevation: float = 360.0, 
        min_azimuth: float = 0.0,   
        max_azimuth: float = 360.0,
        time_step: timedelta = timedelta(minutes=10),  # Шаг по времени
    ) -> list:      

        latitude = observer["radar_latitude"] 
        longitude = observer["radar_longitude"] 
        height = observer["radar_height"]

        observer_location = EarthLocation(lat=latitude * u.deg, lon=longitude * u.deg, height=height * u.m)

        visibility_times = []
        current_time = day.replace(hour=0, minute=0, second=0, microsecond=0)
        end_time = day.replace(hour=23, minute=59, second=59, microsecond=999999)

        is_visible = False
        visibility_start = None

        while current_time <= end_time:
            astro_time = Time(current_time)
            
            error_code, teme_p, teme_v = satellite.sgp4(astro_time.jd1, astro_time.jd2)
            if error_code != 0:
                raise RuntimeError(SGP4_ERRORS[error_code])

            teme_p = CartesianRepresentation(teme_p * u.km)
            teme_v = CartesianDifferential(teme_v * u.km / u.s)
            teme = TEME(teme_p.with_differentials(teme_v), obstime=astro_time)

            itrs_geo = teme.transform_to(ITRS(obstime=astro_time))
            altaz_frame = AltAz(obstime=astro_time, location=observer_location)
            satellite_altaz = itrs_geo.transform_to(altaz_frame)

            azimuth = satellite_altaz.az.degree
            elevation = satellite_altaz.alt.degree

            if (
                min_elevation <= elevation <= max_elevation
                and min_azimuth <= azimuth <= max_azimuth
            ):
                if not is_visible:
                    
                    visibility_start = current_time
                    is_visible = True
            else:
                if is_visible:
                    
                    visibility_times.append(VisionTime(Begin=visibility_start.timestamp(), End=current_time.timestamp()))
                    is_visible = False

            current_time += time_step

        if is_visible:
            visibility_times.append(VisionTime(Begin=visibility_start.timestamp(), End=end_time.timestamp()))

        return visibility_times
