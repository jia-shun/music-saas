package utils

var (
	IdVerify       = Rules{"ID": {NotEmpty()}}
	LoginVerify    = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify = Rules{"Username": {NotEmpty()}, "Email": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	MusicVerify    = Rules{"UserId": {NotEmpty()}, "MusicName": {NotEmpty()}, "CustomerName": {NotEmpty()}, "BeganAt": {NotEmpty()}, "FinishedAt": {NotEmpty()}}
)
