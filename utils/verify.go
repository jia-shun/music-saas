package utils

var (
	IdVerify             = Rules{"ID": {NotEmpty()}}
	LoginVerify          = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify       = Rules{"Username": {NotEmpty()}, "Email": {NotEmpty()}, "Password": {NotEmpty()}}
	PageInfoVerify       = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	MusicVerify          = Rules{"SongName": {NotEmpty()}, "CustomerName": {NotEmpty()}, "Price": {NotEmpty()}, "BeganAt": {NotEmpty()}, "FinishedAt": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
)
