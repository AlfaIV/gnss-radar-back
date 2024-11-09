package store

import (
	"context"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
	sq "github.com/Masterminds/squirrel"
)

const (
	gnssTable = "gnss_coords"
)

type IGnssStore interface {
	ListGnssCoords(ctx context.Context, params ListParams) ([]*ListResult, error)
	RinexList(ctx context.Context) ([]*model.RinexResults, error)
}

type GnssStore struct {
	storage *Storage
}

func NewGnssStore(storage *Storage) *GnssStore {
	return &GnssStore{
		storage: storage,
	}
}

type ListParams struct {
	X float64
	Y float64
	Z float64
}

type ListResult struct {
	ID            string  `db:"id"`
	SatelliteID   string  `db:"satellite_id"`
	SatelliteName string  `db:"satellite_name"`
	X             float64 `db:"x"`
	Y             float64 `db:"y"`
	Z             float64 `db:"z"`
}

func (g *GnssStore) ListGnssCoords(ctx context.Context, params ListParams) ([]*ListResult, error) {
	query := g.storage.Builder().
		Select("id, satellite_id, satellite_name, x, y, z").
		From(gnssTable)

	if params.X != 0 {
		query = query.Where(sq.Eq{"x": params.X})
	}
	if params.Y != 0 {
		query = query.Where(sq.Eq{"y": params.Y})
	}
	if params.Z != 0 {
		query = query.Where(sq.Eq{"z": params.Z})
	}

	var coords []*ListResult
	if err := g.storage.db.Selectx(ctx, &coords, query); err != nil {
		return nil, postgresError(err)
	}

	return coords, nil
}

func (g *GnssStore) RinexList(ctx context.Context) ([]*model.RinexResults, error) {
	jsonData := []*model.RinexResults{
		{
			Header: &model.Header{
				RinexVersion:       "2.11",
				FileType:           "OBSERVATION DATA",
				PgmRunByDate:       "CONVBIN 2.4.3 20241106 095923 UTC",
				Comments:           []string{"format: Javad GREIS", "log: ../tmp_rinex2_hourly/YARO_2024_11_06_08.jps"},
				MarkerName:         "",
				MarkerNumber:       "",
				ObserverAgency:     "",
				RecInfo:            "",
				AntInfo:            "",
				ApproxPositionXyz:  []float64{2626719.2512, 2195458.1382, 5363666.0981},
				AntennaDeltaHen:    []float64{0.0000, 0.0000, 0.0000},
				WavelengthFactL1L2: []int{1, 1},
				TypesOfObs:         []string{"C1", "L1", "P1", "P2", "L2", "C2", "C5", "L5"},
				Interval:           10.000,
				TimeOfFirstObs:     "2024 11 06 08 59 50.0000000 GPS",
				TimeOfLastObs:      "2024 11 06 08 59 50.0000000 GPS",
				EndOfHeader:        true,
			},
			Observations: []*model.Observation{
				{
					Time:      "2024 11 06 08 59 50.0000000",
					EpochFlag: 0,
					Satellites: []*model.Satellite{
						{
							SatelliteID:  "23R",
							Observations: []string{"21637955.781", "115748497.2511", "21637957.379"},
						},
						{
							SatelliteID:  "23R",
							Observations: []string{"19723101.973", "105246184.4061", "19723098.435"},
						},
						{
							SatelliteID:  "22G",
							Observations: []string{"22766526.105", "121529156.8331", "22766524.885", "22766542.459", "94522675.6331", "22766539.847"},
						},
						{
							SatelliteID:  "16G",
							Observations: []string{"22303210.999", "117204258.0021", "22303209.896", "22303216.932", "91328033.4211"},
						},
						{
							SatelliteID:  "11G",
							Observations: []string{"23210320.069", "121971045.1601", "23210320.441", "23210322.845", "95042373.3311", "23210322.281", "23210330.328", "91082273.6561"},
						},
						{
							SatelliteID:  "06S",
							Observations: []string{"22580005.523", "118658672.7071", "22580005.781", "22580013.192", "92461282.6061", "22580012.928", "22580016.610", "88608732.2241"},
						},
						{
							SatelliteID:  "36G",
							Observations: []string{"42698660.033", "224383954.4171", "42698671.734", "167558951.1841"},
						},
						{
							SatelliteID:  "31R",
							Observations: []string{"23697554.619", "124531426.1341", "23697555.768", "23697562.681", "97037457.3321", "23697564.734"},
						},
						{
							SatelliteID:  "09G",
							Observations: []string{"21704140.137", "115898995.3911", "21704139.952", "21704155.235", "90143723.5771", "21704154.731"},
						},
						{
							SatelliteID:  "04S",
							Observations: []string{"20198608.366", "106144377.4681", "20198607.665", "20198612.965", "82709895.0251", "20198613.097", "20198617.666", "79263647.6451"},
						},
						{
							SatelliteID:  "27G",
							Observations: []string{"42093623.760", "221203304.4561", "42093645.171", "165184229.9381"},
						},
						{
							SatelliteID:  "07G",
							Observations: []string{"23117956.159", "121485787.6041", "23117956.168", "23117965.291", "94664269.7691", "23117965.423"},
						},
						{
							SatelliteID:  "06R",
							Observations: []string{"43006100.412", "225998406.6231", "43006124.803", "168764996.4521"},
						},
						{
							SatelliteID:  "17G",
							Observations: []string{"23183933.938", "124148883.1081", "23183929.689", "23183941.609", "96560237.9341", "23183941.843"},
						},
						{
							SatelliteID:  "11G",
							Observations: []string{"22758365.505", "119595850.0051", "22758364.780", "22758377.695", "93191533.4211", "22758378.609", "22758382.941", "89308544.0321"},
						},
						{
							SatelliteID:  "06S",
							Observations: []string{"42183829.071", "221677774.9391"},
						},
						{
							SatelliteID:  "36G",
							Observations: []string{"24446368.591", "128466528.5021", "24446366.328", "24446370.693", "100103783.4071"},
						},
						{
							SatelliteID:  "21G",
							Observations: []string{"22186059.844", "118555649.8941", "22186058.975", "22186070.877", "92209976.1571", "22186070.898"},
						},
						{
							SatelliteID:  "17G",
							Observations: []string{"20159372.252", "105938230.3921", "20159371.667", "20159377.217", "82549266.1111", "20159377.804", "20159379.855", "79109712.9511"},
						},
						{
							SatelliteID:  "36G",
							Observations: []string{"43094771.706", "226463949.3301"},
						},
						{
							SatelliteID:  "19G",
							Observations: []string{"19609198.442", "104969558.3841", "19609199.252", "19609208.569", "81642989.0881", "19609208.989"},
						},
						{
							SatelliteID:  "20G",
							Observations: []string{"19833957.468", "105949631.2851", "19833957.537", "19833967.769", "82405321.5061", "19833967.832"},
						},
						{
							SatelliteID:  "24G",
							Observations: []string{"22492129.204", "118197052.4691", "22492128.520", "22492138.692", "92101631.6171", "22492137.919", "22492142.173", "88264066.4751"},
						},
						{
							SatelliteID:  "23R",
							Observations: []string{"23243078.307", "124334761.901", "23243080.966"},
						},
						{
							SatelliteID:  "06R",
							Observations: []string{"21633253.668", "115439046.769", "21633253.278"},
						},
						{
							SatelliteID:  "16G",
							Observations: []string{"23998309.120", "128419771.441", "23998309.072", "23998324.698", "99882054.881", "23998324.952"},
						},
						{
							SatelliteID:  "11G",
							Observations: []string{"21923664.733", "115209766.128", "21923663.791", "21923673.750", "89773889.039"},
						},
						{
							SatelliteID:  "16G",
							Observations: []string{"23424519.852", "123096630.383", "23424516.884", "23424524.127", "95919443.589", "23424522.428", "23424531.355", "91922797.244"},
						},
						{
							SatelliteID:  "23R",
							Observations: []string{"24453451.191", "128503533.158", "24453450.555", "24453466.978", "100132562.228", "24453469.047", "24453473.966", "95960366.409"},
						},
						{
							SatelliteID:  "06R",
							Observations: []string{"42736359.492", "224582024.345", "42736372.074", "167706855.657"},
						},
						{
							SatelliteID:  "09G",
							Observations: []string{"19674729.455", "105062138.974", "19674728.715", "19674739.645", "81715078.893", "19674740.188"},
						},
						{
							SatelliteID:  "23G",
							Observations: []string{"21459301.640", "112769325.037", "21459302.056", "21459312.855", "87872178.458", "21459310.843", "21459317.637", "84210833.102"},
						},
						{
							SatelliteID:  "24G",
							Observations: []string{"42140500.115", "221449640.742", "42140522.734", "165368224.536"},
						},
						{
							SatelliteID:  "23G",
							Observations: []string{"21459110.045", "112768551.155", "21459110.399", "21459115.217", "87871632.597", "21459115.391"},
						},
						{
							SatelliteID:  "36G",
							Observations: []string{"23600946.305", "124023891.606", "23600945.882", "23600958.488", "96642021.833", "23600957.607", "23600963.687", "92615276.176"},
						},
						{
							SatelliteID:  "06R",
							Observations: []string{"43044614.114", "226200813.895", "43044639.224", "168916147.294"},
						},
						{
							SatelliteID:  "20G",
							Observations: []string{"20826002.217", "111522298.111", "20826001.503", "20826012.125", "86739563.905", "20826011.753"},
						},
						{
							SatelliteID:  "24G",
							Observations: []string{"22444434.682", "120020508.181", "22444435.129", "22444449.165", "93349281.012", "22444449.372"},
						},
						{
							SatelliteID:  "24G",
							Observations: []string{"42219944.076", "221867534.076"},
						},
						{
							SatelliteID:  "24G",
							Observations: []string{"22459782.002", "118026955.097", "22459779.733", "22459789.012", "91969049.478"},
						},
						{
							SatelliteID:  "20G",
							Observations: []string{"20067559.460", "105455729.757", "20067558.644", "20067565.306", "82173286.724", "20067565.276", "20067568.705", "78749398.299"},
						},
						{
							SatelliteID:  "36G",
							Observations: []string{"43131931.362", "226659202.146"},
						},
						{
							SatelliteID:  "19G",
							Observations: []string{"19132570.585", "102418109.420", "19132570.750", "19132580.268", "79658524.723", "19132580.892"},
						},
						{
							SatelliteID:  "26G",
							Observations: []string{"22370985.665", "119250026.543", "22370984.373"},
						},
						{
							SatelliteID:  "20G",
							Observations: []string{"20677759.178", "110457042.088", "20677758.516", "20677771.362", "85911078.807", "20677771.886"},
						},
						{
							SatelliteID:  "26G",
							Observations: []string{"23932031.753", "125763737.830", "23932031.816", "23932044.647", "97997735.733", "23932045.849", "23932049.435", "93914496.482"},
						},
					},
				},
			},
		},
	}

	return jsonData, nil
}

func (g *GnssStore) List1(ctx context.Context, params ListParams) ([]*model.Gnss, error) {
	jsonData := []*model.Gnss{
		{
			ID:            "PC06",
			SatelliteID:   "PC06",
			SatelliteName: "PC06",
			Coordinates: &model.CoordsResults{
				X: "-16806.320344",
				Y: "29291.120310",
				Z: "-25355.710938",
			},
		},
		{
			ID:            "PC07",
			SatelliteID:   "PC07",
			SatelliteName: "PC07",
			Coordinates: &model.CoordsResults{
				X: "-6959.418476",
				Y: "39332.954409",
				Z: "-13000.851001",
			},
		},
		{
			ID:            "PC08",
			SatelliteID:   "PC08",
			SatelliteName: "PC08",
			Coordinates: &model.CoordsResults{
				X: "-1908.204600",
				Y: "21553.224987",
				Z: "36203.881809",
			},
		},
		{
			ID:            "PC09",
			SatelliteID:   "PC09",
			SatelliteName: "PC09",
			Coordinates: &model.CoordsResults{
				X: "-11202.586298",
				Y: "28046.331947",
				Z: "-29182.143554",
			},
		},
		{
			ID:            "PC10",
			SatelliteID:   "PC10",
			SatelliteName: "PC10",
			Coordinates: &model.CoordsResults{
				X: "-917.431406",
				Y: "41238.966109",
				Z: "-6711.991412",
			},
		},
		{
			ID:            "PC11",
			SatelliteID:   "PC11",
			SatelliteName: "PC11",
			Coordinates: &model.CoordsResults{
				X: "-16138.177056",
				Y: "-3913.891460",
				Z: "-22348.411693",
			},
		},
		{
			ID:            "PC12",
			SatelliteID:   "PC12",
			SatelliteName: "PC12",
			Coordinates: &model.CoordsResults{
				X: "-997.099233",
				Y: "-19759.345910",
				Z: "-19638.934483",
			},
		},
		{
			ID:            "PC13",
			SatelliteID:   "PC13",
			SatelliteName: "PC13",
			Coordinates: &model.CoordsResults{
				X: "5858.392549",
				Y: "25505.986419",
				Z: "33308.911170",
			},
		},
		{
			ID:            "PC14",
			SatelliteID:   "PC14",
			SatelliteName: "PC14",
			Coordinates: &model.CoordsResults{
				X: "-17706.605729",
				Y: "-14691.268566",
				Z: "15829.680477",
			},
		},
		{
			ID:            "PC16",
			SatelliteID:   "PC16",
			SatelliteName: "PC16",
			Coordinates: &model.CoordsResults{
				X: "-22387.055407",
				Y: "28560.640995",
				Z: "-21454.026667",
			},
		},
	}

	return jsonData, nil
}
