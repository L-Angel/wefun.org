{{template "base/base.html" .}}
{{define "meta"}}
    <meta name="description" content="Beego Web is official blog and documentation website for beego app web framework" />
    <meta name="keywords" content="Go, golang, beego, API documentation, blog, app web framework">
    <title>{{i18n .Lang "homepage"}} - WeFun: {{i18n .Lang "app_intro"}}</title>
{{end}}
{{define "head"}}{{end}}
{{define "body"}}
<div class="main-header">
	
</div>
{{ range .Blogs}}
<div class="item_list">
			<!-- 内容区left -->
			<!--
			<a href="http://www.wefun.org/md/" title="" class="list_left">
				<img class="list_pic" src="http://blog.daocloud.io/wp-content/uploads/2016/02/4afbfbedab64034f2227605babc379310b551d80.jpg-250x250.png" class="attachment-post-thumbnail wp-post-image" alt="4afbfbedab64034f2227605babc379310b551d80.jpg" alt="pic1">
			</a>
		-->
			<!-- 内容区right -->
			<div class="list_right">
				<h1 class="list_title">{{.B_Title}}</h1>
				<p>
					<span>Posted on 
						<a href="#">
							<time class="entry-date published" datetime="2016-04-15T15:41:20+00:00">{{.B_ReleaseTime}}}</time>
						</a>
					</span>
					<span>&nbsp;by <a href="#" class="entry-date">{{.B_Author}}}</a></span>
				</p>
				<div class="list_content">
					<p>{{.B_Content}}}</p>
				</div>
				<p>
					<span class="cat-links">
						Posted in <a href="http://www.wefun.org/md/category/%e5%b9%b2%e8%b4%a7/" rel="category tag">{{.B_Type}}</a>
					</span>
					<span class="tags-links">
						Tagged <a href="http://www.wefun.org/md/tags/docker-%e5%b9%b2%e8%b4%a7/" rel="tag">Docker 干货</a>
					</span>
					<span class="comments-link"><a href="http://blog.daocloud.io/docker-%e4%b8%8d%e5%86%8d%e6%98%af%e5%8d%95%e4%b8%80%e7%9a%84-docker-%e8%ae%b0-docker-1-11-0-%e7%89%88%e6%9c%ac%e5%8f%98%e6%9b%b4/#respond" title="Comment on Docker 不再是单一的 Docker——记 Docker 1.11.0 版本变更">Leave a comment</a></span>

				</p>
			</div>
		</div> <!-- end of item list -->
		{{end}}
{{end}}

