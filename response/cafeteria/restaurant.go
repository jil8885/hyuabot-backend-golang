package cafeteria

import "github.com/hyuabot-developers/hyuabot-backend-golang/model/cafeteria"

type RestaurantListResponse struct {
	RestaurantList []RestaurantItemResponse `json:"restaurant"`
}

type RestaurantItemResponse struct {
	RestaurantID int        `json:"id"`
	Name         string     `json:"name"`
	MenuList     []MenuList `json:"menu"`
}

type MenuList struct {
	TimeType string     `json:"time"`
	MenuList []MenuItem `json:"menu"`
}

type MenuItem struct {
	Food  string `json:"food"`
	Price string `json:"price"`
}

func CreateRestaurantListResponse(restaurantList []cafeteria.RestaurantItem) RestaurantListResponse {
	var restaurant []RestaurantItemResponse
	for _, restaurantItem := range restaurantList {
		restaurant = append(restaurant, CreateRestaurantItemResponse(restaurantItem))
	}
	return RestaurantListResponse{RestaurantList: restaurant}
}

func CreateRestaurantItemResponse(restaurantItem cafeteria.RestaurantItem) RestaurantItemResponse {
	return RestaurantItemResponse{
		RestaurantID: restaurantItem.RestaurantID,
		Name:         restaurantItem.Name,
		MenuList:     CreateMenuList(restaurantItem.MenuList),
	}
}

func CreateMenuList(menuList []cafeteria.Menu) []MenuList {
	var menu = make(map[string][]MenuItem)
	for _, menuItem := range menuList {
		_, exists := menu[menuItem.TimeType]
		if !exists {
			menu[menuItem.TimeType] = make([]MenuItem, 0)
		}
		menu[menuItem.TimeType] = append(menu[menuItem.TimeType], MenuItem{Food: menuItem.Menu, Price: menuItem.Price})
	}
	var menuListResponse = make([]MenuList, 0)
	for key, value := range menu {
		menuListResponse = append(menuListResponse, MenuList{TimeType: key, MenuList: value})
	}
	return menuListResponse
}
