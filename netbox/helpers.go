package netbox

import (
	"github.com/netbox-community/go-netbox/netbox/models"
)


func expandTags(input []interface{}) []*models.NestedTag {
	if len(input) == 0 {
		return nil
	}

	results := make([]*models.NestedTag, 0)

	for _, item := range input {
		values := item.(map[string]interface{})
		result := &models.NestedTag{}

		if val, ok := values["id"]; ok {
			result.ID = int64(val.(int))
		}

		if val, ok := values["name"]; ok {
			name := val.(string)
			result.Name = &name
		}

		if val, ok := values["slug"]; ok {
			slug := val.(string)
			result.Slug = &slug
		}

		if val, ok := values["color"]; ok {
			result.Color = val.(string)
		}

		results = append(results, result)
	}

	return results
}

func flattenTags(input []*models.NestedTag) []interface{} {
	if input == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)

	for _, item := range input {
		values := make(map[string]interface{})

		values["id"] = item.ID
		values["name"] = item.Name
		values["slug"] = item.Slug
		values["color"] = item.Color

		result = append(result, values)
	}

	return result
}
