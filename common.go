package http_http2

var Assets []string

const (
	StaticDir = "/static/"
)
const HTML = `<!doctype html>
<html lang="en">
<head>
    <link href="/static/favicon.svg" rel="icon"  type="image/x-icon"/>
    <link rel="stylesheet" href="/static/bootstrap.min.css">
    <style>
        img {width: 30px;}
    </style>
</head>
<body class="bg-dark text-white">
	<div class="container-fluid">
		<h1>Hello HTTP and HTTP2</h1>

		<img src="/static/gopher001.png"><img src="/static/gopher002.png"><img src="/static/gopher003.png"><img src="/static/gopher004.png">
		<img src="/static/gopher005.png"><img src="/static/gopher006.png"><img src="/static/gopher007.png"><img src="/static/gopher008.png">
		<img src="/static/gopher009.png"><img src="/static/gopher010.png"><img src="/static/gopher011.png"><img src="/static/gopher012.png">
		<img src="/static/gopher013.png"><img src="/static/gopher014.png"><img src="/static/gopher015.png"><img src="/static/gopher016.png">
		<img src="/static/gopher017.png"><img src="/static/gopher018.png"><img src="/static/gopher019.png"><img src="/static/gopher020.png">
		<img src="/static/gopher021.png"><img src="/static/gopher022.png"><img src="/static/gopher023.png"><img src="/static/gopher024.png">
		<img src="/static/gopher025.png"><img src="/static/gopher026.png"><img src="/static/gopher027.png"><img src="/static/gopher028.png">
		<img src="/static/gopher029.png"><img src="/static/gopher030.png"><img src="/static/gopher031.png"><img src="/static/gopher032.png">
		<img src="/static/gopher033.png"><img src="/static/gopher034.png"><img src="/static/gopher035.png"><img src="/static/gopher036.png">
		<img src="/static/gopher037.png"><img src="/static/gopher038.png"><img src="/static/gopher039.png"><img src="/static/gopher040.png">
		<img src="/static/gopher041.png"><img src="/static/gopher042.png"><img src="/static/gopher043.png"><img src="/static/gopher044.png">
		<img src="/static/gopher045.png"><img src="/static/gopher046.png"><img src="/static/gopher047.png"><img src="/static/gopher048.png">
		<img src="/static/gopher049.png"><img src="/static/gopher050.png"><img src="/static/gopher051.png"><img src="/static/gopher052.png">
		<img src="/static/gopher053.png"><img src="/static/gopher054.png"><img src="/static/gopher055.png"><img src="/static/gopher056.png">
		<img src="/static/gopher057.png"><img src="/static/gopher058.png"><img src="/static/gopher059.png"><img src="/static/gopher060.png">
		<img src="/static/gopher061.png"><img src="/static/gopher062.png"><img src="/static/gopher063.png"><img src="/static/gopher064.png">
		<img src="/static/gopher065.png"><img src="/static/gopher066.png"><img src="/static/gopher067.png"><img src="/static/gopher068.png">
		<img src="/static/gopher069.png"><img src="/static/gopher070.png"><img src="/static/gopher071.png"><img src="/static/gopher072.png">
		<img src="/static/gopher073.png"><img src="/static/gopher074.png"><img src="/static/gopher075.png"><img src="/static/gopher076.png">
		<img src="/static/gopher077.png"><img src="/static/gopher078.png"><img src="/static/gopher079.png"><img src="/static/gopher080.png">
		<img src="/static/gopher081.png"><img src="/static/gopher082.png"><img src="/static/gopher083.png"><img src="/static/gopher084.png">
		<img src="/static/gopher085.png"><img src="/static/gopher086.png"><img src="/static/gopher087.png"><img src="/static/gopher088.png">
		<img src="/static/gopher089.png"><img src="/static/gopher090.png"><img src="/static/gopher091.png"><img src="/static/gopher092.png">
		<img src="/static/gopher093.png"><img src="/static/gopher094.png"><img src="/static/gopher095.png"><img src="/static/gopher096.png">
		<img src="/static/gopher097.png"><img src="/static/gopher098.png"><img src="/static/gopher099.png"><img src="/static/gopher100.png">
		<img src="/static/gopher101.png"><img src="/static/gopher102.png"><img src="/static/gopher103.png"><img src="/static/gopher104.png">
		<img src="/static/gopher105.png"><img src="/static/gopher106.png"><img src="/static/gopher107.png"><img src="/static/gopher108.png">
		<img src="/static/gopher109.png"><img src="/static/gopher110.png"><img src="/static/gopher111.png"><img src="/static/gopher112.png">
		<img src="/static/gopher113.png"><img src="/static/gopher114.png"><img src="/static/gopher115.png"><img src="/static/gopher116.png">
		<img src="/static/gopher117.png"><img src="/static/gopher118.png"><img src="/static/gopher119.png"><img src="/static/gopher120.png">
		<img src="/static/gopher121.png"><img src="/static/gopher122.png"><img src="/static/gopher123.png"><img src="/static/gopher124.png">
		<img src="/static/gopher125.png"><img src="/static/gopher126.png"><img src="/static/gopher127.png"><img src="/static/gopher128.png">
		<img src="/static/gopher129.png"><img src="/static/gopher130.png"><img src="/static/gopher131.png"><img src="/static/gopher132.png">
		<img src="/static/gopher133.png"><img src="/static/gopher134.png"><img src="/static/gopher135.png"><img src="/static/gopher136.png">
		<img src="/static/gopher137.png"><img src="/static/gopher138.png"><img src="/static/gopher139.png"><img src="/static/gopher140.png">
		<img src="/static/gopher141.png"><img src="/static/gopher142.png"><img src="/static/gopher143.png"><img src="/static/gopher144.png">
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
		"/static/gopher007.png",
		"/static/gopher008.png",
		"/static/gopher009.png",
		"/static/gopher010.png",
		"/static/gopher011.png",
		"/static/gopher012.png",
		"/static/gopher013.png",
		"/static/gopher014.png",
		"/static/gopher015.png",
		"/static/gopher016.png",
		"/static/gopher017.png",
		"/static/gopher018.png",
		"/static/gopher019.png",
		"/static/gopher020.png",
		"/static/gopher021.png",
		"/static/gopher022.png",
		"/static/gopher023.png",
		"/static/gopher024.png",
		"/static/gopher025.png",
		"/static/gopher026.png",
		"/static/gopher027.png",
		"/static/gopher028.png",
		"/static/gopher029.png",
		"/static/gopher030.png",
		"/static/gopher031.png",
		"/static/gopher032.png",
		"/static/gopher033.png",
		"/static/gopher034.png",
		"/static/gopher035.png",
		"/static/gopher036.png",
		"/static/gopher037.png",
		"/static/gopher038.png",
		"/static/gopher039.png",
		"/static/gopher040.png",
		"/static/gopher041.png",
		"/static/gopher042.png",
		"/static/gopher043.png",
		"/static/gopher044.png",
		"/static/gopher045.png",
		"/static/gopher046.png",
		"/static/gopher047.png",
		"/static/gopher048.png",
		"/static/gopher049.png",
		"/static/gopher050.png",
		"/static/gopher051.png",
		"/static/gopher052.png",
		"/static/gopher053.png",
		"/static/gopher054.png",
		"/static/gopher055.png",
		"/static/gopher056.png",
		"/static/gopher057.png",
		"/static/gopher058.png",
		"/static/gopher059.png",
		"/static/gopher060.png",
		"/static/gopher061.png",
		"/static/gopher062.png",
		"/static/gopher063.png",
		"/static/gopher064.png",
		"/static/gopher065.png",
		"/static/gopher066.png",
		"/static/gopher067.png",
		"/static/gopher068.png",
		"/static/gopher069.png",
		"/static/gopher070.png",
		"/static/gopher071.png",
		"/static/gopher072.png",
		"/static/gopher073.png",
		"/static/gopher074.png",
		"/static/gopher075.png",
		"/static/gopher076.png",
		"/static/gopher077.png",
		"/static/gopher078.png",
		"/static/gopher079.png",
		"/static/gopher080.png",
		"/static/gopher081.png",
		"/static/gopher082.png",
		"/static/gopher083.png",
		"/static/gopher084.png",
		"/static/gopher085.png",
		"/static/gopher086.png",
		"/static/gopher087.png",
		"/static/gopher088.png",
		"/static/gopher089.png",
		"/static/gopher090.png",
		"/static/gopher091.png",
		"/static/gopher092.png",
		"/static/gopher093.png",
		"/static/gopher094.png",
		"/static/gopher095.png",
		"/static/gopher096.png",
		"/static/gopher097.png",
		"/static/gopher098.png",
		"/static/gopher099.png",
		"/static/gopher100.png",
		"/static/gopher101.png",
		"/static/gopher102.png",
		"/static/gopher103.png",
		"/static/gopher104.png",
		"/static/gopher105.png",
		"/static/gopher106.png",
		"/static/gopher107.png",
		"/static/gopher108.png",
		"/static/gopher109.png",
		"/static/gopher110.png",
		"/static/gopher111.png",
		"/static/gopher112.png",
		"/static/gopher113.png",
		"/static/gopher114.png",
		"/static/gopher115.png",
		"/static/gopher116.png",
		"/static/gopher117.png",
		"/static/gopher118.png",
		"/static/gopher119.png",
		"/static/gopher120.png",
		"/static/gopher121.png",
		"/static/gopher122.png",
		"/static/gopher123.png",
		"/static/gopher124.png",
		"/static/gopher125.png",
		"/static/gopher126.png",
		"/static/gopher127.png",
		"/static/gopher128.png",
		"/static/gopher129.png",
		"/static/gopher130.png",
		"/static/gopher131.png",
		"/static/gopher132.png",
		"/static/gopher133.png",
		"/static/gopher134.png",
		"/static/gopher135.png",
		"/static/gopher136.png",
		"/static/gopher137.png",
		"/static/gopher138.png",
		"/static/gopher139.png",
		"/static/gopher140.png",
		"/static/gopher141.png",
		"/static/gopher142.png",
		"/static/gopher143.png",
		"/static/gopher144.png",
	}
}
