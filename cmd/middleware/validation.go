package middleware

// todo redo
// func ValidateRelayConfig(handler config.RelayConfigHandler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if err := toml.ValidateRelayConfig(handler.GetRelayConfig()); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	})
// }
