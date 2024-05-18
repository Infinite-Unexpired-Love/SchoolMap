package curd

import (
	"TGU-MAP/models"
	"gorm.io/gorm"
	"reflect"
)

type ListItemStub struct {
	s *CascadeStub
}

func NewListItemStub(db *gorm.DB) *ListItemStub {
	return &ListItemStub{s: NewCascadeStub(db, models.ListItem{})}
}

func (stub *ListItemStub) setChildren(item *models.ListItem, itemMap map[uint][]models.ListItem) {
	children := itemMap[item.ID]
	for i := range children {
		stub.setChildren(&children[i], itemMap)
	}
	item.Children = children
}

func (stub *ListItemStub) setChildrenIterative(root *models.ListItem, itemMap map[uint][]models.ListItem) {
	stack := []*models.ListItem{root}

	for len(stack) > 0 {
		// Pop
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Get children
		children := itemMap[current.ID]
		current.Children = children

		// Push children onto stack
		for i := len(children) - 1; i >= 0; i-- {
			stack = append(stack, &children[i])
		}
	}
}

// FetchData 返回树形结构数据
func (stub *ListItemStub) FetchData() (*[]models.ListItem, *models.CustomError) {
	pdata, err := stub.s.FetchData()
	if err != nil {
		return nil, err
	}
	data := reflect.ValueOf(pdata).Elem().Interface()
	items := data.([]models.ListItem)

	itemMap := make(map[uint][]models.ListItem)
	for _, item := range items {
		parentID := item.ParentID
		itemMap[parentID] = append(itemMap[parentID], item)
	}

	var rootItems []models.ListItem

	for _, item := range itemMap[0] {
		stub.setChildren(&item, itemMap)
		rootItems = append(rootItems, item)
	}
	return &rootItems, nil
}

// InsertData 插入构造好的数据并返回成功插入的节点数
func (stub *ListItemStub) InsertData(data *[]models.ListItem) (uint, *models.CustomError) {
	var insertedCount uint

	// 队列用于BFS
	queue := make([]*models.ListItem, len(*data))
	for i := range *data {
		queue[i] = &(*data)[i]
	}

	for len(queue) > 0 {
		// 处理队列中的当前层次的所有节点
		currentLevelSize := len(queue)
		for i := 0; i < currentLevelSize; i++ {
			current := queue[i]

			// 插入当前节点
			err := stub.InsertNodeByID(current, current.ParentID)
			if err != nil {
				return insertedCount, err
			}
			insertedCount++

			// 将子节点加入队列
			for j := range current.Children {
				current.Children[j].ParentID = current.ID
				queue = append(queue, &current.Children[j])
			}
		}
		// 移除已经处理完的当前层次的节点
		queue = queue[currentLevelSize:]
	}

	return insertedCount, nil
}

func (stub *ListItemStub) InsertNodeByPath(item *models.ListItem, path ...string) *models.CustomError {
	return stub.s.InsertNodeByPath(item, path...)
}

func (stub *ListItemStub) InsertNodeByID(item *models.ListItem, parentID uint) *models.CustomError {
	return stub.s.InsertNodeByID(item, parentID)
}

func (stub *ListItemStub) UpdateNodeByPath(item *models.ListItem, path ...string) *models.CustomError {
	return stub.s.UpdateNodeByPath(item, path...)
}

func (stub *ListItemStub) UpdateNodeByID(item *models.ListItem, elemID uint) *models.CustomError {
	return stub.s.UpdateNodeByID(item, elemID)
}

func (stub *ListItemStub) DeleteNodeByID(elemID uint) *models.CustomError {
	return stub.s.DeleteNodeByID(elemID)
}

func (stub *ListItemStub) DeleteNodeByPath(path ...string) *models.CustomError {
	return stub.s.DeleteNodeByPath(path...)
}
