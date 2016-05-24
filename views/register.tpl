{{template "base/base.html" .}}
{{define "meta"}}
	<meta charset="UTF-8">
	<title>Register Page</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
{{end}}
{{define "head"}}
	<link rel="stylesheet" href="http://apps.bdimg.com/libs/fontawesome/4.4.0/css/font-awesome.min.css">
	<link rel="stylesheet" href="/static_source/css/loginregister.css">
{{end}}
{{define "body"}}
	<div class="register-wrap">
		<div class="register-part">
			<h3>注册账号</h3>
			{{.Success}}
			<div class="register-info">
				<form action="">
					<div class="form-group">
						<input type="text" class="" placeholder="请输入用户名" required>
						<i class="fa fa-user"></i>
					</div>
					<div class="form-group">
						<input type="email" class="" placeholder="请输入邮箱" required>
						<i class="fa fa-envelope"></i>
					</div>
					<div class="form-group">
						<input type="password" class="" placeholder="请输入密码" required>
						<i class="fa fa-lock"></i>
					</div>
					<div class="form-group">
						<input type="text" name="captcha" class="captcha-input" placeholder="请输入验证码" required>
						<i class="fa fa-lock"></i>
						{{create_captcha}}
					</div>
					<input class="register-btn" value="注 册" type="submit">
				</form>
				
			</div>
			<div class="split-line"></div>
			<div class="login-third-part">
				<span>第三方帐号登录</span>
				<a href="javascript:;" class="sina" target="_blank"></a>
				<a href="javascript:;" class="baidu" target="_blank"></a>
				<a href="javascript:;" class="qq" target="_blank"></a>
				<a href="javascript::" class="github" target="_blank"></a>
           	</div>
			</div>
		</div>
	</div>
{{end}}