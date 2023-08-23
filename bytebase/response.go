package bytebase

/*
[
  {
    "code": "number",
    "content": "string",
    "status": "string",
    "title": "string"
  }
]
*/

type Response struct {
	Code int `json:"code"`
	Content string `json:"content"`
	Status string `json:"status"`
	Title string `json:"title"`
}