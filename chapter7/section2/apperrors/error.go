package apperrors

type MyAppError struct {
	// ErrCode -> レスポンスとログに表示するエラーコード
	// Message -> レスポンスに表示するエラーメッセージ
	// error -> ログに表示する生の内部エラー

	ErrCode
	Message string
	Err     error
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// errors.Is/errors.Asを使えるようにUnwrapメソッドを定義
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
