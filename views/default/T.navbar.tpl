{{define "navbar"}}

  
  <a class="navbar-brand" href="/">ShinichR 博客</a>
    <ul class="nav navbar-nav">
        <li {{if .IsHome}} class="active" {{end}}><a href="/">首页</a></li>
        <li {{if .IsCategory}} class="active" {{end}}><a href="/category">分类</a></li>
        <li {{if .IsTopic}}  class="active" {{end}}><a href="/topic">文章</a></li>
     </ul>
 
<div class ="pull-right">
    <ul class ="nav navbar-nav">
        {{if .IsLogin}}
          <li><a href="/login">退出</a>
        {{else}}
          <li><a href="/login">登录</a>
        {{end}}
    </ul>
</div>
{{end}}

