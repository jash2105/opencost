package customcost

import (
	"fmt"

	"github.com/opencost/opencost/core/pkg/filter/ast"
	"github.com/opencost/opencost/core/pkg/filter/matcher"
	"github.com/opencost/opencost/core/pkg/filter/transform"
)

func NewCustomCostMatchCompiler() *matcher.MatchCompiler[*CustomCost] {
	passes := []transform.CompilerPass{
		transform.UnallocatedReplacementPass(),
	}

	return matcher.NewMatchCompiler(
		customCostFieldMap,
		customCostSliceFieldMap,
		customCostMapFieldMap,
		passes...,
	)
}

// Maps fields from a custom cost to a string value based on an identifier
func customCostFieldMap(cc *CustomCost, identifier ast.Identifier) (string, error) {
	if cc == nil {
		return "", fmt.Errorf("cannot map to nil custom cost")
	}
	if identifier.Field == nil {
		return "", fmt.Errorf("cannot map field from identifier with nil field")
	}
	switch CustomCostProperty(identifier.Field.Name) {
	case CustomCostZoneProp:
		return cc.Zone, nil
	case CustomCostAccountNameProp:
		return cc.AccountName, nil
	case CustomCostChargeCategoryProp:
		return cc.ChargeCategory, nil
	case CustomCostResourceTypeProp:
		return cc.ResourceType, nil
	case CustomCostDomainProp:
		return cc.Domain, nil
	}

	return "", fmt.Errorf("failed to find string identifier on CustomCost: %s", identifier.Field.Name)
}

// Maps slice fields from an asset to a []string value based on an identifier
func customCostSliceFieldMap(cc *CustomCost, identifier ast.Identifier) ([]string, error) {
	return nil, fmt.Errorf("custom costs have no slice fields")
}

// Maps map fields from a custom cost to a map[string]string value based on an identifier
func customCostMapFieldMap(cc *CustomCost, identifier ast.Identifier) (map[string]string, error) {
	return nil, fmt.Errorf("custom costs have no map fields")
}
