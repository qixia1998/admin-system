// 菜单数据层

package dao

import (
	"admin-api-golang/api/entity"
	"admin-api-golang/common/util"
	. "admin-api-golang/pkg/db"
	"time"
)

// 根据菜单名称查询
func GetSysMenuByName(menuName string) (sysMenu entity.SysMenu) {
	Db.Where("menu_name = ?", menuName).First(&sysMenu)
	return sysMenu
}

// 查询新增选项列表
func QuerySysMenuVoList() (sysMenuVo []entity.SysMenuVo) {
	Db.Table("sys_menu").Select("id, menu_name AS label, parent_id").Scan(&sysMenuVo)
	return sysMenuVo
}

// 新增菜单
func CreateSysMenu(addSysMenu entity.SysMenu) bool {
	sysMenuByName := GetSysMenuByName(addSysMenu.MenuName)
	if sysMenuByName.ID != 0 {
		return false
	}
	// 目录
	if addSysMenu.MenuType == 1 {
		sysMenu := entity.SysMenu{
			ParentId:   0,
			MenuName:   addSysMenu.MenuName,
			Icon:       addSysMenu.Icon,
			MenuType:   addSysMenu.MenuType,
			Url:        addSysMenu.Url,
			MenuStatus: addSysMenu.MenuStatus,
			Sort:       addSysMenu.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	} else if addSysMenu.MenuType == 2 {
		sysMenu := entity.SysMenu{
			ParentId:   addSysMenu.ParentId,
			MenuName:   addSysMenu.MenuName,
			Icon:       addSysMenu.Icon,
			MenuType:   addSysMenu.MenuType,
			MenuStatus: addSysMenu.MenuStatus,
			Value:      addSysMenu.Value,
			Url:        addSysMenu.Url,
			Sort:       addSysMenu.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	} else if addSysMenu.MenuType == 3 {
		sysMenu := entity.SysMenu{
			ParentId:   addSysMenu.ParentId,
			MenuName:   addSysMenu.MenuName,
			MenuType:   addSysMenu.MenuType,
			MenuStatus: addSysMenu.MenuStatus,
			Value:      addSysMenu.Value,
			Sort:       addSysMenu.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
		return true
	}
	return false
}

// 根据id查询菜单详情
func GetSysMenu(Id int) (sysMenu entity.SysMenu) {
	Db.First(&sysMenu, Id)
	return sysMenu
}

// 修改菜单
func UpdateSysMenu(menu entity.SysMenu) (sysMenu entity.SysMenu) {
	Db.First(&sysMenu, menu.ID)
	sysMenu.ParentId = menu.ParentId
	sysMenu.MenuName = menu.MenuName
	sysMenu.Icon = menu.Icon
	sysMenu.Value = menu.Value
	sysMenu.MenuType = menu.MenuType
	sysMenu.Url = menu.Url
	sysMenu.MenuStatus = menu.MenuStatus
	sysMenu.Sort = menu.Sort
	Db.Save(&sysMenu)
	return sysMenu
}
func GetSysRoleMenu(id uint) (sysRoleMenu entity.SysRoleMenu) {
	Db.Where("menu_id = ?", id).First(&sysRoleMenu)
	return sysRoleMenu
}

// 删除菜单
func DeleteSysMenu(dto entity.SysMenuIdDto) bool {
	// 菜单已分配角色不能删除
	sysRoleMenu := GetSysRoleMenu(dto.Id)
	if sysRoleMenu.MenuId > 0 {
		return false
	}
	Db.Where("parent_id = ?", dto.Id).Delete(&entity.SysMenu{})
	Db.Delete(&entity.SysMenu{}, dto.Id)
	return true
}

// 查询菜单列表
func GetSysMenuList(MenuName string, MenuStatus string) (sysMenu []*entity.SysMenu) {
	curDb := Db.Table("sys_menu").Order("sort")
	if MenuName != "" {
		curDb = curDb.Where("menu_name = ?", MenuName)
	}
	if MenuStatus != "" {
		curDb = curDb.Where("menu_status = ?", MenuStatus)
	}
	curDb.Find(&sysMenu)
	return sysMenu
}
