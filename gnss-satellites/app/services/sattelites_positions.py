from datetime import datetime,timedelta

import numpy as np
from datetime import datetime, timedelta, timezone
from skyfield.api import load, Topos, EarthSatellite,utc
from skyfield.toposlib import wgs84
from skyfield.positionlib import Geocentric

from app.core.config import configs
from app.entities.sattellites import TLE
from app.schemas.sattelites_position import (
    SatellitesTimeRequest,
    SatellitesPositionResponce,
    SatellitesTimeResponce,
    RadarPositionGeograthRequest,
    VisionTime,
    SatelliteTime,
    SatellitePosition,
)

from app.repositories import S3Repository


class SatellitesPositions:
    def __init__(self, s3_repository: S3Repository):

        self.s3_repository =  s3_repository

        print(f"Проверка S3: {self.s3_repository.check_s3_connection()}" )

        self.TLE_array = self.load_sattelites_tle()

        self.ts = load.timescale()
        self.time_step = timedelta(minutes=5)

        self.mask_visible = 0




    def load_sattelites_tle(self) -> list:

        #Тут надо поменять логику на работу с S3
        tle_file = configs.TLE_PATH
        TLE_array = []
        with open(tle_file, "r") as file:
            for line in file:
                words = line.split()
                if words[0] == "1":
                    TLE_array[-1].line1 = line
                elif words[0] == "2":
                    TLE_array[-1].line2 = line
                else:
                    grouping = words[0]
                    satellite_name = " ".join(words[1:])
                    TLE_array.append(TLE(satellite_name,grouping))
        return TLE_array


    def get_sattelites_positions(
        self, radar: RadarPositionGeograthRequest,
    ) -> SatellitesPositionResponce:

        satellite_positions = []

        for satellite in self.TLE_array:
            if radar.satellites_name:
                if satellite.name in radar.satellites_name:
                    sattelite_props = self.get_sattelite_positions(
                        radar.inspection_time, satellite, radar
                    )

                    satellite_positions.append(
                        SatellitePosition(
                            Group = satellite.group,
                            Name = satellite.name,
                            Azimuth = sattelite_props["Azimuth"],
                            Elevation = sattelite_props["Elevation"],
                            Range = sattelite_props["Range"],
                        )
                    )
            else:
                sattelite_props = self.get_sattelite_positions(
                    radar.inspection_time, satellite, radar
                )

                satellite_positions.append(
                    SatellitePosition(
                        Group = satellite.group,
                        Name = satellite.name,
                        Azimuth = sattelite_props["Azimuth"],
                        Elevation = sattelite_props["Elevation"],
                        Range = sattelite_props["Range"],
                    )
                )

        return SatellitesPositionResponce(
            Satellites = satellite_positions,
        )

    def get_sattelite_positions(
        self, current_time: datetime,  satellite: object, observer: RadarPositionGeograthRequest
    ) -> SatellitePosition:
        
    
        current_time = datetime.fromtimestamp(current_time, tz=timezone.utc)
        skyfield_time = self.ts.from_datetime(current_time)

        observer = wgs84.latlon(observer.radar_latitude, observer.radar_longitude, observer.radar_height)
    
        # Получаем положение спутника
        satellite = EarthSatellite(satellite.line1.strip(), satellite.line2.strip(), satellite.name.strip())
        position = satellite.at(skyfield_time)
        subpoint = wgs84.subpoint(position)

        difference = satellite - observer
        topocentric = difference.at(skyfield_time)
        
        # Получаем азимут и угол места
        elevation, azimuth, distance = topocentric.altaz()

        return {
            "Azimuth": round(azimuth.degrees, 2),
            "Range": round(distance.km, 2),
            "Elevation": round(elevation.degrees, 2),
            "Longitude": round(subpoint.longitude.degrees, 2),
            "Latitude": round(subpoint.latitude.degrees, 2),
            "Height": round(subpoint.elevation.km, 2),
        }
    

    def get_sattelites_times_vison(
        self, radar: SatellitesTimeRequest
    ) -> SatellitesTimeResponce:
        
        begin_time = datetime.fromtimestamp(radar.begin_time, tz=timezone.utc)

        end_time = datetime.fromtimestamp(radar.end_time, tz=timezone.utc)
        
        observer = wgs84.latlon(
            radar.radar_latitude,
            radar.radar_longitude,
            radar.radar_height
        )

        satellite_time = []
           

        for satellite in self.TLE_array:
            
            if radar.satellites_name:

                if radar.satellites_name.__contains__(satellite.name):

                    satellite_data = EarthSatellite(
                        satellite.line1.strip(),
                        satellite.line2.strip(),
                        satellite.name
                    )

                    visibility_times = self.visibility_times(
                        begin_time,end_time, satellite_data, observer
                    )
                    
                    satellite_time.append(
                        SatelliteTime(
                            Group = satellite.group,
                            Name = satellite.name,
                            Time=visibility_times,
                        )
                    )

            else:
                satellite_data = EarthSatellite(
                    satellite.line1.strip(),
                    satellite.line2.strip(),
                    satellite.name.strip()
                )

                visibility_times = self.visibility_times(
                    begin_time,end_time, satellite_data, observer
                )
                
                satellite_time.append(
                    SatelliteTime(
                        Group = satellite.group,
                        Name = satellite.name,
                        Time=visibility_times,
                    )
                )

        return SatellitesTimeResponce(Satellites=satellite_time)


    def visibility_times(
        self,
        start_time: datetime,
        end_time: datetime,
        satellite: EarthSatellite,
        observer: Geocentric,
        min_elevation = 0,
    ) -> list[VisionTime]:
        
        current_time = start_time
        visibility_periods = []
        in_visibility = False
        period_start = None
        
        while current_time <= end_time:
            skyfield_time = self.ts.from_datetime(current_time)
            
            difference = satellite - observer
            topocentric = difference.at(skyfield_time)
            elevation = topocentric.altaz()[0].degrees
            
            
            if elevation >= min_elevation:
                if not in_visibility:
                    
                    period_start = current_time
                    in_visibility = True
            else:
                if in_visibility:
                    
                    visibility_periods.append(
                        VisionTime(
                            Begin=period_start.timestamp(),
                            End=current_time.timestamp()
                        )
                    )
                    in_visibility = False
            
            current_time += self.time_step
        
        if in_visibility:
            visibility_periods.append(
                VisionTime(
                    Begin=period_start.timestamp(),
                    End=end_time.timestamp()
                )
            )
            
        return visibility_periods
