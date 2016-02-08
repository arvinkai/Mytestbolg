    {{define "navbar"}}
    <div class="navbar">
        <div class="navbar-inner">

            <a class="brand" href="/">我的博客</a>
            <ul class="nav">
                <li {{if .IsHome}}class= "active"{{end}}><a href="/">首页</a></li>
                <li {{if .IsCategor}}class= "active"{{end}}><a href="/category">分类</a></li>
                <li {{if .IsTopic}}class= "active"{{end}}><a href="/topic">文章</a></li>
            </ul>
            <div class="pull-right">
                <ul class="nav ">
                    {{if .IsLogin}}
                    <button  class="btn" onclick="return exitAccount()">退出</button>
                    <script type="text/javascript">
                    function exitAccount(){
                        window.location.href = "/exit";
                        return true;
                    }
                    </script>
                    {{else}}
                        <li><a href="/login">管理员登录</a></li>
                    {{end}}
                </ul>
            </div>
        </div>
    </div>
    {{end}}
