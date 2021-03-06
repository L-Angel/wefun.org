{{ template "base/base.html" .}}
	{{define "meta"}}
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>{{i18n .Lang "homepage"}} - WeFun: {{i18n .Lang "app_intro"}}</title>
	{{end}}
	{{define  "head"}}
	<link rel="stylesheet" href="http://apps.bdimg.com/libs/fontawesome/4.4.0/css/font-awesome.min.css">
	<link rel="stylesheet" href="/static_source/css/loginregister.css">
	{{end}}
	{{ define "body"}}
	<header></header>
	<div class="login-wrap">
		<div class="login-part">
			<h3>账号登陆</h3>
			<div class="login-info">
				<form onsubmit="retrun false;" class="login_form">
					<div class="form-group">
						<input type="text" name="username" placeholder="用户名/邮箱/手机号"  autocomplete="off">
						<i class="fa fa-user"></i>
						<div class="error-tip username-error">
							<i class="triangleup"></i>
							<i class="fa fa-info-circle"></i>
							<span></span>
						</div>
					</div>
					<div class="form-group">
						<input type="password" name="password" placeholder="请输入密码"  autocomplete="off">
						<i class="fa fa-lock"></i>

						<div class="error-tip password-error">
							<i class="triangleup"></i>
							<i class="fa fa-info-circle"></i>
							<span></span>
						</div>
					</div>
					<div class="form-group">
						<input type="text" name="captcha" class="captcha-input" placeholder="请输入验证码"  autocomplete="off">
						<i class="fa fa-lock"></i>

						<div class="error-tip captcha-error">
							<i class="triangleup"></i>
							<i class="fa fa-info-circle"></i>
							<span></span>
						</div>
						{{create_captcha}}
					</div>
					<div class="login-forget-password">
						<span>
							<input type="checkbox" id="rememberMe">
							<label for="rememberMe">下次自动登录</label>
						</span>
						<a href="javascript:;">忘记密码</a>
					</div>
					<input class="login-btn" value="登 录" type="submit" >
				</form>

			</div>
			<div class="split-line"></div>
			<div class="login-third-part">
				<span>第三方帐号登录</span>
				<a href="javascript:;" class="sina" target="_blank"></a>
				<a href="javascript:;" class="baidu" target="_blank"></a>
				<a href="javascript:;" class="qq" target="_blank"></a>
				<a href="javascript::" class="github" target="_blank"></a>

				<div class="login-register"><span>还没有帐号？<a href="javascript:;" target="_blank">立即注册</a></span></div>
			</div>
		</div>
	</div>
	<script src="/static_source/js/jquery.min.js"></script>
	<script src="/static_source/js/login.js"></script>
	{{end}}
