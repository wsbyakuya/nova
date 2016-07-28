package report

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"text/template"
)

func (r *Reporter) ExportHTML(open bool) {
	for i, _ := range r.Msgs {
		r.Msgs[i].Body = formatJSON(r.Msgs[i].Body)
	}
	r.ReportText = r.Report()
	report, err := template.New("report").Parse(HTMLTPL)
	if err != nil {
		panic(err)
	}
	f, fileErr := os.Create("report.html")
	if fileErr != nil {
		panic(fileErr)
	}
	defer f.Close()
	err = report.Execute(f, r)
	if err != nil {
		panic(err)
	}

	if open {
		defer openHTML()
	}
}

func openHTML() {
	cmd := exec.Command("C:/Program Files (x86)/Google/Chrome/Application/chrome.exe")
	cmd.Args = append(cmd.Args, "report.html")
	cmd.Run()
}

func formatJSON(s string) string {
	buf := []byte{}
	dst := bytes.NewBuffer(buf)
	json.Indent(dst, []byte(s), "", "  ")
	buf = dst.Bytes()
	return string(buf)
}

const HTMLTPL = `<!doctype html>
<HTML>
<HEAD>
<meta charset="utf-8" />
<style type="text/css">

.container {
	width: 800px;	
	text-align:center;
	margin:0 auto;	
}
.item {
	padding-top: 10px;
	margin-bottom: 20px;
	padding-left: 20px;	
	border:solid 1px #4CAF50;
	text-align:middle;
	overflow: auto;
	display: block;
}
.top{
	margin-top: 5px;
	margin-bottom: 10px;
	overflow: auto;
}
.middle{
	overflow: auto;
}
.content{
	text-align:left;
	margin-top: 20px;
	margin-bottom: 15px;
}
.url{
	float: left;
	width: 80%;
	text-align:left;	
}
.statusCode{
	float: right;
	width: 20%;
	color: #F44336;
	font-size: larger;
}
.pass{
	float: left;
	width: 39%;
	text-align:left;
    color: #009688;
    font-size: x-large;
}
.created{
	float: left;	
	width: 30%;
}
.time{
	float: right;
	width: 20%;
}
</style>
<script type="text/javascript">
	function spread(event){
		var thisDiv = event.target;
		nextNode = thisDiv.nextElementSibling;
		nextNode.style.display="block";
		thisDiv.style.display="none";
	}
	function shrink(event){
		var thisDiv = event.target;
		parentNode = thisDiv.parentNode;
		previousNode = parentNode.previousElementSibling;
		parentNode.style.display="none";
		previousNode.style.display="block";
	}
</script>
</HEAD>
<BODY>
<div class="container">
<pre>{{.ReportText}}</pre>
{{$isSpread := .IsSpread}}
{{range .Msgs}}
	<div class="item">
		<div class="top">
			<div class="url">{{.Url}}</div>
			<div class="statusCode">{{.StatusCode}}</div>
		</div>
		<div class="middle">
			<div class="pass">{{if .Pass}}Pass{{else}}Fail{{end}}</div>
			<div class="created"></div>
			<div class="time">{{.Time}} ms</div>
		</div>
		<div class="content">
			<div style="display:{{if $isSpread}}none{{else}}block{{end}}" onclick="spread(event)">Spread body&nbsp|&nbsp{{.ItemNum}} item{{if .ItemNum}}s{{end}}</div>
			<div style="display:{{if $isSpread}}block{{else}}none{{end}};" onclick="shrink(event)">
				<pre>{{.Body}}</pre>
			</div>
		</div>
	</div>
{{end}}
</div>
</BODY>
</HTML>`
