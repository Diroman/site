// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"lawSite/restapi/operations/tools"
	"net/http"

	interpose "github.com/carbocation/interpose/middleware"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"lawSite/restapi/operations"
	"lawSite/restapi/operations/auth"
	"lawSite/restapi/operations/post"
	"lawSite/restapi/operations/user"

	bearer "lawSite/internal/auth"
	server "lawSite/internal/server_handler"
)

//go:generate swagger generate server --target ../../lawSite --name LawSite --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.LawSiteAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.LawSiteAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.BearerAuth = func(s string) (interface{}, error) {
		email, _ := bearer.ParseToken(s)
		return email, nil
	}
	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.AuthRegisterUserHandler = auth.RegisterUserHandlerFunc(server.ServerClient.CreateUser)
	api.AuthLoginUserHandler = auth.LoginUserHandlerFunc(server.ServerClient.LoginUser)
	api.AuthIsAuthHandler = auth.IsAuthHandlerFunc(server.ServerClient.IsAuth)
	api.AuthSendResetLinkHandler = auth.SendResetLinkHandlerFunc(server.ServerClient.SendResetLink)
	api.AuthCheckResetTokenHandler = auth.CheckResetTokenHandlerFunc(server.ServerClient.CheckResetToken)
	api.AuthUpdatePasswordHandler = auth.UpdatePasswordHandlerFunc(server.ServerClient.UpdatePassword)

	api.UserGetUserInfoHandler = user.GetUserInfoHandlerFunc(server.ServerClient.GetUserInfo)
	api.UserUserPostsAddHandler = user.UserPostsAddHandlerFunc(server.ServerClient.AddPostToFavorite)
	api.UserUserPostsDeleteHandler = user.UserPostsDeleteHandlerFunc(server.ServerClient.DeletePostFromFavorite)
	api.UserGetUserPostsHandler = user.GetUserPostsHandlerFunc(server.ServerClient.GetUserRegisteredPosts)
	api.UserUpdateUserInfoHandler = user.UpdateUserInfoHandlerFunc(server.ServerClient.UpdateUserInfo)

	api.UserGetUserCreatedPostsHandler = user.GetUserCreatedPostsHandlerFunc(server.ServerClient.GetPostsCreatedByUser)

	api.PostGetNearestPostHandler = post.GetNearestPostHandlerFunc(server.ServerClient.GetNearestEvent)
	api.PostGetRecommendedPostHandler = post.GetRecommendedPostHandlerFunc(server.ServerClient.GetRecommendedEvent)
	api.PostGetPostByIDHandler = post.GetPostByIDHandlerFunc(server.ServerClient.GetPostByID)
	api.PostGetPostListHandler = post.GetPostListHandlerFunc(server.ServerClient.GetPostList)
	api.PostAddNewPostHandler = post.AddNewPostHandlerFunc(server.ServerClient.CreatePost)
	api.PostDeletePostHandler = post.DeletePostHandlerFunc(server.ServerClient.DeletePost)
	api.PostUpdatePostHandler = post.UpdatePostHandlerFunc(server.ServerClient.UpdatePost)
	api.PostUpdateLinkInPostHandler = post.UpdateLinkInPostHandlerFunc(server.ServerClient.UpdatePostYoutubeLink)
	api.PostCheckPostInFavoriteHandler = post.CheckPostInFavoriteHandlerFunc(server.ServerClient.CheckPostInFavorite)

	api.PostGetMaterialByPostIDHandler = post.GetMaterialByPostIDHandlerFunc(server.ServerClient.GetMaterialsByID)
	api.PostAddNewMaterialsForPostHandler = post.AddNewMaterialsForPostHandlerFunc(server.ServerClient.AddMaterialsToPost)
	api.PostDeletePostMaterialHandler = post.DeletePostMaterialHandlerFunc(server.ServerClient.DeleteMaterialsFromPost)

	api.PostGetPostParticipantsHandler = post.GetPostParticipantsHandlerFunc(server.ServerClient.GetPostsParticipants)

	api.ToolsSendFeedbackHandler = tools.SendFeedbackHandlerFunc(server.ServerClient.SendFeedback)

	api.AddMiddlewareFor("GET", "/api/auth/is_auth", authMiddleware)
	api.AddMiddlewareFor("GET", "/api/user/info", authMiddleware)
	api.AddMiddlewareFor("GET", "/api/user/posts", authMiddleware)
	api.AddMiddlewareFor("POST", "/api/user/info/update", authMiddleware)
	api.AddMiddlewareFor("POST", "/api/user/posts/add", authMiddleware)
	api.AddMiddlewareFor("POST", "/api/user/posts/delete", authMiddleware)
	api.AddMiddlewareFor("POST", "/api/user/posts/own", authMiddleware)
	api.AddMiddlewareFor("POST", "/api/post/add", authMiddleware)
	api.AddMiddlewareFor("POST", "/api/post/delete", authMiddleware)
	api.AddMiddlewareFor("POST", "/api/post/update", authMiddleware)
	api.AddMiddlewareFor("POST", "/api/post/{id}/participants", authMiddleware)

	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	logViaLogrus := interpose.NegroniLogrus()
	return logViaLogrus(handler)
}

func authMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		code, err := bearer.ParseToken(token)
		if err != nil {
			if code != nil {
				w.WriteHeader(code.(int))
			} else {
				w.WriteHeader(400)
			}
			w.Write([]byte(err.Error()))
			return
		}
		handler.ServeHTTP(w, r)
	})
}
