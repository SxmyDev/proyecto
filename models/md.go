package models

type Proveedor struct {
    IDProveedor  int    `gorm:"primaryKey;column:IDPROVEEDOR"`
    RUC          string `gorm:"size:11;column:RUC"`
    RazonSocial  string `gorm:"size:255;column:RAZONSOCIAL"`
    Descripcion  string `gorm:"size:255;column:DESCRIPCION"`
    Email        string `gorm:"size:50;column:EMAIL"`
    Password     string `gorm:"size:50;column:PASSWORD"`
}

func (Proveedor) TableName() string {
    return "PROVEEDOR"
}
