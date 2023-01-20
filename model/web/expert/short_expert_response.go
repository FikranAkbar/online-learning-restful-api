package expert

type ShortExpertResponse struct {
	ID    uint   `json:"id,omitempty" mapstructure:"id"`
	Name  string `json:"name,omitempty" mapstructure:"name"`
	Photo string `json:"photo,omitempty" mapstructure:"photo"`
}
