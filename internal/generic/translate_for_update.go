package generic

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func translateForUpdate(remoteString, id string, attribute map[string]schema.Attribute, ccToTf map[string]string) (string, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(remoteString), &result)
	if err != nil {
		return "", err
	}
	result["ID"] = id
	err = translateMap(result, attribute, ccToTf)
	if err != nil {
		return "", err
	}
	marshal, err := json.Marshal(result)
	return string(marshal), err

}

func translateMap(data interface{}, attribute map[string]schema.Attribute, ccToTf map[string]string) error {
	switch v := data.(type) {
	case map[string]interface{}: // JSON 对象
		for key, val := range v {
			var ccKey string
			var ok bool
			if ccKey, ok = ccToTf[key]; !ok {
				delete(v, key)
				continue
			}
			if attributeValue, ok := attribute[ccKey]; ok {
				//if readonly attribute delete
				if attributeValue.IsComputed() && !attributeValue.IsRequired() && !attributeValue.IsOptional() {
					delete(v, key)
					continue
				}
				switch attributeValue.GetType().(type) {
				case types.ListType:
					convert, ok := attributeValue.(schema.ListNestedAttribute)
					if ok {
						translateMap(val, convert.NestedObject.Attributes, ccToTf)
					}
				case types.SetType:
					convert, ok := attributeValue.(schema.SetNestedAttribute)
					if ok {
						translateMap(val, convert.NestedObject.Attributes, ccToTf)
					}
				case types.MapType:
					convert, ok := attributeValue.(schema.MapNestedAttribute)
					if ok {
						translateMap(val, convert.NestedObject.Attributes, ccToTf)
					}
				case types.ObjectType:
					convert, ok := attributeValue.(schema.SingleNestedAttribute)
					if ok {
						translateMap(val, convert.Attributes, ccToTf)
					}
				default:
					continue
				}
			} else {
				delete(v, key)
				continue
			}

		}
	case []interface{}: // JSON 数组
		for _, val := range v {
			translateMap(val, attribute, ccToTf)
		}
	default: // 基本类型（string、float64、bool、nil）
	}
	return nil
}
