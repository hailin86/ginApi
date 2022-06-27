package model

import "time"

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Status int `json:"status"`
	//...
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`

}


//增
func AddProduct(model *Product) (int,error)  {
	err := DB.Create(&model).Error
	if err != nil {
		return 0,err
	}
	return model.Id,nil
}

//批量添加
func BatchAddProduct(ml []*Product) (err error)  {
	err = DB.Create(&ml).Error
	return
}

//删
func DelProduct(model *Product) error  {
	err := DB.Delete(model).Error
	return err
}

//改
func UpdateProduct (model *Product) error {
	err := DB.Save(model).Error
	return err
}


//查
//通过主键查询
func GetProductById(id int) (model *Product,err error)  {
	err = DB.First(&model,id).Error
	if err != nil {
		return nil,err
	}
	return model,nil
}

//通过普通字段查询
func GetProductByName(name string) (model *Product,err error)  {
	err = DB.Where("name = ?",name).First(&model).Error
	if err != nil {
		return nil,err
	}
	return model,nil
}

//列表查询
func ListProductByStatus(status int) (ml []*Product,err error) {
	err = DB.Where("status = ?",status).Find(&ml).Error
	if err != nil {
		return nil,err
	}
	return ml,nil
}


type MyProduct struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
//模拟选中字段查询 //比如我要查询的列表只想有 id name 字段
func ListProductByIds(ids []int) (ml []*MyProduct,err error) {
	//下面这两种写法都可以
	err = DB.Table("product").Select([]string{"id","name"}).Where("id in ?",ids).Find(&ml).Error
	//err = DB.Model(Product{}).Select([]string{"id","name"}).Where("id in ?",ids).Find(&ml).Error
	if err != nil {
		return nil,err
	}
	return ml,nil
}


//分页查询
func ListProductByPage(page,limit int)(ml []*Product,err error)  {
	offset := (page -1) *limit
	err = DB.Limit(limit).Offset(offset).Order("id desc").Find(&ml).Error
	if err != nil {
		return nil,err
	}
	return ml,nil
}

