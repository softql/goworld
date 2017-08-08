package entity

import . "github.com/xiaonanln/goworld/engine/common"

func CreateSpaceLocally(kind int) EntityID {
	return createEntity(_SPACE_ENTITY_TYPE, nil, Position{}, "", map[string]interface{}{
		_SPACE_KIND_ATTR_KEY: kind,
	}, nil, nil, ccCreate)
}

func CreateSpaceAnywhere(kind int) {
	createEntityAnywhere(_SPACE_ENTITY_TYPE, map[string]interface{}{
		_SPACE_KIND_ATTR_KEY: kind,
	})
}
