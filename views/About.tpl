{{template "base/base.html" .}}
{{define "meta"}}
    <meta name="description" content="Beego Web is official blog and documentation website for beego app web framework" />
    <meta name="keywords" content="Go, golang, beego, API documentation, blog, app web framework">
    <title>{{i18n .Lang "homepage"}} - WeFun: {{i18n .Lang "app_intro"}}</title>
{{end}}
{{define "head"}}{{end}}
{{define "body"}}
<div class="main-header">
About wefun organization
</br>
We are just ordinary people who loves programming and freedom. We want to make our life more and more beautiful.

</div>
{{end}}
