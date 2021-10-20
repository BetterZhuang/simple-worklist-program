/**
    @author: zzg
    @date: 2021/10/20 19:51
    @dir_path: Models
    @note:
**/

package Models

import "bubble/dao"

//Todo Model
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

// AddTodo 增加todo
func AddTodo(todo *Todo)(err error)  {
	err = dao.DB.Create(&todo).Error
	//响应
	if err != nil {
		return err
	}
	return
}

// QueryAllTodo 查询list
func QueryAllTodo()(todoList []*Todo, err error)  {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return todoList,nil
}

// QueryByID 查询byID
func QueryByID(id string)(todo *Todo, err error)  {
	todo = new(Todo)
	err = dao.DB.Where("id=?",id).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo,nil
}

// SaveModify 保存更新
func SaveModify(todo *Todo)(err error)  {
	err = dao.DB.Save(&todo).Error
	return
}

// DeleteTodo 删除
func DeleteTodo(id string)(err error)  {
	err = dao.DB.Where("id=?",id).Delete(Todo{}).Error  //针对带有DeletedAt会执行软删除
	//err = dao.DB.Where("id=?",id).Unscoped().Delete(Todo{}).Error //硬删除
	return
}