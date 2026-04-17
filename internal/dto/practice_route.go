package dto

type ListRoutesRequest struct {
	Category string `form:"category" binding:"omitempty"`
}

type RouteResponse struct {
	ID              int64            `json:"id"`
	Name            string           `json:"name"`
	Title           string           `json:"title"`
	Icon            string           `json:"icon"`
	Color           string           `json:"color"`
	Description     string           `json:"description"`
	TargetLevel     string           `json:"targetLevel"`
	SuitableFor     []string         `json:"suitableFor"`
	Skills          []string         `json:"skills"`
	InterviewWeight string           `json:"interviewWeight"`
	Phases          []PhaseResponse  `json:"phases"`
}

type PhaseResponse struct {
	Phase       string   `json:"phase"`
	Duration    string   `json:"duration"`
	Topics      []string `json:"topics"`
	Description string   `json:"description"`
}

type GetRouteRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type ListRoutesResponse struct {
	Routes []RouteResponse `json:"routes"`
	Total  int64           `json:"total"`
}
