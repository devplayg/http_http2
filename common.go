package http_http2

var Assets []string

const (
	StaticDir = "/static/"
)
const HTML = `<!doctype html>
<html lang="en">
<head>
    <link rel="stylesheet" href="/static/bootstrap.min.css">
    <style>
        img {width: 30px;}
    </style>
</head>
<body class="bg-dark text-white">
	<div class="container-fluid">
		<h1>Hello HTTP and HTTP2</h1>

		<img src="/static/gopher001.png"><img src="/static/gopher002.png">
		<img src="/static/gopher003.png"><img src="/static/gopher004.png">
		<img src="/static/gopher005.png"><img src="/static/gopher006.png">
	</div>
	
	<script src="/static/jquery-3.4.1.slim.min.js"></script>
	<script src="/static/popper.min.js"></script>
	<script src="/static/bootstrap.min.js"></script>

</body>
</html>`

func init() {
	Assets = []string{
		"/static/bootstrap.min.css",
		"/static/jquery-3.4.1.slim.min.js",
		"/static/popper.min.js",
		"/static/bootstrap.min.js",
		"/static/gopher001.png",
		"/static/gopher002.png",
		"/static/gopher003.png",
		"/static/gopher004.png",
		"/static/gopher005.png",
		"/static/gopher006.png",
	}
}
