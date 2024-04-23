package coinbasemanager

import (
	"strconv"
	"testing"

	"github.com/dmgcoin/dmgcoin/domain/consensus/model/externalapi"
	"github.com/dmgcoin/dmgcoin/domain/consensus/utils/constants"
	"github.com/dmgcoin/dmgcoin/domain/dagconfig"
)

func TestCalcDeflationaryPeriodBlockSubsidy(t *testing.T) {
	const secondsPerMonth = 2629800
	const secondsPerHalving = secondsPerMonth * 12
	const deflationaryPhaseDaaScore = 86400 //secondsPerMonth * 6
	const deflationaryPhaseBaseSubsidy = 0.0176 * constants.SompiPerDmgcoin
	coinbaseManagerInterface := New(
		nil,
		0,
		0,
		0,
		&externalapi.DomainHash{},
		deflationaryPhaseDaaScore,
		deflationaryPhaseBaseSubsidy,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil)
	coinbaseManagerInstance := coinbaseManagerInterface.(*coinbaseManager)

	tests := []struct {
		name                 string
		blockDaaScore        uint64
		expectedBlockSubsidy uint64
	}{
		{
			name:                 "start of deflationary phase",
			blockDaaScore:        deflationaryPhaseDaaScore,
			expectedBlockSubsidy: deflationaryPhaseBaseSubsidy,
		},
		{
			name:                 "after one halving",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving,
			expectedBlockSubsidy: deflationaryPhaseBaseSubsidy / 2,
		},
		{
			name:                 "after two halvings",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*2,
			expectedBlockSubsidy: deflationaryPhaseBaseSubsidy / 4,
		},
		{
			name:                 "after five halvings",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*5,
			expectedBlockSubsidy: deflationaryPhaseBaseSubsidy / 32,
		},
		// {
		// 	name:                 "after 32 halvings",
		// 	blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*32,
		// 	expectedBlockSubsidy: deflationaryPhaseBaseSubsidy / 4294967296,
		// },
		// {
		// 	name:                 "just before subsidy depleted",
		// 	blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*35,
		// 	expectedBlockSubsidy: 1,
		// },
		// {
		// 	name:                 "after subsidy depleted",
		// 	blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*36,
		// 	expectedBlockSubsidy: 0,
		// },
	}

	for _, test := range tests {
		blockSubsidy := coinbaseManagerInstance.calcDeflationaryPeriodBlockSubsidy(test.blockDaaScore)
		if blockSubsidy != test.expectedBlockSubsidy {
			t.Errorf("TestCalcDeflationaryPeriodBlockSubsidy: test '%s' failed. Want: %d, got: %d",
				test.name, test.expectedBlockSubsidy, blockSubsidy)
		}
	}
}

func TestBuildSubsidyTable(t *testing.T) {
	const deflationaryPhaseDaaScore = 86400
	const secondsPerMonth = 2629800

	deflationaryPhaseBaseSubsidy := dagconfig.MainnetParams.DeflationaryPhaseBaseSubsidy
	if deflationaryPhaseBaseSubsidy != 0.0176*constants.SompiPerDmgcoin {
		t.Errorf("TestBuildSubsidyTable: table generation function was not updated to reflect "+
			"the new base subsidy %d. Please fix the constant above and replace subsidyByDeflationaryMonthTable "+
			"in coinbasemanager.go with the printed table", deflationaryPhaseBaseSubsidy)
	}
	coinbaseManagerInterface := New(
		nil,
		0,
		0,
		0,
		&externalapi.DomainHash{},
		0,
		deflationaryPhaseBaseSubsidy,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil)
	coinbaseManagerInstance := coinbaseManagerInterface.(*coinbaseManager)

	var subsidyTable []uint64
	for M := uint64(0); ; M++ {
		subsidy := coinbaseManagerInstance.calcDeflationaryPeriodBlockSubsidyFloatCalc(M)
		subsidyTable = append(subsidyTable, subsidy)
		if subsidy == 0 {
			break
		}
	}

	tableStr := "\n{\t"
	for i := 0; i < len(subsidyTable); i++ {
		subsidy := float64(subsidyTable[i]) / float64(constants.SompiPerDmgcoin)
		tableStr += strconv.Itoa(deflationaryPhaseDaaScore+(secondsPerMonth*i)) + " : " + strconv.FormatFloat(subsidy, 'f', 10, 64) + ", "
		//tableStr += strconv.FormatUint(subsidyTable[i], 10) + ", "
		if (i+1)%25 == 0 {
			tableStr += "\n\t"
		}
	}
	tableStr += "\n}"
	t.Logf(tableStr)
}
