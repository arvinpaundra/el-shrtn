package request

type (
	CreateLink struct {
		OriginLink string `json:"origin_link" validate:"required,url"`
		Duration   string `json:"expired_at" validate:"required,oneof=30m 1h 12h 1d 5d 30d -"`
	}
)
